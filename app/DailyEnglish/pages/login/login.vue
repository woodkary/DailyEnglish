<template>
	<view>
		<view class="all-container">
			<image class="background" src="../../static/login1.svg"></image>

			<view class="container">
				<span1>Sign In</span1>
				<view class="white-container1">
					<span>账号</span>
					<input class="search-box" type="text" v-model="username" placeholder="请输入账号">
				</view>
				<view class="white-container2">
					<span>密码</span>
					<input class="search-box" type="password" v-model="password" placeholder="请输入密码">
					<a href="forgot-password.html" class="forgot-password-link">忘记密码?</a>

				</view>
				<button class="login-button" @click="login">登录</button>
				<!-- 	<button class="register-button">注册</button><button class="forget">忘记密码？</button>
				<button class="button1"></button>
				<span class="text">登录代表你同意用户协议、隐私政策和儿童隐私政策</span> -->
				<span class="text">have no account?<a>click here</a></span>
			</view>
		</view>
	</view>

</template>

<script>
	export default {
		data() {
			return {
				username: '',
				password: '',
				remember: false
			}
		},
		beforeMount() {
			//获取本地存储的用户名和密码
			let username = uni.getStorageSync('username');
			let password = uni.getStorageSync('password');
			let remember = uni.getStorageSync('remember');
			if (username && password && remember) {
				this.username = username;
				this.password = password;
				this.remember = remember;
			}
		},
		methods: {
			autoLogin() {
				this.remember = !this.remember;
				console.log(this.remember);
				console.log(this.username);
				console.log(this.password);
			},
			login() {
				// 登录逻辑
				let username = this.username;
				let password = this.password;
				let remember = this.remember;
				uni.request({
					url: 'http://localhost:8080/api/users/login',
					data: {
						username: username,
						password: password,
						remember: remember
					},
					method: 'POST',
					success: (res) => {
						if (remember) {
							let token = res.data.token;
							uni.setStorageSync('username');
							uni.setStorageSync('password');
							uni.setStorageSync('remember');
							uni.setStorageSync('token', token);
						}

						uni.navigateTo({
							//TODO: 跳转到首页，或处理其他逻辑
							url: '/pages/index/index'
						});
					},
					fail: (res) => {
						//TODO: 处理登录失败逻辑
						uni.showToast({
							title: '登录失败',
							icon: 'none'
						});
					}
				});

			}
		}
	}
</script>

<style>
	.all-container {
		height: 100%;
		width: 100%;
		background-color: #fed8c3;
		position: absolute;
	}

	.background {
		background-color: transparent;
		margin-top: 3rem;
		margin-left: 3rem;
	}

	.container {
		background-color: #ffffff;
		width: 100%;
		margin-top: 2rem;
		height: 60%;
		/*上边圆角*/
		border-top-left-radius: 2rem;
		border-top-right-radius: 2rem;
		display: flex;
		flex-direction: column;
	}

	.container span1 {
		margin-top: 5%;
		font-size: 2.7rem;
		margin-left: 8%;
		font-weight: 500;
	}

	@font-face {
		font-family: '仓耳渔阳体';
		src: url('../../static/TsangerYuYangT_W05_W05.ttf') format('truetype');
		/* 如果有其他格式，也可以添加其他src */
	}

	.white-container1 {
		margin-top: 2rem;
	}

	.white-container1 span {
		margin-left: 2.4rem;
		color: #838383;
	}

	.white-container1 input {
		width: 86%;
		height: 3.3rem;
		border-radius: 4rem;
		background-color: #f0f3f1;
		margin-left: 1.4rem;
		margin-top: 0.4rem;
	}

	.white-container1 input:hover {
		background-color: #eff0ef;
	}

	.white-container2 {
		margin-top: 1rem;
	}

	.white-container2 span {
		margin-left: 2.4rem;
		color: #838383;
	}

	.white-container2 input {
		width: 86%;
		height: 3.3rem;
		border-radius: 4rem;
		background-color: #f0f3f1;
		margin-left: 1.4rem;
		margin-top: 0.4rem;
	}

	.white-container2 input:hover {
		background-color: #eff0ef;
	}
	.forgot-password-link {
	    color: #f57b56; /* 蓝色字体 */
	    text-decoration: none; /* 去除下划线 */
		font-size: 1rem;
	    margin-left: 16rem; /* 添加一些左边距 */
	}
	

	.login-button {
		background-color: #44564a;
		margin-top: 2.7rem;
		color: white;
		width: 76%;
		height: 3.3rem;
		border-radius: 4rem;
		text-align: center;
		/*文本垂直居中 */
		line-height: 3.3rem;
	}
	.text{
		margin-top: 1rem;
		margin-left: 6.4rem;
		color: #636363;
	}
	.text a{
		color: #6b7f73;
	}
</style>