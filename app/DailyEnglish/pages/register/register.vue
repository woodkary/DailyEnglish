<template>
	<view>

		<view class="background">
			<span class="span1"><span class="sign">Sign</span><br>Up</span>
			<image class="pic" src="../../static/register1.svg"></image>
		</view>
		<view class="input-container" id="1">
			<span>昵称</span>
			<input class="input" v-model="username" type="text" placeholder="请输入昵称">
		</view>
		<view class="input-container" id="2">
			<span>邮箱</span>
			<input class="input" id="emailInput" ref="email" v-model="email" type="text" placeholder="请输入邮箱">
			<button class="vtBtn" ref="sendCodeBtn" id="sendCodeBtn" @click="sendCode">发送验证码</button>
		</view>

		<view class="input-container" id="3">
			<span style="left:2rem;">验证码</span>
			<input v-model="verifyCode" id="verifyCodeInput" class="input" @blur="checkInput" type="text" placeholder="请输入验证码">
		</view>

		<view class="input-container" id="4">
			<span>密码</span>
			<input class="input" v-model="password" type="password" placeholder="请输入密码">
		</view>
		<view class="input-container" id="5">
			<input class="input" id="passwordInput" v-model="password2" @blur="checkPasswordInput" type="password" placeholder="再次输入密码">
		</view>
		<button class="login-button" @click="register">注册</button>
		<span class="text">Already have  account?<a>click here to login</a></span>
	</view>
</template>

<script>
	export default {
    data() {
      return {
        username: '',
        email: '',
        password: '',
        password2: '',
        initialVerifyCodeInput: '',
        verifyCode: ''
      }
    },
    methods: {
      checkEmail() {
        let regex = /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
        let emailInput = document.getElementById('emailInput');
        if (!regex.test(this.email)) {
          emailInput.classList.add('inputActive');
          setTimeout(() => {
            emailInput.classList.remove('inputActive');
          }, 2000);
          return false;
        }
        return true;
      },
      sendCode() {
        if (!this.checkEmail()) {
          return;
        }
        let btn = document.getElementById('sendCodeBtn');
        uni.request({
          url: 'http://localhost:8080/api/register/sendCode',
          data: {
            email: this.email
          },
          method: 'POST',
          success: (res) => {
            console.log(res);
            if(res.statusCode === 200){
              let vCode = res.data.data;
              let codeAndExpiry = {
                vCode: vCode,
                expiry: new Date().getTime() + 1000 * 60 * 5//5分钟有效
              };
              //存储验证码和过期时间
              uni.setStorageSync('codeAndExpiry', codeAndExpiry);
              let timeLeft=60;
              btn.disabled=true;
              //按钮可使用状态倒计时
              let timer=setInterval(() => {
                timeLeft--;
                btn.innerText = `${timeLeft}秒后请重试`;
                //倒计时结束，按钮恢复可用状态
                if(timeLeft<=0){
                  clearInterval(timer);
                  btn.innerText='发送验证码';
                  btn.disabled=false;
                }
              }, 1000);
            }else if(data.code==409) {//邮箱已注册
              uni.showToast({
                title: '邮箱已注册',
              });
            }else if(data.code==400){//请求参数错误
              uni.showToast({
                title: '请求参数错误',
              });
            }else{
              uni.showToast({
                title: '发送失败',
              });
            }
          },
          fail: (res) => {
            console.log(res);
            uni.showToast({
              title: '发送失败',
            });
          }
        });
      },
      checkInput(){
        let verifyCode=this.verifyCode;
        let verifyCodeInput=document.getElementById('verifyCodeInput');
        if(verifyCode!==''&&!this.checkVerifyCode(verifyCode)){
          verifyCodeInput.classList.add('inputActive');
          setTimeout(() => {
            verifyCodeInput.classList.remove('inputActive');
          }, 2000);
          return false;
        }
        return true;
      },
      checkVerifyCode(verifyCode) {
        let codeAndExpiry = uni.getStorageSync('codeAndExpiry');
        if (codeAndExpiry) {
          let now = new Date().getTime();
          if (now < codeAndExpiry.expiry) {
            if (verifyCode == codeAndExpiry.vCode) {
              return true;
            }
          }
        }
        return false;
      },
      //密码输入框失去焦点时检查密码是否一致
      checkPasswordInput(){
        let password=this.password;
        let password2=this.password2;
        let passwordInput=document.getElementById('passwordInput');
        if(password!==''&&password2!==''&&password!==password2){
          passwordInput.classList.add('inputActive');
          setTimeout(() => {
            passwordInput.classList.remove('inputActive');
          }, 2000);
          uni.showToast({
            title: '两次输入的密码不一致',
          });
          return false;
        }
        return true;
      },
      //注册
      register() {
        //检查密码是否一致
        /*if(!this.checkPasswordInput()){
          return;
        }*/
        let username=this.username;
        let email=this.email;
        let password=this.password;
        uni.request({
          url: 'http://localhost:8080/api/users/register',
          data: {
            username: username,
            email: email,
            password: password
          },
          method: 'POST',
          success: (res) => {
            console.log(res);
            if(res.statusCode === 200){
              uni.showToast({
                title: '注册成功',
              });
              setTimeout(() => {
                uni.navigateBack();
              }, 1000);
            }else if(data.code==409) {//用户名已注册
              uni.showToast({
                title: '用户名已注册',
              });
            }else if(data.code==400){//请求参数错误
              uni.showToast({
                title: '请求参数错误',
              });
            }else{
              uni.showToast({
                title: '注册失败',
              });
            }
          },
          fail: (res) => {
            console.log(res);
            uni.showToast({
              title: '注册失败',
            });
          }
        });
      }
    }
  }
