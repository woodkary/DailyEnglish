//这些数据可以从后端获取，这里假设已经获取到了
let teamAndStudents = {
    "team1": ["student1", "student2", "student3"],
    "team2": ["student4", "student5", "student6"],
    "team3": ["student7", "student8", "student9"]
}
const questions=[
    {name:"完型",max: 20},
    {name:"阅读",max: 40},
    {name:"翻译",max:10},
    {name:"小作文",max:20},
    {name:"大作文",max:20}
];
let studentAverageGrades=[
    {name: 'student1', value: [12, 2, 5, 20, 16]}, 
    {name: 'student2', value: [20, 18, 3, 1, 7]},
    {name: 'student3', value: [12, 32, 4, 7, 2]},
    {name: 'student4', value: [8, 26, 3, 2, 10]},
    {name: 'student5', value: [3, 22, 4, 13, 3]},
    {name: 'student6', value: [15, 13, 8, 4, 16]},
    {name: 'student7', value: [7, 4, 8, 13, 19]},
    {name: 'student8', value: [14, 26, 7, 8, 15]},
    {name: 'student9', value: [16, 10, 6, 6, 16]}
]
//最近7次考试的名称
let examNames=["2021年秋季期末考试","2021年春季期末考试","2021年夏季期末考试","2021年秋季期中考试","2021年春季期中考试","2021年夏季期中考试","2021年秋季期末考试"]
//最近7次考试的成绩数据
let studentRankChangeData=[
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


loadTeamAndStudents();
//根据学生排名变化，获取折线图的数据
function getStackedLineSeries(studentNames){
    let result=[];
    for(let i=0;i<studentNames.length;i++){
        let studentRankChange=studentRankChangeData.find(item=>item.name===studentNames[i]);
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
    let studentAverageGrade=studentAverageGrades.find(item=>item.name===studentName);
    return studentAverageGrade.value;
}
//加载团队和学生数据
function loadTeamAndStudents() {
    let teamSelect=document.getElementById("teamSelect");
    teamSelect.innerHTML="";
    let flag=false;
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
        }
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
//根据选择的团队加载学生数据
function loadStudents(selectedTeam){
    let studentSelect = document.getElementById("studentSelect");
    studentSelect.innerHTML = "";
    let flag=false;
    for (let student of teamAndStudents[selectedTeam]) {
        let newOption = document.createElement("option");
        newOption.value = student;
        newOption.text = student;
        studentSelect.add(newOption);
        if(!flag){
            //默认选择第一个学生
            studentSelect.value=student;
            flag=true;
            loadRadarChart(student);
        }
        //监听学生选择下拉框的change事件，以改变学生的成绩雷达图
        studentSelect.addEventListener("change", function() {
            let selectedStudent = studentSelect.value;
            loadRadarChart(selectedStudent);
        });
    }
}
function loadRadarChart(studentName){
    let chartDom = document.getElementById('radar-chart');
    let myChart = echarts.init(chartDom);
    let option;
    option = {
        legend: {
            data: [studentName]
        },
        radar: {
            // shape: 'circle',
            indicator: questions
        },
        series: [
            {
                name: '成绩',
                type: 'radar',
                data: [
                    getAverageGrade(studentName)
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