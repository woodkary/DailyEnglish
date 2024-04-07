package main

import (
	controlsql "DailyEnglish/Control_SQL"

	service "DailyEnglish/Service"

	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main1() {

	//redis连接
	client := redis.NewClient(&redis.Options{
		Addr:     "r-bp1jdmrszl1yd6xxdipd.redis.rds.aliyuncs.com:6379", // Redis 服务器地址
		Password: "MIMAhezhanghao1yang",                                // Redis 服务器密码
		DB:       255,                                                  // 使用的 Redis 数据库编号
	})

	defer client.Close()
	// 定义和初始化 team 变量
	team := controlsql.Team{
		Name:           "TeamA",
		ID:             34,
		TotalMembers:   9,
		AdminCount:     3,
		RecentExamDate: "2024-04-01",
		Last7DaysAttendance: struct {
			Count int
			Rate  float64
		}{Count: 9, Rate: 1.00},
		Members: []controlsql.Member{
			{Username: "张三", JoinDate: "2024-01-01", AttendanceDays: 90, IsAdmin: true},
			{Username: "李四", JoinDate: "2024-01-15", AttendanceDays: 88, IsAdmin: false},
			{Username: "Charlie", JoinDate: "2024-02-01", AttendanceDays: 87, IsAdmin: false},
			{Username: "Alice", JoinDate: "2024-02-15", AttendanceDays: 89, IsAdmin: false},
			{Username: "Bob", JoinDate: "2024-03-01", AttendanceDays: 86, IsAdmin: false},
			{Username: "小明", JoinDate: "2024-03-15", AttendanceDays: 91, IsAdmin: false},
			{Username: "李白", JoinDate: "2024-04-01", AttendanceDays: 92, IsAdmin: false},
			{Username: "龙王", JoinDate: "2024-04-15", AttendanceDays: 93, IsAdmin: false},
			{Username: "牢大", JoinDate: "2024-05-01", AttendanceDays: 94, IsAdmin: false},
		},
	}

	// 保存团队信息到数据库
	err := controlsql.SaveTeam(client, team)
	if err != nil {
		fmt.Println("保存团队信息失败:", err)
		return
	}

	// 成功保存团队信息后输出成功提示
	fmt.Println("团队信息保存成功！")
}
func main() {

	//mysql连接
	db, er := sql.Open("mysql", "root:123456@tcp(47.107.81.75:3306)/Daily-English")
	if er != nil {
		log.Fatal(er)
	}
	defer db.Close()

	controlsql.InsertUserInfo(db, "小明", "10086", "12344", "123456@qq.com", 2024000123, 19, 1, "2024-04-01")

}
func main2() {

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

	r.GET("/getToken/:userName", func(c *gin.Context) {
		str := c.Param("userName")
		fmt.Println(str)
		str1, err := service.GenerateToken(str)
		str2, err := service.ParseToken(str1)
		c.JSON(200, gin.H{
			"code":        "200",
			"msg":         "123456",
			"userName":    str,
			"token":       str1,
			"error":       err,
			"tokenParsed": str2,
		})
	})

	r.GET("/api/team_manager/index", func(c *gin.Context) {
		//@TODO
		//添加发送前端需要的json数据
		c.JSON(200, gin.H{"code": "200", "msg": "成功", "completed": 80, "uncompleted": 20, "exam": "exam"})
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
