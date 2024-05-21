<template>
	<view>
		<view class="today-container">
			<span class="title">今日考试</span>
			<span class="no-exam">今日暂无考试，<navigator>前往复习</navigator></span>
			<view class="todo-exam">
				<view class="row1">
					<text class="exam-name">第一单元第一次小测</text>
					<text class="exam-time"> 20：00 ~ 21：00</text>
					<text class="exam-info">共20题</text>
				</view>
				<view class="row2">
					<button class="todo-btn">提醒我</button>
					<button class="todo-btn">去考试</button>
				</view>
			</view>
			<view class="todo-exam">
				<view class="row1">
					<text class="exam-name">第一单元第一次小测</text>
					<text class="exam-time"> 20：00 ~ 21：00</text>
					<text class="exam-info">共20题</text>
				</view>
				<view class="row2">
					<button class="todo-btn">提醒我</button>
					<button class="todo-btn">去考试</button>
				</view>
			</view>
		</view>
		<view class="history-container">
			<view class="_row1">
				<span class="title">所有考试</span>
				<input />
			</view>
			<view class="_row2">
				<button class="choice selected">时间顺序</button>
				<button class="choice">时间逆序</button>
				<button class="choice">成绩顺序</button>
				<button class="choice">成绩逆序</button>
			</view>
			<view class="finished-exam">
				<image class="level" src="@/static/score1.svg"></image>
				<view class="row1">
					<text class="exam-name">第一单元第一次小测</text>
					<text class="exam-date"> 2023年1月1日</text>
					<text class="exam-info">共20题</text>
				</view>
				<view class="row22">
					<view>
						<span class="score">95</span>
						<span style="margin-left: 8rpx;font-size: 24px;">分</span>
					</view>
					<button class="todetail-btn">考试详情</button>
				</view>
			</view>
			<view class="finished-exam">
				<view class="row1">
					<text class="exam-name">第一单元第一次小测</text>
					<text class="exam-date"> 2023年1月1日</text>
					<text class="exam-info">共20题</text>
				</view>
				<view class="row22">
					<view>
						<span class="score">95</span>
						<span style="margin-left: 8rpx;font-size: 24px;">分</span>
					</view>
					<button class="todetail-btn">考试详情</button>
				</view>
			</view>
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
	@font-face {
		font-family: "pingfang";
		src: url('@/static/PingFang Medium_downcc.otf');
	}

	body {
		height: 100vh;
		width: 100%;
		background-color: #fbfbfb;
	}

	.today-container {
		width: 100%;
		display: flex;
		flex-direction: column;
		position: relative;
	}

	.title {
		margin-top: 20rpx;
		margin-left: 4%;
		font-size: 22px;
		font-weight: bold;
		font-family: 'pingfang';
	}

	.history-container {
		margin-top: 40rpx;
		width: 100%;
		display: flex;
		flex-direction: column;
		position: relative;
	}

	.no-exam {
		margin-left: 40rpx;
		margin-top: 30rpx;
		display: flex;

		navigator {
			color: blue;
		}
	}

	.todo-exam {
		display: flex;
		/* 开启flexbox */
		width: 92%;
		border-radius: 12px;
		background-color: white;
		margin-left: 4%;
		margin-top: 20rpx;
		height: 200rpx;
		box-shadow: rgba(149, 157, 165, 0.2) 0px 8px 24px;
	}

	.row1,
	.row2 {
		display: flex;
		/* 开启flexbox */
		flex-direction: column;
		/* 子元素纵向排列 */
	}

	.row1 {
		text-align: left;
		/* 文本向左对齐 */
		width: 60%;
	}

	.row2 {
		width: 40%;
	}

	.exam-name {
		margin-top: 18rpx;
		margin-left: 34rpx;
		font-size: 19px;
		font-family: 'pingfang';
	}

	.exam-time,
	.exam-date {
		margin-top: 10rpx;
		margin-left: 34rpx;
		font-size: 19px;
		color: #456DE7;
		font-weight: 550;
		font-family: 'pingfang';
	}

	.exam-info {
		margin-top: 12rpx;
		margin-left: 34rpx;
		font-size: 15px;
		font-weight: 550;
		color: #A7A7A7;
		font-family: 'pingfang';
	}

	.todo-btn {
		background-color: #5f89fc;
		color: white;
		width: 240rpx;
		height: 60rpx;
		margin-top: 20rpx;
		line-height: 60rpx;
		margin-bottom: 13rpx;
		margin-right: 40rpx;
		border-radius: 12px;
	}

	._row1 {
		width: 100%;
		display: flex;
	}

	._row2 {
		width: 100%;
		display: flex;
		margin-top: 20rpx;

	}

	.choice {
		width: 130rpx;
		height: 40rpx;
		line-height: 40rpx;
		text-align: center;
		padding: 0;
		font-size: 13px;
		background-color: #F0F0F0;
		color: #A7A7A7;
		margin-left: 5rpx;
		/* 减少左侧外边距 */
		margin-right: 5rpx;
		/* 减少右侧外边距 */
	}

	._row2>.choice:first-child {
		margin-left: 25rpx;
	}

	.selected {
		background-color: #e2e7f9;
		color: #456DE7;
	}

	.finished-exam {
		position: relative;
		display: flex;
		width: 92%;
		border-radius: 12px;
		background-color: white;
		margin-left: 4%;
		margin-top: 40rpx;
		height: 200rpx;
		box-shadow: rgba(149, 157, 165, 0.2) 0px 8px 24px;
	}

	.level {
		position: absolute;
		right: 0;
		top:-15px;
		width: 200rpx;
		height: 200rpx;
		z-index: 0;
		opacity: 0.7;
	}

	.score {
		font-size: 46px;
		font-family: 'pingfang';
		margin-left: 20rpx;
		
	}

	.todetail-btn {
		background-color: #5f89fc;
		color: white;
		width: 240rpx;
		height: 60rpx;
		margin-top: 2rpx;
		border-radius: 10px;
		line-height: 60rpx;
		margin-bottom: 13rpx;
		margin-right: 40rpx;
		z-index: 0;
	}
</style>