// Selecting form and input elements
const form = document.querySelector("form");
const passwordInput = document.querySelector(".passwordInput");
let state=0;
const MAX_TEAMNUM = 10;

// Function to display error messages
const showError = (field, errorText) => {
    field.classList.add("error");
    const errorElement = document.createElement("small");
    errorElement.classList.add("error-text");
    errorElement.innerText = errorText;
    field.closest(".form-group").appendChild(errorElement);
}

// Function to handle form submission
const handleFormData = (e) => {
    e.preventDefault();

    // Retrieving input elements
    const teamnameInput = document.getElementById("teamname");
    const teamnumInput = document.getElementById("teamnum");

    // Getting trimmed values from input fields
    const teamname = teamnameInput.value.trim();
    const teamnum = parseInt(teamnumInput.value.trim());



    // Clearing previous error messages
    document.querySelectorAll(".form-group .error").forEach(field => field.classList.remove("error"));
    document.querySelectorAll(".error-text").forEach(errorText => errorText.remove());

    // Performing validation checks
    if (teamname === "") {
        showError(teamnameInput, "请输入完整团队名称");
    }
    if (isNaN(teamnum) || teamnum < 1 || teamnum > MAX_TEAMNUM) {
        showError(teamnumInput, "请输入有效的团队编号（1-"+MAX_TEAMNUM+"）");
    }


    // Checking for any remaining errors before form submission
    const errorInputs = document.querySelectorAll(".form-group .error");
    if (errorInputs.length > 0) return;

    if(state===0){
        state=1;
        passwordInput.style.display = "block";

    }
    // Submitting the form
    else if(state===1){
        //Todo: 验证密码
        //todo: 后端生成团队邀请码
        let invitationCode = generateInvitationCode();
        console.log(invitationCode);
        let codeAndExpiry = {
            code: invitationCode,
            expiry: new Date().getTime() + 1000*60*5 //  5 minutes
        };
        localStorage.setItem("codeAndExpiry", JSON.stringify(codeAndExpiry));
        form.submit();
    }
}
// Toggling invitation code generation
function generateInvitationCode(){
    let code = "#";
    const possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    for (let i = 0; i < 10; i++) {
        code += possible.charAt(Math.floor(Math.random() * possible.length));
    }
    return code;
}
function getInvitationCode(){
    let codeAndExpiry = localStorage.getItem("codeAndExpiry");
    if(codeAndExpiry){
        let codeAndExpiryObj = JSON.parse(codeAndExpiry);
        let now = new Date().getTime();
        if(now < codeAndExpiryObj.expiry){
            return codeAndExpiryObj.code;
        }
    }
    getNewCode(codeAndExpiry);
}
function getNewCode(codeAndExpiry){
    let newCode = generateInvitationCode();
    codeAndExpiry = {
        code: newCode,
        expiry: new Date().getTime() + 1000*60*5 //  5 minutes
    };
    localStorage.setItem("codeAndExpiry", JSON.stringify(codeAndExpiry));
    return newCode;
}

// Handling form submission event
form.addEventListener("submit", handleFormData);