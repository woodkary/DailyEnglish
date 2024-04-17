package controlsql

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

type Team struct {
	Name           string // 团队名
	ID             int    // 团队ID
	TotalMembers   int    // 团队总人数
	AdminCount     int    // 管理员人数
	RecentExamDate string // 最近一场考试日期

	Members []Member // 团队成员列表
}

type Member struct {
	Username       string // 成员用户名
	JoinDate       string // 加入团队日期
	AttendanceDays int    // 打卡天数
	IsAdmin        bool   // 是否是团队管理员
	AttendanceRate string //打卡率
	AttendanceNum  int    //打卡单词数量
}

type AttendanceRecord struct {
	Date             string         // 日期
	TeamName         string         // 团队名
	AttendanceCount  int            // 打卡人数
	MemberAttendance map[string]int // 成员打卡情况，键为用户名，值为是否打卡，1是0否
	AttendanceRate   float64        // 打卡率
}

type TeamRequest struct {
	TeamName string // 团队名
	flag     string // 0加入申请 1管理员申请
	Username string // 申请者用户名
	Time     string // 申请时间
	Message  string // 申请留言内容
}

type Notification struct {
	ID       string // 通知ID
	flag     string //已处理标志，0未处理1已经处理
	Title    string // 通知标题
	Content  string // 通知内容
	Time     string // 通知时间
	TeamName string // 通知团队名
}

type ExamResult struct {
	TeamName string // 团队名
	ExamName string
	Scores   map[string]int // 成员分数，键为用户名，值为分数
	Rankings map[string]int // 成员排名，键为用户名，值为排名
}

type ExamInfo struct {
	ID            int            // 考试ID
	date          string         //日期
	Name          string         // 考试名称
	QuestionCount int            // 试题数量
	Questions     []string       // 试题内容
	AverageScore  float64        // 考试平均分
	PassRate      float64        // 及格率
	TopSix        map[string]int // 前6名成员用户名及分数，键为用户名，值为分数
}

