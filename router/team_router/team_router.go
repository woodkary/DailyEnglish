package teamrouter

import (
	controlsql "DailyEnglish/db"
	middlewares "DailyEnglish/middlewares"
	utils "DailyEnglish/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func tokenAuthMiddleware() gin.HandlerFunc {
	return middlewares.TokenAuthMiddleware("TeamManager")
}

func InitTeamRouter(r *gin.Engine, db *sql.DB, rdb *redis.Client, es *elasticsearch.Client) {
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

		yyyy, err := strconv.Atoi(request.Year)
		if err != nil {
			c.JSON(400, "请求参数错误")
			return
		}
		mm, err := strconv.Atoi(request.Month)
		if err != nil {
			c.JSON(400, "请求参数错误")
			return
		}
		user, _ := c.Get("user")
		TeamManagerClaims, ok := user.(*utils.TeamManagerClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		// 查询每个团队
		var Item []controlsql.ExamInfo
		for teamID := range TeamManagerClaims.Team {
			// 查询该团队所有考试信息 包括ID Name Date
			examInfo, err := controlsql.SearchExamInfoByTeamID(db, teamID)
			if err != nil {
				c.JSON(500, "服务器错误")
				log.Panic(err)
				return
			}
			Item = append(Item, examInfo...)
		}

		// fmt.Println(Item)
		// fmt.Println("year = ", yyyy, "month = ", mm)

		type response struct {
			Code      string   `json:"code"`      // 响应代码
			Msg       string   `json:"msg"`       // 响应消息
			Exam_date []string `json:"exam_date"` // 有考试的日期
		}
		var Response response
		Response.Exam_date = make([]string, 0)

		// 找到所有团队中所有考试时间为request所给参数的考试对应的日期
		for _, exam := range Item {
			examDate, err := time.Parse("2006-01-02", exam.ExamDate)
			if err != nil {
				log.Println("Error parsing date:", err)
				continue
			}
			// fmt.Println("now parsing: ", examDate.Year(), examDate.Month())

			if examDate.Year() == yyyy && examDate.Month() == time.Month(mm) {
				Response.Exam_date = append(Response.Exam_date, exam.ExamDate)
			}
		}
		// fmt.Println(Response.Exam_date)
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
		TeamManagerClaims, ok := user.(*utils.TeamManagerClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		Item := make(map[int][]controlsql.ExamInfo)
		for teamID := range TeamManagerClaims.Team {
			examInfo, err := controlsql.SearchExamInfoByTeamIDAndDate(db, teamID, request.Date)
			if err != nil {
				c.JSON(500, "服务器错误")
				log.Panic(err)
				return

			}
			Item[teamID] = examInfo
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
		for team_id, items := range Item {
			for _, exam := range items {
				var examInfo struct {
					TeamName string `json:"team_name"`
					TeamID   string `json:"team_id"`
					ExamID   string `json:"exam_id"`
					ExamName string `json:"exam_name"`
				}
				teamname := TeamManagerClaims.Team[team_id]
				examInfo.TeamID = strconv.Itoa(team_id)
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
		levelNums := utils.CalculateUserLevel(ScoresInExam)

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
			c.JSON(500, "服务器错误5")
			log.Panic(err)
			return
		}
		fmt.Println(qid)

		for i := 0; i < QuestionNum; i++ {
			qd[i], err = controlsql.SearchQuestionStatistics(db, request.ExamID, qid[i])
			if err != nil {
				c.JSON(500, "服务器错误6")
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
			Code       int        `json:"code"`        // 状态码
			Msg        string     `json:"msg"`         // 消息
			ExamDetail ExamDetail `json:"exam_detail"` // 考试详情
		}
		var Response response
		Response.Code = 200
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
		TeamManagerClaims, ok := user.(*utils.TeamManagerClaims) // 将 user 转换为 *TeamManagerClaims 类型
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
		for teamID, teamname := range TeamManagerClaims.Team {
			var Team team
			Team.TeamID = teamID
			Team.TeamName = teamname
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
		TeamManagerClaims, ok := user.(*utils.TeamManagerClaims) // 将 user 转换为 *UserClaims 类型
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
		ManageInfo, err := controlsql.SearchManagerInfoByManagerID(db, TeamManagerClaims.ManagerID)
		if err != nil {
			log.Panic(err)
			c.JSON(500, "服务器错误")
			return
		}
		response.Name = ManageInfo.ManagerName
		response.Phone = ManageInfo.ManagerPhone
		response.Partment = ManageInfo.ManagerPartment
		response.Email = ManageInfo.ManagerEmail
		// //输入manager_id，返回manager_Id,teamMap
		// managerId, teams, err := controlsql.GetTokenParamsByManagerId(db, TeamManagerClaims.ManagerID)
		// if err != nil {
		// 	log.Panic(err)
		// 	c.JSON(500, "服务器错误")
		// 	return
		// }
		// //根据manager_id和teamMap，生成新的token
		// newToken, err := utils.GenerateToken_TeamManager(managerId, teams)
		// if err != nil {
		// 	log.Panic(err)
		// 	c.JSON(500, "服务器错误")
		// 	return
		// }
		// 查询团队信息，主要为了获取团队人数
		for teamID := range TeamManagerClaims.Team {
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
	// 刷新团队码
	r.POST("/api/team_manage/refresh_team_code", tokenAuthMiddleware(), func(c *gin.Context) {
		fmt.Println("刷新团队码")
		type Request struct {
			TeamID int `json:"team_id"` // 团队ID
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
			return
		}
		fmt.Println("获取的团队ID为:", request.TeamID)
		NewCode := utils.EncryptIC(request.TeamID, 114514)
		fmt.Println("获取的邀请码为:", NewCode)
		var response struct {
			Code           string `json:"code"`            // 状态码
			Msg            string `json:"msg"`             // 消息
			InvitationCode string `json:"invitation_code"` // 团队码
		}
		response.Code = "200"
		response.Msg = "成功"
		response.InvitationCode = NewCode
		c.JSON(200, response)
	})

	// 获取考试题目
	r.POST("/api/team_manage/new_exam/all_questions", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			Index int `json:"index"` // 要获取的题目的索引
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "请求参数错误"})
			return
		}
		type Question struct {
			QuestionID         int               `json:"question_id"`
			QuestionType       string            `json:"question_type"`
			QuestionDifficulty string            `json:"question_difficulty"`
			QuestionGrade      string            `json:"question_grade"`
			QuestionContent    string            `json:"question_content"`
			QuestionChoices    map[string]string `json:"question_choices"`
			QuestionAnswer     string            `json:"question_answer"`
			FullScore          int               `json:"full_score"`
		}
		type response struct {
			Code      int        `json:"code"`      // 状态码
			Msg       string     `json:"msg"`       // 消息
			Questions []Question `json:"questions"` // 题目列表
		}
		var Response response
		var QuestionTypeDict = map[int]string{
			1: "单选题",
			2: "填空题",
			3: "写作题",
			4: "填空题",
			5: "简答题",
		}
		var gradeDescriptions = map[int]string{
			1:  "小学一年级",
			2:  "小学二年级",
			3:  "小学三年级",
			4:  "小学四年级",
			5:  "小学五年级",
			6:  "小学六年级",
			7:  "初中一年级",
			8:  "初中二年级",
			9:  "初中三年级",
			10: "高中一年级",
			11: "高中二年级",
			12: "高中三年级",
			13: "四级",
			14: "六级",
		}
		var difficultyDescriptions = map[int]string{
			1: "容易",
			2: "中等",
			3: "困难",
		}

		// 获取需要查询的题目ID
		questionIDs := make([]int, 0, 50)
		for i := request.Index; i < request.Index+50; i++ {
			questionIDs = append(questionIDs, i)
		}

		// 批量从Elasticsearch获取题目信息
		questionsInfo, err := controlsql.GetQuestionsInfoFromES(es, questionIDs)
		if err != nil {
			log.Printf("Error getting questions info from ES: %s", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "服务器错误",
			})
			return
		}

		// 如果Elasticsearch中没有找到所有需要的题目，则从MySQL中查询并插入到Elasticsearch
		if len(questionsInfo) < len(questionIDs) {
			// 找到缺失的题目ID
			foundIDs := make(map[int]bool)
			for _, q := range questionsInfo {
				foundIDs[q.Question_id] = true
			}
			missingIDs := make([]int, 0)
			for _, id := range questionIDs {
				if !foundIDs[id] {
					missingIDs = append(missingIDs, id)
				}
			}

			// 从MySQL中查询缺失的题目
			missingQuestionsInfo, err := controlsql.GetQuestionsInfoFromDB(db, missingIDs)
			if err != nil {
				log.Printf("Error getting questions info from DB: %s", err)
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器错误"})
				return
			}

			// 将缺失的题目插入到Elasticsearch
			if err := controlsql.StoreQuestionsInfoToES(es, missingQuestionsInfo); err != nil {
				log.Printf("Error storing questions info to ES: %s", err)
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器错误"})
				return
			}

			// 合并结果
			questionsInfo = append(questionsInfo, missingQuestionsInfo...)
		}

		// 构建响应数据
		for _, question := range questionsInfo {
			var q Question
			q.QuestionID = question.Question_id
			q.QuestionType = QuestionTypeDict[question.Questiontype]
			q.QuestionDifficulty = difficultyDescriptions[question.QuestionDifficulty]
			q.QuestionGrade = gradeDescriptions[question.QuestionGrade]
			q.QuestionContent = question.QuestionContent
			q.QuestionChoices = question.Options
			q.QuestionAnswer = question.QuestionAnswer
			q.FullScore = 5
			Response.Questions = append(Response.Questions, q)
		}

		Response.Code = 200
		Response.Msg = "成功"
		c.JSON(http.StatusOK, Response)
	})
	//发布考试
	r.POST("/api/team_manage/new_exam", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			ExamName    string `json:"exam_name"`    // 考试名称
			ExamDate    string `json:"exam_date"`    // 考试日期
			TeamName    string `json:"team_name"`    // 团队名称
			Exam_clock  string `json:"exam_clock"`   // 考试时间
			QuestionIDs []int  `json:"question_ids"` // 题目ID
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
			return
		}
		user, _ := c.Get("user")
		TeamManagerClaims, ok := user.(*utils.TeamManagerClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误,请联系管理员")
			return
		}
		var teamID int
		for id, name := range TeamManagerClaims.Team {
			if name == request.TeamName {
				teamID = id
				break
			}
		}
		question_num := len(request.QuestionIDs)
		var question_id string
		for i := 0; i < question_num; i++ {
			question_id += strconv.Itoa(request.QuestionIDs[i])
			if i != question_num-1 {
				question_id += "-"
			}
		}
		exam_id, err := controlsql.InsertExamInfo(db, request.ExamName, request.ExamDate, request.Exam_clock, question_num, question_id, teamID)
		if err != nil {
			log.Panic(err)
			c.JSON(500, "服务器错误")
			return
		}
		// 初始化QuestionStatistics
		// err = controlsql.InitQuestionStatistics(db, exam_id, question_num))
		// if err != nil {
		// 	log.Panic(err)
		// 	c.JSON(500, "服务器错误")
		// 	return
		// }

		c.JSON(200, "发布成功")
		//@TODO具体逻辑待议
		//发布考试后在考试的当天设置定时任务，检查是否需要更新数据库，需要则更新并停止任务，不需要则继续等待
		//定时任务的时间为考试时间的当天
		//定时任务的内容为检查是否需要更新数据库
		ticker := time.NewTicker(1 * time.Hour)
		go func() {
			for {
				select {
				case <-ticker.C:
					//检查是否需要更新数据库
					isNeed, err := controlsql.CalculateRank(db, exam_id)
					if err != nil {
						log.Panic(err)
						return
					}
					if isNeed {
						//更新数据库
						err = controlsql.FreshRank(db, exam_id)
						ticker.Stop()
						return
					}
				}
			}
		}()

	})
	//创建团队
	r.POST("/api/team_manage/create_team", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			TeamName string `json:"team_name"` // 团队名称
			MaxNum   int    `json:"max_num"`   // 最大成员数量
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
			return
		}
		user, _ := c.Get("user")
		TeamManagerClaims, ok := user.(*utils.TeamManagerClaims)
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		err := controlsql.RegisterTeam(db, request.TeamName, TeamManagerClaims.ManagerID, request.MaxNum)
		if err != nil {
			log.Panic(err)
			c.JSON(500, "服务器错误")
			return
		}
		type CreateTeamResponse struct {
			Code  int    `json:"code"`  // 状态码
			Msg   string `json:"msg"`   // 消息
			Token string `json:"token"` // 创建团队后产生新的的token
		}
		//输入manager_id，返回manager_Id,teamMap
		managerId, teams, err := controlsql.GetTokenParamsByManagerId(db, TeamManagerClaims.ManagerID)
		if err != nil {
			log.Panic(err)
			c.JSON(500, "服务器错误")
			return
		}
		//根据manager_id和teamMap，生成新的token
		newToken, err := utils.GenerateToken_TeamManager(managerId, teams)
		if err != nil {
			log.Panic(err)
			c.JSON(500, "服务器错误")
			return
		}
		var response CreateTeamResponse
		response.Code = 200
		response.Msg = "创建成功"
		response.Token = newToken
		c.JSON(200, response)
	})
	r.DELETE("/api/team_manage/member_manage/delete", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			TeamName string `json:"team_name"` // 团队名
			UserName string `json:"user_name"` // 用户名
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
			return
		}
		// 从TeamName找到对应TeamID
		TargetTeamID, err := controlsql.SearchTeamIDByTeamName(db, request.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
		}
		// 以TeamID和UserName删除对应团队中的用户

		err = controlsql.DeleteTeammember(db, TargetTeamID, request.UserName)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
		}
		c.JSON(200, "删除成员成功")
	})
	//根据Token中的ManagerID和Team，获取所有team的所有学生所有题型平均分，以及学生其各题型平均分、排名变化
	r.GET("/api/team_manage/exam_situation/teams_and_students_grade", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		TeamManagerClaims, ok := user.(*utils.TeamManagerClaims) // 将 user 转换为 *TeamManagerClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		type Response struct {
			Code                 int                       `json:"code"` // 状态码
			Msg                  string                    `json:"msg"`  // 消息
			TeamAndStudents      *controlsql.CustomMap     `json:"team_and_students"`
			StudentAverageScores []controlsql.AverageScore `json:"student_average_scores"` // 学生各题型平均分
			TeamAverageScores    []controlsql.AverageScore `json:"team_average_scores"`    // 团队各题型平均分
			ExamNames            []string                  `json:"exam_names"`             //前三个月考试名称
			StudentRankScores    []controlsql.RankScore    `json:"student_rank_scores"`    // 学生各题型排名变化
		}
		var response Response
		teamMemberMap, studentIds, err := controlsql.SearchTeamMemberByTeamID(db, TeamManagerClaims.Team)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
		}
		response.TeamAndStudents = teamMemberMap
		//查询学生和团队的各题型平均分
		response.StudentAverageScores, err = controlsql.SearchStudentAverageScoresByStudentIDs(db, rdb, studentIds)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
		}
		fmt.Println(response.StudentAverageScores)
		response.TeamAverageScores, err = controlsql.SearchTeamAverageScoresByTeamMap(rdb, TeamManagerClaims.Team)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
		}
		fmt.Println(response.TeamAverageScores)
		//查最近的五次考试名称 todo 先返回空数据
		response.ExamNames = []string{}
		response.StudentRankScores = []controlsql.RankScore{}
		fmt.Println(response)
		fmt.Println(response.TeamAndStudents)
		response.Code = 200
		response.Msg = "成功"
		c.JSON(200, response)
	})
}

//创建团队 加入团队 删除成员 搜索成员
// utils- 每日更新打卡内容
