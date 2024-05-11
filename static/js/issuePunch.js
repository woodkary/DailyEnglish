const modal=document.querySelector('.modal');
const modelOverlay=document.querySelector('.modal-overlay');
const calendarContainer=document.querySelector('.calendar-container');
const examSelectBtn=document.querySelector('.exam-select-btn');
let isDisplay=false;
document.addEventListener('keydown',(event)=>{
    if(event.keyCode===77) {//if the key pressed is M
        if(modal.style.display==='none'){
            modal.style.display='flex';//show
            modelOverlay.style.display='flex';//show
            isDisplay=true;
        }else{
            modal.style.display='none';//hide
            modelOverlay.style.display='none';//hide
            isDisplay=false;
        }
    }
});
document.addEventListener('click', (event) => {
    if (isDisplay&&event.target === calendarContainer) {
        modal.style.display = 'none';
        modelOverlay.style.display = 'none';
        isDisplay=false;
    }
});
examSelectBtn.addEventListener('click',()=>{
    modal.style.display='flex';//show
    modelOverlay.style.display='flex';//show
    isDisplay=true;
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
        modal.style.display='none';
        modelOverlay.style.display='none';
        isDisplay=false;
    }
});
const getAllTeamInfo=()=>{
    let teamInfoJson=locaStorage.getItem("team_info");
    if(teamInfoJson){
        return JSON.parse(teamInfoJson);
    }else {
        return null;
    }
}
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