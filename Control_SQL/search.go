package controlsql

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

// 2 通过团队名查询团队信息
func GetTeamInfo(client *redis.Client, teamName string) (Team, error) {
	// 构建 Redis Key
	teamKey := "team:" + teamName

	// 从 Redis 中获取团队信息
	teamJSON, err := client.Get(teamKey).Result()
	if err != nil {
		return Team{}, err
	}

	// 解析 JSON 格式的团队信息
	var team Team
	err = json.Unmarshal([]byte(teamJSON), &team)
	if err != nil {
		return Team{}, err
	}

	// 获取团队成员信息的键名
	memberKeys, err := client.Keys("team:" + teamName + ":member:*").Result()
	if err != nil {
		return Team{}, err
	}

	// 遍历每个成员的键名，获取成员信息并添加到团队成员列表中
	for _, key := range memberKeys {
		// 从 Redis 中获取成员信息
		memberJSON, err := client.HGetAll(key).Result()
		if err != nil {
			return Team{}, err
		}

		// 解析 JSON 格式的成员信息
		var member Member
		member.Username = strings.TrimPrefix(key, "team:"+teamName+":member:")
		member.JoinDate = memberJSON["join_date"]
		member.AttendanceDays, _ = strconv.Atoi(memberJSON["attendance_days"])
		member.IsAdmin, _ = strconv.ParseBool(memberJSON["is_admin"])
		member.AttendanceRate = memberJSON["attendance_rate"]
		member.AttendanceNum, _ = strconv.Atoi(memberJSON["attendance_num"])

		// 添加成员到团队成员列表中
		team.Members = append(team.Members, member)
	}

	// 输出团队信息
	fmt.Println("团队信息：", team)

	return team, nil
}

// 2.0通过团队名查询成员列表
func GetTeamMembers(client *redis.Client, teamName string) ([]Member, error) {
	// 查询团队成员信息
	membersJSON, err := client.Get("team:" + teamName + ":members").Result()
	if err != nil {
		return nil, err
	}

	// 解析成员信息
	var members []Member
	err = json.Unmarshal([]byte(membersJSON), &members)
	if err != nil {
		return nil, err
	}

	return members, nil
}

// 2.1 根据用户名查询该用户加入的所有团队
func GetJoinedTeams(client *redis.Client, username string) ([]string, error) {
	// 从 Redis 获取指定用户名的团队信息
	val, err := client.Get(username).Result()
	if err != nil {
		return nil, err
	}

	// 解析 JSON 数据到 Myteamseam 结构体
	var userTeams Myteams
	err = json.Unmarshal([]byte(val), &userTeams)
	if err != nil {
		return nil, err
	}

	// 返回团队数组
	return userTeams.Teams, nil
}

// 2.2 【团队管理个人中心】根据用户名查询邮箱和密码

func GetUserInfoByEmailPwd(db *sql.DB, username string) (string, string, string, error) {
	// 准备查询语句
	query := "SELECT email, pwd, phone FROM user_info WHERE username = ?"
	// 执行查询操作
	row := db.QueryRow(query, username)

	// 从查询结果中获取邮箱和密码
	var email, pwd, phone string
	err := row.Scan(&email, &pwd, &phone)
	if err != nil {
		return "", "", "", fmt.Errorf("error getting user info: %v", err)
	}

	return email, pwd, phone, nil
}

