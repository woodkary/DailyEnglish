package db

import (
	utils "DailyEnglish/utils"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
)

// 定义词性的常量
const (
	Verb         = "verb"
	Noun         = "noun"
	Pronoun      = "pronoun"
	Adjective    = "adjective"
	Adverb       = "adverb"
	Preposition  = "preposition"
	Conjunction  = "conjunction"
	Interjection = "interjection"
)

// 创建词性映射
var posMap = map[string]string{
	"v.":      Verb,
	"n.":      Noun,
	"pron.":   Pronoun,
	"adj.":    Adjective,
	"adv.":    Adverb,
	"prep.":   Preposition,
	"conj.":   Conjunction,
	"interj.": Interjection,
}
var gradeMap = map[int]string{
	1: "小学",
	2: "初中",
	3: "高中",
	4: "四级",
	5: "六级",
	6: "考研",
	7: "托福",
	8: "雅思",
	9: "GRE",
}
var gradeScoreMap = map[int]float64{
	1: 100,
	2: 100,
	3: 25,
	4: 106.5,
	5: 106.5,
	6: 20,
	7: 30,
	8: 9,
	9: 6,
}

// 定义meanings结构体
type Meanings struct {
	Verb         []string `json:"verb"`
	Noun         []string `json:"noun"`
	Pronoun      []string `json:"pronoun"`
	Adjective    []string `json:"adjective"`
	Adverb       []string `json:"adverb"`
	Preposition  []string `json:"preposition"`
	Conjunction  []string `json:"conjunction"`
	Interjection []string `json:"interjection"`
}

// MarshalJSON 自定义序列化方法
func (m Meanings) MarshalJSON() ([]byte, error) {
	type Alias Meanings
	return json.Marshal(&struct {
		Alias
		Verb         []string `json:"verb"`
		Noun         []string `json:"noun"`
		Pronoun      []string `json:"pronoun"`
		Adjective    []string `json:"adjective"`
		Adverb       []string `json:"adverb"`
		Preposition  []string `json:"preposition"`
		Conjunction  []string `json:"conjunction"`
		Interjection []string `json:"interjection"`
	}{
		Alias:        (Alias)(m),
		Verb:         ensureNotNil(m.Verb),
		Noun:         ensureNotNil(m.Noun),
		Pronoun:      ensureNotNil(m.Pronoun),
		Adjective:    ensureNotNil(m.Adjective),
		Adverb:       ensureNotNil(m.Adverb),
		Preposition:  ensureNotNil(m.Preposition),
		Conjunction:  ensureNotNil(m.Conjunction),
		Interjection: ensureNotNil(m.Interjection),
	})
}

// ensureNotNil 确保切片不为 nil
func ensureNotNil(slice []string) []string {
	if slice == nil {
		return []string{}
	}
	return slice
}

// 初始化meanings结构体
func newMeanings() *Meanings {
	return &Meanings{
		Verb:         []string{},
		Noun:         []string{},
		Pronoun:      []string{},
		Adjective:    []string{},
		Adverb:       []string{},
		Preposition:  []string{},
		Conjunction:  []string{},
		Interjection: []string{},
	}
}

// 将输入字符串转换为meanings结构体
func parseMeanings(input string) *Meanings {
	meanings := newMeanings()
	//先根据~号分隔各词性
	parts := strings.Split(input, "~")

	for _, part := range parts {
		//再根据.号分隔词性和词义
		posMeaning := strings.SplitN(part, ":", 2)
		if len(posMeaning) == 2 {
			//将vt和vi转换为v
			if posMeaning[0] == "vt" {
				posMeaning[0] = "v"
			} else if posMeaning[0] == "vi" {
				posMeaning[0] = "v"
			}
			pos := posMeaning[0] + "."
			//最后根据中文逗号分隔词义
			meaning := strings.Split(posMeaning[1], "；")
			if posName, ok := posMap[pos]; ok {
				switch posName {
				case Verb:
					meanings.Verb = meaning
				case Noun:
					meanings.Noun = meaning
				case Pronoun:
					meanings.Pronoun = meaning
				case Adjective:
					meanings.Adjective = meaning
				case Adverb:
					meanings.Adverb = meaning
				case Preposition:
					meanings.Preposition = meaning
				case Conjunction:
					meanings.Conjunction = meaning
				case Interjection:
					meanings.Interjection = meaning
				}
			}
		}
	}
	return meanings
}

