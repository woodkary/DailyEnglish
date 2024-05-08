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
func RegisterUser(db *sql.DB, username string, email string, pwd string, sex int, phone string, birthday date, register_date date) int {
	stmt, err := db.Prepare("INSERT INTO users(username, email, pwd, sex, phone, birthday, register_date) VALUES(?,?)")
	if err != nil {
		return 0
	}
	res, err := stmt.Exec(username, email, pwd, sex, phone, birthday, register_date)
	if err != nil {
		return 0
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0
	}
	return int(id)
}
