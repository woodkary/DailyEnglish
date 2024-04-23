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
		data() {
			return {

			}
		},
		methods: {

		}
	}
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