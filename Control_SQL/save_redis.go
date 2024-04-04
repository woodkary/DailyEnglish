package controlsql

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

type Team struct {
	Name                string // 团队名
	ID                  int    // 团队ID
	TotalMembers        int    // 团队总人数
	AdminCount          int    // 管理员人数
	RecentExamDate      string // 最近一场考试日期
	Last7DaysAttendance struct {
		Count int     // 最近7天的打卡人数
		Rate  float64 // 最近7天的打卡率
	}
	Members []Member // 团队成员列表
}

type Member struct {
	Username       string // 成员用户名
	JoinDate       string // 加入团队日期
	AttendanceDays int    // 打卡天数
	IsAdmin        bool   // 是否是团队管理员
}

type AttendanceRecord struct {
	Date             string         // 日期
	TeamName         string         // 团队名
	AttendanceCount  int            // 打卡人数
	MemberAttendance map[string]int // 成员打卡情况，键为用户名，值为打卡单词数量
	AttendanceRate   float64        // 打卡率
}

type AdminRequest struct {
	TeamName string // 团队名
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
	TeamName string         // 团队名
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

// 保存团队信息
func SaveTeam(client *redis.Client, team Team) error {
	// 使用哈希数据结构保存团队信息
	_, err := client.HMSet("team:"+team.Name, map[string]interface{}{
		"id":                           team.ID,
		"total_members":                team.TotalMembers,
		"admin_count":                  team.AdminCount,
		"recent_exam_date":             team.RecentExamDate,
		"last_7_days_attendance_count": team.Last7DaysAttendance.Count,
		"last_7_days_attendance_rate":  team.Last7DaysAttendance.Rate,
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

// 保存团队申请信息
func SaveAdminRequest(client *redis.Client, request AdminRequest) error {
	// 使用哈希数据结构保存申请信息
	_, err := client.HMSet("admin_request:"+request.TeamName+":user:"+request.Username, map[string]interface{}{
		"time":    request.Time,
		"message": request.Message,
	}).Result()
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

// 保存考试成绩
func SaveExamResult(client *redis.Client, examResult ExamResult) error {
	// 使用哈希数据结构保存考试成绩
	for username, score := range examResult.Scores {
		_, err := client.HSet("exam:"+examResult.TeamName+":user:"+username, "score", score).Result()
		if err != nil {
			return err
		}
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
