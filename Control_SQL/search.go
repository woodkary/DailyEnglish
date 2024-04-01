package controlsql

import (
	"fmt"
	"database/sql"
	"github.com/go-redis/redis" // 导入 redis 驱动
)



// 1. 通过用户名查询用户信息
func GetUserInfoByUsername(db *sql.DB, username string) (UserInfo, error) {
    var userInfo UserInfo
    query := "SELECT * FROM user_info WHERE username = ?"
    err := db.QueryRow(query, username).Scan(&userInfo.Username, &userInfo.ID, &userInfo.Phone, &userInfo.Pwd, &userInfo.Email, &userInfo.Age, &userInfo.Sex, &userInfo.RegisterDate)
    if err != nil {
        return UserInfo{}, fmt.Errorf("error getting user info by username: %v", err)
    }
    return userInfo, nil
}

// 2. 通过团队名查询团队信息
func GetTeamInfoByName(client *redis.Client, teamName string) (TeamInfo, error) {
    teamInfo, err := client.HGetAll("team_info").Result()
    if err != nil {
        return TeamInfo{}, err
    }
    if val, ok := teamInfo[teamName]; ok {
        var info TeamInfo
        err := json.Unmarshal([]byte(val), &info)
        if err != nil {
            return TeamInfo{}, err
        }
        return info, nil
    }
    return TeamInfo{}, fmt.Errorf("team %s not found", teamName)
}

// 3. 通过团队名查询团队打卡信息
func GetAttendanceByTeamName(client *redis.Client, teamName string) (AttendanceInfo, error) {
    key := fmt.Sprintf("attendance:*:%s", teamName)
    keys, err := client.Keys(key).Result()
    if err != nil {
        return AttendanceInfo{}, err
    }
    var attendance AttendanceInfo
    for _, k := range keys {
        data, err := client.Get(k).Bytes()
        if err != nil {
            return AttendanceInfo{}, err
        }
        var info AttendanceInfo
        err = json.Unmarshal(data, &info)
        if err != nil {
            return AttendanceInfo{}, err
        }
        attendance = info
        break
    }
    return attendance, nil
}

// 4. 通过用户名和团队名查询该用户在该团队的打卡信息
func GetAttendanceByUsernameAndTeamName(client *redis.Client, username string, teamName string) (AttendanceInfo, error) {
    key := fmt.Sprintf("attendance:*:%s", teamName)
    keys, err := client.Keys(key).Result()
    if err != nil {
        return AttendanceInfo{}, err
    }
    var attendance AttendanceInfo
    for _, k := range keys {
        data, err := client.Get(k).Bytes()
        if err != nil {
            return AttendanceInfo{}, err
        }
        var info AttendanceInfo
        err = json.Unmarshal(data, &info)
        if err != nil {
            return AttendanceInfo{}, err
        }
        if _, ok := info.Attendance[username]; ok {
            attendance = info
            break
        }
    }
    return attendance, nil
}

// 5. 通过用户名，团队名和考试名称查询该用户考试成绩
func GetScoresByUsernameTeamNameAndExamName(client *redis.Client, username string, teamName string, examName string) (int, error) {
    key := fmt.Sprintf("exam_scores:%s:%s", examName, teamName)
    score, err := client.ZScore(key, username).Result()
    if err != nil {
        return 0, err
    }
    return int(score), nil
}

// 6. 通过团队名和考试名称查询该团队该考试的成绩信息
func GetScoresByTeamNameAndExamName(client *redis.Client, teamName string, examName string) (map[string]int, error) {
    key := fmt.Sprintf("exam_scores:%s:%s", examName, teamName)
    scores, err := client.ZRangeWithScores(key, 0, -1).Result()
    if err != nil {
        return nil, err
    }
    result := make(map[string]int)
    for _, score := range scores {
        result[score.Member.(string)] = int(score.Score)
    }
    return result, nil
}

// 7. 通过团队名查询该团队的通知信息
func GetNotificationByTeamName(client *redis.Client, teamName string) ([]NotificationInfo, error) {
    key := fmt.Sprintf("notification:*")
    keys, err := client.Keys(key).Result()
    if err != nil {
        return nil, err
    }
    var notifications []NotificationInfo
    for _, k := range keys {
        data, err := client.HGetAll(k).Result()
        if err != nil {
            return nil, err
        }
        if data["team_name"] == teamName {
            var info NotificationInfo
            info.Title = data["title"]
            info.Content = data["content"]
            info.Time = data["time"]
            info.TeamName = data["team_name"]
            notifications = append(notifications, info)
        }
    }
    return notifications, nil
}

// 8. 通过考试名获取该厂考试信息
func GetExamInfoByExamName(client *redis.Client, examName string) (ExamInfo, error) {
    keys, err := client.Keys(fmt.Sprintf("exam_info:*")).Result()
    if err != nil {
        return ExamInfo{}, err
    }
    for _, key := range keys {
        data, err := client.HGetAll(key).Result()
        if err != nil {
            return ExamInfo{}, err
        }
        if data["exam_name"] == examName {
            var info ExamInfo
            info.ExamName = data["exam_name"]
            info.TotalQuestions, _ = strconv.Atoi(data["total_questions"])
            info.AverageScore = data["average_score"]
            info.PassRate = data["pass_rate"]
            return info, nil
        }
    }
    return ExamInfo{}, fmt.Errorf("exam %s not found", examName)
}

// 9. 通过日期查询当天所有考试信息
func GetExamsByDate(client *redis.Client, date string) ([]ExamInfo, error) {
    var exams []ExamInfo
    keys, err := client.Keys(fmt.Sprintf("exam_info:%s", date)).Result()
    if err != nil {
        return nil, err
    }
    for _, key := range keys {
        data, err := client.HGetAll(key).Result()
        if err != nil {
            return nil, err
        }
        var info ExamInfo
        info.ExamName = data["exam_name"]
        info.TotalQuestions, _ = strconv.Atoi(data["total_questions"])
        info.AverageScore = data["average_score"]
        info.PassRate = data["pass_rate"]
        exams = append(exams, info)
    }
    return exams, nil
}

// 10. 通过用户名查询该用户是否团队管理员
func IsUserTeamAdmin(client *redis.Client, username string) (bool, error) {
    keys, err := client.Keys(fmt.Sprintf("team_members:*")).Result()
    if err != nil {
        return false, err
    }
    for _, key := range keys {
        _, err := client.HGet(key, username).Result()
        if err == nil {
            return true, nil
        }
    }
    return false, nil
}

// 11. 通过团队名查询该团队所有管理员信息
func GetTeamAdminsByTeamName(client *redis.Client, teamName string) ([]string, error) {
    key := fmt.Sprintf("team_members:%s", teamName)
    admins, err := client.HKeys(key).Result()
    if err != nil {
        return nil, err
    }
    return admins, nil
}

// 12. 通过团队名查询该团队所有团队管理员申请信息
func GetAdminRequestsByTeamName(client *redis.Client, teamName string) ([]AdminRequestInfo, error) {
    key := fmt.Sprintf("admin_requests:%s:*", teamName)
    keys, err := client.Keys(key).Result()
    if err != nil {
        return nil, err
    }
    var requests []AdminRequestInfo
    for _, k := range keys {
        data, err := client.HGetAll(k).Result()
        if err != nil {
            return nil, err
        }
        var request AdminRequestInfo
        request.Username = strings.Split(k, ":")[2]
        request.RequestTime = data["request_time"]
        request.Message = data["message"]
        requests = append(requests, request)
    }
    return requests, nil
}
4FSFSF