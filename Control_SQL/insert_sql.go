package controlsql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

// InsertUserInfo 向用户基础信息表中插入用户信息
func InsertUserInfo(db *sql.DB) error {
	// 准备 SQL 语句
	query := "INSERT INTO user_info (username, id, phone, pwd, email, age, sex, register_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	// 准备参数
	username := "小明"
	id := 123                             // 假设id为123
	phone := "123456789"                  // 假设手机号为123456789
	pwd := "12344"                        // 假设密码
	email := "xiaoming@example.com"       // 假设邮箱为xiaoming@example.com
	age := 25                             // 假设年龄为25
	sex := 1                              // 假设性别为男性（1代表男性）
	registerDate := "2024-03-28 12:00:00" // 假设注册时间为2024-03-28 12:00:00

	// 执行 SQL 语句
	_, err := db.Exec(query, username, id, phone, pwd, email, age, sex, registerDate)
	if err != nil {
		return err
	}

	return nil
}
