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
            sessionStorage.setItem("token", data.token);
            /*window.location.href = 'http://localhost:8080/api/team_manager/index';*/
            window.location.href = 'index';//跳转到主页,为了能展示，先暂存
        } else {
            var message = document.getElementById("verification-message");
            message.textContent = data.message;
            message.style.color = "red";
        }
    }).then(error => {
        console.log(error);
    })
}
