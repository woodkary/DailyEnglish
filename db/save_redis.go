package controlsql

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

type Myteams struct {
	Username string
	Teams    []string
}

type Team struct { //hash存储
	Name           string // 团队名
	ID             int    // 团队ID
	TotalMembers   int    // 团队总人数
	AdminCount     int    // 管理员人数
	RecentExamDate string // 最近一场考试日期

	Members []Member // 团队成员列表
}

type Member struct { //hash
	Username       string // 成员用户名
	JoinDate       string // 加入团队日期
	AttendanceDays int    // 打卡天数
	IsAdmin        bool   // 是否是团队管理员
	AttendanceRate string //打卡率
	AttendanceNum  int    //打卡单词数量
}

type AttendanceRecord struct { //hash
	Date             string         // 日期
	TeamName         string         // 团队名
	AttendanceCount  int            // 打卡人数
	MemberAttendance map[string]int // 成员打卡情况，键为用户名，值为是否打卡，1是0否
	AttendanceRate   float64        // 打卡率
}

type TeamRequest struct { //list  根据团队名和id查
	Flag     string // 0加入申请 1管理员申请
	TeamName string // 团队名
	Username string // 申请者用户名
	Time     string // 申请时间
	Message  string // 申请留言内容
}

type Notification struct { //zset 集合 根据团队名和flag查
	ID       string // 通知ID
	flag     string //已处理标志，0未处理1已经处理
	Title    string // 通知标题
	Content  string // 通知内容
	Time     string // 通知时间
	TeamName string // 通知团队名
}

type ExamResult struct { //键值对的字符串string
	TeamName string // 团队名
	ExamName string
	Scores   map[string]int // 成员分数，键为用户名，值为分数
	Rankings map[string]int // 成员排名，键为用户名，值为排名
}

type ExamInfo struct { //hash
	ID            int            // 考试ID
	Date          string         //日期
	Name          string         // 考试名称
	QuestionCount int            // 试题数量
	Questions     []string       // 试题内容
	AverageScore  float64        // 考试平均分
	PassRate      float64        // 及格率
	TopSix        map[string]int // 前6名成员用户名及分数，键为用户名，值为分数
}

// 函数用于保存 Myteamseam 结构体到 Redis 数据库
func savemyteams(redisClient *redis.Client, team *Myteams) error {
	// 将 Myteamseam 结构体转换为 JSON 字符串
	jsonData, err := json.Marshal(team)
	if err != nil {
		return err
	}

	// 保存 JSON 字符串到 Redis
	err = redisClient.Set(team.Username, jsonData, 0).Err()
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
// (1)根据团队名和用户名，将用户加入团队中，并更新团队成员信息和成绩信息
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

// (2)函数用于将团队名添加到指定用户名的 Teams 数组中
func Addmyteams(client *redis.Client, username, teamName string) error {
	// 从 Redis 获取指定用户名的团队信息
	val, err := client.Get(username).Result()
	if err != nil {
		return err
	}

	// 解析 JSON 数据到 Myteamseam 结构体
	var userTeams Myteams
	err = json.Unmarshal([]byte(val), &userTeams)
	if err != nil {
		return err
	}

	// 将团队名添加到 Teams 数组中
	userTeams.Teams = append(userTeams.Teams, teamName)

	// 保存更新后的团队信息到 Redis
	err = savemyteams(client, &userTeams)
	if err != nil {
		return err
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
	_, err := client.HMSet("exam_info:"+examInfo.Name, map[string]interface{}{
		"Date":           examInfo.Date,
		"Name":           examInfo.Name,
		"Question_count": examInfo.QuestionCount,
		"Average_score":  examInfo.AverageScore,
		"Pass_rate":      examInfo.PassRate,
	}).Result()
	client.Set("exam_name:"+examInfo.Name, examInfo.ID, 0)

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