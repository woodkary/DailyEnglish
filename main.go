package main

import (
	middlewares "DailyEnglish/middlewares"
	adminrouter "DailyEnglish/router/admin_router"
	teamrouter "DailyEnglish/router/team_router"
	userrouter "DailyEnglish/router/user_router"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
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
	r := gin.Default()
	r.Static("static/team_manager", "./static")
	// r.Static("static/team_manager/css", "./static/css")
	// r.Static("static/team_manager/js", "./static/js")
	// r.LoadHTMLFiles("./static/login.html", "./static/register.html", "./static/forgot_password.html", "./static/index.html", "./static/404.html")
	adminrouter.InitAdminRouter(r, db)
	teamrouter.InitTeamRouter(r, db)
	go func() {
		r1 := gin.Default()
		userrouter.InitUserRouter(r1, db)
		r1.Use(middlewares.Cors())
		r1.Run(":8080")
	}()
	r.Use(middlewares.Cors())
	r.Run(":8081")
}
