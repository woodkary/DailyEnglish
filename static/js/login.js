let switchCtn = document.querySelector("#switch-cnt");
let switchC1 = document.querySelector("#switch-c1");
let switchC2 = document.querySelector("#switch-c2");
let switchCircle = document.querySelectorAll(".switch_circle");
let switchBtn = document.querySelectorAll(".switch-btn");
let aContainer = document.querySelector("#a-container");
let bContainer = document.querySelector("#b-container");
let allButtons = document.querySelectorAll(".submit");

let getButtons = (e) => e.preventDefault()
let changeForm = (e) => {
    // 修改类名
    switchCtn.classList.add("is-gx");
    setTimeout(function () {
        switchCtn.classList.remove("is-gx");
    }, 1500)
    switchCtn.classList.toggle("is-txr");
    switchCircle[0].classList.toggle("is-txr");
    switchCircle[1].classList.toggle("is-txr");

    switchC1.classList.toggle("is-hidden");
    switchC2.classList.toggle("is-hidden");
    aContainer.classList.toggle("is-txl");
    bContainer.classList.toggle("is-txl");
    bContainer.classList.toggle("is-z");
}
// 点击切换
let shell = (e) => {
    for (let i = 0; i < allButtons.length; i++)
        allButtons[i].addEventListener("click", getButtons);
    for (let i = 0; i < switchBtn.length; i++)
        switchBtn[i].addEventListener("click", changeForm)
}
const toast = document.querySelector(".toast")
closeIcon = document.querySelector(".close")
progress = document.querySelector(".progress");
window.addEventListener("load", shell);
let timer1;

closeIcon.addEventListener("click", function () {
    toast.classList.remove("active");
    progress.classList.remove("active");
    clearTimeout(timer1);
});

window.onload = function () {
    // 点击波纹效果
    const rippleContainer = document.querySelector('body');
    console.log(rippleContainer);

    rippleContainer.addEventListener('click', function(e) {
        const circle = document.createElement('div');
        /*circle.style.zIndex=100;*/
        //排除按钮、链接、输入框、图片
        if(e.target.tagName === 'BUTTON' || e.target.tagName === 'A'|| e.target.tagName === 'INPUT'|| e.target.tagName === 'IMG'){
            return;
        }
        circle.classList.add('ripple-effect');

        // 设置波纹效果的大小
        /*const diameter = Math.max(rippleContainer.clientWidth, rippleContainer.clientHeight);*/
        const diameter = 50;
        circle.style.width = circle.style.height = `${diameter}px`;

        // 根据点击位置设置波纹效果的位置
        const rect = rippleContainer.getBoundingClientRect();
        circle.style.top = `${e.clientY - rect.top - diameter / 2}px`;
        circle.style.left = `${e.clientX - rect.left - diameter / 2}px`;

        rippleContainer.appendChild(circle);

        // 波纹动画结束后删除元素
        circle.addEventListener('animationend', function() {
            circle.classList.remove('ripple-effect');
            circle.remove();
        });
    });
}

