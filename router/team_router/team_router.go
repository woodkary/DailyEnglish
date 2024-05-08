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
	"github.com/go-redis/redis"
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

func InitTeamRouter(r *gin.Engine, client *redis.Client, db *sql.DB) {
	//考试情况数据
	r.GET("/api/team_manage/exam_situation/data", tokenAuthMiddleware(), func(c *gin.Context) {
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
			now := time.Now()

			examDate, err := time.Parse("2006-01-02", exam.ExamDate)
			if err != nil {
				log.Println("Error parsing date:", err)
				continue
			}

			if examDate.Year() == now.Year() && examDate.Month() == now.Month() {
				Response.Exam_date = append(Response.Exam_date, exam.ExamDate)
			}
		}
		Response.Code = "200"
		Response.Msg = "成功"
		c.JSON(200, Response)
	})
	//获取某日管理的团队的所有考试信息
	r.POST("/api/team_manage/exam_situation/exam_data", tokenAuthMiddleware(), func(c *gin.Context) {
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
		var Item []controlsql.ExamInfo
		for _, teamID := range userClaims.TeamID {
			examInfo, err := controlsql.SearchExamInfoByTeamIDAndDate(db, teamID, request.Date)
			if err != nil {
				c.JSON(500, "服务器错误")
				log.Panic(err)
				return

			}
			Item = append(Item, examInfo...)
		}
		// ExamsResponse 结构体表示包含多个考试的响应
		type response struct {
			Code  string `json:"code"` // 响应代码
			Msg   string `json:"msg"`  // 响应消息
			Exams []struct {
				ExamID   string `json:"exam_id"`   // 考试ID
				ExamName string `json:"exam_name"` // 考试名称
				ExamDate string `json:"exam_date"` // 考试日期
			} `json:"exams"` // 考试列表
		}
		var Response response
		Response.Code = "200"
		Response.Msg = "成功"
		for _, exam := range Item {
			if exam.ExamDate == request.Date {
				var examData struct {
					ExamID   string `json:"exam_id"`
					ExamName string `json:"exam_name"`
					ExamDate string `json:"exam_date"`
				}

				examData.ExamID = strconv.Itoa(exam.ExamID)
				examData.ExamName = exam.ExamName
				examData.ExamDate = exam.ExamDate
				Response.Exams = append(Response.Exams, examData)
			}
		}
		c.JSON(200, Response)
	})

	//获取单次考试详情
	r.POST("/api/team_manage/exam_situation/exam_detail", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			ExamID string `json:"exam_id"` // 考试名称
			TeamID string `json:"team_id"` // 团队名称
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
			return
		}

		examInfo, err := controlsql.GetExamInfoByID(client, "Exam1")
		if err != nil {
			c.JSON(500, "服务器错误")
			log.Panic(err)
			return
		}

		ScoresInExam, err := controlsql.GetScoresInExambyExamID(client, "Exam1")
		if err != nil {
			c.JSON(500, "服务器错误")
			log.Panic(err)
			return
		}
		levelNums := service.CalculateUserLevel(ScoresInExam)

		type UserResult struct {
			Attend   string `json:"attend"`   // 考试参与情况
			Username string `json:"username"` // 用户名
			Score    string `json:"score"`    // 得分
			FailNum  string `json:"fail_num"` // 错题数量
			Progress string `json:"progress"` // 进步分数 (相距上次)
		}

		QuestionNum := controlsql.GetQuestionNum(client, "Exam1") // 考试题目数量
		var qd = make([][5]int, QuestionNum)                      // 考试题目详情
		for i := 0; i < QuestionNum; i++ {
			for j := 0; j < 5; j++ {
				//TODO
				// 查询考试每一题正确选项，选A人数，选B人数，C,D
			}
		}

		type ExamDetail struct {
			ID             string       `json:"exam_id"`          // 考试ID
			Name           string       `json:"exam_name"`        // 考试名称
			UserLevels     []int        `json:"user_levels"`      // 用户等级
			QuestionDetail [][5]int     `json:"question_details"` // 考试题目详情
			UserResult     []UserResult `json:"user_result"`      // 考试参与人员得分情况
		}
		type response struct {
			Code       string     `json:"code"`        // 状态码
			Msg        string     `json:"msg"`         // 消息
			ExamDetail ExamDetail `json:"exam_detail"` // 考试详情
		}
		var Response response
		Response.Code = "200"
		Response.Msg = "成功"
		Response.ExamDetail.ID = strconv.Itoa(examInfo.ID)
		Response.ExamDetail.Name = examInfo.Name
		Response.ExamDetail.UserLevels = levelNums
		Response.ExamDetail.QuestionDetail = qd
		for _, score := range ScoresInExam {
			var userResult UserResult
			userResult.Username = score.Username
			userResult.Score = score.Score
			userResult.FailNum = "0"  //
			userResult.Progress = "0" //
			Response.ExamDetail.UserResult = append(Response.ExamDetail.UserResult, userResult)
		}
		c.JSON(200, Response)
	})

	//打卡详情界面
	r.GET("/api/team_manage/punch_statistics/data", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		userClaims, ok := user.(*service.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}

		Item1, err := controlsql.GetTeamMembersAttendance1(client, userClaims.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
		}
		Item2, err := controlsql.GetTeamMembersAttendanceByDate(client, userClaims.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
		}
		Item3, err := controlsql.GetTeamMembersAttendanceNum(client, userClaims.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
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
	r.GET("/api/team_manage/member_manage/data", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		userClaims, ok := user.(*service.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}

		Item, err := controlsql.GetTeamInfo(client, userClaims.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
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
	// //成员删除
	// r.POST("/api/team_manage/member_manage/delete", tokenAuthMiddleware(), func(c *gin.Context) {
	// 	type Request struct {
	// 		Username string `json:"username"` // 要删除的成员的用户名
	// 		Teamname string `json:"teamname"` // 团队名
	// 	}
	// 	var request Request
	// 	if err := c.ShouldBind(&request); err != nil {
	// 		c.JSON(400, "请求参数错误")
	// 	}
	// 	err := controlsql.DeleteUserFromTeam(client, request.Teamname, request.Username)
	// 	if err != nil {
	// 		c.JSON(500, "服务器错误")
	// 	}
	// 	c.JSON(200, "删除成功")
	// })
	//搜索成员
	r.POST("/api/team_manage/member_manage/search", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			Username string `json:"username"` // 用户名
			Teamname string `json:"teamname"` // 团队名
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
			return
		}
		Item, err := controlsql.GetTeamMembers(client, request.Teamname)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
		}
		type Member struct {
			Name     string `json:"name"`      // 成员姓名
			Right    string `json:"right"`     // 成员权限
			Time     string `json:"time"`      // 加入组织时间
			PunchDay string `json:"punch_day"` // 打卡天数
		}
		type response struct {
			Code    string   `json:"code"`    // 状态码
			Msg     string   `json:"msg"`     // 消息
			Members []Member `json:"members"` // 团队成员列表
		}
		var Response response
		for _, m := range Item {
			if m.Username == request.Username { // 检查成员名称是否与请求中的用户名匹配
				var member Member
				member.Name = m.Username
				member.PunchDay = strconv.Itoa(m.AttendanceDays)
				member.Time = m.JoinDate
				if m.IsAdmin {
					member.Right = "Admin"
				} else {
					member.Right = "Member"
				}
				Response.Members = append(Response.Members, member)
				break // 如果找到了匹配的成员，就没有必要继续循环
			}
		}

		if len(Response.Members) == 0 {
			// 如果没有找到匹配的成员，返回空列表
			Response.Code = "404"
			Response.Msg = "未找到成员"
		} else {
			Response.Code = "200"
			Response.Msg = "成功"
		}
		c.JSON(200, Response)
	})

	//获取审核申请页面信息
	r.GET("/api/team_manage/request_manage/data", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		userClaims, ok := user.(*service.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}

		Item1, err := controlsql.GetTeamRequestsByFlag(client, userClaims.TeamName, "0")

		if err != nil {
			c.JSON(500, "服务器错误")
		}
		Item2, err := controlsql.GetTeamInfo(client, userClaims.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		type Request struct {
			Name     string `json:"name"`      // 请求者姓名
			Time     string `json:"time"`      // 请求时间
			LeaveMsg string `json:"leave_msg"` // 留言信息
		}

		// Response 定义了响应的信息
		type response struct {
			Code       string    `json:"code"`        // 状态码
			Msg        string    `json:"msg"`         // 消息
			MemberNum  int       `json:"member_num"`  // 成员数量
			ManagerNum int       `json:"manager_num"` // 管理员数量
			Requests   []Request `json:"request"`     // 请求列表
		}
		var Response response
		Response.Code = "200"
		Response.Msg = "成功"
		Response.ManagerNum = Item2.AdminCount
		Response.MemberNum = Item2.TotalMembers
		for _, r := range Item1 {
			var request Request
			request.Name = r.Username
			request.Time = r.Time
			request.LeaveMsg = r.Message
			Response.Requests = append(Response.Requests, request)
		}
		c.JSON(200, Response)
	})
	//获取个人中心界面所需信息
	r.GET("/api/team_manage/personal_center/data", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		userClaims, ok := user.(*service.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		Item1, Item2, Item3, err := controlsql.GetUserInfoByEmailPwd(db, userClaims.UserName)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		type User struct {
			Name string `json:"name"` // 用户姓名
			// 用户所属全部团队
			Teams    []string `json:"teams"`
			Right    string   `json:"right"` // 用户权限
			Email    string   `json:"email"` // 用户邮箱
			Password string   `json:"pwd"`   //用户密码
			Phone    string   `json:"phone"` // 用户手机号
		}
		// Response 定义了响应的信息
		type Response struct {
			Code string `json:"code"` // 状态码
			Msg  string `json:"msg"`  // 消息
			User User   `json:"user"` // 用户信息
		}
		var response Response
		response.User.Email = Item1
		response.User.Password = Item2
		response.User.Phone = Item3
		response.User.Name = userClaims.UserName
		response.User.Teams, err = controlsql.GetJoinedTeams(client, userClaims.UserName)
		if err != nil {
			c.JSON(500, "服务器错误")
		}

		response.Code = "200"
		response.Msg = "成功"
		c.JSON(200, response)
	})
	//打卡信息发布界面信息
	r.GET("/api/team_manage/task_daily/data", tokenAuthMiddleware(), func(c *gin.Context) {
		c.Get("user")
		Item, err := controlsql.QueryBooksBy(db, "大一", "简单", 0)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		type Book struct { //书籍信息
			Title      string `json:"title"`       // 书名
			LearnerNum int    `json:"learner_num"` // 学习人数
			Describe   string `json:"describe"`    // 描述
			WordsNum   int    `json:"words_num"`   // 单词数
		}
		type response struct { //响应信息
			Code  string `json:"code"`  // 状态码
			Msg   string `json:"msg"`   // 消息
			Books []Book `json:"books"` // 书籍列表
		}
		var Response response
		for _, item := range Item {
			var book Book
			book.Title = item.Title
			book.LearnerNum = item.LearnerNum
			book.Describe = item.Describe
			book.WordsNum = item.WordsNum
			Response.Books = append(Response.Books, book)
		}
		Response.Code = "200"
		Response.Msg = "成功"
		c.JSON(200, Response)
	})
	//团队管理员选择打卡发布页面信息
	r.POST("/api/team_manage/task_daily/deliver_punch", func(c *gin.Context) {
		type Request struct {
			Grade          string `json:"grade"`          //年级
			Difficulty     string `json:"difficulty"`     //难度
			Flag           string `json:"flag"`           //按什么排序，0是最新，1是最热
			Time_expected  string `json:"time_expected"`  //打卡时长预计
			Words_expected string `json:"words_expected"` //每日单词预计
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		flag, _ := strconv.Atoi(request.Flag)
		Item, err := controlsql.QueryBooksBy(db, request.Grade, request.Difficulty, flag)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		type Book struct { //书籍信息
			Name       string `json:"name"`        // 书名
			LearnerNum int    `json:"learner_num"` // 学习人数
			Describe   string `json:"describe"`    // 描述
			WordsNum   int    `json:"words_num"`   // 单词数
		}
		type Response struct { //响应信息
			Code  string `json:"code"`  // 状态码
			Msg   string `json:"msg"`   // 消息
			Books []Book `json:"books"` // 书籍列表
		}
		var response Response
		for _, item := range Item {
			var book Book
			book.Name = item.Title
			book.Describe = item.Describe
			book.LearnerNum = item.LearnerNum
			book.WordsNum = item.WordsNum
			response.Books = append(response.Books, book)
		}
		response.Code = "200"
		response.Msg = "成功"
		c.JSON(http.StatusOK, response)
	})
}
