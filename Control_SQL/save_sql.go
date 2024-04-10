package controlsql

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

// InsertUserInfo 向用户基础信息表插入用户数据
func InsertUserInfo(db *sql.DB, username, phone, pwd, email string, id, age, sex int, registerDate string) error {
	// 准备插入语句
	query := "INSERT INTO user_info (username, id, phone, pwd, email, age, sex, register_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	// 执行插入操作
	_, err := db.Exec(query, username, id, phone, pwd, email, age, sex, registerDate)
	if err != nil {
		return fmt.Errorf("error inserting user info: %v", err)
	}
	return nil
}

// InsertBookInfo 向书籍表插入书籍数据
func InsertBooks(db *sql.DB, title, describe, grade, difficulty, date string, learnerNum, finishNum, bookType, id, wordsNum int) error {
	// 准备插入语句
	query := "INSERT INTO books (title, learner_num, finish_num, type, `describe`, id, words_num, grade, difficulty, date) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	// 执行插入操作
	_, err := db.Exec(query, title, learnerNum, finishNum, bookType, describe, id, wordsNum, grade, difficulty, date)
	if err != nil {
		return fmt.Errorf("error inserting book info: %v", err)
	}
	return nil
}

// InsertCET4Word 向 CET4 单词表插入单词数据
func InsertCet4_dictionary(db *sql.DB, words string, wordID int, soundmark, describe, question1, question2 string) error {
	// 准备插入语句
	query := "INSERT INTO cet4_dictionary (words, word_id, soundmark, `describe`, question_1, question_2) VALUES (?, ?, ?, ?, ?, ?)"
	// 执行插入操作
	_, err := db.Exec(query, words, wordID, soundmark, describe, question1, question2)
	if err != nil {
		return fmt.Errorf("error inserting CET4 word: %v", err)
	}
	return nil
}

// InsertMistake 向错题表插入单个错题数据
func InsertMistake(db *sql.DB, username, question string) error {
	// 准备插入语句
	query := "INSERT INTO mistakes (username, question) VALUES (?, ?)"
	// 执行插入操作
	_, err := db.Exec(query, username, question)
	if err != nil {
		return fmt.Errorf("error inserting mistake: %v", err)
	}
	return nil
}

// InsertMistakes 向错题表插入多条错题数据，用于一次性插入多个错题
func InsertMistakes(db *sql.DB, username string, questions []string) error {
	// 准备插入语句
	query := "INSERT INTO mistakes (username, question) VALUES "
	// 构建插入值
	var valueStrings []string
	var values []interface{}
	for _, q := range questions {
		valueStrings = append(valueStrings, "(?, ?)")
		values = append(values, username, q)
	}
	query += strings.Join(valueStrings, ", ")
	// 执行插入操作
	_, err := db.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting mistakes: %v", err)
	}
	return nil
}

// InsertNotebook 向单词收藏本插入生词数据
func InsertNotebook(db *sql.DB, username string, words []string) error {
	// 准备插入语句
	query := "INSERT INTO notebook (words, username) VALUES "
	// 构建插入值
	var valueStrings []string
	var values []interface{}
	for _, word := range words {
		valueStrings = append(valueStrings, "(?, ?)")
		values = append(values, word, username)
	}
	query += strings.Join(valueStrings, ", ")
	// 执行插入操作
	_, err := db.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error inserting words to notebook: %v", err)
	}
	return nil
}

// InsertUserStudyInfo 向用户学习信息表插入用户学习信息
func InsertUserStudyInfo(db *sql.DB, username, learnBook, finishBook string, wordsNum, wordsIndex int) error {
	// 准备插入语句
	query := "INSERT INTO user_study (username, learn_book, finish_book, words_num, words_index) VALUES (?, ?, ?, ?, ?)"
	// 执行插入操作
	_, err := db.Exec(query, username, learnBook, finishBook, wordsNum, wordsIndex)
	if err != nil {
		return fmt.Errorf("error inserting user study info: %v", err)
	}
	return nil
}
