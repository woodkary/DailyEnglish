package controlsql

import (
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

// 1. 通过用户名查询用户信息
func GetUserByUsername(client *redis.Client, username string) (Member, error) {
	// 查询用户信息
	// 使用 Key 格式为 "member:{username}" 进行查询
	memberInfo, err := client.HGetAll("member:" + username).Result()
	if err != nil {
		return Member{}, err
	}

	// 解析用户信息并返回
	attendanceDays, err := strconv.Atoi(memberInfo["attendance_days"])
	if err != nil {
		return Member{}, err
	}
	isAdmin := memberInfo["is_admin"] == "true"

	member := Member{
		Username:       memberInfo["username"],
		JoinDate:       memberInfo["join_date"],
		AttendanceDays: attendanceDays,
		IsAdmin:        isAdmin,
	}

	return member, nil
}

// 通过团队名查询团队信息
func GetTeamInfo(client *redis.Client, teamName string) (Team, error) {
	// 查询团队信息
	teamInfo, err := client.HGetAll("team:" + teamName).Result()
	if err != nil {
		return Team{}, err
	}

	// 解析团队信息并返回
	team := Team{
		Name:           teamName,
		ID:             0,  // Initialize with default value
		TotalMembers:   0,  // Initialize with default value
		AdminCount:     0,  // Initialize with default value
		RecentExamDate: "", // Initialize with default value
		Last7DaysAttendance: struct {
			Count int
			Rate  float64
		}{}, // Initialize with default value
		Members: nil, // Initialize with default value
	}

	// Parse integer values
	if teamInfo["id"] != "" {
		team.ID, _ = strconv.Atoi(teamInfo["id"])
	}
	if teamInfo["total_members"] != "" {
		team.TotalMembers, _ = strconv.Atoi(teamInfo["total_members"])
	}
	if teamInfo["admin_count"] != "" {
		team.AdminCount, _ = strconv.Atoi(teamInfo["admin_count"])
	}
	if teamInfo["last_7_days_attendance_count"] != "" {
		team.Last7DaysAttendance.Count, _ = strconv.Atoi(teamInfo["last_7_days_attendance_count"])
	}
	if teamInfo["last_7_days_attendance_rate"] != "" {
		team.Last7DaysAttendance.Rate, _ = strconv.ParseFloat(teamInfo["last_7_days_attendance_rate"], 64)
	}

	return team, nil
}

// 3. 通过团队名查询团队打卡信息
func GetTeamAttendanceByTeamName(client *redis.Client, teamName string) (AttendanceRecord, error) {
	// 查询团队打卡信息
	// 使用 Key 格式为 "attendance:{date}:team:{teamName}" 进行查询
	attendanceInfo, err := client.HGetAll("attendance:" + teamName).Result()
	if err != nil {
		return AttendanceRecord{}, err
	}

	// 解析团队打卡信息并返回
	attendanceCount, err := strconv.Atoi(attendanceInfo["attendance_count"])
	if err != nil {
		return AttendanceRecord{}, err
	}
	attendanceRate, err := strconv.ParseFloat(attendanceInfo["attendance_rate"], 64)
	if err != nil {
		return AttendanceRecord{}, err
	}

	attendanceRecord := AttendanceRecord{
		Date:             attendanceInfo["date"],
		TeamName:         teamName,
		AttendanceCount:  attendanceCount,
		AttendanceRate:   attendanceRate,
		MemberAttendance: make(map[string]int),
	}

	// 查询每个成员的打卡情况
	memberKeys, err := client.Keys("attendance:" + attendanceRecord.Date + ":team:" + teamName + ":member:*").Result()
	if err != nil {
		return AttendanceRecord{}, err
	}

	// 遍历每个成员并查询打卡情况
	for _, memberKey := range memberKeys {
		username := strings.TrimPrefix(memberKey, "attendance:"+attendanceRecord.Date+":team:"+teamName+":member:")
		wordCount, err := client.HGet(memberKey, "word_count").Result()
		if err != nil {
			return AttendanceRecord{}, err
		}
		wordCountInt, err := strconv.Atoi(wordCount)
		if err != nil {
			return AttendanceRecord{}, err
		}
		attendanceRecord.MemberAttendance[username] = wordCountInt
	}

	return attendanceRecord, nil
}

// 3.1 查询团队成员的打卡情况
func GetTeamMembersAttendance(client *redis.Client, teamName string) (map[string]Member, error) {
	// 获取团队成员信息的键名
	memberKeys, err := client.Keys("team:" + teamName + ":member:*").Result()
	if err != nil {
		return nil, err
	}

	// 初始化团队成员打卡情况的map
	teamMembersAttendance := make(map[string]Member)

	// 遍历每个成员的键名，获取成员的打卡情况
	for _, key := range memberKeys {
		memberData, err := client.HGetAll(key).Result()
		if err != nil {
			return nil, err
		}

		// 解析成员打卡情况
		var member Member
		member.Username = strings.TrimPrefix(key, "team:"+teamName+":member:")
		member.JoinDate = memberData["join_date"]
		member.AttendanceDays, _ = strconv.Atoi(memberData["attendance_days"])
		member.IsAdmin, _ = strconv.ParseBool(memberData["is_admin"])
		member.AttendanceRate = memberData["attendance_rate"]

		// 将成员信息存入map
		teamMembersAttendance[member.Username] = member
	}

	return teamMembersAttendance, nil
}

// 4. 通过用户名和团队名查询该用户在该团队的打卡信息
func GetUserAttendanceByTeamName(client *redis.Client, username string, teamName string) (int, error) {
	// 查询该用户在该团队的打卡信息
	// 使用 Key 格式为 "attendance:{date}:team:{teamName}:member:{username}" 进行查询
	wordCount, err := client.HGet("attendance:"+time.Now().Format("2006-01-02")+":team:"+teamName+":member:"+username, "word_count").Result()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(wordCount)
}

// 5. 通过用户名，团队名和考试名称查询该用户考试成绩
func GetUserExamScore(client *redis.Client, username string, teamName string, examName string) (int, error) {
	// 查询该用户考试成绩
	// 使用 Key 格式为 "exam:{teamName}:user:{username}" 进行查询
	score, err := client.HGet("exam:"+teamName+":user:"+username, "score").Result()
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(score)
}

// 6. 通过团队名和考试名称查询该团队该考试的成绩信息
func GetTeamExamResult(client *redis.Client, teamName string, examName string) (ExamResult, error) {
	// 查询该团队该考试的成绩信息
	// 使用 Key 格式为 "exam_result:{teamName}:{examName}" 进行查询
	examResultInfo, err := client.HGetAll("exam_result:" + teamName + ":" + examName).Result()
	if err != nil {
		return ExamResult{}, err
	}

	// 解析考试成绩信息并返回
	examResult := ExamResult{
		TeamName: teamName,
		Scores:   make(map[string]int),
		Rankings: make(map[string]int),
	}

	// 解析成员分数和排名
	for username, score := range examResultInfo {
		if strings.HasPrefix(username, "score:") {
			username = strings.TrimPrefix(username, "score:")
			examResult.Scores[username], _ = strconv.Atoi(score)
		}
		if strings.HasPrefix(username, "rank:") {
			username = strings.TrimPrefix(username, "rank:")
			examResult.Rankings[username], _ = strconv.Atoi(score)
		}
	}

	return examResult, nil
}

// 7. 通过团队名查询该团队的通知信息
func GetTeamNotifications(client *redis.Client, teamName string) ([]Notification, error) {
	// 查询该团队的通知信息
	// 使用 Key 格式为 "notifications:{teamName}" 进行查询
	notifications, err := client.ZRange("notifications:"+teamName, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	// 解析通知信息并返回
	var teamNotifications []Notification
	for _, notification := range notifications {
		parts := strings.Split(notification, "|")
		teamNotifications = append(teamNotifications, Notification{
			ID:       parts[0],
			Title:    parts[1],
			Content:  parts[2],
			Time:     parts[3],
			TeamName: teamName,
		})
	}

	return teamNotifications, nil
}

// 8. 通过考试名获取考试id，然后再通过id获取该场考试信息
func GetExamInfoByExamName(client *redis.Client, examName string) (ExamInfo, error) {
	// 查询考试id
	examIDStr, err := client.Get("exam_id:" + examName).Result()
	if err != nil {
		return ExamInfo{}, err
	}

	examID, err := strconv.Atoi(examIDStr)
	if err != nil {
		return ExamInfo{}, err
	}

	// 使用考试id查询考试信息
	// 使用 Key 格式为 "exam_info:{examID}" 进行查询
	examInfo, err := client.HGetAll("exam_info:" + strconv.Itoa(examID)).Result()
	if err != nil {
		return ExamInfo{}, err
	}

	// 解析考试信息并返回
	questionCount, err := strconv.Atoi(examInfo["question_count"])
	if err != nil {
		return ExamInfo{}, err
	}
	averageScore, err := strconv.ParseFloat(examInfo["average_score"], 64)
	if err != nil {
		return ExamInfo{}, err
	}
	passRate, err := strconv.ParseFloat(examInfo["pass_rate"], 64)
	if err != nil {
		return ExamInfo{}, err
	}

	exam := ExamInfo{
		ID:            examID,
		Name:          examInfo["name"],
		QuestionCount: questionCount,
		AverageScore:  averageScore,
		PassRate:      passRate,
		TopTen:        make(map[string]int),
		Questions:     []string{},
	}

	// 查询前十名成员信息
	topTenMembers, err := client.HGetAll("exam_info:" + strconv.Itoa(examID) + ":top_ten").Result()
	if err != nil {
		return ExamInfo{}, err
	}
	for username, score := range topTenMembers {
		scoreInt, err := strconv.Atoi(score)
		if err != nil {
			return ExamInfo{}, err
		}
		exam.TopTen[username] = scoreInt
	}

	// 查询试题内容
	questions, err := client.HGetAll("exam_info:" + strconv.Itoa(examID) + ":questions").Result()
	if err != nil {
		return ExamInfo{}, err
	}
	for _, question := range questions {
		exam.Questions = append(exam.Questions, question)
	}

	return exam, nil
}

// 9. 通过日期查询当天所有考试信息
func GetExamsByDate(client *redis.Client, date string) ([]ExamInfo, error) {
	// 查询当天所有考试信息
	// 使用 Key 格式为 "exams:{date}" 进行查询
	examIDs, err := client.SMembers("exams:" + date).Result()
	if err != nil {
		return nil, err
	}

	// 查询每场考试的信息并返回
	var exams []ExamInfo
	for _, examID := range examIDs {
		exam, err := GetExamInfoByExamName(client, examID)
		if err != nil {
			return nil, err
		}
		exams = append(exams, exam)
	}

	return exams, nil
}

// 10. 通过用户名查询该用户是否团队管理员
func IsUserTeamAdmin(client *redis.Client, username string) (bool, error) {
	// 查询用户是否团队管理员
	// 使用 Key 格式为 "admin:{username}" 进行查询
	isAdmin, err := client.Get("admin:" + username).Result()
	if err != nil {
		return false, err
	}

	return isAdmin == "true", nil
}

// 11. 通过团队名查询该团队所有管理员信息
func GetTeamAdmins(client *redis.Client, teamName string) ([]string, error) {
	// 查询该团队所有管理员信息
	// 使用 Key 格式为 "team_admins:{teamName}" 进行查询
	admins, err := client.SMembers("team_admins:" + teamName).Result()
	if err != nil {
		return nil, err
	}

	return admins, nil
}

// 12. 通过团队名查询该团队所有团队管理员申请信息
func GetTeamAdminRequests(client *redis.Client, teamName string) ([]AdminRequest, error) {
	// 查询该团队所有团队管理员申请信息
	// 使用 Key 格式为 "admin_request:{teamName}:user:*" 进行查询
	requestKeys, err := client.Keys("admin_request:" + teamName + ":user:*").Result()
	if err != nil {
		return nil, err
	}

	// 解析申请信息并返回
	var adminRequests []AdminRequest
	for _, requestKey := range requestKeys {
		requestData, err := client.HGetAll(requestKey).Result()
		if err != nil {
			return nil, err
		}
		adminRequests = append(adminRequests, AdminRequest{
			TeamName: teamName,
			Username: requestData["username"],
			Time:     requestData["time"],
			Message:  requestData["message"],
		})
	}

	return adminRequests, nil
}
