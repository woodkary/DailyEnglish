package controlsql

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql" // 导入 MySQL 驱动
)

// SearchUserByUsername 按用户名查找用户是否存在，存在则返回true，不存在则返回false
func SearchUserByUsername(db *sql.DB, username string) bool {
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM user_info WHERE username = ?)", username).Scan(&exists)
	if err != nil {
		return false // 查询过程中出现错误，假定用户不存在
	}
	return exists
}

// InsertUser 向用户基础信息表插入用户信息
func InsertUser(db *sql.DB, username string, pwd string, email string) error {
	_, err := db.Exec("INSERT INTO user_info (username, pwd, email) VALUES (?, ?, ?)", username, pwd, email)
	return err
}

// CheckUser 检查用户名和密码是否匹配，匹配则返回true，不匹配则返回false
func CheckUser(db *sql.DB, username string, pwd string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM user_info WHERE username = ? AND pwd = ?", username, pwd).Scan(&count)
	if err != nil || count != 1 {
		return false
	}
	return true
}

// InsertRandomData 向所有表中插入随机数据
func InsertRandomData(db *sql.DB) error {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 向用户基础信息表插入随机数据
	for i := 1; i <= 10; i++ {
		username := fmt.Sprintf("user%d", i)
		phone := fmt.Sprintf("12345678%d", i)
		pwd := fmt.Sprintf("password%d", i)
		email := fmt.Sprintf("user%d@example.com", i)
		age := rand.Intn(50) + 18
		sex := rand.Intn(2)
		registerDate := time.Now().Format("2006-01-02 15:04:05")

		_, err := db.Exec("INSERT INTO user_info (username, id, phone, pwd, email, age, sex, register_date) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
			username, i, phone, pwd, email, age, sex, registerDate)
		if err != nil {
			return err
		}
	}

	// 向团队表插入随机数据
	for i := 1; i <= 5; i++ {
		teamname := fmt.Sprintf("team%d", i)
		teamID := i
		teammate := sql.NullString{String: fmt.Sprintf("member%d_member%d", rand.Intn(10)+1, rand.Intn(10)+1), Valid: true}
		teamManager := sql.NullString{String: fmt.Sprintf("manager%d", rand.Intn(10)+1), Valid: true}

		_, err := db.Exec("INSERT INTO team (teamname, team_id, teammate, team_manager) VALUES (?, ?, ?, ?)",
			teamname, teamID, teammate, teamManager)
		if err != nil {
			return err
		}
	}

	// 向系统管理员表插入随机数据
	for i := 1; i <= 3; i++ {
		username := fmt.Sprintf("admin%d", i)
		id := i

		_, err := db.Exec("INSERT INTO system_manager (username, id) VALUES (?, ?)",
			username, id)
		if err != nil {
			return err
		}
	}

	return nil
}

// UserInfo 结构体用于存储用户信息
type UserInfo struct {
	Username     string
	ID           int
	Phone        string
	Pwd          string
	Email        string
	Age          int
	Sex          int
	RegisterDate string
}

// TeamInfo 结构体用于存储团队信息
type TeamInfo struct {
	Teamname    string
	TeamID      int
	Teammate    sql.NullString
	TeamManager sql.NullString
}

// SystemManager 结构体用于存储系统管理员信息
type SystemManager struct {
	Username string
	ID       int
}

// QueryUserInfo 查询用户基础信息表
func QueryUserInfo(db *sql.DB) ([]UserInfo, error) {
	rows, err := db.Query("SELECT * FROM user_info;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userInfos []UserInfo
	for rows.Next() {
		var userInfo UserInfo
		if err := rows.Scan(&userInfo.Username, &userInfo.ID, &userInfo.Phone, &userInfo.Pwd, &userInfo.Email, &userInfo.Age, &userInfo.Sex, &userInfo.RegisterDate); err != nil {
			return nil, err
		}
		userInfos = append(userInfos, userInfo)
	}

	return userInfos, nil
}

// QueryTeam 查询团队表
func QueryTeam(db *sql.DB) ([]TeamInfo, error) {
	rows, err := db.Query("SELECT * FROM team;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teamInfos []TeamInfo
	for rows.Next() {
		var teamInfo TeamInfo
		if err := rows.Scan(&teamInfo.Teamname, &teamInfo.TeamID, &teamInfo.Teammate, &teamInfo.TeamManager); err != nil {
			return nil, err
		}
		teamInfos = append(teamInfos, teamInfo)
	}

	return teamInfos, nil
}

// QuerySystemManager 查询系统管理员名单
func QuerySystemManager(db *sql.DB) ([]SystemManager, error) {
	rows, err := db.Query("SELECT * FROM system_manager;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var systemManagers []SystemManager
	for rows.Next() {
		var systemManager SystemManager
		if err := rows.Scan(&systemManager.Username, &systemManager.ID); err != nil {
			return nil, err
		}
		systemManagers = append(systemManagers, systemManager)
	}

	return systemManagers, nil
}
