<template>
  <view class="container">
    <view class="head">
      <image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
      <span>我的团队</span>
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
        members: [
          {
            userName: "张三",
            userSex: 1,//1:男 0:女
          },
          {
            userName: "李四",
            userSex: 0,//1:男 0:女
          }
        ]

			}
		},
		methods: {
			
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
}

.head {
  display: flex;
  height: 2rem;
  width: 100vw;
  background-color: transparent;
}

.back-icon {
  width: 2rem;
  /*图标宽度 */
  height: 2rem;

}

span {
  margin-left: 7rem;
  font-size: 40rpx;
  font-weight: bold;
  margin-top: 2rpx;
}
.container1 {
  height: 7rem;
  width:100%;
  margin-bottom:20px;
  display: flex;
  /* 使用 flexbox 以方便布局 */

  border-radius: 10px;
  box-shadow: rgba(0, 0, 0, 0.45) 0px 10px 20px -20px;
  background-color: white;
}
.team-name{
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;

}
.team-name-text{
  font-size: 20rpx;
  font-weight: lighter;
}

</style>
