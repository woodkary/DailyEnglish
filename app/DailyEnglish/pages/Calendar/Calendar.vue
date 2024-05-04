<template>
	<view class="body">
		<view class="color">
		</view>
		<view>
			<view class="calendar">
				<view class="head">
					<text class="date">2024年5月</text>
					<button class="last">
						<!-- 					<image src="../../static/last.png"></image> -->
					</button>
					<button class="next">
						<!-- 					<image src="../../static/next.png"></image> -->
					</button>
				</view>
				<view class="week">
					<text class="week1">日</text>
					<text class="week2">一</text>
					<text class="week3">二</text>
					<text class="week4">三</text>
					<text class="week5">四</text>
					<text class="week6">五</text>
					<text class="week7">六</text>
				</view>
				<view class="day">
					<view class="date-item" v-for="(date, index) in dates" :key="index" :class="{
                      'clickable': date.hasExam,
				              'sunday': date.dayOfWeek === 0,
				              'saturday': date.dayOfWeek === 6
				            }" @click="handleClick(date)">
						{{ date.value }}
					</view>
				</view>
			</view>
		</view>
		<view>
			<view class="examMsg">
				<text class="title">5月4日</text>
				<view class="card-container">
					<view class="card" id="daka">
						<image src="../../static/done.svg"></image>
						<text class="title">打卡计划:</text>
						<text class="state">未完成</text>
					</view>
					<view class="card" id="daka">
						<image src="../../static/done.svg"></image>
						<text class="title">打卡计划:</text>
						<text class="state">已完成</text>
					</view>
					<!-- <view class="card" id="exam">
					<image src="../../static/todo.svg"></image>
					<text class="time">9:40</text>
					<text class="course">语文</text>
					<text class="score">得分：90</text>
				</view> -->
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				year: 2024,
				month: 5,
				dates: [], // 存储当前月份的日期
				// 考试信息
				examMsg: new Set([
					'2024-5-1',
					'2024-5-13',
					'2024-5-22',
					'2024-5-29'
				]) // 存储当前月份的考试日期
			}
		},

		beforeMount() {
			this.generateDates();
		},
		methods: {
			//TODO 发送网络请求获取考试信息
			requestExamMsg() {

			},
			handleClick(date) {
				if (date.hasExam) {
					console.log('考试日期：', date.value);
					//TODO 跳转到考试页面及其他操作
				}

			},
			generateDates() {
				const firstDay = new Date(this.year, this.month - 1, 1); // 获取当前月份的第一天
				const firstDayOfWeek = firstDay.getDay(); // 获取当前月份的第一天是星期几
				const totalDays = new Date(this.year, this.month, 0).getDate(); // 获取当前月份的总天数

				// 初始化日期数组
				this.dates = [];

				// 添加空白日期（用于填充第一天之前的空白）
				for (let i = 0; i < firstDayOfWeek; i++) {
					this.dates.push({
						value: '',
						dayOfWeek: '',
						hasExam: false

					});
				}
				// 添加日期
				for (let i = 1; i <= totalDays; i++) {
					const dayOfWeek = (firstDayOfWeek + i - 1) % 7; // 计算当前日期对应的星期几（0 表示星期日，1 表示星期一，以此类推）
					const dateStr = `${this.year}-${this.month}-${i}`;
					this.dates.push({
						value: i,
						dayOfWeek: dayOfWeek,
						hasExam: this.examMsg.has(dateStr) // 判断当前日期是否有考试
					});
				}
			}
		},
	}
</script>

<style>
	.body {
		height: 100vh;
		z-index: -1;
	}

	.color {
		background-color: #144de4;
		position: absolute;
		top: 0;
		width: 100%;
		height: 500rpx;
		z-index: -1;
	}

	.calendar {
		position: absolute;
		width: 90%;
		height: 750rpx;
		top: 90rpx;
		margin-left: 5%;
		background-color: transparent;
		z-index: 1;

		box-shadow: 5px 5px 25px rgb(0, 0, 0, 0.1);
	}

	.head {
		display: flex;
		margin-top: 30rpx;
		margin-bottom: 40rpx;
		font-size: 60rpx;
		font-weight: 500;
		color: #fff;
	}

	.week {
		display: flex;
		justify-content: space-around;
		text-align: center;
		/* margin-bottom: 30rpx; */
		font-size: 45rpx;
		font-weight: 550;
		background-color: white;
		border-top-left-radius: 10rpx;
		border-top-right-radius: 10rpx;
	}

	.week {

		.week1,
		.week7 {
			color: #aa916e;
		}

		.week2,
		.week3,
		.week4,
		.week5,
		.week6 {
			color: #b58b4b;
		}
	}

	.day {
		display: flex;
		flex-wrap: wrap;
		background-color: white;
	}

	.date-item {
		/* 设置盒模型为 border-box，这样边框和内边距都不会影响元素的最终宽度 */
		box-sizing: border-box;
		width: 96.4rpx;
		/* 每个日期占据日历宽度的1/7 */
		text-align: center;
		/* 设置日期元素之间的垂直间距 */
		font-size: 40rpx;
		height: 96.4rpx;
		margin-bottom: 10rpx;
		line-height: 96.4rpx;
		color: #919597;

		position: relative;
	}

	.clickable {
		color: black;
		/* border: 1px solid red; */
		/* border-radius: 50%; */
		z-index: 1;
	}

	.clickable::before {
		content: '';
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		width: 90%;
		/* 调整圆圈的大小 */
		height: 90%;
		border-radius: 50%;
		background-color: #f1f4fb;
		/* 设置圆圈的颜色 */
		z-index: -1;
		/* 将圆圈置于日期数字下方 */
		/* opacity: 0.5; */
		/* 设置透明度，以便可以看到下方的数字 */

	}

	.examMsg {
		background-color: #fff;
		position: absolute;
		top: 880rpx;
		/* 初始位置 */
		width: 100%;
		height: calc(100vh - 900rpx);
		/* 设置高度 */
		justify-content: space-between;
		/* 使图片和文本之间有空间 */
		display: flex;
		flex-direction: column;
	}

	.examMsg .title {
		margin-left: 30rpx;
		margin-top: 40rpx;
		font-size: 40rpx;
		font-weight: 550;
	}

	.examMsg .card-container {
		display: flex;
		flex-direction: column;
		position: absolute;
		top: 130rpx;
		width: 100%;
	}

	.examMsg .card {
		display: flex;
		width: 90%;
		/* border: 0.1px solid #d7d7d7; */
		margin-bottom: 40rpx;
		margin-left: 5%;
		height: 140rpx;
		/* 		  box-shadow: 0 4px 4px rgba(192, 192, 192, 0.2), 4px 0 4px rgba(191, 191, 191, 0.2); */
		/* box-shadow: 10px 10px 25px rgb(0, 0, 0, 0.1); */
	}

	.examMsg .card image {
		width: 70rpx;
		height: 70rpx;
		margin-left: 60rpx;
		margin-top: 40rpx;
	}

	#daka {
		background-color: #f1f4fb;
	}

	#daka .title {
		margin-left: 40rpx;
		margin-top: 40rpx;
		font-size: 35rpx;
		font-weight: 500;
	}

	#daka .state {
		margin-left: 40rpx;
		margin-top: 40rpx;
		font-size:35rpx;
	}
</style>