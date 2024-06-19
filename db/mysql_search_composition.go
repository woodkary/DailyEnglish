package db

import (
	utils "DailyEnglish/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Composition_completion struct {
	TitleID      string `json:"title_id"`
	TeamID       string `json:"team_id"`
	Title        string `json:"title"`
	Word_num     string `json:"word_num"`
	Requirement  string `json:"requirement"`
	Publish_date string `json:"publish_date"`
	Grade        string `json:"grade"`
	Tag          string `json:"tag"`
	Team_Name    string `json:"team_name"`
	Submit_num   int    `json:"submit_num"`
	Member_num   int    `json:"member_num"`
}
type WritingTask struct {
	TitleID      string `json:"title_id"`
	Title        string `json:"title"`
	Manager_name string `json:"manager_name"`
	Word_num     string `json:"word_num"`
	Requirement  string `json:"requirement"`
	Publish_date string `json:"publish_date"`
	Submit_date  string `json:"submit_date"`
	Grade        string `json:"grade"`
	Tag          string `json:"tag"`
	Machine_mark int    `json:"score"`
}

type EssayResult struct {
	TitleID    int    `json:"title_id"`
	Title      string `json:"title"`
	RawEssay   string `json:"raw_essay"`
	Word_cnt   string `json:"word_cnt"`
	Requirment string `json:"requirment"`
	//机器评分与评价
	Machine_mark     int    `json:"machine_mark"`
	Machine_evaluate string `json:"machine_evaluate"`
	//教师评分与评价
	Teacher_mark     int              `json:"teacher_mark"`
	Teacher_evaluate string           `json:"teacher_evaluate"`
	MajorScore       utils.MajorScore `json:"majorScore"`
	//逐句点评
	SentsFeedback []SentsFeedback `json:"sents_feedback"`
}
type SentsFeedback struct {
	ParaId                int                   `json:"para_id"`
	SentId                int                   `json:"sent_id"`
	RawSent               string                `json:"raw_sent"`
	ErrorPosInfos         []utils.ErrorPosInfos `json:"errorPosInfos"`
	SentFeedback          string                `json:"sent_feedback"`
	CorrectedSent         string                `json:"corrected_sent"`
	IsContainGrammarError bool                  `json:"is_contain_grammar_error"`
	IsValidLangSent       bool                  `json:"is_valid_lang_sent"`
}

//根据manager_id查找发布的所有任务以及完成情况

// 查找用户的写作任务，训练任务和已完成任务
func GetUserWritingTask(db *sql.DB, user_id int) ([]WritingTask, []WritingTask, []WritingTask, error) {
	// 定义结果集
	Tasks := []WritingTask{}
	TrainingTasks := []WritingTask{}
	FinishedTasks := []WritingTask{}

	// 查询用户的team_id
	var team_id int
	err := db.QueryRow("SELECT team_id FROM `user-team` WHERE user_id = ?", user_id).Scan(&team_id)
	if err != nil {
		return nil, nil, nil, err
	}

	// 查询团队和系统发布的所有title_ids
	var title_ids []int
	var finish_ids []int

	// 执行合并查询
	rows, err := db.Query(`
        SELECT title_id, team_id 
        FROM composition 
        WHERE team_id = ? OR team_id = 0`, team_id)
	if err != nil {
		return nil, nil, nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var title_id, t_id int
		err := rows.Scan(&title_id, &t_id)
		if err != nil {
			return nil, nil, nil, err
		}
		if t_id == team_id {
			title_ids = append(title_ids, title_id)
		} else {
			finish_ids = append(finish_ids, title_id)
		}
	}

	// 定义一个函数来处理任务
	processTasks := func(ids []int, tag string) ([]WritingTask, []WritingTask, error) {
		var tasks []WritingTask
		var finishedTasks []WritingTask

		if len(ids) == 0 {
			return tasks, finishedTasks, nil
		}

		// 构建查询字符串
		query := `
            SELECT c.title_id, c.composition_title, c.manager_id, c.word_num, c.composition_require, 
                   c.publish_date, c.grade, e.respond_date, e.machine_mark
            FROM composition c
            LEFT JOIN composition_evaluate e ON c.title_id = e.title_id AND e.user_id = ?
            WHERE c.title_id IN (` + strings.Join(strings.Split(strings.Repeat("?", len(ids)), ""), ",") + `)`

		args := make([]interface{}, len(ids)+1)
		args[0] = user_id
		for i, id := range ids {
			args[i+1] = id
		}

		// 执行批量查询
		rows, err := db.Query(query, args...)
		if err != nil {
			return nil, nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var task WritingTask
			var manager_id, grade int
			var respond_date sql.NullString
			var machine_mark sql.NullInt32

			err := rows.Scan(&task.TitleID, &task.Title, &manager_id, &task.Word_num, &task.Requirement,
				&task.Publish_date, &grade, &respond_date, &machine_mark)
			if err != nil {
				return nil, nil, err
			}

			task.Grade = gradeMap[grade]
			task.Tag = tag

			if tag == "任务" {
				err = db.QueryRow("SELECT manager_name FROM manager_info WHERE manager_id = ?", manager_id).Scan(&task.Manager_name)
				if err != nil {
					return nil, nil, err
				}
			} else {
				task.Manager_name = "系统"
			}

			if respond_date.Valid {
				task.Submit_date = respond_date.String
				task.Machine_mark = int(machine_mark.Int32)
				finishedTasks = append(finishedTasks, task)
			} else {
				tasks = append(tasks, task)
			}
		}

		return tasks, finishedTasks, nil
	}

	// 处理团队任务
	teamTasks, finishedTeamTasks, err := processTasks(title_ids, "任务")
	if err != nil {
		return nil, nil, nil, err
	}
	Tasks = append(Tasks, teamTasks...)
	FinishedTasks = append(FinishedTasks, finishedTeamTasks...)

	// 处理训练任务
	trainingTasks, finishedTrainingTasks, err := processTasks(finish_ids, "训练")
	if err != nil {
		return nil, nil, nil, err
	}
	TrainingTasks = append(TrainingTasks, trainingTasks...)
	FinishedTasks = append(FinishedTasks, finishedTrainingTasks...)

	// 按照提交日期对已完成任务进行排序
	sort.Slice(FinishedTasks, func(i, j int) bool {
		iTime, _ := time.Parse("2006-01-02", FinishedTasks[i].Submit_date)
		jTime, _ := time.Parse("2006-01-02", FinishedTasks[j].Submit_date)
		return iTime.After(jTime)
	})

	return Tasks, TrainingTasks, FinishedTasks, nil
}

