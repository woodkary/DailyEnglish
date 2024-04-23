package main

import (
	controlsql "DailyEnglish/Control_SQL"
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 数据库连接信息
	username := "mimahezhanghao1yang"
	password := "MIMAhezhanghao1yang"
	hostname := "rm-wz9p61j3qlj6lg69fpo.mysql.rds.aliyuncs.com"
	port := "3306"
	dbname := "dailyenglish"

	// 构建数据库连接字符串
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)

	// 连接数据库
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 检查数据库连接是否成功
	err = db.Ping()
	if err != nil {
		panic("Failed to connect to the database")
	}

	fmt.Println("Successfully connected to the database")

	// Redis连接
	client := redis.NewClient(&redis.Options{
		Addr:     "r-bp1jdmrszl1yd6xxdipd.redis.rds.aliyuncs.com:6379", // Redis服务器地址
		Password: "MIMAhezhanghao1yang",                                // Redis服务器密码
		DB:       255,                                                  // 使用的Redis数据库编号
	})

	defer client.Close()

	// 检查连接是否成功
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("连接Redis失败:", err)
	} else {
		fmt.Println("连接Redis成功:", pong)
	}

	//数据库测试
	// 调用 insertData 函数插入测试数据
	controlsql.InsertData(client)
	//controlsql.InsertUserInfo(db, "小明", "10086", "12344", "123456@qq.com", 2024000123, 19, 1, "2024-04-01")
	//数据库测试

}
