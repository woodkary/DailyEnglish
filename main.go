package main

import (
	controlsql "DailyEnglish/Control_SQL"
	teamrouter "DailyEnglish/router/team_router"

	service "DailyEnglish/Service"

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

	defer client.Close()

	//数据库测试

	controlsql.InsertUserInfo(db, "小明", "10086", "12344", "123456@qq.com", 2024000123, 19, 1, "2024-04-01")
	//数据库测试

	r := gin.Default()
	r.Static("static/team_manager", "./static")
	// r.Static("static/team_manager/css", "./static/css")
	// r.Static("static/team_manager/js", "./static/js")
	//r.LoadHTMLFiles("./static/login.html", "./static/register.html", "./static/forgot_password.html", "./static/index.html", "./static/404.html")

	service.TestAES()

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

	//重定向至登录页面
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/static/team_manager/login.html")
		//c.String(http.StatusOK, "Welcome to Daily English!")
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

		if !controlsql.SearchUserByUsername(db, data.Username) {
			c.JSON(http.StatusForbidden, gin.H{
				"code":    "403",
				"message": "用户不存在",
			})
		} else if controlsql.CheckUser(db, data.Username, data.Pwd) {
			c.JSON(http.StatusOK, gin.H{
				"code":    "200",
				"message": "登录成功",
				"token":   "123456",
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
