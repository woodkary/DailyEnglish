window.onload = function() {
    const ranges = document.querySelectorAll('.container input[type=range]');
    ranges.forEach(range => {
        range.style.setProperty('--value', `${range.value}%`);
        // 获取对应的输入数字
        const numberInput = range.parentElement.querySelector('input[type=number]');
        // 更新输入框的值
        numberInput.value = range.value;
    });
    addInputListener();
}

function addInputListener(){
    // 获取所有的滑块
    const ranges = document.querySelectorAll('.container input[type=range]');

    // 加入监听事件
    ranges.forEach(range => {
        range.addEventListener('input', () => {
            // 更新滑块的值
            range.style.setProperty('--value', `${range.value}%`);

            // 获取对应的输入数字
            const numberInput = range.parentElement.querySelector('input[type=number]');

            // 更新输入框的值
            numberInput.value = range.value;
        });
    });
}