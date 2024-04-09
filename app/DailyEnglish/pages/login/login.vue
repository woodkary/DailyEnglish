<template>
	<view>
		<image class="background" src="../../static/background.png"></image>
		<image class="color" src="../../static/color.png"></image>  
		<view class="logo"></view>
		<view class="white-container1">  
			<input class="search-box" type="text" placeholder="请输入账号">  
		</view>  
		<view class="white-container2">
			<input class="search-box" type="text" placeholder="请输入密码">  
		</view>
		<button class="auto"></button>
		<span class="auto1">自动登录</span>
		<button class="forget">忘记密码？</button>
		<button class="login-button">登录</button>
		<button class="register-button">注册</button>
		<button class="button1"></button>
		<span class="text">登录代表你同意用户协议、隐私政策和儿童隐私政策</span>
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
.background {  
	width: 100vw;
    height: 100vh; /* 高度等于视口高度 */  
}  
.color{
	width: 100vw;
	height: 100vh; /* 高度等于视口高度 */  
		position: absolute; /* 绝对定位以覆盖整个容器 */  
		top: 0;  
		left: 0;  
		z-index: 1; /* 设置较低的z-index值 */ 
		opacity: 0.9;
}

.logo {  
	width: 10rem; /* 设置白色框的宽度 */  
	height: 10rem; /* 设置白色框的高度 */  
	background-color: white; /* 设置背景颜色为白色 */  
	border: 1px solid black; /* 设置边框为1像素的黑色实线 */  
	position: absolute; /* 绝对定位以覆盖在图片之上 */  
	top:15rem; 
	left: 12rem;  
	transform: translate(-50%, -50%); /* 将框居中 */  
	z-index: 2; /* 设置最高的z-index值，确保在color之上 */  
}  

.white-container1 {    
	width: 20rem; /* 设置容器的宽度 */    
	height: 2rem; /* 设置容器的高度，根据需要调整 */    
	background-color: white; /* 设置背景颜色为白色 */    
	border: 1px solid black; /* 设置边框为1像素黑色实线 */    
	position: absolute; /* 绝对定位以覆盖在图片之上 */    
	top: 22rem; /* 顶部距离视口50% */  
	left: 2rem; /* 左边距离视口50% */   
	z-index: 2; /* 设置z-index值，确保在color之上 */  
	display: flex; /* 设置为弹性布局 */  
	align-items: center; /* 垂直居中 */  
	justify-content: center; /* 水平居中 */  
	border-radius: 10px; /* 设置边框圆角 */
	opacity: 0.65;
}  

.white-container2 {    
	width: 20rem; /* 设置容器的宽度 */    
	height: 2rem; /* 设置容器的高度，根据需要调整 */    
	background-color: white; /* 设置背景颜色为白色 */    
	border: 1px solid black; /* 设置边框为1像素黑色实线 */    
	position: absolute; /* 绝对定位以覆盖在图片之上 */    
	top: 27rem; /* 顶部距离视口50% */  
	left: 2rem; /* 左边距离视口50% */   
	z-index: 2; /* 设置z-index值，确保在color之上 */  
	display: flex; /* 设置为弹性布局 */  
	align-items: center; /* 垂直居中 */  
	justify-content: center; /* 水平居中 */  
	border-radius: 10px; /* 设置边框圆角 */
	opacity: 0.65;
} 
  
.search-box {    
	width: 100%; /* 设置搜索框宽度为容器的80% */    
	height: 2rem; /* 设置搜索框高度 */       
	border-radius: 10px; /* 设置边框圆角 */    
	border: none; /* 设置边框颜色 */    
	font-size: 0.8rem; /* 设置字体大小 */  
	outline: none; /* 移除默认轮廓 */  
	margin-left: 1rem ;
}  

.auto{
	width: 50rpx;
	height: 50rpx;
	background: transparent; /* 设置背景颜色为白色 */
	border: 1px solid white; /* 设置边框为1像素黑色实线 */    
	position: absolute; /* 绝对定位以覆盖在图片之上 */    
	top: 31rem; 
	left: 2rem;   
	z-index: 2; /* 设置z-index值，确保在color之上 */  
	border-radius: 50%; /* 设置边框圆角 */
	scale: 0.5;
}

.auto1{
	position: absolute;
	z-index: 2;
	color: white;
	top: 31.3rem;
	left: 4rem;
	font-size: 25rpx;
}

.forget{
	background: transparent; /* 设置背景颜色为白色 */
	border: none; /* 设置边框为1像素黑色实线 */    
	color: white;
	position: absolute; /* 绝对定位以覆盖在图片之上 */    
	top: 30.8rem; 
	left: 17rem;   
	font-size: 25rpx;
	z-index: 2; /* 设置z-index值，确保在color之上 */  
	cursor: pointer;
}

.login-button{
	background-color: #75C2FD; /* 设置背景颜色为白色 */
	border: none; /* 设置边框为1像素黑色实线 */    
	color: white;
	position: absolute; /* 绝对定位以覆盖在图片之上 */    
	top: 35rem; 
	left: 4rem;
	width: 15rem;
	height: 3rem;
	font-size: 35rpx;
	z-index: 2; /* 设置z-index值，确保在color之上 */  
	cursor: pointer;
	border-radius: 20rpx;
}

.register-button{
	background: transparent; /* 设置背景颜色为白色 */
	border: none; /* 设置边框为1像素黑色实线 */    
	color: white;
	position: absolute; /* 绝对定位以覆盖在图片之上 */    
	top: 39rem; 
	left: 10rem;   
	font-size: 25rpx;
	z-index: 2; /* 设置z-index值，确保在color之上 */  
	cursor: pointer;
}

.button1{
	width: 50rpx;
	height: 50rpx;
	background: transparent; /* 设置背景颜色为白色 */
	border: 1px solid white; /* 设置边框为1像素黑色实线 */    
	position: absolute; /* 绝对定位以覆盖在图片之上 */    
	top: 42rem; 
	left: 3rem;   
	z-index: 2; /* 设置z-index值，确保在color之上 */  
	border-radius: 50%; /* 设置边框圆角 */
	scale: 0.5;
}

.text{
	position: absolute;
	z-index: 2;
	color: white;
	top: 42.3rem;
	left: 4.5rem;
	font-size: 20rpx;
}
</style>
