<template>
	<view class="container">
		<image class="back-icon" src="../../static/back.svg"></image>
		<view class="edit">
			<view class="title">
				<text>修改个人信息</text>
			</view>
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
		margin-top: 80rpx;
		width:100%;
	}
	.tilte {
		font-size:40rpx;
		
	}
	.photo {
		width: 200rpx;
		height: 200rpx;
		margin-left: 50%;
		margin-top: 40rpx;
	}
</style>