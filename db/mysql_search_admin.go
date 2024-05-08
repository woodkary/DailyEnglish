package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func EmailIsRegistered(db *sql.DB, email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE email =?", email).Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func UserExists(db *sql.DB, username string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE username =?", username).Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

// 插入用户 数据库字段有username string, email string, pwd string, sex int, phone string, birthday date, register_date date
// RegisterUser 向 user_info 表中插入用户数据
func RegisterUser(db *sql.DB, username, email, password string, sex int, phone string, birthday string, registerDate string) error {
	// 准备插入语句
	stmt, err := db.Prepare("INSERT INTO user_info(username, email, pwd, sex, phone, birthday, register_date) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 执行插入语句
	_, err = stmt.Exec(username, email, password, sex, phone, birthday, registerDate)
	if err != nil {
		return err
	}

	return nil
}
