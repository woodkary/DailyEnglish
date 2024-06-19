/*type Request struct {
    TeamId      int    `json:"team_id"`      // 发布作文的团队ID，如果是0，则代表选择的是题库自带的作文
    Title       string `json:"title"`        // 作文题目
    MinWordNum  int    `json:"min_word_num"` // 最少字数要求
    MaxWordNum  int    `json:"max_word_num"` // 最多字数要求
    Requirement string `json:"requirement"`  // 作文要求
    Grade       string `json:"grade"`        // 作文等级
}*/
let requestParams={
    team_id: 0,
    title: "",
    min_word_num: 0,
    max_word_num: 0,
    requirement: "",
    grade: "小学"
}
let teamInfo={};
//页面加载，首先从本地缓存获取teamId和name的映射关系
document.addEventListener("DOMContentLoaded", function() {
    init();

});
function modalTeamSelect(){
    let teamSelect=document.getElementById("modalTeamSelect");
    teamSelect.innerHTML="";
    let flag=false;
    let teams=JSON.parse(localStorage.getItem("team_info"));
    teamInfo=teams;
    if(teams==null||teams.length===0){
        //如果没有团队数据，则提示用户
        return;
    }
    for (let team in teams) {
        //在团队选择下拉框中添加团队选项
        let newOption=document.createElement("option");
        newOption.value=team;
        newOption.text=teams[team];
        teamSelect.add(newOption);
        if(!flag){
            //默认选择第一个团队
            teamSelect.value=team;
            flag=true;
            teamSelect.addEventListener("change", function() {
                requestParams.team_id=parseInt(teamSelect.value);
            });
        }
    }
}
function init(){
    //从本地缓存中获取teamId和name的映射关系
    let teamInfo=JSON.parse(localStorage.getItem("team_info"));
    let select=document.getElementById("teamSelect");
    select.innerHTML="";
    let flag=false;
    //渲染select
    for(let teamId in teamInfo){
        let option=document.createElement("option");
        option.value=teamId;
        option.textContent=teamInfo[teamId];
        select.add(option);
        if(!flag){
            requestParams.team_id=parseInt(teamId);
            flag=true;
        }
    }
    //change事件
    select.addEventListener("change", function() {
        requestParams.team_id=parseInt(select.value);
    });
    requestSystemEssays();
    modalTeamSelect();
}
function setTitle(title){
    requestParams.title=title;
}
function setMinWordNum(minWordNum){
    requestParams.min_word_num=parseInt(minWordNum);
}
function setMaxWordNum(maxWordNum){
    requestParams.max_word_num=parseInt(maxWordNum);
}
function pressEnter(event,input){
    if(event.keyCode==13){
        setRequirement(input);
    }
}
function setRequirement(requirement){
    requestParams.requirement=requirement;
}
function setGrade(grade){
    requestParams.grade=grade;
}
function request(){
    console.log(requestParams);
    //发送请求
    fetch("http://localhost:8081/api/team_manage/composition_mission", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            'Authorization': `Bearer ${localStorage.getItem("token")}`
        },
        body: JSON.stringify(requestParams)
    }).then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error("Network response was not ok");
        }
    }).then(data => {
        console.log(data);
        alert("发布成功");
    }).catch(error => {
        console.error("Error:", error);
        alert("发布失败");
    });
}
//这是后端传来的所有系统作文，
//todo publishDate不知道是什么
let systemEssays=[
    {
        titleId: "1",
        title: "小学生作文",
        grade: "小学",
        wordNum:"100~200",
        requirement: "哈哈哈",
        publishDate: "2021-01-01"
    },
    {
        titleId: "2",
        title: "小学生作文2",
        grade: "小学",
        wordNum:"150~250",
        requirement: "fuck you",
        publishDate: "2022-01-01"
    }
]
const essayDivMap= {
    "全部":document.getElementById("全部"),
    "小学":document.getElementById("小学"),
    "初中":document.getElementById("初中"),
    "高中":document.getElementById("高中"),
    "四级":document.getElementById("四级"),
    "六级":document.getElementById("六级"),
    "考研":document.getElementById("考研"),
    "托福":document.getElementById("托福"),
    "雅思":document.getElementById("雅思"),
    "GRE":document.getElementById("GRE")
}
//清空essayDivMap中所有元素的内联样式
for(let key in essayDivMap){
    essayDivMap[key].innerHTML="";
}
function requestSystemEssays(){
    //发送请求
    fetch("http://localhost:8081/api/team_manage/composition_mission/system_compositions", {
        method: "GET",
        headers: {
            'Authorization': `Bearer ${localStorage.getItem("token")}`
        },
    }).then(response => {
        if (response.ok) {
            return response.json();
        } else {
            throw new Error("Network response was not ok");
        }
    }).then(data => {
        console.log(data);
        systemEssays=[];
        data.compositions.forEach(composition => {
            systemEssays.push({
                titleId: composition.title_id,
                title: composition.title,
                grade: composition.grade,
                wordNum: composition.word_num,
                requirement: composition.requirement,
                publishDate: composition.publish_date
            });
        });

        renderSystemEssays(systemEssays,"全部");
        renderSystemEssays(systemEssays,"小学");
        renderSystemEssays(systemEssays,"初中");
        renderSystemEssays(systemEssays,"高中");
        renderSystemEssays(systemEssays,"四级");
        renderSystemEssays(systemEssays,"六级");
        renderSystemEssays(systemEssays,"考研");
        renderSystemEssays(systemEssays,"托福");
        renderSystemEssays(systemEssays,"雅思");
        renderSystemEssays(systemEssays,"GRE");
    }).catch(error => {
        console.error("Error:", error);
    });
}
function openTab(event, tabId) {
    // Hide all tab contents
    var tabContents = document.querySelectorAll('.tab-content');
    tabContents.forEach(function (content) {
        content.style.display = 'none';
    });

    // Remove 'active' class from all tab buttons
    var tabButtons = document.querySelectorAll('.tab-button');
    tabButtons.forEach(function (button) {
        button.classList.remove('active');
    });

    // Show the selected tab content and add 'active' class to the clicked tab button
    document.getElementById(tabId).style.display = 'block';
    event.currentTarget.classList.add('active');
}
function openSubTab(event, subtabId) {
    var subtabContents = document.querySelectorAll('.subtab-content');
    subtabContents.forEach(function (content) {
        content.style.display = 'none';
    });

    // Remove 'active' class from all tab buttons
    var subtabButtons = document.querySelectorAll('.subtab-button');
    subtabButtons.forEach(function (button) {
        button.classList.remove('active');
    });

    // Show the selected tab content and add 'active' class to the clicked tab button
    document.getElementById(subtabId).style.display = 'block';
    event.currentTarget.classList.add('active');
    //把essayDivMap中不是当前tab的作文隐藏掉
    for(let key in essayDivMap) {
        if (key != subtabId) {
            essayDivMap[key].style.display = "none";
        } else {
            essayDivMap[key].style.display = "block";
        }
    }
}
/*    var gradeMap = map[int]string{
        1: "小学",
            2: "初中",
            3: "高中",
            4: "四级",
            5: "六级",
            6: "考研",
            7: "托福",
            8: "雅思",
            9: "GRE",
    }*/

