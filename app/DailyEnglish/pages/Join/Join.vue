<template>
	<view class="container">
		<view class="head">
			<image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
			<span>加入团队</span>
			<button>确认</button>
		</view>

		<view class="input-wrapper">
			<input class="uni-input" placeholder="请输入团队码" v-model="inputClearValue" @input="clearInput" />
			<image class="uni-icon" src="../../static/not-done2.svg" v-if="showClearIcon" @click="clearIcon">
			</image>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				inputClearValue: '',
				showClearIcon: false
			}
		},
		methods: {
			handleBack() {
				uni.navigateBack({
					delta: 1
				})
			},
			clearInput() {
				this.showClearIcon = this.inputClearValue.length > 0;
			},
			clearIcon() {
				this.inputClearValue = '';
				this.showClearIcon = false;
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

	button {
		margin-left: 5rem;
		background-color: transparent;
		line-height: 2rem;
		height: 2rem;
		font-size: 21px;
		align-items: center;
		border: none;
		color:#456DE7;
	}

	button::after {
		border: none;
	}

	/* 
	input {
		width: 90%;
		background-color: rgb(227, 227, 227);
		margin-top: 2rem;
		height: 2.6rem;
	} */
	.input-wrapper {
		display: flex;
		padding-left: 16px;
		padding-right: 13px;
		padding-top: 16px;
		padding-bottom: 8px;
		margin-top: 2rem ;
		flex-direction: row;
		flex-wrap: nowrap;
		background-color: #e3e3e3;
		width: 80%;
		height: 2.2rem;
		border-radius: 5px;
	}

	.uni-input {
		height: 28px;
		font-size: 22px;
		padding: 0px;
		flex: 1;
		background-color: #e3e3e3;
		color: #000;
	}

	.uni-icon {
		width: 24px;
		height: 24px;
	}
</style>