// 2.3 【团队管理个人中心】根据用户名和团队名查询团队权限
func GetIsAdminByTeamAndUsername(team Team, username string) (bool, error) {
	// 遍历团队成员列表
	for _, member := range team.Members {
		// 如果用户名匹配，则返回该成员的 IsAdmin 属性
		if member.Username == username {
			return member.IsAdmin, nil
		}
	}
	// 如果未找到匹配的用户，则返回错误
	return false, fmt.Errorf("user '%s' not found in team '%s'", username, team.Name)
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

// 3.1 查询团队成员的打卡情况  （测试成功）
func GetTeamMembersAttendance1(client *redis.Client, teamName string) (map[string]Member, error) {
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
	// 打印团队成员的打卡情况
	fmt.Println("Team Members Attendance:")
	for username, member := range teamMembersAttendance {
		fmt.Printf("Username: %s\n", username)
		fmt.Printf("Join Date: %s\n", member.JoinDate)
		fmt.Printf("Attendance Days: %d\n", member.AttendanceDays)
		fmt.Printf("Is Admin: %t\n", member.IsAdmin)
		fmt.Printf("Attendance Rate: %s\n", member.AttendanceRate)
	}
	return teamMembersAttendance, nil
}

// 3.2 查询当前日期的团队成员的打卡情况  （已测试成功）
func GetTeamMembersAttendanceByDate(client *redis.Client, teamName string) (map[string]int, error) {
	// 获取当前日期
	currentDate := "2024-04-10"
	// 构建哈希键模式
	pattern := "attendance:" + currentDate + ":team:" + teamName + ":member:*"

	// 从 Redis 中获取符合模式的所有键
	keys, err := client.Keys(pattern).Result()
	if err != nil {
		return nil, err
	}

	// 初始化结果映射
	membersAttendance := make(map[string]int)

	// 遍历符合模式的键并获取相应的值
	for _, key := range keys {
		// 获取用户名
		username := strings.TrimPrefix(key, "attendance:"+currentDate+":team:"+teamName+":member:")
		// 获取打卡情况
		wordCount, err := client.HGet(key, "word_count").Int()
		if err != nil {
			return nil, err
		}
		// 将用户名及其对应的打卡情况添加到结果映射中
		membersAttendance[username] = wordCount
	}
	// 输出结果映射
	fmt.Println("Members Attendance:")
	for username, wordCount := range membersAttendance {
		fmt.Printf("%s: %d\n", username, wordCount)
	}
	return membersAttendance, nil
}

// 3.3// 根据团队名查找所有成员的打卡单词数量//返回一个 map，其中键是成员的用户名，值是对应的打卡单词数量（测试成功）
func GetTeamMembersAttendanceNum(client *redis.Client, teamName string) (map[string]int, error) {
	// 获取团队成员信息的键名
	memberKeys, err := client.Keys("team:" + teamName + ":member:*").Result()
	if err != nil {
		return nil, err
	}

	// 初始化团队成员打卡单词数量的map
	teamMembersAttendanceNum := make(map[string]int)

	// 遍历每个成员的键名，获取成员的打卡单词数量
	for _, key := range memberKeys {
		// 从Redis中获取成员的打卡单词数量
		memberAttendanceNum, err := client.HGet(key, "attendance_num").Result()
		if err != nil {
			return nil, err
		}

		// 解析成员的打卡单词数量
		attendanceNum, _ := strconv.Atoi(memberAttendanceNum)

		// 将成员的打卡单词数量存入map
		memberUsername := strings.TrimPrefix(key, "team:"+teamName+":member:")
		teamMembersAttendanceNum[memberUsername] = attendanceNum
	}
	fmt.Println("团队成员打卡单词数量：", teamMembersAttendanceNum)
	return teamMembersAttendanceNum, nil
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

// 5.1通过团队名和考试名称查询该团队该考试的成绩信息
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
		ExamName: examName,
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

// 5.2根据团队名查询该团队所有考试的考试名称，再通过考试名称查询相应考试日期、平均分、通过率，并按照考试日期排序（测试成功）
func QueryTeamExams(client *redis.Client, teamName string) ([]map[string]string, error) {
	// 查询该团队参加的所有考试名称
	examNames, err := client.Keys("exam_result:" + teamName + ":*").Result()
	if err != nil {
		return nil, err
	}

	// 查询每场考试的详细信息
	var exams []map[string]string
	for _, examKey := range examNames {
		// 提取考试名称
		examName := strings.TrimPrefix(examKey, "exam_result:"+teamName+":")
		// 查询考试详细信息
		examInfo, err := GetExamInfoByName(client, examName)
		if err != nil {
			return nil, err
		}

		// 将 ExamInfo 转换为 map[string]string
		examMap := map[string]string{
			"ID":            strconv.Itoa(examInfo.ID),
			"Date":          examInfo.Date,
			"Name":          examInfo.Name,
			"QuestionCount": strconv.Itoa(examInfo.QuestionCount),
			"AverageScore":  strconv.FormatFloat(examInfo.AverageScore, 'f', -1, 64),
			"PassRate":      strconv.FormatFloat(examInfo.PassRate, 'f', -1, 64),
		}

		// 将 TopSix 转换为 map[string]string
		topSixMap := make(map[string]string)
		for username, score := range examInfo.TopSix {
			topSixMap[username] = strconv.Itoa(score)
		}
		examMap["TopSix"] = fmt.Sprintf("%v", topSixMap)

		// 将 Questions 转换为 map[string]string
		questionsMap := make(map[string]string)
		for i, question := range examInfo.Questions {
			questionsMap[strconv.Itoa(i)] = question
		}
		examMap["Questions"] = fmt.Sprintf("%v", questionsMap)

		exams = append(exams, examMap)
	}
	// 按日期升序排序
	sort.Slice(exams, func(i, j int) bool {
		date1, _ := time.Parse("2006-01-02", exams[i]["Date"])
		date2, _ := time.Parse("2006-01-02", exams[j]["Date"])
		return date1.Before(date2)
	})
	return exams, nil
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

// 8. 通过考试名获取考试  （测试成功）
func GetExamInfoByName(client *redis.Client, examName string) (*ExamInfo, error) {
	// 查询考试ID
	examID, err := client.Get("exam_name:" + examName).Int()
	if err != nil {
		return nil, err
	}

	// 查询考试信息
	examInfo := ExamInfo{
		ID:   examID,
		Name: examName,
	}
	fmt.Println("exam_info:" + examInfo.Name)
	examInfoMap, err := client.HGetAll("exam_info:" + examInfo.Name).Result()
	fmt.Println(examInfoMap)
	if err != nil {
		return nil, err
	}

	// 解析考试信息
	examInfo.Name = examInfoMap["name"]
	examInfo.Date = examInfoMap["date"]
	examInfo.QuestionCount, _ = strconv.Atoi(examInfoMap["question_count"])
	examInfo.AverageScore, _ = strconv.ParseFloat(examInfoMap["average_score"], 64)
	examInfo.PassRate, _ = strconv.ParseFloat(examInfoMap["pass_rate"], 64)

	// 查询前六名成员信息
	topSixMap, err := client.HGetAll("exam_info:" + strconv.Itoa(examID) + ":top_six").Result()
	if err != nil {
		return nil, err
	}
	examInfo.TopSix = make(map[string]int)
	for username, scoreStr := range topSixMap {
		score, _ := strconv.Atoi(scoreStr)
		examInfo.TopSix[username] = score
	}

	// 查询试题内容
	questionsMap, err := client.HGetAll("exam_info:" + strconv.Itoa(examID) + ":questions").Result()
	if err != nil {
		return nil, err
	}
	examInfo.Questions = make([]string, len(questionsMap))
	for i, question := range questionsMap {
		index, _ := strconv.Atoi(i)
		examInfo.Questions[index] = question
	}
	return &examInfo, nil
}

// 9. 通过日期查询当天所有考试信息
func GetExamsByDate(client *redis.Client, date string) ([]ExamInfo, error) { return nil, nil }

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

// 12. 通过团队名和flag（0加入/1管理员）查询该团队申请信息

func GetTeamRequestsByFlag(client *redis.Client, teamName string, Flag string) ([]TeamRequest, error) {
	// 从 Redis 中根据团队名和标志获取数据
	val, err := client.Get(teamName).Bytes()
	if err != nil {
		return nil, err
	}

	// 解析 JSON 格式的数据
	var requests []TeamRequest
	err = json.Unmarshal(val, &requests)
	if err != nil {
		return nil, err
	}

	// 根据标志筛选出符合条件的 TeamRequest
	var filteredRequests []TeamRequest
	for _, req := range requests {
		if req.Flag == Flag {
			filteredRequests = append(filteredRequests, req)
		}
	}

	return filteredRequests, nil
}

// 13根据团队名和日期查询最近7天的打卡情况
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

// 14根据团队名，查询该团队所有成员中，打卡天数前6名的成员名，以及他们各自的打卡天数和打卡率，其中需要把 float64 的打卡率数据转化成 string
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

// 15查询考试 id 最大的考试 ExamInfo 前六名信息 TopSix
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
