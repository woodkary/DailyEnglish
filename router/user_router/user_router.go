package userrouter

import (
<<<<<<< Updated upstream
	controlsql "DailyEnglish/Control_SQL"
	service "DailyEnglish/services"
=======
	controlsql "DailyEnglish/db"
	utils "DailyEnglish/utils"
>>>>>>> Stashed changes
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

<<<<<<< Updated upstream
func User_manager(r *gin.Engine, client *redis.Client, db *sql.DB) {

	//注册接口
	r.POST("/api/team_manager/register", func(c *gin.Context) {
		type regdata struct {
			Username   string `json:"username"`
			Pwd        string `json:"pwd"`
			Email      string `json:"email"`
			VerifyCode string `json:"verify_code"`
=======
func InitUserRouter(r *gin.Engine, db *sql.DB) {
	//发送验证码
	r.POST("/api/register/sendCode", func(c *gin.Context) {
		type response struct {
			Email string `json:"email"`
		}
		var data response
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(400, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		fmt.Print("data.Email: ", data.Email, "\n")
		// 验证邮箱是否已注册
		if controlsql.EmailIsRegistered_User(db, data.Email) {
			c.JSON(http.StatusConflict, gin.H{
				"code": "409",
				"msg":  "邮箱已注册",
			})
			return
		}
		// 初始化 Config 结构体
		config := utils.Config{
			EmailFrom: "834479572@qq.com", // 设置固定的发送者邮箱地址
			SmtpHost:  "smtp.qq.com",
			SmtpPort:  587,
			SmtpUser:  "834479572@qq.com",
			SmtpPass:  "bmqdkfqwluctbefh",
		}
		// 生成验证码
		Vcode := utils.RandomNcode(6)
		// 发送验证码
		//err := utils.SendCode(data.Email, Vcode, config)
		err := utils.SendVerificationCode(data.Email, Vcode, config)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "验证码发送失败",
			})
			return
		}
		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "验证码发送成功",
			"data": Vcode,
		})
	})
	//注册
	r.POST("/api/user/register", func(c *gin.Context) {
		type regdata struct {
			Username string `json:"username"`
			Pwd      string `json:"password"`
			Email    string `json:"email"`
>>>>>>> Stashed changes
		}

		var data regdata
		fmt.Println("Username:", data.Username, "Pwd:", data.Pwd, "Email:", data.Email)
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
<<<<<<< Updated upstream
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
=======
		//验证用户是否已注册
		if controlsql.UserExists_User(db, data.Username) {
			c.JSON(http.StatusConflict, gin.H{
				"code": "409",
				"msg":  "用户已注册",
			})
			return
>>>>>>> Stashed changes
		}
		Key := "123456781234567812345678" //密钥
		cryptoPwd := utils.AesEncrypt(data.Pwd, Key)
		//注册用户
		err := controlsql.RegisterUser_User(db, data.Username, cryptoPwd, data.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "注册成功",
		})
	})
<<<<<<< Updated upstream

	r.GET("/api/team_manager/login", func(c *gin.Context) {
		c.File("/static/team_manager/login.html")
	})

	//重定向至登录页面
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/static/team_manager/login.html")
	})
	//登录接口
	r.POST("/api/team_manager/login", func(c *gin.Context) {
		type logdata struct {
=======
	//登录
	r.POST("/api/user/login", func(c *gin.Context) {
		type logindata struct {
>>>>>>> Stashed changes
			Username string `json:"username"`
			Pwd      string `json:"password"`
		}
		var data logindata
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  "400",
				"error": err.Error(),
			})
			return
		}
<<<<<<< Updated upstream

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
=======
		//验证用户是否存在
		if !controlsql.UserExists_User(db, data.Username) {
			c.JSON(403, gin.H{
				"code": "403",
				"msg":  "用户不存在",
			})
			return
		}
		//验证密码是否正确
		isMatch := controlsql.CheckUser_User(db, data.Username, data.Pwd)
		if !isMatch {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "密码错误",
>>>>>>> Stashed changes
			})
			return
		}
<<<<<<< Updated upstream
=======
		//生成token
		userid, err := controlsql.GetUserID(db, data.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}
		team_id, team_name, err := controlsql.GetTokenParams_User(db, userid)

		if err != nil && err.Error() != "sql: no rows in result set" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}

		token, _ := utils.GenerateToken_User(userid, team_id, team_name)
		c.JSON(http.StatusOK, gin.H{
			"code":  "200",
			"msg":   "登录成功",
			"token": token,
		})
>>>>>>> Stashed changes
	})
}
