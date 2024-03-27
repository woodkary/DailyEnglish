package main

import (
	controlsql "DailyEnglish/Control_SQL"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, er := sql.Open("mysql", "root:123456@tcp(47.107.81.75:3306)/daily_english")
	if er != nil {
		log.Fatal(er)
	}
	defer db.Close()
	users, err := controlsql.QueryUserInfo(db)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(users); i++ {
		fmt.Println(users[i])
	}
}
