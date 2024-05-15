package userrouter

import (
	controlsql "DailyEnglish/db"
	utils "DailyEnglish/utils"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.Engine, db *sql.DB) {
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

}
