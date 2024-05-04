<template>
	<view class="body">
		<view class="color">
		</view>
		<view>
			<view class="calendar">
				<view class="head">
					<text class="date">{{ year }}年{{ month }}月</text>
					<button class="last-or-next" @click="subMonth">
						<image class="icon" src="../../static/last.svg"></image>
					</button>
					<button class="last-or-next" @click="addMonth">
						<image class="icon" src="../../static/next.svg"></image>
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
            <span class="badge" v-if="date.hasExam"></span>
					</view>
				</view>
			</view>
		</view>
		<view>
			<view class="examMsg">
				<text class="title">{{ chosenMonth }}月{{ chosenDay }}日</text>
				<view class="card-container">
					<view class="card" v-if="getChosenDateFromDates()==1" id="daka">
						<image src="../../static/not-done.svg"></image>
						<text class="title">打卡计划:</text>
						<text class="state">未完成</text>
					</view>
          <view v-else-if="getChosenDateFromDates()==0">
            <view  class="card" id="daka">
              <image src="../../static/done.svg"></image>
              <text class="title">打卡计划:</text>
              <text class="state">已完成</text>
            </view>
            <view class="words">
              <!--todo 为这个span添加样式-->
              <span v-for="(word, index) in punchWords" :key="index">{{ word.word }}：
                {{ word.meanings.verb!=null?'v.':'' }}{{ word.meanings.verb }}
                {{word.meanings.noun!=null?'n.':'' }} {{ word.meanings.noun }}
                {{ word.meanings.adj!=null?'adj.':'' }} {{ word.meanings.adj }}
                {{ word.meanings.adv!=null?'adv.':'' }} {{ word.meanings.prep }}
                {{ word.meanings.prep!=null?'prep.':'' }} {{ word.meanings.adv }}
              </span>
            </view>
          </view>
          <view class="card" v-else id="daka">
            <image src="../../static/not-done.svg"></image>
            <text class="title">打卡计划:</text>
            <text class="state">已过期</text>
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
				punchMsg: 63398853,
        //0000 0011 1100 0111 0110 0011 1100 0101,
        //存储前32天打卡信息，
        //每一位表示一天，0表示未打卡，1表示已打卡
        chosenYear: 2024, // 选中的年份
        chosenMonth: 5, // 选中的月份
        chosenDay: 4, // 选中的日期
        punchWords: [{
          word: "refuse",
          meanings: {
            verb: "拒绝,谢绝",
            noun: "废物",
            adj: "扔掉的，无用的",
            adv: null,
            prep: null,
           }
          },
          {
            word: "objective",
            meanings: {
              verb: null,
              noun: "目的，目标，<语法>宾格，物镜",
              adj: "客观的，<语法>宾格的，真实的，目标的",
              adv: null,
              prep: null,
             }
          }
        ]
			}
		},

		beforeMount() {
			this.generateDates();
		},
		methods: {
			//TODO 发送网络请求获取考试信息
			requestExamMsg() {

			},
      //判断是否有未完成的打卡计划
      getChosenDateFromDates(){
        let date=new Date(this.chosenYear,this.chosenMonth-1,this.chosenDay);
        let diffDays=Math.floor((new Date()-date)/(24*60*60*1000));
        if(diffDays>=0&&diffDays<32)
          return this.punchMsg>>diffDays&1;
        return -1;
      },
      subMonth() {
        this.month--;
        if (this.month < 1) {
          this.month = 12;
          this.year--;
        }
        this.generateDates();
      },
      addMonth() {
        this.month++;
        if (this.month > 12) {
          this.month = 1;
          this.year++;
        }
        this.generateDates();
      },
			handleClick(date) {
        let year=date.date.getFullYear();
        let month=date.date.getMonth()+1;
        let day=date.date.getDate();
        console.log(year,month,day);
        this.chosenYear=year;
        this.chosenMonth=month;
        this.chosenDay=day;
				if (date.hasExam) {
					console.log('未完成打卡计划日期：', date.value);
					//TODO 跳转到考试页面及其他操作
				}else{
          console.log('已完成打卡计划日期：', date.value);
          //TODO 从后端获取当天打卡的所有单词
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
          let date=new Date()-i*24*60*60*1000;
					this.dates.push({
            date:date,
						value: '',
						dayOfWeek: '',
						hasExam: false

					});
				}
				// 添加日期
				for (let i = 1; i <= totalDays; i++) {
					let dayOfWeek = (firstDayOfWeek + i - 1) % 7; // 计算当前日期对应的星期几（0 表示星期日，1 表示星期一，以此类推）
          let date=new Date(this.year, this.month-1, i);
          let today=new Date();
          //计算当前日期与今天的差值，并判断是否有考试
          let diffDays=Math.floor((today-date)/(24*60*60*1000));
          let hasExam=diffDays>=0&&diffDays<32?this.punchMsg>>diffDays&1:false;
					this.dates.push({
            date:date,
						value: i,
						dayOfWeek: dayOfWeek,
						hasExam: hasExam // 判断当前日期是否有考试
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
  .badge {
    position: absolute;
    top: 0;
    right: 0;
    width: 16px;
    height: 16px;
    border-radius: 50%;
    background-color: red;
    /* 使用伪元素创建圆点 */
    &::before {
      content: '';
      position: absolute;
      top: 50%;
      left: 50%;
      width: 12px;
      height: 12px;
      border-radius: 50%;
      background-color: inherit;
      transform: translate(-50%, -50%);
    }
  }
  .last-or-next {
    background-color: transparent;
    border: none;
    outline: none;

  }
  .icon {
    margin-top: 20rpx;
    width: 60rpx;
    height: 60rpx;
    filter: invert(1);
  }
</style>