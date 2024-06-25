//这些数据可以从后端获取，这里假设已经获取到了
let teamAndStudents = {
    "team1": ["student1", "student2", "student3"],
    "team2": ["student4", "student5", "student6"],
    "team3": ["student7", "student8", "student9"]
}
//所有题目的数据
const questions=[
    {name:"选择",max: 5},
    {name:"填空",max: 5},
    {name:"作文",max:5},
    {name:"阅读",max:100},
    {name:"完型",max:200},
    {name:"翻译",max:100}
];
//所有学生不同题型的平均成绩数据
let studentAverageScores=[
    {name: 'student1', value: [3, 2.5, 5, 6, 20, 16]},
    {name: 'student2', value: [2, 1.5, 3, 6, 1, 7]},
    {name: 'student3', value: [1, 3.5, 4, 10, 7, 2]},
    {name: 'student4', value: [3, 4, 3, 2, 2, 10]},
    {name: 'student5', value: [4, 2, 4, 3, 13, 3]},
    {name: 'student6', value: [5, 1.5, 1, 3, 4, 16]},
    {name: 'student7', value: [2, 0.5, 1, 10, 13, 19]},
    {name: 'student8', value: [3, 2.6, 3, 5, 8, 15]},
    {name: 'student9', value: [1, 6.7, 4, 6, 6, 16]}
]
let teamAverageScores=[
    {name: 'team1', value: [5, 4.76, 3, 70, 90, 77]},
    {name: 'team2', value: [4.21, 4.2, 3.9, 91, 78, 65]},
    {name: 'team3', value: [4, 3.87, 2.56, 92, 79, 69]}
]
//最近7次考试的名称
let examNames=["2021年秋季期末考试","2021年春季期末考试","2021年夏季期末考试","2021年秋季期中考试","2021年春季期中考试","2021年夏季期中考试","2021年秋季期末考试"]
//最近7次考试的排名数据
let studentRankChanges=[
    {name: 'student1', data: [120, 132, 101, 134, 90, 230, 210]},
    {name: 'student2', data: [100, 120, 110, 130, 120, 200, 220]},
    {name: 'student3', data: [110, 125, 105, 120, 115, 210, 200]},
    {name: 'student4', data: [100, 120, 110, 130, 120, 200, 220]},
    {name: 'student5', data: [110, 125, 105, 120, 115, 210, 200]},
    {name: 'student6', data: [100, 120, 110, 130, 120, 200, 220]},
    {name: 'student7', data: [110, 125, 105, 120, 115, 210, 200]},
    {name: 'student8', data: [100, 120, 110, 130, 120, 200, 220]},
    {name: 'student9', data: [110, 125, 105, 120, 115, 210, 200]}
]
//将后端获取的团队和学生数据转换成前端需要的格式
function transformTeamData(teams) {
    let teamAndStudents = {};

    teams.forEach(team => {
        let teamKey = team.team_name;
        teamAndStudents[teamKey] = team.members.map(member => member.name);
    });

    return teamAndStudents;
}
function teamAndStudentsInit(){
    //初始化团队和学生数据
    fetch('http://localhost:8081/api/team_manage/exam_situation/teams_and_students_grade',{
        method: 'GET',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    }).then(res => {
        if(res.status==401||res.status=="401"){
            //如果token失效，则跳转到登录页面
            alert("登录信息已过期，请重新登录");
            localStorage.removeItem('token');
            window.location.href="./login&register.html";
            return null;
        }
        return res.json();
    })
   .then(data => {
       if(data.code===200||data.code=="200"){
           //直接获取数据，无需转换格式
           teamAndStudents=data.team_and_students;
           studentAverageScores=data.student_average_scores;
           teamAverageScores=data.team_average_scores;
           examNames=data.exam_names;
           studentRankChanges=data.student_rank_scores;
           loadTeamAndStudents();
       }
   }).catch(err => console.log(err));
}
teamAndStudentsInit();
//loadTeamAndStudents();
//根据学生排名变化，获取折线图的数据
function getStackedLineSeries(studentNames){
    let result=[];
    for(let i=0;i<studentNames.length;i++){
        let studentRankChange=studentRankChanges.find(item=>item.name===studentNames[i]);
        let data=studentRankChange.data;
        let series={
            name: studentNames[i],
            type: 'line',
            stack: 'Total',
            data: data
        };
        result.push(series);
    }
    return result;
}
//根据学生名获取平均分
function getAverageGrade(studentName){
    return studentAverageScores.find(item => item.name === studentName)
}
//根据团队名获取平均分
function getTeamAverageGrade(teamName){
    return teamAverageScores.find(item => item.name === teamName)
}
//加载团队和学生数据
function loadTeamAndStudents() {
    let teamSelect=document.getElementById("teamSelect");
    teamSelect.innerHTML="";
    let flag=false;
    if(teamAndStudents==null||teamAndStudents.length===0){
        //如果没有团队数据，则提示用户
        return;
    }
    for (let team in teamAndStudents) {
        //在团队选择下拉框中添加团队选项
        let newOption=document.createElement("option");
        newOption.value=team;
        newOption.text=team;
        teamSelect.add(newOption);
        if(!flag){
            //默认选择第一个团队
            teamSelect.value=team;
            flag=true;
            loadStudents(team);
            //加载该团队的学生的成绩折线图
            let stuNames=teamAndStudents[team];
            loadStackedLineChart(stuNames);
            //监听团队选择下拉框的change事件，以改变学生选择下拉框的选项
            teamSelect.addEventListener("change", function() {
                let selectedTeam=teamSelect.value;
                loadStudents(selectedTeam);
                //加载该团队的学生的成绩折线图
                let stuNames=teamAndStudents[selectedTeam];
                loadStackedLineChart(stuNames);
            });
        }
    }
}
//根据选择的团队加载学生数据
function loadStudents(selectedTeam){
    let studentSelect = document.getElementById("studentSelect");
    studentSelect.innerHTML = "";
    let flag=false;
    if(teamAndStudents[selectedTeam]==null||teamAndStudents[selectedTeam].length===0){
        //如果没有学生数据，则提示用户
        return;
    }
    teamAndStudents[selectedTeam].forEach(student => {
        let newOption = document.createElement("option");
        newOption.value = student;
        newOption.text = student;
        studentSelect.add(newOption);
        if(!flag){
            //默认选择第一个学生，并加载选择的团队
            studentSelect.value=student;
            flag=true;
            loadRadarChart(student,selectedTeam);
            //监听学生选择下拉框的change事件，以改变学生的成绩雷达图
            studentSelect.addEventListener("change", function() {
                let selectedStudent = studentSelect.value;
                loadRadarChart(selectedStudent,selectedTeam);
            });
        }
    });
}
function loadRadarChart(studentName,teamName){
    let chartDom = document.getElementById('radar-chart');
    let myChart = echarts.init(chartDom);
    let option;
    option = {
        legend: {
            top: 0,
            data: [studentName,teamName]
        },
        grid: {
            containLabel: true
        },
        radar: {
            // shape: 'circle',
            radius: '60%', // 减小雷达图的半径
            indicator: questions
        },
        series: [
            {
                name: '成绩',
                type: 'radar',
                data: [
                    getAverageGrade(studentName),
                    getTeamAverageGrade(teamName)
                ]
            }
        ]
    };
    option && myChart.setOption(option);
}
function loadStackedLineChart(studentNames){
    let chartDom = document.getElementById('stacked-line-chart');
    let myChart = echarts.init(chartDom);
    let option;

    option = {
        tooltip: {
            trigger: 'axis'
        },
        legend: {
            //所有学生姓名
            data: studentNames
        },
        grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
        },
        toolbox: {
            feature: {
                saveAsImage: {}
            }
        },
        xAxis: {
            type: 'category',
            boundaryGap: false,
            //最近7次考试的名称
            data: examNames
        },
        yAxis: {
            type: 'value'
        },
        //折线图数据
        series: getStackedLineSeries(studentNames)
    };
    option && myChart.setOption(option);
}