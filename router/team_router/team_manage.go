package teamrouter

import (
	service "DailyEnglish/Service"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Team_manager(r *gin.Engine, client *redis.Client, db *sql.DB) {
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
		//下面给定死的
		response.Code = "200"
		response.Msg = "成功"
		response.Punch.Punched = 45
		response.Punch.PunchNum = []string{"45", "40", "37", "42", "48", "22", "46"}
		response.Punch.PunchRate = []string{"0.9", "0.8", "0.7", "0.8", "0.96", "0.5", "0.93"}
		Item1 := struct {
			Name      string `json:"name"`
			PunchRate string `json:"punch_rate"`
			PunchDay  string `json:"punch_day"`
		}{"OTTO", "87%", "666"}
		response.Punch.PunchLB = append(response.Punch.PunchLB, Item1)
		response.Punch.PunchLB = append(response.Punch.PunchLB, Item1)
		response.Punch.PunchLB = append(response.Punch.PunchLB, Item1)
		response.Punch.PunchLB = append(response.Punch.PunchLB, Item1)
		response.Punch.PunchLB = append(response.Punch.PunchLB, Item1)
		response.Punch.PunchLB = append(response.Punch.PunchLB, Item1)
		response.Exam.Time = "三月三十五日"
		Item2 := struct {
			Name      string `json:"name"`
			ExamRank  string `json:"exam_rank"`
			ExamScore string `json:"exam_score"`
		}{"OTTO", "第一名", "99"}
		response.Exam.ExamLB = append(response.Exam.ExamLB, Item2)
		Item2.ExamRank = "第二名"
		Item2.ExamScore = "97"
		response.Exam.ExamLB = append(response.Exam.ExamLB, Item2)
		Item2.ExamRank = "第三名"
		Item2.ExamScore = "87"
		response.Exam.ExamLB = append(response.Exam.ExamLB, Item2)
		Item2.ExamRank = "第四名"
		Item2.ExamScore = "83"
		response.Exam.ExamLB = append(response.Exam.ExamLB, Item2)
		Item2.ExamRank = "第五名"
		Item2.ExamScore = "74"
		response.Exam.ExamLB = append(response.Exam.ExamLB, Item2)
		Item2.ExamRank = "第六名"
		Item2.ExamScore = "55"
		response.Exam.ExamLB = append(response.Exam.ExamLB, Item2)
		response.Notice.NoticeJoin = 5
		Item3 := struct {
			NoticeData string `json:"notice_data"`
			NoticeTime string `json:"notice_time"`
		}{"您有新增的团队加入申请，请及时审核", "1分钟前"}
		response.Notice.NoticeRecent = append(response.Notice.NoticeRecent, Item3)
		response.Notice.NoticeRecent = append(response.Notice.NoticeRecent, Item3)
		response.Notice.NoticeRecent = append(response.Notice.NoticeRecent, Item3)
		response.Notice.NoticeRecent = append(response.Notice.NoticeRecent, Item3)
		response.Notice.NoticeRecent = append(response.Notice.NoticeRecent, Item3)
		response.Notice.NoticeRecent = append(response.Notice.NoticeRecent, Item3)
		c.JSON(200, response)
	})
	r.GET("/api/team_manage/exam_situation/data", func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/static/team_manager/login.html")
		}
		_, err := service.ParseToken(token)
		if err != nil {
			c.JSON(400, "登录信息无效或过期")
		}

	})
}
