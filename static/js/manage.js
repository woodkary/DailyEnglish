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
            let members=data.members;
            totalPage=Math.ceil(members.length/PAGE_SIZE);
            //在获取数据后再创建分页按钮
            pagination.innerHTML='';
            createPaginationButtons();
            //分页
            const startIndex=(page-1)*PAGE_SIZE;
            const endIndex=startIndex+PAGE_SIZE;
            members=members.slice(startIndex,endIndex);
            members.forEach(member => {
                const row=document.createElement("tr");
                //todo: 获取到后端数据后，显示对应的内容

                //创建按钮1
                const buttonCell=document.createElement("td");
                const button=document.createElement("button");
                button.classList.add("BlueButton");//按钮样式
                const span=document.createElement("span");//按钮文字
                span.innerText="详情";
                span.classList.add("UnderlinedBlue");//按钮文字样式
                button.appendChild(span);
                //创建按钮2
                const button2=document.createElement("button");
                button2.classList.add("RedButton");//按钮样式
                const span2=document.createElement("span");//按钮文字
                span2.innerText="删除";
                span2.classList.add("UnderlinedRed");//按钮文字样式
                button2.appendChild(span2);
                //将按钮1添加到单元格中
                buttonCell.appendChild(button);//将按钮添加到单元格中
                //将按钮2添加到单元格中
                buttonCell.appendChild(button2);//将按钮添加到单元格中
                //将单元格1添加到行中
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
