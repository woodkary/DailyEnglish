package db

import (
	utils "DailyEnglish/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type Exam_score struct {
	UserAnswer string `json:"selectedChoice"`
	UserScore  int    `json:"score"`
}

// 根据email查询user是否存在
func EmailIsRegistered_User(db *sql.DB, email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user_info WHERE email =?", email).Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

// 根据username查询user是否存在
func UserExists_User(db *sql.DB, username string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user_info WHERE username =?", username).Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func GetWordByWordId(db *sql.DB, word_id int) (map[string]interface{}, error) {
	var pronunciation string
	var meanings string
	var word string
	err := db.QueryRow("SELECT pronunciation,meanings,word FROM word WHERE word_id =?", word_id).Scan(&pronunciation, &meanings, &word)
	if err != nil {
		return nil, err
	}
	// Construct the word data structure
	wordData := map[string]interface{}{
		"word_id":       word_id,
		"spelling":      word,
		"pronunciation": pronunciation,
		"meanings":      meanings,
	}
	return wordData, nil
}

// 插入用户 数据库字段有username string, email string
func RegisterUser_User(db *sql.DB, username string, password string, email string) error {
	fmt.Print("RegisterUser_User")
	// 准备插入语句
	userid := utils.GenerateID(time.Now(), 1145141919810)
	stmt, err := db.Prepare("INSERT INTO user_info(user_id ,username, email, pwd,sex,phone,birthday,register_date) VALUES( ?, ?, ?, ?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 执行插入语句
	_, err = stmt.Exec(userid, username, email, password, -1, "12345678901", "2000-01-01", utils.GetCurrentDate())
	if err != nil {
		return err
	}

	return nil
}

// 验证用户密码正确性
func CheckUser_User(db *sql.DB, username, password string) bool {
	var row string
	db.QueryRow("SELECT pwd FROM user_info WHERE username =?", username).Scan(&row)

	decryptrow := utils.AesDecrypt(row, "123456781234567812345678")

	return password == decryptrow
}

// 根据username查询userid
func GetUserID(db *sql.DB, username string) (int, error) {
	var userid int
	err := db.QueryRow("SELECT user_id FROM user_info WHERE username =?", username).Scan(&userid)
	if err != nil {
		return 0, err
	}
	return userid, nil
}

// 根据userid查询team_id和team_name
func GetTokenParams_User(db *sql.DB, user_id int) (int, string, error) {
	var team_id int
	var team_name string
	err := db.QueryRow("SELECT team_id FROM `user-team` WHERE user_id =?", user_id).Scan(&team_id)
	if err != nil {
		return 0, "", err
	}
	err = db.QueryRow("SELECT team_name FROM team_info WHERE team_id =?", team_id).Scan(&team_name)
	if err != nil {
		return 0, "", err
	}
	return team_id, team_name, nil
}

// 词书
type BookInfo struct {
	BookID    int
	BookName  string
	WordNum   int
	Diffculty int
	Grade     int
	Describe  string
}

// 查询所有词书
func GetAllBooks(db *sql.DB) ([]BookInfo, error) {
	var books []BookInfo
	rows, err := db.Query("SELECT book_id,book_name,grade,word_num,difficulty,`describe` FROM book_info")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var book BookInfo
		err := rows.Scan(&book.BookID, &book.BookName, &book.Grade, &book.WordNum, &book.Diffculty, &book.Describe)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// 根据user_id查询user_study用户学习信息
type UserStudy struct {
	BookLearning   string
	WordNumTotal   int
	WordNumLearned int
	Days_left      int
	PunchNum       int
	IsPunched      bool
}

// 添加用户学习信息
func AddUserBook(db *sql.DB, user_id int, book_id int) error {
	// 首先检查是否已经存在相同的记录
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user_study WHERE user_id = ? AND book_id = ?", user_id, book_id).Scan(&count)
	if err != nil {
		return err
	}

	// 如果已经存在记录，则返回"已完成"
	if count > 0 {
		return errors.New("已完成")
	}

	// 如果不存在，则准备并执行插入语句
	stmt, err := db.Prepare("INSERT INTO user_study(user_id,book_id,plan_num,study_day) VALUES(?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user_id, book_id, 20, 0)
	if err != nil {
		return err
	}

	return nil
}
func GetUserStudy(db *sql.DB, user_id int) (UserStudy, error) {
	var userStudy UserStudy
	var book_id int
	err := db.QueryRow("SELECT book_id,plan_num,study_day FROM user_study WHERE user_id =?", user_id).Scan(&book_id, &userStudy.PunchNum, &userStudy.WordNumLearned)
	if err != nil {
		return userStudy, err
	}
	err = db.QueryRow("SELECT word_num,book_name FROM book_info WHERE book_id =?", book_id).Scan(&userStudy.WordNumTotal, &userStudy.BookLearning)
	if err != nil {
		return userStudy, err
	}
	userStudy.Days_left = (userStudy.WordNumTotal - userStudy.WordNumLearned) / userStudy.PunchNum
	var date string
	err = db.QueryRow("SELECT last_punchdate FROM user_punch WHERE user_id =?", user_id).Scan(&date)
	if err != nil {
		if err == sql.ErrNoRows {
			// 说明当前用户没有打卡记录，则返回默认值
			log.Println("当前的用户可能一次打卡都没有")
			userStudy.IsPunched = false
		}
		return userStudy, err
	}
	userStudy.IsPunched = date == utils.GetCurrentDate()

	return userStudy, nil
}

// Exam_info
type Exam struct {
	ExamID      int
	ExamName    string
	Exam_clock  string
	ExamDate    string
	QuestionNum int
	ExamScore   int
	ExamRank    int
	QuestionID  string
}

// 根据exam_id查询exam_info
func GetExamInfoByExamID(db *sql.DB, exam_id int) (Exam, error) {
	var exam Exam
	err := db.QueryRow("SELECT exam_name,exam_date,exam_clock,question_num,question_id FROM exam_info WHERE exam_id =?", exam_id).Scan(&exam.ExamName, &exam.ExamDate, &exam.Exam_clock, &exam.QuestionNum, &exam.QuestionID)
	if err != nil {
		return exam, err
	}
	return exam, nil
}

// 根据user_id,team_id查询考试信息
func GetExamInfo(db *sql.DB, user_id int, team_id int) ([]Exam, error) {
	var exams []Exam
	rows, err := db.Query("SELECT exam_id,exam_name,exam_date,exam_clock,question_num FROM exam_info WHERE team_id =?", team_id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var exam Exam
		err := rows.Scan(&exam.ExamID, &exam.ExamName, &exam.ExamDate, &exam.Exam_clock, &exam.QuestionNum)
		if err != nil {
			return nil, err
		}
		//每个exam_id查examRank和examScore
		err = db.QueryRow("SELECT exam_score,exam_rank from `user-exam_score` WHERE user_id =? AND exam_id =?", user_id, exam.ExamID).Scan(&exam.ExamScore, &exam.ExamRank)
		if err != nil && err.Error() != "sql: no rows in result set" {
			return nil, err
		}
		exams = append(exams, exam)
	}
	return exams, nil
}

type QuestionDetail struct {
	Question_id int
	Question    string
	Answer      string
	UserAnswer  string
	Options     []string
	Score       int
}
type QuestionInfo struct {
	Question_id        int
	Questiontype       int
	QuestionDifficulty int
	QuestionContent    string
	QuestionAnswer     string
	QuestionGrade      int
	Options            map[string]string
}

// 根据question_id查询questiondetail
func GetQuestionInfo(db *sql.DB, question_id int) (QuestionInfo, error) {
	var question_info QuestionInfo
	var question_type, question_diffculty, question_grade int
	var content string
	var answer string
	err := db.QueryRow("SELECT question_type,question_difficulty,question_grade,question_content,question_answer FROM question_info WHERE question_id =?", question_id).Scan(&question_type, &question_diffculty, &question_grade, &content, &answer)
	if err != nil {
		return QuestionInfo{}, err
	}
	if question_type == 1 { //选择题
		content_list := strings.Split(content, "：")
		question_info.QuestionContent = content_list[0]
		question_info.Options = make(map[string]string)
		options := strings.Split(content_list[1], " ")
		for _, option := range options {
			m := strings.Split(option, ".")
			question_info.Options[m[0]] = m[1]
		}
		question_info.QuestionAnswer = answer
		question_info.Question_id = question_id
		question_info.Questiontype = question_type
		question_info.QuestionDifficulty = question_diffculty
		question_info.QuestionGrade = question_grade
		return question_info, nil
	} else if question_type == 2 { //填空题
		content_list := strings.Split(content, "：")
		question_info.QuestionContent = content_list[0]
		question_info.QuestionAnswer = answer
		question_info.Question_id = question_id
		question_info.Questiontype = question_type
		question_info.QuestionDifficulty = question_diffculty
		question_info.QuestionGrade = question_grade
		question_info.Options = make(map[string]string)
		return question_info, nil
	}
	return QuestionInfo{}, nil

}

// 根据user_id和exam_id查询单场考试详情
func GetExamDetail(db *sql.DB, user_id int, exam_id int) ([]QuestionDetail, error) {
	var questionDetails []QuestionDetail
	var questions string
	err := db.QueryRow("SELECT question_id FROM exam_info WHERE exam_id =?", exam_id).Scan(&questions)
	fmt.Println(exam_id)
	if err != nil {
		return nil, err
	}
	questions_list := strings.Split(questions, "-")
	userAnwser := make(map[int]string)
	var ans string
	err = db.QueryRow("SELECT user_answer from `user-exam_score` WHERE user_id =? AND exam_id =?", user_id, exam_id).Scan(&ans)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	if err != sql.ErrNoRows {
		ans_list := strings.Split(ans, "-")
		for _, item := range ans_list {
			a := strings.Split(item, ":")
			question_id, _ := strconv.Atoi(a[0])
			answer := a[1]
			userAnwser[question_id] = answer
		}
		for _, question := range questions_list {
			var questionDetail QuestionDetail
			questionid, _ := strconv.Atoi(question)
			questionDetail.Question_id = questionid
			var content string
			var questionType int //1选择题 2填空题
			err := db.QueryRow("SELECT question_type,question_content,quetion_answer FROM question_info WHERE question_id =?", questionid).Scan(&questionType, &content, &questionDetail.Answer)
			if err != nil {
				return nil, err
			}
			if questionType == 1 {
				content_list := strings.Split(content, "：")
				questionDetail.Question = content_list[0]
				questionDetail.Options = strings.Split(content_list[1], " ")
			} else if questionType == 2 {
				questionDetail.Question = content
				questionDetail.Options = []string{""}
			}
			questionDetail.UserAnswer = userAnwser[questionid]
			//每个question_id查score
			if userAnwser[questionid] == questionDetail.Answer {
				questionDetail.Score = 5
			} else {
				questionDetail.Score = 0
			}
			questionDetails = append(questionDetails, questionDetail)
		}
		return questionDetails, nil
	}
	for _, question := range questions_list {
		var questionDetail QuestionDetail
		questionid, _ := strconv.Atoi(question)
		questionDetail.Question_id = questionid
		var content string
		var questionType int //1选择题 2填空题
		err := db.QueryRow("SELECT question_type,question_content,question_answer FROM question_info WHERE question_id =?", questionid).Scan(&questionType, &content, &questionDetail.Answer)
		if err != nil {
			return nil, err
		}
		if questionType == 1 {
			content_list := strings.Split(content, "：")
			questionDetail.Question = content_list[0]
			questionDetail.Options = strings.Split(content_list[1], " ")
		} else if questionType == 2 {
			questionDetail.Question = content
			questionDetail.Options = []string{""}
		}
		questionDetail.UserAnswer = "未作答"
		//每个question_id查score
		questionDetail.Score = 0
		questionDetails = append(questionDetails, questionDetail)
	}
	return questionDetails, nil
}
func InsertUserScore(db *sql.DB, user_id int, exam_id int, user_answer string, score int) error {
	stmt, err := db.Prepare("INSERT INTO `user-exam_score`(user_id,exam_id,exam_score,user_answer,exam_rank) VALUES(?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user_id, exam_id, score, user_answer, 0)
	if err != nil && err.Error() != "Error 1062 (23000): Duplicate entry '32-29364224' for key 'user-exam_score.PRIMARY'" {
		return err
	}
	if err != nil && err.Error() == "Error 1062 (23000): Duplicate entry '32-29364224' for key 'user-exam_score.PRIMARY'" {
		stmt, err := db.Prepare("UPDATE `user-exam_score` SET exam_score = ?,user_answer = ? WHERE user_id = ? AND exam_id = ?")
		if err != nil {
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(score, user_answer, user_id, exam_id)
		if err != nil {
			return err
		}
	}
	return nil
}

type Word struct {
	WordID       int               `json:"word_id"`
	Word         string            `json:"word"`
	PhoneticUS   string            `json:"phonetic_us"`
	WordQuestion map[string]string `json:"word_question"`
	Answer       string            `json:"answer"`
}

// 从数据库中查询，并且生成用户打卡内容
func GetUserPunchContent(db *sql.DB, userID int) ([]Word, error) {
	// 查询用户当前学习的bookID
	// 查询用户计划的打卡词数
	var bookID int
	var plan_num int
	var learn_index int
	err := db.QueryRow("SELECT book_id,plan_num,learn_index FROM user_study WHERE user_id = ?", userID).Scan(&bookID, &plan_num, &learn_index)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	// 查找该book从learned_index以后plan_num个未学习过的词的WordID
	var wordIDs []int
	query := "SELECT word_id FROM word_book WHERE book_id = ? AND word_id > ? AND word_id <= ?"
	rows, err := db.Query(query, bookID, learn_index, learn_index+plan_num)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	for rows.Next() {
		var wordID int
		err := rows.Scan(&wordID)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
		wordIDs = append(wordIDs, wordID)
	}

	// 查询每个 wordID 对应的 Word 对象
	var WordList []Word
	var objectQuestion string
	for _, wordID := range wordIDs {
		var object Word
		object.WordID = wordID
		err := db.QueryRow("SELECT word,phonetic_us,word_question,answer FROM word WHERE word_id = ?", wordID).Scan(&object.Word, &object.PhoneticUS, &objectQuestion, &object.Answer)
		if err != nil {
			log.Panic(err)
			return nil, err
		}

		fmt.Println("objectQuestion: ", objectQuestion, "parsed: ", objectQuestion[1:])
		// 将objectQuestion字符串的首位：字符忽略，并以空格划分为四个子字符串，形如A.1 B.2 C.3 D.4
		objectQuestionList := strings.Split(objectQuestion[1:], " ")
		object.WordQuestion = make(map[string]string)
		for _, question := range objectQuestionList {
			m := strings.Split(question, ".")
			object.WordQuestion[m[0]] = m[1]
		}
		fmt.Println(object.WordQuestion)

		WordList = append(WordList, object)
	}
	return WordList, nil
}

// 更新一次打卡信息
func UpdateUserPunch(db *sql.DB, userID int, today string) error {
	// 查询当前用户的打卡记录
	query := "SELECT is_punch, last_punchdate FROM user_punch WHERE user_id = ?"
	var isPunch int64
	var lastPunchdate string
	err := db.QueryRow(query, userID).Scan(&isPunch, &lastPunchdate)
	if err != nil {
		if err == sql.ErrNoRows {
			// 说明当前用户没有打卡记录，则插入一条新的记录
			insertQuery, err := db.Prepare("INSERT INTO user_punch(user_id, is_punch, last_punchdate) VALUES(?,?,?)")
			if err != nil {
				log.Panic(err)
				return err
			}
			defer insertQuery.Close()
			_, err = insertQuery.Exec(userID, 0x01, today)
			if err != nil {
				log.Panic(err)
				return err
			}
		}
		return err
	}

	// 解析最后打卡日期
	lastPunchTime, err := time.Parse("2006-01-02", lastPunchdate)
	if err != nil {
		log.Panic(err)
		return err
	}

	// 计算今天和最后打卡日期之间的天数差
	todayTime, err := time.Parse("2006-01-02", today)
	if err != nil {
		log.Panic(err)
		return err
	}
	dayDiff := int(todayTime.Sub(lastPunchTime).Hours() / 24)
	fmt.Println("dayDiff:", dayDiff)

	// 根据天数差移位
	isPunch <<= int64(dayDiff)
	//把最低位设为1，表示今天打卡
	isPunch |= 1
	fmt.Println("isPunch:", isPunch)

	// 更新数据库中的记录
	updateQuery, err := db.Prepare("UPDATE user_punch SET is_punch = ?, last_punchdate = ? WHERE user_id = ?")
	fmt.Println("00000000", updateQuery)
	if err != nil {
		log.Panic(err)
		return err
	}
	fmt.Println("1111111", updateQuery)
	defer updateQuery.Close()
	_, err = updateQuery.Exec(isPunch, today, userID)
	if err != nil {
		log.Panic(err)
		return err
	}

	fmt.Printf("User %d punch record updated successfully.\n", userID)
	return nil
}

// redis------user_id:question_type:["score","num"]
// 向redis中插入学生的题目总分和题目数量
func UpdateStudentRDB(db *sql.DB, rdb *redis.Client, userID int, examResult map[int]Exam_score) (map[int]float64, error) {
	averageScores := make(map[int]float64)
	ctx := context.Background()
	//先遍历map，获取题目id和题目分数
	for questionID, questionResult := range examResult {
		var questionType int
		// 查询题目类型
		err := db.QueryRow("SELECT question_type FROM question_info WHERE question_id = ?", questionID).Scan(&questionType)
		if err != nil {
			log.Printf("Failed to query question type for questionID %d: %v", questionID, err)
			continue
		}

		// 构造 Redis 键
		key := fmt.Sprintf("%d:%d", userID, questionType)
		fmt.Println("key:", key)

		// 使用 Redis 事务确保原子性
		_, err = rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			// 获取当前的 score 和 num
			score, err := rdb.HGet(ctx, key, "score").Int()
			if err == redis.Nil {
				score = 0
			} else if err != nil {
				return err
			}
			fmt.Println("score:", score)

			num, err := rdb.HGet(ctx, key, "num").Int()
			if err == redis.Nil {
				num = 0
			} else if err != nil {
				return err
			}
			fmt.Println("num:", num)

			// 更新 score 和 num
			score += questionResult.UserScore
			num += 1

			// 设置新的 score 和 num
			err = pipe.HSet(ctx, key, map[string]interface{}{
				"score": score,
				"num":   num,
			}).Err()
			if err != nil {
				return err
			}
			fmt.Println("已更新redis", score)
			//在返回值的map中记录平均分
			averageScores[questionType] = float64(score) / float64(num)
			return nil
		})

		if err != nil {
			log.Printf("Failed to update Redis for questionType %d: %v", questionType, err)
			continue
		}
	}

	return averageScores, nil
}
