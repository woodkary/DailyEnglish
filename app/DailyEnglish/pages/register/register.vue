<template>
	<view>

		<view class="background">
			<span class="span1"><span class="sign">Sign</span><br>Up</span>
			<image class="pic" src="../../static/register1.svg"></image>
		</view>
		<view class="input-container" id="1">
			<span>昵称</span>
			<input class="input" type="text" placeholder="请输入昵称">
		</view>
		<view class="input-container" id="2">
			<span>邮箱</span>
			<input class="input" type="text" placeholder="请输入邮箱">
			<button class="vtBtn">发送验证码</button>
		</view>

		<view class="input-container" id="3">
			<span style="left:2rem;">验证码</span>
			<input class="input" type="text" placeholder="请输入验证码">
		</view>

		<view class="input-container" id="4">
			<span>密码</span>
			<input class="input" type="password" placeholder="请输入密码">
		</view>
		<view class="input-container" id="5">
			<input class="input" type="password" placeholder="再次输入密码">
		</view>
		<button class="login-button">注册</button>
		<span class="text">Already have  account?<a>click here to login</a></span>
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