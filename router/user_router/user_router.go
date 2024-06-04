package userrouter

import (
	controlsql "DailyEnglish/db"
	middlewares "DailyEnglish/middlewares"
	utils "DailyEnglish/utils"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		//验证用户是否已注册
		if controlsql.UserExists_User(db, data.Email) {
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
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}
		team_id, team_name, err := controlsql.GetTokenParams_User(db, userid)
		fmt.Println(team_id, team_name)
		if err != nil && err.Error() != "sql: no rows in result set" {
			log.Panic(err)
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
	//选择词书界面
	r.GET("/api/users/navigate_books", tokenAuthMiddleware(), func(c *gin.Context) {
		//查询词书信息
		Item, err := controlsql.GetAllBooks(db)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		type Book struct {
			BookID            int    `json:"book_id"`
			BookName          string `json:"book_name"`
			WordNum           int    `json:"word_num"`
			Grade             int    `json:"grade"`
			Grade_description string `json:"grade_description"`
			Describe          string `json:"description"`
		}
		type Response struct {
			Code  int    `json:"code"`
			Msg   string `json:"msg"`
			Books []Book `json:"books"`
		}
		var response Response
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
		for _, item := range Item {
			var book Book
			book.BookID = item.BookID
			book.BookName = item.BookName
			book.WordNum = item.WordNum
			book.Grade = item.Grade
			book.Grade_description = gradeDescriptions[item.Grade]
			book.Describe = item.Describe
			response.Books = append(response.Books, book)
		}
		response.Code = 200
		response.Msg = "成功"
		c.JSON(200, response)
	})
	//选择单词界面
	//选择词书
	r.POST("/api/users/navigate_books", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			BookID int `json:"book_id"`
		}
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		//添加用户词书
		err := controlsql.AddUserBook(db, UserClaims.UserID, request.BookID)
		if err != nil {
			// 检查错误是否为"已完成"
			if err.Error() == "已完成" {
				c.JSON(200, gin.H{
					"code": "200",
					"msg":  "您已设置词书",
				})
			} else {
				// 其他错误
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": "500",
					"msg":  "服务器内部错误",
				})
			}
			return
		}
		//todo 向user_punch-learn表插入或者更新一项数据
		/*[
		{
			"user_id": 1,
			"learned_index": 50,//目前打卡到的单词本的下标
			"punch_num": 20,//打卡总单词数
			"review_num": 20,//复习总单词数
			"date": "2024-05-22"//第一次选择词书的日期
		}
		]*/
		c.JSON(200, gin.H{
			"code": "200",
			"msg":  "设置词书成功",
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
		if err != nil && err != sql.ErrNoRows {
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
			Code      int       `json:"code"`
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
		response.Code = 200
		response.Msg = "成功"
		if err == sql.ErrNoRows {
			response.Code = 404
			response.Msg = "您还没有打卡"
		}
		c.JSON(200, response)
	})
	//打卡
	r.GET("/api/main/take_punch", tokenAuthMiddleware(), func(c *gin.Context) {
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
	// 打卡结果提交
	r.POST("/api/main/punched", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		type Request struct {
			PunchResult map[int]string `json:"punch_result"`
		}
		fmt.Println("接收到的打卡结果为", c.PostForm("punch_result"))
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}

		//TODO 将打卡结果存入数据库
		userId := UserClaims.UserID //获取用户id
		fmt.Println("打卡的用户id为", userId)
		//更新用户学习进度
		err := controlsql.UpdateUserPunch(db, userId, time.Now().Format("2006-01-02"))
		if err != nil && err != sql.ErrNoRows {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "服务器内部错误",
			})
			return
		}

		type Response struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		}
		var response Response
		response.Code = 200
		response.Msg = "成功"
		c.JSON(http.StatusOK, response)

	})

	//复习
	r.GET("/api/main/take_review", tokenAuthMiddleware(), func(c *gin.Context) {
		//查询复习单词，这里写死先
		//复习单词
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
		word.Word = "abondon"
		word.PhoneticUS = "[əˈbændən]"
		word.WordQuestion = make(map[string]string)
		word.WordQuestion["A"] = "放弃"
		word.WordQuestion["B"] = "保留"
		word.WordQuestion["C"] = "拒绝"
		word.WordQuestion["D"] = "接受"
		word.Answer = "A"
		response.WordList = append(response.WordList, word)
		word.WordID = 2
		word.Word = "abroad"
		word.PhoneticUS = "[əˈbrɔːd]"
		word.WordQuestion = make(map[string]string)
		word.WordQuestion["A"] = "国内"
		word.WordQuestion["B"] = "国外"
		word.WordQuestion["C"] = "国际"
		word.WordQuestion["D"] = "国内外"
		word.Answer = "B"
		response.WordList = append(response.WordList, word)
		word.WordID = 3
		word.Word = "absorb"
		word.PhoneticUS = "[əbˈzɔːrb]"
		word.WordQuestion = make(map[string]string)
		word.WordQuestion["A"] = "吸收"
		word.WordQuestion["B"] = "排放"
		word.WordQuestion["C"] = "吸引"
		word.WordQuestion["D"] = "排斥"
		word.Answer = "A"
		response.WordList = append(response.WordList, word)
		word.WordID = 4
		word.Word = "bargain"
		word.PhoneticUS = "[ˈbɑːrɡɪn]"
		word.WordQuestion = make(map[string]string)
		word.WordQuestion["C"] = "讨价还价"
		word.WordQuestion["B"] = "交易"
		word.WordQuestion["A"] = "协商"
		word.WordQuestion["D"] = "交易"
		word.Answer = "C"
		response.WordList = append(response.WordList, word)
		word.WordID = 5
		word.Word = "satisfy"
		word.PhoneticUS = "[ˈsætɪsfaɪ]"
		word.WordQuestion = make(map[string]string)
		word.WordQuestion["C"] = "满足"
		word.WordQuestion["B"] = "满意"
		word.WordQuestion["A"] = "满足"
		word.WordQuestion["D"] = "满足"
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
			log.Panic(err)
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
			//选择今日以前的考试
			if item.ExamDate >= utils.GetCurrentDate() {
				continue
			}
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
	//查询今日考试
	r.POST("api/exams/exam_date", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			ExamDate string `json:"date"`
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 400,
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
		fmt.Println(request.ExamDate, UserClaims.UserID, UserClaims.TeamID)
		Item, err := controlsql.GetExamInfo(db, UserClaims.UserID, UserClaims.TeamID)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
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
			//选择今日的考试
			if item.ExamDate != request.ExamDate {
				continue
			}
			fmt.Println(item.ExamDate, request.ExamDate)
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
	r.POST("/api/exams/examination_details", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			ExamID string `json:"exam_id"`
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			log.Panic(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		fmt.Println(request.ExamID)
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		examId, _ := strconv.Atoi(request.ExamID)
		//查询考试信息
		Item, err := controlsql.GetExamDetail(db, UserClaims.UserID, examId)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		Item2, err := controlsql.GetExamInfoByExamID(db, examId)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		type Question struct {
			QuestionID    int               `json:"question_id"`
			QuestionIndex int               `json:"question_index"`
			QuestionDesc  string            `json:"question_decription"`
			Options       map[string]string `json:"choices"`
			MyAnswer      string            `json:"my_answer"`
			CorrectAnswer string            `json:"correct_answer"`
			Score         int               `json:"score"`
			FullScore     int               `json:"full_score"`
		}
		type Response struct {
			ExamDate    string     `json:"exam_date"`
			QuestionNum int        `json:"question_num"`
			CorrectNum  int        `json:"correct_num"`
			Score       int        `json:"score"`
			Questions   []Question `json:"questions"`
		}
		var response Response
		response.ExamDate = Item2.ExamDate
		response.QuestionNum = len(Item)
		response.CorrectNum = 0
		response.Score = 0
		for i, item := range Item {
			var question Question
			question.QuestionID = item.Question_id
			question.QuestionIndex = i + 1
			question.QuestionDesc = item.Question
			// 扩展item中的Questions内的Choices List["","","",""]为map[["A","B","C","D"],["","","",""]]
			question.Options = make(map[string]string)
			i = 0
			for i, option := range item.Options {
				key := (string)(i + 65) // 65是'A'的ASCII码
				question.Options[key] = option
			}

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
	//考试页面,用户开始考试需要向用户发送考试题目
	r.POST("/api/exams/take_examination", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			ExamID int `json:"exam_id"`
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		type Question struct {
			QuestionID         int               `json:"question_id"`
			QuestionType       int               `json:"question_type"`
			QuestionDifficulty int               `json:"question_difficulty"`
			QuestionGrade      int               `json:"question_grade"`
			QuestionContent    string            `json:"question_content"`
			QuestionChoices    map[string]string `json:"question_choices"`
			QuestionAnswer     string            `json:"question_answer"`
			FullScore          int               `json:"full_score"`
		}
		type Response struct {
			Code         string     `json:"code"`
			Msg          string     `json:"msg"`
			QuestionNum  int        `json:"question_num"`
			QuestionList []Question `json:"question_list"`
		}
		var response Response
		//查询考试信息
		Item1, err := controlsql.GetExamInfoByExamID(db, request.ExamID)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		response.QuestionNum = Item1.QuestionNum
		//查询考试题目
		questionIds := strings.Split(Item1.QuestionID, "-")
		for _, questionId := range questionIds {
			questionId = strings.TrimSpace(questionId)
			questionID, err := strconv.Atoi(questionId)
			if err != nil {
				log.Panic(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": "500",
					"msg":  "服务器内部错误"})
				return
			}
			Item2, err := controlsql.GetQuestionInfo(db, questionID)
			if err != nil {
				log.Panic(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": "500",
					"msg":  "服务器内部错误"})
				return
			}
			var question Question
			question.QuestionID = Item2.Question_id
			question.QuestionType = Item2.Questiontype
			question.QuestionDifficulty = Item2.QuestionDifficulty
			question.QuestionGrade = Item2.QuestionGrade
			question.QuestionContent = Item2.QuestionContent
			question.QuestionChoices = Item2.Options
			question.QuestionAnswer = Item2.QuestionAnswer
			question.FullScore = 5
			response.QuestionList = append(response.QuestionList, question)
		}
		response.Code = "200"
		response.Msg = "成功"
		c.JSON(200, response)
	})
	//提交考试@TODO
	r.POST("/api/exams/submitExamResult", tokenAuthMiddleware(), func(c *gin.Context) {
		type Exam_score struct {
			UserAnswer string `json:"selectedChoice"`
			UserScore  int    `json:"score"`
		}
		type Request struct {
			Exam_result map[int]Exam_score `json:"selectedChoiceAndScore"`
			Exam_id     int                `json:"exam_id"`
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		fmt.Println(request)
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		//计算总分
		var totalScore int
		var user_answer string
		for _, item := range request.Exam_result {
			totalScore += item.UserScore
			user_answer += item.UserAnswer + "-"
		}
		user_answer = strings.TrimRight(user_answer, "-")
		//插入考试信息
		err := controlsql.InsertUserScore(db, UserClaims.UserID, request.Exam_id, user_answer, totalScore)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		c.JSON(200, gin.H{
			"code": "200",
			"msg":  "提交成功",
		})
	})
	// 我的团队
	r.GET("/api/users/my_team", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		fmt.Println("now searching teamid ", UserClaims.TeamID)
		//查询用户所属团队
		Item, err := controlsql.SearchTeamInfo(db, UserClaims.TeamID)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		type User struct {
			UserID   int    `json:"user_id"`
			UserName string `json:"user_name"`
			UserSex  int    `json:"user_sex"`
		}
		type TeamInfo struct {
			TeamID      int    `json:"team_id"`
			TeamName    string `json:"team_name"`
			ManagerID   int    `json:"manager_id"`
			ManagerName string `json:"manager_name"`
			TeamSize    int    `json:"member_num"`
			MemberList  []User `json:"member_list"`
		}
		type Response struct {
			Code     int      `json:"code"`
			Msg      string   `json:"msg"`
			TeamInfo TeamInfo `json:"team"`
		}

		var response Response
		var teamInfo TeamInfo
		teamInfo.TeamID = Item.Teamid
		teamInfo.TeamName = Item.Teamname
		teamInfo.ManagerID = Item.Managerid
		teamInfo.ManagerName = Item.Managername
		teamInfo.TeamSize = Item.Teamsize
		teamInfo.MemberList = make([]User, 0)

		for _, member := range Item.Memberlist {
			var user User
			user.UserID = member.Userid
			user.UserName = member.Username
			user.UserSex = member.Usersex
			teamInfo.MemberList = append(teamInfo.MemberList, user)
		}
		response.Code = 200
		response.Msg = "成功"
		c.JSON(http.StatusOK, response)
	})

	// 加入团队
	r.POST("/api/users/my_team/join_team", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			InvitationCode string `json:"invitation_code"`
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
		//解密邀请码
		fmt.Println("邀请码为:", request.InvitationCode)
		TargetTeamID, _ := utils.DecryptIC(request.InvitationCode, 114514)
		fmt.Println("解密的团队码为:", TargetTeamID)
		//utils.TestICD()

		//查询是否有该ID的团队
		exist, _ := controlsql.CheckTeam(db, TargetTeamID)
		full, _ := controlsql.IsTeamFull(db, TargetTeamID)

		if !exist {
			c.JSON(http.StatusNotFound, gin.H{
				"code": "404",
				"msg":  "邀请码无效",
			})
			return
		} else if full {
			// 该团队是否已满
			c.JSON(http.StatusForbidden, gin.H{
				"code": "403",
				"msg":  "该团队已满",
			})
			return
		}
		// 插入该成员
		insertOK, err := controlsql.JoinTeam(db, UserClaims.UserID, TargetTeamID)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		if !insertOK {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		c.JSON(200, gin.H{
			"code": "200",
			"msg":  "成功加入团队",
		})
	})

	// 获取某一天的所有考试信息
	r.POST("/api/exams/exams_date", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			Date string `json:"date"`
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
		UserClaims, ok := user.(*utils.UserClaims)
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		//查询该日期的考试信息
		Item, err := controlsql.SearchExaminfoByTeamIDAndDate222(db, UserClaims.UserID, request.Date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		type Exam struct {
			ExamID      int    `json:"exam_id"`
			ExamName    string `json:"name"`
			StartTime   string `json:"start_time"`
			Duration    int    `json:"duration"`
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
			exam.ExamID = item.ExamID
			exam.ExamName = item.ExamName
			exam.StartTime = item.StartTime
			exam.Duration = item.Duration
			exam.QuestionNum = item.QuestionNum
			response.Exams = append(response.Exams, exam)
		}
		response.Code = 200
		response.Msg = "成功"
		c.JSON(200, response)
	})
}