// 根据用户id和titleid查用户作文结果
func GetEssayResult(db *sql.DB, title_ID int, user_id int) (EssayResult, error) {
	essayResult := EssayResult{}

	// 合并查询获取composition和composition_evaluate的数据
	query := `
        SELECT c.composition_title, c.word_num, c.composition_require, e.evaluate_id, e.machine_mark, e.machine_evaluate, 
               e.teacher_mark, e.teacher_evaluate, e.major_score, e.rawessay
        FROM composition c
        LEFT JOIN composition_evaluate e ON c.title_id = e.title_id
        WHERE c.title_id = ? AND e.user_id = ?`

	var major_score string
	var evaluate_id int

	err := db.QueryRow(query, title_ID, user_id).Scan(&essayResult.Title, &essayResult.Word_cnt, &essayResult.Requirment,
		&evaluate_id, &essayResult.Machine_mark, &essayResult.Machine_evaluate,
		&essayResult.Teacher_mark, &essayResult.Teacher_evaluate, &major_score, &essayResult.RawEssay)
	if err != nil {
		return EssayResult{}, err
	}

	// 将major_score JSON字符串转为MajorScore结构
	err = json.Unmarshal([]byte(major_score), &essayResult.MajorScore)
	if err != nil {
		return essayResult, err
	}

	// 查询evaluate_everysentence表中的所有数据
	rows, err := db.Query(`
        SELECT paraid, sentid, rawsent, errorposinfo, sentfeedback, correctedsent, is_containgrammarerror, is_validlangsent 
        FROM evaluate_everysentence 
        WHERE evaluate_id = ?`, evaluate_id)
	if err != nil {
		return essayResult, err
	}
	defer rows.Close()

	var sentsFeedback []SentsFeedback
	for rows.Next() {
		var sentFeedback SentsFeedback
		var errorPosInfo string

		err := rows.Scan(&sentFeedback.ParaId, &sentFeedback.SentId, &sentFeedback.RawSent, &errorPosInfo, &sentFeedback.SentFeedback,
			&sentFeedback.CorrectedSent, &sentFeedback.IsContainGrammarError, &sentFeedback.IsValidLangSent)
		if err != nil {
			return essayResult, err
		}

		// 将errorPosInfo JSON字符串转为ErrorPosInfos结构
		err = json.Unmarshal([]byte(errorPosInfo), &sentFeedback.ErrorPosInfos)
		if err != nil {
			return essayResult, err
		}

		sentsFeedback = append(sentsFeedback, sentFeedback)
	}
	essayResult.SentsFeedback = sentsFeedback

	return essayResult, nil
}

// 说实话，实在懒得再把这些搬进utils包里了，直接写在这里吧。
func parseMeaningsFromMap(meanings map[string]interface{}) *Meanings {
	meaningMap := Meanings{
		Verb:         toStringSlice(meanings["verb"]),
		Noun:         toStringSlice(meanings["noun"]),
		Pronoun:      toStringSlice(meanings["pronoun"]),
		Adjective:    toStringSlice(meanings["adjective"]),
		Adverb:       toStringSlice(meanings["adverb"]),
		Preposition:  toStringSlice(meanings["preposition"]),
		Conjunction:  toStringSlice(meanings["conjunction"]),
		Interjection: toStringSlice(meanings["interjection"]),
	}
	return &meaningMap
}

