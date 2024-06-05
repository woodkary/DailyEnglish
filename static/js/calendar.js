let todaysDate = new Date();
let year = todaysDate.getFullYear(),
    month = todaysDate.getMonth() + 1;
let dates;
getExamDates();
//这个是主要函数
function getExamDates(){
    let exam_dates;
    //如果本地存储中没有考试日期，则从后端获取考试日期
    if(!exam_dates) {
        exam_dates = new Set();
        fetch('http://localhost:8081/api/team_manage/exam_situation/calendar', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            },
            body: JSON.stringify({
                year: year+"",
                month: month+""
            })
        }).then(response => {
            if (response.ok) {
                return response.json();
            } else {
                throw new Error('Network response was not ok');
            }
        }).then(data => {
            if (data.code == 200||data.code=="200") {
                let exam_date = data.exam_date;
                for (let i = 0; i < exam_date.length; i++) {
                    exam_dates.add(exam_date[i]);
                }
                dates = generateDates();
                renderCalendar(dates,exam_dates);
                getTodayExamData();
            } else {
                console.log(data.msg);
                console.log("错误码为：" + data.code)
            }
        }).catch(error => {
            console.error('Error:', error);
        });
    }else{
        //本地存储中有考试日期，则直接渲染日历并获取当天考试数据
        exam_dates=JSON.parse(exam_dates);
        exam_dates=new Set(exam_dates);
        dates = generateDates();
        renderCalendar(dates,exam_dates);
        getTodayExamData();
    }
}

function fromDateToStr(date) {
    let year = date.getFullYear();
    let month = date.getMonth() + 1;
    let zero=month<10?'0':'';
    let day = date.getDate();
    return year + "-" + +zero + month + "-" + day;
}


