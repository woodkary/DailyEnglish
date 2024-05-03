<template>
	<view class="body">
		<view>
			<view class="calendar">
				<view class="head">
					<button class="last">
						<!-- 					<image src="../../static/last.png"></image> -->
					</button>
					<text class="date">2024年5月</text>
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
      requestExamMsg(){

      },
      handleClick(date) {
        if(date.hasExam){
          console.log('考试日期：', date.value);
          console.log('考试分数：', this.examMsg[date.value]);
          //TODO 跳转到考试页面
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
		background-color: rgb(245, 197, 127);
		height: 100vh;
		z-index: -1;
	}

	.calendar {
		position: absolute;
		width: 100%;
		height: 840rpx;
		top: 40rpx;
		background-color: transparent;
		border: 1px solid #000;
		;
	}

	.head {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-top: 20rpx;
		margin-bottom: 40rpx;
		font-size: 60rpx;
		font-weight: bold;
	}

	.week {
		display: flex;
		justify-content: space-around;
		text-align: center;
		margin-bottom: 40rpx;
		font-size: 45rpx;
		font-weight: 550;
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
	}

	.date-item {
		width: 14.28%;
		/* 每个日期占据日历宽度的1/7 */
		text-align: center;
		margin-bottom: 55rpx;
		/* 设置日期元素之间的垂直间距 */
		font-size: 40rpx;
	}

	.clickable {
		color: red;
	}

	.sunday {
		color: #aa916e;
		/* 星期日红色字体 */
	}

	.saturday {
		color: #aa916e;
		/* 星期六蓝色字体 */
	}
</style>