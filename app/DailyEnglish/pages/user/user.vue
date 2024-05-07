<template>
	<view>
		<view class="personal-information" >
			<image class="head" src="../../static/pikachu.jpg" @click="GotoPersonal_information"></image> 
			<span class="username">user</span>
		</view>
		 <view class="container1">
			 <view class="words-amount">
				 <span class="number1">120</span>
				 <span class="words1">打卡单词数</span>
			 </view>
			 <view class="dakatianshu">
			 	<span class="number2">12</span>
			 	<span class="words2">总打卡天数</span>
			 </view>
			 <view class="lianxudakatianshu">
			 	<span class="number3">7</span>
			 	<span class="words3">连续打卡天数</span>
			 </view>
		 </view>
		 <view class="container2" @click="goToCalendar">
			 <image class="calendar" src="../../static/calendar.png"></image>
			 <span class="word2">我的日历</span>
			 <image class="right" src="../../static/right.png"></image>
		 </view>
		 <view class="container3">
		 			 <image class="wordbook" src="../../static/wordbook.png"></image>
		 			 <span class="word3">我的单词本</span>
		 			 <image class="right" src="../../static/right.png"></image>
		 </view>
		 <view class="container4">
		 			 <image class="setting" src="../../static/setting.png"></image>
		 			 <span class="word4">设置</span>
		 			 <image class="right" src="../../static/right.png"></image>
		 </view>
		 <view class="container5">
		 			 <image class="feedback" src="../../static/feedback.png"></image>
		 			 <span class="word5">反馈</span>
		 			 <image class="right" src="../../static/right.png"></image>
		 </view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				
			}
		},
		methods: {
			goToCalendar() {
			    // 使用 uniapp 提供的路由跳转方法进行跳转
			    uni.navigateTo({
			      url: '../Calendar/Calendar'
			    });
			  },
			  GotoPersonal_information(){
				  uni.navigateTo({
				    url: '../personal-information/personal-information'
				  });
			  }
		}
	}
</script>

<style>
	body{
		background-color: #f6f6f6;
	}
	.personal-information{
		height: 20rem;
		margin: auto; /* 使用margin自动居中，并添加上下间距 */
		width: 90%;
	    display: flex; /* 使用 flexbox 以方便布局 */
		flex-direction: column; /* 将 flex 子项垂直堆叠 */  
	    align-items: center; /* 如果需要水平居中 */  
	}
	.head{
		margin-top: 3rem;
		width: 9rem;
		height: 9rem;
		border-radius: 50%;
	}
	.username{
		margin-top: 1rem;
		font-size: 30px;
	}
	.container1{
		height: 5rem;
		margin: 20px auto; /* 使用margin自动居中，并添加上下间距 */
		width: 100%;
		margin-bottom: 2rem;
		background-color: #f6f6f6;
		display: flex; /* 使用 flexbox 以方便布局 */
		align-items: center; /* 如果需要水平居中 */
	}
	.words-amount{
		height: 100%;
		width: 25%;
		margin-left: 6%;
		background-color: white;
		border-radius: 2rem;
		display: flex;
		flex-direction: column; /* 将 flex 子项垂直堆叠 */
		align-items: center; /* 如果需要水平居中 */
	}
	.number1{
		margin-top: 0.7rem;
		margin-bottom: 0.3rem;
		font-size: 30px;
	}
	.words1{
		font-size: 12px;
	}
	.dakatianshu{
		height: 100%;
		width: 25%;
		margin-left: 6%;
		background-color: white;
		border-radius: 2rem;
		display: flex;
		flex-direction: column; /* 将 flex 子项垂直堆叠 */
		align-items: center; /* 如果需要水平居中 */
	}
	.number2{
		margin-top: 0.7rem;
		margin-bottom: 0.3rem;
		font-size: 30px;
	}
	.words2{
		font-size: 12px;
	}
	.lianxudakatianshu{
		height: 100%;
		width: 25%;
		margin-left: 6%;
		background-color: white;
		border-radius: 2rem;
		display: flex;
		flex-direction: column; /* 将 flex 子项垂直堆叠 */
		align-items: center; /* 如果需要水平居中 */
	}
	.number3{
		margin-top: 0.7rem;
		margin-bottom: 0.3rem;
		font-size: 30px;
	}
	.words3{
		font-size: 12px;
	}
	.container2{
		height: 4rem;
		margin: 1px auto; /* 使用margin自动居中，并添加上下间距 */
		width: 100%;
		background-color: white;
		display: flex; /* 使用 flexbox 以方便布局 */  
		align-items: center; /* 如果需要水平居中 */  
		justify-content: space-between; /* 在主轴上分配空间，使得子项之间有间隔 */ 
		border-bottom: 2px solid #f6f6f6;
	}
	.calendar{
		width: 2rem;
		height: 2rem;
		margin-left: 0.3rem;
	}
	.word2{
		margin-left: 0.3rem;
		flex-grow: 1;
		font-size: 18px;
	}
	.right{
		width: 1rem;
		height: 1rem;
		margin-left: auto;
		margin-right: 1rem;
	}
	.container3{
		height: 4rem;
		margin: 1px auto; /* 使用margin自动居中，并添加上下间距 */
		width: 100%;
		background-color: white;
		display: flex; /* 使用 flexbox 以方便布局 */  
		align-items: center; /* 如果需要水平居中 */  
		justify-content: space-between; /* 在主轴上分配空间，使得子项之间有间隔 */ 
		border-bottom: 2px solid #f6f6f6;
	}
	.wordbook{
		width: 2rem;
		height: 2rem;
		margin-left: 0.3rem;
	}
	.word3{
		margin-left: 0.3rem;
		flex-grow: 1;
		font-size: 18px;
	}
	.container4{
		height: 4rem;
		margin: 1px auto; /* 使用margin自动居中，并添加上下间距 */
		width: 100%;
		background-color: white;
		display: flex; /* 使用 flexbox 以方便布局 */  
		align-items: center; /* 如果需要水平居中 */  
		justify-content: space-between; /* 在主轴上分配空间，使得子项之间有间隔 */ 
		border-bottom: 2px solid #f6f6f6;
	}
	.setting{
		width: 2rem;
		height: 2rem;
		margin-left: 0.3rem;
	}
	.word4{
		margin-left: 0.3rem;
		flex-grow: 1;
		font-size: 18px;
	}
	.container5{
		height: 4rem;
		margin: 1px auto; /* 使用margin自动居中，并添加上下间距 */
		width: 100%;
		background-color: white;
		display: flex; /* 使用 flexbox 以方便布局 */  
		align-items: center; /* 如果需要水平居中 */  
		justify-content: space-between; /* 在主轴上分配空间，使得子项之间有间隔 */ 
		border-bottom: 2px solid #f6f6f6;
	}
	.feedback{
		width: 2rem;
		height: 2rem;
		margin-left: 0.3rem;
	}
	.word5{
		margin-left: 0.3rem;
		flex-grow: 1;
		font-size: 18px;
	}
</style>
