function createTableRow(message, time) {
    // 创建行元素
    const row = document.createElement('div');
    row.className = 'table-row';

    // 创建图标元素
    const icon = document.createElement('span');
    icon.className = 'table-icon';
    // 根据文本内容设置图标
    if (message.includes('Success')) {
        icon.textContent = '✅';
    } else if (message.includes('Error')) {
        icon.textContent = '❌';
    } else {
        icon.textContent = '🔹';
    }

    // 创建消息元素
    const messageElement = document.createElement('span');
    messageElement.className = 'table-message';
    messageElement.textContent = message;

    // 创建时间元素
    const timeElement = document.createElement('span');
    timeElement.className = 'table-time';
    timeElement.textContent = time;

    // 组合行内的元素
    row.appendChild(icon);
    row.appendChild(messageElement);
    row.appendChild(timeElement);

    return row;
}