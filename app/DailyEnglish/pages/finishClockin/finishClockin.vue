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
		    <view class="border2" v-for="exam in exams" :key="exam.exam_id">
		      <view class="exam-info">
		        <h3>{{ exam.name }}</h3> <!-- 使用考试标题 -->
		        <h4>开始时间：{{ exam.start_time }}</h4> <!-- 使用考试时间 -->
		        <view class="small-btn" @click="toExam(exam)">去考试</view>
		      </view>
		      <image src="../../static/gotoexam.png" style="width: 150px;height: 150px;"></image>
		    </view>
		  </view>
		<view class="btn" style="background-color: #456de7;font-size: 26px;" @click="toHome">完成</view>
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
        operation:0,
        // 考试列表
				exams: [
				        /*{
                  exam_id: 1,
                  name: '第一单元第一次小测',
                  start_time: '20:00',
                  duration: 60,
                  info:"共10题",
                  question_num: 10
                },
				        {
                  exam_id: 2,
                  name: '第二单元第一次小测',
                  start_time: '21:00',
                  duration: 60,
                  info:"共10题",
                  question_num: 10
                }*/
				      ],
					  
			}
		},
    onLoad(event) {
      //获取题目数量
      this.todayLearned = event.questionNum;
      this.operation=parseInt(event.operation);
      // 从本地缓存中获取今天的学习天数
      const consecutivePunchDay = uni.getStorageSync("consecutivePunchDay");
      if (consecutivePunchDay) {
        this.consecutivePunchDay = consecutivePunchDay;
      }else{
        //发送请求获取今天的学习天数
        uni.request({
          url: 'http://localhost:8080/api/users/my_punches',
          method: 'GET',
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          },
          success: (res) => {
            //token失效
            if(res.statusCode === 401){
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
            const data = res.data;
            if (data.code === 0) {
              this.consecutivePunchDay = data.consecutive_punch_day;
              uni.setStorageSync("consecutivePunchDay", data.consecutive_punch_day);
            }
          },
          fail: (err) => {
            console.log(err);
          }
        });
      }
      // 从服务器获取今天的考试记录
      uni.request({
        url: 'http://localhost:8080/api/exams/exams_date',
        method: 'POST',
        header: {
          'Authorization': `Bearer ${uni.getStorageSync('token')}`
        },
        data: {
          date: this.getExamDate(new Date())
        },
        success: (res) => {
            //token失效
            if(res.statusCode === 401){
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
          if (res.data.code == 200) {
            this.exams = this.transformExams(res.data.exams);
            this.haveExam = this.exams.length > 0;
            this.examCnt = this.exams.length;
          }
        },
        fail: (err) => {
          console.log(err);
        }
      });
    },
		methods: {
      toHome() {
        if(this.operation==0)
        uni.setStorageSync("toReview", true);
        else if(this.operation==1)
          uni.setStorageSync("reviewed", true);
        uni.switchTab({
          url: `../home/home`
        });
      },
      transformExams(exams) {
        if(exams==null){//防止空数组报错
          return [];
        }
        return exams.map(exam => {
          // 将日期从 "yyyy/mm/dd" 转换为 "年月日"
          const dateParts = exam.exam_date.split('-');
          const dateInChineseFormat = `${dateParts[0]}年${dateParts[1]}月${dateParts[2]}日`;

          return {
            exam_id: exam.exam_id,
            name: exam.exam_name,
            date: dateInChineseFormat,
            info: `共${exam.question_num}题`,
            questionNum: exam.question_num,
            score: exam.exam_score
          };
        });
      },
      //由date类型转为类似于'2022-01-03'字符串类型
      getExamDate(date) {
        // 获取年、月、日
        const year = date.getFullYear();
        const month = date.getMonth() + 1; // 月份是从0开始的，所以需要+1
        const day = date.getDate();

        // 格式化月份和日期，确保它们总是两位数
        const formattedMonth = month < 10 ? '0' + month : month;
        const formattedDay = day < 10 ? '0' + day : day;

        // 拼接成"YYYY-MM-DD"格式的字符串
        return `${year}-${formattedMonth}-${formattedDay}`;
      },
			toExam(exam) {
        uni.setStorageSync("startExam", JSON.stringify(exam));
        uni.navigateTo({
          url: `../startexam/startexam`
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