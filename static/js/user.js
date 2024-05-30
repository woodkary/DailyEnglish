const toast = document.querySelector(".toast")
closeIcon = document.querySelector(".close")
progress = document.querySelector(".progress");

let allTeams = [];
//TODO 在team选项卡打开时，把allTeam的内容更新上去

function getPersonalInfo(){
    let token=localStorage.getItem('token');
    fetch('http://localhost:8081/api/team_manage/personal_center/data',{
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
            nameP.textContent = data.name;
            emailP.textContent = data.email;
            phoneP.textContent = data.phone;
            data.team.forEach(t => {
                //假设每个团队10人，这里可以根据实际情况调整
                allTeams.push({teamName:t.team_name,memberNum:t.member_num,teamId:t.team_id});
            });
            renderTeamInfo();
        }).catch(error => {
        //初始化默认的个人信息
        console.log(error);
        let nameP = document.getElementById('name');
        let emailP = document.getElementById('email');
        let phoneP = document.getElementById('phone');
        nameP.textContent = '未登录';
        emailP.textContent = '未登录';
        phoneP.textContent = '未登录';
        allTeams.push({teamName:'未加入任何团队',memberNum:0,teamId:0});
        renderTeamInfo();
    });
}
async function renderTeamInfo(){
    let teamInfoChart=document.querySelector('#tab_2');
    teamInfoChart.innerHTML='';//清空原有内容
    for (const t of allTeams) {
        let card = document.createElement('card');
        card.classList.add('team-card');
        let img=document.createElement('img');
        img.src='./image/team.svg';
        img.classList.add('group-img');
        img.alt='team';
        card.appendChild(img);
        //显示团队信息
        let teamDetails=document.createElement('div');
        teamDetails.classList.add('group-details');
        let teamInfo=document.createElement('div');
        teamInfo.classList.add('group-info');
        let teamNameSpan=document.createElement('span');
        teamNameSpan.classList.add('group-name');
        teamNameSpan.textContent=t.teamName;
        let memberNumSpan=document.createElement('span');
        memberNumSpan.classList.add('col-md-6');
        memberNumSpan.classList.add('group-peopleAmount');
        memberNumSpan.textContent=t.memberNum+'人';
        teamInfo.appendChild(teamNameSpan);
        teamInfo.appendChild(memberNumSpan);
        teamDetails.appendChild(teamInfo);
        //显示邀请码
        let teamCodeDiv=document.createElement('div');
        teamCodeDiv.classList.add('group-code');
        let teamCodeP=document.createElement('p');
        teamCodeP.classList.add('teamcode');
        teamCodeP.textContent='邀请码：';
        let teamCodeSpan=document.createElement('span');
        let codeMap = JSON.parse(sessionStorage.getItem('codeMap'));
        //创建map保存团队名和邀请码
        if (codeMap == null) {
            codeMap = {};
            sessionStorage.setItem('codeMap', JSON.stringify(codeMap));
        }
        teamCodeSpan.textContent=await getInvitationCode(t.teamName,t.teamId);
        teamCodeP.appendChild(teamCodeSpan);
        let copyBtn=document.createElement('i');
        copyBtn.classList.add('uil');
        copyBtn.classList.add('uil-copy');
        copyBtn.classList.add('copyBtn');
        copyBtn.addEventListener('click', () => {
            let code = copyBtn.parentNode.querySelector('span').textContent;
            navigator.clipboard.writeText(code).then(r => { // 复制成功
                console.log('复制成功');
                let title=toast.querySelector('.text-1');
                let message=toast.querySelector('.text-2');
                title.textContent='复制成功';
                message.textContent='邀请码已复制到剪贴板！';

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
        teamCodeP.appendChild(copyBtn);
        let generateBtn=document.createElement('i');
        generateBtn.classList.add('uil');
        generateBtn.classList.add('uil-redo');
        generateBtn.classList.add('generateBtn');
        generateBtn.addEventListener('click', async () => {
            let teamname = generateBtn.parentNode.parentNode.parentNode.querySelector('.group-name').textContent;
            let newCode = await generateTeamCode(teamname, t.teamId);
            let code = generateBtn.parentNode.querySelector('span');
            code.textContent = newCode;
            let title = toast.querySelector('.text-1');
            let message = toast.querySelector('.text-2');
            title.textContent = '重置邀请码成功';
            message.textContent = '邀请码已重置';
            if (toast.classList.contains("active")) {
                toast.classList.remove("active");
                progress.classList.remove("active");
                clearTimeout(timer1);
                setTimeout(() => {
                    toast.classList.add("active");
                    progress.classList.add("active");
                    timer1 = setTimeout(() => {
                        progress.style.width = "100%";
                        let title = toast.querySelector('.text-1');
                        let message = toast.querySelector('.text-2');
                        title.textContent = '';
                        message.textContent = '';
                        toast.classList.remove("active");
                        progress.classList.remove("active");
                    }, 5000)
                }, 300);
            } else {
                toast.classList.add("active");
                progress.classList.add("active");
                timer1 = setTimeout(() => {
                    progress.style.width = "100%";
                    let title = toast.querySelector('.text-1');
                    let message = toast.querySelector('.text-2');
                    title.textContent = '';
                    message.textContent = '';
                    toast.classList.remove("active");
                    progress.classList.remove("active");
                }, 5000);
            }
        });
        teamCodeP.appendChild(generateBtn);
        teamCodeDiv.appendChild(teamCodeP);
        teamDetails.appendChild(teamCodeDiv);
        card.appendChild(teamDetails);
        teamInfoChart.appendChild(card);
    }
}

window.onload=function(){
    console.log('这里什么都没有，骗你的，哈哈哈！');
    let aboutSpan=document.querySelector('.tabs').querySelector('#about');
    aboutSpan.dispatchEvent(new Event('click'));
    closeIcon.addEventListener("click", () => {
        toast.classList.remove("active");

        setTimeout(() => {
            progress.classList.remove("active");
        }, 300);

        clearTimeout(timer1);
    });
    getPersonalInfo();
    initializeInput();
    /*toggleToast();*/
}
//初始化输入框
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
//TODO 改为向后端请求邀请码，而不是自己生成
async function generateInvitationCode(teamId) {
    const response = await fetch('http://localhost:8081/api/team_manage/refresh_team_code', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': 'Bearer ' + localStorage.getItem('token')
        },
        body: JSON.stringify({
            team_id: teamId,
        })
    });
    code = (await response.json()).invitation_code;
    return code;
}
async function getInvitationCode(teamname, teamId) {
    let codeMap = JSON.parse(sessionStorage.getItem('codeMap'));
    let codeAndExpiry = codeMap[teamname];
    if (codeAndExpiry == null) {
        let code = await generateInvitationCode(teamId);
        let expiry = new Date();
        expiry.setDate(expiry.getTime() + 1000 * 60 * 5); // 设置5分钟后过期
        codeMap[teamname] = {code: code, expiry: expiry};
        sessionStorage.setItem('codeMap', JSON.stringify(codeMap));
        return code;
    } else {
        let expiry = new Date(codeAndExpiry.expiry);
        let now = new Date();
        if (expiry < now) { // 过期
            let code = await generateInvitationCode(teamId);
            codeMap[teamname] = {code: code, expiry: new Date(now.getTime() + 1000 * 60 * 5)}; // 续期5分钟
            sessionStorage.setItem('codeMap', JSON.stringify(codeMap));
            return code;
        } else {
            return codeAndExpiry.code;
        }
    }
}
async function generateTeamCode(teamname, teamId) {
    let newCode = await generateInvitationCode(teamId);
    let codeMap = JSON.parse(sessionStorage.getItem('codeMap'));
    codeMap[teamname] = {code: newCode, expiry: new Date(new Date().getTime() + 1000 * 60 * 5)}; // 5分钟后过期
    sessionStorage.setItem('codeMap', JSON.stringify(codeMap));
    return newCode;
}
