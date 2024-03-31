function logout(){
    sessionStorage.removeItem("token");
    window.location.href = 'login.html';
}
function myFunction() {
    document.getElementById("myDropdown").classList.toggle("show");
}

// 点击下拉菜单以外的地方时隐藏下拉内容
window.onclick = function(event) {
    if (!event.target.matches('.dropbtn')) {
        var dropdowns = document.getElementsByClassName("dropdown-content");
        var i;
        for (i = 0; i < dropdowns.length; i++) {
            var openDropdown = dropdowns[i];
            if (openDropdown.classList.contains('show')) {
                openDropdown.classList.remove('show');
            }
        }
    }
}