function renderSystemEssays(essays, grade) {
    let container = document.getElementById(grade);
    container.innerHTML = ''; // Clear existing content

    // Filter essays based on grade
    essays.forEach((essay, index) => {
        if (grade == essay.grade || grade == "全部") {
            let card = document.createElement('div');
            card.className = 'card';

            let line1 = document.createElement('div');
            line1.className = 'line1';
            line1.innerHTML = `<span style="color:#456de7;margin-right: 12px;">[题目]</span><span>${essay.title}</span>`;

            let line2 = document.createElement('div');
            line2.className = 'line2';
            line2.innerHTML = `<div class="wordcnt">字数：<span>${essay.wordNum}</span></div><div class="req">要求:<span>${essay.requirement}</span></div>`;

            let line3 = document.createElement('div');
            line3.className = 'line3';
            line3.innerHTML = `<span class="time">上传时间：<span>${essay.publishDate}</span></span><div class="submit-btn2"><button type="submit" id="button-${grade}-${index}">发布</button></div>`;

            card.appendChild(line1);
            card.appendChild(line2);
            card.appendChild(line3);

            container.appendChild(card);

            // Add event listener to the button
            let button = document.getElementById(`button-${grade}-${index}`);
            button.addEventListener('click', () => {
                requestParams.title=essay.title;
                let wordNumArr=essay.wordNum.split("~");
                requestParams.min_word_num=parseInt(wordNumArr[0]);
                requestParams.max_word_num=parseInt(wordNumArr[1]);
                requestParams.requirement=essay.requirement;
                requestParams.grade=essay.grade;
                request();
            });
        }
    });
}
