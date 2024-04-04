let page = 1;
const PAGE_SIZE = 6;
const BUTTON_NUM=5;
let totalPage=0;
const pagination = document.getElementById("pagination");
window.onload = function () {
    page=new URLSearchParams(window.location.search).get("page")||1;
    const tableBody = document.getElementById('table-body');
    tableBody.innerHTML = '';
    allTestStatistics();
}
function allTestStatistics() {
    let token = localStorage.getItem('token');
    let url = "http://localhost:8080/api/team_manage/exam_situation/data";
    fetch(url, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'token': token
        }
    })
       .then(response => response.json())
       .then(data => {
            console.log(data);
            const tableBody=document.getElementById("table-body");
            let exams=data.exams;
            totalPage=Math.ceil(exams.length/PAGE_SIZE);
            //在获取数据后再创建分页按钮
            pagination.innerHTML='';
            createPaginationButtons();
           //分页
            const startIndex=(page-1)*PAGE_SIZE;
            const endIndex=startIndex+PAGE_SIZE;
            exams=exams.slice(startIndex,endIndex);
            exams.forEach(exam => {
                const row=document.createElement("tr");
                //获取考试名称
                const name=document.createElement("td");
                name.innerText=exam.name;
                //获取考试时间
                const time=document.createElement("td");
                time.innerText=exam.time;
                //获取考试地点
                const full_score=document.createElement("td");
                full_score.innerText=exam.full_score;
                //获取考试平均分
                const average_score=document.createElement("td");
                average_score.innerText=exam.average_score;
                //获取考试通过率
                const pass_rate=document.createElement("td");
                pass_rate.innerText=exam.pass_rate;

                //创建按钮
                const buttonCell=document.createElement("td");
                const button=document.createElement("button");
                button.classList.add("BlueButton");//按钮样式
                const span=document.createElement("span");//按钮文字
                span.innerText="查看详情";
                span.classList.add("UnderlinedBlue");//按钮文字样式
                button.appendChild(span);
                //按钮点击事件
                button.addEventListener("click",()=>{
                    window.location.href="/test-detail.html?name="+exam.name;
                });
                buttonCell.appendChild(button);//将按钮添加到单元格中
                //将单元格添加到行中
                row.appendChild(name);
                row.appendChild(time);
                row.appendChild(full_score);
                row.appendChild(average_score);
                row.appendChild(pass_rate);
                row.appendChild(buttonCell);
                //将行添加到表格中
                tableBody.appendChild(row);
            })
        }).catch(error => {
        console.log(error);
    });
}
//分页按钮
function createPaginationButtons() {
    const pagination = document.getElementById("pagination");

    const prevButton = document.createElement("button");
    prevButton.innerText = '<';
    prevButton.addEventListener("click",()=>{
        page=1;
        const tableBody = document.getElementById('table-body');
        tableBody.innerHTML = '';
        allTestStatistics();
    });
    //TODO: 同步机制；前一页按钮需要在页数按钮之前加入。
    pagination.appendChild(prevButton);
    updatePagination(BUTTON_NUM);
    const nextButton = document.createElement("button");
    nextButton.innerText = '>';
    nextButton.addEventListener("click",()=>{
        page=totalPage;
        const tableBody = document.getElementById('table-body');
        tableBody.innerHTML = '';
        allTestStatistics();
    });
    pagination.appendChild(nextButton);
}

function updatePagination(pageSize) {
    const totalPages = totalPage; // 总页数
    let startPage = Math.max(1, page - Math.floor(pageSize / 2));
    let endPage = Math.min(totalPages, startPage + pageSize - 1);

    if (endPage - startPage < pageSize - 1) {
        startPage = Math.max(1, endPage - pageSize + 1);
    }
    const pagination = document.getElementById('pagination');

    for (let i = startPage; i <= endPage; i++) {
        const button = document.createElement('button');
        button.textContent = i;
        button.addEventListener('click', () => {
            page = i;
            const tableBody = document.getElementById('table-body');
            tableBody.innerHTML = '';
            allTestStatistics();
        });

        if (i === page) {
            button.classList.add('active');
        }

        pagination.appendChild(button);
    }
}
