package main

import (
	controlsql "DailyEnglish/Control_SQL"
	service "DailyEnglish/Service"
	teamrouter "DailyEnglish/router/team_router"
	"database/sql"
	"fmt"
	"log"
	"net/http"

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

	//controlsql.InsertUserInfo(db, "小明", "10086", "12344", "123456@qq.com", 2024000123, 19, 1, "2024-04-01")
	//数据库测试

	r := gin.Default()
	r.Static("static/team_manager", "./static")
	// r.Static("static/team_manager/css", "./static/css")
	// r.Static("static/team_manager/js", "./static/js")
	//r.LoadHTMLFiles("./static/login.html", "./static/register.html", "./static/forgot_password.html", "./static/index.html", "./static/404.html")
	// r.LoadHTMLGlob("./static/*.html")

	//注册接口
	r.POST("/api/team_manager/register", func(c *gin.Context) {
		type regdata struct {
			Username   string `json:"username"`
			Pwd        string `json:"pwd"`
			Email      string `json:"email"`
			VerifyCode string `json:"verify_code"`
		}
		var data regdata
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if controlsql.SearchUserByUsername(db, data.Username) {
			//TODO
			//验证码还妹搞
			if controlsql.InsertUser(db, data.Username, data.Pwd, data.Email) != nil {
				c.JSON(http.StatusOK, gin.H{
					"code":    "200",
					"message": "注册成功",
				})
			} else {
				c.JSON(http.StatusConflict, gin.H{
					"code":  "409",
					"error": "用户已注册",
				})
			}
		}
	})

	r.GET("/api/team_manager/login", func(c *gin.Context) {
		c.File("/static/team_manager/login.html")
	})

	//重定向至登录页面
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/static/team_manager/login.html")
	})
	// //登录界面验证header头部
	// r.GET("/static/team_manager/login.html", func(c *gin.Context) {

	// })
	//登录接口
	r.POST("/api/team_manager/login", func(c *gin.Context) {
		type logdata struct {
			Username string `json:"username"`
			Pwd      string `json:"password"`
		}
		var data logdata
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  "400",
				"error": err.Error(),
			})
			return
		}

		if !controlsql.SearchUserByUsername(db, data.Username) {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    "403",
				"message": "用户不存在",
			})
		} else if controlsql.CheckUser(db, data.Username, data.Pwd) {
			teamName, err := controlsql.GetJoinedTeams(client, data.Username)
			//输出所有teamname
			fmt.Println("teamName:", teamName)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":  "500",
					"error": "服务器错误",
				})
				return
			}
			token, err := service.GenerateToken(data.Username, teamName[0])

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":  "500",
					"error": "服务器错误",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "登录成功",
				"token":   token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  "401",
				"error": "用户名或密码有误",
			})
		}
	})
	teamrouter.Team_manager(r, client, db)
	r.Run(":8080")

	users, err := controlsql.QueryUserInfo(db)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}

func main2() {
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

	// 调用 insertData 函数插入测试数据
	err = controlsql.InsertData(client)
	if err != nil {
		fmt.Println("Failed to insert test data:", err)
		return
	}

	fmt.Println("Test data inserted successfully.")

}
