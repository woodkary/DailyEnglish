const calendarModal=document.querySelector('#calendar-modal');
const clockModal=document.querySelector('#clock-modal');
const modelOverlay=document.querySelector('.modal-overlay');
const calendarContainer=document.querySelector('.calendar-container');
const clockContainer = document.querySelector(".clock-container");
const examSelectBtn=document.querySelector('.exam-select-btn');
const timeSelectBtns=document.querySelectorAll('.time-select-btn');
let isDisplay=false;
/*document.addEventListener('keydown',(event)=>{
    if(event.keyCode===77) {//if the key pressed is M
        if(calendarModal.style.display==='none'){
            calendarModal.style.display='flex';//show
            modelOverlay.style.display='flex';//show
            isDisplay=true;
        }else{
            calendarModal.style.display='none';//hide
            modelOverlay.style.display='none';//hide
            isDisplay=false;
        }
    }else if(event.keyCode===84) {//if the key pressed is T
        if(clockModal.style.display==='none') {
            clockModal.style.display = 'flex';//show
            modelOverlay.style.display = 'flex';//show
            isDisplay = true;
        }else{
            clockModal.style.display = 'none';//hide
            modelOverlay.style.display = 'none';//hide
            isDisplay = false;
        }
    }
});*/
document.addEventListener('click', (event) => {
    if (isDisplay&&event.target === calendarContainer) {
        calendarModal.style.display = 'none';
        modelOverlay.style.display = 'none';
        isDisplay=false;
    }
    if(isDisplay&&event.target === clockContainer){
        clockModal.style.display = 'none';
        modelOverlay.style.display = 'none';
        isDisplay=false;
    }
});
examSelectBtn.addEventListener('click',()=>{
    calendarModal.style.display='flex';//show
    modelOverlay.style.display='flex';//show
    isDisplay=true;
});
timeSelectBtns.forEach(btn=>{
    btn.addEventListener('click',()=> {
        clockModal.style.display = 'flex';//show
        modelOverlay.style.display = 'flex';//show
        isDisplay = true;
    });
});

const daysTag = document.querySelector(".days"),
    currentDate = document.querySelector(".current-date"),
    prevNextIcon = document.querySelectorAll(".icons span");

// getting new date, current year and month
let date = new Date(),
    currYear = date.getFullYear(),
    currMonth = date.getMonth();

// storing full name of all months in array
const months = ["January", "February", "March", "April", "May", "June", "July",
    "August", "September", "October", "November", "December"];

const renderCalendar = () => {
    let firstDayofMonth = new Date(currYear, currMonth, 1).getDay(), // getting first day of month
        lastDateofMonth = new Date(currYear, currMonth + 1, 0).getDate(), // getting last date of month
        lastDayofMonth = new Date(currYear, currMonth, lastDateofMonth).getDay(), // getting last day of month
        lastDateofLastMonth = new Date(currYear, currMonth, 0).getDate(); // getting last date of previous month
    let liTag = "";

    for (let i = firstDayofMonth; i > 0; i--) { // creating li of previous month last days
        liTag += `<li class="inactive">${lastDateofLastMonth - i + 1}</li>`;
    }

    for (let i = 1; i <= lastDateofMonth; i++) { // creating li of all days of current month
        // adding active class to li if the current day, month, and year matched
        let isToday = i === date.getDate() && currMonth === new Date().getMonth()
        && currYear === new Date().getFullYear() ? "active" : "";
        liTag += `<li class="${isToday}">${i}</li>`;
    }

    for (let i = lastDayofMonth; i < 6; i++) { // creating li of next month first days
        liTag += `<li class="inactive">${i - lastDayofMonth + 1}</li>`
    }
    currentDate.innerText = `${months[currMonth]} ${currYear}`; // passing current mon and yr as currentDate text
    daysTag.innerHTML = liTag;
}
// 为每个天添加点击事件监听器
// 使用事件委托，在daysTag上添加一个点击事件监听器
daysTag.addEventListener('click', (event) => {
    // 检查点击的目标是否是li元素
    if (event.target.tagName === 'LI' && !event.target.classList.contains('inactive')) {
        // 获取点击的日期
        const clickedDate = parseInt(event.target.innerText);
        let newDate = new Date(currYear, currMonth, clickedDate); // creating new date object with clicked date
        let today=new Date();
        // 检查是否选择了过去的时间
        if(newDate.getTime()<today.getTime()){
            alert("不能选择过去的时间！");
            return;
        }
        // 执行点击事件的处理逻辑
        console.log(`Clicked date: ${toDateString(newDate)}`);
        // 在这里添加您的点击处理逻辑
        examSelectBtn.children.item(0).textContent=toDateString(newDate);
        // 关闭弹窗
        calendarModal.style.display='none';
        modelOverlay.style.display='none';
        isDisplay=false;
    }
});
//获取本地存储的团队信息
const getAllTeamInfo=()=>{
    let teamInfoJson=localStorage.getItem("team_info");
    if(teamInfoJson){
        return JSON.parse(teamInfoJson);
    }else {
        return null;
    }
}
const teamSelect=document.querySelector("#team-select");
teamSelect.addEventListener("click",()=>{
    /*alert("wwefse");*/
    let teamInfo=getAllTeamInfo();
    let teamNameArray;
    if(teamInfo) {
        //从本地获取
        teamNameArray = Object.values(teamInfo);
    }else{
        //默认选项
        teamNameArray = ["团队1","团队2","团队3","团队4","团队5","团队6","团队7","团队8","团队9","团队10"];
    }
    // 清空原有选项
    teamSelect.innerHTML="<option disabled selected>请选择发布的团队</option>";
    // 重新渲染选项
    for (let i = 0; i < teamNameArray.length; i++) {
        let option = document.createElement("option");
        option.value = teamNameArray[i];
        option.text = teamNameArray[i];
        teamSelect.add(option);
    }
});
const toDateString=(date)=>{
    return date.getFullYear() + "年" + (date.getMonth() + 1) + "月" + date.getDate()+"日";
}
renderCalendar();

