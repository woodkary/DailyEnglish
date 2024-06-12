package userrouter

import (
	controlsql "DailyEnglish/db"
	middlewares "DailyEnglish/middlewares"
	utils "DailyEnglish/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"golang.org/x/net/context"
)

func tokenAuthMiddleware() gin.HandlerFunc {
	return middlewares.TokenAuthMiddleware("User")
}

// 升级http连接
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	connections = make(map[string]*websocket.Conn) //
	mutex       = &sync.Mutex{}                    //互斥锁
)

// FormatWordData formats the word data into the desired string format
func FormatWordData(wordData map[string]interface{}) string {
	var formattedData string
	formattedData = "'{"
	for key, value := range wordData {
		formattedData += fmt.Sprintf("%s:", key)
		switch v := value.(type) {
		case string:
			formattedData += fmt.Sprintf("'%s',", v)
		default:
			formattedData += jsonValue(v) + ","
		}
	}
	// Remove the trailing comma
	if len(formattedData) > 1 {
		formattedData = formattedData[:len(formattedData)-1]
	}
	formattedData += "}'"
	return formattedData
}

// jsonValue converts value to JSON format
func jsonValue(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}
func InitUserRouter(r *gin.Engine, db *sql.DB, rdb *redis.Client, es *elasticsearch.Client) {
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
		//将验证码存入 Redis
		ctx := context.Background()                         // 创建一个空的 context
		key := fmt.Sprintf("%s:%s", "app", data.Email)      // key前缀为app:邮箱
		err = rdb.Set(ctx, key, Vcode, time.Minute*5).Err() // 验证码有效期5分钟,更新时替换
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "验证码存储失败",
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
			Code     string `json:"code"`
		}

		var data regdata
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		//验证验证码
		ctx := context.Background()
		key := fmt.Sprintf("%s:%s", "app", data.Email)
		code, rerr := rdb.Get(ctx, key).Result()
		if rerr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "验证码已过期",
			})
			return
		}
		if code != data.Code {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "401",
				"msg":  "验证码错误",
			})
			return
		}
		// 验证码验证成功后尝试删除验证码，即使删除失败也不会影响流程
		rerr = rdb.Del(ctx, key).Err()
		if rerr != nil {
			fmt.Printf("删除验证码失败：%v\n", rerr)
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
		isChoosed := controlsql.CheckUserBook(db, userid)
		token, _ := utils.GenerateToken_User(userid, team_id, team_name)
		c.JSON(http.StatusOK, gin.H{
			"code":      "200",
			"msg":       "登录成功",
			"token":     token,
			"isChoosed": isChoosed,
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

		err = controlsql.AddUserPunchLearn(db, UserClaims.UserID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}
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
	//修改词书
	r.PUT("/api/users/navigate_books", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			BookID int `json:"book_id"`
		}
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims)
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

		//修改用户词书
		err := controlsql.UpdateUserBook(db, UserClaims.UserID, request.BookID)
		if err != nil {
			// 其他错误
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}
		//todo 向user_punch-learn表更新一项数据
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
			"msg":  "修改词书成功",
		})
	})

	//主页面
	r.POST("/api/punch/main_menu", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		type Request struct {
			Time int `json:"time"`
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		//查询用户信息
		Item, _ := controlsql.GetUserStudy(db, UserClaims.UserID)

		Item2, err := controlsql.GetReviewWordID(db, UserClaims.UserID)
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
		response.TaskToday.PunchNum = Item.PunchNum
		response.TaskToday.IsPunched = Item.IsPunched
		response.TaskToday.ReviewNum = len(Item2)
		response.TaskToday.BookLearning = Item.BookLearning
		response.TaskToday.WordNumLearned = Item.WordNumLearned
		response.TaskToday.WordNumTotal = Item.WordNumTotal
		response.TaskToday.DaysLeft = Item.Days_left

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

		// 从UserClaims中获取用户id
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}

		wordlist, err := controlsql.GetUserPunchContent(db, UserClaims.UserID)
		if err != nil && err != sql.ErrNoRows {
			log.Panic(err)
			c.JSON(500, "服务器内部错误")
			return
		}

		var aword Word
		for _, word := range wordlist {
			aword.WordID = word.WordID
			aword.Word = word.Word
			aword.PhoneticUS = word.PhoneticUS
			aword.WordQuestion = word.WordQuestion
			aword.Answer = word.Answer
			response.WordList = append(response.WordList, aword)
		}
		fmt.Println("打卡单词列表 ", response.WordList)
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
			PunchResult map[int]bool `json:"punch_result"`
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

		// 将打卡结果存入数据库
		userId := UserClaims.UserID //获取用户id
		fmt.Println("打卡的用户id为", userId)
		//更新用户学习进度
		err := controlsql.UpdateUserPunch(db, userId, time.Now().Format("2006-01-02"), rdb, request.PunchResult)
		if err != nil && err != sql.ErrNoRows {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "服务器内部错误",
			})
			return
		}
		//更新learned_index
		err = controlsql.UpdateUserLearnIndex(db, userId)
		if err != nil {
			if err.Error() == "请先选择词书" {
				c.JSON(http.StatusNotFound, gin.H{
					"code": 404,
					"msg":  "请先选择词书",
				})
				return
			}
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 500,
				"msg":  "服务器内部错误",
			})
			return
		}
		//打卡后插入学习记录并更新复习时间
		for k, v := range request.PunchResult {
			err := controlsql.InsertUserMemory(db, userId, k, v)
			if err != nil && err != sql.ErrNoRows {
				log.Panic(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg":  "服务器内部错误",
				})
				return
			}
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
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
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
		//数据库中查找word_id
		wordIDs, err := controlsql.GetReviewWordID(db, UserClaims.UserID)
		if err != nil && err != sql.ErrNoRows {
			log.Panic(err)
			c.JSON(500, "服务器内部错误")
			return
		}
		//查询单词信息
		words, err := controlsql.GetWordByWordID(db, wordIDs)
		if err != nil && err != sql.ErrNoRows {
			log.Panic(err)
			c.JSON(500, "服务器内部错误")
			return
		}
		var aword Word
		for _, word := range words {
			aword.WordID = word.WordID
			aword.Word = word.Word
			aword.PhoneticUS = word.PhoneticUS
			aword.WordQuestion = word.WordQuestion
			aword.Answer = word.Answer
			response.WordList = append(response.WordList, aword)
		}
		fmt.Println("复习单词列表 ", response.WordList)
		response.Code = 200
		response.Msg = "成功"
		c.JSON(200, response)
	})
	// 复习结果提交
	r.POST("/api/main/reviewed", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(500, "服务器错误")
			return
		}
		type Request struct {
			PunchResult map[int]bool `json:"punch_result"`
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
		for k, v := range request.PunchResult {
			err := controlsql.UpdatetUserMemory(db, userId, k, v)
			if err != nil && err != sql.ErrNoRows {
				log.Panic(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 500,
					"msg":  "服务器内部错误",
				})
				return
			}
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
		fmt.Println("查询日期为:", request.Date)
		Item, err := controlsql.SearchExaminfoByTeamIDAndDate222(db, UserClaims.TeamID, UserClaims.UserID, request.Date)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		fmt.Println("查询结果为:", Item)
		type Exam struct {
			ExamID      string `json:"exam_id"`
			ExamName    string `json:"exam_name"`
			StartTime   string `json:"start_time"`
			ExamDate    string `json:"exam_date"`
			Duration    int    `json:"duration"`
			QuestionNum int    `json:"question_num"`
			ExamScore   int    `json:"exam_score"`
		}
		type Response struct {
			Code  int    `json:"code"`
			Msg   string `json:"msg"`
			Exams []Exam `json:"exams"`
		}
		var response Response
		for _, item := range Item {
			var exam Exam
			exam.ExamID = strconv.Itoa(item.ExamID)
			exam.ExamName = item.ExamName
			exam.StartTime = item.StartTime
			exam.Duration = item.Duration
			exam.QuestionNum = item.QuestionNum
			exam.ExamDate = item.ExamDate
			exam.ExamScore = item.ExamFullScore
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
			ExamID string `json:"exam_id"`
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
		examId, _ := strconv.Atoi(request.ExamID)
		//查询考试信息
		Item1, err := controlsql.GetExamInfoByExamID(db, examId)
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
	//提交考试
	//redis------studentId:question_type:["score","num"]
	r.POST("/api/exams/submitExamResult", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			Exam_result map[int]controlsql.Exam_score `json:"selectedChoiceAndScore"`
			Exam_id     string                        `json:"exam_id"`
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
		exam_id, _ := strconv.Atoi(request.Exam_id)
		//插入考试信息
		err := controlsql.InsertUserScore(db, UserClaims.UserID, exam_id, user_answer, totalScore)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}

		// 更新该场考试的 question_statistics 表

		//向redis插入学生各题型总分信息
		averageScores, err := controlsql.UpdateStudentRDB(db, rdb, UserClaims.UserID, UserClaims.TeamID, request.Exam_result)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
		}
		fmt.Println("各题型平均分为：", averageScores)
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

	// 获取生词本的单词
	//从redis获取释义，从mysql获取拼写和读音
	r.GET("/api/words/get_starbk", tokenAuthMiddleware(), func(c *gin.Context) {
		// 从请求上下文中获取用户信息
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器错误",
			})
			return
		}
		ctx := context.Background()
		userID := UserClaims.UserID
		// 构造Redis键的匹配模式
		pattern := fmt.Sprintf("word:%d:*", userID)

		// 使用SCAN命令获取所有匹配的键，并构成keys切片
		var cursor uint64
		var keys []string
		for {
			var err error
			var result []string
			result, cursor, err = rdb.Scan(ctx, cursor, pattern, 0).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": "500",
					"msg":  "服务器内部错误",
				})
				return
			}
			keys = append(keys, result...)
			if cursor == 0 {
				break
			}
		}

		// 定义单词数据结构
		type WordData struct {
			WordID        int                  `json:"word_id"`
			Spelling      string               `json:"spelling"`
			Pronunciation string               `json:"pronunciation"`
			Meanings      *controlsql.Meanings `json:"meanings"`
		}

		// 并行获取所有单词的详细信息
		var words []WordData
		var wg sync.WaitGroup
		wordChan := make(chan WordData, len(keys)) // 用于存储单词数据的通道
		errChan := make(chan error, len(keys))     // 用于存储错误的通道

		// 遍历所有键，启动goroutine并行处理
		//每有一个收藏的单词，就会增加一个goroutine
		for _, key := range keys {
			wg.Add(1)
			go func(key string) {
				defer wg.Done()
				wordData := WordData{
					Meanings: new(controlsql.Meanings),
				}

				// 从Redis哈希中获取字段值
				values, err := rdb.HGetAll(ctx, key).Result()
				if err != nil {
					errChan <- err
					return
				}

				// 解析各词性字段值到Meanings结构体中
				wordData.Meanings.Adjective = strings.Split(values["adjective"], ",")
				wordData.Meanings.Adverb = strings.Split(values["adverb"], ",")
				wordData.Meanings.Conjunction = strings.Split(values["conjunction"], ",")
				wordData.Meanings.Interjection = strings.Split(values["interjection"], ",")
				wordData.Meanings.Noun = strings.Split(values["noun"], ",")
				wordData.Meanings.Preposition = strings.Split(values["preposition"], ",")
				wordData.Meanings.Pronoun = strings.Split(values["pronoun"], ",")
				wordData.Meanings.Verb = strings.Split(values["verb"], ",")

				// 从键中提取wordID
				var wordID int
				_, err = fmt.Sscanf(key, "word:%d:%d", &userID, &wordID)
				if err != nil {
					errChan <- err
					return
				}
				wordData.WordID = wordID

				// 从MySQL根据wordID获取单词发音和单词拼写
				pronAndSpelling, err := controlsql.GetWordInfo(db, wordID)
				if err != nil {
					errChan <- err
					return
				}

				wordData.Pronunciation = pronAndSpelling[0]
				wordData.Spelling = pronAndSpelling[1]

				// 将处理好的单词数据发送到通道
				wordChan <- wordData
			}(key)
		}

		// 等待所有goroutine完成，并关闭通道
		go func() {
			wg.Wait()
			close(wordChan)
			close(errChan)
		}()

		// 收集结果和错误
		for wd := range wordChan {
			words = append(words, wd)
		}
		if err := <-errChan; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}

		// 返回结果
		c.JSON(http.StatusOK, gin.H{
			"code":  200,
			"msg":   "获取用户存储的单词成功",
			"words": words,
		})
	})

	// 获取单词详情
	r.GET("/api/words/get_word_detail", tokenAuthMiddleware(), func(c *gin.Context) {
		type Request struct {
			WordID int `json:"word_id"`
		}
		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}

		// 定义 mean_sen_meansen 结构体
		type mean_sen_meansen struct {
			dtm map[string]string
		}

		// 定义 phra_mea 结构体
		type phra_mea struct {
			pam map[string]string
		}

		// 构造返回数据结构
		type Response struct {
			Code              int                `json:"code"`
			Msg               string             `json:"msg"`
			Word_Spelling     string             `json:"word_spelling"`
			Word_Phonetic     string             `json:"word_phonetic"`
			Word_Distortion   map[string]string  `json:"word_distortion"`
			Detailed_Meanings []mean_sen_meansen `json:"detailed_meanings"`
			Phrases           []phra_mea         `json:"phrases"`
		}

		// 从MySQL根据wordID获取单词信息
		wordDetail, err := controlsql.GetWordDetailByWordId(db, request.WordID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}

		var response Response
		response.Code = 200
		response.Msg = "成功"
		response.Word_Spelling = wordDetail.Word
		response.Word_Phonetic = wordDetail.Pronounciation
		response.Word_Distortion = map[string]string{
			"过去式":    wordDetail.Morpheme1,
			"第三人称单数": wordDetail.Morpheme2,
		}

		response.Detailed_Meanings = make([]mean_sen_meansen, 0)
		var item1, item2 mean_sen_meansen
		item1.dtm = make(map[string]string)
		item1.dtm["chinese_meaning"] = wordDetail.Word_Meaning1
		item1.dtm["example_sentence"] = wordDetail.Sentence1
		item1.dtm["sentence_meaning"] = wordDetail.Sentence_Meaning1
		response.Detailed_Meanings = append(response.Detailed_Meanings, item1)
		item2.dtm = make(map[string]string)
		item2.dtm["chinese_meaning"] = wordDetail.Word_Meaning2
		item2.dtm["example_sentence"] = wordDetail.Sentence2
		item2.dtm["sentence_meaning"] = wordDetail.Sentence_Meaning2
		response.Detailed_Meanings = append(response.Detailed_Meanings, item2)

		response.Phrases = make([]phra_mea, 0)
		var item3, item4 phra_mea
		item3.pam = make(map[string]string)
		item3.pam["phrase"] = wordDetail.Phrase1
		item3.pam["meaning"] = wordDetail.Phrase_Meaning1
		response.Phrases = append(response.Phrases, item3)
		item4.pam = make(map[string]string)
		item4.pam["phrase"] = wordDetail.Phrase2
		item4.pam["meaning"] = wordDetail.Phrase_Meaning2
		response.Phrases = append(response.Phrases, item4)

		c.JSON(http.StatusOK, response)

	})

	// 添加生词
	r.POST("/api/words/add_new_word", tokenAuthMiddleware(), func(c *gin.Context) {
		type request struct {
			Username string `json:"username"`
			WordId   int    `json:"word_id"`
		}
		var req request
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器错误",
			})
			return
		}
		ctx := context.Background()
		// 获取单词信息
		wordData, err := controlsql.GetWordByWordId(db, req.WordId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}

		// 添加生词到 Redis
		//key为word:userID:wordID
		key := fmt.Sprintf("word:%d:%d", UserClaims.UserID, req.WordId)

		// 将Meanings转换为map以便于存储到Redis中，并按字段名称字典序排序
		meaningsMap := map[string]interface{}{
			"adjective":    strings.Join(wordData.Meanings.Adjective, ","),
			"adverb":       strings.Join(wordData.Meanings.Adverb, ","),
			"conjunction":  strings.Join(wordData.Meanings.Conjunction, ","),
			"interjection": strings.Join(wordData.Meanings.Interjection, ","),
			"noun":         strings.Join(wordData.Meanings.Noun, ","),
			"preposition":  strings.Join(wordData.Meanings.Preposition, ","),
			"pronoun":      strings.Join(wordData.Meanings.Pronoun, ","),
			"verb":         strings.Join(wordData.Meanings.Verb, ","),
		}

		// 创建一个有序切片存储字段和值，确保所有字段都被初始化
		orderedFields := []string{"adjective", "adverb", "conjunction", "interjection", "noun", "preposition", "pronoun", "verb"}
		orderedMap := make(map[string]interface{})
		for _, field := range orderedFields {
			if value, exists := meaningsMap[field]; exists {
				orderedMap[field] = value
			} else {
				orderedMap[field] = ""
			}
		}

		// 使用HMSet按顺序插入字段和值
		err = rdb.HMSet(ctx, key, orderedMap).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"code": "200",
			"msg":  "生词添加成功",
		})
	})
	//个人中心页面
	r.GET("/api/users/my_punches", tokenAuthMiddleware(), func(c *gin.Context) {
		user, _ := c.Get("user")
		UserClaims, ok := user.(*utils.UserClaims) // 将 user 转换为 *UserClaims 类型
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器错误",
			})
			return
		}
		//查询用户的打卡记录
		userPunchInfo, err := controlsql.GetUserCenter(db, UserClaims.UserID)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		type Response struct {
			Code                int    `json:"code"`
			Msg                 string `json:"msg"`
			PunchWordNum        int    `json:"punch_word_num"`        //打卡单词数
			TotalPunchDay       int    `json:"total_punch_day"`       //总打卡天数
			ConsecutivePunchDay int    `json:"consecutive_punch_day"` //连续打卡天数
		}
		var response Response
		response.Code = 200
		response.Msg = "成功"
		response.PunchWordNum = userPunchInfo.PunchWordNum
		response.TotalPunchDay = userPunchInfo.TotalPunchDay
		response.ConsecutivePunchDay = userPunchInfo.ConsecutivePunchDay
		c.JSON(http.StatusOK, response)
	})
	//测试接口
	r.POST("/test/users/check_punch_finish", func(c *gin.Context) {
		type Request struct {
			UserID int `json:"user_id"`
			BookID int `json:"book_id"`
		}

		var request Request
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		//查询用户的打卡记录
		userPunchInfo, err := controlsql.CheckUserPunchFinish(db, request.UserID, request.BookID)
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		type Response struct {
			Code          int    `json:"code"`
			Msg           string `json:"msg"`
			UserPunchInfo int    `json:"user_punch_info"`
		}
		var response Response
		response.Code = 200
		response.Msg = "成功"
		response.UserPunchInfo = userPunchInfo
		c.JSON(http.StatusOK, response)
	})
	//查找单词
	r.POST("/api/users/search_words", func(ctx *gin.Context) {
		type Request struct {
			Input string `json:"input"`
		}

		var req Request
		if err := ctx.ShouldBind(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": "400",
				"msg":  "请求参数错误",
			})
			return
		}
		//根据input，搜索词库中的所有匹配的单词
		words, err := controlsql.SearchWords(db, es, req.Input)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": "500",
				"msg":  "服务器内部错误"})
			return
		}
		type Response struct {
			Code  int                  `json:"code"`
			Msg   string               `json:"msg"`
			Words []controlsql.EngWord `json:"words"`
		}
		var response Response
		response.Code = 200
		response.Msg = "成功"
		response.Words = words
		ctx.JSON(http.StatusOK, response)
	})
}
