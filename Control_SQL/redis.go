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
