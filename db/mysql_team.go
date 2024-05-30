package db

import (
	utils "DailyEnglish/utils"
	"database/sql"
	"fmt"
	"log"
	"time"
)

// （新建团队）插入团队信息
func RegisterTeam(db *sql.DB, teamname string, managerid, maxnum int) error {
	// 准备插入语句
	teamid := utils.GenerateID(time.Now(), 1)
	// 获取当前日期
	now := time.Now()
	// 格式化日期为字符串
	today := now.Format("2006-01-02")

	stmt, err := db.Prepare("INSERT INTO team_info(team_id,manager_id,team_name, member_num,build_date) VALUES( ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 执行插入语句
	_, err = stmt.Exec(teamid, managerid, teamname, maxnum, today)
	if err != nil {
		return err
	}

	return nil
}

//删除团队
//DELETE FROM employees WHERE employee_id = 123;

func DeleteTeam(db *sql.DB, teamid int) error {

	stmt, err := db.Prepare("DELETE FROM team_info WHERE team_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// 执行插入语句
	_, err = stmt.Exec(teamid)
	if err != nil {
		return err
	}

	return nil
}

// 删除团队成员
func DeleteTeammember(db *sql.DB, teamid int, username string) error {

	userid, errr := GetUserID(db, username)
	if errr != nil {
		return errr
	}
	stmt, err := db.Prepare("DELETE FROM user-team WHERE user_id = ? AND team_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userid, teamid)
	if err != nil {
		return err
	}

	return nil
}

// 查找成员是否存在，返回bool
func CheckTeammember(db *sql.DB, username string, teamid int) (bool, error) {
	userid, errr := GetUserID(db, username)
	if errr != nil {
		return false, errr
	}
	// SQL 查询语句
	query := "SELECT COUNt(*) FROM user-team WHERE user_id = ? AND team_id = ?"

	var count int
	// 执行查询
	err := db.QueryRow(query, userid, teamid).Scan(&count)
	if err != nil {
		return false, err
	}

	// 如果 count 大于 0，说明记录存在
	return count > 0, nil
}

// 用户加入团队
func JoinTeam(db *sql.DB, userid int, teamid int) (bool, error) {

	// 获取当前日期
	now := time.Now()
	// 格式化日期为字符串
	today := now.Format("2006-01-02")
	stmt, err := db.Prepare("INSERT INTO `user-team` (user_id,team_id,join_date) values (?,?,?)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	// 执行插入语句
	_, err = stmt.Exec(userid, teamid, today)
	if err != nil {
		return false, err
	}

	return true, nil
}

// 根据团队id查询团队是否存在
func CheckTeam(db *sql.DB, teamid int) (bool, error) {

	// SQL 查询语句
	query := "SELECT COUNT(*) FROM team_info WHERE team_id = ?"

	var count int
	// 执行查询
	err := db.QueryRow(query, teamid).Scan(&count)
	if err != nil {
		log.Panic(err)
		return false, err
	}

	// 如果 count 大于 0，说明team存在
	return count > 0, nil
}

// 查询团队是否已满
func IsTeamFull(db *sql.DB, teamid int) (bool, error) {
	// 查询member_num
	query := "SELECT member_num FROM team_info WHERE  team_id = ?"
	var member_num int
	err := db.QueryRow(query, teamid).Scan(&member_num)
	if err != nil {
		return false, err
	}

	// count team内成员数量
	query = "SELECT COUNT(*) FROM user-team WHERE  team_id = ?"
	var count int
	err = db.QueryRow(query, teamid).Scan(&count)
	if err != nil {
		return false, err
	}
	return count == member_num, nil
}

// 查询团队信息

type Team struct {
	Teamid      int
	Managerid   int
	Teamname    string
	Managername string
	Teamsize    int
	Memberlist  []Member
}
type Member struct {
	Userid   int
	Username string
	Usersex  int
}

func SearchTeamInfo(db *sql.DB, teamid int) (Team, error) {
	var team Team
	fmt.Println("teamId是", teamid)
	err := db.QueryRow("SELECT manager_id,team_name FROM team_info WHERE team_id = ?").Scan(&team.Managerid, &team.Teamname)
	// 查询数据库以获取信息
	team.Teamid = teamid
	if err != nil {
		log.Panic(err)
		return Team{}, err
	}
	err = db.QueryRow("SELECT manager_name FROM manager_info WHERE manager_id = ?", team.Managerid).Scan(&team.Managername)
	if err != nil {
		log.Panic(err)
		return Team{}, err
	}
	err = db.QueryRow("SELECT COUNT(*) FROM user-team WHERE team_id = ?", teamid).Scan(&team.Teamsize)
	if err != nil {
		log.Panic(err)
		return Team{}, err
	}

	var users []Member
	// 查询数据库以获取用户名称
	rows, err := db.Query("SELECT user_id  FROM user-team WHERE team_id = ?", teamid)
	if err != nil {
		log.Panic(err)
		return Team{}, err
	}
	defer rows.Close()
	var user Member
	// 遍历结果集并收集
	for rows.Next() {
		var userID int

		if err := rows.Scan(&userID); err != nil {
			return Team{}, err
		}
		err = db.QueryRow("SELECT user_id,username,sex FROM user_info WHERE user_id = ?", userID).Scan(&user.Userid, &user.Username, &user.Usersex)
		if err != nil {
			log.Panic(err)
			return Team{}, err
		}

		users = append(users, user)
	}
	team.Memberlist = users
	return team, nil

}
