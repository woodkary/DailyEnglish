package controlsql

import (
	"fmt"
	"time"

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

type Question struct {
	Description string // 题目描述
	Answer      string // 答案
}

// SaveExam 保存考试信息到 Redis
// SaveExam 保存考试信息到 Redis
func SaveExam(client *redis.Client, examID, examName string, examDate time.Time, questions map[int]Question) error {

	// 保存考试名称
	err := client.Set(fmt.Sprintf("Exam:%s:Name", examID), examName, 0).Err()
	if err != nil {
		return err
	}

	// 保存考试日期
	err = client.Set(fmt.Sprintf("Exam:%s:Date", examID), examDate.Format("2006-01-02"), 0).Err()
	if err != nil {
		return err
	}

	// 保存题目数量
	err = client.Set(fmt.Sprintf("Exam:%s:QuestionCount", examID), len(questions), 0).Err()
	if err != nil {
		return err
	}

	// 保存每个题目的描述和答案
	for number, question := range questions {
		descriptionKey := fmt.Sprintf("Exam:%s:Question:%d:Description", examID, number)
		err := client.Set(descriptionKey, question.Description, 0).Err()
		if err != nil {
			return err
		}

		answerKey := fmt.Sprintf("Exam:%s:Question:%d:Answer", examID, number)
		err = client.Set(answerKey, question.Answer, 0).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
