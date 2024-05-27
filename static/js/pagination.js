/*
 * @Date: 2024-04-17 14:17:35
 */
let BUTTON_NUM = 5; // 每页显示的按钮数
let page = 1; // 当前页数
let totalPage = 100; // 总页数，不知道怎么获取，先写死
const QUESTION_PER_PAGE = 10; // 每页题目数
let pageAndQuestionsMap={};//存放每页题目的字典，key为页数，value为题目数组
let allQuestionIds=new Set();//存放所有确认选择的题目的id集合
const questionTypeDict = {
    1: "单选题",
    2: "多选题",
    3: "判断题",
    4: "填空题",
    5: "简答题"
};
const difficultyDescriptions = {
    1: "容易",
    2: "中等",
    3: "困难",
}
//页面加载时，先获取第一页到第五页道题目，共50道题目
document.addEventListener("DOMContentLoaded",()=>{
    getQuestion(1);
});
//分页按钮
createPaginationButtons();
function createPaginationButtons() {
    const startBtns = document.querySelector(".startBtns");
    let startBtn = document.createElement("button");
    startBtn.classList.add('button');
    let i = document.createElement("i");
    i.classList.add("fa", "fa-angles-left");
    startBtn.appendChild(i);
    startBtn.addEventListener("click", () => {
        page = 1;
        updatePagination(BUTTON_NUM);
        getQuestion(page);
    });
    startBtns.appendChild(startBtn);
    let prevButton = document.createElement("button");
    prevButton.classList.add('button');
    let i2 = document.createElement("i");
    i2.classList.add("fa", "fa-angle-left");
    prevButton.appendChild(i2);
    prevButton.addEventListener("click", () => {
        page--;
        if (page < 1) {
            page = 1;
        }
        updatePagination(BUTTON_NUM);
        getQuestion(page);
    });
    startBtns.appendChild(prevButton);
    updatePagination(BUTTON_NUM);
    const endBtns = document.querySelector(".endBtns");
    const nextButton = document.createElement("button");
    nextButton.classList.add('button');
    let i3 = document.createElement("i");
    i3.classList.add("fa", "fa-angle-right");
    nextButton.appendChild(i3);
    nextButton.addEventListener("click", () => {
        page++;
        if (page > totalPage) {
            page = totalPage;
        }
        updatePagination(BUTTON_NUM);
        getQuestion(page);
    });
    endBtns.appendChild(nextButton);

    const endBtn = document.createElement("button");
    endBtn.classList.add('button');
    let i4 = document.createElement("i");
    i4.classList.add("fa", "fa-angles-right");
    endBtn.appendChild(i4);
    endBtn.addEventListener("click", () => {
        page = totalPage;
        //更新分页按钮
        updatePagination(BUTTON_NUM);
        //根据page请求题目
        getQuestion(page);
    });
    endBtns.appendChild(endBtn);
}
//更新分页按钮和题目列表
function updatePagination(pageSize) {
    const totalPages = totalPage; // 总页数
    let startPage = Math.max(1, page - Math.floor(pageSize / 2));
    let endPage = Math.min(totalPages, startPage + pageSize - 1);

    if (endPage - startPage < pageSize - 1) {
        startPage = Math.max(1, endPage - pageSize + 1);
    }
    const pagination = document.getElementById('pagination');
    pagination.innerHTML = '';

    for (let i = startPage; i <= endPage; i++) {
        const a = document.createElement('a');
        a.textContent = i;
        a.classList.add('link');
        a.addEventListener('click', () => {
            page = i;
            updatePagination(BUTTON_NUM);
            getQuestion(page);
        });

        if (i === page) {
            a.classList.add('active');
        }

        pagination.appendChild(a);
    }
}
//请求体参数为index，以page传入，向后端请求所有题目
function getQuestion(page) {
    //如果当前页数已经请求过题目，则直接从pageAndQuestionsMap中获取
    if(pageAndQuestionsMap[page]) {
        createQuestionTable(pageAndQuestionsMap[page]);
        return;
    }
    fetch('http://localhost:8081/api/team_manage/new_exam/all_questions', {
        method: 'POST',
        body: JSON.stringify({
            index: page,
        }),
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+localStorage.getItem('token')
        }
    }).then(res => res.json()).then(data => {
        if(data.code === 200){
            //获取到题目后，先以十道题目分开
            let questions=data.questions;
            const responseTotalPages=5;//响应总是返回5页
            for(let i=0;i<responseTotalPages;i++){
                let startIndex=i*QUESTION_PER_PAGE;
                let endIndex=startIndex+QUESTION_PER_PAGE;
                let questionSubset=questions.slice(startIndex,endIndex);
                //创建表格
                createQuestionTable(questionSubset);
                //获取当前题目数组的开始页数
                let startPage=i+page;
                //存入pageAndQuestionsMap
                pageAndQuestionsMap[startPage]=questionSubset;
            }
        }
    }).catch(err => {
        console.log(err);
    })
}

