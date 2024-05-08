package DB

import (
	"database/sql"
	_ "strconv"
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

// 2根据team_id查询该团队所有的exam_id,exam_name,exam_date
func SearchExamsByTeamID(db *sql.DB, teamID int) ([]int, []string, []string, error) {
	var examIDs []int
	var examNames []string
	var examDates []string

	// 查询数据库以获取团队的考试信息
	rows, err := db.Query("SELECT exam_id, exam_name, exam_date FROM exam_info WHERE team_id = ?", teamID)
	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()

	// 遍历结果集并收集考试信息
	for rows.Next() {
		var examID int
		var examName string
		var examDate string
		if err := rows.Scan(&examID, &examName, &examDate); err != nil {
			return nil, nil, nil, err
		}
		examIDs = append(examIDs, examID)
		examNames = append(examNames, examName)
		examDates = append(examDates, examDate)
	}

	// 检查遍历过程中是否出错
	if err := rows.Err(); err != nil {
		return nil, nil, nil, err
	}

	return examIDs, examNames, examDates, nil
}

type ExamDetail struct {
	ID             string       `json:"exam_id"`          // 考试ID
	Name           string       `json:"exam_name"`        // 考试名称
	UserLevels     []int        `json:"user_levels"`      // 用户等级
	QuestionDetail [][5]int     `json:"question_details"` // 考试题目详情
	UserResult     []UserResult `json:"user_result"`      // 考试参与人员得分情况
}

type UserResult struct {
	Attend   string `json:"attend"`   // 考试参与情况
	Username string `json:"username"` // 用户名
	Score    string `json:"score"`    // 得分
	FailNum  string `json:"fail_num"` // 错题数量
	Progress string `json:"progress"` // 进步分数 (相距上次)
}

// 3根据exam_id查询exam_detail
func SearchExamInfoByExamID(db *sql.DB, examID string) ExamDetail {
	return ExamDetail{}
}

// 4根据exam_id查询题目数量
func SearchQuestionNumByExamID(db *sql.DB, examID string) int {
	return 0
}

// 5根据exam_id和question_id查询题目详情
// 题目详情包括正确答案，选A人数，选B人数，选C人数，选D人数
func SearchQuestionDetailByExamIDAndQuestionID(db *sql.DB, examID string, questionID int) [5]int {
	return [5]int{}
}
