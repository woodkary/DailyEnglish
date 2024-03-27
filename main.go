package main

import (
	controlsql "DailyEnglish/Control_SQL"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//mysql连接
	db, er := sql.Open("mysql", "root:123456@tcp(47.107.81.75:3306)/daily_english")
	if er != nil {
		log.Fatal(er)
	}
	defer db.Close()

	//redis连接
	client := redis.NewClient(&redis.Options{
		Addr:     "r-bp1jdmrszl1yd6xxdipd.redis.rds.aliyuncs.com:6379", // Redis 服务器地址
		Password: "MIMAhezhanghao1yang",                                // Redis 服务器密码
		DB:       255,                                                  // 使用的 Redis 数据库编号
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	defer client.Close()

	r := gin.Default()
	r.Static("api/team_manager/static", "./static")
	r.Static("api/team_manager/css", "./static/css")
	r.Static("api/team_manager/js", "./static/js")
	//重定向至登录页面
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/team_manager/login")
		//c.String(http.StatusOK, "Welcome to Daily English!")
	})

	//注册页面
	r.GET("/api/team_manager/register", func(c *gin.Context) {
		c.File("./static/register.html")
	})

	//注册接口
	r.POST("/api/team_manager/register", func(c *gin.Context) {
		var user controlsql.UserInfo
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if controlsql.QueryUserInfo(db) == nil {
			if controlsql.InsertUser(db, user) {
				c.JSON(http.StatusOK, gin.H{
					"message": "Register success",
				})
			} else {
				c.JSON(http.StatusConflict, gin.H{
					"error": "Username already exists",
				})
			}
		}
	})

	//登录页面
	r.GET("/api/team_manager/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "./static/login.html", nil)
	})

	//登录接口
	r.POST("/api/team_manager/login", func(c *gin.Context) {
		var user controlsql.UserInfo
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if controlsql.QueryUserInfo(db) == nil {
			if controlsql.CheckUser(db, user.Username, user.Pwd) {
				c.JSON(http.StatusOK, gin.H{
					"message": "Login success",
				})

			} else {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Username or password error",
				})
			}
		}

	})

	//忘记密码页面
	r.GET("/api/team_manager/forgot_password", func(c *gin.Context) {
		c.File("./static/forgot_password.html")
	})

	r.Run(":9090")

	users, err := controlsql.QueryUserInfo(db)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