</script>

<style>
	.background {
		height: 14rem;
		background-color: #fed8c3;
		border-bottom-left-radius: 15%;
		border-bottom-right-radius: 15%;
	}

	.pic {
		margin-left: 8rem;
		top: 2rem;
		width: 14rem;
		transform: scaleX(-1);
	}

	.span1 {
		position: absolute;
		margin-top: 25%;
		font-size: 2.7rem;
		margin-left: 6%;
		font-weight: 600;
		color: white;
		display: block;
		/* 让 <span> 元素变成块级元素 */
		text-align: right;
		/* 文本向右对齐 */
	}

	.sign {
		font-size: 3.2rem;
	}

	.input-container {
		margin-top: 1.5rem;
		display: flex;
		align-items: center;
	}

  .input-container .inputActive {
    border: 1px solid #e74c3c;
  }

	.input-container input {
		width: 86%;
		height: 3.3rem;
		border-radius: 0.7rem;
		background-color: #f0f3f1;
		margin-left: 1.4rem;
		margin-top: 0.4rem;
	}

	.input-container {
		position: relative;
		display: flex;
		margin-top: 1.5rem;
		height: 3.3rem;
		width: 95%;
	}

	.input-container input {
		padding-left: 4rem;
		padding-right: 6rem;
		background-color: transparent;
		border: 1px solid #c4c7ce;
		/*输入的字大小*/
		font-size: 1.1rem;

	}

	.vtBtn {
		position: absolute;
		right: 1rem;
		background-color: transparent;
		white-space: nowrap;
		border-radius: 0%;
		height: 3.3rem;
		width: 6rem;
		top: 6%;
		line-height: 3.3rem;
		font-size: 1rem;
		/* 将行高设置为按钮的高度，实现垂直居中 */
		box-shadow: -0.1rem 0 0 #e3e5e7;
		/* 使用盒子阴影模拟左边框 */

	}

	.vtBtn::after {
		border: none;

	}


	.input-container span {
		position: absolute;
		left: 2.2rem;
		top: 35%;
		color: #212121;
		font-size: 1rem;
	}

	.login-button {
		margin-top: 2.2rem;
		width: 18rem;
		height: 3rem;
		font-size: 35rpx;
		font-weight: bold;
		color:white;
		background-color: #ff9b28;
		
	}

	.text {
		color: #636363;
		font-size: 27rpx;
		margin-left: 4rem;
	}
	.text a{
		color: #f9b732;
	}
</style>