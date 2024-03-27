// 假设从后端接收到的数据格式如下：
let cardData = [
    { word: "四级词汇", date: "2024-03-27", wordOfTheDay: 20, completionCount: [10,50] },
    { word: "四级词汇", date: "2024-03-28", wordOfTheDay: 20, completionCount: [15,50] },
    { word: "四级词汇", date: "2024-03-29", wordOfTheDay: 20, completionCount: [20,50] }
];

let examData=[
    {date:"2024-03-27",range:"四级词汇",avgScore:20,completionCount:[10,50]},
    {date: "2024-03-28", range: "四级词汇", avgScore: 20, completionCount: [15, 50] },
    {date: "2024-03-29", range: "四级词汇", avgScore: 20, completionCount: [20, 50] }
];
window.onload = function () {
    setCardData();
    setExamData();

}





function setCardData(){
    // 获取表格的 tbody 元素
    let tbody = document.getElementById('cardData');
    // 遍历从后端接收到的数据，并将每一条数据插入表格中
    cardData.forEach(data => {
        let row = tbody.insertRow(); // 插入新的行
        let wordCell = row.insertCell(0); // 插入打卡词汇单元格
        let dateCell = row.insertCell(1); // 插入打卡日期单元格
        let wordOfTheDayCell = row.insertCell(2); // 插入打卡单词单元格
        let completionCountCell = row.insertCell(3); // 插入完成人数单元格
        let detailCountCell=row.insertCell(4); //插入详情单元格
        let emptyRow=tbody.insertRow();//插入空行
        emptyRow.insertCell(0);

        // 设置单元格的文本内容
        wordCell.textContent = data.word;
        dateCell.textContent = data.date;
        wordOfTheDayCell.textContent = data.wordOfTheDay+"人";
        completionCountCell.textContent = data.completionCount[0]+"/"+data.completionCount[1]+"人";
        detailCountCell.innerHTML = "<a href='signInDetail.html'>详情></a>";//设置详情单元格的跳转链接，页面未指定

        // 设置单元格的样式
        wordCell.style.fontFamily = "黑体";
        dateCell.style.fontFamily = "黑体";
        wordOfTheDayCell.style.fontFamily = "黑体";
        completionCountCell.style.fontFamily = "黑体";
        detailCountCell.style.fontFamily = "黑体";

    });

}
function setExamData(){
    let tbody = document.getElementById('examData');
    examData.forEach(data => {
        let row = tbody.insertRow(); // 插入新的行
        let dateCell = row.insertCell(0); // 插入打卡日期单元格
        let rangeCell = row.insertCell(1); // 插入打卡范围单元格
        let avgScoreCell = row.insertCell(2); // 插入平均分单元格
        let completionCountCell = row.insertCell(3); // 插入完成人数单元格
        let detailCountCell=row.insertCell(4); //插入详情单元格
        let emptyRow=tbody.insertRow();//插入空行
        emptyRow.insertCell(0);

        // 设置单元格的文本内容
        dateCell.textContent = data.date;
        rangeCell.textContent = data.range;
        avgScoreCell.textContent = data.avgScore+"分";
        completionCountCell.textContent = data.completionCount[0]+"/"+data.completionCount[1]+"人";
        detailCountCell.innerHTML = "<a href='examDetail.html'>详情></a>";//设置考试详情单元格的跳转链接，页面未指定

        // 设置单元格的样式
        dateCell.style.fontFamily = "黑体";
        rangeCell.style.fontFamily = "黑体";
        avgScoreCell.style.fontFamily = "黑体";
        completionCountCell.style.fontFamily = "黑体";
        detailCountCell.style.fontFamily = "黑体";

    });
}
