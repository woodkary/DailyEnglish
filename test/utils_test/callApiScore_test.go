package test

import (
	utils "DailyEnglish/utils"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestImageCorrect(t *testing.T) {
	Base64IMG, err := utils.ReadFileAsBase64("D:/屏幕截图 2024-06-17 133710.png")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}

	result := utils.CallApiScore(Base64IMG, 6, "A Proposal to the School Library")
	// 打印返回结果
	fmt.Print(result)
}
func TestGetImg(t *testing.T) {
	username := "mimahezhanghao1yang"
	password := "MIMAhezhanghao1yang"
	hostname := "rm-wz9p61j3qlj6lg69fpo.mysql.rds.aliyuncs.com"
	port := "3306"
	dbname := "dailyenglish"

	// 构建数据库连接字符串
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)

	// 连接数据库
	err := error(nil)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	utils.GetOSSSecret(db)
	url, err := utils.GetImageFromOSS("2024/06/18/1718688821.jpg")
	fmt.Println(url)
	if err != nil {
		log.Fatal(err)
	}
}
