<template>
	<view>
		<view class="today-container">
			<span class="title">今日考试</span>
			<span v-if="exams.length === 0" class="no-exam">今日暂无考试，<navigator>前往复习</navigator></span>
			<view class="todo-exam" v-for="exam in exams" :key="exam.name">
				<view class="row1">
					<text class="exam-name">{{ exam.name }}</text>
					<text class="exam-time"> {{ formatTimeRange(exam.start_time, exam.duration) }}</text>
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
				<button class="choice selected" @click="sortExamsBy('date')">时间顺序</button>
				<button class="choice" @click="sortExamsReverse('date')">时间逆序</button>
				<button class="choice" @click="sortExamsBy('score')">成绩顺序</button>
				<button class="choice" @click="sortExamsReverse('score')">成绩逆序</button>
			</view>
			<view class="finished-exam" v-for="exam in finishedExams" :key="exam.date">
				<image class="level" src="@/static/score1.svg" v-if="exam.score>= 80"></image>
				<image class="level" src="@/static/score2.svg" v-else-if="exam.score >= 60&&exam.score<80"></image>
				<image class="level" src="@/static/score3.png" v-else></image>
				<!-- <image class="level" src="@/static/score1.svg" v-else></image> -->
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
			return {
				exams: [/*{
						exam_id: 1,
						name: '第一单元第一次小测',
						start_time: '20:00',
						duration: 60,
						info: '共20题',
						questionNum: 20,
					},
					{
						exam_id: 2,
						name: '第一单元第一次小测',
						start_time: '20:00',
						duration: 60,
						info: '共20题',
						questionNum: 20
					},
					// ...更多的考试对象*/
				],
				finishedExams: [/*{
						exam_id: 3,
						name: '第一单元第一次小测',
						date: '2023年1月1日',
						info: '共20题',
						questionNum: 20,
						score: 95
					},
					{
						exam_id: 4,
						name: '第一单元第二次小测',
						date: '2023年1月1日',
						info: '共20题',
						questionNum: 20,
						score: 35
					},
					{
						exam_id: 5,
						name: '第二单元第一次小测',
						date: '2023年1月1日',
						info: '共20题',
						questionNum: 20,
						score: 70
					},
					// ...更多的考试结果对象*/
				]
			}

		},
		onLoad() {
			this.getTodayExams();
			this.getPreviousExams();
		},
		methods: {
      getDateByString(dateStr){
        // 正则表达式匹配中文字符串中的年、月、日
        const regex = /(\d+)年(\d+)月(\d+)日/;
        const match = dateStr.match(regex);

        if (!match) {
          // 如果字符串不匹配预期的格式，则返回null
          return null;
        }

        // 提取年、月、日
        const year = parseInt(match[1], 10);
        const month = parseInt(match[2], 10) - 1; // 月份是从0开始计数的
        const day = parseInt(match[3], 10);

        // 使用提取的年、月、日创建Date对象
        return new Date(year, month, day);
      },
      sortExamsBy(param){
        this.finishedExams.sort((a, b) => {
          if(param === 'date')
            return this.getDateByString(a.date) - this.getDateByString(b.date);
          else if(param ==='score')
            return b.score - a.score;
        });
      },
      sortExamsReverse(param){
        this.finishedExams.sort((a, b) => {
          if(param === 'date')
            return this.getDateByString(b.date) - this.getDateByString(a.date);
          else if(param ==='score')
            return a.score - b.score;
        });
      },
			formatTimeRange(start_time, duration) {
				// 解析开始时间
				const [startHour, startMinute] = start_time.split(':').map(Number);

				// 计算结束时间
				const endMinute = (startMinute + duration) % 60;
				const endHour = startHour + Math.floor((startMinute + duration) / 60);
				const endTime = `${endHour.toString().padStart(2, '0')}:${endMinute.toString().padStart(2, '0')}`;

				// 返回格式化的时间范围字符串
				return `${start_time} ~${endTime}`;
			},
			takeExam(exam) {
				uni.setStorageSync("startExam", JSON.stringify(exam));
				uni.navigateTo({
					url: `../startexam/startexam`
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
      }
,
			getTodayExams() {
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
						}
					},
					fail: (err) => {
						console.log(err);
					}
				});
			},
			getPreviousExams() {
				// 从服务器获取之前的考试记录
				uni.request({
					url: 'http://localhost:8080/api/exams/previous_examinations',
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
						this.finishedExams = this.transformExams(res.data.exams);
					}
				});
			},

			viewDetails(exam) {
				// 跳转到考试详情页面
				uni.navigateTo({
					url: `../exam_details/exam_details?exam_id=${exam.exam_id}&exam_name=${exam.name}`
				});
			},

			transformExams(exams) {
        if(exams==null){//防止空数组报错
          return [];
        }
				return exams.map(exam => {
					// 将日期从 "yyyy/mm/dd" 转换为 "年月日"
					const dateParts = exam.exam_date.split('-');
          console.log(dateParts); // 输出: );
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
		top: -15px;
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