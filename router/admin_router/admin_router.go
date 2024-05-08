package adminrouter

import (
	controlsql "DailyEnglish/db"
	utils "DailyEnglish/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAdminRouter(r *gin.Engine, db *sql.DB) {

	//注册&登录页面
	r.GET("/api/team_manager/login", func(c *gin.Context) {
		c.File("/static/team_manager/login&register.html")
	})

	//重定向至注册&登录页面
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/static/team_manager/login&register.html")
	})

	// 发送验证码接口
	r.POST("/api/team_manager/send_code", func(c *gin.Context) {
		// 解析 JSON 数据
		type response struct {
			Email string `json:"email"`
		}
		var data response
		if err := c.ShouldBindJSON(&data); err != nil {
			fmt.Print(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		// 验证邮箱是否已注册
		if controlsql.EmailIsRegistered(db, data.Email) {
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

	//注册接口
	r.POST("/api/team_manager/register", func(c *gin.Context) {
		type regdata struct {
			Username string `json:"username"`
			Pwd      string `json:"pwd"`
			Email    string `json:"email"`
		}
		var data regdata
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		if !controlsql.UserExists(db, data.Username) {
			//验证码由前端完成判定

			Key := "123456781234567812345678" //密钥
			cryptoPwd := utils.AesEncrypt(data.Pwd, Key)
			//获取系统当前日期
			//RegisterDate := utils.GetCurrentDate()

			if controlsql.RegisterUser(db, data.Username, data.Email, cryptoPwd, "10086") != nil {
				c.JSON(http.StatusOK, gin.H{
					"code": "200",
					"msg":  "注册成功",
				})
			} else {
				c.JSON(http.StatusConflict, gin.H{
					"code": "500",
					"msg":  "服务器内部错误",
				})
			}
		} else {
			c.JSON(http.StatusConflict, gin.H{
				"code": "409",
				"msg":  "用户名已被注册",
			})
		}
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
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}

		if !controlsql.UserExists(db, data.Username) {
			c.JSON(http.StatusForbidden, gin.H{
				"code": "403",
				"msg":  "用户不存在",
			})
		} else if controlsql.CheckUser(db, data.Username, data.Pwd) {
			item1, item2s, err := controlsql.GetTokenParams(db, data.Username)
			//输出所有teamid
			for _, item2 := range item2s {
				fmt.Println(item2)
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": "500",
					"msg":  "服务器内部错误",
				})
				return
			}

			//生成token
			token, err := utils.GenerateToken(item1, item2s)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": "500",
					"msg":  "服务器错误",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":  "200",
				"msg":   "登录成功",
				"token": token,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "用户名或密码有误",
			})
		}
	})

}
