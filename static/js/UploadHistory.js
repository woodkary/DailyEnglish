document.addEventListener('DOMContentLoaded', function () {
    // 获取点击的图标元素
    var toggleFilter = document.querySelector('.toggle-filter');

    // 获取 i 元素
    var filterIcon = toggleFilter.querySelector('i');

    // 获取需要隐藏和显示的 filter 容器
    var filterContainer = document.querySelector('.filter');

    // 添加点击事件
    toggleFilter.addEventListener('click', function () {
        // 切换 filter 容器的显示和隐藏
        if (filterContainer.style.display === 'none') {
            filterContainer.style.display = 'flex';
            filterContainer.style.maxHeight = '1000px'; // 使用足够大的值来确保内容可以完全展开
            filterIcon.classList.remove('uil-angle-down');
            filterIcon.classList.add('uil-angle-up');
        } else {
            filterContainer.style.display = 'none';
            filterContainer.style.maxHeight = '0'; // 隐藏内容
            filterIcon.classList.remove('uil-angle-up');
            filterIcon.classList.add('uil-angle-down');
        }
    });
    requestCompositions();
});
function setProgress(progress,Submit_num, Member_num, titleID) {
    const progressBar = document.querySelector(`#progress-text-${titleID}`).parentElement.querySelector('.progress-bar');
    const progressText = document.getElementById(`progress-text-${titleID}`);

    const radius = 90;
    const circumference = 2 * Math.PI * radius;
    const offset = circumference - (progress / 100 * circumference);

    progressBar.style.strokeDashoffset = offset;
    progressText.innerHTML = `提交人数<br>${Submit_num}/${Member_num}`;
}


let compositionCompletions = [
    {
        TitleID: "1234567890",
        TeamID: "1234567890",
        Title: "Composition 1",
        Word_num: "1000",
        Requirement: "100",
        Publish_date: "2021-01-01",
        Grade: "A",
        Tag: "Tag1, Tag2",
        Team_Name: "Team 1",
        Submit_num: 10,
        Member_num: 10
    },
    {
        TitleID: "123456789",
        TeamID: "1234567890",
        Title: "Composition 2",
        Word_num: "1000",
        Requirement: "100",
        Publish_date: "2021-01-01",
        Grade: "A",
        Tag: "Tag1, Tag2",
        Team_Name: "Team 1",
        Submit_num: 20,
        Member_num: 40
    },
]
function requestCompositions() {
    compositionCompletions = [];
    fetch('http://localhost:8081/api/team_manage/composition_mission/history',{
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    }).then(response => {
        if(response.status===401){
            alert("请先登录！");
            window.location.href = "/login&register.html";
            return;
        }
        if (response.ok) {
            return response.json();
        } else {
            throw new Error("Network response was not ok");
        }
    }).then(data => {
        console.log(data);
        data.compositions.forEach(comp => {
            compositionCompletions.push({
                TitleID: comp.title_id,
                TeamID: comp.team_id,
                Title: comp.title,
                Word_num: comp.word_num,
                Requirement: comp.requirement,
                Publish_date: comp.publish_date,
                Grade: comp.grade,
                Tag: comp.tag,
                Team_Name: comp.team_name,
                Submit_num: comp.submit_num,
                Member_num: comp.member_num
            });
        });
        renderCompositions(compositionCompletions);
    }).catch(error => {
        console.error('Error:', error);
    });
}
function renderCompositions(compositions) {
    const container = document.getElementById('history-container');
    container.innerHTML = '<span class="history-title">发布历史</span>'; // 清空容器内容

    compositions.forEach(comp => {
        const progress = (comp.Submit_num / comp.Member_num) * 100;

        const itemHTML = `
                <div class="history-item">
                    <div class="row1">
                        <div class="line1">
                            <span style="color:#456de7;margin-right: 12px;">[题目]</span>
                            <span>${comp.Title}</span>
                        </div>
                        <div class="line2">
                            <div class="wordcnt">字数：<span>${comp.Word_num}</span></div>
                            <div class="req">要求:<span>${comp.Requirement}</span></div>
                        </div>
                        <div class="line3">
                            <span class="time">发布时间：<span>${comp.Publish_date}</span></span>
                        </div>
                    </div>
                    <div class="row2">
                        <div class="progress-container">
                            <svg class="progress-circle" width="200" height="200" viewBox="0 0 200 200">
                                <circle class="progress-bg" cx="100" cy="100" r="90"></circle>
                                <circle class="progress-bar" cx="100" cy="100" r="90"></circle>
                            </svg>
                            <div class="progress-text" id="progress-text-${comp.TitleID}">0%</div>
                        </div>
                        <button class="submit-btn2" data-title="${comp.Title}" data-grade="${comp.Grade}" data-tags="${comp.Tag}">成绩详情</button>
                    </div>
                </div>
            `;

        const itemElement = document.createElement('div');
        itemElement.innerHTML = itemHTML;
        container.appendChild(itemElement);

        setProgress(progress,comp.Submit_num, comp.Member_num, comp.TitleID);

        // 添加点击事件到按钮
        const btn = itemElement.querySelector('.submit-btn2');
        btn.addEventListener('click', function() {
            //传输题目、字数、要求
            window.location.href = `./StudentUploads.html?title_id=${comp.TitleID}&title=${comp.Title}&word_num=${comp.Word_num}&requirement=${comp.Requirement}`;
        });
    });
}
/*
type Composition_completion struct {
    TitleID      string `json:"title_id"`
    TeamID       string `json:"team_id"`
    Title        string `json:"title"`
    Word_num     string `json:"word_num"`
    Requirement  string `json:"requirement"`
    Publish_date string `json:"publish_date"`
    Grade        string `json:"grade"`
    Tag          string `json:"tag"`
    Team_Name    string `json:"team_name"`
    Submit_num   int    `json:"submit_num"`
    Member_num   int    `json:"member_num"`
}*/
