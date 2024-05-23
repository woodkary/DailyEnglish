const canvas = document.getElementById('pie-chart');
const ctx = canvas.getContext('2d');
const centerX = canvas.width / 2;
const centerY = canvas.height / 2;
const radius = Math.min(centerX, centerY);
const innerRadius = radius * 0.6; // 内部半径
// 后端提供的数据：未完成和已完成的数量
let unfinishedTasks = 20;
let finishedTasks = 80;
let totalTasks = 0;//总任务数
let unfinishedAngle = 0;//未完成任务的角度
let finishedAngle = 0;//已完成任务的角度
//加载时执行，获取未完成和已完成的数量
window.onload = function () {
    //获取token
    let token = sessionStorage.getItem("token");
    fetch("http://localhost:8080/api/team_manager/index", {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+token//在请求头设置 token
        }
    }).then(response => {
        return response.json();
    }).then(data => {
        console.log(data);
        unfinishedTasks = data.uncompleted;
        finishedTasks = data.completed;
        // 计算角度
        totalTasks = unfinishedTasks + finishedTasks;//总任务数
        unfinishedAngle = (Math.PI * 2 * unfinishedTasks) / totalTasks;//未完成任务的角度
        finishedAngle = (Math.PI * 2 * finishedTasks) / totalTasks;//已完成任务的角度
        // 在获取到新的unfinishedTasks和finishedTasks值后调用drawPieChart方法重新绘制饼图
        drawPieChart();
    })
}
// 显示信息的HTML元素
const tooltip = document.getElementById('tooltip');


// 绘制饼图
function drawPieChart() {
    // 绘制未完成部分
    ctx.beginPath();
    ctx.moveTo(centerX + innerRadius * Math.cos(-Math.PI / 2), centerY + innerRadius * Math.sin(-Math.PI / 2));
    ctx.arc(centerX, centerY, innerRadius, -Math.PI / 2, -Math.PI / 2 + unfinishedAngle);
    ctx.lineTo(centerX + radius * Math.cos(-Math.PI / 2 + unfinishedAngle), centerY + radius * Math.sin(-Math.PI / 2 + unfinishedAngle));
    ctx.arc(centerX, centerY, radius, -Math.PI / 2 + unfinishedAngle, -Math.PI / 2, true);
    ctx.closePath();
    ctx.lineWidth = 3; // 设置边框宽度
    ctx.strokeStyle = '#FFFFFF'; // 设置边框颜色
    ctx.stroke(); // 绘制边框
    ctx.fillStyle = '#FFE76F'; // 未完成部分的颜色，灰蓝色
    ctx.fill();


    // 绘制已完成部分
    ctx.beginPath();
    ctx.moveTo(centerX + innerRadius * Math.cos(-Math.PI / 2 + unfinishedAngle), centerY + innerRadius * Math.sin(-Math.PI / 2 + unfinishedAngle));
    ctx.arc(centerX, centerY, innerRadius, -Math.PI / 2 + unfinishedAngle, -Math.PI / 2 + unfinishedAngle + finishedAngle);
    ctx.lineTo(centerX + radius * Math.cos(-Math.PI / 2 + unfinishedAngle + finishedAngle), centerY + radius * Math.sin(-Math.PI / 2 + unfinishedAngle + finishedAngle));
    ctx.arc(centerX, centerY, radius, -Math.PI / 2 + unfinishedAngle + finishedAngle, -Math.PI / 2 + unfinishedAngle, true);
    ctx.closePath();
    ctx.lineWidth = 3; // 设置边框宽度
    ctx.strokeStyle = '#FFFFFF'; // 设置边框颜色
    ctx.stroke(); // 绘制边框
    ctx.fillStyle = '#002EA6'; // 已完成部分的颜色，蓝色
    ctx.fill();

    // 添加事件监听器来检测鼠标悬停
    canvas.addEventListener('mousemove', handleMouseMove);
    canvas.addEventListener('mouseout', handleMouseOut);
}

// 绘制图例
const legend = document.getElementById('legend');
legend.innerHTML = `
        <div class="legend-item">
        <span class="unfinished-color"></span>未完成任务
        </div>
        <div class="legend-item">
            <span class="finished-color"></span>已完成任务
        </div>
        `;

// 处理鼠标悬停事件
function handleMouseMove(event) {
    const mouseX = event.clientX - canvas.getBoundingClientRect().left;
    const mouseY = event.clientY - canvas.getBoundingClientRect().top;
    const angle = Math.atan2(mouseY - centerY, mouseX - centerX) + Math.PI / 2;
    let degree = (angle * 180) / Math.PI;
    if (degree < 0) {
        degree = 360 + degree; // Convert negative degree to positive
    }
    let percentage, people;
    if (degree >= 0 && degree < unfinishedAngle * (180 / Math.PI)) {
        // 悬停在未完成部分
        percentage = (unfinishedTasks / totalTasks * 100).toFixed(2);
        people = unfinishedTasks;
    } else if (degree >= unfinishedAngle * (180 / Math.PI) && degree < (unfinishedAngle + finishedAngle) * (180 / Math.PI)) {
        // 悬停在已完成部分
        percentage = (finishedTasks / totalTasks * 100).toFixed(2);
        people = finishedTasks;
    }
    showTooltip(`${people}人 (${percentage}%)`, event.pageX, event.pageY);
}

// 处理鼠标移出事件
function handleMouseOut() {
    hideTooltip();
}

// 显示提示框
function showTooltip(content, x, y) {
    tooltip.style.display = 'block';
    tooltip.style.left = x + 'px';
    tooltip.style.top = y + 'px';
    tooltip.innerText = content;
}

// 隐藏提示框
function hideTooltip() {
    tooltip.style.display = 'none';
}