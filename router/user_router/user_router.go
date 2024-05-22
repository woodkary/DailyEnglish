package userrouter

import (
	controlsql "DailyEnglish/db"
	middlewares "DailyEnglish/middlewares"
	utils "DailyEnglish/utils"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func tokenAuthMiddleware() gin.HandlerFunc {
	return middlewares.TokenAuthMiddleware("User")
}
func InitUserRouter(r *gin.Engine, db *sql.DB) {
	//发送验证码
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
		fmt.Print("data.Email: ", data.Email, "\n")
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
	//注册
	r.POST("/api/user/register", func(c *gin.Context) {
		type regdata struct {
			Username string `json:"username"`
			Pwd      string `json:"password"`
			Email    string `json:"email"`
		}

		var data regdata
		fmt.Println("Username:", data.Username, "Pwd:", data.Pwd, "Email:", data.Email)
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		//验证用户是否已注册
		if controlsql.UserExists_User(db, data.Username) {
			c.JSON(http.StatusConflict, gin.H{
				"code": "409",
				"msg":  "用户已注册",
			})
			return
		}
		Key := "123456781234567812345678" //密钥
		cryptoPwd := utils.AesEncrypt(data.Pwd, Key)
		//注册用户
		err := controlsql.RegisterUser_User(db, data.Username, cryptoPwd, data.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "注册成功",
		})
	})
	//登录
	r.POST("/api/user/login", func(c *gin.Context) {
		type logindata struct {
			Username string `json:"username"`
			Pwd      string `json:"password"`
		}
		var data logindata
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		//验证用户是否存在
		if !controlsql.UserExists_User(db, data.Username) {
			c.JSON(403, gin.H{
				"code": "403",
				"msg":  "用户不存在",
			})
			return
		}
		//验证密码是否正确
		isMatch := controlsql.CheckUser_User(db, data.Username, data.Pwd)
		if !isMatch {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "密码错误",
			})
			return
		}
		//生成token
		userid, err := controlsql.GetUserID(db, data.Username)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}
		team_id, team_name, err := controlsql.GetTokenParams_User(db, userid)

		if err != nil && err.Error() != "sql: no rows in result set" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}

		token, _ := utils.GenerateToken_User(userid, team_id, team_name)
		c.JSON(http.StatusOK, gin.H{
			"code":  "200",
			"msg":   "登录成功",
			"token": token,
		})
	})
	//主页面
	r.GET("/api/punch/main_menu", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		//查询用户信息
		Item, err := controlsql.GetUserStudy(db, UserClaims.UserID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		type TaskToday struct {
			BookLearning   string `json:"book_learning"`
			WordNumLearned int    `json:"word_num_learned"`
			WordNumTotal   int    `json:"word_num_total"`
			DaysLeft       int    `json:"days_left"`
			PunchNum       int    `json:"punch_num"`
			ReviewNum      int    `json:"review_num"`
			IsPunched      bool   `json:"ispunched"`
		}
		type Response struct {
			Code      string    `json:"code"`
			Msg       string    `json:"msg"`
			TaskToday TaskToday `json:"task_today"`
		}
		var response Response
		response.TaskToday.BookLearning = Item.BookLearning
		response.TaskToday.WordNumLearned = Item.WordNumLearned
		response.TaskToday.WordNumTotal = Item.WordNumTotal
		response.TaskToday.DaysLeft = Item.Days_left
		response.TaskToday.PunchNum = Item.PunchNum
		response.TaskToday.ReviewNum = 10 //这里写死的@TODO去找那些单词需要复习
		response.TaskToday.IsPunched = Item.IsPunched
		response.Code = "200"
		response.Msg = "成功"
		c.JSON(200, response)
	})
	//打卡
	r.GET("/api/main/take_punch", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		_, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		//查询打卡单词，这里写死先
		//打卡单词
		type Word struct {
			WordID       int               `json:"word_id"`
			Word         string            `json:"word"`
			PhoneticUS   string            `json:"phonetic_us"`
			WordQuestion map[string]string `json:"word_question"`
			Answer       string            `json:"answer"`
		}
		type Response struct {
			Code     int    `json:"code"`
			Msg      string `json:"msg"`
			WordList []Word `json:"word_list"`
		}
		var response Response
		var word Word
		word.WordID = 1
		word.Word = "apple"
		word.PhoneticUS = "[ˈæpl]"
		word.WordQuestion = make(map[string]string)
		word.WordQuestion["A"] = "苹果"
		word.WordQuestion["B"] = "香蕉"
		word.WordQuestion["C"] = "橘子"
		word.WordQuestion["D"] = "梨"
		word.Answer = "A"
		response.WordList = append(response.WordList, word)
		word.WordID = 2
		word.Word = "banana"
		word.PhoneticUS = "[bəˈnænə]"
		word.WordQuestion = make(map[string]string)
		word.WordQuestion["A"] = "苹果"
		word.WordQuestion["B"] = "香蕉"
		word.WordQuestion["C"] = "橘子"
		word.WordQuestion["D"] = "梨"
		word.Answer = "B"
		response.WordList = append(response.WordList, word)
		word.WordID = 3
		word.Word = "orange"
		word.PhoneticUS = "[ˈɔːrɪndʒ]"
		word.WordQuestion = make(map[string]string)
		word.WordQuestion["A"] = "苹果"
		word.WordQuestion["B"] = "香蕉"
		word.WordQuestion["C"] = "橘子"
		word.WordQuestion["D"] = "梨"
		word.Answer = "C"
		response.WordList = append(response.WordList, word)
		word.WordID = 4
		word.Word = "pear"
		word.PhoneticUS = "[per]"
		word.WordQuestion = make(map[string]string)
		word.WordQuestion["A"] = "苹果"
		word.WordQuestion["B"] = "香蕉"
		word.WordQuestion["C"] = "橘子"
		word.WordQuestion["D"] = "梨"
		word.Answer = "D"
		response.WordList = append(response.WordList, word)
		word.WordID = 5
		word.Word = "grape"
		word.PhoneticUS = "[ɡreɪp]"
		word.WordQuestion = make(map[string]string)
		word.WordQuestion["A"] = "苹果"
		word.WordQuestion["B"] = "香蕉"
		word.WordQuestion["C"] = "葡萄"
		word.WordQuestion["D"] = "梨"
		word.Answer = "C"
		response.WordList = append(response.WordList, word)
		response.Code = 200
		response.Msg = "成功"
		c.JSON(200, response)
	})
	//历次考试页面
	r.GET("/api/exams/previous_examinations", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		Item, err := controlsql.GetExamInfo(db, UserClaims.UserID, UserClaims.TeamID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		type Exam struct {
			ExamID      int    `json:"exam_id"`
			ExamName    string `json:"exam_name"`
			ExamDate    string `json:"exam_date"`
			ExamScore   int    `json:"exam_score"`
			ExamRank    int    `json:"exam_rank"`
			QuestionNum int    `json:"question_num"`
		}
		type Response struct {
			Code  int    `json:"code"`
			Msg   string `json:"msg"`
			Exams []Exam `json:"exams"`
		}
		var response Response
		for _, item := range Item {
			var exam Exam
			exam.ExamDate = item.ExamDate
			exam.ExamID = item.ExamID
			exam.ExamName = item.ExamName
			exam.ExamRank = item.ExamRank
			exam.ExamScore = item.ExamScore
			exam.QuestionNum = item.QuestionNum
			response.Exams = append(response.Exams, exam)
		}
		response.Code = 200
		response.Msg = "成功"
		c.JSON(200, response)
	})
	//单次考试详情
	r.GET("/api/exams/examination_details", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			ExamID   int    `json:"exam_id"`
			ExamName string `json:"exam_name"`
			ExamDate string `json:"exam_date"`
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		//查询考试信息
		Item, err := controlsql.GetExamDetail(db, UserClaims.UserID, request.ExamID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		type Question struct {
			QuestionID    int      `json:"question_id"`
			QuestionIndex int      `json:"question_index"`
			QuestionDesc  string   `json:"question_decription"`
			Options       []string `json:"choices"`
			MyAnswer      string   `json:"my_answer"`
			CorrectAnswer string   `json:"correct_answer"`
			Score         int      `json:"score"`
			FullScore     int      `json:"full_score"`
		}
		type Response struct {
			ExamDate    string     `json:"exam_date"`
			QuestionNum int        `json:"question_num"`
			CorrectNum  int        `json:"correct_num"`
			Score       int        `json:"score"`
			Questions   []Question `json:"questions"`
		}
		var response Response
		response.ExamDate = request.ExamDate
		response.QuestionNum = len(Item)
		response.CorrectNum = 0
		response.Score = 0
		for i, item := range Item {
			var question Question
			question.QuestionID = item.Question_id
			question.QuestionIndex = i + 1
			question.QuestionDesc = item.Question
			question.Options = item.Options
			question.MyAnswer = item.UserAnswer
			question.CorrectAnswer = item.Answer
			question.Score = item.Score
			response.Score += item.Score
			if item.UserAnswer == item.Answer {
				response.CorrectNum++
			}
			question.FullScore = 5
			response.Questions = append(response.Questions, question)
		}
		c.JSON(200, response)
	})

}
