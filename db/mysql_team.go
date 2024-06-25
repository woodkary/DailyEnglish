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
	teamid := utils.GenerateID()
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
	Teamid      int      `json:"team_id"`
	Managerid   int      `json:"manager_id"`
	Teamname    string   `json:"team_name"`
	Managername string   `json:"manager_name"`
	Teamsize    int      `json:"member_num"`
	Memberlist  []Member `json:"member_list"`
}
type Member struct {
	Userid   int    `json:"user_id"`
	Username string `json:"user_name"`
	Usersex  int    `json:"user_sex"`
}

func SearchTeamInfo(db *sql.DB, teamid int) (Team, error) {
	var team Team
	team.Teamid = teamid

	// 合并查询获取 team 信息和 manager 信息
	query := `
		SELECT t.manager_id, t.team_name, m.manager_name
		FROM team_info t
		JOIN manager_info m ON t.manager_id = m.manager_id
		WHERE t.team_id = ?`
	err := db.QueryRow(query, teamid).Scan(&team.Managerid, &team.Teamname, &team.Managername)
	if err != nil {
		return Team{}, fmt.Errorf("failed to get team and manager info: %v", err)
	}

	// 获取团队大小
	err = db.QueryRow("SELECT COUNT(*) FROM `user-team` WHERE team_id = ?", teamid).Scan(&team.Teamsize)
	if err != nil {
		return Team{}, fmt.Errorf("failed to get team size: %v", err)
	}

	// 获取团队成员信息
	query = `
		SELECT ui.user_id, ui.username, ui.sex
		FROM user_info ui
		JOIN ` + "`user-team`" + ` ut ON ui.user_id = ut.user_id
		WHERE ut.team_id = ?`
	rows, err := db.Query(query, teamid)
	if err != nil {
		return Team{}, fmt.Errorf("failed to get team members: %v", err)
	}
	defer rows.Close()

	var users []Member
	for rows.Next() {
		var user Member
		if err := rows.Scan(&user.Userid, &user.Username, &user.Usersex); err != nil {
			return Team{}, fmt.Errorf("failed to scan user info: %v", err)
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return Team{}, fmt.Errorf("rows iteration error: %v", err)
	}

	team.Memberlist = users

	return team, nil
}
