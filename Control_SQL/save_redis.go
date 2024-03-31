package controlsql

import (
	"fmt"

	"github.com/go-redis/redis" // 导入 redis 驱动
)

//打卡信息
func SaveAttendance(client *redis.Client, date string, teamName string, attendance map[string]int, attendanceRate string) error {
    attendanceInfo := AttendanceInfo{
        TeamName:      teamName,
        Date:          date,
        Attendance:    attendance,
        AttendanceRate: attendanceRate,
    }
    jsonData, err := json.Marshal(attendanceInfo)
    if err != nil {
        return err
    }
    key := fmt.Sprintf("attendance:%s:%s", date, teamName)
    err = client.Set(key, jsonData, 0).Err()
    if err != nil {
        return err
    }
    return nil
}


// SaveTeam 存储团队信息到 Redis
func SaveTeam(client *redis.Client, teamName string, teamID int, totalMembers int, adminCount int, adminFlag bool, lastExamDate string, recent7DaysAttendance int, attendanceRate string, members map[string]string) error {
    // Save team info
    err := client.HSet("team_info", teamName, json.Marshal(map[string]interface{}{
        "team_id":                 teamID,
        "total_members":           totalMembers,
        "admin_count":             adminCount,
        "admin_flag":              adminFlag,
        "last_exam_date":          lastExamDate,
        "recent_7_days_attendance": recent7DaysAttendance,
        "attendance_rate":         attendanceRate,
    })).Err()
    if err != nil {
        return err
    }

    // Save team members
    for username, joinDate := range members {
        err := client.HSet(fmt.Sprintf("team_members:%s", teamName), username, joinDate).Err()
        if err != nil {
            return err
        }
    }
    return nil
}


// SaveScores 保存考试成绩信息到 Redis
func SaveScores(client *redis.Client, examID string, teamName string, scores map[string]int) error {
    key := fmt.Sprintf("exam_scores:%s:%s", examID, teamName)
    for username, score := range scores {
        err := client.ZAdd(key, &redis.Z{Score: float64(score), Member: username}).Err()
        if err != nil {
            return err
        }
    }
    return nil
}


// SaveExam 保存考试信息到 Redis
func SaveExam(client *redis.Client, examID string, examName string, totalQuestions int, averageScore string, passRate string, topTen map[string]map[string]interface{}) error {
    key := fmt.Sprintf("exam_info:%s", examID)
    err := client.HSet(key, map[string]interface{}{
        "exam_name":       examName,
        "total_questions": totalQuestions,
        "average_score":   averageScore,
        "pass_rate":       passRate,
    }).Err()
    if err != nil {
        return err
    }

    // Save top ten
    for rank, data := range topTen {
        err := client.HSet(fmt.Sprintf("%s:top_ten", key), rank, data).Err()
        if err != nil {
            return err
        }
    }
    return nil
}

//通知信息
func SaveNotification(client *redis.Client, notificationID int, title string, content string, time string, teamName string) error {
    key := fmt.Sprintf("notification:%d", notificationID)
    err := client.HSet(key, map[string]interface{}{
        "title":      title,
        "content":    content,
        "time":       time,
        "team_name":  teamName,
    }).Err()
    if err != nil {
        return err
    }
    return nil
}

//团队管理员申请
func SaveAdminRequest(client *redis.Client, teamName string, username string, requestTime string, message string) error {
    key := fmt.Sprintf("admin_requests:%s:%s", teamName, username)
    err := client.HSet(key, map[string]interface{}{
        "request_time": requestTime,
        "message":      message,
    }).Err()
    if err != nil {
        return err
    }
    return nil
}
