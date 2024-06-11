package db

import (
	service "DailyEnglish/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// 重写map的序列化函数
type CustomMap struct {
	Data map[string][]string
}

func (m CustomMap) MarshalJSON() ([]byte, error) {
	// 创建一个临时map来存储序列化结果
	tempMap := make(map[string]interface{})
	for k, v := range m.Data {
		// 如果value是空数组，将其转换为interface{}类型的空数组
		if v == nil {
			tempMap[k] = []string{}
		} else {
			tempMap[k] = v
		}
	}
	// 序列化临时map
	return json.Marshal(tempMap)
}

type AverageScore struct {
	Name  string    `json:"name"`  // 学生名或团队名
	Value []float64 `json:"value"` // 各题型平均分
}
type RankScore struct {
	Name string `json:"name"` // 学生名
	Data []int  `json:"data"` // 各题型排名
}

// 0插入单词
type NewWord struct {
	Word          string `json:"word"`
	Pronunciation string `json:"pronunciation"`
	PartsOfSpeech string `json:"parts_of_speech"`
	Examples      string `json:"examples"`
	Phrases       string `json:"phrases"`
	RelatedWords  string `json:"related_words"`
}

func UpdateWordID(db *sql.DB) error {
	new_id := 1
	//查询数据库中所有word_id并更新
	rows, err := db.Query("SELECT word_id FROM word")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var word_id int
		rows.Scan(&word_id)
		_, err := db.Exec("UPDATE word SET word_id = ? WHERE word_id = ?", new_id, word_id)
		if err != nil {
			return err
		}
		new_id++
	}
	return nil
}

func InsertWords(db *sql.DB, filename string) error {
	// 读取文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// 解析JSON数据
	var NewWords []NewWord
	err = json.Unmarshal(data, &NewWords)
	if err != nil {
		return err
	}

	// 插入数据到数据库
	//word_id := 101
	for _, word := range NewWords {
		//先查询该单词是否已经存在
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM word WHERE word = ?", word.Word).Scan(&count)
		if err != nil {
			return err
		}
		//如果存在更新信息
		if count > 0 {
			_, err = db.Exec("UPDATE word SET pronunciation = ?, meanings = ?, morpheme = ?, example_sentence = ?, phrases = ? WHERE word = ?", word.Pronunciation, word.PartsOfSpeech, word.RelatedWords, word.Examples, word.Phrases, word.Word)
			if err != nil {
				log.Printf("Failed to update word %s: %v", word.Word, err)
			}
			continue
		}
		// //如果不存在插入新单词
		// _, err := db.Exec(
		// 	"INSERT INTO word (word_id,word, pronunciation, meanings, morpheme, example_sentence, phrases) VALUES (?,?, ?, ?, ?, ?, ?)",
		// 	word_id,
		// 	word.Word,
		// 	word.Pronunciation,
		// 	word.PartsOfSpeech,
		// 	word.RelatedWords,
		// 	word.Examples,
		// 	word.Phrases,
		// )
		// if err != nil {
		// 	log.Printf("Failed to insert word %s: %v", word.Word, err)
		// } else {
		// 	word_id++
		// }
	}

	return nil
}
func UpdateWords(newdb *sql.DB, db *sql.DB) error {
	//查询mysql中的word表
	rows, err := db.Query("SELECT word_id,word FROM word")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var word_id int
		var word string
		if err := rows.Scan(&word_id, &word); err != nil {
			return err
		}
		//查询sqlite中的VOC_TB表
		var difficulty int
		err = newdb.QueryRow("SELECT difficulty FROM VOC_TB WHERE spelling = ?", word).Scan(&difficulty)
		if err != nil && err != sql.ErrNoRows {
			return err
		}
		if err == sql.ErrNoRows {
			fmt.Println("word: ", word, " not found in VOC_TB\n")
			continue
		}
		//更新mysql中的word表
		_, err = db.Exec("UPDATE word SET difficulty = ? WHERE word_id = ?", difficulty, word_id)
		if err != nil {
			return err
		}
	}
	return nil
}

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

// 2.3 clock duration questionnum
type Examinfo struct {
	ExamID        int
	ExamName      string
	ExamDate      string
	StartTime     string
	Duration      int
	QuestionNum   int
	ExamFullScore int
}
type ExamDTO struct {
	ExamID       int
	ExamName     string
	ExamDate     string
	ExamClock    string
	ExamDuration int
	QuestionNum  int
}

