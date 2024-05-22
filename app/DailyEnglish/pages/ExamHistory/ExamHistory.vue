<template>
	<view>
		<view class="today-container">
			<span class="title">今日考试</span>
			<span v-if="exams.length === 0" class="no-exam">今日暂无考试，<navigator>前往复习</navigator></span>
      <view class="todo-exam" v-for="exam in exams" :key="exam.name">
        <view class="row1">
          <text class="exam-name">{{ exam.name }}</text>
          <text class="exam-time"> {{ exam.time }}</text>
          <text class="exam-info">{{ exam.info }}</text>
        </view>
        <view class="row2">
          <button class="todo-btn" @click="reminder(exam)">提醒我</button>
          <button class="todo-btn" @click="takeExam(exam)">去考试</button>
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
      <view class="finished-exam" v-for="exam in finishedExams" :key="exam.date">
        <image class="level" src="@/static/score1.svg"></image>
        <view class="row1">
          <text class="exam-name">{{ exam.name }}</text>
          <text class="exam-date"> {{ exam.date }}</text>
          <text class="exam-info">{{ exam.info }}</text>
        </view>
        <view class="row22">
          <view>
            <span class="score">{{ exam.score }}</span>
            <span style="margin-left: 8rpx;font-size: 24px;">分</span>
          </view>
          <button class="todetail-btn" @click="viewDetails(exam)">考试详情</button>
        </view>
      </view>
		</view>
	</view>
</template>

<script>
	export default {
    data() {
      return{
        exams: [
          {
            exam_id: 1,
            name: '第一单元第一次小测',
            time: '20:00 ~ 21:00',
            info: '共20题'
          },
          {
            exam_id: 2,
            name: '第一单元第一次小测',
            time: '20:00 ~ 21:00',
            info: '共20题'
          },
          // ...更多的考试对象
        ],
        finishedExams: [
          {
            exam_id: 3,
            name: '第一单元第一次小测',
            date: '2023年1月1日',
            info: '共20题',
            score: 95
          },
          {
            exam_id: 4,
            name: '第一单元第二次小测',
            date: '2023年1月1日',
            info: '共20题',
            score: 95
          },
          {
            exam_id: 5,
            name: '第二单元第一次小测',
            date: '2023年1月1日',
            info: '共20题',
            score: 95
          },
          // ...更多的考试结果对象
        ]
      }

    },
		methods: {
      getPreviousExams() {
        // 从服务器获取上一次考试记录
        uni.request({
          url: '/api/exams/previous_examinations',
          success: (res) => {
            this.exams = this.transformExams(res.data.exams);
          }
        });
      },

      viewDetails(exam){
        // 跳转到考试详情页面
        uni.navigateTo({
          url: `../exam_details/exam_details?exam_id=${exam.exam_id}&exam_name=${exam.name}`
        });
      },

      transformExams(exams) {
        return exams.map(exam => {
          // 将日期从 "yyyy/mm/dd" 转换为 "年月日"
          const dateParts = exam.exam_date.split('/');
          const dateInChineseFormat = `${dateParts[0]}年${dateParts[1]}月${dateParts[2]}日`;

          return {
            exam_id: exam.exam_id,
            name: exam.exam_name,
            date: dateInChineseFormat,
            info: `共${exam.question_num}题`,
            score: exam.exam_score
          };
        });
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