package controlsql

import (
	"github.com/go-redis/redis"
)

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
		Flag:     "0",
		TeamName: "每日背单词小组",

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
		Flag:     "0",
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
		Flag:     "1",
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
		Flag:     "1",
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
		Flag:     "1",
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
		Flag:     "0",
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
		Flag:     "0",
		TeamName: "每日背单词小组",

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
