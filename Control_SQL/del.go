package controlsql

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

// 1 退出团队
func delmyteams(client *redis.Client, username, teamName string) error {
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

	// 查找团队名在 Teams 数组中的索引
	index := -1
	for i, name := range userTeams.Teams {
		if name == teamName {
			index = i
			break
		}
	}

	// 如果找到了团队名，则从 Teams 数组中删除
	if index != -1 {
		userTeams.Teams = append(userTeams.Teams[:index], userTeams.Teams[index+1:]...)
	} else {
		return fmt.Errorf("team %s not found in user's teams", teamName)
	}

	// 保存更新后的团队信息到 Redis
	err = savemyteams(client, &userTeams)
	if err != nil {
		return err
	}

	return nil
}
func delmember(client *redis.Client, teamName, username string) error {
	// 构建 Redis 键名
	teamKey := "Team:" + teamName
	memberKey := "Member:" + username

	// 删除成员信息
	err := client.HDel(teamKey, memberKey).Err()
	if err != nil {
		return err
	}

	return nil
}
