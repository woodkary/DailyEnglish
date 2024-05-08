package DB

import (
	"database/sql"
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
// QuestionStatistics 结构体用于保存题目统计信息
type QuestionStatistics struct {
	ExamID         int
	QuestionID     int
	QuestionAnswer string
	ANum           int
	BNum           int
	CNum           int
	DNum           int
}

// SearchQuestionStatistics 根据考试ID和题目ID查询题目统计信息
func SearchQuestionStatistics(db *sql.DB, examID, questionID int) (QuestionStatistics, error) {
	var questionStats QuestionStatistics

	// 查询题目统计信息
	err := db.QueryRow("SELECT A_num, B_num, C_num, D_num FROM quetion_statistics WHERE exam_id = ? AND question_id = ?", examID, questionID).Scan(&questionStats.ANum, &questionStats.BNum, &questionStats.CNum, &questionStats.DNum)
	if err != nil {
		return QuestionStatistics{}, err
	}

	// 查询题目答案
	err = db.QueryRow("SELECT quetion_answer FROM quetion_info WHERE question_id = ?", questionID).Scan(&questionStats.QuestionAnswer)
	if err != nil {
		return QuestionStatistics{}, err
	}

	// 填充其他字段
	questionStats.ExamID = examID
	questionStats.QuestionID = questionID

	return questionStats, nil
}
