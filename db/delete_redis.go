package controlsql

import (
	"fmt"

	"github.com/go-redis/redis" // 导入 redis 驱动
)

// DeleteUserFromTeamc 从 Redis 中删除团队成员
func DeleteUserFromTeam(client *redis.Client, teamID string, username string) error {
	// 删除团队成员
	err := client.SRem(fmt.Sprintf("Team:%s:Members", teamID), username).Err()
	if err != nil {
		return err
	}

	return nil
}

// DeleteTeamInfoRedis 从 Redis 中删除团队信息，包括团队名、团队ID、团队成员的用户名和临近的考试ID
func DeleteTeamInfo(client *redis.Client, teamID string) error {
	// 删除团队名
	err := client.Del(fmt.Sprintf("Team:%s:Name", teamID)).Err()
	if err != nil {
		return err
	}

	// 删除团队ID
	err = client.Del(fmt.Sprintf("Team:%s:ID", teamID)).Err()
	if err != nil {
		return err
	}

	// 删除团队成员
	err = client.Del(fmt.Sprintf("Team:%s:Members", teamID)).Err()
	if err != nil {
		return err
	}

	// 删除临近的考试ID
	err = client.Del(fmt.Sprintf("Team:%s:UpcomingExamID", teamID)).Err()
	if err != nil {
		return err
	}

	return nil
}

// DeleteScores 从 Redis 中删除考试成绩信息
func DeleteScores(client *redis.Client, examID string) error {
	// 删除考试名称
	err := client.Del(fmt.Sprintf("Exam:%s:Name", examID)).Err()
	if err != nil {
		return err
	}

	// 删除每个团队的成绩
	err = client.Del(fmt.Sprintf("Exam:%s:Team:*:Score", examID)).Err()
	if err != nil {
		return err
	}

	return nil
}

// DeleteExam 从 Redis 中删除考试信息
func DeleteExam(client *redis.Client, examID string) error {
	// 删除考试名称
	err := client.Del(fmt.Sprintf("Exam:%s:Name", examID)).Err()
	if err != nil {
		return err
	}

	// 删除考试日期
	err = client.Del(fmt.Sprintf("Exam:%s:Date", examID)).Err()
	if err != nil {
		return err
	}

	// 删除题目数量
	err = client.Del(fmt.Sprintf("Exam:%s:NumOfQuestions", examID)).Err()
	if err != nil {
		return err
	}

	return nil
}
