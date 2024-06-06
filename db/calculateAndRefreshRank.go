package db

import (
	"database/sql"
	"log"
	"strings"
	"time"
)

// 检查是否需要更新数据库中的排名
func CalculateRank(db *sql.DB, exam_id int) (bool, error) { //目前只有在所有人都提交了考试或者超时了才会计算排名，具体逻辑可以根据需求修改
	//先检查是否所有用户提交了考试
	//查询班级人数
	var class_num int
	//根据exam_id查询team_id,再根据team_id查询member_num
	err := db.QueryRow("SELECT member_num FROM team_info WHERE team_id = (SELECT team_id FROM exam_info WHERE exam_id = ?)", exam_id).Scan(&class_num)
	if err != nil {
		return true, err
	}
	//查询已提交人数
	var submit_num int
	err = db.QueryRow("SELECT COUNT(DISTINCT user_id) FROM `user-exam_score` WHERE exam_id = ?", exam_id).Scan(&submit_num)
	if err != nil {
		return true, err
	}
	if submit_num == class_num { //所有人都提交了可以计算排名
		return true, nil
	}
	//没有提交再检查是否超时了
	var exam_clock string
	err = db.QueryRow("SELECT exam_clock FROM exam_info WHERE exam_id = ?", exam_id).Scan(&exam_clock)
	if err != nil {
		return true, err
	}
	end_time, _ := time.Parse(strings.Split(exam_clock, "~")[1], "09:00")
	// 获取当前时间
	currentTime := time.Now()

	// 获取当前日期的时间部分
	currentTimeWithoutDate := time.Date(0, 1, 1, currentTime.Hour(), currentTime.Minute(), currentTime.Second(), 0, time.UTC)

	// 比较当前时间与目标时间
	if currentTimeWithoutDate.Before(end_time) { //当前时间在考试结束时间之前
		return false, nil
	}
	//超时了
	return true, nil
}

// 计算排名并更新数据库
func FreshRank(db *sql.DB, exam_id int) error {
	//先把该考试所有人的分数取出来
	rows, err := db.Query("SELECT user_id, exam_score FROM `user-exam_score` WHERE exam_id = ?", exam_id)
	if err != nil {
		log.Panic(err)
		return err
	}
	defer rows.Close()
	var user_score = make(map[int]int)
	for rows.Next() {
		var user_id, score int
		if err := rows.Scan(&user_id, &score); err != nil {
			return err
		}
		user_score[user_id] = score
	}
	if err := rows.Err(); err != nil {
		return err
	}
	//排序
	type kv struct {
		Key   int
		Value int
	}
	var ss []kv
	for k, v := range user_score {
		ss = append(ss, kv{k, v})
	}
	//排序
	for i := 0; i < len(ss); i++ {
		for j := i + 1; j < len(ss); j++ {
			if ss[i].Value < ss[j].Value {
				ss[i], ss[j] = ss[j], ss[i]
			}
		}
	}
	//更新排名
	for i, v := range ss {
		_, err := db.Exec("UPDATE `user-exam_score` SET rank = ? WHERE user_id = ? AND exam_id = ?", i+1, v.Key, exam_id)
		if err != nil {
			log.Panic(err)
			return err
		}
	}
	//未参考的人不参与排名
	//计算平均分，计算考试的满分，再以此计算及格率，并将用户的成绩以解析成字符串"user_id:score,user_id:score"形式存入exam_info表中
	var question_num int
	err = db.QueryRow("SELECT question_num FROM exam_info WHERE exam_id = ?", exam_id).Scan(&question_num)
	if err != nil {
		log.Panic(err)
		return err
	}
	var full_score int
	err = db.QueryRow("SELECT full_score FROM exam_info WHERE exam_id = ?", exam_id).Scan(&full_score)
	if err != nil {
		log.Panic(err)
		return err
	}
	//及格分数用总分的60%
	pass_score := int(float64(full_score) * 0.6)
	//计算平均分
	var sum_score int
	for _, v := range ss {
		sum_score += v.Value
	}
	average_score := sum_score / len(ss)
	//计算及格率
	var pass_num int
	for _, v := range ss {
		if v.Value >= pass_score {
			pass_num++
		}
	}
	pass_rate := float64(pass_num) / float64(len(ss))
	//将用户的成绩以解析成字符串"user_id:score,user_id:score"形式存入exam_info表中
	var user_score_str string
	for _, v := range ss {
		user_score_str += string(v.Key) + ":" + string(v.Value) + ","
	}
	//将得到的数据插入exam_info表中
	stmt, err := db.Prepare("INSERT INTO `exam_score`(exam_id,max_score,average_score,pass_rate,exam_score) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(exam_id, full_score, average_score, pass_rate, user_score_str)
	if err != nil {
		return err
	}
	return nil
}