prevNextIcon.forEach(icon => { // getting prev and next icons
    icon.addEventListener("click", () => { // adding click event on both icons
        // if clicked icon is previous icon then decrement current month by 1 else increment it by 1
        currMonth = icon.id === "prev" ? currMonth - 1 : currMonth + 1;

        if(currMonth < 0 || currMonth > 11) { // if current month is less than 0 or greater than 11
            // creating a new date of current year & month and pass it as date value
            date = new Date(currYear, currMonth, new Date().getDate());
            currYear = date.getFullYear(); // updating current year with new date year
            currMonth = date.getMonth(); // updating current month with new date month
        } else {
            date = new Date(); // pass the current date as date value
        }
        renderCalendar(); // calling renderCalendar function
    });
});

const currentTime = clockContainer.querySelector("h1"),
    content = clockContainer.querySelector(".content"),
    selectMenu = clockContainer.querySelectorAll("select"),
    setAlarmBtn = clockContainer.querySelector("button");
let alarmTime, isAlarmSet,
    ringtone = new Audio("./files/ringtone.mp3");
for (let i = 12; i > 0; i--) {
    i = i < 10 ? `0${i}` : i;
    let option = `<option value="${i}">${i}</option>`;
    selectMenu[0].firstElementChild.insertAdjacentHTML("afterend", option);
}
for (let i = 59; i >= 0; i--) {
    i = i < 10 ? `0${i}` : i;
    let option = `<option value="${i}">${i}</option>`;
    selectMenu[1].firstElementChild.insertAdjacentHTML("afterend", option);
}
for (let i = 2; i > 0; i--) {
    let ampm = i == 1 ? "AM" : "PM";
    let option = `<option value="${ampm}">${ampm}</option>`;
    selectMenu[2].firstElementChild.insertAdjacentHTML("afterend", option);
}
setInterval(() => {
    let date = new Date(),
        h = date.getHours(),
        m = date.getMinutes(),
        s = date.getSeconds(),
        ampm = "AM";
    if(h >= 12) {
        h = h - 12;
        ampm = "PM";
    }
    h = h == 0 ? h = 12 : h;
    h = h < 10 ? "0" + h : h;
    m = m < 10 ? "0" + m : m;
    s = s < 10 ? "0" + s : s;
    currentTime.innerText = `${h}:${m}:${s} ${ampm}`;
    if (alarmTime === `${h}:${m} ${ampm}`) {
        ringtone.play();
        ringtone.loop = true;
    }
});
function setAlarm() {
    if (isAlarmSet) {
        alarmTime = "";
        ringtone.pause();
        content.classList.remove("disable");
        setAlarmBtn.innerText = "设置时间";
        return isAlarmSet = false;
    }
    let time = `${selectMenu[0].value}:${selectMenu[1].value} ${selectMenu[2].value}`;
    if (time.includes("Hour") || time.includes("Minute") || time.includes("AM/PM")) {
        return alert("请选择正确的时间！");
    }
    alarmTime = time;
    isAlarmSet = true;
    content.classList.add("disable");
    setAlarmBtn.innerText = "清除时间";
}
setAlarmBtn.addEventListener("click", setAlarm);

/*"questions": [
    {
        "question_id": "example001",
        "question_type": "1",
        "question_difficulty": "3",
        "question_grade": "13",
        "question_content": "Are u OK?",
        "question_choices": [
            "Yes",
            "No",
            "I dont know",
            "貴様のナメクジ野郎"
        ],
        "question_answer": "Yes",
        "full_score": 5
    }
]*/
getQuestionTable=(questions)=>{
    let tableBody=document.querySelector("#tableBody");
    let questionIdArray=[];//创建表格，并返回问题id的数组
    for(let i=0;i<questions.length;i++){
        let questionId=questions[i].question_id;
        questionIdArray.push(questionId);
        let tr=document.createElement("tr");
        //先创建选择按钮
        let tdInput=document.createElement("td");
        let input=document.createElement("input");
        input.type="checkbox";
        input.name="checkbox";
        input.value=questionId;
        tdInput.appendChild(input);
        tr.appendChild(tdInput);

    }
}