type Exam_score struct {
	UserAnswer string `json:"selectedChoice"`
	UserScore  int    `json:"score"`
}
type WordData struct {
	WordID   int    `json:"word_id"`
	Spelling string `json:"spelling"`
	Meanings *Meanings
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

func GetWordByWordId(db *sql.DB, word_id int) (*WordData, error) {
	var pronunciation string
	var meanings string
	var word string
	err := db.QueryRow("SELECT pronunciation,meanings,word FROM word WHERE word_id =?", word_id).Scan(&pronunciation, &meanings, &word)
	if err != nil {
		return nil, err
	}
	// Construct the word data structure
	wordData := &WordData{
		WordID:   word_id,
		Spelling: word,
		Meanings: parseMeanings(meanings),
	}
	return wordData, nil
}

// 定义 Word_Detail 结构体 (对应6-11日 MySQL 数据库中的 Word 表内所有字段, 但不包括word_question)
type Word_Detail struct {
	WordID            int
	Word              string
	Pronounciation    string
	Meanings          string
	Morpheme1         string
	Morpheme2         string
	Word_Meaning1     string
	Sentence1         string
	Sentence_Meaning1 string
	Word_Meaning2     string
	Sentence2         string
	Sentence_Meaning2 string
	Phrase1           string
	Phrase_Meaning1   string
	Phrase2           string
	Phrase_Meaning2   string
	Difficulty        int
}

func GetWordDetailByWordId(db *sql.DB, word_id int) (*Word_Detail, error) {
	var word_detail Word_Detail
	err := db.QueryRow("SELECT word_id,word,pronunciation,meanings,morpheme1,morpheme2,word_meaning1,sentence1,sentence_meaning1,word_meaning2,sentence2,sentence_meaning2,phrase1,phrase_meaning1,phrase2,phrase_meaning2,difficulty FROM word WHERE word_id =?", word_id).Scan(&word_detail.WordID, &word_detail.Word, &word_detail.Pronounciation, &word_detail.Meanings, &word_detail.Morpheme1, &word_detail.Morpheme2, &word_detail.Word_Meaning1, &word_detail.Sentence1, &word_detail.Sentence_Meaning1, &word_detail.Word_Meaning2, &word_detail.Sentence2, &word_detail.Sentence_Meaning2, &word_detail.Phrase1, &word_detail.Phrase_Meaning1, &word_detail.Phrase2, &word_detail.Phrase_Meaning2, &word_detail.Difficulty)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	return &word_detail, nil
}

// 插入用户 数据库字段有username string, email string
func RegisterUser_User(db *sql.DB, username string, password string, email string) error {
	fmt.Print("RegisterUser_User")
	// 准备插入语句
	userid := utils.GenerateID()
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

// 添加用户学习词书信息
func AddUserBook(db *sql.DB, user_id int, book_id int) error {
	// 首先检查是否已经存在相同的记录
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user_study WHERE user_id = ?", user_id).Scan(&count)
	if err != nil {
		return err
	}
	fmt.Println("count:", count)
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
	//默认计划10个单词每天学习
	_, err = stmt.Exec(user_id, book_id, 10, 0)
	if err != nil {
		return err
	}

	return nil
}
func AddUserPunchLearn(db *sql.DB, user_id int) error {
	//向user_punch-learn插入一项记录
	//由于其几乎所有字段都有默认值，所以只需根据user_id插入即可
	stmt2, err := db.Prepare("INSERT INTO `user_punch-learn`(user_id,date) VALUES(?,?)")
	if err != nil {
		return err
	}
	defer stmt2.Close()
	_, err = stmt2.Exec(user_id, utils.GetCurrentDate())
	if err != nil {
		return err
	}
	return nil
}

// 更新用户学习词书信息
func UpdateUserBook(db *sql.DB, user_id int, book_id int) error {
	plan_num := 20 // 默认计划每天学习10个单词
	// 首先检查用户是否已经选择一本词书
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user_study WHERE user_id = ?", user_id).Scan(&count)
	if err != nil {
		return err
	}
	// 如果已经存在，则更新
	if count > 0 {
		stmt, err := db.Prepare("UPDATE user_study SET plan_num =?,study_day =? WHERE user_id =?")
		if err != nil {
			log.Panic(err)
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(plan_num, book_id, user_id)
		if err != nil {
			log.Panic(err)
			return err
		}
	} else {
		// 如果不存在，则插入
		stmt, err := db.Prepare("INSERT INTO user_study(user_id,book_id,plan_num) VALUES(?,?,?)")
		if err != nil {
			log.Panic(err)
			return err
		}
		defer stmt.Close()
		_, err = stmt.Exec(user_id, book_id, plan_num)
		if err != nil {
			log.Panic(err)
			return err
		}
	}
	return nil
}
func GetUserStudy(db *sql.DB, user_id int) (UserStudy, error) {
	var userStudy UserStudy
	var book_id int
	//根据user_id，先查询book_id,plan_num,learn_index
	var plan_num int
	var learn_index int
	err := db.QueryRow("SELECT book_id,plan_num,learn_index FROM user_study WHERE user_id = ?", user_id).Scan(&book_id, &plan_num, &learn_index)
	if err != nil {
		log.Panic(err)
		return userStudy, err
	}
	// 查找该book从learned_index以后plan_num个未学习过的词,并作为punchNum

	query := "SELECT COUNT(*) FROM word_book WHERE book_id = ? AND word_id > ? AND word_id <= ?"
	err = db.QueryRow(query, book_id, learn_index, learn_index+plan_num).Scan(&userStudy.PunchNum)
	if err != nil {
		return userStudy, err
	}
	fmt.Println("打卡总次数punchNum:", userStudy.PunchNum)
	// 查找WordNumLearned_该用户已学词数
	err = db.QueryRow("SELECT book_id,learn_index FROM user_study WHERE user_id =?", user_id).Scan(&book_id, &userStudy.WordNumLearned)
	if err != nil {
		return userStudy, err
	}
	// 查找WordNumTotal_该词书总词数,BookLearning_该词书名称
	err = db.QueryRow("SELECT word_num,book_name FROM book_info WHERE book_id =?", book_id).Scan(&userStudy.WordNumTotal, &userStudy.BookLearning)
	if err != nil {
		return userStudy, err
	}
	//计算Days_left_剩余天数,PunchNum_打卡次数
	userStudy.Days_left = (userStudy.WordNumTotal - userStudy.WordNumLearned) / 10 //每个用户每天计划打卡10个单词——这是固定死的
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
	ExamID      string
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
	fmt.Println(exam_id)
	err := db.QueryRow("SELECT exam_name,exam_date,exam_clock,question_num,question_id FROM exam_info WHERE exam_id =?", exam_id).Scan(&exam.ExamName, &exam.ExamDate, &exam.Exam_clock, &exam.QuestionNum, &exam.QuestionID)
	if err != nil {
		return exam, err
	}
	exam.ExamID = fmt.Sprintf("%d", exam_id)
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

// 从Elasticsearch批量查询题目信息
func GetQuestionsInfoFromES(es *elasticsearch.Client, question_ids []int) ([]QuestionInfo, error) {
	var questionsInfo []QuestionInfo

	/*
		POST /questions/_mget
			{
				"docs": [
					{ "_id": "1" },
					{ "_id": "2" }
				]
			}

	*/
	//由上述指令批量查询question_ids对应的题目信息
	var buf bytes.Buffer
	buf.WriteString(`{"docs":[`)
	for i, id := range question_ids {
		if i > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf(`{"_id":"%d"}`, id))
	}
	buf.WriteString(`]}`)
	//使用mget指令批量查询
	res, err := es.Mget(bytes.NewReader(buf.Bytes()), es.Mget.WithContext(context.Background()), es.Mget.WithIndex("questions"))
	if err != nil {
		return nil, fmt.Errorf("error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("error: %s", res.String())
	}
	//解析返回值
	var bulkResponse struct {
		Docs []struct {
			Found  bool         `json:"found"`
			Source QuestionInfo `json:"_source"`
		} `json:"docs"`
	}

	if err := json.NewDecoder(res.Body).Decode(&bulkResponse); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}
	//读取返回值中的_source字段，并存入questionsInfo
	for _, doc := range bulkResponse.Docs {
		if doc.Found {
			questionsInfo = append(questionsInfo, doc.Source)
		}
	}

	return questionsInfo, nil
}

// 将题目信息批量存入Elasticsearch
func StoreQuestionsInfoToES(es *elasticsearch.Client, questionsInfo []QuestionInfo) error {
	var buf bytes.Buffer
	for _, questionInfo := range questionsInfo {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, questionInfo.Question_id, "\n"))
		data, err := json.Marshal(questionInfo)
		if err != nil {
			return fmt.Errorf("error marshaling question info: %s", err)
		}
		buf.Grow(len(meta) + len(data))
		buf.Write(meta)
		buf.Write(data)
		buf.WriteByte('\n')
	}
	// 使用bulk指令批量插入
	req := esapi.BulkRequest{
		Index: "questions",
		Body:  &buf,
	}

	res, err := req.Do(context.Background(), es)
	if err != nil {
		return fmt.Errorf("error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("error: %s", res.String())
	}

	return nil
}

// 从MySQL查询题目信息并返回
func GetQuestionsInfoFromDB(db *sql.DB, question_ids []int) ([]QuestionInfo, error) {
	var questionsInfo []QuestionInfo

	for _, id := range question_ids {
		questionInfo, err := GetQuestionInfo(db, id)
		if err != nil {
			return nil, fmt.Errorf("error getting question info from DB: %s", err)
		}
		questionsInfo = append(questionsInfo, questionInfo)
	}

	return questionsInfo, nil
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

func GetExamDetail(db *sql.DB, user_id int, exam_id int) ([]QuestionDetail, error) {
	var questionDetails []QuestionDetail
	var questions string

	// 获取问题列表
	err := db.QueryRow("SELECT question_id FROM exam_info WHERE exam_id =?", exam_id).Scan(&questions)
	if err != nil {
		return nil, err
	}
	questions_list := strings.Split(questions, "-")

	// 获取用户答案
	var ans string
	err = db.QueryRow("SELECT user_answer FROM `user-exam_score` WHERE user_id =? AND exam_id =?", user_id, exam_id).Scan(&ans)
	userAnswer := make(map[int]string)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == nil {
		ans_list := strings.Split(ans, "-")
		for i, question := range questions_list {
			question_id, _ := strconv.Atoi(question)
			if i < len(ans_list) {
				userAnswer[question_id] = ans_list[i]
			}
		}
	}

	// 批量查询问题详细信息
	query := "SELECT question_id, question_type, question_content, question_answer FROM question_info WHERE question_id IN (?)"
	query, args, err := sqlx.In(query, questions_list)
	if err != nil {
		return nil, err
	}
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 解析查询结果
	for rows.Next() {
		var questionDetail QuestionDetail
		var content string
		var questionType int

		err := rows.Scan(&questionDetail.Question_id, &questionType, &content, &questionDetail.Answer)
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

		questionDetail.UserAnswer = userAnswer[questionDetail.Question_id]
		if questionDetail.UserAnswer == questionDetail.Answer {
			questionDetail.Score = 5
		} else {
			questionDetail.Score = 0
		}
		questionDetails = append(questionDetails, questionDetail)
	}

	if err = rows.Err(); err != nil {
		return nil, err
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

// 根据wordIDs查询word
func GetWordByWordID(db *sql.DB, wordIDs []int) ([]Word, error) {
	var WordList []Word
	var objectQuestion string
	for _, wordID := range wordIDs {
		var object Word
		object.WordID = wordID
		err := db.QueryRow("SELECT word,pronunciation,word_question,answer FROM word WHERE word_id = ?", wordID).Scan(&object.Word, &object.PhoneticUS, &objectQuestion, &object.Answer)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
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
		err := db.QueryRow("SELECT word,pronunciation,word_question,answer FROM word WHERE word_id = ?", wordID).Scan(&object.Word, &object.PhoneticUS, &objectQuestion, &object.Answer)
		if err != nil {
			log.Panic(err)
			return nil, err
		}
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
func UpdateUserPunch(db *sql.DB, userID int, today string, rdb *redis.Client, punchResult map[int]bool) error {
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
			lastPunchdate = today
		} else {
			return err
		}
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
	//更新用户learn_index和study_day
	var old_index int
	var plan_num int
	var study_day int
	db.QueryRow("SELECT learn_index,plan_num,study_day FROM user_study WHERE user_id = ?", userID).Scan(&old_index, &plan_num, &study_day)
	new_index := old_index + plan_num
	study_day++
	updateQuery, err = db.Prepare("UPDATE user_study SET learn_index = ?,study_day = ? WHERE user_id = ?")
	if err != nil {
		log.Panic(err)
		return err
	}
	defer updateQuery.Close()
	_, err = updateQuery.Exec(new_index, study_day, userID)
	if err != nil {
		log.Panic(err)
		return err
	}
	fmt.Printf("User %d punch record updated successfully.\n", userID)

	//计算isPunch中连续的1的个数，这是用户连续打卡到的天数
	count := 0
	maxCount := 0
	for isPunch > 0 {
		if isPunch&0x01 == 1 {
			count++
			if count > maxCount {
				maxCount = count
			}
		} else {
			count = 0
		}
		isPunch >>= 1
	}
	count = maxCount
	fmt.Println("连续打卡天数:", count)
	//更新user_study表中的continuous_study字段，其值为现有值和count的最大值
	//表示连续打卡天数
	updateQuery, err = db.Prepare("UPDATE user_study SET continuous_study = GREATEST(continuous_study,?) WHERE user_id = ?")
	if err != nil {
		log.Panic(err)
		return err
	}
	defer updateQuery.Close()
	_, err = updateQuery.Exec(count, userID)
	if err != nil {
		log.Panic(err)
		return err
	}

	// 构建键名
	redisKey := fmt.Sprintf("punchResult:%d:%s", userID, today)

	// 准备哈希的键值对
	hashData := make(map[string]interface{})
	for k, v := range punchResult {
		// 由于Redis的哈希字段和值都是字符串，所以这里需要将整数键转换为字符串
		keyStr := strconv.Itoa(k)
		// 布尔值在Go中转换为字符串时，"true"会被转换为"1"，"false"会被转换为"0"
		valueStr := strconv.FormatBool(v)
		hashData[keyStr] = valueStr
	}

	// 使用HMSet将所有键值对存入Redis哈希
	_, err = rdb.HMSet(context.Background(), redisKey, hashData).Result()
	if err != nil {
		log.Panic(err)
		return err
	}

	fmt.Printf("User %d punch record updated successfully.\n", userID)
	return nil
}
func UpdateUserLearnIndex(db *sql.DB, userID int) error {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Fatalf("Recovered from panic: %v", p)
		}
	}()

	// `user_punch-learn`查询当前用户的学习进度
	var learnedIndex int
	err = tx.QueryRow("SELECT learned_index FROM `user_punch-learn` WHERE user_id = ?", userID).Scan(&learnedIndex)
	if err != nil {
		if err == sql.ErrNoRows {
			tx.Rollback()
			return errors.New("请先选择词书")
		}
		tx.Rollback()
		return err
	}

	var planNum int
	//从user_study表中查询该用户计划的词数
	err = tx.QueryRow("SELECT plan_num FROM user_study WHERE user_id = ?", userID).Scan(&planNum)
	if err != nil {
		tx.Rollback()
		log.Panic(err)
		return err
	}
	//更新用户learn_index和study_day
	var old_index int
	var plan_num int
	var study_day int
	db.QueryRow("SELECT learn_index,plan_num,study_day FROM user_study WHERE user_id = ?", userID).Scan(&old_index, &plan_num, &study_day)
	new_index := old_index + plan_num
	study_day++
	updateQuery, err := db.Prepare("UPDATE user_study SET learn_index = ?,study_day = ? WHERE user_id = ?")
	if err != nil {
		log.Panic(err)
		return err
	}
	defer updateQuery.Close()
	_, err = updateQuery.Exec(new_index, study_day, userID)
	if err != nil {
		log.Panic(err)
		return err
	}

	// 将learnedIndex+planNum更新到user_punch-learn，并将user_punch-learn的punch_num加上计划的词数
	updateQuery, err = tx.Prepare("UPDATE `user_punch-learn` SET learned_index = ?, punch_num = punch_num + ? WHERE user_id = ?")
	if err != nil {
		tx.Rollback()
		log.Panic(err)
		return err
	}
	defer updateQuery.Close()

	_, err = updateQuery.Exec(learnedIndex+planNum, planNum, userID)
	if err != nil {
		tx.Rollback()
		log.Panic(err)
		return err
	}

	// 将user_study表中的learn_index也更新为learnedIndex+planNum，同时将study_day+1，表示多学了一天
	updateQuery, err = tx.Prepare("UPDATE user_study SET learn_index = ?, study_day = study_day + 1 WHERE user_id = ?")
	if err != nil {
		tx.Rollback()
		log.Panic(err)
		return err
	}
	defer updateQuery.Close()

	_, err = updateQuery.Exec(learnedIndex+planNum, userID)
	if err != nil {
		tx.Rollback()
		log.Panic(err)
		return err
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Printf("User %d learn index updated successfully.\n", userID)
	return nil
}

// 一位用户提交考试结果，更新该场考试的question_statistics表
func UpdateQuestionStatistics(db *sql.DB, examID int, question_result map[int]string) error {
	// 循环遍历question_result的键question_id
	// 判断question_result的值为A还是B还是C还是D，为对应的字段加1
	for questionID, answer := range question_result {
		switch answer {
		case "A":
			updatequery, _ := db.Prepare("UPDATE question_statistics SET A_num VALUES A_num + 1 WHERE exam_id = ? AND question_id = ?")
			defer updatequery.Close()
			_, err := updatequery.Exec(examID, questionID)
			if err != nil {
				log.Panic(err)
				return err
			}
		case "B":
			updatequery, _ := db.Prepare("UPDATE question_statistics SET B_num VALUES B_num + 1 WHERE exam_id = ? AND question_id = ?")
			defer updatequery.Close()
			_, err := updatequery.Exec(examID, questionID)
			if err != nil {
				log.Panic(err)
				return err
			}
		case "C":
			updatequery, _ := db.Prepare("UPDATE question_statistics SET C_num VALUES C_num + 1 WHERE exam_id = ? AND question_id = ?")
			defer updatequery.Close()
			_, err := updatequery.Exec(examID, questionID)
			if err != nil {
				log.Panic(err)
				return err
			}
		case "D":
			updatequery, _ := db.Prepare("UPDATE question_statistics SET D_num VALUES D_num + 1 WHERE exam_id = ? AND question_id = ?")
			defer updatequery.Close()
			_, err := updatequery.Exec(examID, questionID)
			if err != nil {
				log.Panic(err)
				return err
			}
		default:
			log.Printf("Invalid answer %s for question %d", answer, questionID)
		}
		fmt.Printf("Question %d answer %s updated successfully.\n", questionID, answer)
	}
	return nil
}

// 打卡后插入学习记录并更新复习时间
func InsertUserMemory(db *sql.DB, userID int, wordID int, isCorret bool) error {
	//先根据wordID查询difficulty
	var difficulty int
	err := db.QueryRow("SELECT difficulty FROM word WHERE word_id = ?", wordID).Scan(&difficulty)
	if err != nil {
		log.Panic(err)
		return err
	}
	//计算下次复习时间
	interval_history := "0"
	var feedback_history string
	if isCorret {
		feedback_history = "1"
	} else {
		feedback_history = "0"
	}
	bestInterval := utils.CalculateBestInterval(difficulty, interval_history, feedback_history)
	nextReviewTime := time.Now().AddDate(0, 0, bestInterval).Format("2006-01-02")
	//插入记忆记录
	insertQuery, err := db.Prepare("INSERT INTO user_word_memory(user_id,word_id,learn_times,interval_history,feedback_history,interval_days,is_memory,review_date,difficulty) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Panic(err)
		return err
	}
	defer insertQuery.Close()
	_, err = insertQuery.Exec(userID, wordID, 1, interval_history, feedback_history, bestInterval, 0, nextReviewTime, difficulty)
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

// 复习后更新学习记录
func UpdatetUserMemory(db *sql.DB, userID int, wordID int, isCorret bool) error {
	//先根据wordID查询difficulty,interval_history,feedback_history,review_date,interval_days
	var difficulty int
	var interval_history string
	var feedback_history string
	var review_date string
	var interval_days int
	err := db.QueryRow("SELECT difficulty,interval_history,feedback_history,review_date,interval_days FROM user_word_memory WHERE user_id = ? AND word_id = ?", userID, wordID).Scan(&difficulty, &interval_history, &feedback_history, &review_date, &interval_days)
	if err != nil {
		log.Panic(err)
		return err
	}
	//更新interval_history
	real_reviewDate := time.Now()
	t, _ := time.Parse("2006-01-02", review_date)
	real_interval_days := interval_days + int(real_reviewDate.Sub(t).Hours()/24)
	interval_history = interval_history + "," + strconv.Itoa(real_interval_days)
	//更新feedback_history
	if isCorret {
		feedback_history = feedback_history + "," + "1"
	} else {
		feedback_history = feedback_history + "," + "0"
	}
	bestInterval := utils.CalculateBestInterval(difficulty, interval_history, feedback_history)
	nextReviewTime := time.Now().AddDate(0, 0, bestInterval).Format("2006-01-02")
	//更新记忆记录
	updateQuery, err := db.Prepare("UPDATE user_word_memory SET learn_times = learn_times + 1,interval_history = ?,feedback_history = ?,interval_days = ?,review_date = ? WHERE user_id = ? AND word_id = ?")
	if err != nil {
		log.Panic(err)
		return err
	}
	defer updateQuery.Close()
	_, err = updateQuery.Exec(interval_history, feedback_history, bestInterval, nextReviewTime, userID, wordID)
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}

// redis------user_id:question_type:["score","num"]
// 向redis中插入学生的题目总分和题目数量
func UpdateStudentRDB(db *sql.DB, rdb *redis.Client, userID int, teamID int, examResult map[int]Exam_score) (map[int]float64, error) {
	averageScores := make(map[int]float64)
	ctx := context.Background()
	username := ""
	teamName := ""

	for questionID, questionResult := range examResult {
		var questionType int
		// 查询题目类型
		err := db.QueryRow("SELECT question_type FROM question_info WHERE question_id = ?", questionID).Scan(&questionType)
		if err != nil {
			log.Printf("Failed to query question type for questionID %d: %v", questionID, err)
			continue
		}

		// 构造学生的 Redis 键
		userKey := fmt.Sprintf("studentAverage:%d:%d", userID, questionType)
		// 构造团队的 Redis 键
		teamKey := fmt.Sprintf("teamAverage:%d:%d", teamID, questionType)

		// 使用 Redis 事务确保原子性
		_, err = rdb.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			isEmptyUser := false
			// 获取并更新学生的 score 和 num，如果没有则初始化为 0，并从mysql中查询username
			userScore, err := rdb.HGet(ctx, userKey, "score").Int()
			if err == redis.Nil {
				isEmptyUser = true
				userScore = 0
			} else if err != nil {
				return err
			}

			userNum, err := rdb.HGet(ctx, userKey, "num").Int()
			if err == redis.Nil {
				isEmptyUser = true
				userNum = 0
			} else if err != nil {
				return err
			}
			//如果没有查到username，则从mysql中查询
			if isEmptyUser {
				err = db.QueryRow("SELECT username FROM user_info WHERE user_id = ?", userID).Scan(&username)
				if err != nil {
					log.Printf("Failed to query username for userID %d: %v", userID, err)
					return err
				}
			}

			// 更新学生的 score 和 num
			userScore += questionResult.UserScore
			userNum += 1

			// 设置学生的新的username score 和 num
			if isEmptyUser {
				err = pipe.HSet(ctx, userKey, map[string]interface{}{
					"username": username,
					"score":    userScore,
					"num":      userNum,
				}).Err()
			} else {
				err = pipe.HSet(ctx, userKey, map[string]interface{}{
					"score": userScore,
					"num":   userNum,
				}).Err()
			}
			if err != nil {
				return err
			}
			isEmptyTeam := false

			// 获取并更新团队的 score 和 num
			teamScore, err := rdb.HGet(ctx, teamKey, "score").Int()
			if err == redis.Nil {
				isEmptyTeam = true
				teamScore = 0
			} else if err != nil {
				return err
			}

			teamNum, err := rdb.HGet(ctx, teamKey, "num").Int()
			if err == redis.Nil {
				isEmptyTeam = true
				teamNum = 0
			} else if err != nil {
				return err
			}
			//如果没有查到teamName，则从mysql中查询
			if isEmptyTeam {
				err = db.QueryRow("SELECT team_name FROM team_info WHERE team_id = ?", teamID).Scan(&teamName)
				if err != nil {
					log.Printf("Failed to query teamName for teamID %d: %v", teamID, err)
					return err
				}
			}

			// 更新团队的 score 和 num
			teamScore += questionResult.UserScore
			teamNum += 1

			// 设置团队的新的 score 和 num
			if isEmptyTeam {
				err = pipe.HSet(ctx, teamKey, map[string]interface{}{
					"team_name": teamName,
					"score":     teamScore,
					"num":       teamNum,
				}).Err()
			} else {
				err = pipe.HSet(ctx, teamKey, map[string]interface{}{
					"score": teamScore,
					"num":   teamNum,
				}).Err()
			}
			if err != nil {
				return err
			}

			// 在返回值的 map 中记录学生的平均分
			averageScores[questionType] = float64(userScore) / float64(userNum)
			return nil
		})

		if err != nil {
			log.Printf("Failed to update Redis for questionType %d: %v", questionType, err)
			continue
		}
	}

	return averageScores, nil
}

// 根据wordId获取单词发音和单词拼写
func GetWordInfo(db *sql.DB, wordId int) ([]string, error) {
	var pronunciation string
	var word string
	err := db.QueryRow("SELECT pronunciation,word FROM word WHERE word_id = ?", wordId).Scan(&pronunciation, &word)
	if err != nil {
		return nil, err
	}
	return []string{pronunciation, word}, nil
}

// 根据user_id查询用户是否以及选了词书
func CheckUserBook(db *sql.DB, user_id int) int {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user_study WHERE user_id = ?", user_id).Scan(&count)
	if err != nil {
		log.Panic(err)
		return 0
	}
	if count == 0 {
		return 0
	}
	return 1
}

// 查询review_date<=now的word_id列表
func GetReviewWordID(db *sql.DB, user_id int) ([]int, error) {
	var wordIDs []int
	nowdate := utils.GetCurrentDate()
	rows, err := db.Query("SELECT word_id FROM user_word_memory WHERE user_id = ? AND review_date <= ?", user_id, nowdate)
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
	return wordIDs, nil
}

type UserPunchInfo struct {
	PunchWordNum        int `json:"punch_word_num"`        //打卡单词数
	TotalPunchDay       int `json:"total_punch_day"`       //总打卡天数
	ConsecutivePunchDay int `json:"consecutive_punch_day"` //连续打卡天数
}

// 查询个人中心页面
func GetUserCenter(db *sql.DB, user_id int) (UserPunchInfo, error) {
	var punchWordNum int = 0
	var totalPunchDay int = 0
	var consecutivePunchDay int = 0
	//首先从user_punch-learn查punch_num作为打卡单词数
	err := db.QueryRow("SELECT punch_num FROM `user_punch-learn` WHERE user_id = ?", user_id).Scan(&punchWordNum)
	if err != nil && err != sql.ErrNoRows {
		log.Panic(err)
		return UserPunchInfo{}, err
	}
	//再从user_study查study_day和continuous_study作为总打卡天数和连续打卡天数
	err = db.QueryRow("SELECT study_day,continuous_study FROM user_study WHERE user_id = ?", user_id).Scan(&totalPunchDay, &consecutivePunchDay)
	if err != nil && err != sql.ErrNoRows {
		log.Panic(err)
		return UserPunchInfo{}, err
	}
	return UserPunchInfo{PunchWordNum: punchWordNum, TotalPunchDay: totalPunchDay, ConsecutivePunchDay: consecutivePunchDay}, nil
}

// 查询该用户当前词书是否被打卡完
func CheckUserPunchFinish(db *sql.DB, user_id int, book_id int) (int, error) {
	// 查询user_study的learn_num -- 当前词书的学习进度
	var learn_num int
	err := db.QueryRow("SELECT learn_index FROM user_study WHERE user_id = ?", user_id).Scan(&learn_num)
	if err != nil {
		log.Panic(err)
		return 0, err
	}
	// 查询book_info的word_num -- 当前词库的单词数量
	var word_num int
	err = db.QueryRow("SELECT word_num FROM book_info WHERE book_id = ?", book_id).Scan(&word_num)
	if err != nil {
		log.Panic(err)
		return 0, err
	}

	// 如果learn_num等于word_num，则表示该用户当前词书已被打卡完
	if learn_num == word_num {
		return 1, nil
	}
	return 0, nil
}

type EngWord struct {
	WordID        int       `json:"word_id"`
	Spelling      string    `json:"spelling"`
	Pronunciation string    `json:"pronunciation"`
	Meanings      *Meanings `json:"meanings"`
}

// 先在es中搜索，如果没有搜索到，再在mysql中搜索
// 返回EngWord数组
func SearchWords(db *sql.DB, es *elasticsearch.Client, input string) ([]EngWord, error) {
	ctx := context.Background()

	// Step 1: Search in Elasticsearch
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"wildcard": map[string]interface{}{
				"spelling": fmt.Sprintf("*%s*", input),
			},
		},
	}
	/*查询es中所有包含input的词，返回结果
	POST /dailyenglish/_search
		{
			"query":{
				"wildcard":{
				"spelling":"*input*"
				}
			}
		}
	*/

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Panicf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex("dailyenglish"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		if !strings.Contains(err.Error(), "index_not_found_exception") {
			log.Panicf("Error searching for documents: %s", err)
		} else {
			log.Printf("Index not found: %v", err)
		}
	}

	defer res.Body.Close()

	var words []EngWord

	//检查es是否有结果
	if res.IsError() {
		log.Printf("Failed to search documents: %s", res.Status())
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		}
		//查找json中hits下的hits数组
		hits := r["hits"].(map[string]interface{})["hits"].([]interface{})
		for _, hit := range hits {
			//解析hit下的_source字段
			source := hit.(map[string]interface{})["_source"].(map[string]interface{})
			//_source字段结构和EngWord结构相同，直接解析为EngWord
			word := EngWord{
				WordID:        (int)(source["word_id"].(float64)),
				Spelling:      source["spelling"].(string),
				Pronunciation: source["pronunciation"].(string),
				Meanings:      parseMeaningsFromMap(source["meanings"].(map[string]interface{})),
			}
			words = append(words, word)
			// fmt.Println("从Elasticsearch搜索的单词:", engWordToString(&word))
		}
		if len(words) > 0 {
			return words, nil
		}
	}

	// 如果es没有结果，则在mysql中搜索
	queryStr := fmt.Sprintf("SELECT word_id, word, pronunciation, meanings FROM word WHERE word LIKE '%%%s%%'", input)
	rows, err := db.Query(queryStr)
	if err != nil {
		return nil, fmt.Errorf("MySQL query failed: %w", err)
	}
	defer rows.Close()

	//在搜索的同时，将搜索结果插入到请求体bulkBuf中，准备批量插入es
	var bulkBuf bytes.Buffer
	for rows.Next() {
		var word EngWord
		var meanings string

		if err := rows.Scan(&word.WordID, &word.Spelling, &word.Pronunciation, &meanings); err != nil {
			return nil, fmt.Errorf("Failed to scan MySQL row: " + err.Error())
		}

		word.Meanings = parseMeanings(meanings)
		// fmt.Println("从sql中搜索的单词:", engWordToString(&word))
		words = append(words, word)

		// 先插入{"index":{"_index":"dailyenglish","_id":1}}部分
		meta := []byte(fmt.Sprintf(`{ "index" : { "_index" : "dailyenglish", "_id" : "%d" } }%s`, word.WordID, "\n"))
		bulkBuf.Write(meta)

		doc, err := json.Marshal(word)
		if err != nil {
			return nil, fmt.Errorf("Failed to marshal document" + err.Error())
		}
		// 再插入{"word_id":1,"spelling":"input","pronunciation":"input","meanings":{"verb":[]}}部分
		bulkBuf.Write(doc)
		bulkBuf.Write([]byte("\n"))
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("MySQL rows error: %w", err)
	}

	// 如果bulkBuf有内容，则批量插入es
	if bulkBuf.Len() > 0 {
		res, err := es.Bulk(
			bytes.NewReader(bulkBuf.Bytes()),
			es.Bulk.WithContext(ctx),
			es.Bulk.WithIndex("dailyenglish"),
			es.Bulk.WithRefresh("true"),
		)
		if err != nil {
			return nil, fmt.Errorf("Failed to execute bulk insert: " + err.Error())
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("Bulk insert error: %s", res.String())
		}
	}

	return words, nil
}

func GetThirdPunchResultForDay(rdb *redis.Client, ctx context.Context, userID int, today string) (map[int]bool, error) {
	//先构造punchResult:userID:today的键
	punchResultKey := fmt.Sprintf("punchResult:%d:%s", userID, today)
	//根据punchResultKey获取对应的hash集合，这通常只有一个元素，即今天的打卡结果
	punchResult, err := rdb.HGetAll(ctx, punchResultKey).Result()
	if err != nil {
		return nil, err
	}
	//如果punchResult为空，则说明今天还没有打卡，返回空map
	if len(punchResult) == 0 {
		return nil, nil
	}
	//如果punchResult不为空，则解析出每道题的结果，并返回
	resultMap := make(map[int]bool)
	for key, value := range punchResult {
		//key是题目ID，value是bool值，转换成int
		questionID, err := strconv.Atoi(key)
		if err != nil {
			return nil, fmt.Errorf("failed to parse question id: %s", key)
		}
		result, err := strconv.ParseBool(value)
		if err != nil {
			return nil, fmt.Errorf("failed to parse result: %s", value)
		}
		resultMap[questionID] = result
	}
	return resultMap, nil
}

type WritingTask struct {
	TitleID      string `json:"title_id"`
	Title        string `json:"title"`
	Manager_name string `json:"manager_name"`
	Word_num     string `json:"word_num"`
	Requirement  string `json:"requirement"`
	Publish_date string `json:"publish_date"`
	Submit_date  string `json:"submit_date"`
	Grade        string `json:"grade"`
	Tag          string `json:"tag"`
	Machine_mark int    `json:"score"`
}

func GetUserWritingTask(db *sql.DB, user_id int) ([]WritingTask, []WritingTask, []WritingTask, error) {
	// 定义结果集
	Tasks := []WritingTask{}
	TrainingTasks := []WritingTask{}
	FinishedTasks := []WritingTask{}

	// 查询用户的team_id
	var team_id int
	err := db.QueryRow("SELECT team_id FROM `user-team` WHERE user_id = ?", user_id).Scan(&team_id)
	if err != nil {
		return nil, nil, nil, err
	}

	// 查询团队和系统发布的所有title_ids
	var title_ids []int
	var finish_ids []int

	// 执行合并查询
	rows, err := db.Query(`
        SELECT title_id, team_id 
        FROM composition 
        WHERE team_id = ? OR team_id = 0`, team_id)
	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var title_id, t_id int
		err := rows.Scan(&title_id, &t_id)
		if err != nil {
			return nil, nil, nil, err
		}
		if t_id == team_id {
			title_ids = append(title_ids, title_id)
		} else {
			finish_ids = append(finish_ids, title_id)
		}
	}

	// 定义一个函数来处理任务
	processTasks := func(ids []int, tag string) ([]WritingTask, []WritingTask, error) {
		var tasks []WritingTask
		var finishedTasks []WritingTask

		if len(ids) == 0 {
			return tasks, finishedTasks, nil
		}

		// 构建查询字符串
		query := `
            SELECT c.title_id, c.composition_title, c.manager_id, c.word_num, c.composition_require, 
                   c.publish_date, c.grade, e.respond_date, e.machine_mark
            FROM composition c
            LEFT JOIN composition_evaluate e ON c.title_id = e.title_id AND e.user_id = ?
            WHERE c.title_id IN (` + strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",") + `)`

		args := make([]interface{}, len(ids)+1)
		args[0] = user_id
		for i, id := range ids {
			args[i+1] = id
		}

		// 执行批量查询
		rows, err := db.Query(query, args...)
		if err != nil {
			return nil, nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var task WritingTask
			var manager_id, grade int
			var respond_date sql.NullString
			var machine_mark sql.NullInt32

			err := rows.Scan(&task.TitleID, &task.Title, &manager_id, &task.Word_num, &task.Requirement,
				&task.Publish_date, &grade, &respond_date, &machine_mark)
			if err != nil {
				return nil, nil, err
			}

			task.Grade = gradeMap[grade]
			task.Tag = tag

			if tag == "任务" {
				err = db.QueryRow("SELECT manager_name FROM manager_info WHERE manager_id = ?", manager_id).Scan(&task.Manager_name)
				if err != nil {
					return nil, nil, err
				}
			} else {
				task.Manager_name = "系统"
			}

			if respond_date.Valid {
				task.Submit_date = respond_date.String
				task.Machine_mark = int(machine_mark.Int32)
				finishedTasks = append(finishedTasks, task)
			} else {
				tasks = append(tasks, task)
			}
		}

		return tasks, finishedTasks, nil
	}

	// 处理团队任务
	teamTasks, finishedTeamTasks, err := processTasks(title_ids, "任务")
	if err != nil {
		return nil, nil, nil, err
	}
	Tasks = append(Tasks, teamTasks...)
	FinishedTasks = append(FinishedTasks, finishedTeamTasks...)

	// 处理训练任务
	trainingTasks, finishedTrainingTasks, err := processTasks(finish_ids, "训练")
	if err != nil {
		return nil, nil, nil, err
	}
	TrainingTasks = append(TrainingTasks, trainingTasks...)
	FinishedTasks = append(FinishedTasks, finishedTrainingTasks...)

	// 按照提交日期对已完成任务进行排序
	sort.Slice(FinishedTasks, func(i, j int) bool {
		iTime, _ := time.Parse("2006-01-02", FinishedTasks[i].Submit_date)
		jTime, _ := time.Parse("2006-01-02", FinishedTasks[j].Submit_date)
		return iTime.After(jTime)
	})

	return Tasks, TrainingTasks, FinishedTasks, nil
}

