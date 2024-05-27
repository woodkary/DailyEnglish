/*
 * @Date: 2024-04-17 14:17:35
 */
let BUTTON_NUM = 5; // 每页显示的按钮数
let page = 1; // 当前页数
let totalPage = 10; // 总页数
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
function getQuestion(index) {
    fetch('/api/team_manage/new_exam/all_questions', {
        method: 'POST',
        body: JSON.stringify({
            index: index,
        }),
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+localStorage.getItem('token')
        }
    }).then(res => res.json()).then(data => {
        if(data.code === 200){
            //获取到题目后，创建表格
            createQuestionTableAndGetQuestionIds(data.questions);
        }
    }).catch(err => {
        console.log(err);
    })
}
function createQuestionTableAndGetQuestionIds(questions) {
    let tableBody=document.querySelector("#tableBody");
    //先清空原有内容
    tableBody.innerHTML="";
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
    //将问题id数组返回
    return questionIdArray;
}