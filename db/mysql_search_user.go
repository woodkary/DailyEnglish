package db

import (
	utils "DailyEnglish/utils"
	"database/sql"
	"strconv"
	"strings"
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
	// 准备插入语句
	userid := utils.GenerateID()
	//userid := 114514
	stmt, err := db.Prepare("INSERT INTO manager_info(manager_id ,manager_name, email, pwd) VALUES( ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 执行插入语句
	_, err = stmt.Exec(userid, username, email, password)
	if err != nil {
		return err
	}

	return nil
}

// 验证用户密码正确性
func CheckUser_User(db *sql.DB, username, password string) bool {
	var row string
	db.QueryRow("SELECT pwd FROM user_info WHERE username =?", username).Scan(&row)
	utils.TestAES()

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
	err := db.QueryRow("SELECT team_id FROM user_team WHERE use_id =?", user_id).Scan(&team_id)
	if err != nil {
		return 0, "", err
	}
	err = db.QueryRow("SELECT team_name FROM team_info WHERE team_id =?", team_id).Scan(&team_name)
	if err != nil {
		return 0, "", err
	}
	return team_id, team_name, nil
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
		err = db.QueryRow("SELECT exam_score,exam_rank from user-exam score WHERE user_id =?,exam_id =?", use_id, exam.ExamID).Scan(&exam.ExamScore, exam.ExamRank)
		if err != nil {
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

// 根据user_id和exam_id查询单场考试详情
func GetExamDetail(db *sql.DB, user_id int, exam_id int) ([]QuestionDetail, error) {
	var questionDetails []QuestionDetail
	var questions string
	err := db.QueryRow("SELECT question_id FROM question_info WHERE exam_id =?", exam_id).Scan(&questions)
	if err != nil {
		return nil, err
	}
	questions_list := strings.Split(questions, "-")
	userAnwser := make(map[int]string)
	var ans string
	err = db.QueryRow("SELECT user_answer from user-exam_score WHERE user_id =?,exam_id =?", user_id, exam_id).Scan(&ans)
	if err != nil {
		return nil, err
	}
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
		err := db.QueryRow("SELECT question_content,quetion_answer FROM question_info WHERE question_id =?", questionid).Scan(&content, &questionDetail.Answer)
		content_list := strings.Split(content, "：")
		questionDetail.Question = content_list[0]
		questionDetail.Options = strings.Split(content_list[1], " ")
		if err != nil {
			return nil, err
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
