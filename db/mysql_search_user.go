package db

import (
	utils "DailyEnglish/utils"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

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
	err := db.QueryRow("SELECT book_id,plan_num,learned_num FROM user_study WHERE user_id =?", user_id).Scan(&book_id, &userStudy.PunchNum, &userStudy.WordNumLearned)
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
func GetExamInfo(db *sql.DB, use_id int, team_id int) ([]Exam, error) {
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
		err = db.QueryRow("SELECT exam_score,exam_rank from `user-exam_score` WHERE user_id =? AND exam_id =?", use_id, exam.ExamID).Scan(&exam.ExamScore, &exam.ExamRank)
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
	if err.Error() != "sql: no rows in result set" {
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
	if err.Error() == "Error 1062 (23000): Duplicate entry '32-29364224' for key 'user-exam_score.PRIMARY'" {
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

// 查询用户是否已有打卡计划
func GetPunchPlan(db *sql.DB, user_id int) (bool, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user_study WHERE user_id =?", user_id).Scan(&count)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
