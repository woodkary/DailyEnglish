package DB

import "database/sql"

//1根据manager_id查所有team_id和team_name
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

//2根据team_id查询该团队所有的exam_id,exam_name,exam_date
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


//3根据exam_id查询ExamInfo
func SearchExamInfoByExamID(db *sql.DB, examID string) (string, error) {
	ID := examID.Atoi()

	// 查询数据库以获取考试信息
	row := db.QueryRow("SELECT exam_name FROM exam_info WHERE exam_id = ?", ID)
	if err := row.Scan(&examName); err != nil {
		return "", err
	}

