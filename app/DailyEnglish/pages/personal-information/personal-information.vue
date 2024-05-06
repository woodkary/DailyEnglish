<template>
	<view class="container">
		<image class="back-icon" src="../../static/back.svg"></image>
		<view class="edit">
			<image class="photo" src="../../static/logo.png"></image>
			<view class="items">
				<view class="item" id="username">
					<text>用户名</text>
					<input placeholder="wwswad">
				</view>
				<view class="item" id="email">
					<text>邮箱</text>
					<input placeholder="3242412@xy.com">
				</view>
				<view class="item" id="phone">
					<text>手机号</text>
					<input placeholder="12312312312">
				</view>
				<view class="item" id="birthday">
					<text>生日</text>
					<input placeholder="1999-01-01">
				</view>
				<view class="item" id="team">
					<text>团队</text>
					<span>无</span>
				</view>
			</view>
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
	.container {
		display: flex;
		width: 100%;
		height: 100%;
	}

	.back-icon {
		width: 2rem;
		/*图标宽度 */
		height: 2rem;
		/*图标高度 */
		position: absolute;
		/*绝对定位 */
		top: 0.8rem;
		/*距离顶部20px */
		left: 0.5rem;
		/*距离左侧20px */
		cursor: pointer;
		/*鼠标移上去显示小手 */
	}

	.edit {
		position: absolute;
		top: 80px;
		/* border: 1rpx solid black; */
		width: 100%;
		height:1400rpx;
	}

	.photo {
		width: 280rpx;
		height:280rpx;
		top:80rpx;
		left: 240rpx;
		border-radius: 50%;
		box-shadow: 0px 10px 25px rgb(0, 0, 0, 0.2);
	}
	.items {
		width: 100%;
		margin-top: 180rpx;
		/* border: 1rpx solid black; */
		display: flex;
		flex-direction: column;
	}
	.item {
		width: 100%;
		display: flex; 
		justify-content: space-between; /* 内容左右分布 */
		padding: 45rpx 20rpx ; /* 上下左右边距 */
		font-size:40rpx;
		border-bottom: 1rpx solid #e3e3e3;
		/* gap: 20rpx; /* 间距 */
	}
	
	.item text {
		flex-shrink: 0; /* 防止text收缩 */
		margin-left: 50rpx;
		color:#b2b2b2;
		font-weight:500;
	}
	.item input ,
	.item span{
		width: 200rpx;
		margin-right: 70rpx;
		flex-grow: 1; /* 让input或span占据剩余空间 */
		text-align: right; /* 文本右对齐 */
		font-weight: 500;
		color:black;
	}
	.item input {
		border: none;
		font-size: 40rpx;
	}
	.item :last-child{
		border-bottom: none;
	}
</style>