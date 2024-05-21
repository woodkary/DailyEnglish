package userrouter

import (
	controlsql "DailyEnglish/db"
	utils "DailyEnglish/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		}

		var data regdata
		fmt.Println("Username:", data.Username, "Pwd:", data.Pwd, "Email:", data.Email)
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		//验证用户是否已注册
		if controlsql.UserExists_User(db, data.Username) {
			c.JSON(http.StatusConflict, gin.H{
				"code": "409",
				"msg":  "用户已注册",
			})
			return
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
	//登录
	r.POST("/api/user/login", func(c *gin.Context) {
		type logindata struct {
			Username string `json:"username"`
			Pwd      string `json:"password"`
		}
		var data logindata
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
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
			})
			return
		}
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
	})
}
