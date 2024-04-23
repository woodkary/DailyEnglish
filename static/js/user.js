const toast = document.querySelector(".toast")
closeIcon = document.querySelector(".close")
progress = document.querySelector(".progress");

window.onload=function(){
    console.log('这里什么都没有，骗你的，哈哈哈！');
    closeIcon.addEventListener("click", () => {
        toast.classList.remove("active");

        setTimeout(() => {
            progress.classList.remove("active");
        }, 300);

        clearTimeout(timer1);
    });
    initializeCodeMap();
    initializeTeamCode();
    initializeInput();
    /*toggleToast();*/
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
    for(let i=0;i<copyBtns.length;i++) {
        let copyBtn = copyBtns[i];
        copyBtn.addEventListener('click', () => {
            let code = copyBtn.parentNode.querySelector('span').textContent;
            navigator.clipboard.writeText(code).then(r => { // 复制成功
                console.log('复制成功');
                let title=toast.querySelector('.text-1');
                let message=toast.querySelector('.text-2');
                title.textContent='复制成功';
                message.textContent='邀请码已复制到剪贴板，请妥善保管！';

                if(toast.classList.contains("active")){
                    toast.classList.remove("active");
                    progress.classList.remove("active");
                    clearTimeout(timer1);
                    setTimeout(() => {
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
                    },300);
                }else{
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
                    },5000);
                }


            });
        });
    }
    let generateBtns = document.querySelectorAll('.generateBtn');
    for(let i=0;i<generateBtns.length;i++) {
        let generateBtn = generateBtns[i];
        generateBtn.addEventListener('click', () => {
            let teamname = generateBtn.parentNode.parentNode.parentNode.querySelector('.group-name').textContent;
            let newCode = generateTeamCode(teamname);
            let code = generateBtn.parentNode.querySelector('span');
            code.textContent = newCode;
            let title=toast.querySelector('.text-1');
            let message=toast.querySelector('.text-2');
            title.textContent='重置邀请码成功';
            message.textContent='邀请码已重置,请妥善保存';
            if(toast.classList.contains("active")){
                toast.classList.remove("active");
                progress.classList.remove("active");
                clearTimeout(timer1);
                setTimeout(() => {
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
                },300);
            }else{
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
                },5000);
            }
        })
    }
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
    if (codeAndExpiry == null) {
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

