package main

import (
	"DailyEnglish/middlewares"
	adminrouter "DailyEnglish/router/admin_router"
	teamrouter "DailyEnglish/router/team_router"
	userrouter "DailyEnglish/router/user_router"
	"database/sql"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db  *sql.DB
	rdb *redis.Client
	es  *elasticsearch.Client
)

func init() {
	// 数据库连接信息
	username := "mimahezhanghao1yang"
	password := "MIMAhezhanghao1yang"
	hostname := "rm-wz9p61j3qlj6lg69fpo.mysql.rds.aliyuncs.com"
	port := "3306"
	dbname := "dailyenglish"

	// 构建数据库连接字符串
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, hostname, port, dbname)

	// 连接数据库
	err := error(nil)
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	// redis 连接信息
	rdb = redis.NewClient(&redis.Options{
		Addr:     "r-bp1jdmrszl1yd6xxdipd.redis.rds.aliyuncs.com:6379", // Redis服务器地址
		Username: "mimahezhanghao1yang",                                // Redis账号
		Password: "MIMAhezhanghao1yang",                                // Redis密码
		DB:       0,                                                    // 选择的数据库
	})

	// 测试Redis连接
	_, err = rdb.Ping(rdb.Context()).Result()
	if err != nil {
		panic(err)
	}

	// 连接es
	esURL := "https://8af9afd9e4bf4d88b97b14488467361d.us-central1.gcp.cloud.es.io"
	esAPIKey := "SEZ3cUI1QUJaclpXZ01wZGhPckE6UlRkcjZXeENRQjJXaEhISnF2eTBZQQ=="
	cfg := elasticsearch.Config{
		APIKey: esAPIKey,
		Addresses: []string{
			esURL,
		},
	}
	es, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	//测试es连接
	_, err = es.Ping()
	if err != nil {
		log.Fatalf("Error pinging Elasticsearch: %s", err)
	}
}

func main() {
	defer rdb.Close()
	defer db.Close()
	defer es.Indices.Close([]string{"dailyenglish", "questions"})

	// 启动生产者
	go middlewares.RunProducer()

	// 启动消费者
	go middlewares.RunConsumer()

	r := gin.Default()
	r.Use(middlewares.Cors())
	r.Static("static/team_manager", "./static")
	adminrouter.InitAdminRouter(r, db, rdb)
	teamrouter.InitTeamRouter(r, db, rdb, es)
	go func() {
		r1 := gin.Default()
		r1.Use(middlewares.Cors()) //跨域
		userrouter.InitUserRouter(r1, db, rdb, es)
		r1.Run(":8080")
	}()
	log.Println("Server is running at :8081")
	r.Run(":8081")
}
