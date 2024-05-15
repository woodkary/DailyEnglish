package db

import (
	"database/sql"
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
