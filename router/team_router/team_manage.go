package teamrouter

import (
	controlsql "DailyEnglish/Control_SQL"
	service "DailyEnglish/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

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

func Team_manager(r *gin.Engine, client *redis.Client, db *sql.DB) {
	//主页数据
	r.GET("/api/team_manage/index/data", tokenAuthMiddleware(), func(c *gin.Context) {
		// 定义JSON响应的结构体
		type Response struct {
			Code  string `json:"code"`
			Msg   string `json:"msg"`
			Punch struct {
				Total_punchrate string   `json:"total_punchrate"`
				Punched         int      `json:"punched"`
				PunchNum        []string `json:"punch_num"`
				PunchRate       []string `json:"punch_rate"`
				PunchLB         []struct {
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
		response.Punch.Total_punchrate = "87%"
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
	//主页打卡详情
	r.GET("/api/team_manage/index/sign_in_details", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		userClaims, ok := user.(*service.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		Item, err := controlsql.GetTeamMembersAttendanceByDate(client, userClaims.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		// SignInDetails 结构体对应于JSON数组中的一个对象
		type SignInDetails struct {
			Name     string `json:"name"`
			IsSignIn string `json:"issignIn"`
		}
		// Response 结构体对应于JSON响应
		type Response struct {
			Code    string          `json:"code"`
			Msg     string          `json:"msg"`
			Details []SignInDetails `json:"sign_in_details"`
		}
		var response Response
		for key, value := range Item {
			var sign_in_details SignInDetails
			sign_in_details.Name = key
			if value == 1 {
				sign_in_details.IsSignIn = "是"
			} else {
				sign_in_details.IsSignIn = "否"
			}
			response.Details = append(response.Details, sign_in_details)
		}
		c.JSON(200, response)
	})
	//考试情况数据
	r.GET("/api/team_manage/exam_situation/data", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		userClaims, ok := user.(*service.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		Item, err := controlsql.QueryTeamExams(client, userClaims.TeamName)
		if err != nil {
			c.JSON(500, "服务器错误")
			return
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
	//获取单次考试详情
	r.POST("/api/team_manage/exam_situation/exam_detail", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			ExamName string `json:"exam_name"` // 考试名称
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(400, "请求参数错误")
			return
		}

		examInfo, err := controlsql.GetExamInfoByName(client, "Exam1")
		fmt.Println(examInfo)
		fmt.Println("AverageScore", examInfo.AverageScore)
		fmt.Println("PassRate", examInfo.PassRate)
		fmt.Println("QuestionCount", examInfo.QuestionCount)
		fmt.Println("ID", examInfo.ID)
		fmt.Println("Name", examInfo.Name)
		fmt.Println("TopSix", examInfo.TopSix)
		fmt.Println(examInfo.Questions)
		//fmt.Println("TotalScore",examInfo.TotalScore)
		//fmt.Println("TotalUser",examInfo.TotalUser)

		if err != nil {
			c.JSON(500, "服务器错误")
			log.Panic(err)
			return
		}
		type UserResult struct {
			Attend   string `json:"attend"`   // 考试参与情况
			Username string `json:"username"` // 用户名
			Score    string `json:"score"`    // 得分
			FailNum  string `json:"fail_num"` // 错题数量

		}
		type ExamDetail struct {
			ID          string       `json:"exam_id"`     // 考试ID
			Name        string       `json:"exam_name"`   // 考试名称
			UserResults []UserResult `json:"user_result"` // 考试参与人员及得分
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
		//TODO 这里应该查询数据库获取考试所有参与人员等情况
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
		Item1, Item2, err := controlsql.GetUserInfoByEmailPwd(db, userClaims.UserName)
		if err != nil {
			c.JSON(500, "服务器错误")
		}
		type User struct {
			Name     string `json:"name"`  // 用户姓名
			Team     string `json:"team"`  // 用户所属团队
			Right    string `json:"right"` // 用户权限
			Email    string `json:"email"` // 用户邮箱
			Password string `json:"pwd"`   //用户密码
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
		response.User.Name = userClaims.UserName
		response.User.Team = userClaims.TeamName
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
