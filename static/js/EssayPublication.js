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
    fetch("http://localhost:8081/api/team_manage/composition", {
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