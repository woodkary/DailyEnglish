const buttons = document.getElementsByTagName("button"),
    toast = document.querySelector(".toast")
closeIcon = document.querySelector(".close")
    progress = document.querySelector(".progress");

let timer1, timer2;
//对所有按钮添加点击事件
for (let i = 0; i < buttons.length; i++) {
    buttons[i].addEventListener("click", () => {
        toast.classList.add("active");
        progress.classList.add("active");

        timer1 = setTimeout(() => {
            progress.style.width = "100%";
        })
    })
}

closeIcon.addEventListener("click", () => {
    toast.classList.remove("active");

    setTimeout(() => {
        progress.classList.remove("active");
    }, 300);

    clearTimeout(timer1);
    clearTimeout(timer2);
});