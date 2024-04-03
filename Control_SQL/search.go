package controlsql

import (
	"sort"
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

// 根据团队名查询该团队所有考试的考试名称、考试日期、平均分、通过率
func QueryTeamExams(client *redis.Client, teamName string) ([]map[string]string, error) {
	// 获取所有考试的键名
	keys, err := client.Keys("exam_info:*").Result()
	if err != nil {
		return nil, err
	}

	// 保存结果的切片
	var examInfos []map[string]string

	// 遍历所有考试键名，提取考试信息
	for _, key := range keys {
		examInfo, err := client.HGetAll(key).Result()
		if err != nil {
			return nil, err
		}

		// 检查考试是否属于指定团队
		if examInfo["team_name"] == teamName {
			// 转换通过率为百分比形式
			passRate, err := strconv.ParseFloat(examInfo["pass_rate"], 64)
			if err != nil {
				return nil, err
			}
			passRateStr := strconv.FormatFloat(passRate*100, 'f', 2, 64) + "%"

			// 构建考试信息的映射
			exam := map[string]string{
				"Name":         examInfo["name"],
				"Date":         examInfo["date"],
				"AverageScore": examInfo["average_score"],
				"PassRate":     passRateStr,
			}

			// 将考试信息添加到结果切片中
			examInfos = append(examInfos, exam)
		}
	}

	return examInfos, nil
}

// 7. 通过团队名查询该团队的通知信息
// 根据团队名查询flag为0的通知，并按时间排序
func QueryUnprocessedNotifications(client *redis.Client, teamName string) ([]Notification, error) {
	// 查询有序集合中的通知信息，按时间从小到大排序
	notifications, err := client.ZRangeByScore("notifications:"+teamName, redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  -1,
	}).Result()
	if err != nil {
		return nil, err
	}

	// 将通知信息解析为结构体
	var unprocessedNotifications []Notification
	for _, notificationStr := range notifications {
		notificationParts := strings.Split(notificationStr, "|")
		// 检查通知是否已处理，如果 flag 为 "0" 则未处理
		if len(notificationParts) >= 5 && notificationParts[1] == "0" {
			notification := Notification{
				ID:       notificationParts[0],
				flag:     notificationParts[1],
				Title:    notificationParts[2],
				Content:  notificationParts[3],
				Time:     notificationParts[4],
				TeamName: teamName,
			}
			unprocessedNotifications = append(unprocessedNotifications, notification)
		}
	}
	return unprocessedNotifications, nil
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
		TopSix:        make(map[string]int),
		Questions:     []string{},
	}

	// 查询前十名成员信息
	topSixMembers, err := client.HGetAll("exam_info:" + strconv.Itoa(examID) + ":top_six").Result()
	if err != nil {
		return ExamInfo{}, err
	}
	for username, score := range topSixMembers {
		scoreInt, err := strconv.Atoi(score)
		if err != nil {
			return ExamInfo{}, err
		}
		exam.TopSix[username] = scoreInt
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

// 根据团队名和日期查询最近7天的打卡情况
func QueryAttendance(client *redis.Client, teamName string, date time.Time) (map[string]string, error) {
	attendanceData := make(map[string]string)
	for i := 0; i < 7; i++ {
		// 计算查询日期
		queryDate := date.AddDate(0, 0, -i)
		attendanceKey := "attendance:" + queryDate.Format("2006-01-02") + ":team:" + teamName
		result, err := client.HGetAll(attendanceKey).Result()
		if err != nil {
			return nil, err
		}
		// 将打卡率从字符串转换为百分比形式，例如 "0.71" 转换成 "71%"
		attendanceRateStr := result["attendance_rate"]
		attendanceRate, err := strconv.ParseFloat(attendanceRateStr, 64)
		if err != nil {
			return nil, err
		}
		attendanceRatePercent := strconv.FormatFloat(attendanceRate*100, 'f', 2, 64) + "%"
		attendanceData[queryDate.Format("2006-01-02")] = result["attendance_count"] + "|" + attendanceRatePercent
	}
	return attendanceData, nil
}

// 根据团队名，查询该团队所有成员中，打卡天数前6名的成员名，以及他们各自的打卡天数和打卡率，其中需要把 float64 的打卡率数据转化成 string
func GetTopSixAttendanceMembers(client *redis.Client, teamName string) (map[string]int, map[string]string, error) {
	// 查询团队成员列表
	members, err := client.Keys("team:" + teamName + ":member:*").Result()
	if err != nil {
		return nil, nil, err
	}

	// 初始化前6名成员信息
	topSixMembers := make(map[string]int)
	topSixAttendanceRate := make(map[string]string)

	// 遍历所有成员，统计打卡天数并记录前6名信息
	for _, memberKey := range members {
		// 获取成员打卡天数
		attendanceDaysStr, err := client.HGet(memberKey, "attendance_days").Result()
		if err != nil {
			return nil, nil, err
		}

		attendanceDays, _ := strconv.Atoi(attendanceDaysStr)
		username := memberKey[len("team:"+teamName+":member:"):]
		topSixMembers[username] = attendanceDays
	}

	// 对打卡天数进行排序，获取前6名成员
	type member struct {
		Username       string
		AttendanceDays int
	}

	var sortedMembers []member
	for username, attendanceDays := range topSixMembers {
		sortedMembers = append(sortedMembers, member{username, attendanceDays})
	}
	sort.Slice(sortedMembers, func(i, j int) bool {
		return sortedMembers[i].AttendanceDays > sortedMembers[j].AttendanceDays
	})

	// 取前6名成员的打卡率并转化为字符串
	for i := 0; i < len(sortedMembers) && i < 6; i++ {
		username := sortedMembers[i].Username
		attendanceDays := sortedMembers[i].AttendanceDays
		attendanceRate := float64(attendanceDays) / 7.0
		attendanceRateStr := strconv.FormatFloat(attendanceRate, 'f', 2, 64)
		topSixAttendanceRate[username] = attendanceRateStr
	}

	return topSixMembers, topSixAttendanceRate, nil
}

// 查询考试 id 最大的考试 ExamInfo 前六名信息 TopSix
func GetTopSixExamScores(client *redis.Client) (map[string]int, error) {
	// 查询考试信息列表
	examInfos, err := client.Keys("exam_info:*").Result()
	if err != nil {
		return nil, err
	}

	// 初始化前六名信息
	topSixScores := make(map[string]int)

	// 遍历所有考试，获取每场考试的前六名成绩
	for _, examKey := range examInfos {
		// 获取考试名称
		examName, err := client.HGet(examKey, "name").Result()
		if err != nil {
			return nil, err
		}

		// 获取考试成绩列表
		examScores, err := client.HGetAll("exam:" + examName).Result()
		if err != nil {
			return nil, err
		}

		// 遍历成绩列表，记录前六名成绩
		for username, scoreStr := range examScores {
			score, _ := strconv.Atoi(scoreStr)
			if len(topSixScores) < 6 || score > topSixScores[username] {
				topSixScores[username] = score
			}
		}
	}

	return topSixScores, nil
}
