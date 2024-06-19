let studentUploads=[
    {
        userId: 1,

    }
]
function createTable(data) {
    // 清空表格内容
    document.getElementById('tableBody').innerHTML = '';
    // 创建表头
    const tableHead = document.querySelector('thead tr');
    const headers = Object.keys(data[0]); // 获取第一个数据对象的所有属性作为表头
    headers.forEach(header => {
        const th = document.createElement('th');
        th.textContent = header;
        tableHead.appendChild(th);
    });

    // 创建表格行
    data.forEach(item => {
        const row = document.createElement('tr');
        headers.forEach(header => {
            const td = document.createElement('td');
            td.textContent = item[header]; // 设置单元格内容
            row.appendChild(td);
        });
        document.getElementById('tableBody').appendChild(row);
    });
}