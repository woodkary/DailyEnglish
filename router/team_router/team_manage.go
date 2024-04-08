package teamrouter

import (
	controlsql "DailyEnglish/Control_SQL"
	service "DailyEnglish/Service"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func Team_manager(r *gin.Engine, client *redis.Client, db *sql.DB) {
	//主页数据
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
	//考试情况数据
	r.GET("/api/team_manage/exam_situation/data", func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/static/team_manager/login.html")
		}
		user, err := service.ParseToken(token)
		if err != nil {
			c.JSON(400, "登录信息无效或过期")
		}

		Item, err := controlsql.QueryTeamExams(client, user.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		// ExamInfo 结构体表示单个考试的信息
		type ExamInfo struct {
			Name         string `json:"name"`          // 考试名称
			Time         string `json:"time"`          // 考试时间
			FullScore    string `json:"full_score"`    // 满分
			AverageScore string `json:"average_score"` // 平均分
			PassRate     string `json:"pass_rate"`     // 通过率
		}

		// ExamsResponse 结构体表示包含多个考试的响应
		type response struct {
			Code  string     `json:"code"`  // 响应代码
			Msg   string     `json:"msg"`   // 响应消息
			Exams []ExamInfo `json:"exams"` // 考试列表
		}
		var Response response
		var examinfo ExamInfo
		for _, item := range Item {
			examinfo.Name = item["Name"]
			examinfo.FullScore = "100"
			examinfo.AverageScore = item["AverageScore"]
			examinfo.PassRate = item["PassRate"]
			examinfo.Time = item["Date"]
			Response.Exams = append(Response.Exams, examinfo)
		}
		Response.Code = "200"
		Response.Msg = "成功"
		c.JSON(200, Response)
	})
	//打卡详情界面
	r.GET("/api/team_manage/punch_statistics/data", func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/static/team_manager/login.html")
		}
		user, err := service.ParseToken(token)
		if err != nil {
			c.JSON(400, "登录信息无效或过期")
		}

		Item1, err := controlsql.GetTeamMembersAttendance1(client, user.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		Item2, err := controlsql.GetTeamMembersAttendanceByDate(client, user.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		Item3, err := controlsql.GetTeamMembersAttendanceNum(client, user.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		type punchStatistic struct {
			Name      string `json:"name"`
			IsPunched string `json:"ispunched"`
			PunchDay  string `json:"punch_day"`
			PunchWord string `json:"punch_word"`
			PunchRate string `json:"punch_rate"`
		}
		type response struct {
			Code             string           `json:"code"` // 响应代码
			Msg              string           `json:"msg"`  // 响应消息
			Punch_statistics []punchStatistic `json:"punch_statistics"`
		}
		var Response response
		var punchstatistic punchStatistic
		Response.Code = "200"
		Response.Msg = "成功"
		for _, member := range Item1 {
			punchstatistic.Name = member.Username
			punchstatistic.PunchDay = strconv.Itoa(member.AttendanceDays)
			punchstatistic.PunchRate = member.AttendanceRate
			if Item2[member.Username] == 1 {
				punchstatistic.IsPunched = "是"
			} else {
				punchstatistic.IsPunched = "否"
			}
			punchstatistic.PunchWord = strconv.Itoa(Item3[member.Username])
			Response.Punch_statistics = append(Response.Punch_statistics, punchstatistic)
		}
		c.JSON(200, Response)
	})
	//成员管理页面
	r.GET("/api/team_manage/member_manage/data", func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/static/team_manager/login.html")
		}
		user, err := service.ParseToken(token)
		if err != nil {
			c.JSON(400, "登录信息无效或过期")
		}

		Item, err := controlsql.GetTeamInfo(client, user.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		type Member struct {
			Name     string `json:"name"`      // 成员姓名
			Right    string `json:"right"`     // 成员权限
			Time     string `json:"time"`      // 时间
			PunchDay string `json:"punch_day"` // 打卡天数
		}
		type response struct {
			Code    string   `json:"code"`    // 状态码
			Msg     string   `json:"msg"`     // 消息
			Members []Member `json:"members"` // 团队成员列表
		}
		var Response response
		for _, m := range Item.Members {
			var member Member
			member.Name = m.Username
			member.PunchDay = strconv.Itoa(m.AttendanceDays)
			member.Time = m.JoinDate
			if m.IsAdmin {
				member.Right = "管理员"
			} else {
				member.Right = "成员"
			}
			Response.Members = append(Response.Members, member)
		}
		Response.Code = "200"
		Response.Msg = "成功"
		c.JSON(200, Response)
	})
	//获取审核申请页面信息
	r.GET("/api/team_manage/request_manage/data", func(c *gin.Context) {

	})
}
