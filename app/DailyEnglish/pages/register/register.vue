<template>
	<view>
		<image class="background" src="../../static/background.png"></image>
		<image class="color" src="../../static/color.png"></image>
		<view class="logo-container">
			<image class="r_logo1" src="../../static/r_logo1.png"></image>
			<image class="r_logo2" src="../../static/r_logo2.png"></image>
			<image class="r_logo3" src="../../static/r_logo3.png"></image>
		</view>
		<view class="white-container" :style="{ marginTop: '0rem' }" id="1">
			<input class="search-box" type="text" placeholder="请输入账号">
		</view>
		<view class="white-container" :style="{ marginTop: '4rem' }" id="2">
			<input class="search-box" type="text" placeholder="请输入邮箱">
		</view>
		<view class="white-container" :style="{ marginTop: '8rem' , width:'10rem'}" id="3">
			<input class="search-box" type="text" placeholder="请输入验证码">
		</view>
		<button class="verifi-button">获取验证码</button>
		<view class="white-container" :style="{ marginTop: '12rem' }" id="4">
			<input class="search-box" type="text" placeholder="请输入密码">
		</view>
		<view class="white-container" :style="{ marginTop: '16rem' }" id="5">
			<input class="search-box" type="text" placeholder="再次输入密码">
		</view>
		<button class="login-button">注册</button>
		<button class="button1"></button>
		<span class="text">已有账号？点此登录</span>
	</view>
</template>

<script>
	export default {
<<<<<<< Updated upstream
		data() {
			return {

			}
		},
		methods: {

		}
	}
=======
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
          url: '/api/register/sendCode',
          data: {
            email: this.email
          },
          withCredentials: false,
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
            }else if(res.statusCode==409) {//邮箱已注册
              uni.showToast({
                title: '邮箱已注册',
                icon: 'error'
              });
            }else if(res.statusCode==400){//请求参数错误
              uni.showToast({
                title: '请求参数错误',
                icon: 'error'
              });
            }else{
              uni.showToast({
                title: '发送失败',
                icon: 'error'
              });
            }
          },
          fail: (res) => {
            console.log(res);
            uni.showToast({
              title: '发送失败',
              icon: 'error'
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
            icon: 'error'
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
          url: '/api/users/register',
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
                icon: 'success'
              });
              setTimeout(() => {
                uni.navigateBack();
              }, 1000);
            }else if(data.code==409) {//用户名已注册
              uni.showToast({
                title: '用户名已注册',
                icon: 'error'
              });
            }else if(data.code==400){//请求参数错误
              uni.showToast({
                title: '请求参数错误',
                icon: 'error'
              });
            }else{
              uni.showToast({
                title: '注册失败',
                icon: 'error'
              });
            }
          },
          fail: (res) => {
            console.log(res);
            uni.showToast({
              title: '注册失败',
              icon: 'error'
            });
          }
        });
      }
    }
  }
>>>>>>> Stashed changes
</script>

<style>
	.background {
		width: 100vw;
		height: 100vh;
		/* 高度等于视口高度 */
		position: absolute;
	}

	.color {
		width: 100vw;
		height: 100vh;
		/* 高度等于视口高度 */
		position: absolute;
		/* 绝对定位以覆盖整个容器 */
		top: 0;
		left: 0;
		z-index: 1;
		/* 设置较低的z-index值 */
		opacity: 0.9;
	}

	.logo-container {
		position: absolute;
		top: 266rpx;
		/* 您希望的对齐位置 */
		left: 74rpx;
		display: flex;
		flex-direction: column;
		align-items: flex-start;
		/* 如果您希望图标左对齐 */
	}

	.r_logo1 {
		left: 0rpx;
		top: 0rpx;
		width: 400rpx;
		height: 80rpx;
		z-index: 2;
	}

	.r_logo2 {
		left: 0rpx;
		top: 70rpx;
		width: 350rpx;
		height: 68rpx;
		z-index: 2;
	}

	.r_logo3 {
		left: 440rpx;
		top: -48rpx;
		width: 146rpx;
		height: 112rpx;
		z-index: 2;
	}

	.white-container {
		width: 19.5rem;
		/* 设置容器的宽度 */
		height: 2rem;
		/* 设置容器的高度，根据需要调整 */
		background-color: white;
		/* 设置背景颜色为白色 */
		border: 1px solid black;
		/* 设置边框为1像素黑色实线 */
		position: absolute;
		/* 绝对定位以覆盖在图片之上 */
		top: 18rem;
		/* 顶部距离视口50% */
		left: 2rem;

		/* 左边距离视口50% */
		z-index: 2;
		/* 设置z-index值，确保在color之上 */
		display: flex;
		/* 设置为弹性布局 */
		align-items: center;
		/* 垂直居中 */
		justify-content: center;
		/* 水平居中 */
		border-radius: 10px;
		/* 设置边框圆角 */
		opacity: 0.65;
	}

	.search-box {
		width: 100%;
		/* 设置搜索框宽度为容器的80% */
		height: 2rem;
		/* 设置搜索框高度 */
		border-radius: 10px;
		/* 设置边框圆角 */
		border: none;
		/* 设置边框颜色 */
		font-size: 0.8rem;
		/* 设置字体大小 */
		outline: none;
		/* 移除默认轮廓 */
		margin-left: 1rem;
	}

	.login-button {
		background-color: #75C2FD;
		/* 设置背景颜色为白色 */
		border: none;
		/* 设置边框为1像素黑色实线 */
		color: white;
		position: absolute;
		/* 绝对定位以覆盖在图片之上 */
		top: 38rem;
		left: 3rem;
		width: 18rem;
		height: 3rem;
		font-size: 35rpx;
		z-index: 2;
		/* 设置z-index值，确保在color之上 */
		cursor: pointer;
		border-radius: 20rpx;
	}

	.text {
		position: absolute;
		z-index: 2;
		color: white;
		top: 42.3rem;
		left: 9rem;
		font-size: 20rpx;
	}
	.verifi-button{
		background-color: blue;
		/* 设置背景颜色为白色 */
		border: none;
		/* 设置边框为1像素黑色实线 */
		color: white;
		position: absolute;
		/* 绝对定位以覆盖在图片之上 */
		top: 26rem;
		left: 14rem;
		height: 2rem;
		background-color: blue;
		text-align: center;
		color: white;
		border: none;
		border-radius: 5px;
		cursor: pointer;
		vertical-align: middle;
		z-index: 2;
		/* 设置z-index值，确保在color之上 */
		cursor: pointer;
		border-radius: 20rpx;
		line-height: 2rem;
	}
</style>