function renderCalendar(dates,exam_dates) {
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
        let hasExam;
        if((hasExam = exam_dates.has(fromDateToStr(date)))) {
            dayDiv.addEventListener('click', () => {
                let eventDiv = document.querySelector('.event');
                eventDiv.innerHTML = '';
                // 清空之前的考试信息
                let dateH1 = document.createElement('h1');
                // 设置日期标题
                dateH1.textContent = date.toLocaleDateString();
                eventDiv.appendChild(dateH1);
                // 点击日期时，获取日期对应的考试信息并渲染到页面
                /*let exams = date_Exam_Map[date];*/
                //看本地存储中是否有考试信息
                let exams=sessionStorage.getItem(date.toLocaleDateString());
                exams=JSON.parse(exams);
                if (exams) {
                    for (let i = 0; i < exams.length; i++) {
                        // 创建一个 card 元素作为考试信息的容器
                        let cardDiv = document.createElement('div');
                        cardDiv.className = 'card';
                        eventDiv.appendChild(cardDiv);
                        //创建一个标记
                        let image = document.createElement('img');
                        image.src = './image/done.svg';
                        image.id = 'finish';
                        cardDiv.appendChild(image);
                        // 创建一个 p 元素作为考试团队名称
                        let teamP = document.createElement('p');
                        teamP.textContent = exams[i].team_name+"\t"+exams[i].exam_name;
                        teamP.className = 'team';
                        cardDiv.appendChild(teamP);
                        /*// 创建一个 p 元素作为考试时间
                        let timeP = document.createElement('p');
                        timeP.textContent = exams[i].time;
                        timeP.className = 'time';
                        cardDiv.appendChild(timeP);*/
                        //创建跳转按钮
                        let jumpBtn = document.createElement('button');
                        jumpBtn.className = 'toDetail';
                        let jumpSvg = document.createElement('img');
                        jumpSvg.src = './image/jump.svg';
                        jumpSvg.alt = 'Jump';
                        jumpBtn.appendChild(jumpSvg);
                        cardDiv.appendChild(jumpBtn);
                        jumpBtn.addEventListener('click', () => {
                            window.location.href = './test-statistics.html?date=' + date.toLocaleDateString() + "&team_id=" + exams[i].team_id+"&exam_id="+exams[i].exam_id;
                        });
                    }
                    //创建加考试按钮
                    let addBtn = document.createElement('button');
                    addBtn.className = 'addEvent';
                    addBtn.textContent = 'add event';
                    eventDiv.appendChild(addBtn);
                } else {
                    fetchExamData(date);
                }
            });
        }else{
            dayDiv.addEventListener('click', () => {
                renderDefaultExamData(date);
            });
        }
        dayDiv.className = 'day'; // 设置类名
        dayDiv.classList.add(hasExam?'has_exam':'no_exam'); // 设置类名为 date
        let today=new Date();
        if (date.getFullYear() === today.getFullYear() && date.getMonth() + 1 === today.getMonth() + 1 && date.getDate() === today.getDate()) {
            dayDiv.classList.add('today_exam'); // 设置类名为 today
        }
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
        let dateToCheck=new Date(firstDay - (i + 1) * 24 * 60 * 60 * 1000)
        dates.push(new Date(Date.UTC(dateToCheck.getFullYear(), dateToCheck.getMonth(), dateToCheck.getDate())));
    }
    // 填充本月的日期
    for (let i = 1; i <= totalDays; i++) {
        let dateToCheck=new Date(year, month - 1, i);
        dates.push(new Date(Date.UTC(dateToCheck.getFullYear(), dateToCheck.getMonth(), dateToCheck.getDate())));
    }

    // 填充下个月的日期
    for (let i = 0; dates.length < 42; i++) {
        let dateToCheck=new Date(year, month, i + 1);
        dates.push(new Date(Date.UTC(dateToCheck.getFullYear(), dateToCheck.getMonth(), dateToCheck.getDate())));
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
    //重新获取考试日期并渲染日历
    getExamDates();
    setMonth();
}
function examFactory(team_id,exam_id,exam_name,team_name){
    return {
        team_id:team_id,
        exam_id:exam_id,
        exam_name:exam_name,
        team_name:team_name
    };
}
function fetchExamData(date){
    fetch('http://localhost:8081/api/team_manage/exam_situation/exam_date', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
        body: JSON.stringify({
            date: fromDateToStr(date)
        })
    }).then(response  =>{
        if(response.ok){
            return response.json();
        }else{
            renderDefaultExamData(date);
        }
    }).then(data => {
        //获取这个日期的考试信息
        let exams = data.exams;
        let exam_list = [];
        for(let i=0;i<exams.length;i++){
            let exam = examFactory(exams[i].team_id,exams[i].exam_id,exams[i].exam_name,exams[i].team_name);
            exam_list.push(exam);
        }
        /*date_Exam_Map[date] = exam_list;*/
        //存储考试信息到本地存储
        sessionStorage.setItem(date.toLocaleDateString(),JSON.stringify(exam_list));
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
    /*let exams = date_Exam_Map[date];*/
    //看本地存储中是否有该日期对应的考试信息
    let exams=sessionStorage.getItem(date.toLocaleDateString());
    exams=JSON.parse(exams);
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
        teamP.textContent = exams[i].team_name+"\t"+exams[i].exam_name;
        cardDiv.appendChild(teamP);
/*        // 创建一个 p 元素作为考试时间
        let timeP = document.createElement('p');
        timeP.className = 'time';
        timeP.textContent = exams[i].time;
        cardDiv.appendChild(timeP);*/
        //创建跳转按钮
        let jumpBtn=document.createElement('button');
        jumpBtn.className='toDetail';
        let jumpSvg=document.createElement('img');
        jumpSvg.src='./image/jump.svg';
        jumpSvg.alt='Jump';
        jumpBtn.appendChild(jumpSvg);
        cardDiv.appendChild(jumpBtn);
        jumpBtn.addEventListener('click',()=>{
            window.location.href='./test-statistics.html?date='+date.toLocaleDateString()+"&team_id="+exams[i].team_id+"&exam_id="+exams[i].exam_id;
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
    /*let exams = date_Exam_Map[date];*/
    //看本地存储中是否有该日期对应的考试信息
    let exams=sessionStorage.getItem(date.toLocaleDateString());
    exams=JSON.parse(exams);
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
            teamP.textContent = exams[i].team_name+"\t"+exams[i].exam_name;
            cardDiv.appendChild(teamP);
            /*// 创建一个 p 元素作为考试时间
            let timeP = document.createElement('p');
            timeP.className = 'time';
            timeP.textContent = exams[i].time;
            cardDiv.appendChild(timeP);*/
            //创建跳转按钮
            let jumpBtn = document.createElement('button');
            jumpBtn.className = 'toDetail';
            let jumpSvg = document.createElement('img');
            jumpSvg.src = './image/jump.svg';
            jumpSvg.alt = 'Jump';
            jumpBtn.appendChild(jumpSvg);
            cardDiv.appendChild(jumpBtn);
            jumpBtn.addEventListener('click', () => {
                window.location.href = './test-statistics.html?date=' + date.toLocaleDateString()+"&team_id="+exams[i].team_id+"&exam_id="+exams[i].exam_id;
            });
        }
    }else{
        fetchExamData(date);
    }
}
function renderDefaultExamData(date){
/*    let eventDiv=document.querySelector('.event');
    eventDiv.innerHTML='';
    // 清空之前的考试信息
    let dateH1=document.createElement('h1');
    // 设置日期标题
    dateH1.textContent=date.toLocaleDateString();
    eventDiv.appendChild(dateH1);
    // 点击日期时，获取日期对应的考试信息并渲染到页面
    let exams = [
        examFactory('1','1','qwfqwfqwf','saasfasf'),
        examFactory('2','2','qwfqfwb','dfhdfhdfndfn'),
        examFactory('2','3','sdbsdb','dffhbdfndfn')
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
        teamP.textContent = exams[i].team_name+"\t"+exams[i].exam_name;
        teamP.className = 'team';
        cardDiv.appendChild(teamP);
        /!*!// 创建一个 p 元素作为考试时间
        let timeP = document.createElement('p');
        timeP.textContent = exams[i].time;
        timeP.className = 'time';
        cardDiv.appendChild(timeP);*!/
        //创建跳转按钮
        let jumpBtn = document.createElement('button');
        jumpBtn.className = 'toDetail';
        let jumpSvg = document.createElement('img');
        jumpSvg.src = './image/jump.svg';
        jumpSvg.alt = 'Jump';
        jumpBtn.appendChild(jumpSvg);
        cardDiv.appendChild(jumpBtn);
        jumpBtn.addEventListener('click', () => {
            window.location.href = './test-statistics.html?date=' + date.toLocaleDateString()+"&team_id="+exams[i].team_id+"&exam_id="+exams[i].exam_id;
        });
    }
    */
    let eventDiv = document.querySelector('.event');
    eventDiv.innerHTML = '';
    // 清空之前的考试信息
    let dateH1 = document.createElement('h1');
    // 设置日期标题
    dateH1.textContent = "今天没有考试";
    eventDiv.appendChild(dateH1);
}