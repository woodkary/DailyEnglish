package test

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	teamrouter "DailyEnglish/router/team_router" // 替换为你的项目路径

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type response struct {
	Code      string   `json:"code"`      // 响应代码
	Msg       string   `json:"msg"`       // 响应消息
	Exam_date []string `json:"exam_date"` // 有考试的日期
}

func TestExamSituationCalendar(t *testing.T) {
	// 创建一个包含 year 和 month 的 map
	data := map[string]string{
		"year":  "2024",
		"month": "05",
	}

	// 将 map 转换为 JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Error encoding JSON: %s", err)
	}

	// 创建一个 gin.Engine 实例
	r := gin.Default()

	// 初始化你的路由
	teamrouter.InitTeamRouter(r, &sql.DB{}, &redis.Client{}, &elasticsearch.Client{})

	// 创建一个模拟的 HTTP 服务器
	server := httptest.NewServer(r)
	defer server.Close()

	// 创建一个请求
	req, err := http.NewRequest("POST", server.URL+"/team/exam_situation_calendar", strings.NewReader(string(jsonData)))
	if err != nil {
		t.Fatalf("Error creating request: %s", err)
	}

	req.Header.Set("Authorization", "Bearer mock_token")
	req.Header.Set("Content-Type", "application/json")

	// 发送请求并获取响应
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("Error making request: %s", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %s", err)
	}

	// 解析响应体
	var res response
	if err := json.Unmarshal(body, &r); err != nil {
		t.Fatalf("Error unmarshalling response: %s", err)
	}

	// 打印 response
	t.Logf("Response: %+v\n", r)

	// 断言 response
	if res.Code != "200" {
		t.Errorf("Expected Code to be '200', but got '%s'", res.Code)
	}

	if res.Msg != "Success" {
		t.Errorf("Expected Msg to be 'Success', but got '%s'", res.Msg)
	}

	if len(res.Exam_date) != 2 {
		t.Errorf("Expected length of Exam_date to be 2, but got %d", len(res.Exam_date))
	}
}
