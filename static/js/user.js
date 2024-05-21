<<<<<<< Updated upstream
=======
const toast = document.querySelector(".toast")
closeIcon = document.querySelector(".close")
progress = document.querySelector(".progress");

let allTeams = [];
//TODO 在team选项卡打开时，把allTeam的内容更新上去

function getPersonalInfo(){
    let token=localStorage.getItem('token');
    fetch('http://47.113.117.103:8080/api/team_manage/personal_center/data',{
        method: 'GET',
        headers: {
            'Authorization': 'Bearer '+token
        }
    })
   .then(response => response.json())
   .then(data => {
        console.log(data);
        let nameP = document.getElementById('name');
        let emailP = document.getElementById('email');
        let phoneP = document.getElementById('phone');
        nameP.textContent = data.user.name;
        emailP.textContent = data.user.email;
        phoneP.textContent = data.user.phone;
        data.user.teams.forEach(team => {
            //假设每个团队10人，这里可以根据实际情况调整
            //TODO 如果后端更新了查人数，就要用team来查，而不是用10
            allTeams.push({team:team,num:10});
        })
    })
}

>>>>>>> Stashed changes
window.onload=function(){
    console.log('这里什么都没有，骗你的，哈哈哈！');
    initializeCodeMap();
    toggleToast();
    initializeInput();
    initializeTeamCode();
}
//先放入一个map存储所有的code和过期时间
function initializeCodeMap () {
    let codeMap = JSON.parse(localStorage.getItem('codeMap'));
    if (codeMap == null) {
        codeMap = {};
        localStorage.setItem('codeMap', JSON.stringify(codeMap));
    }
    let groupInfos = document.querySelectorAll('.team-card');
    groupInfos.forEach(groupInfo => {
        let teamname = groupInfo.querySelector('.group-name').textContent;
        let code = getInvitationCode(teamname);
        let teamcode = groupInfo.querySelector('.teamcode');
        teamcode.querySelector('span').textContent = code;
    });
}
function initializeInput () {
    var editBtn = document.getElementById('editProfileBtn');
    var saveBtn = document.getElementById('saveBtn');
    var profileContent = document.getElementById('myTabContent');

    if(editBtn!== null) {
        editBtn.addEventListener('click', function (evt) {
            evt.preventDefault();
            // 将所有的p元素变为输入框
            var ps = profileContent.getElementsByTagName('p');
            for (var i = ps.length - 1; i >= 0; i--) { // 使用倒序循环
                var p = ps[i];
                var input = document.createElement('input');
                input.type = 'text';
                input.value = p.textContent;
                input.classList.add('message-input');
                p.parentNode.replaceChild(input, p);
            }
            // 显示保存按钮
            if(saveBtn!== null) {
                saveBtn.style.display = 'block';
                saveBtn.addEventListener('click', function (evt) {
                    evt.preventDefault();
                    // 将所有的输入框变回p元素，并保存输入框中的文本
                    var inputs = profileContent.getElementsByTagName('input');
                    for (var i = inputs.length - 1; i >= 0; i--) {
                        var input = inputs[i];
                        var p = document.createElement('p');
                        p.textContent = input.value;
                        input.parentNode.replaceChild(p, input);
                    }

                    // 隐藏保存按钮
                    saveBtn.style.display = 'none';
                });
            }
        });
    }
}
/*document.addEventListener('DOMContentLoaded', );*/

function initializeTeamCode () {
    let copyBtns = document.querySelectorAll('.copyBtn');
    copyBtns.forEach(copyBtn => {
        copyBtn.addEventListener('click', () => {
            let code = copyBtn.parentNode.querySelector('span').textContent;
            navigator.clipboard.writeText(code).then(r => { // 复制成功
                copyBtn.textContent = 'Copied!';
                setTimeout(() => {
                    copyBtn.textContent = '复制团队码';
                }, 500);
            });
        });
    });
    let generateBtns = document.querySelectorAll('.generateBtn');
    generateBtns.forEach(generateBtn => {
        generateBtn.addEventListener('click', () => {
            let teamname = generateBtn.parentNode.parentNode.querySelector('.group-name').textContent;
            let newCode = generateTeamCode(teamname);
            let code = generateBtn.parentNode.querySelector('.teamcode').querySelector('span');
            code.textContent = newCode;
        });
    });
}
// Toggling invitation code generation

function generateInvitationCode() {
    let code = "#";
    let possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    for (let i = 0; i < 10; i++) {
        code += possible.charAt(Math.floor(Math.random() * possible.length));
    }
    return code;
}
function getInvitationCode(teamname) {
    let codeMap = JSON.parse(localStorage.getItem('codeMap'));
    let codeAndExpiry = codeMap[teamname];
    if (codeAndExpiry === null) {
        let code = generateInvitationCode();
        let expiry = new Date();
        expiry.setDate(expiry.getTime() + 1000*60*5); // 设置5分钟后过期
        codeMap[teamname] = { code: code, expiry: expiry };
        localStorage.setItem('codeMap', JSON.stringify(codeMap));
        return code;
    } else {
        let expiry = new Date(codeAndExpiry.expiry);
        let now = new Date();
        if (expiry < now) { // 过期
            let code = generateInvitationCode();
            codeMap[teamname] = { code: code, expiry: new Date(now.getTime() + 1000*60*5) }; // 续期5分钟
            localStorage.setItem('codeMap', JSON.stringify(codeMap));
            return code;
        } else {
            return codeAndExpiry.code;
        }
    }
}
function generateTeamCode(teamname) {
    let newCode = generateInvitationCode();
    let codeMap = JSON.parse(localStorage.getItem('codeMap'));
    codeMap[teamname] = { code: newCode, expiry: new Date(new Date().getTime() + 1000*60*5) }; // 5分钟后过期
    localStorage.setItem('codeMap', JSON.stringify(codeMap));
    return newCode;
}

function toggleToast(){
    let toast = document.getElementById('toast');
    let overlay = document.getElementById('overlay');
    let showToastBtn = document.getElementById('showToast');
    let toastClose=document.getElementById('toastClose');
    showToastBtn.addEventListener('click', function () {
        overlay.style.display = 'block';
        toast.style.display = 'block';
        /*//两秒后自动关闭toast
        setTimeout(() => {
            toast.style.display = 'none';
            overlay.style.display = 'none';
        }, 2000);*/
    });
    toastClose.addEventListener('click', function () {
        toast.style.display = 'none';
        overlay.style.display = 'none';
    });
}

