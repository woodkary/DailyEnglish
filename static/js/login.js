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
                    'Authorization': 'Bearer' + token
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