/*{
    "code": 200,
    "msg": "成功",
    "questions": [
        {
            "question_id": "example001",
            "question_type": "1",
            "question_difficulty": "3",
            "question_grade": "13",
            "question_content": "Are u OK?",
            "question_choices":{
                "A": "Yes",
                "B": "No",
                "C": "I dont know",
                "D": "Hi"
            },
        ,
        "question_answer": "Yes",
        "full_score": 5
        }
    ]
}*/
function createQuestionTable(questions) {
    let tableBody=document.querySelector("#tableBody");
    //先清空原有内容
    tableBody.innerHTML="";
    for(let i=0;i<questions.length;i++){
        let questionId=questions[i].question_id;
        let tr=document.createElement("tr");
        //先创建选择按钮
        let tdInput=document.createElement("td");
        let input=document.createElement("input");
        input.type="checkbox";
        input.name="checkbox";
        input.value=questionId;
        input.addEventListener("click",()=>{
            if(input.checked){
                //如果选中，则添加到allQuestionIds中
                allQuestionIds.add(questionId);
            }else{
                //如果取消选中，则从allQuestionIds中删除
                allQuestionIds.delete(questionId);
            }
        });
        tdInput.appendChild(input);
        tr.appendChild(tdInput);
        //再创建题目类型展示框
        let tdType=document.createElement("td");
        //从字典中获取题目类型对应的文字
        tdType.innerText=questionTypeDict[questions[i].question_type];
        tr.appendChild(tdType);
        //再创建题目内容展示框
        let tdContent=document.createElement("td");
        tdContent.innerText=questions[i].question_content;
        tr.appendChild(tdContent);
        //再创建题目难度展示框
        let tdDifficulty=document.createElement("td");
        //从字典中获取题目难度对应的文字
        tdDifficulty.innerText=difficultyDescriptions[questions[i].question_difficulty];
        tr.appendChild(tdDifficulty);
    }
}
//提交按钮
let publishBtn=document.querySelector(".publish-button");
publishBtn.addEventListener("click",()=>{
    let examSelectBtn=document.querySelector('.exam-select-btn')
    //选中的考试日期
    let examDate=convertDateToISOFormat(examSelectBtn.children.item(0).textContent);
    if(examDate==='error'){
        alert("您没有选择考试日期！");
        return;
    }
    //获取选择的团队
    let teamSelect=document.querySelector('#team-select');
    let teamName=teamSelect.options[teamSelect.selectedIndex].value;
    if(teamName==='请选择团队'){
        alert("您没有选择团队！");
        return;
    }
    //获取题目id数组
    let questionIds=Array.from(allQuestionIds);
    if(questionIds.length===0){
        alert("您没有选择题目！");
        return;
    }
    let examName=document.querySelector('#exam-name-input').value;
    if(examName.trim()==''){
        alert("考试名称不能为空！");
        return;
    }
    let examStartTime=document.querySelector('#exam-start-time').textContent;
    examStartTime=from12To24(examStartTime);
    let examEndTime=document.querySelector('#exam-end-time').textContent;
    examEndTime=from12To24(examEndTime);
    fetch('http://localhost:8081/api/team_manage/new_exam', {
        method: 'POST',
        body: JSON.stringify({
            team_name: teamName,
            exam_name: examName,
            exam_date: examDate,
            exam_clock: `${examStartTime}~${examEndTime}`,
            question_ids: questionIds,
        }),
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+localStorage.getItem('token')
        }
    }).then(res => res.json()).then(data => {
        if (data.code === 200) {
            alert("发布成功！");
        }
    }).catch(err => {
        console.log(err);
    });
});
function convertDateToISOFormat(dateString) {
    // 正则表达式匹配'YYYY年MM月DD日'格式
    const match = dateString.match(/(\d+)年(\d+)月(\d+)日/);

    // 如果匹配失败，返回错误
    if (!match) {
        return 'error';
    }
    // 假设dateString的格式为'YYYY年MM月DD日'
    const year = dateString.match(/\d+/g)[0];
    const month = dateString.match(/\d+/g)[1].padStart(2, '0');
    const day = dateString.match(/\d+/g)[2].padStart(2, '0');

    // 返回ISO格式日期'YYYY-MM-DD'
    return `${year}-${month}-${day}`;
}
function from12To24(time) {
    let timeAndPeriod = time.split(' ');
    let [hour, minute] = timeAndPeriod[0].split(':');
    let period = timeAndPeriod[1];
    if (period === 'PM' && hour !== '12') {
        hour = parseInt(hour, 10) + 12;
    }
    if (period === 'AM' && hour === '12') {
        hour = '00';
    }

    return `${hour}:${minute}`;
}