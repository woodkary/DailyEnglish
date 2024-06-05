package adminrouter

import (
	controlsql "DailyEnglish/db"
	utils "DailyEnglish/utils"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func InitAdminRouter(r *gin.Engine, db *sql.DB, rdb *redis.Client) {

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
		type request struct {
			Email string `json:"email"`
		}
		var data request
		if err := c.ShouldBindJSON(&data); err != nil {
			fmt.Print(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		// 验证邮箱是否已注册
		if controlsql.EmailIsRegistered_TeamManager(db, data.Email) {
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
		// 将验证码存入 Redis
		ctx := context.Background()// 创建一个空的 context
		key := fmt.Sprintf("%s:%s", "web", data.Email)// key前缀为web:邮箱
		err = rdb.Set(ctx, key, Vcode, time.Minute*5).Err() // 验证码有效期5分钟,更新时替换
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "验证码存储失败",
			})
			return
		}

		// 返回成功响应
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "验证码发送成功",
			"data": Vcode,
			"email": data.Email,
		})
	})

	//注册接口
	r.POST("/api/team_manager/register", func(c *gin.Context) {
		type regdata struct {
			Username string `json:"username"`
			Pwd      string `json:"password"`
			Email    string `json:"email"`
			Code     string `json:"code"`
		}

		var data regdata
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		//验证验证码
		ctx := context.Background()
		key := fmt.Sprintf("%s:%s", "web", data.Email)
		code, err := rdb.Get(ctx, key).Result()
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{
				"code": "403",
				"msg":  "验证码已过期",
			})
			return
		}
		if code != data.Code {
			c.JSON(http.StatusForbidden, gin.H{
				"code": "403",
				"msg":  "验证码错误",
			})
			return
		}
		// 验证码验证成功后尝试删除验证码，即使删除失败也不会影响流程
        err = rdb.Del(ctx, key).Err()
        if err != nil {
            fmt.Printf("删除验证码失败：%v\n", err)
        }
		
		//验证用户名是否已被注册
		if !controlsql.AdminManagerExists(db, data.Username) {
			
			// fmt.Println("Pwd:", data.Pwd)
			Key := "123456781234567812345678" //密钥
			cryptoPwd := utils.AesEncrypt(data.Pwd, Key)
			// fmt.Println("cryptoPwd:", cryptoPwd)
			//获取系统当前日期
			//RegisterDate := utils.GetCurrentDate()

			if controlsql.RegisterUser(db, data.Username, data.Email, cryptoPwd, "10086") == nil {
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

		if !controlsql.AdminManagerExists(db, data.Username) {
			c.JSON(http.StatusForbidden, gin.H{
				"code": "403",
				"msg":  "用户不存在",
			})
		} else if controlsql.CheckTeamManager(db, data.Username, data.Pwd) {
			item1, item2s, err := controlsql.GetTokenParams_TeamManager(db, data.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": "500",
					"msg":  "服务器内部错误",
				})
				return
			}
			fmt.Println(item2s)
			//生成token
			token, err := utils.GenerateToken_TeamManager(item1, item2s)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": "500",
					"msg":  "服务器错误",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":      "200",
				"msg":       "登录成功",
				"token":     token,
				"team_info": item2s,
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "用户名或密码有误",
			})
		}
	})
}
