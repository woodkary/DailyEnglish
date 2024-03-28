package controlsql

import (
	"database/sql"
)

// BookInfo 结构体用于存储书籍信息
type Books struct {
	Title      string
	LearnerNum int
	FinishNum  int
	Type       int
	Describe   string
	ID         int
	WordsNum   int
}

// CET4WordInfo 结构体用于存储 CET4 单词信息
type Cet4_dictionary struct {
	Words     string
	WordID    int
	Soundmark string
	Describe  string
	Question1 string
	Question2 string
}

// MistakeInfo 结构体用于存储错题信息
type Mistake struct {
	Username string
	Question string
}

// NotebookInfo 结构体用于存储单词收藏本信息
type Notebook struct {
	Words    string
	Username string
}

// UserStudyInfo 结构体用于存储用户学习信息
type User_Study struct {
	Username   string
	LearnBook  string
	FinishBook string
	WordsNum   int
	WordsIndex int
}

// QueryUserInfo 查询用户信息
func QueryUser_Info(db *sql.DB) ([]UserInfo, error) {
	rows, err := db.Query("SELECT * FROM user_info")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userInfos []UserInfo
	for rows.Next() {
		var userInfo UserInfo
		if err := rows.Scan(&userInfo.Username, &userInfo.ID, &userInfo.Phone, &userInfo.Pwd, &userInfo.Email, &userInfo.Age, &userInfo.Sex, &userInfo.RegisterDate); err != nil {
			return nil, err
		}
		userInfos = append(userInfos, userInfo)
	}
	return userInfos, nil
}

// QueryBookInfo 查询书籍信息
func QueryBooks(db *sql.DB) ([]Books, error) {
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookInfos []Books
	for rows.Next() {
		var bookInfo Books
		if err := rows.Scan(&bookInfo.Title, &bookInfo.LearnerNum, &bookInfo.FinishNum, &bookInfo.Type, &bookInfo.Describe, &bookInfo.ID, &bookInfo.WordsNum); err != nil {
			return nil, err
		}
		bookInfos = append(bookInfos, bookInfo)
	}
	return bookInfos, nil
}

// QueryCET4WordInfo 查询 CET4 单词信息
func QueryCet4_dictionary(db *sql.DB) ([]Cet4_dictionary, error) {
	rows, err := db.Query("SELECT * FROM cet4_dictionary")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cet4WordInfos []Cet4_dictionary
	for rows.Next() {
		var cet4WordInfo Cet4_dictionary
		if err := rows.Scan(&cet4WordInfo.Words, &cet4WordInfo.WordID, &cet4WordInfo.Soundmark, &cet4WordInfo.Describe, &cet4WordInfo.Question1, &cet4WordInfo.Question2); err != nil {
			return nil, err
		}
		cet4WordInfos = append(cet4WordInfos, cet4WordInfo)
	}
	return cet4WordInfos, nil
}

// QueryMistakeInfo 查询错题信息
func QueryMistake(db *sql.DB) ([]Mistake, error) {
	rows, err := db.Query("SELECT * FROM mistakes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mistakeInfos []Mistake
	for rows.Next() {
		var mistakeInfo Mistake
		if err := rows.Scan(&mistakeInfo.Username, &mistakeInfo.Question); err != nil {
			return nil, err
		}
		mistakeInfos = append(mistakeInfos, mistakeInfo)
	}
	return mistakeInfos, nil
}

// QueryNotebookInfo 查询单词收藏本信息
func QueryNotebook(db *sql.DB) ([]Notebook, error) {
	rows, err := db.Query("SELECT * FROM notebook")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notebookInfos []Notebook
	for rows.Next() {
		var notebookInfo Notebook
		if err := rows.Scan(&notebookInfo.Words, &notebookInfo.Username); err != nil {
			return nil, err
		}
		notebookInfos = append(notebookInfos, notebookInfo)
	}
	return notebookInfos, nil
}

// QueryUserStudyInfo 查询用户学习信息
func QueryUser_Study(db *sql.DB) ([]User_Study, error) {
	rows, err := db.Query("SELECT * FROM user_study")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userStudyInfos []User_Study
	for rows.Next() {
		var userStudyInfo User_Study
		if err := rows.Scan(&userStudyInfo.Username, &userStudyInfo.LearnBook, &userStudyInfo.FinishBook, &userStudyInfo.WordsNum, &userStudyInfo.WordsIndex); err != nil {
			return nil, err
		}
		userStudyInfos = append(userStudyInfos, userStudyInfo)
	}
	return userStudyInfos, nil
}
