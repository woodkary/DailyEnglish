<template>
	<view>
		<view class="container">
			<view class="left-container">
				<h3 class="exam-number">第一单元第一次小测</h3>
				<p class="date">2023年1月1日</p>
				<view style="display: flex; justify-content: flex-start; margin-top: 25px; margin-bottom: 20px;">
					<p class="question-number">共20题</p>
					<p class="correct-number">答对19/20题</p>
				</view>
			</view>
			<view class="right-container">
				<p class="point">95</p>
				<p class="small-text">分</p>
			</view>
		</view>
		<view class="container2">
			<view class="header-container">
				<h3>考试题目</h3>
				<view class="search">
					<image class="search-icon" src="/static/search.svg"></image>
					<input placeholder="搜索">
				</view>
			</view>
			<view class="button-group">
				<button class="button1">题目顺序</button>
				<button class="button2">题目逆序</button>
				<button class="button3">分数顺序</button>
				<button class="button4">分数逆序</button>
			</view>
		</view>
		<view class="title-container">
			<view style="    display: flex; justify-content: space-between; width: 90%;">
				<p class="title-number">1.</p>
				<p class="wrong-title-point">0/5</p>
			</view>
			<p class="sentence1">—____ is your brother?</p>
			<p class="sentence2">—He is a doctor.</p>
			<p class="answer">A.What B.Who C.Where D. How</p>
			<view style="    display: flex; justify-content: space-between; width: 90%; margin-top: 20px;">
				<p class="wrong-answer">我的答案:B</p>
				<p class="true-answer">正确答案:A</p>
			</view>
			<button class="big-button">查看解析</button>
		</view>
		<view class="title-container">
			<view style="    display: flex; justify-content: space-between; width: 90%;">
				<p class="title-number">1.</p>
				<p class="correct-title-point">5/5</p>
			</view>
			<p class="sentence1">—____is the nearest post office, please?</p>
			<p class="sentence2">—lt’s about half an hour’s walk from here.</p>
			<p class="answer">A.How far B.How long C. How often D. How soon</p>
			<view style="    display: flex; justify-content: space-between; width: 90%; margin-top: 20px;">
				<p class="correct-answer">我的答案:A</p>
				<p class="true-answer">正确答案:A</p>
			</view>
			<button class="big-button">查看解析</button>
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
			
		}
	}
</script>

<style>
	.container{
		width: 90%;
		margin: 0 auto;
		margin-top: 1rem;
		display: flex;  
		align-items: stretch; /* 拉伸子元素以匹配最高的子元素 */  
		justify-content: space-between; /* 水平对齐 */  
		border: 2px solid #e6e6e6;
		border-radius: 10px;
		padding: 15px;
		box-shadow: 0 0 0 2px rgb(247 127 0 / 10%);
	}
	.left-container{
		flex: 0 0 60%;
		display: flex;
		flex-direction: column;
		/* 设置为列方向，实现从上往下排列 */
		align-items: stretch;
		/* 子项默认拉伸以填充整个容器 */
	}
	.right-container {  
	  flex: 0 0 40%; /* 右边占40%宽度 */  
	  display: flex;  
	    flex-direction: column; /* 垂直方向 */  
	    justify-content: center; /* 垂直居中 */  
	    align-items: center; /* 水平居中 */  
	    text-align: center; /* 水平文本居中（对于单行文本）*/ 
	  /* 其他样式... */  
	} 
	.exam-number{
		margin-top: 10px;
	}
	.date{
		margin-top: 25px;
		color: #456de7;
		font-size: 20px;
	}
	.question-number{
		color: #a7a7a7;
		margin-right: 10px;
		font-size: 18px;
	}
	.correct-number{
		color: #a7a7a7;
		font-size: 18px;
	}
	.point{
		  font-size: 80px; /* 大字号的数字字体大小 */  
		  font-style: italic; /* 设置为斜体 */
			margin-left: -40px;
			color: #3FC681;
	}
	.small-text{
			font-style: italic; /* 设置为斜体 */ 
			 font-size: 26px;
			 margin-top: -40px;
			 margin-left: 80px;
			 color: #3FC681;
	}
	.container2 {  
	    width: 95%;  
	    margin: 0 auto;  
	    margin-top: 1rem;
		  margin-bottom: 1rem;
	    display: flex;  
	    flex-direction: column; /* 如果需要垂直堆叠 */  
	}  
	  
	.header-container {  
	    display: flex;  
	    align-items: center;  
	    justify-content: space-between; /* 根据需要，如果h3和search之间需要间隔 */
		  margin-bottom: 10px;
	}  
	  
	.search {  
	    flex: 1; /* 让.search占据剩余空间 */  
	    background-color: #ffffff;  
	    border-radius: 50rpx;  
	    display: flex;  
	    padding: 5rpx;  
	    align-items: center;  
	    /* 移除margin-left，或根据需要设置 */  
	    border: 1px solid #e6e6e6;
		box-shadow: 0 0 0 2px rgb(247 127 0 / 10%);
		margin-left: 30px;
	}  
	  
	.search-icon {  
	    width: 25px;  
	    height: 25px;  
	    margin-right: 20rpx;  
	}  
	  
	.search-input {  
	    flex: 1; /* 让输入框占据.search的剩余空间 */  
	    border: none;  
	    padding: 0 10rpx; /* 根据需要添加内边距 */  
	}
	.button-group{
		    display: flex;  
		    justify-content: flex-start; /* 让按钮从左边开始排列 */ 
			 width: 80%;
	}
	.button-group button {  
	    margin-right: 10px; /* 按钮之间的间距 */  
	    padding: 0 10px; /* 内边距，控制按钮的大小 */  
	    border: none; /* 移除默认的边框 */  
	    border-radius: 5px; /* 圆角 */  
	    background-color: #ededed; /* 背景色 */  
	    color: gray; /* 文字颜色 */  
	    font-size: 12px; /* 文字大小 */  
	    cursor: pointer; /* 鼠标悬停时变为小手形状 */  
	}
	.button-group .button1{
		color: #2e47e7;
	 }
	 .title-container{
		width: 90%;
		margin: 0 auto;
		 border: 2px solid #e6e6e6;
		 border-radius: 10px;
		 padding: 15px;
		 box-shadow: 0 0 0 2px rgb(247 127 0 / 10%);
		 margin-bottom: 20px;
	 }
	 .wrong-title-point{
		 font-size: 36px;
		font-style: italic; /* 设置为斜体 */ 
		 color: #ee4e41;
	 }
	 .correct-title-point{
		 font-size: 36px;
		 font-style: italic; /* 设置为斜体 */ 
		  color: #3fc681;
	 }
	 .sentence1{
		 margin-left: 30px;
		 margin-bottom: 5px;
	 }
	 .sentence2{
		 margin-left: 30px;
		 margin-bottom: 5px;
	 }
	 .answer{
	 		 margin-left: 30px;
	 }
	 .wrong-answer{
		 color: red;
	 }
	 .correct-answer{
	 		 color: #3fc681;
	 }
	 .true-answer{
		 color: #2e47e7;
	 }
	 .big-button{
		 width: 90%;
		 background-color: #2e47e7;
		 color: white;
		 margin-top: 20px;
		 font-size: 14px;
	 }
</style>
