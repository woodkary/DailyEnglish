package userrouter

import (
	controlsql "DailyEnglish/Control_SQL"
	service "DailyEnglish/services"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-redis/redis"

	"github.com/gin-gonic/gin"
)

func User_manager(r *gin.Engine, client *redis.Client, db *sql.DB) {

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
		c.File("/static/team_manager/login&register.html")
	})

	//重定向至登录页面
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/static/team_manager/login&register.html")
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
}
