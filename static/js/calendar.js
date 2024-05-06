let year=2024,
    month=5;
let dates=generateDates();
renderCalendar(dates);
setMonth();
function renderCalendar(dates) {
    const calendarDays = document.querySelector('.days'); // 获取日历的 days 容器
    calendarDays.innerHTML = ''; // 清空容器，准备重新渲染
    let week;

    dates.forEach((date,index) => {
        if(index%7===0){
            week = document.createElement('div'); // 创建一个 div 元素作为一周
            week.className = 'week'; // 设置类名
            calendarDays.appendChild(week); // 将周添加到日历的 days 容器中
        }
        const dayDiv = document.createElement('div'); // 创建一个 div 元素
        dayDiv.className = 'day'; // 设置类名
        // 判断是否为本月的日期
        if(date.getMonth() !== month-1){
            dayDiv.classList.add('notThisMonth'); // 设置类名为 notThisMonth
        }

        const dayNumber = document.createElement('span'); // 创建一个 span 元素用于显示日期数字
        dayNumber.className = 'day-number'; // 设置类名
        dayNumber.textContent = date.getDate(); // 设置日期数字

        dayDiv.appendChild(dayNumber); // 将日期数字添加到 dayDiv 中
        week.appendChild(dayDiv); // 将 dayDiv 添加到日历的 days 容器中
    });
}
function getMonthString(month){
    switch(month){
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
function setMonth(){
    let monthText=document.querySelector('#month');
    monthText.textContent=getMonthString(month);
}
function generateDates(){
    let dates; // 存储当前月份的日期
    const firstDay = new Date(year, month - 1, 1); // 获取当前月份的第一天
    const firstDayOfWeek = firstDay.getDay(); // 获取当前月份的第一天是星期几
    const totalDays = new Date(year, month, 0).getDate(); // 获取当前月份的总天数

    // 初始化日期数组
    dates = [];
    // 填充上个月的日期
    for (let i = firstDayOfWeek-1; i>=0; i--) {
        dates.push(new Date(firstDay - (i+1) * 24 * 60 * 60 * 1000));
    }
    // 填充本月的日期
    for (let i = 1; i <= totalDays; i++) {
        dates.push(new Date(year, month - 1, i));
    }

    // 填充下个月的日期
    for (let i = 0; dates.length<42; i++) {
        dates.push(new Date(year,month,i+1));
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
    }else{
        let today = new Date();
        year = today.getFullYear();
        month = today.getMonth() + 1;
    }
    dates = generateDates();
    renderCalendar(dates);
    setMonth();
}