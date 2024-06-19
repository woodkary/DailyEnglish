let studentUploads=[
    {
        evaluateId: 1,
        studentId: 1,
        studentName: '张三',
        respondDate: '2021-05-10',
        wordCount: 100,
        score: 80
    },
    {
        evaluateId: 2,
        studentId: 3,
        studentName: '李四',
        respondDate: '2021-05-10',
        wordCount: 100,
        score: 80
    }
]
function getChartObject(studentUploads){
    let res=[];
    studentUploads.forEach(item=> {
        res.push({
            提交日期: item.respondDate,
            姓名: item.studentName,
            词数: item.wordCount,
            得分: item.score
        });
    });
    return res;
}
createTable(getChartObject(studentUploads));
function createTable(data) {
    // 清空表格内容
    document.getElementById('tableBody').innerHTML = '';
    // 创建表头
    const tableHead = document.querySelector('thead tr');
    const headers = Object.keys(data[0]); // 获取第一个数据对象的所有属性作为表头
    headers.forEach(header => {
        const th = document.createElement('th');
        th.textContent = header;
        tableHead.appendChild(th);
    });
    //创建操作按钮表头
    const th = document.createElement('th');
    th.textContent = '操作';
    tableHead.appendChild(th);

    // 创建表格行
    data.forEach(item => {
        const row = document.createElement('tr');
        headers.forEach(header => {
            const td = document.createElement('td');
            td.textContent = item[header]; // 设置单元格内容
            row.appendChild(td);
        });
        document.getElementById('tableBody').appendChild(row);
        //创建一个操作按钮单元格
        const td = document.createElement('td');
        const button = document.createElement('button');
        button.textContent = '详情';
        button.addEventListener('click', () => {
            // 跳转到TeacherEssayMark页面
            window.location.href = `TeacherEssayMark.html?studentId=${item.studentId}`;
        });
        td.appendChild(button);
        row.appendChild(td);
    });

}