type EssayResult struct {
	TitleID    int    `json:"title_id"`
	Title      string `json:"title"`
	RawEssay   string `json:"raw_essay"`
	Word_cnt   string `json:"word_cnt"`
	Requirment string `json:"requirment"`
	//机器评分与评价
	Machine_mark     int    `json:"machine_mark"`
	Machine_evaluate string `json:"machine_evaluate"`
	//教师评分与评价
	Teacher_mark     int              `json:"teacher_mark"`
	Teacher_evaluate string           `json:"teacher_evaluate"`
	MajorScore       utils.MajorScore `json:"majorScore"`
	//逐句点评
	SentsFeedback []SentsFeedback `json:"sents_feedback"`
}
type SentsFeedback struct {
	ParaId                int                   `json:"para_id"`
	SentId                int                   `json:"sent_id"`
	RawSent               string                `json:"raw_sent"`
	ErrorPosInfos         []utils.ErrorPosInfos `json:"errorPosInfos"`
	SentFeedback          string                `json:"sent_feedback"`
	CorrectedSent         string                `json:"corrected_sent"`
	IsContainGrammarError bool                  `json:"is_contain_grammar_error"`
	IsValidLangSent       bool                  `json:"is_valid_lang_sent"`
}

func GetEssayResult(db *sql.DB, title_ID int, user_id int) (EssayResult, error) {
	essayResult := EssayResult{}

	// 合并查询获取composition和composition_evaluate的数据
	query := `
        SELECT c.composition_title, c.word_num, c.composition_require, e.evaluate_id, e.machine_mark, e.machine_evaluate, 
               e.teacher_mark, e.teacher_evaluate, e.major_score, e.rawessay
        FROM composition c
        LEFT JOIN composition_evaluate e ON c.title_id = e.title_id
        WHERE c.title_id = ? AND e.user_id = ?`

	var major_score string
	var evaluate_id int

	err := db.QueryRow(query, title_ID, user_id).Scan(&essayResult.Title, &essayResult.Word_cnt, &essayResult.Requirment,
		&evaluate_id, &essayResult.Machine_mark, &essayResult.Machine_evaluate,
		&essayResult.Teacher_mark, &essayResult.Teacher_evaluate, &major_score, &essayResult.RawEssay)
	if err != nil {
		return EssayResult{}, err
	}

	// 将major_score JSON字符串转为MajorScore结构
	err = json.Unmarshal([]byte(major_score), &essayResult.MajorScore)
	if err != nil {
		return essayResult, err
	}

	// 查询evaluate_everysentence表中的所有数据
	rows, err := db.Query(`
        SELECT paraid, sentid, rawsent, errorposinfo, sentfeedback, correctedsent, is_containgrammarerror, is_validlangsent 
        FROM evaluate_everysentence 
        WHERE evaluate_id = ?`, evaluate_id)
	if err != nil {
		return essayResult, err
	}
	defer rows.Close()

	var sentsFeedback []SentsFeedback
	for rows.Next() {
		var sentFeedback SentsFeedback
		var errorPosInfo string

		err := rows.Scan(&sentFeedback.ParaId, &sentFeedback.SentId, &sentFeedback.RawSent, &errorPosInfo, &sentFeedback.SentFeedback,
			&sentFeedback.CorrectedSent, &sentFeedback.IsContainGrammarError, &sentFeedback.IsValidLangSent)
		if err != nil {
			return essayResult, err
		}

		// 将errorPosInfo JSON字符串转为ErrorPosInfos结构
		err = json.Unmarshal([]byte(errorPosInfo), &sentFeedback.ErrorPosInfos)
		if err != nil {
			return essayResult, err
		}

		sentsFeedback = append(sentsFeedback, sentFeedback)
	}
	essayResult.SentsFeedback = sentsFeedback

	return essayResult, nil
}

