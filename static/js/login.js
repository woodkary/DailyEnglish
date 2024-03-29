function login() {
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
        return response.json();
    }).then(data => {
        if (data.code == 200) {
            console.log(data);
            sessionStorage.setItem("token", data.token);
            window.location.href = 'http://localhost:8080/api/team_manager/static/index.html';
        } else {
            var message = document.getElementById("verification-message");
            message.textContent = data.message;
            message.style.color = "red";
        }
    })
}