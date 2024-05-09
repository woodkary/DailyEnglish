package teamrouter

import (
	controlsql "DailyEnglish/db"
	service "DailyEnglish/utils"
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func tokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, "未提供令牌")
			c.Abort()
			return
		}
		// 检查头是否以"Bearer"开头
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, "令牌格式错误")
			c.Abort()
			return
		}
		// 提取令牌
		token := strings.TrimPrefix(authHeader, "Bearer ")
		user, err := service.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "令牌无效")
			c.Abort()
			return
		}

		// 将用户信息存储在context中，后续的处理器可以使用
		c.Set("user", user)
		c.Next()
	}
}

func InitTeamRouter(r *gin.Engine, db *sql.DB) {
	//考试情况数据
	r.POST("/api/team_manage/exam_situation/calendar", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			Year  string `json:"year"`  // 年份
			Month string `json:"month"` // 月份
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
			return
		}
		requestYear, err := strconv.Atoi(request.Year)
		if err != nil {
			log.Println("Error parsing year:", err)
		}

		requestMonth, err := strconv.Atoi(request.Month)
		if err != nil {
			log.Println("Error parsing month:", err)

		}
		user, _ := c.Get("user")
		userClaims, ok := user.(*service.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		//TODO这里是查询数据库获取数据
		var Item []controlsql.ExamInfo
		for _, teamID := range userClaims.TeamID {
			examInfo, err := controlsql.SearchExamInfoByTeamID(db, teamID)
			if err != nil {
				c.JSON(500, "服务器错误")
				log.Panic(err)
				return
			}
			Item = append(Item, examInfo...)
		}

		// ExamsResponse 结构体表示包含多个考试的响应
		type response struct {
			Code      string   `json:"code"`      // 响应代码
			Msg       string   `json:"msg"`       // 响应消息
			Exam_date []string `json:"exam_date"` // 有考试的日期
		}
		var Response response

		//TODO 将查询到的考试信息转换为响应的结构体
		for _, exam := range Item {
			examDate, err := time.Parse("2006-01-02", exam.ExamDate)
			if err != nil {
				log.Println("Error parsing date:", err)
				continue
			}

			if examDate.Year() == requestYear && examDate.Month() == time.Month(requestMonth) {
				Response.Exam_date = append(Response.Exam_date, exam.ExamDate)
			}
		}
		Response.Code = "200"
		Response.Msg = "成功"
		c.JSON(200, Response)
	})
	//获取某日管理的团队的所有考试信息
	r.POST("/api/team_manage/exam_situation/exam_date", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			Date string `json:"date"` // 日期
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
			return
		}
		user, _ := c.Get("user")
		userClaims, ok := user.(*service.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		//TODO这里是查询数据库获取数据
		var Item [][]controlsql.ExamInfo
		for _, teamID := range userClaims.TeamID {
			examInfo, err := controlsql.SearchExamInfoByTeamIDAndDate(db, teamID, request.Date)
			if err != nil {
				c.JSON(500, "服务器错误")
				log.Panic(err)
				return

			}
			Item = append(Item, examInfo)
		}
		// ExamsResponse 结构体表示包含多个考试的响应
		type response struct {
			Code  string `json:"code"` // 响应代码
			Msg   string `json:"msg"`  // 响应消息
			Exams []struct {
				TeamName string `json:"team_name"` // 团队名称
				TeamID   string `json:"team_id"`   // 团队ID
				ExamID   string `json:"exam_id"`   // 考试ID
				ExamName string `json:"exam_name"` // 考试名称
			} `json:"exams"` // 考试列表
		}
		var Response response
		Response.Code = "200"
		Response.Msg = "成功"
		for i, items := range Item {
			for _, exam := range items {
				var examInfo struct {
					TeamName string `json:"team_name"`
					TeamID   string `json:"team_id"`
					ExamID   string `json:"exam_id"`
					ExamName string `json:"exam_name"`
				}
				teamname, err := controlsql.SearchTeamNameByTeamID(db, userClaims.TeamID[i])
				if err != nil {
					c.JSON(500, "服务器错误")
					log.Panic(err)
					return
				}
				examInfo.TeamID = strconv.Itoa(userClaims.TeamID[i])
				examInfo.TeamName = teamname
				examInfo.ExamID = strconv.Itoa(exam.ExamID)
				examInfo.ExamName = exam.ExamName
				Response.Exams = append(Response.Exams, examInfo)
			}
		}
		c.JSON(200, Response)
	})
	//获取单次考试详情
	r.POST("/api/team_manage/exam_situation/exam_detail", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			ExamID int `json:"exam_id"` // 考试名称
			TeamID int `json:"team_id"` // 团队名称
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
			return
		}

		ScoresInExam, err := controlsql.SearchExamScoreByExamID(db, request.ExamID)
		if err != nil {
			c.JSON(500, "服务器错误1")
			log.Panic(err)
			return
		}
		levelNums := service.CalculateUserLevel(ScoresInExam)

		type UserResult struct {
			Username string `json:"username"` // 用户名
			Score    int    `json:"score"`    // 得分
			Progress int    `json:"progress"` // 进步名次 (相距上次)
		}

		userIDs, err := controlsql.SearchUserIDByTeamID(db, request.TeamID)
		if err != nil {
			c.JSON(500, "服务器错误2")
			log.Panic(err)
			return
		}

		userres := make([]UserResult, 0)
		for _, userID := range userIDs {
			item1, item2, item3, err := controlsql.SearchClosestExamByTeamIDAndExamID(db, request.TeamID, request.ExamID, userID)
			if err != nil {
				c.JSON(500, "服务器错误3")
				log.Panic(err)
				return
			}
			userres = append(userres, UserResult{Username: item1, Score: item2, Progress: item3})
		}

		QuestionNum, err := controlsql.SearchQuestionNumByExamID(db, request.ExamID) // 考试题目数量
		if err != nil {
			c.JSON(500, "服务器错误4")
			log.Panic(err)
			return
		}

		var qd = make([][]int, QuestionNum)                                  // 考试题目详情
		qid, err := controlsql.SearchQuestionIDsByExamID(db, request.ExamID) // 考试题目ID
		if err != nil {
			c.JSON(500, "服务器错误")
			log.Panic(err)
			return
		}

		for i := 0; i < QuestionNum; i++ {
			qd[i], err = controlsql.SearchQuestionStatistics(db, request.ExamID, qid[i])
			if err != nil {
				c.JSON(500, "服务器错误")
				log.Panic(err)
				return
			}
		}

		type ExamDetail struct {
			ID             string       `json:"exam_id"`          // 考试ID
			Name           string       `json:"exam_name"`        // 考试名称
			UserLevels     []int        `json:"user_levels"`      // 用户等级
			QuestionDetail [][]int      `json:"question_details"` // 考试题目详情
			UserResult     []UserResult `json:"user_result"`      // 考试参与人员得分情况
		}
		ExamName, err := controlsql.SearchExamNameByExamID(db, request.ExamID)
		if err != nil {
			c.JSON(500, "服务器错误")
			log.Panic(err)
			return
		}

		type response struct {
			Code       string     `json:"code"`        // 状态码
			Msg        string     `json:"msg"`         // 消息
			ExamDetail ExamDetail `json:"exam_detail"` // 考试详情
		}
		var Response response
		Response.Code = "200"
		Response.Msg = "成功"
		Response.ExamDetail.ID = strconv.Itoa(request.ExamID)
		Response.ExamDetail.Name = ExamName
		Response.ExamDetail.UserLevels = levelNums[:]
		Response.ExamDetail.QuestionDetail = qd
		Response.ExamDetail.UserResult = make([]UserResult, 0)
		Response.ExamDetail.UserResult = append(Response.ExamDetail.UserResult, userres...)
		c.JSON(200, Response)
	})

	//成员管理页面
	r.GET("/api/team_manage/member_manage/data", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		userClaims, ok := user.(*service.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		type Member struct {
			ID    int    `json:"id"`    // 成员ID
			Name  string `json:"name"`  // 成员姓名
			Phone string `json:"phone"` // 成员手机号
			Email string `json:"email"` // 成员邮箱
		}
		type team struct {
			TeamName string   `json:"team_name"` // 团队名
			TeamID   int      `json:"team_id"`   // 团队ID
			Members  []Member `json:"members"`   // 成员列表
		}
		type response struct {
			Code string `json:"code"` // 状态码
			Msg  string `json:"msg"`  // 消息
			Team []team `json:"team"` // 团队列表
		}
		var Response response
		for _, teamID := range userClaims.TeamID {
			var Team team
			Team.TeamID = teamID
			Team.TeamName, _ = controlsql.SearchTeamNameByTeamID(db, teamID)
			users, err := controlsql.SearchUserIDByTeamID(db, teamID)
			if err != nil {
				c.JSON(500, "服务器错误")
				return
			}
			for _, userID := range users {
				var Member Member
				Member.ID = userID
				Member.Name, Member.Phone, Member.Email, err = controlsql.SearchUserNameAndPhoneByUserID(db, userID)
				if err != nil {
					c.JSON(500, "服务器错误")
					return
				}
				Team.Members = append(Team.Members, Member)
			}
			Response.Team = append(Response.Team, Team)
		}
		Response.Code = "200"
		Response.Msg = "成功"
		c.JSON(200, Response)
	})
	//成员删除
	r.POST("/api/team_manage/member_manage/delete", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			TeamID int `json:"team_id"` // 要删除的成员的用户名
			UserID int `json:"user_id"` // 团队名
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
		}
		err := controlsql.DeleteUserTeamByUserIDAndTeamID(db, request.TeamID, request.UserID)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		c.JSON(200, "删除成功")
	})
	//获取个人中心界面所需信息
	r.GET("/api/team_manage/personal_center/data", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		userClaims, ok := user.(*service.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}

		type Team struct {
			TeamName  string `json:"team_name"`  // 团队名
			TeamID    int    `json:"team_id"`    // 团队ID
			MemberNum int    `json:"member_num"` // 成员数量
		}
		// Response 定义了响应的信息
		type Response struct {
			Code     string `json:"code"`     // 状态码
			Msg      string `json:"msg"`      // 消息
			Name     string `json:"name"`     // 用户名
			Phone    string `json:"phone"`    // 手机号
			Partment string `json:"partment"` // 部门
			Email    string `json:"email"`    // 邮箱
			Team     []Team `json:"team"`     // 团队列表
		}
		var response Response
		// 查询用户信息
		ManageInfo, err := controlsql.SearchManagerInfoByManagerID(db, userClaims.UserID)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
		}
		response.Name = ManageInfo.ManagerName
		response.Phone = ManageInfo.ManagerPhone
		response.Partment = ManageInfo.ManagerPartment
		response.Email = ManageInfo.ManagerEmail
		// 查询团队信息
		for _, teamID := range userClaims.TeamID {
			var team Team
			team.TeamID = teamID
			team.TeamName, team.MemberNum, err = controlsql.SearchTeamInfoByTeamID(db, teamID)
			if err != nil {
				c.JSON(500, "服务器错误")
				return
			}
			response.Team = append(response.Team, team)
		}
		response.Code = "200"
		response.Msg = "成功"
		c.JSON(200, response)
	})

}