// 说实话，实在懒得再把这些搬进utils包里了，直接写在这里吧。
func parseMeaningsFromMap(meanings map[string]interface{}) *Meanings {
	meaningMap := Meanings{
		Verb:         toStringSlice(meanings["verb"]),
		Noun:         toStringSlice(meanings["noun"]),
		Pronoun:      toStringSlice(meanings["pronoun"]),
		Adjective:    toStringSlice(meanings["adjective"]),
		Adverb:       toStringSlice(meanings["adverb"]),
		Preposition:  toStringSlice(meanings["preposition"]),
		Conjunction:  toStringSlice(meanings["conjunction"]),
		Interjection: toStringSlice(meanings["interjection"]),
	}
	return &meaningMap
}

// toStringSlice converts an interface to a []string.
func toStringSlice(value interface{}) []string {
	if value == nil {
		return []string{}
	}
	if slice, ok := value.([]interface{}); ok {
		result := make([]string, len(slice))
		for i, v := range slice {
			result[i] = v.(string)
		}
		return result
	}
	return []string{}
}

// 向数据库存入一段机器评分作文的数据
func InsertEssayScore(db *sql.DB, userID int, titleID int, url string, result utils.Response) error {
	//先插入一条记录到composition_score表
	insertQuery, err := db.Prepare("INSERT INTO composition_evaluate(user_id,title_id,composition_url,respond_date,machine_evaluate,machine_mark,rawessay,major_score) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Panic(err)
		return err
	}
	defer insertQuery.Close()
	var grade int
	err = db.QueryRow("SELECT grade FROM composition WHERE title_id = ?", titleID).Scan(&grade)
	if err != nil {
		log.Panic(err)
		return err
	}
	//将分数转换为百分制
	TotalScore := int(result.Result.TotalScore / gradeScoreMap[grade] * 100)
	//将result.Result.EssayFeedback.MajorScore json格式转为字符串
	majorScore, err := json.Marshal(result.Result.MajorScore)
	if err != nil {
		log.Panic(err)
		return err
	}
	_, err = insertQuery.Exec(userID, titleID, url, time.Now().Format("2006-01-02"), result.Result.EssayAdvice, TotalScore, result.Result.RawEssay, string(majorScore))
	if err != nil {
		log.Panic(err)
		return err
	}
	//查询刚插入的记录的id
	var scoreID int
	err = db.QueryRow("SELECT evaluate_id FROM composition_evaluate WHERE user_id = ? AND title_id = ? AND composition_url = ?", userID, titleID, url).Scan(&scoreID)
	if err != nil {
		log.Panic(err)
		return err
	}
	//再更新evaluate_everysentence表
	for _, sentence := range result.Result.EssayFeedback.SentsFeedback {
		if len(sentence.ErrorPosInfos) > 0 {
			insertQuery, err := db.Prepare("INSERT INTO evaluate_everysentence(evaluate_id,paraid,sentid,rawsent,errorposinfo,sentfeedback,correctedsent,is_containgrammarerror,is_validlangsent) VALUES(?,?,?,?,?,?,?,?,?)")
			if err != nil {
				log.Panic(err)
				return err
			}
			defer insertQuery.Close()
			var is_containgrammarerror int
			var is_validlangsent int
			if sentence.IsContainGrammarError {
				is_containgrammarerror = 1
			} else {
				is_containgrammarerror = 0
			}
			if sentence.IsValidLangSent {
				is_validlangsent = 1
			} else {
				is_validlangsent = 0
			}
			errorPosInfo, err := json.Marshal(sentence.ErrorPosInfos)
			if err != nil {
				log.Panic(err)
				return err
			}
			_, err = insertQuery.Exec(scoreID, sentence.ParaId, sentence.SentId, sentence.RawSent, string(errorPosInfo), sentence.SentFeedback, sentence.CorrectedSent, is_containgrammarerror, is_validlangsent)
			if err != nil {
				log.Panic(err)
				return err
			}
		}

	}
	//composition表的提交人数submit_count+1
	_, err = db.Exec("UPDATE composition SET submit_count = submit_count + 1 WHERE title_id = ?", titleID)
	if err != nil {
		log.Panic(err)
		return err
	}
	return nil
}
func GetEssayTitle(db *sql.DB, titleId int) (string, int, error) {
	var title string
	var grade int
	fmt.Print("titleId:", titleId)
	err := db.QueryRow("SELECT composition_title,grade FROM composition WHERE title_id = ?", titleId).Scan(&title, &grade)
	if err != nil {
		return "", 0, err
	}
	return title, grade, nil
}
