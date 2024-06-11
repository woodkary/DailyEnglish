<template>
	<view>
		<view class="all-container">
			<image class="background" src="../../static/login1.svg"></image>

			<view class="container">
				<span1>Sign In</span1>
				<view class="white-container1">
					<span>账号</span>
					<input id="username" class="search-box" type="text" v-model="username" placeholder="请输入账号">
				</view>
				<view class="white-container2">
					<span>密码</span>
					<view class="password-container">
						<input id="password" class="search-box" type="password" v-model="password" placeholder="请输入密码">
						<img ref="errorIcon" class="error-icon" src="../../static/errorCross.svg">
					</view>
				</view>
				<view class="line">
					<view class="checkbox-container">
						<checkbox id="tmp-28" class="promoted-input-checkbox" />
						<label for="tmp-28">
							自动登录
						</label>
					</view>
					
					<view class="forgot-password-link">忘记密码?</view>
				</view>
				<button class="login-button" @click="login">登录</button>
				
				<span class="text">have no account?
					<router-link to="../register/register">click here</router-link>
				</span>
				
				<span style="color: #636363;margin-left:20px;margin-top:20px">登录代表你同意用户协议、隐私政策和儿童隐私政策</span>
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
				let flag = true;
				// 登录逻辑
				let username = this.username;
				if (!username) {
					this.$nextTick(() => {
						let usernameInput = document.getElementById('username');
						usernameInput.classList.add('inputActive');
						setTimeout(() => {
							usernameInput.classList.remove('inputActive');
						}, 2000);
					});
					flag = false;
				}
				let password = this.password;
				if (!password) {
					this.$nextTick(() => {
						let passwordInput = document.getElementById('password');
						passwordInput.classList.add('inputActive');
						setTimeout(() => {
							passwordInput.classList.remove('inputActive');
						}, 2000);
					});
					flag = false;
				}
				if (!flag) {
					return;
				}
				let remember = this.remember;
				uni.clearStorage()
				uni.request({
					url: 'http://localhost:8080/api/user/login',
					data: {
						username: username,
						password: password,
						remember: remember
					},
					method: 'POST',
					success: (res) => {
						//token失效
						if (res.statusCode === 401) {
							uni.removeStorageSync('token');
							uni.showToast({
								title: '登录已过期，请重新登录',
								icon: 'none',
								duration: 2000
							});
							uni.navigateTo({
								url: '../login/login'
							});
						}
						if (res.statusCode == 200) {
							if (remember) {
								uni.setStorageSync('username');
								uni.setStorageSync('password');
								uni.setStorageSync('remember');
							}
							let token = res.data.token;
							console.log(token);
							uni.setStorageSync('token', token);
							if (res.data.isChoosed) {
								uni.switchTab({
									url: '../home/home'
								});
							} else {
								uni.navigateTo({
									url: '../Welcome/Welcome' +
										`?operation=${res.data.isChoosed ? 1 : 0}`
								});
							}
							uni.setStorageSync("consecutivePunchDay", 11)
						} else if (res.statusCode == 401) { //密码错误
							let passwordInput = document.getElementById('password');
							passwordInput.classList.add('inputActive');
							setTimeout(() => {
								passwordInput.classList.remove('inputActive');
							}, 2000);
							uni.showToast({
								title: '密码错误',
								icon: 'none'
							});
						} else if (res.statusCode == 403) { //用户不存在
							let usernameInput = document.getElementById('username');
							usernameInput.classList.add('inputActive');
							setTimeout(() => {
								usernameInput.classList.remove('inputActive');
							});
							uni.showToast({
								title: '用户不存在',
								icon: 'none'
							}, 1000);
						}
					},
					fail: (res) => {
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
/* 		position: absolute; */
	}

	.background {
		background-color: transparent;
		margin-top: 4rem;
		margin-left: 3rem;
	}

	.container {
		background-color: #ffffff;
		width: 100%;
		margin-top: 3rem;
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

	.line{
		display: flex;
		
	}
	.checkbox-container {
		margin-left: 2.4rem;
		margin-top: 1rem;
		display:block;
		width:40%;
		align-items: center;
		justify-content: center;
	}
	.checkbox-container label {
		color: #838383;
		font-size:1rem;
	}

	.forgot-password-link {
		color: #f57b56;
		text-decoration: none;
		font-size:1rem;
		display: block;
		margin-top: 1rem;
		margin-left: 5rem;
		/* margin-left: 16rem; */
	}

	.password-container {
		display: flex;
	}

	.error-icon {
		position: fixed;
		right: -10%;
		bottom: 17.5%;
		font-size: 16px;
		/* 图标大小 */
		transform: scale(0.15);
		/* 缩放 */
		transition: opacity 0.2s ease-in-out;
		/* 添加过渡效果 */
		opacity: 0;
		/* 初始不显示 */
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

	.text {
		margin-top: 1rem;
		margin-left: 6.4rem;
		color: #636363;
	}

	.text a {
		color: #6b7f73;
	}

	input,
	input::placeholder {
		font-size: 32rpx;
		font-family: Arial, sans-serif;
		padding-left: 32rpx;
	}

	.search-box {
		transition: all 0.3s ease-in-out;
	}

	.inputActive {
		border: 1px solid #e74c3c;
	}
</style>