// toStringSlice converts an interface to a []string.
func toStringSlice(value interface{}) []string {
	if value == nil {
		return []string{}
	}
	if slice, ok := value.([]interface{}); ok {
		result := make([]string, len(slice))
		for i, v := range slice {
			result[i] = v.(string)
		}
		return result
	}
	return []string{}
}

// 向数据库存入一段机器评分作文的数据
func InsertEssayScore(db *sql.DB, userID int, titleID int, url string, result utils.Response) error {
	//先插入一条记录到composition_score表
	insertQuery, err := db.Prepare("INSERT INTO composition_evaluate(user_id,title_id,composition_url,respond_date,machine_evaluate,machine_mark,rawessay,major_score) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Panic(err)
		return err
	}
	defer insertQuery.Close()
	var grade int
	err = db.QueryRow("SELECT grade FROM composition WHERE title_id = ?", titleID).Scan(&grade)
	if err != nil {
		log.Panic(err)
		return err
	}
	//将分数转换为百分制
	TotalScore := int(result.Result.TotalScore / gradeScoreMap[grade] * 100)
	//将result.Result.EssayFeedback.MajorScore json格式转为字符串
	majorScore, err := json.Marshal(result.Result.MajorScore)
	if err != nil {
		log.Panic(err)
		return err
	}
	_, err = insertQuery.Exec(userID, titleID, url, time.Now().Format("2006-01-02"), result.Result.EssayAdvice, TotalScore, result.Result.RawEssay, string(majorScore))
	if err != nil {
		log.Panic(err)
		return err
	}
	//查询刚插入的记录的id
	var scoreID int
	err = db.QueryRow("SELECT evaluate_id FROM composition_evaluate WHERE user_id = ? AND title_id = ? AND composition_url = ?", userID, titleID, url).Scan(&scoreID)
	if err != nil {
		log.Panic(err)
		return err
	}
	//再更新evaluate_everysentence表
	for _, sentence := range result.Result.EssayFeedback.SentsFeedback {
		if len(sentence.ErrorPosInfos) > 0 {
			insertQuery, err := db.Prepare("INSERT INTO evaluate_everysentence(evaluate_id,paraid,sentid,rawsent,errorposinfo,sentfeedback,correctedsent,is_containgrammarerror,is_validlangsent) VALUES(?,?,?,?,?,?,?,?,?)")
			if err != nil {
				log.Panic(err)
				return err
			}
			defer insertQuery.Close()
			var is_containgrammarerror int
			var is_validlangsent int
			if sentence.IsContainGrammarError {
				is_containgrammarerror = 1
			} else {
				is_containgrammarerror = 0
			}
			if sentence.IsValidLangSent {
				is_validlangsent = 1
			} else {
				is_validlangsent = 0
			}
			errorPosInfo, err := json.Marshal(sentence.ErrorPosInfos)
			if err != nil {
				log.Panic(err)
				return err
			}
			_, err = insertQuery.Exec(scoreID, sentence.ParaId, sentence.SentId, sentence.RawSent, string(errorPosInfo), sentence.SentFeedback, sentence.CorrectedSent, is_containgrammarerror, is_validlangsent)
			if err != nil {
				log.Panic(err)
				return err
			}
		}

	}
	return nil
}
func GetEssayTitle(db *sql.DB, titleId int) (string, int, error) {
	var title string
	var grade int
	fmt.Print("titleId:", titleId)
	err := db.QueryRow("SELECT composition_title,grade FROM composition WHERE title_id = ?", titleId).Scan(&title, &grade)
	if err != nil {
		return "", 0, err
	}
	return title, grade, nil
}

// 根据Team map[int]string查询管理员发布的所有作文
func GetAllComposition(db *sql.DB, Team map[int]string) ([]Composition_completion, error) {
	var composition_completions []Composition_completion
	for team_id, team_name := range Team {
		var composition_completion Composition_completion
		composition_completion.TeamID = strconv.Itoa(team_id)
		composition_completion.Team_Name = team_name
		rows, err := db.Query("SELECT title_id,composition_title,word_num,composition_require,publish_date,grade FROM composition WHERE team_id = ?", team_id)
		if err != nil {
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			var grade, title_ID int
			err := rows.Scan(&title_ID, &composition_completion.Title, &composition_completion.Word_num, &composition_completion.Requirement, &composition_completion.Publish_date, &grade)
			if err != nil {
				return nil, err
			}
			composition_completion.TitleID = strconv.Itoa(title_ID)
			composition_completion.Grade = gradeMap[grade]
			//查询提交人数
			var submit_num int
			err = db.QueryRow("SELECT COUNT(*) FROM composition_evaluate WHERE title_id = ?", composition_completion.TitleID).Scan(&submit_num)
			if err != nil {
				return nil, err
			}
			//查询团队人数
			var member_num int
			err = db.QueryRow("SELECT COUNT(*) FROM `user-team` WHERE team_id = ?", team_id).Scan(&member_num)
			if err != nil {
				return nil, err
			}
			composition_completion.Submit_num = submit_num
			composition_completion.Member_num = member_num
			composition_completions = append(composition_completions, composition_completion)
		}

	}
	return composition_completions, nil
}
