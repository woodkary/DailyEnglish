package DB

import (
	"database/sql"
	"strconv"
	"strings"
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

// 8.1 根据user_ids查询user_names和user_phones
func SearchUserNameAndPhoneByUserID(db *sql.DB, userID int) (string, string, error) {
	var userName string
	var userPhone string

	// 查询数据库以获取用户名称
	err := db.QueryRow("SELECT username, phone FROM user_info WHERE user_id = ?", userID).Scan(&userName, &userPhone)
	if err != nil {
		return "", "", err
	}

	return userName, userPhone, nil
}
