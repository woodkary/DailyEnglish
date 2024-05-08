package controlsql

import (
	"database/sql"
	"fmt"
)

// DeleteUserInfo 删除用户信息
func DeleteUserInfo(db *sql.DB, username string) error {
	// 准备删除语句
	query := "DELETE FROM user_info WHERE username = ?"
	// 执行删除操作
	_, err := db.Exec(query, username)
	if err != nil {
		return fmt.Errorf("error deleting user info: %v", err)
	}
	return nil
}

// DeleteBookInfo 删除书籍信息
func DeleteBookInfo(db *sql.DB, title string) error {
	// 准备删除语句
	query := "DELETE FROM books WHERE title = ?"
	// 执行删除操作
	_, err := db.Exec(query, title)
	if err != nil {
		return fmt.Errorf("error deleting book info: %v", err)
	}
	return nil
}

// DeleteCET4Word 删除 CET4 单词信息
func DeleteCET4Word(db *sql.DB, word string) error {
	// 准备删除语句
	query := "DELETE FROM cet4_dictionary WHERE words = ?"
	// 执行删除操作
	_, err := db.Exec(query, word)
	if err != nil {
		return fmt.Errorf("error deleting CET4 word: %v", err)
	}
	return nil
}

// DeleteMistake 删除错题信息
func DeleteMistake(db *sql.DB, username, question string) error {
	// 准备删除语句
	query := "DELETE FROM mistakes WHERE username = ? AND question = ?"
	// 执行删除操作
	_, err := db.Exec(query, username, question)
	if err != nil {
		return fmt.Errorf("error deleting mistake: %v", err)
	}
	return nil
}

// DeleteNotebookWord 删除生词信息
func DeleteNotebookWord(db *sql.DB, username, words string) error {
	// 准备删除语句
	query := "DELETE FROM notebook WHERE username = ? AND words = ?"
	// 执行删除操作
	_, err := db.Exec(query, username, words)
	if err != nil {
		return fmt.Errorf("error deleting notebook word: %v", err)
	}
	return nil
}

// DeleteUserStudyInfo 删除用户学习信息
func DeleteUserStudyInfo(db *sql.DB, username string) error {
	// 准备删除语句
	query := "DELETE FROM user_study WHERE username = ?"
	// 执行删除操作
	_, err := db.Exec(query, username)
	if err != nil {
		return fmt.Errorf("error deleting user study info: %v", err)
	}
	return nil
}
