<template>
	<view>
		<view class="personal-information" @click="GotoPersonal_information">
			<image class="head" src="../../static/pikachu.jpg"></image>
			<view style="display:flex;flex-direction: column;">
				<span class="username">user</span>
				
			</view>
			<image class="right1" src="../../static/back.svg"></image>

		</view>

		
		<view class="container1">
			<view class="words-amount">
				<span class="number1">{{ punch_word_num }}</span>
				<span class="words1">打卡单词数</span>
			</view>
			
			<view class="lianxudakatianshu">
				<span class="number3">{{ consecutive_punch_day }}</span>
				<span class="words3">连续打卡天数</span>
			</view>
			
			<view class="dakatianshu">
				<span class="number2">{{ total_punch_day }}</span>
				<span class="words2">总打卡天数</span>
			</view>
		</view>
		
		<view class="container2" @click="goToCalendar">
			<image class="calendar" src="../../static/calendar.png"></image>
			<span class="word2">我的日历</span>
			<image class="right" src="../../static/back.svg"></image>
		</view>
		
		<view class="container3">
			<image class="wordbook" src="../../static/biji2.svg"></image>
			<span class="word3">我的单词本</span>
			<image class="right" src="../../static/back.svg"></image>
		</view>
		<view class="container3">
			<image class="team" src="../../static/team.svg"></image>
			<span class="word31">我的团队</span>
			<image class="right" src="../../static/back.svg"></image>
		</view>
		<view class="container4">
			<image class="setting" src="../../static/setting.png"></image>
			<span class="word4">设置</span>
			<image class="right" src="../../static/back.svg"></image>
		</view>
		<view class="container5">
			<image class="feedback" src="../../static/feedback.png"></image>
			<span class="word5">反馈</span>
			<image class="right" src="../../static/back.svg"></image>
		</view>
		<view class="container5">
			<image class="feedback" src="../../static/like.svg"></image>
			<span class="word5">给我们好评</span>
			<image class="right" src="../../static/back.svg"></image>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				punch_word_num: 120,
				total_punch_day: 12,
				consecutive_punch_day: 7
			}
		},
		// onLoad() {
		//   //获取一些打卡数据
		//   uni.request({
		//     url: '/api/users/my_punches',
		//     method: 'GET',
		//     header: {
		//       'Authorization': 'Bearer '+ uni.getStorageSync('token')
		//     },
		//     success: (res) => {
		//       console.log(res.data);
		//       this.punch_word_num = res.data.punch_word_num;
		//       this.total_punch_day = res.data.total_punch_day;
		//       this.consecutive_punch_day = res.data.consecutive_punch_day;
		//     },
		//     fail: (err) => {
		//       console.log(err);
		//     }
		//   });
		// },
		methods: {
			goToCalendar() {
				// 使用 uniapp 提供的路由跳转方法进行跳转
				uni.navigateTo({
					url: '../Calendar/Calendar'
				});
			},
			GotoPersonal_information() {
				uni.navigateTo({
					url: '../personal-information/personal-information'
				});
			}
		}
	}
</script>

