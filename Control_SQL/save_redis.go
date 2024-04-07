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
	AttendanceRate string //打卡率
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
	Name          string         // 考试名称
	QuestionCount int            // 试题数量
	Questions     []string       // 试题内容
	AverageScore  float64        // 考试平均分
	PassRate      float64        // 及格率
	TopTen        map[string]int // 前十名成员用户名及分数，键为用户名，值为分数
}

// 保存团队信息
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
			"attendance_rate": member.AttendanceRate, // 保存成员的打卡率
		}).Result()
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

// 保存团队管理员申请信息
func SaveAdminRequest(client *redis.Client, request AdminRequest) error {
	// 使用哈希数据结构保存管理员申请信息
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
		Member: fmt.Sprintf("%s|%s|%s|%s", notification.ID, notification.Title, notification.Content, notification.Time),
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
		"name":           examInfo.Name,
		"question_count": examInfo.QuestionCount,
		"average_score":  examInfo.AverageScore,
		"pass_rate":      examInfo.PassRate,
	}).Result()
	if err != nil {
		return err
	}

	// 保存前十名成员信息
	for username, score := range examInfo.TopTen {
		_, err := client.HSet("exam_info:"+strconv.Itoa(examInfo.ID)+":top_ten", username, score).Result()
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
