let examName = "";
window.onload = function() {
    let urlParams = new URLSearchParams(window.location.search);
    examName = urlParams.get('name');
    console.log(examName);
    buttonHref();
}
function buttonHref() {
    let buttons = document.getElementsByClassName("BlueButton");
    for (let i = 0; i < buttons.length; i++) {
        buttons[i].addEventListener("click", function() {
            let parentTd = this.parentNode.parentNode;
            let name=parentTd.getElementsByTagName("td")[0].innerText;
            window.location.href = "personal-test-details.html?name="+name;

        });
    }
}