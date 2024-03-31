package controlsql

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis" // 导入 redis 驱动
)

// FindTeamInfoByName 通过团队名查找团队信息
func FindTeamInfoByName(client *redis.Client, teamName string) (map[string]string, error) {
	teamInfo := make(map[string]string)

	// 查找团队ID
	teamID, err := client.Get(fmt.Sprintf("TeamName:%s:ID", teamName)).Result()
	if err != nil {
		return nil, err
	}
	teamInfo["ID"] = teamID

	// 查找团队成员
	members, err := client.SMembers(fmt.Sprintf("Team:%s:Members", teamID)).Result()
	if err != nil {
		return nil, err
	}
	teamInfo["Members"] = strings.Join(members, ", ")

	// 查找临近的考试ID
	upcomingExamID, err := client.Get(fmt.Sprintf("Team:%s:UpcomingExamID", teamID)).Result()
	if err != nil {
		return nil, err
	}
	teamInfo["UpcomingExamID"] = upcomingExamID

	return teamInfo, nil
}

// FindScoresByExamAndTeam 通过考试ID和团队名找到该考试该团队所有成员的成绩
func FindScoresByExamAndTeam(client *redis.Client, examID, teamID string) (map[string]int, error) {
	scores := make(map[string]int)

	// 获取该团队该考试下所有成员的成绩
	memberScores, err := client.HGetAll(fmt.Sprintf("Exam:%s:Team:%s:Score", examID, teamID)).Result()
	if err != nil {
		return nil, err
	}

	// 将成绩转换为整数类型
	for username, scoreStr := range memberScores {
		score, err := strconv.Atoi(scoreStr)
		if err != nil {
			return nil, err
		}
		scores[username] = score
	}

	return scores, nil
}

// FindExamInfoByID 通过考试ID找到考试信息
func FindExamInfoByID(client *redis.Client, examID string) (map[string]string, error) {
	examInfo := make(map[string]string)

	// 查找考试名称
	examName, err := client.Get(fmt.Sprintf("Exam:%s:Name", examID)).Result()
	if err != nil {
		return nil, err
	}
	examInfo["Name"] = examName

	// 查找考试日期
	examDate, err := client.Get(fmt.Sprintf("Exam:%s:Date", examID)).Result()
	if err != nil {
		return nil, err
	}
	examInfo["Date"] = examDate

	// 查找题目数量
	numOfQuestionsStr, err := client.Get(fmt.Sprintf("Exam:%s:NumOfQuestions", examID)).Result()
	if err != nil {
		return nil, err
	}
	examInfo["NumOfQuestions"] = numOfQuestionsStr

	return examInfo, nil
}
