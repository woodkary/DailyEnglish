<template>
	<view>
		
		<view class="title-container">
			<image class="back-icon" src="../../static/back.svg"></image>
			<span>考试</span>
		</view>
		<view class="exam-info">
			<h2>第一单元第一次小测</h2>
			<h3>20：00 ~ 21：00</h3>
			<view class="circle">
				<span class="exam-time">考试时间60分钟</span>
				<span class="exam-num">共二十题</span>
			</view>
			<view class="start-exam-btn">去考试</view>
		</view>
		
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
	body {
		background-image: url(@/static/daka-background.png);
		background-size: cover;
		
	}
	.title-container {
		display: flex;
		line-height: 80rpx;
	}

	.back-icon {
		width: 60rpx;
		margin-right: 80rpx;
		height: 80rpx;
		margin-top: 20rpx;
		margin-bottom: 20rpx;
	}

	.title-container span {
		font-family: 'Source Han Sans', 'Microsoft YaHei', Arial, sans-serif;
		font-weight: bold;
		font-size: 18px;
		margin-top: 20rpx;
		margin-left: 180rpx;
	}

	.exam-info {
		text-align: center;
		/* 将文本水平居中 */
	}

	.exam-info h2 {
		margin-top: 100rpx;
	}

	.exam-info h3 {
		margin-top: 30rpx;
		color: #456DE7;
	}

	.circle {
		position: relative;
		display: flex;
		/* 使用 Flexbox 布局 */
		align-items: center;
		/* 垂直居中 */
		justify-content: center;
		/* 水平居中 */
		flex-direction: column;
		/* 垂直排列 */
		padding: 20px;
		width: 200px;
		/* 设置相等的宽高 */
		height: 200px;
		border-radius: 50%;
		border: 35px solid rgba(213, 251, 252, 0.7);
		/* 圆形边框，透明度为 0.7 */
		color: #000;
		/* 文本颜色 */
		margin: 20rpx auto; /* 水平居中 */
	}

	.exam-time,
	.exam-num {
		text-align: center;
		font-size: 24px;
	}
	.exam-num{
		color: #b0b0b0;
	}
	.start-exam-btn {
			background-color: #0276ff; /* 蓝色背景 */
			color: #fff; /* 白色文字 */
			padding: 10px 20px; /* 添加适当的内边距 */
			border-radius: 5px; /* 圆角 */ 
			margin: 50px 40px; /* 与 .circle 之间的间距 */
			text-align: center; /* 文本居中 */
			cursor: pointer; /* 鼠标指针样式 */
			font-size: 26px;
			 box-shadow: 0 2px 10px rgba(0, 0, 0, 0.16), 0 3px 6px rgba(0, 0, 0, 0.1);
		}
</style>