<template>
	<view class="container">
		<view class="head">
			<image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
			<span class="title">我的团队</span>
			<image class="join" src="../../static/add.png"  @click="goToJoin">+</image>
		</view>
		<view class="container1">
			<view class="team-icon-container">
				<span class="team-initial">M</span>
			</view>
			<view class="team-info">
				<view class="team-name">
					<text class="team-name-text">{{teamName}}</text>
				</view>
				<view class="people">
					<text class="people-text">团队人数：{{memberNum}}</text>
				</view>
			</view>
		</view>
		<view class="container2">
			<view class="team-captain-row">
				<span>团队队长</span>
			</view>
			<view class="name-row">
				<view class="people-icon-container">
					<span class="people-initial">{{firstLetter}}</span>
				</view>
				<span class="manager-name">{{managerName}}</span>
			</view>
		</view>
		<view class="container3">
			<view class="team-captain-row">
				<span>团队成员</span>
			</view>
			<view v-for="(item, index) in members" :key="index" class="name-row">
				<view class="people-icon-container">
					<span class="people-initial">{{item.userName[0]}}</span>
				</view>
				<span class="manager-name">{{item.userName}}</span>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				teamName: "春田花花幼稚园",
				managerName: "佐藤太郎",
				memberNum: 50,
				members: [{
						userName: "张三",
						userSex: 1, //1:男 0:女
					},
					{
						userName: "李四",
						userSex: 0, //1:男 0:女
					}
				]

			}
		},
		onLoad() {
			//获取所有团队成员
			uni.request({
				url: "http://localhost:8080/api/users/my_team",
				method:'GET',
				header:{
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
					if(res.data.code==200){
						//获取成功
						let teamInfo=res.data.team;
						this.teamName=teamInfo.team_name;
						this.managerName=teamInfo.manager_name;
						this.memberNum=teamInfo.member_num;
						this.members=[];
						teamInfo.member_list.forEach((member)=>{
							this.members.push({
								userName: member.user_name,
								userSex: member.user_sex
							});
						});
						// 获取managerName的第一个字符并更新到data中
						this.setData({
							firstLetter: this.data.managerName[0]
						});
					}
				},
				fail: (error) => {
					console.log(error);
				}
			})
			// // 获取managerName的第一个字符并更新到data中
			// this.setData({
			// 	firstLetter: this.data.managerName[0]
			// });
		},
		methods: {
      handleBack() {
        uni.navigateBack();
      },
			goToJoin(){
				uni.navigateTo({
					url: '../Join/Join'
				})
			}
		}
	}
</script>

<style>
	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
		height: 100vh;
		width: 100vw;
		background-color: #fafafa;
	}

	.head {
		display: flex;
		height: 2rem;
		width: 100vw;
		background-color: white;
	}

	.back-icon {
		width: 2rem;
		/*图标宽度 */
		height: 2rem;

	}

	.title {
		margin-left: 7rem;
		font-size: 48rpx;
		font-weight: bold;
		margin-top: 2rpx;
	}

	.container1 {
		display: flex;
		width: 100vw;
		padding: 20px;
		/* 调整容器内边距 */
	}

	.team-icon-container {
		width: 50px;
		/* 设置圆圈的直径 */
		height: 50px;
		/* 设置圆圈的直径 */
		border-radius: 50%;
		/* 使其成为圆形 */
		background-color: #4CAF50;
		/* 圆圈的背景颜色 */
		display: flex;
		align-items: center;
		justify-content: center;
		margin-right: 10px;
		/* 圆圈和文字之间的间距 */
		margin-left: 30px;
	}

	.team-initial {
		color: white;
		/* 文字颜色 */
		font-size: 24px;
		/* 文字大小 */
		font-weight: bold;
		/* 文字加粗 */
	}

	.join{
		width: 25px;
		/* 设置圆圈的直径 */
		height: 25px;
		/* 设置圆圈的直径 */
		display: flex;
		align-items: center;
		justify-content: center;
		margin-left: 25%;
		margin-top: 2%;
	}
	
	.team-info {
		margin-left: 20px;
		/* 调整图片和文字之间的间距 */
		display: flex;
		flex-direction: column;
	}

	.team-name-text,
	{
	font-size: 16px;
	/* 调整文字大小 */
	color: black;
	}

	.people {
		margin-top: 5px;
		/* 调整文字和文字之间的间距 */
	}

	.people-text {
		font-size: 14px;
		/* 调整文字大小 */
	}

	.container2 {
		display: flex;
		flex-direction: column;
		width: 100vw;
		padding: 20px;
	}

	.team-captain-row {
		margin-bottom: 10px;
		margin-left: 20px;
	}

	.name-row {
		display: flex;
		align-items: center;
		width: 100%;
		background-color: white;
		padding: 10px;
	}

	.people-icon-container {
		width: 50px;
		/* 设置圆圈的直径 */
		height: 50px;
		/* 设置圆圈的直径 */
		border-radius: 50%;
		/* 使其成为圆形 */
		background-color: white;
		/* 圆圈的背景颜色 */
		display: flex;
		align-items: center;
		justify-content: center;
		margin-right: 10px;
		/* 圆圈和文字之间的间距 */
		margin-left: 30px;
		border: 1px solid #b2b2b2;
	}

	.people-initial {
		color:black;
		/* 文字颜色 */
		font-size: 24px;
		/* 文字大小 */
		/* font-weight: bold; */
		/* 文字加粗 */
	}

	.manager-name {
		font-size: 18px;
		margin-left: 10px;
	}
	
	.container3 {
		display: flex;
		flex-direction: column;
		width: 100vw;
		padding: 20px;
		.name-row {
			display: flex;
			align-items: center;
			width: 100%;
			background-color: white;
			padding: 10px;
			border-bottom: 1px solid #e6e6e6;
		}
	}
</style>