func InsertData(client *redis.Client) error {
	// 保存团队信息
	team := Team{
		Name:           "每日背单词小组",
		ID:             9,
		TotalMembers:   5,
		AdminCount:     1,
		RecentExamDate: "2024-04-10",
		Members: []Member{
			{Username: "小明", JoinDate: "2024-01-01", AttendanceDays: 90, IsAdmin: true, AttendanceRate: "90%", AttendanceNum: 500},
			{Username: "小红", JoinDate: "2024-01-15", AttendanceDays: 85, IsAdmin: false, AttendanceRate: "85%", AttendanceNum: 480},
			{Username: "小蓝", JoinDate: "2024-02-01", AttendanceDays: 80, IsAdmin: false, AttendanceRate: "80%", AttendanceNum: 450},
			{Username: "张三", JoinDate: "2024-02-15", AttendanceDays: 75, IsAdmin: false, AttendanceRate: "75%", AttendanceNum: 430},
			{Username: "李四", JoinDate: "2024-03-01", AttendanceDays: 70, IsAdmin: false, AttendanceRate: "70%", AttendanceNum: 400},
		},
	}
	err := SaveTeam(client, team)
	if err != nil {
		return err
	}

	// 保存打卡信息
	attendanceRecord := AttendanceRecord{
		Date:             "2024-04-10",
		TeamName:         "每日背单词小组",
		AttendanceCount:  5,
		MemberAttendance: map[string]int{"小明": 1, "小红": 1, "小蓝": 1, "张三": 1, "李四": 0},
		AttendanceRate:   0.8,
	}
	err = SaveAttendanceRecord(client, attendanceRecord)
	if err != nil {
		return err
	}

	// 保存团队申请信息
	teamRequest := TeamRequest{
		TeamName: "每日背单词小组",
		flag:     "0",
		Username: "小橙",
		Time:     "2024-04-10",
		Message:  "I want to join the team.",
	}
	err = SaveTeamRequest(client, teamRequest)
	if err != nil {
		return err
	}
	// 保存团队申请信息
	teamRequest2 := TeamRequest{
		TeamName: "每日背单词小组",
		flag:     "0",
		Username: "小绿",
		Time:     "2023-04-10",
		Message:  "I want to join the team.",
	}
	err = SaveTeamRequest(client, teamRequest2)
	if err != nil {
		return err
	}
	// 保存团队申请信息
	teamRequest3 := TeamRequest{
		TeamName: "每日背单词小组",
		flag:     "1",
		Username: "小红",
		Time:     "2024-04-10",
		Message:  "我要当管理.",
	}
	err = SaveTeamRequest(client, teamRequest3)
	if err != nil {
		return err
	}
	// 保存团队申请信息
	teamRequest4 := TeamRequest{
		TeamName: "每日背单词小组",
		flag:     "1",
		Username: "小橙",
		Time:     "2024-04-10",
		Message:  "我也要当管理",
	}
	err = SaveTeamRequest(client, teamRequest4)
	if err != nil {
		return err
	}
	// 保存团队申请信息
	teamRequest5 := TeamRequest{
		TeamName: "每日背单词小组",
		flag:     "1",
		Username: "张三",
		Time:     "2024-05-10",
		Message:  "我不当普通成员",
	}
	err = SaveTeamRequest(client, teamRequest5)
	if err != nil {
		return err
	}
	// 保存团队申请信息
	teamRequest6 := TeamRequest{
		TeamName: "每日背单词小组",
		flag:     "0",
		Username: "小橙",
		Time:     "2024-04-10",
		Message:  "我要进部",
	}
	err = SaveTeamRequest(client, teamRequest6)
	if err != nil {
		return err
	}
	// 保存团队申请信息
	teamRequest7 := TeamRequest{
		TeamName: "每日背单词小组",
		flag:     "0",
		Username: "李四",
		Time:     "2024-09-10",
		Message:  "让我进去",
	}
	err = SaveTeamRequest(client, teamRequest7)
	if err != nil {
		return err
	}

	// 保存通知信息
	notification1 := Notification{
		ID:       "1",
		flag:     "0",
		Title:    "Notification Title",
		Content:  "This is a notification content.",
		Time:     "2024-04-10",
		TeamName: "每日背单词小组",
	}
	err = SaveNotification(client, notification1)
	if err != nil {
		return err
	}
	notification2 := Notification{
		ID:       "2",
		flag:     "0",
		Title:    "Notification Title",
		Content:  "This is a notification content.",
		Time:     "2024-04-10",
		TeamName: "每日背单词小组",
	}
	err = SaveNotification(client, notification2)
	if err != nil {
		return err
	}
	notification3 := Notification{
		ID:       "3",
		flag:     "0",
		Title:    "Notification Title",
		Content:  "This is a notification content.",
		Time:     "2024-04-10",
		TeamName: "每日背单词小组",
	}
	err = SaveNotification(client, notification3)
	if err != nil {
		return err
	}
	notification4 := Notification{
		ID:       "4",
		flag:     "0",
		Title:    "Notification Title",
		Content:  "This is a notification content.",
		Time:     "2024-04-10",
		TeamName: "每日背单词小组",
	}
	err = SaveNotification(client, notification4)
	if err != nil {
		return err
	}
	notification5 := Notification{
		ID:       "5",
		flag:     "0",
		Title:    "Notification Title",
		Content:  "This is a notification content.",
		Time:     "2024-04-10",
		TeamName: "每日背单词小组",
	}
	err = SaveNotification(client, notification5)
	if err != nil {
		return err
	}
	notification6 := Notification{
		ID:       "6",
		flag:     "0",
		Title:    "Notification Title",
		Content:  "This is a notification content.",
		Time:     "2024-04-10",
		TeamName: "每日背单词小组",
	}
	err = SaveNotification(client, notification6)
	if err != nil {
		return err
	}

	// 保存考试成绩信息
	examResult := ExamResult{
		TeamName: "每日背单词小组",
		ExamName: "Exam1",
		Scores:   map[string]int{"小明": 90, "小红": 85, "小蓝": 80, "张三": 75, "李四": 70},
		Rankings: map[string]int{"小明": 1, "小红": 2, "小蓝": 3, "张三": 4, "李四": 5},
	}
	err = SaveExamResult(client, examResult)
	if err != nil {
		return err
	}
	// 保存考试成绩信息
	examResult2 := ExamResult{
		TeamName: "每日背单词小组",
		ExamName: "Exam2",
		Scores:   map[string]int{"小明": 90, "小红": 85, "小蓝": 80, "张三": 75, "李四": 70},
		Rankings: map[string]int{"小明": 1, "小红": 2, "小蓝": 3, "张三": 4, "李四": 5},
	}
	err = SaveExamResult(client, examResult2)
	if err != nil {
		return err
	}
	// 保存考试成绩信息
	examResult3 := ExamResult{
		TeamName: "每日背单词小组",
		ExamName: "Exam3",
		Scores:   map[string]int{"小明": 90, "小红": 85, "小蓝": 80, "张三": 75, "李四": 70},
		Rankings: map[string]int{"小明": 1, "小红": 2, "小蓝": 3, "张三": 4, "李四": 5},
	}
	err = SaveExamResult(client, examResult3)
	if err != nil {
		return err
	}
	// 保存考试信息
	examInfo1 := ExamInfo{
		ID:            1,
		date:          "2024-04-10",
		Name:          "Exam1",
		QuestionCount: 5,
		Questions:     []string{"Question 1", "Question 2", "Question 3", "Question 4", "Question 5"},
		AverageScore:  80.0,
		PassRate:      0.8,
		TopSix:        map[string]int{"小明": 90, "小红": 85, "小蓝": 80, "张三": 75, "李四": 70, "小黄": 30},
	}
	examInfo2 := ExamInfo{
		ID:            2,
		date:          "2024-02-10",
		Name:          "Exam2",
		QuestionCount: 5,
		Questions:     []string{"Question 1", "Question 2", "Question 3", "Question 4", "Question 5"},
		AverageScore:  70.0,
		PassRate:      0.8,
		TopSix:        map[string]int{"小明": 90, "小红": 75, "小蓝": 70, "张三": 65, "李四": 50, "小黄": 30},
	}
	examInfo3 := ExamInfo{
		ID:            3,
		date:          "2024-03-10",
		Name:          "Exam3",
		QuestionCount: 5,
		Questions:     []string{"Question 1", "Question 2", "Question 3", "Question 4", "Question 5"},
		AverageScore:  85.0,
		PassRate:      0.8,
		TopSix:        map[string]int{"小明": 100, "小红": 90, "小蓝": 80, "张三": 75, "李四": 70, "小黄": 30},
	}
	err = SaveExamInfo(client, examInfo1)
	if err != nil {
		return err
	}
	err = SaveExamInfo(client, examInfo2)
	if err != nil {
		return err
	}
	err = SaveExamInfo(client, examInfo3)
	if err != nil {
		return err
	}

	return nil
}

