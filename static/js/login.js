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
    fetch('http://localhost:8080/api/team_manager/login', {
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
            let token = data.token;
            console.log(token);
            let url = 'http://localhost:8080/api/team_manage/personal_center/data';
            fetch(url, {
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
            })
            /*window.location.href = 'http://localhost:8080/api/team_manager/index';*/
            window.location.href = 'index.html';//跳转到主页,为了能展示，先暂存
        } else {
            var message = document.getElementById("verification-message");
            message.textContent = data.message;
            message.style.color = "red";
        }
    }).then(error => {
        console.log(error);
    })

}
