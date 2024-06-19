let titleId="6";
let title="学生上传作文";
let wordNum=100;
let requirement="sdcfsdvsdvsdvsd";
//从url参数获取titleId字符串
let urlParams = new URLSearchParams(window.location.search);
if (urlParams.has('title_id')&&urlParams.has('title')&&urlParams&&urlParams.has('word_num')&&urlParams.has('requirement')) {
    titleId = urlParams.get('title_id');
    title = urlParams.get('title');
    wordNum = urlParams.get('word_num');
    requirement = urlParams.get('requirement');
    document.getElementById('title').textContent = title;
    document.getElementById('wordNum').textContent = wordNum;
    document.getElementById('requirement').textContent = requirement;
}
console.log(titleId);
requestStudentUploads(titleId);
//这是根据titleId获取学生上传的作文数据，请从composition_evaluate中查找对应数据
let studentUploads=[
    {
        evaluateId: 1,
        studentId: 1,
        studentName: '张三',
        respondDate: '2021-05-10',
        machineScore: 80,
        teacherScore: 90,
    },
    {
        evaluateId: 2,
        studentId: 3,
        studentName: '李四',
        respondDate: '2021-05-10',
        machineScore: 80,
        teacherScore: 90,
    }
]
function requestStudentUploads(titleId) {
    //打印titleId的类型
    console.log(typeof titleId);
    fetch(`http://localhost:8081/api/team_manage/composition_mission/submission_records`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
        body: JSON.stringify({
            title_id: titleId
        })
    }).then(res => {
        if(res.status === 401){
            alert('请先登录！');
            window.location.href = 'login&register.html';
            return;
        }
        if (res.ok) {
            return res.json();
        } else {
            throw new Error('Network response was not ok');
        }
    }).then(data => {
        console.log(data);
        studentUploads = [];
        data.records.forEach(item => {
            studentUploads.push({
                evaluateId: item.evaluate_id,
                studentId: item.student_id,
                studentName: item.student_name,
                respondDate: item.respond_date,
                machineScore: item.machine_score,
                teacherScore: item.teacher_score,
            });
        });
        createTable(getChartObject(studentUploads));
    }).catch(error => {
        console.error('Error:', error);
    });
}
function getChartObject(studentUploads){
    let res=[];
    studentUploads.forEach(item=> {
        res.push({
            提交日期: item.respondDate,
            姓名: item.studentName,
            机器评分: item.machineScore,
            教师评分: item.teacherScore
        });
    });
    return res;
}
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