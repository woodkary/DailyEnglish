<template>
	<view>
		<view class="background">
			<span>Go for it！</span>
			<span style="color: grey;font-weight: lighter;margin-top: 20rpx;font-size: 20px;">今日已学<span
					class="study-num">{{ todayLearned }}</span>个单词</span>
			<image src="../../static/sanyueqi.png" class="sanyueqi"></image>
		</view>
		<view class="center-container">
			<view class="border1">
				<span style="font-size: 20px;margin-left: -160px;">你已连续学习</span><br>
				<span class="study-day">{{ consecutivePunchDay }}</span>
				<span style="color: #6200EE;font-size: 20px;margin-left:-10px;">天</span>
				<image src="../../static/flash.png" class="flash"></image>
				<view class="btn">查看我的足迹</view>
			</view>
		</view>
		<h3 v-show="haveExam" style="margin-left: 40px;margin-top:20px;">你今天有<span>{{examCnt}}</span>场考试</h3>
		<view class="center-container2" v-show="haveExam">
		    <view class="border2" v-for="exam in exams" :key="exam.id">
		      <view class="exam-info">
		        <h3>{{ exam.title }}</h3> <!-- 使用考试标题 -->
		        <h4>开始时间：{{ exam.time }}</h4> <!-- 使用考试时间 -->
		        <view class="small-btn" @click="toExam">去考试</view>
		      </view>
		      <image src="../../static/gotoexam.png" style="width: 150px;height: 150px;"></image>
		    </view>
		  </view>
		<view class="btn" style="background-color: #456de7;font-size: 26px;">完成</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
        // 今天学了多少个单词
        todayLearned: 10,
        // 今天是否有考试
				haveExam: true,
        // 连续学习天数
        consecutivePunchDay: 10,
        // 今天考试数量
				examCnt: 2,
        // 考试列表
				exams: [
				        { id: 1, title: '第一单元第一次小测', time: '20:00 ' },
				        { id: 2, title: '第二单元第一次小测', time: '21:00 ' }
				      ],
					  
			}
		},
		methods: {
			toExam() {
				this.$router.push('/exam')
			}
		}
	}
</script>

<style>
	@font-face {
		font-family: "pingfang";
		src: url('@/static/PingFang Medium_downcc.otf');
	}

	body {
		font-family: "pingfang";
	}

	.background {
		background-color: rgb(121, 245, 255, 0.33);
		display: flex;
		/* 使用 Flexbox 布局 */
		justify-content: center;
		/* 水平居中 */
		align-items: center;
		/* 垂直居中 */
		flex-direction: column;
		/* 垂直布局 */
		border: 0 solid #000;
		/* 添加边框，2px宽度，黑色 */
		border-radius: 0 0 30px 30px;
		/* 左下、右下圆角30度 */
		padding: 10px;
		/* 添加内边距 */
		z-index: 1;
	}

	.background span {
		margin-top: 60rpx;
		font-size: 24px;
		font-weight: bold;
	}

	.study-num {
		color: #000;
		font-weight: bold;
	}

	.sanyueqi {
		width: 180px;
		height: 180px;
	}

	.center-container {
		display: flex;
		justify-content: center;
		align-items: center;
		margin-top: -50px;
		z-index: 2;
	}

	.border1 {
		width: 80%;
		text-align: center;
		background: #ffffff;
		/* 文本居中 */
		border: 2px solid #ccc;
		/* 添加边框，2px宽度，灰色 */
		box-shadow: rgba(0, 0, 0, 0.1) 0px 5px 12px 0px;
		padding: 20px;
		/* 添加内边距 */
		border-radius: 18px;
		z-index: 2;
	}

	.study-day {
		font-size: 42px;
		margin-left: -140px;
		margin-right: 20px;
		color: #6200EE;
		font-weight:1000;
	}

	.flash {
		position: absolute;
		right: 60px;
		top:280px;
		width: 90px;
		height: 90px;
	}

	.center-container2 {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		margin-top: 20px;
		row-gap: 20px;
		z-index: 2;
	}

	.border2 {
		display: flex;
		/* 添加 Flex 布局 */
		justify-content: space-between;
		/* 左右排列 */
		align-items: center;
		/* 垂直居中 */
		width: 78%;
		text-align: center;
		background: #ffffff;
		/* 文本居中 */
		border: 2px solid #ccc;
		/* 添加边框，2px宽度，灰色 */
		box-shadow: rgba(0, 0, 0, 0.1) 0px 5px 12px 0px;
		/* 添加阴影，水平和垂直偏移均为2px，模糊半径为5px，透明度为0.2 */
		padding-left:20px;
		padding-right:20px;
		/* 添加内边距 */
		border-radius: 18px;
		z-index: 2;
	}

	.exam-info {
		text-align: center;
		/* 将文本水平居中 */
	}

	.exam-info h3 {
		margin-top:0rpx;
		font-size: 20px;
	}

	.exam-info h4 {
		margin-top: 20rpx;
		color: #456DE7;
	}

	.btn {
		background-color: #6200ee;
		/* 蓝色背景 */
		color: #fff;
		/* 白色文字 */
		padding: 10px 5px;
		/* 添加适当的内边距 */
		border-radius: 12px;
		/* 圆角 */
		margin-top: 20px;
		/* 与 .circle 之间的间距 */
		text-align: center;
		/* 文本居中 */
		cursor: pointer;
		/* 鼠标指针样式 */
		font-size: 24px;
		width: 80%;
		margin-left: 10%;
		height: 30px;
		line-height:30px;
	}

	.small-btn {
		background-color: #456DE7;
		/* 蓝色背景 */
		color: #fff;
		/* 白色文字 */
		padding: 5px 2.5px;
		/* 添加适当的内边距 */
		border-radius: 5px;
		/* 圆角 */
		margin-top:10px;
		/* 与 .circle 之间的间距 */
		text-align: center;
		/* 文本居中 */
		cursor: pointer;
		/* 鼠标指针样式 */
		font-size: 20px;
	}
</style>