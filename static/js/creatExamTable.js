/*
 * @Date: 2024-04-09 14:03:06
 */
function createExamTable(data) {
    // 创建表格元素
    const table = document.createElement('table');
    table.className = 'borderless-table';

    // 创建表头
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    data.headers.forEach(headerText => {
        const header = document.createElement('th');
        header.textContent = headerText;
        headerRow.appendChild(header);
    });
    thead.appendChild(headerRow);
    table.appendChild(thead);

    // 创建表体
    const tbody = document.createElement('tbody');
    data.rows.forEach(rowData => {
        const row = document.createElement('tr');
        rowData.forEach(cellData => {
            const cell = document.createElement('td');
            cell.textContent = cellData;
            row.appendChild(cell);
        });
        tbody.appendChild(row);
    });
    table.appendChild(tbody);

    // 将表格添加到容器中
    const container = document.getElementById('exam-table');
    container.appendChild(table);
}
//获取最新考试名称并渲染表格
function getLatestExamNameAndRenderTable() {
    console.log(localStorage.getItem('token'));
    fetch('http://47.113.117.103:8080/api/team_manage/exam_situation/data',{
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+localStorage.getItem('token')
        }
    })
        .then(response => response.json())
        .then(data => {
            const latestExam = data.exams[0].name;
            let latestExamDate=document.getElementById('latest-exam-date');
            latestExamDate.textContent=data.exams[0].time;
            fetch('http://47.113.117.103:8080/api/team_manage/exam_situation/exam_detail', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + localStorage.getItem('token')
                },
                body: JSON.stringify({
                    exam_name: latestExam
                })
            })
                    .then(response => response.json())
                    .then(data => {
                        const examResult = data.response.result;
                        // render table using examResult
                        let headers = ['成员', '考试排名', '考试分数'];
                        let rows = [];
                        for (let i = 0; i < examResult.length; i++) {
                            const team = examResult[i];
                            const row = [team.user_id, i + 1, team.user_score];
                            rows.push(row);
                        }
                        let scoreTable = {headers: headers, rows: rows};
                        createExamTable(scoreTable);
                    }).catch(error => {
                        console.log(error);
                        // 示例数据
                        const score_data = {
                            headers: ['成员', '考试排名', '考试分数'],
                            rows: [
                                ['otto', '第1名', '100'], ['otto', '第2名', '99'],
                                ['otto', '第3名', '98'], ['otto', '第4名', '97'], ['otto', '第5名', '96'], ['otto', '第6名', '95']
                                // 更多数据行
                            ]
                        };
                        // 调用函数创建表格
                        createExamTable(score_data);
                    })
            })
        .catch(error => {
                console.log(error);
                // 示例数据
                const score_data = {
                    headers: ['成员', '考试排名', '考试分数'],
                    rows: [
                        ['otto', '第1名', '100'], ['otto', '第2名', '99'],
                        ['otto', '第3名', '98'], ['otto', '第4名', '97'], ['otto', '第5名', '96'], ['otto', '第6名', '95']
                        // 更多数据行
                    ]
                };
                // 调用函数创建表格
                createExamTable(score_data);
    });
}


