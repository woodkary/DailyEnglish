package main

import (
	controlsql "DailyEnglish/Control_SQL"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "r-bp1jdmrszl1yd6xxdipd.redis.rds.aliyuncs.com:6379", // Redis 服务器地址
	// 	Password: "MIMAhezhanghao1yang",                                // Redis 服务器密码
	// 	DB:       255,                                                  // 使用的 Redis 数据库编号
	// })
	// pong, err := client.Ping().Result()
	// fmt.Println(pong, err)
	// defer client.Close()

	r := gin.Default()
	r.Static("api/team_manager/static", "./static")
	r.Static("api/team_manager/css", "./static/css")
	r.Static("api/team_manager/js", "./static/js")
	//重定向至登录页面
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/api/team_manager/login.html")
		//c.String(http.StatusOK, "Welcome to Daily English!")
	})

	//注册页面
	r.GET("/api/team_manager/register", func(c *gin.Context) {
		c.File("./static/register.html")
	})

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

	//登录页面
	r.GET("/api/team_manager/login", func(c *gin.Context) {
		c.File("./static/login.html")
	})

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

		if controlsql.SearchUserByUsername(db, data.Username) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "登录成功",
				"token":   "123456",
			})
		} else if controlsql.CheckUser(db, data.Username, data.Pwd) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "登录成功",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  "401",
				"error": "用户名或密码有误",
			})
		}
	})

	//忘记密码页面
	r.GET("/api/team_manager/forgot_password", func(c *gin.Context) {
		c.File("./static/forgot_password.html")
	})

	r.Run(":8080")

	users, err := controlsql.QueryUserInfo(db)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
