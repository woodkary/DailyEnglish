package controlsql

import (
	"fmt"

	"github.com/go-redis/redis" // 导入 redis 驱动
)

// //////////////////////////////////////////////////////////////
// nosql
// StoreTeamInfoRedis 存储团队信息到 Redis，包括团队名（主键）、团队ID、团队成员的用户名和临近的考试ID
func StoreTeamInfoRedis(client *redis.Client, teamName, teamID string, members []string, upcomingExamID string) error {

	// 存储团队名（主键）
	err := client.Set(fmt.Sprintf("Team:%s:Name", teamID), teamName, 0).Err()
	if err != nil {
		return err
	}

	// 存储团队ID
	err = client.Set(fmt.Sprintf("Team:%s:ID", teamID), teamID, 0).Err()
	if err != nil {
		return err
	}

	// 存储团队成员的用户名
	for _, member := range members {
		err = client.SAdd(fmt.Sprintf("Team:%s:Members", teamID), member).Err()
		if err != nil {
			return err
		}
	}

	// 存储临近的考试ID
	err = client.Set(fmt.Sprintf("Team:%s:UpcomingExamID", teamID), upcomingExamID, 0).Err() // 注意这里的参数类型应为 interface{}
	if err != nil {
		return err
	}

	return nil
}

// SaveScores 保存考试成绩信息到 Redis
// 接收考试ID examID、考试名称 examName 和一个映射 teams，该映射的键是团队ID，值是另一个映射，其键是用户名，值是该用户的考试成绩。
func SaveScores(client *redis.Client, examID, examName string, teams map[string]map[string]int) error {

	// 保存考试名称
	err := client.Set(fmt.Sprintf("Exam:%s:Name", examID), examName, 0).Err()
	if err != nil {
		return err
	}

	// 保存每个团队的成绩
	for teamID, scores := range teams {
		for username, score := range scores {
			key := fmt.Sprintf("Exam:%s:Team:%s:Score", examID, teamID)
			err := client.HSet(key, username, score).Err()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// SaveExam 保存考试信息到 Redis
// SaveExam 保存考试信息到 Redis
func SaveExam(client *redis.Client, examID, examName, examDate string, numOfQuestions int) error {

	// 保存考试名称
	err := client.Set(fmt.Sprintf("Exam:%s:Name", examID), examName, 0).Err()
	if err != nil {
		return err
	}

	// 保存考试日期
	err = client.Set(fmt.Sprintf("Exam:%s:Date", examID), examDate, 0).Err()
	if err != nil {
		return err
	}

	// 保存题目数量
	err = client.Set(fmt.Sprintf("Exam:%s:NumOfQuestions", examID), numOfQuestions, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
