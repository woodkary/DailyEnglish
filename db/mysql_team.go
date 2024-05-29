package db

import (
	utils "DailyEnglish/utils"
	"database/sql"
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
func JoinTeam(db *sql.DB, username string, teamid int) (bool, error) {
	//先判断团队是否满人
	// SQL 查询语句
	query1 := "SELECT COUNt(*) FROM user-team WHERE  team_id = ?"

	var count1 int
	// 执行查询
	err := db.QueryRow(query1, teamid).Scan(&count1)
	if err != nil {
		return false, err
	}
	// SQL 查询语句
	query2 := "SELECT member_num FROM team_info WHERE  team_id = ?"

	var count2 int
	// 执行查询
	err = db.QueryRow(query2, teamid).Scan(&count2)
	if err != nil {
		return false, err
	}
	//团队已满不可加入
	if count1 == count2 {
		return false, err
	}

	//未满则用户可加团队
	userid, errr := GetUserID(db, username)
	if errr != nil {
		return false, errr
	}
	// 获取当前日期
	now := time.Now()
	// 格式化日期为字符串
	today := now.Format("2006-01-02")
	stmt, err := db.Prepare("INSERT INTO user-team(user_id,team_id,join_date) values (?,?,?)")
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
