package teamrouter

import (
	"DailyEnglish/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Team_manager(r *gin.Engine) {
	r.GET("/api/team_manage/index/data", func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/static/team_manager/login.html")
		}
		_, err := service.ParseToken(token)
		if err != nil {
			c.JSON(400, "登录信息无效或过期")
		}
		// 定义JSON响应的结构体
		type Response struct {
			Code  string `json:"code"`
			Msg   string `json:"msg"`
			Punch struct {
				Punched   int      `json:"punched"`
				PunchNum  []string `json:"punch_num"`
				PunchRate []string `json:"punch_rate"`
				PunchLB   []struct {
					Name      string `json:"name"`
					PunchRate string `json:"punch_rate"`
					PunchDay  string `json:"punch_day"`
				} `json:"punch_LB"`
			} `json:"Punch"`
			Exam struct {
				Time   string `json:"time"`
				ExamLB []struct {
					Name      string `json:"name"`
					ExamRank  string `json:"exam_rank"`
					ExamScore string `json:"exam_score"`
				} `json:"exam_LB"`
			} `json:"exam"`
			Notice struct {
				NoticeJoin   int `json:"notice_join"`
				NoticeRecent []struct {
					NoticeData string `json:"notice_data"`
					NoticeTime string `json:"notice_time"`
				} `json:"notice_recent"`
			} `json:"notice"`
		}
		var response Response
		//@TODO 添加数据库查找数据放入response中的逻辑
		c.JSON(200, response)
	})
}
