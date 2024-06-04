<template>
	<view>
		<view class="body-head" style="display:flex">
			<image class="back" src="../../static/back.svg"></image>
			<image class="logout" src="../../static/logout.svg"></image>
		</view>
		<view class="personal-information" >
			<image class="head" src="../../static/pikachu.jpg"></image>
			<span class="username">user</span>
			<view style="display:flex;flex-direction: column;">

			</view>
			<button class="discription" @click="GotoPersonal_information">Edit Profile</button>

		</view>

		<view class="tail">
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
				<view class="calendar-wrapper">
					<image class="calendar" src="../../static/calendar.png"></image>
				</view>
				<span class="word2">我的日历</span>
				<image class="right" src="../../static/back.svg"></image>
			</view>

			<view class="container3">
				<view class="wordbook-wrapper">
					<image class="wordbook" src="../../static/biji2.svg"></image>
				</view>
				<span class="word3">我的单词本</span>
				<image class="right" src="../../static/back.svg"></image>
			</view>
			<view class="container3" @click="goToTeam">
				<view class="team-wrapper">
					<image class="team" src="../../static/team.svg"></image>
				</view>
				<span class="word31">我的团队</span>
				<image class="right" src="../../static/back.svg"></image>
			</view>
			<view class="container4">
				<view class="setting-wrapper">
				<image class="setting" src="../../static/setting.png"></image>
				</view>
				<span class="word4">设置</span>
				<image class="right" src="../../static/back.svg"></image>
			</view>
			<view style="height: 22px;"></view>
			<!-- <view class="container5">
				<view class="feedback-wrapper">
				<image class="feedback" src="../../static/feedback.png"></image>
				</view>
				<span class="word5">反馈</span>
				<image class="right" src="../../static/back.svg"></image>
			</view> -->
		</view>
		<!-- 
		
		<view class="container5">
			<image class="feedback" src="../../static/like.svg"></image>
			<span class="word5">给我们好评</span>
			<image class="right" src="../../static/back.svg"></image>
		</view> -->
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
		onLoad() {
			// 页面加载完成后，获取用户信息
			uni.request({
				url: 'http://localhost:8080/api/users/my_punches',
				method: 'GET',
				header: {
					'Authorization': `Bearer ${uni.getStorageSync('token')}` // 这里需要将 token 放到 header 中
				},
				success: (res) => {
					if (res.statusCode === 200) {
						const data = res.data;
						this.punch_word_num = data.punch_word_num;
						this.total_punch_day = data.total_punch_day;
						this.consecutive_punch_day = data.consecutive_punch_day;
					}
				},
				fail: (err) => {
					console.log(err);
				}
			});
		},

		methods: {
			goToTeam() {
				uni.navigateTo({
					url: '../MyTeam/MyTeam' //Todo: 跳转到团队页面
				});
			},
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
		background-color: #f3f2f5;

	}

	.body-head {
		height: 2rem;
		width: 100%;
		background-color: transparent;
		display: flex;
	}

	.back {
		width: 2rem;
		height: 2rem;
		margin-left: 1rem;
	}

	.logout {
		width: 1.8rem;
		height: 1.8rem;
		margin-left: auto;
		margin-right: 1rem;
	}

	.personal-information {
		position: relative;
		height: 18rem;
		margin: auto;
		width: 100%;
		display: flex;
		flex-direction: column;
		background-color: transparent;
		align-items: center;
	}

	.head {
		margin-top: 2rem;
		width: 8rem;
		height: 8rem;
		border-radius: 25px;
	}

	.username {
		margin-top: 0.5rem;
		font-size: 32px;
	}

	.discription {
		width: 40%;
		height: 2rem;
		margin-top: 1rem;
		border-radius: 12px;
		white-space: nowrap;
		line-height: 2rem;
		text-align: center;
		background-color: transparent;
		border: 1px solid #b4b3b6;
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
		color: #2a9d8fd6;
	}

	.words2 {
		font-size: 12px;
		/* color:#828491; */
	}

	.right1 {
		margin-top: 3.3rem;
		width: 1.4rem;
		height: 1.4rem;
		margin-left: 5rem;
	}

	.tail {
		background-color: white;
		/*左上和右上的圆角*/
		border-top-left-radius: 30px;
		border-top-right-radius: 30px;
		width: 100%;
		margin-top:20px;
	}

	.container1 {
		height: 7rem;
		width: 100%;
		display: flex;
		align-items: center;
		border-radius: 10px;
		background-color: transparent;
	}

	.words-amount {
		height: 100%;
		width: 50%;
		margin-left: 4%;
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
		color: #e9c46cd6;
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
		color: #ba074bd6;
	}

	.words3 {
		font-size: 12px;
		/* color:#828491; */
	}

	.container2 {
		height: 4rem;
		margin: 1px auto;
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.calendar-wrapper {
		margin-left: 2rem;
		display: inline-block;
		padding: 10px;
		border: none;
		background-color: #f3f2f5;
		border-radius: 15px;
		box-sizing: border-box;
		/* 确保padding和border包含在宽高内 */
	}

	.calendar {
		width: 2rem;
		height: 2rem;
		display: block;
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
		margin-right: 2rem;
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

	.wordbook-wrapper {
		margin-left: 2rem;
		display: inline-block;
		padding: 10px;
		border: none;
		background-color: #f3f2f5;
		border-radius: 15px;
		box-sizing: border-box;
		/* 确保padding和border包含在宽高内 */
	}

	.wordbook {
		width: 2rem;
		height: 2rem;
		display: block;
	}

	.team-wrapper {
		margin-left: 2rem;
		display: inline-block;
		padding: 10px;
		border: none;
		background-color: #f3f2f5;
		border-radius: 15px;
		box-sizing: border-box;
		/* 确保padding和border包含在宽高内 */
	}

	.team {
		width: 2rem;
		height: 2rem;
		display: block;
	}

	.word3 {
		margin-left: 0.9rem;
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
		width: 100%;
		background-color: white;
		display: flex;
		align-items: center;
		justify-content: space-between;
	}

	.setting-wrapper {
		margin-left: 2rem;
		display: inline-block;
		padding: 10px;
		border: none;
		background-color: #f3f2f5;
		border-radius: 15px;
		box-sizing: border-box;
		/* 确保padding和border包含在宽高内 */
	}
	.setting {
		width: 2rem;
		height: 2rem;
		display: block;
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

	.feedback-wrapper {
		margin-left: 2rem;
		display: inline-block;
		padding: 10px;
		border: none;
		background-color: #f3f2f5;
		border-radius: 15px;
		box-sizing: border-box;
		/* 确保padding和border包含在宽高内 */
	}
	.feedback {
		width: 2rem;
		height: 2rem;
		display: block;
	}

	.word5 {
		margin-left: 0.9rem;
		flex-grow: 1;
		font-size: 19px;
	}
</style>