function login(event) {
    event.preventDefault();
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    const data = {
        username: username,
        password: password
    };
    fetch('http://47.113.117.103:8080/api/team_manager/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then(response => {
        console.log(response);
        return response.json();
    }).then(data => {
        console.log(data);
        if (data.code == 200) {
            localStorage.setItem("token", data.token);
            //获取管理员所管理的团队id和团队名称，这是个map类型的json数据
            localStorage.setItem("team_info", data.team_info);
            let token = data.token;
            console.log(token);
           /* let url = 'http://47.113.117.103:8080/api/team_manage/personal_center/data';*/
/*            fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': 'Bearer ' + token
                }
            }).then(response => {
                console.log(response);
                return response.json();
            }).then(data => {
                console.log(data);
                if (data.code == 200) {
                    console.log(data.data);
                    localStorage.setItem("name", data.user.name);
                    localStorage.setItem("email", data.user.email);
                    localStorage.setItem("team", data.user.team);
                    localStorage.setItem("right", data.user.right);
                } else {
                    console.log(data.message);
                }
            }).catch(error => {
                console.log(error);
            })*/
            /*window.location.href = 'http://47.113.117.103:8080/api/team_manager/index';*/
            window.location.href = 'user.html';//跳转到主页,为了能展示，先暂存
        } else {
            var message = document.getElementById("verification-message");
            message.textContent = data.message;
            message.style.color = "red";
        }
    }).then(error => {
        console.log(error);
    })

}
function toggleToast(t,m){
    let title=toast.querySelector('.text-1');
    let message=toast.querySelector('.text-2');
    title.textContent=t;
    message.textContent=m;
    toast.classList.add("active");
    progress.classList.add("active");
    timer1 = setTimeout(() => {
        progress.style.width = "100%";
        let title=toast.querySelector('.text-1');
        let message=toast.querySelector('.text-2');
        title.textContent='';
        message.textContent='';
        toast.classList.remove("active");
        progress.classList.remove("active");
    },5000)
}
function register(event) {
    event.preventDefault();
    let username = document.getElementById("username-register").value;
    let email = document.getElementById("email-register").value;
    let password = document.getElementById("password-register").value;
    let passwordRetype = document.getElementById("password-register-retype").value;
    if (password !== passwordRetype) {
        //TODO 不应使用toast，应使用文本提示
        toggleToast("提示","两次密码输入不一致");
        return;
    }
    fetch('http://47.113.117.103:8080/api/team_manager/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: username,
            email: email,
            password: password
        })
    }).then(response => {
        console.log(response);
        if(response.status==200){
            console.log("注册成功");
            toggleToast("提示","注册成功");
        }else{
            console.log("注册失败");
            toggleToast("提示","注册失败");
        }

    })
}
function sendCode(event) {
    event.preventDefault();
    let button = event.target;
    let email = document.getElementById("email-register").value;
    fetch("http://47.113.117.103:8080/api/team_manager/send_code", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({
            email: email
        })
    }).then(response => {
        console.log(response);
        return response.json();
    }).then(data => {
        console.log(data);
        if (data.code == 200) {
            let vCode = data.data;
            let codeAndExpiry = {
                vCode: vCode,
                expiry: new Date().getTime() + 1000 * 60 * 5//5分钟有效
            };
            localStorage.setItem("codeAndExpiry", JSON.stringify(codeAndExpiry));

            // 设置初始倒计时时间（秒）
            let timeLeft = 60;
            // 禁用按钮
            button.disabled = true;
            // 更新按钮文字
            button.textContent = `${timeLeft}秒后请重试`;
            // 设置定时器，每秒更新一次按钮文字
            const timer = setInterval(() => {
                timeLeft -= 1;
                button.textContent = `${timeLeft}秒后请重试`;
                // 当倒计时结束时
                if (timeLeft <= 0) {
                    // 清除定时器
                    clearInterval(timer);
                    // 启用按钮
                    button.disabled = false;
                    // 恢复按钮文字
                    button.textContent = `发送验证码`;
                }
            }, 1000);
        }else if(data.code==409){//邮箱已注册
            //TODO 不应使用toast，应使用文本提示
            toggleToast("提示","邮箱已注册");
        }else if(data.code==400){//请求参数错误
            //TODO 不应使用toast，应使用文本提示
            toggleToast("提示","请求参数错误");
        }else{
            //TODO 不应使用toast，应使用文本提示
            toggleToast("提示","发送验证码失败");
            
        }
    }).catch(error => {
        console.log(error);
    })
}
let initialEmail="";

let emailInput = document.getElementById("email-register");
emailInput.addEventListener("focus", function () {
    initialEmail = emailInput.value;
});
emailInput.addEventListener("blur", function () {
    if (emailInput.value !== initialEmail) {
        if(!checkEmail(emailInput.value)){
            //TODO 不应使用toast，应使用文本提示
            toggleToast("提示","邮箱格式错误");
        }
    }
})
function checkEmail(email) {
    var regex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
    return regex.test(email);
}

let initialVerifyCodeInput = "";
let verifyCodeInput = document.getElementById("verification-code-register");
verifyCodeInput.addEventListener("focus", function () {
    initialVerifyCodeInput = verifyCodeInput.value;
});
verifyCodeInput.addEventListener("blur", function () {
    if (verifyCodeInput.value !== initialVerifyCodeInput) {
        if(!checkVerifyCode(verifyCodeInput.value)){
            //TODO 不应使用toast，应使用文本提示
            toggleToast("提示","验证码错误");
        }else{
            let passwordInput = document.getElementById("password-register");
            let passwordRetypeInput = document.getElementById("password-register-retype");
            let registerBtn = document.getElementById("register-btn");
            passwordInput.style.display = "block";
            passwordRetypeInput.style.display = "block";
            registerBtn.style.display = "block";
        }
    }
})
function checkVerifyCode(verifyCode) {
    let codeAndExpiry = localStorage.getItem("codeAndExpiry");
    if (codeAndExpiry) {
        codeAndExpiry = JSON.parse(codeAndExpiry);
        if (codeAndExpiry.expiry > new Date().getTime()) {
            if (verifyCode === codeAndExpiry.vCode) {
                return true;
            }
        }
    }
    return false;
}
