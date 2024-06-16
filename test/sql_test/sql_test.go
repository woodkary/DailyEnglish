package test

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestWordBook(t *testing.T) {
	// 数据库连接信息
	username := "mimahezhanghao1yang"
	password := "MIMAhezhanghao1yang"
	hostname := "rm-wz9p61j3qlj6lg69fpo.mysql.rds.aliyuncs.com"
	port := "3306"
	dbname := "dailyenglish"
	var db *sql.DB

	// 构建数据库连接字符串
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)

	// 连接数据库
	err := error(nil)
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// 执行SQL语句
	for i := 1; i <= 100; i++ {
		rows, err := db.Query("INSERT INTO word_book (word_id,book_id) VALUES (?,4)", i)
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()
	}
	fmt.Println("Insert 100 rows into word_book table successfully.")
}
