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
        title: "小学生作文",
        wordNum:"100~200",
        requirement: "哈哈哈",
        publishDate: "2021-01-01"
    },
    {
        title: "小学生作文2",
        wordNum:"150~250",
        requirement: "fuck you",
        publishDate: "2022-01-01"
    }
]
renderSystemEssays(systemEssays);
function renderSystemEssays(essays) {
    let container = document.getElementById('subtab1');
    container.innerHTML = '';  // Clear existing content

    essays.forEach(essay => {
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
        line3.innerHTML = `<span class="time">上传时间：<span>${essay.publishDate}</span></span><div class="submit-btn2"><button type="submit">发布</button></div>`;

        card.appendChild(line1);
        card.appendChild(line2);
        card.appendChild(line3);

        container.appendChild(card);
    });
}