<style>
	@font-face {
		font-family: 'w03';
		src: url('@/static/Circe Rounded W03 Bold.otf');
	}
	body {
		background-color: #fcfcfc;
	}

	.personal-information {
		height: 10rem;
		margin: auto;
		/* 使用margin自动居中，并添加上下间距 */
		width: 86%;
		display: flex;
		/* 使用 flexbox 以方便布局 */
	}

	.head {
		margin-top: 3rem;
		width: 5rem;
		height: 5rem;
		border-radius: 30%;
		margin-left: 0rem;
		border: 1px solid #f6f6f6;
		box-shadow: 0 0 0 3px white;
	}

	.username {
		margin-top: 3rem;
		font-size: 38px;
		margin-left: 1rem;
	}

	.dakatianshu {
		height: 100%;
		width: 50%;
		margin-left: 7%;
		display: flex;
		flex-direction: column;
		/* 将 flex 子项垂直堆叠 */
		align-items: center;
	}

	.number2 {
		margin-top: 0.7rem;
		margin-bottom: 0.2rem;
		font-size: 60px;
		font-family: 'w03';
		color:#2a9d8fd6;
	}

	.words2 {
		font-size: 12px;
		/* color:#828491; */
	}

	.right1 {
		margin-top: 4rem;
		width: 1.4rem;
		height: 1.4rem;
		margin-left: 6rem;
		/*翻转180度*/
		transform: rotate(180deg);
	}

	.container1 {
		height: 7rem;
		width:90%;
		margin-bottom:20px;
		margin-left: 5%;
		display: flex;
		/* 使用 flexbox 以方便布局 */
		align-items: center;
		/* 如果需要水平居中 */
		border-radius: 10px;
		box-shadow: rgba(0, 0, 0, 0.45) 0px 10px 20px -20px;
		background-color: white;
	}

	.words-amount {
		height: 100%;
		width: 50%;
		margin-left:7%;
		display: flex;
		flex-direction: column;
		/* 将 flex 子项垂直堆叠 */
		align-items: center;
	}

	.number1 {
		margin-top: 0.7rem;
		margin-bottom: 0.2rem;
		font-size: 60px;
		font-family: 'w03';
		color:#e9c46cd6;
	}

	.words1 {
		font-size: 12px;
		/* color:#828491; */
	}



	.lianxudakatianshu {
		height: 100%;
		width: 50%;
		margin-left: 7%;
		display: flex;
		flex-direction: column;
		/* 将 flex 子项垂直堆叠 */
		align-items: center;
		/* 如果需要水平居中 */
		
	}

	.number3 {
		margin-top: 0.7rem;
		margin-bottom: 0.2rem;
		font-size: 60px;
		font-family: 'w03';
		color:#ba074bd6;
	}

	.words3 {
		font-size: 12px;
		/* color:#828491; */
	}

	.container2 {
		height: 4rem;
		margin: 1px auto;
		margin-top: 60px;
		/* 使用margin自动居中，并添加上下间距 */
		background-color: white;
		display: flex;
		/* 使用 flexbox 以方便布局 */
		align-items: center;
		/* 如果需要水平居中 */
		justify-content: space-between;
		/* 在主轴上分配空间，使得子项之间有间隔 */
		/* border-bottom: 2px solid #f6f6f6; */
	}

	.calendar {
		width: 2rem;
		height: 2rem;
		margin-left: 0.9rem;
	}

	.word2 {
		margin-left: 0.9rem;
		flex-grow: 1;
		font-size: 19px;
	}

	.right {
		width: 1.2rem;
		height: 1.2rem;
		margin-left: auto;
		margin-right: 1rem;
		rotate: 180deg;
	}

	.container3 {
		height: 4rem;
		margin: 1px auto;
		/* 使用margin自动居中，并添加上下间距 */
		width: 100%;
		background-color: white;
		display: flex;
		/* 使用 flexbox 以方便布局 */
		align-items: center;
		/* 如果需要水平居中 */
		justify-content: space-between;
		/* 在主轴上分配空间，使得子项之间有间隔 */
		/* border-bottom: 2px solid #f6f6f6; */
	}

	.wordbook {
		width: 2.5rem;
		height: 2.5rem;
		margin-left: 0.7rem;
	}

	.team {
		width: 2rem;
		height: 2rem;
		margin-left: 0.9rem;
	}

	.word3 {
		margin-left: 0.6rem;
		flex-grow: 1;
		font-size: 19px;
	}

	.word31 {
		margin-left: 1rem;
		flex-grow: 1;
		font-size: 19px;
	}

	.container4 {
		height: 4rem;
		margin: 1px auto;
		margin-top:10px;
		/* 使用margin自动居中，并添加上下间距 */
		width: 100%;
		background-color: white;
		display: flex;
		/* 使用 flexbox 以方便布局 */
		align-items: center;
		/* 如果需要水平居中 */
		justify-content: space-between;
		/* 在主轴上分配空间，使得子项之间有间隔 */
		/* border-bottom: 2px solid #f6f6f6; */
	}

	.setting {
		width: 2rem;
		height: 2rem;
		margin-left: 0.9rem;
	}

	.word4 {
		margin-left: 0.9rem;
		flex-grow: 1;
		font-size: 19px;
	}

	.container5 {
		height: 4rem;
		margin: 1px auto;
		/* 使用margin自动居中，并添加上下间距 */
		width: 100%;
		background-color: white;
		display: flex;
		/* 使用 flexbox 以方便布局 */
		align-items: center;
		/* 如果需要水平居中 */
		justify-content: space-between;
		/* 在主轴上分配空间，使得子项之间有间隔 */
		/* border-bottom: 2px solid #f6f6f6; */
	}

	.feedback {
		width: 2rem;
		height: 2rem;
		margin-left: 0.9rem;
	}

	.word5 {
		margin-left: 0.9rem;
		flex-grow: 1;
		font-size: 19px;
	}
</style>