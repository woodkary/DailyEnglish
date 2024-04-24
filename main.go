package main

import (
	controlsql "DailyEnglish/Control_SQL"
	teamrouter "DailyEnglish/router/team_router"
	userrouter "DailyEnglish/router/user_router"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
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

	// Redis连接
	client := redis.NewClient(&redis.Options{
		Addr:     "r-bp1jdmrszl1yd6xxdipd.redis.rds.aliyuncs.com:6379", // Redis服务器地址
		Password: "MIMAhezhanghao1yang",                                // Redis服务器密码
		DB:       255,                                                  // 使用的Redis数据库编号
	})
	client.Ping().Result()
	defer client.Close()

	//数据库测试
	//teamName := "每日背单词小组"
	controlsql.GetUsernameByEmail(db, "1234567@qq.com")
	//controlsql.InsertUserInfo(db, "小明", "10086", "12344", "1234567@qq.com", 2024000123, 19, 1, "2024-04-01")

	r := gin.Default()
	r.Static("static/team_manager", "./static")
	// r.Static("static/team_manager/css", "./static/css")
	// r.Static("static/team_manager/js", "./static/js")
	//r.LoadHTMLFiles("./static/login.html", "./static/register.html", "./static/forgot_password.html", "./static/index.html", "./static/404.html")
	// r.LoadHTMLGlob("./static/*.html")

	userrouter.InitUserRouter(r, client, db)
	teamrouter.InitTeamRouter(r, client, db)
	r.Run(":8080")

	users, err := controlsql.QueryUserInfo(db)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
