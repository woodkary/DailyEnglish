function createDaKaTable(data) {
    // 创建表格元素
    const table = document.createElement('table');
    table.className = 'borderless-table';

    // 创建表头
    const thead = document.createElement('thead');
    const headerRow = document.createElement('tr');
    data.headers.forEach(headerText => {
        const header = document.createElement('th');
        header.textContent = headerText;
        headerRow.appendChild(header);
    });
    thead.appendChild(headerRow);
    table.appendChild(thead);

    // 创建表体
    const tbody = document.createElement('tbody');
    data.rows.forEach(rowData => {
        const row = document.createElement('tr');
        rowData.forEach(cellData => {
            const cell = document.createElement('td');
            cell.textContent = cellData;
            row.appendChild(cell);
        });
        tbody.appendChild(row);
    });
    table.appendChild(tbody);

    // 将表格添加到容器中
    const container = document.getElementById('daka-table');
    container.appendChild(table);
}

// 调用函数，传入数据
function fetchDataAndCreateTable(url) {
    let token = localStorage.getItem('token');
    console.log(token);
    fetch(url,{
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer '+token
        }
    })
        .then(response => {
            if(response.status === 401){
                //token失效
                alert('登录已过期，请重新登录');
                localStorage.removeItem('token');
                window.location.href = './login.html';
            }
            if (!response.ok) {
                throw new Error('Network response was not ok.');
            }
            return response.json();
        })
        .then(data => {
            let chartData={
                headers: ['成员', '打卡天数', '个人打卡率','学习单词数'],
                rows: []
            }
            data.punch_statistics.forEach(item => {
                let rowData = [item.name, item.punch_day, item.punch_rate, item.punch_word];
                chartData.rows.push(rowData);
            });
            console.log(chartData);
            createDaKaTable(chartData);
            res=chartData;
        })
        .catch(error => {
            console.error('Error fetching and creating table:', error);
            let data = {
                headers: ['成员', '打卡天数', '个人打卡率'],
                rows: [
                    ['otto', '666', '87%'], ['otto', '666', '87%'],
                    ['otto', '666', '87%'], ['otto', '666', '87%'], ['otto', '666', '87%'], ['otto', '666', '87%']
                    // 更多数据行
                ]
            };
            createDaKaTable(data);
        });
}