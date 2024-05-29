package db

import (
	service "DailyEnglish/utils"
	"database/sql"
	"strconv"
	"strings"
	"time"
)

// 1根据manager_id查所有team_id和team_name
func SearchTeamInfoByManagerID(db *sql.DB, managerID int) ([]string, []string, error) {
	var teamIDs []string
	var teamNames []string

	// 查询数据库以获取团队信息
	rows, err := db.Query("SELECT team_id, team_name FROM team_info WHERE manager_id = ?", managerID)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	// 遍历结果集并收集团队ID和团队名称
	for rows.Next() {
		var teamID string
		var teamName string
		if err := rows.Scan(&teamID, &teamName); err != nil {
			return nil, nil, err
		}
		teamIDs = append(teamIDs, teamID)
		teamNames = append(teamNames, teamName)
	}

	// 检查遍历过程中是否出错
	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return teamIDs, teamNames, nil
}

// ExamInfo 结构体用于存储考试信息
type ExamInfo struct {
	ExamID   int
	ExamName string
	ExamDate string
}

// 2.1根据team_id查询该团队所有的exam_id,exam_name,exam_date
func SearchExamInfoByTeamID(db *sql.DB, teamID int) ([]ExamInfo, error) {
	var examInfos []ExamInfo

	// 查询数据库以获取考试信息
	rows, err := db.Query("SELECT exam_id, exam_name, exam_date FROM exam_info WHERE team_id = ?", teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历结果集并收集考试信息
	for rows.Next() {
		var examInfo ExamInfo
		if err := rows.Scan(&examInfo.ExamID, &examInfo.ExamName, &examInfo.ExamDate); err != nil {
			return nil, err
		}
		examInfos = append(examInfos, examInfo)
	}

	// 检查遍历过程中是否出错
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return examInfos, nil
}

// 2.2 根据团队ID和日期查询考试信息
func SearchExamInfoByTeamIDAndDate(db *sql.DB, teamID int, date string) ([]ExamInfo, error) {
	var examInfos []ExamInfo

	// 查询数据库以获取考试信息
	rows, err := db.Query("SELECT exam_id, exam_name, exam_date FROM exam_info WHERE team_id = ? AND exam_date = ?", teamID, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历结果集并收集考试信息
	for rows.Next() {
		var examInfo ExamInfo
		if err := rows.Scan(&examInfo.ExamID, &examInfo.ExamName, &examInfo.ExamDate); err != nil {
			return nil, err
		}
		examInfos = append(examInfos, examInfo)
	}

	// 检查遍历过程中是否出错
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return examInfos, nil
}

// 3 根据exam_id查询exam_score数据表里的exam_score字段
func SearchExamScoreByExamID(db *sql.DB, examID int) (string, error) {
	var examScore string

	// 查询数据库以获取考试成绩
	err := db.QueryRow("SELECT exam_score FROM exam_score WHERE exam_id = ?", examID).Scan(&examScore)
	if err != nil {
		return "", err
	}

	return examScore, nil
}

// 4 根据exam_id查询exam_info数据表里的quetion_num
func SearchQuestionNumByExamID(db *sql.DB, examID int) (int, error) {
	var questionNum int

	// 查询数据库以获取题目数量
	err := db.QueryRow("SELECT question_num FROM exam_info WHERE exam_id = ?", examID).Scan(&questionNum)
	if err != nil {
		return 0, err
	}

	return questionNum, nil
}

// 5 根据exam_id和quetion_id查询quetion_statistics表里的A_num,B_num,C_num,D_num,以及使用quetion_id查询quetion_info里的quetion_answer
func SearchQuestionStatistics(db *sql.DB, examID int, questionID int) ([]int, error) {
	var A_num, B_num, C_num, D_num int = 0, 0, 0, 0
	var correctAnswer string
	// 查询题目统计信息
	err := db.QueryRow("SELECT A_num, B_num, C_num, D_num FROM quetion_statistics WHERE exam_id = ? AND question_id = ?", examID, questionID).Scan(&A_num, &B_num, &C_num, &D_num)
	if err != nil {
		return nil, err
	}

	// 查询题目答案
	err = db.QueryRow("SELECT quetion_answer FROM quetion_info WHERE question_id = ?", questionID).Scan(&correctAnswer)
	if err != nil {
		return nil, err
	}
	ans, err := strconv.Atoi(correctAnswer)
	if err != nil {
		return nil, err
	}

	// 填充字段
	questionStats := []int{ans, A_num, B_num, C_num, D_num}
	return questionStats, nil
}

// 6.1 根据team_id查team_name
func SearchTeamNameByTeamID(db *sql.DB, teamID int) (string, error) {
	var teamName string

	// 查询数据库以获取团队名称
	err := db.QueryRow("SELECT team_name FROM team_info WHERE team_id = ?", teamID).Scan(&teamName)
	if err != nil {
		return "", err
	}

	return teamName, nil
}

// 6.2 SearchExamNameByExamID 根据考试ID查询考试名称
func SearchExamNameByExamID(db *sql.DB, examID int) (string, error) {
	var examName string

	// 查询数据库以获取考试名称
	err := db.QueryRow("SELECT exam_name FROM exam_info WHERE exam_id = ?", examID).Scan(&examName)
	if err != nil {
		return "", err
	}

	return examName, nil
}

// 7 根据exam_id查询exam_info里的quetion_id字段
func SearchQuestionIDsByExamID(db *sql.DB, examID int) ([]int, error) {
	var questionIDStr string

	// 查询数据库以获取题目ID字符串
	err := db.QueryRow("SELECT question_id FROM exam_info WHERE exam_id = ?", examID).Scan(&questionIDStr)
	if err != nil {
		return nil, err
	}

	// 切割字符串以获取各个题目ID
	questionIDStrs := strings.Split(questionIDStr, "-")

	// 创建整数数组用于存储题目ID
	questionIDs := make([]int, len(questionIDStrs))

	// 将字符串转换为整数并存储到数组中
	for i, str := range questionIDStrs {
		id, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		questionIDs[i] = id
	}

	return questionIDs, nil
}

// 8 根据team_id查询user_id
func SearchUserIDByTeamID(db *sql.DB, teamID int) ([]int, error) {
	var userIDs []int

	// 查询数据库以获取用户名称
	rows, err := db.Query("SELECT user_id FROM user_team WHERE team_id = ?", teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历结果集并收集用户名称
	for rows.Next() {
		var userID int
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		userIDs = append(userIDs, userID)
	}

	// 检查遍历过程中是否出错
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userIDs, nil
}

// 8.1 根据user_id查询user_name和user_phone
func SearchUserNameAndPhoneByUserID(db *sql.DB, userID int) (string, string, string, error) {
	var userName string
	var userPhone string
	var userEmail string
	// 查询数据库以获取用户名称
	err := db.QueryRow("SELECT username, phone,email FROM user_info WHERE user_id = ?", userID).Scan(&userName, &userPhone, &userEmail)
	if err != nil {
		return "", "", "", err
	}

	return userName, userPhone, userEmail, nil
}

// 8.2 根据user_id和team_id删除user_team表里的记录
func DeleteUserTeamByUserIDAndTeamID(db *sql.DB, userID int, teamID int) error {
	_, err := db.Exec("DELETE FROM user_team WHERE user_id = ? AND team_id = ?", userID, teamID)
	if err != nil {
		return err
	}
	return nil
}

// 9 根据考试ID和团队ID和userID查询用户名，得分，进步

func SearchClosestExamByTeamIDAndExamID(db *sql.DB, teamID, userID, examID int) (string, int, int, error) {
	var username string
	var score int
	var examRank1 int
	var examRank2 int
	var delta int
	var flag int
	// 查询数据库以获取考试排名
	err := db.QueryRow("SELECT exam_rank FROM user-exam_score WHERE exam_id = ? AND user_id = ?", examID, userID).Scan(&examRank1)
	if err != nil {
		flag = 0
	}
	if err != nil {
		flag = 0
	}

	var closestExamID int

	// 查询数据库以获取最近的另一场考试的ID
	err = db.QueryRow("SELECT exam_id FROM exam_info WHERE team_id = ? AND exam_id != ? AND exam_date < (SELECT exam_date FROM exam_info WHERE exam_id = ?) ORDER BY exam_date DESC LIMIT 1", teamID, examID, examID).Scan(&closestExamID)
	if err != nil {
		flag = 0
	}

	// 查询数据库以获取考试排名
	err = db.QueryRow("SELECT exam_rank FROM user-exam_score WHERE exam_id = ? AND user_id = ?", closestExamID, userID).Scan(&examRank2)
	if err != nil {
		flag = 0
	}

	flag = 1
	if flag == 1 {
		delta = examRank1 - examRank2
	} else {
		delta = 0
	}

	db.QueryRow("SELECT username FROM user_info WHERE user_id = ? ", userID).Scan(&username)
	db.QueryRow("SELECT user_score FROM user-exam_score WHERE exam_id = ? AND user_id = ?", examID, userID).Scan(&score)

	return username, score, delta, nil
}

type ManagerInfo struct {
	ManagerID       int
	ManagerName     string
	ManagerPhone    string
	ManagerEmail    string
	ManagerPartment string
}

// 10 根据manager_id查询manager_info数据表里的manager_name,manager_phone,manager_email,manager_partment
func SearchManagerInfoByManagerID(db *sql.DB, managerID int) (ManagerInfo, error) {
	var managerInfo ManagerInfo

	// 查询数据库以获取管理员信息
	err := db.QueryRow("SELECT manager_name, manager_phone, manager_email, manager_partment FROM manager_info WHERE manager_id = ?", managerID).Scan(&managerInfo.ManagerName, &managerInfo.ManagerPhone, &managerInfo.ManagerEmail, &managerInfo.ManagerPartment)
	if err != nil {
		return ManagerInfo{}, err
	}

	return managerInfo, nil
}

// 10.2 根据teamName查teamId
func SearchTeamIDByTeamName(db *sql.DB, teamName string) (int, error) {
	var teamID int

	// 查询数据库以获取团队ID
	err := db.QueryRow("SELECT team_id FROM team_info WHERE team_name = ?", teamName).Scan(&teamID)
	if err != nil {
		return 0, err
	}

	return teamID, nil
}

// 11 根据team_id查询team_info数据表里team_name,member_num
func SearchTeamInfoByTeamID(db *sql.DB, teamID int) (string, int, error) {
	var teamName string
	var memberNum int

	// 查询数据库以获取团队信息
	err := db.QueryRow("SELECT team_name, member_num FROM team_info WHERE team_id = ?", teamID).Scan(&teamName, &memberNum)
	if err != nil {
		return "", 0, err
	}

	return teamName, memberNum, nil
}

// 12 插入考试
func InsertExamInfo(db *sql.DB, exam_name string, exam_date string, exam_clock string, question_num int, question_id string, team_id int) error {
	now := time.Now()
	exam_id := service.GenerateID(now, 123)
	stmt, err := db.Prepare("INSERT INTO exam_info(exam_id,exam_name,exam_date,exam_clock,question_num,question_id,team_id) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(exam_id, exam_name, exam_date, exam_clock, question_num, question_id, team_id)
	if err != nil {
		return err
	}
	return nil
}

// 13 查询用户打卡信息
func SearchUserpunch(db *sql.DB, userid int) (int, string, error) {
	var lastdate string
	var ispunch int

	// 查询数据库以获取信息
	err := db.QueryRow("SELECT is_punch,last_punchdate FROM user_punch WHERE user_id = ?", userid).Scan(&ispunch, &lastdate)
	if err != nil {
		return 0, "", err
	}

	return ispunch, lastdate, err
}