// 保存每个用户所加入的若干团队
func RecordTeamJoin(redisClient *redis.Client, username string, teamNames []string) error {
	// 使用 Redis 的 Hash 数据结构存储每个用户名加入的团队名
	for _, teamName := range teamNames {
		_, err := redisClient.HSet("user_teams:"+username, teamName, 1).Result()
		if err != nil {
			return err
		}
	}
	return nil
}

// 保存团队信息
func SaveTeam(client *redis.Client, team Team) error {
	// 使用哈希数据结构保存团队信息
	_, err := client.HMSet("team:"+team.Name, map[string]interface{}{
		"id":               team.ID,
		"total_members":    team.TotalMembers,
		"admin_count":      team.AdminCount,
		"recent_exam_date": team.RecentExamDate,
	}).Result()
	if err != nil {
		return err
	}

	// 保存团队成员信息
	for _, member := range team.Members {
		_, err := client.HMSet("team:"+team.Name+":member:"+member.Username, map[string]interface{}{
			"join_date":       member.JoinDate,
			"attendance_days": strconv.Itoa(member.AttendanceDays), // Convert int to string
			"is_admin":        member.IsAdmin,
			"attendance_rate": member.AttendanceRate, // 保存成员的打卡率
			"attendance_num":  member.AttendanceNum,  // 保存成员的打卡单词数量
		}).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

// 新用户加入：
// 根据团队名和用户名，将用户加入团队中，并更新团队成员信息和成绩信息
func AddMemberToTeam(client *redis.Client, teamName string, member Member) error {
	// 构建团队成员在 Redis 中的键名
	memberKey := "team:" + teamName + ":member:" + member.Username

	// 使用哈希数据结构保存团队成员信息
	_, err := client.HMSet(memberKey, map[string]interface{}{
		"join_date":       member.JoinDate,
		"attendance_days": strconv.Itoa(member.AttendanceDays), // Convert int to string
		"is_admin":        member.IsAdmin,
	}).Result()
	if err != nil {
		return err
	}

	// 更新团队总人数
	totalMembersKey := "team:" + teamName + ":total_members"
	_, err = client.Incr(totalMembersKey).Result()
	if err != nil {
		return err
	}

	// 如果成员是管理员，更新管理员人数
	if member.IsAdmin {
		adminCountKey := "team:" + teamName + ":admin_count"
		_, err := client.Incr(adminCountKey).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

// 保存打卡信息
func SaveAttendanceRecord(client *redis.Client, record AttendanceRecord) error {
	// 使用哈希数据结构保存打卡记录
	_, err := client.HMSet("attendance:"+record.Date+":team:"+record.TeamName, map[string]interface{}{
		"attendance_count": record.AttendanceCount,
		"attendance_rate":  record.AttendanceRate,
	}).Result()
	if err != nil {
		return err
	}

	// 保存每个成员的打卡情况
	for username, wordCount := range record.MemberAttendance {
		_, err := client.HSet("attendance:"+record.Date+":team:"+record.TeamName+":member:"+username, "word_count", wordCount).Result()
		if err != nil {
			return err
		}
	}

	return nil
}

// 保存团队申请信息（0加入/1管理员申请）
func SaveTeamRequest(client *redis.Client, request TeamRequest) error {
	// 将 TeamRequest 结构体转换为 JSON 格式
	requestJSON, err := json.Marshal(request)
	if err != nil {
		return err
	}

	// 将 JSON 格式的数据追加到 Redis 列表中
	err = client.RPush(request.TeamName, requestJSON).Err()
	if err != nil {
		return err
	}

	return nil
}

// 保存通知信息
func SaveNotification(client *redis.Client, notification Notification) error {
	// 获取当前时间的 Unix 时间戳，并转换为 float64 类型
	now := float64(time.Now().Unix())

	// 使用有序集合保存通知信息，以通知ID作为分数，保证通知按时间顺序存储和检索
	_, err := client.ZAdd("notifications:"+notification.TeamName, redis.Z{
		Score:  now,
		Member: fmt.Sprintf("%s|%s|%s|%s|%s", notification.ID, notification.flag, notification.Title, notification.Content, notification.Time),
	}).Result()
	if err != nil {
		return err
	}

	return nil
}

// 通知已读
// 根据团队名和通知 ID，将通知的 flag 设为 1
func MarkNotificationAsProcessed(client *redis.Client, teamName string, notificationID string) error {
	// 构建通知在有序集合中的键名
	notificationKey := "notifications:" + teamName

	// 查询通知是否存在，并获取其分数（时间戳）
	score, err := client.ZScore(notificationKey, notificationID).Result()
	if err != nil {
		return err
	}

	// 将通知 ID 的 flag 设为 1，即已处理状态
	_, err = client.ZAdd(notificationKey, redis.Z{
		Score:  score,                               // 使用原来的分数
		Member: fmt.Sprintf("%s|1", notificationID), // 设置 flag 为 1
	}).Result()
	if err != nil {
		return err
	}

	return nil
}

// 保存考试成绩到 Redis 数据库
func SaveExamResult(client *redis.Client, examResult ExamResult) error {
	// 将 examResult 转换为 JSON 格式
	examResultJSON, err := json.Marshal(examResult)
	if err != nil {
		return err
	}

	// 将 examResultJSON 保存到 Redis 中
	err = client.Set("exam_result:"+examResult.TeamName+":"+examResult.ExamName, examResultJSON, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

// 保存考试内容
func SaveExamInfo(client *redis.Client, examInfo ExamInfo) error {
	// 使用哈希数据结构保存考试信息
	_, err := client.HMSet("exam_info:"+strconv.Itoa(examInfo.ID), map[string]interface{}{
		"date":           examInfo.date,
		"name":           examInfo.Name,
		"question_count": examInfo.QuestionCount,
		"average_score":  examInfo.AverageScore,
		"pass_rate":      examInfo.PassRate,
	}).Result()
	if err != nil {
		return err
	}

	// 保存前六名成员信息
	for username, score := range examInfo.TopSix {
		_, err := client.HSet("exam_info:"+strconv.Itoa(examInfo.ID)+":top_six", username, score).Result()
		if err != nil {
			return err
		}
	}

	// 保存试题内容
	for i, question := range examInfo.Questions {
		_, err := client.HSet("exam_info:"+strconv.Itoa(examInfo.ID)+":questions", strconv.Itoa(i), question).Result()
		if err != nil {
			return err
		}
	}

	return nil
}
