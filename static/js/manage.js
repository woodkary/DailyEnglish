let page = 1;
const PAGE_SIZE = 6;//每页显示的数量
const BUTTON_NUM=5;//分页按钮数量
let totalPage=0;
const pagination = document.getElementById("pagination");
window.onload = function () {
    let team=document.getElementById("team-name");
    let localTeam=localStorage.getItem("team");
    team.textContent=localTeam==null?"春田花花幼儿园":localTeam;
    allTestStatistics();
}
function allTestStatistics() {
    let token = localStorage.getItem('token');
    let url = "http://47.113.117.103:8080/api/team_manage/member_manage/data";
    fetch(url, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+token
        }
    })
        .then(response => response.json())
        .then(data => {
            console.log(data);
            let members=data.members;
            totalPage=Math.ceil(members.length/PAGE_SIZE);
            //在获取数据后再创建分页按钮
            pagination.innerHTML='';
            createPaginationButtons();
            //分页
            const startIndex=(page-1)*PAGE_SIZE;
            const endIndex=startIndex+PAGE_SIZE;
            const backendData =[];
            //显示总人数
            let memberNumber=document.getElementById("member-number");
            memberNumber.textContent=members.length;
            //显示管理员数量
            let managerNumber=document.getElementById("manager-number");
            let managerCount=members.filter(member=>member.right=="团队管理员").length;
            managerNumber.textContent=managerCount;
            //将后端数据格式化
            members.forEach(member => {
                //获取到后端数据后，显示对应的内容
                let memberData={};
                memberData.nickname=member.name;
                memberData.teamPermission=member.right;
                memberData.joinDate=member.time;
                memberData.totalCheckIns=member.punch_day;
                backendData.push(memberData);
            });
            generateTable(backendData.slice(startIndex, endIndex));
        }).catch(error => {
        console.log(error);
    });
}
function generateTable(data) {
    const container = document.getElementById('table-container');
    container.innerHTML = ''; // 清空容器

    // 生成表头
    const headerRow = document.createElement('div');
    headerRow.className = 'table-row';
    const headers = ['昵称', '团队权限', '加入组织时间', '累计打卡天数', '操作'];
    headers.forEach(headerText => {
        const headerCell = document.createElement('div');
        headerCell.className = 'table-cell';
        headerCell.textContent = headerText;
        headerRow.appendChild(headerCell);
    });
    container.appendChild(headerRow);

    // 生成行
    data.forEach(item => {
        const row = document.createElement('div');
        row.className = 'table-row';

        Object.keys(item).forEach(key => {
            const cell = document.createElement('div');
            cell.className = 'table-cell';
            cell.textContent = item[key];
            row.appendChild(cell);
        });

        // 添加操作单元格
        const actionsCell = document.createElement('div');
        actionsCell.className = 'table-cell actions-cell';
        const detailBtn = document.createElement('span');
        detailBtn.className = 'button';
        detailBtn.textContent = '详情';
        detailBtn.onclick = () => console.log('详情 clicked');
        const deleteBtn = document.createElement('span');
        deleteBtn.className = 'button';
        deleteBtn.textContent = '删除';
        deleteBtn.onclick = function(){
            let username=item.nickname;
            let teamName=document.getElementById("team-name").textContent;
            console.log(username);
            console.log(teamName);
            let token=localStorage.getItem('token');
            let url="http://47.113.117.103:8080/api/team_manage/member_manage/delete";
            fetch(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer '+token
                },
                body: JSON.stringify({
                    username:username,
                    teamname:teamName
                })
                }).then(response => response.json())
                .then(data => {
                    console.log(data);
                    if(data.code==200){
                        console.log("删除成功");
                        allTestStatistics();
                    }else{
                        console.log("删除失败");
                    }
                }).catch(error => {
                    console.log(error);
                });
        };
        actionsCell.appendChild(detailBtn);
        actionsCell.appendChild(deleteBtn);
        row.appendChild(actionsCell);

        container.appendChild(row);
    });
}

///分页按钮
function createPaginationButtons() {
    const pagination = document.getElementById("pagination");
    // 清空之前的按钮
    pagination.innerHTML = '';

    // 创建并添加前一页按钮
    const prevButton = document.createElement("button");
    prevButton.innerText = '<';
    prevButton.addEventListener("click",()=>{
        page = Math.max(1, page - 1); // 更新页码前需要对页码进行合法性检查
        allTestStatistics();
    });
    pagination.appendChild(prevButton);

    // 创建并添加页数按钮
    updatePagination(BUTTON_NUM); // 根据当前页数和总页数创建页数按钮

    // 创建并添加后一页按钮
    const nextButton = document.createElement("button");
    nextButton.innerText = '>';
    nextButton.addEventListener("click",()=>{
        page = Math.min(totalPage, page + 1); // 更新页码前需要对页码进行合法性检查
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

    for (let i = startPage; i <= endPage; i++) {
        const button = document.createElement('button');
        button.textContent = i;
        button.addEventListener('click', () => {
            page = i;
            allTestStatistics();
        });

        if (i === page) {
            button.classList.add('active');
        }

        pagination.appendChild(button);
    }
}
