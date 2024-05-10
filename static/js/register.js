//发送验证码
function sendVerificationCode() {
    //
    var email = document.getElementById("email").value;
    //错误信息元素
    var errorDiv = document.getElementById("error");
    var data = {
        email: email
    };
    fetch('http://47.113.117.103:8080/api/register/sendCode', {
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
        errorDiv.style.display = "block";
        if (data.code === 200) {
            document.getElementById("verificationCode").classList.remove("error-border");
            errorDiv.style.color = "green";
            errorDiv.textContent = "验证码发送成功";
            return true; //验证码发送成功

        } else {
            document.getElementById("verificationCode").classList.add("error-border");
            return false; //验证码发送失败
        }
    })
}
//验证密码是否输入相同
function validateForm() {
    document.getElementById("username").value;
    var password = document.getElementById("password").value;
    var re_pwd = document.getElementById("re_pwd").value;
    var errorDiv = document.getElementById("duplicatedPassword");

    if (password !== re_pwd) {
        errorDiv.style.display = "block";
        document.getElementById("re_pwd").classList.add("error-border");
        return false; // 密码不一致，显示错误信息
    } else {
        errorDiv.style.display = "none";
        document.getElementById("re_pwd").classList.remove("error-border");
        return true; // 密码一致
    }
}
//注册
function register() {
    var email = document.getElementById("email").value;
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;

    validateForm();

    var errorDiv = document.getElementById("errorReg");
    var data = {
        email: email,
        username: username,
        password: password
    };
    fetch('http://47.113.117.103:8080/api/users/register', {
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
        errorDiv.style.display = "block";
        if (data.code == 200) {
            errorDiv.style.color = "green";
            errorDiv.textContent = "注册成功";
            errorDiv.classList.remove("error-border");
            return true; //注册成功

        } else {
            errorDiv.classList.add("error-border");
            return false; //注册失败
        }
    })
}