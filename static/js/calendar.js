let todaysDate = new Date();
let year = todaysDate.getFullYear(),
    month = todaysDate.getMonth() + 1;
let dates = generateDates();
let date_Exam_Map={};
renderCalendar(dates);
setMonth();
getTodayExamData();
function renderCalendar(dates) {
    const calendarDays = document.querySelector('.days'); // 获取日历的 days 容器
    calendarDays.innerHTML = ''; // 清空容器，准备重新渲染
    let week;

    dates.forEach((date, index) => {
        if (index % 7 === 0) {
            week = document.createElement('div'); // 创建一个 div 元素作为一周
            week.className = 'week'; // 设置类名
            calendarDays.appendChild(week); // 将周添加到日历的 days 容器中
        }
        const dayDiv = document.createElement('div'); // 创建一个 div 元素
        dayDiv.addEventListener('click', () => {
            let eventDiv=document.querySelector('.event');
            eventDiv.innerHTML='';
            // 清空之前的考试信息
            let dateH1=document.createElement('h1');
            // 设置日期标题
            dateH1.textContent=date.toLocaleDateString();
            eventDiv.appendChild(dateH1);
            // 点击日期时，获取日期对应的考试信息并渲染到页面
            let exams = date_Exam_Map[date];
            if(exams){
                for(let i=0;i<exams.length;i++) {
                    // 创建一个 card 元素作为考试信息的容器
                    let cardDiv = document.createElement('div');
                    cardDiv.className = 'card';
                    eventDiv.appendChild(cardDiv);
                    //创建一个标记
                    let image=document.createElement('img');
                    image.src='./image/done.svg';
                    image.id='finish';
                    cardDiv.appendChild(image);
                    // 创建一个 p 元素作为考试团队名称
                    let teamP = document.createElement('p');
                    teamP.textContent = exams[i].team_name;
                    teamP.className = 'team';
                    cardDiv.appendChild(teamP);
                    // 创建一个 p 元素作为考试时间
                    let timeP = document.createElement('p');
                    timeP.textContent = exams[i].time;
                    timeP.className = 'time';
                    cardDiv.appendChild(timeP);
                    //创建跳转按钮
                    let jumpBtn=document.createElement('button');
                    jumpBtn.className='toDetail';
                    let jumpSvg=document.createElement('img');
                    jumpSvg.src='./image/jump.svg';
                    jumpSvg.alt='Jump';
                    jumpBtn.appendChild(jumpSvg);
                    cardDiv.appendChild(jumpBtn);
                    jumpBtn.addEventListener('click',()=>{
                        window.location.href='./test-statistics.html?date='+date.toLocaleDateString()+"&team_name="+exams[i].team_name;
                    });
                }
                //创建加考试按钮
                let addBtn=document.createElement('button');
                addBtn.className='addEvent';
                addBtn.textContent='add event';
                eventDiv.appendChild(addBtn);
            }else{
                fetchExamData(date);
            }
        });
        dayDiv.className = 'day'; // 设置类名
        // 判断是否为本月的日期
        if (date.getMonth() !== month - 1) {
            dayDiv.classList.add('notThisMonth'); // 设置类名为 notThisMonth
        }

        const dayNumber = document.createElement('span'); // 创建一个 span 元素用于显示日期数字
        dayNumber.className = 'day-number'; // 设置类名
        dayNumber.textContent = date.getDate(); // 设置日期数字

        dayDiv.appendChild(dayNumber); // 将日期数字添加到 dayDiv 中
        week.appendChild(dayDiv); // 将 dayDiv 添加到日历的 days 容器中
    });
}
function getMonthString(month) {
    switch (month) {
        case 1:
            return 'January';
        case 2:
            return 'February';
        case 3:
            return 'March';
        case 4:
            return 'April';
        case 5:
            return 'May';
        case 6:
            return 'June';
        case 7:
            return 'July';
        case 8:
            return 'August';
        case 9:
            return 'September';
        case 10:
            return 'October';
        case 11:
            return 'November';
        case 12:
            return 'December';
    }
}
function setMonth() {
    let monthText = document.querySelector('#month');
    monthText.textContent = getMonthString(month);
}
function generateDates() {
    let dates; // 存储当前月份的日期
    const firstDay = new Date(year, month - 1, 1); // 获取当前月份的第一天
    const firstDayOfWeek = firstDay.getDay(); // 获取当前月份的第一天是星期几
    const totalDays = new Date(year, month, 0).getDate(); // 获取当前月份的总天数

    // 初始化日期数组
    dates = [];
    // 填充上个月的日期
    for (let i = firstDayOfWeek - 1; i >= 0; i--) {
        dates.push(new Date(firstDay - (i + 1) * 24 * 60 * 60 * 1000));
    }
    // 填充本月的日期
    for (let i = 1; i <= totalDays; i++) {
        dates.push(new Date(year, month - 1, i));
    }

    // 填充下个月的日期
    for (let i = 0; dates.length < 35; i++) {
        dates.push(new Date(year, month, i + 1));
    }
    return dates;
}
function moveDate(direction) {
    if (direction === 'prev') {
        if (month === 1) {
            year--;
            month = 12;
        } else {
            month--;
        }
    } else if (direction === 'next') {
        if (month === 12) {
            year++;
            month = 1;
        } else {
            month++;
        }
        //跳到今天所在的年份和月份
    } else {
        let today = new Date();
        year = today.getFullYear();
        month = today.getMonth() + 1;
    }
    dates = generateDates();
    renderCalendar(dates);
    setMonth();
}
function examFactory(name,team_name,time,full_score,average_score,pass_rate){
    return {
        name: name,
        team_name: team_name,
        time: time,
        full_score: full_score,
        average_score: average_score,
        pass_rate: pass_rate
    };
}
function fetchExamData(date){
    fetch('http://localhost:8080/api/team_manage/exam_situation/exams_of_date', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
        body: JSON.stringify({
            date: date
        })
    }).then(response  =>{
        if(response.ok){
            return response.json();
        }else{
            renderDefaultExamData(date);
        }
    }).then(data => {
        let exams = data.exams;
        let exam_list = [];
        for(let i=0;i<exams.length;i++){
            let exam = examFactory(exams[i].name,exams[i].team_name,exams[i].time,exams[i].full_score,exams[i].average_score,exams[i].pass_rate);
            exam_list.push(exam);
        }
        date_Exam_Map[date] = exam_list;
        renderExamData(date);
    }).catch(error => {
        renderDefaultExamData(date);
    });
}
function renderExamData(date){
    let eventDiv=document.querySelector('.event');
    eventDiv.innerHTML='';
    // 清空之前的考试信息
    let dateH1=document.createElement('h1');
    // 设置日期标题
    dateH1.textContent=date.toLocaleDateString();
    eventDiv.appendChild(dateH1);
    // 点击日期时，获取日期对应的考试信息并渲染到页面
    let exams = date_Exam_Map[date];
    for(let i=0;i<exams.length;i++) {
        // 创建一个 card 元素作为考试信息的容器
        let cardDiv = document.createElement('div');
        cardDiv.className = 'card';
        eventDiv.appendChild(cardDiv);
        // 创建一个标记
        let image = document.createElement('img');
        image.src = './image/done.svg';
        image.id = 'finish';
        cardDiv.appendChild(image);
        // 创建一个 p 元素作为考试团队名称
        let teamP = document.createElement('p');
        teamP.className = 'team';
        teamP.textContent = exams[i].team_name;
        cardDiv.appendChild(teamP);
        // 创建一个 p 元素作为考试时间
        let timeP = document.createElement('p');
        timeP.className = 'time';
        timeP.textContent = exams[i].time;
        cardDiv.appendChild(timeP);
        //创建跳转按钮
        let jumpBtn=document.createElement('button');
        jumpBtn.className='toDetail';
        let jumpSvg=document.createElement('img');
        jumpSvg.src='./image/jump.svg';
        jumpSvg.alt='Jump';
        jumpBtn.appendChild(jumpSvg);
        cardDiv.appendChild(jumpBtn);
        jumpBtn.addEventListener('click',()=>{
            window.location.href='./test-statistics.html?date='+date.toLocaleDateString()+"&team_name="+exams[i].team_name;
        });
    }
}
function getTodayExamData(){
    let date=todaysDate;
    let eventDiv=document.querySelector('.event');
    eventDiv.innerHTML='';
    // 清空之前的考试信息
    let dateH1=document.createElement('h1');
    // 设置日期标题
    dateH1.textContent=date.toLocaleDateString();
    eventDiv.appendChild(dateH1);
    // 点击日期时，获取日期对应的考试信息并渲染到页面
    let exams = date_Exam_Map[date];
    if(exams) {
        for (let i = 0; i < exams.length; i++) {
            // 创建一个 card 元素作为考试信息的容器
            let cardDiv = document.createElement('div');
            cardDiv.className = 'card';
            eventDiv.appendChild(cardDiv);
            // 创建一个标记
            let image = document.createElement('img');
            image.src = './image/done.svg';
            image.id = 'finish';
            cardDiv.appendChild(image);
            // 创建一个 p 元素作为考试团队名称
            let teamP = document.createElement('p');
            teamP.className = 'team';
            teamP.textContent = exams[i].team_name;
            cardDiv.appendChild(teamP);
            // 创建一个 p 元素作为考试时间
            let timeP = document.createElement('p');
            timeP.className = 'time';
            timeP.textContent = exams[i].time;
            cardDiv.appendChild(timeP);
            //创建跳转按钮
            let jumpBtn = document.createElement('button');
            jumpBtn.className = 'toDetail';
            let jumpSvg = document.createElement('img');
            jumpSvg.src = './image/jump.svg';
            jumpSvg.alt = 'Jump';
            jumpBtn.appendChild(jumpSvg);
            cardDiv.appendChild(jumpBtn);
            jumpBtn.addEventListener('click', () => {
                window.location.href = './test-statistics.html?date=' + date.toLocaleDateString()+"&team_name="+exams[i].team_name;
            });
        }
    }else{
        fetchExamData(date);
    }
}
function renderDefaultExamData(date){
    let eventDiv=document.querySelector('.event');
    eventDiv.innerHTML='';
    // 清空之前的考试信息
    let dateH1=document.createElement('h1');
    // 设置日期标题
    dateH1.textContent=date.toLocaleDateString();
    eventDiv.appendChild(dateH1);
    // 点击日期时，获取日期对应的考试信息并渲染到页面
    let exams = [
        examFactory('考试1','团队1','10:00-12:00',100,80,90),
        examFactory('考试2','团队2','14:00-16:00',100,85,95),
        examFactory('考试3','团队3','16:00-18:00',100,90,95)
    ];
    for(let i=0;i<exams.length;i++) {
        // 创建一个 card 元素作为考试信息的容器
        let cardDiv = document.createElement('div');
        cardDiv.className = 'card';
        eventDiv.appendChild(cardDiv);
        // 创建一个标记
        let image = document.createElement('img');
        image.src = './image/done.svg';
        image.id = 'finish';
        cardDiv.appendChild(image);
        // 创建一个 p 元素作为考试团队名称
        let teamP = document.createElement('p');
        teamP.textContent = exams[i].team_name;
        teamP.className = 'team';
        cardDiv.appendChild(teamP);
        // 创建一个 p 元素作为考试时间
        let timeP = document.createElement('p');
        timeP.textContent = exams[i].time;
        timeP.className = 'time';
        cardDiv.appendChild(timeP);
        //创建跳转按钮
        let jumpBtn = document.createElement('button');
        jumpBtn.className = 'toDetail';
        let jumpSvg = document.createElement('img');
        jumpSvg.src = './image/jump.svg';
        jumpSvg.alt = 'Jump';
        jumpBtn.appendChild(jumpSvg);
        cardDiv.appendChild(jumpBtn);
        jumpBtn.addEventListener('click', () => {
            window.location.href = './test-statistics.html?date=' + date.toLocaleDateString()+"&team_name="+exams[i].team_name;
        });
    }
    
}