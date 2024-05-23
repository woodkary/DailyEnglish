package db

import (
	utils "DailyEnglish/utils"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func EmailIsRegistered_TeamManager(db *sql.DB, email string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM manager_info WHERE email =?", email).Scan(&count)
	if err != nil {
		return false
	}
	if count == 0 {
		return false
	}
	return true
}

func AdminManagerExists(db *sql.DB, username string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM manager_info WHERE manager_name =?", username).Scan(&count)
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
func RegisterUser(db *sql.DB, username, email, password string, phone string) error {
	// 准备插入语句
	//userid := utils.GenerateID()
	userid := 114514
	stmt, err := db.Prepare("INSERT INTO manager_info(manager_id ,manager_name, email, phone, pwd) VALUES( ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 执行插入语句
	_, err = stmt.Exec(userid, username, email, phone, password)
	if err != nil {
		return err
	}

	return nil
}

// 验证用户密码正确性
func CheckTeamManager(db *sql.DB, username, password string) bool {
	var row string
	db.QueryRow("SELECT pwd FROM manager_info WHERE manager_name =?", username).Scan(&row)
	utils.TestAES()

	fmt.Println("row:", row)
	fmt.Println("password:", password)
	decryptrow := utils.AesDecrypt(row, "123456781234567812345678")
	fmt.Println("decryptrow:", decryptrow)

	return password == decryptrow
}

// 根据username获取map[user_id]username和map[team_id]team_name
func GetTokenParams_TeamManager(db *sql.DB, username string) (int, map[int]string, error) {
	var managerID int
	team := make(map[int]string)

	// 查询数据库以获取 manager_id
	err := db.QueryRow("SELECT manager_id FROM manager_info WHERE manager_name = ?", username).Scan(&managerID)
	if err != nil {
		return 0, nil, err
	}
	// 查询数据库以获取与 manager_id 相关的 team_id 列表
	rows, err := db.Query("SELECT team_id,team_name FROM team_info WHERE manager_id = ?", managerID)
	if err != nil {
		return 0, nil, err
	}
	defer rows.Close()

	// 遍历结果集，将 team_id 存储到数组中
	for rows.Next() {
		var teamID int
		var teamName string
		if err := rows.Scan(&teamID, &teamName); err != nil {
			return 0, nil, err
		}
		team[teamID] = teamName
	}
	if err := rows.Err(); err != nil {
		return 0, nil, err
	}

	return managerID, team, nil
}
