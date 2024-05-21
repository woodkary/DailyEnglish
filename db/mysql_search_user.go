package db

import (
	utils "DailyEnglish/utils"
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
func CheckUser_User(db *sql.DB, username, password string) bool {
	var row string
	db.QueryRow("SELECT pwd FROM user_info WHERE username =?", username).Scan(&row)
	utils.TestAES()

	decryptrow := utils.AesDecrypt(row, "123456781234567812345678")

	return password == decryptrow
}