func SearchExaminfoByTeamIDAndDate222(db *sql.DB, teamID int, userID int, date string) ([]Examinfo, error) {
	var examInfos []Examinfo
	fmt.Println("teamID: ", teamID)
	fmt.Println("date: ", date)
	// 查询数据库以获取考试信息

	// 查询数据库以获取考试信息
	rows, err := db.Query("SELECT exam_id, exam_name, exam_date,exam_clock,exam_duration,question_num FROM exam_info WHERE team_id = ? AND exam_date = ?", teamID, date)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	defer rows.Close()
	count := 1

	// 遍历结果集并收集考试信息
	for rows.Next() {
		fmt.Println("count: ", count)
		count++
		var examDto ExamDTO
		if err := rows.Scan(&examDto.ExamID, &examDto.ExamName, &examDto.ExamDate, &examDto.ExamClock, &examDto.ExamDuration, &examDto.QuestionNum); err != nil {
			log.Panic(err)
			return nil, err
		}
		//如果user-exam_score存在该考试，说明用户已经参与了该考试，则不再显示考试信息
		var count int
		err = db.QueryRow("SELECT COUNT(*) FROM `user-exam_score` WHERE exam_id = ? AND user_id = ?", examDto.ExamID, userID).Scan(&count)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
		if count > 0 {
			continue
		}
		var examInfo Examinfo
		examInfo.ExamID = examDto.ExamID
		examInfo.ExamName = examDto.ExamName
		examInfo.ExamDate = examDto.ExamDate
		//分隔考试时间
		startAndEnd := strings.Split(examDto.ExamClock, "[~,-]")

		examInfo.StartTime = startAndEnd[0]
		examInfo.Duration = examDto.ExamDuration
		examInfo.QuestionNum = examDto.QuestionNum
		examInfo.ExamFullScore = examDto.QuestionNum * 5
		examInfos = append(examInfos, examInfo)
	}
	fmt.Println("examInfos: ", examInfos)

	// 检查遍历过程中是否出错
	if err := rows.Err(); err != nil {
		log.Panic(err)
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
	fmt.Println("question_id: ", questionID)
	// 查询题目统计信息
	err := db.QueryRow("SELECT A_num, B_num, C_num, D_num FROM question_statistics WHERE exam_id = ? AND question_id = ?", examID, questionID).Scan(&A_num, &B_num, &C_num, &D_num)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	// 查询题目答案
	err = db.QueryRow("SELECT question_answer FROM question_info WHERE question_id = ?", questionID).Scan(&correctAnswer)
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
	rows, err := db.Query("SELECT user_id FROM `user-team` WHERE team_id = ?", teamID)
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
	err := db.QueryRow("SELECT manager_name, phone, email, partment FROM manager_info WHERE manager_id = ?", managerID).Scan(&managerInfo.ManagerName, &managerInfo.ManagerPhone, &managerInfo.ManagerEmail, &managerInfo.ManagerPartment)
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
func InsertExamInfo(db *sql.DB, exam_name string, exam_date string, exam_clock string, question_num int, question_id string, team_id int) (int, error) {
	exam_id := service.GenerateID()
	fmt.Println(exam_id)
	stmt, err := db.Prepare("INSERT INTO exam_info(exam_id,exam_name,exam_date,exam_clock,question_num,question_id,team_id,exam_duration) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	//将考试时长用exam_clock"09:00~11:00"格式化为int型，单位为分钟
	duration := service.TimeRangeToMinutes(exam_clock)
	_, err = stmt.Exec(exam_id, exam_name, exam_date, exam_clock, question_num, question_id, team_id, duration)
	if err != nil {
		return 0, err
	}
	return int(exam_id), nil
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

// 根据团队id和名称的map查询各团队的各学生，返回map[string][]string，也返回所有学生的id数组
func SearchTeamMemberByTeamID(db *sql.DB, idAndNameMap map[int]string) (*CustomMap, []int, error) {
	teamMemberMap := make(map[string][]string)
	var studentIds []int //这里计算了所有的团队的所有学生的总和
	for teamID, teamName := range idAndNameMap {
		fmt.Println("teamID和teamName:", teamID, teamName)
		//首先查询该团队的所有学生id
		allStudentsId, err := db.Query("SELECT user_id FROM `user-team` WHERE team_id = ?", teamID)
		if err != nil {
			//如果查询不到学生id，说明该团队没有学生，则向map中添加空数组，跳过该团队
			if err == sql.ErrNoRows {
				fmt.Println("该团队没有学生", teamName)
				teamMemberMap[teamName] = []string{}
				continue
			} else {
				//其他错误则返回错误信息
				log.Panic(err)
				return nil, nil, err
			}
		}
		defer allStudentsId.Close()
		var studentNames []string
		fmt.Println("开始查询所有学生姓名")
		for allStudentsId.Next() {
			var studentId int
			//获取每一个studentId
			if err := allStudentsId.Scan(&studentId); err != nil {
				log.Panic(err)
				return nil, nil, err
			}
			studentIds = append(studentIds, studentId)
			//根据学生id查询学生姓名
			var studentName string
			err = db.QueryRow("SELECT username FROM user_info WHERE user_id = ?", studentId).Scan(&studentName)
			if err != nil {
				log.Panic(err)
				return nil, nil, err
			}
			//将学生姓名加入到学生姓名数组中
			studentNames = append(studentNames, studentName)
		}
		//将学生姓名数组加入到teamMemberMap中
		teamMemberMap[teamName] = studentNames
		fmt.Println(teamMemberMap)
	}
	// 创建CustomMap实例并填充数据
	customMap := &CustomMap{
		Data: teamMemberMap,
	}

	return customMap, studentIds, nil
}

// // 根据学生id数组，查询其中所有学生的名字，和各题型平均分数组
// func SearchStudentAverageScoresByStudentIDs(db *sql.DB, studentIds []int) ([]AverageScore, error) {
// 	var averageScores []AverageScore
// 	for _, studentId := range studentIds {
// 		var averageScore AverageScore
// 		//先查询该学生的姓名
// 		err := db.QueryRow("SELECT username FROM user_info WHERE user_id = ?", studentId).Scan(&averageScore.Name)
// 		if err != nil {
// 			return nil, err
// 		}
// 		//查询该学生各题型的平均分
// 		//遍历各题型
// 		for question_type := 1; question_type <= 6; question_type++ {

// 		}
// 	}
// }
