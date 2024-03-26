package controlsql

import (
	"database/sql"
	"fmt"
)

func QueryTables(db *sql.DB) {
	rows, err := db.Query("SHOW TABLES;")
	if err != nil {
		// 直接在这里打印错误信息
		fmt.Println("Error querying tables:", err)
		return
	}
	defer rows.Close()

	var tableName string
	for rows.Next() {
		if err := rows.Scan(&tableName); err != nil {
			// 直接在这里打印错误信息
			fmt.Println("Error scanning table name:", err)
			return
		}
		// 直接在这里打印表名
		fmt.Println("Table found:", tableName)
	}

	// 直接在这里检查是否在循环结束后还有错误
	if err := rows.Err(); err != nil {
		fmt.Println("Error processing rows:", err)
	}
}
