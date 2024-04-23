<template>
	<view class="book-type">
		<view class="book-container">
			<image :src="imageSrc"></image>
			<view class="text-container">
				<span class="book-title">{{ title }}</span>
				<span class="discrip">{{ description }}</span>
				<span class="num">{{ totalWords }}</span>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		props: {
			title: String,
			description: String,
			imageSrc: String,
			totalWords: String
		}
	};
</script>

<style scoped>
	.book-type {
		margin-bottom: 20px;
		border-bottom: 2px solid gray;
		/* 下边框 */
		padding-bottom: 20px;
	}

	.book-container {
		display: flex;
		width: 100%;
		height: 220rpx;

	}

	.book-container image {
		width: 150rpx;
		height: 180rpx;
		margin: 10rpx 40rpx 5rpx 60rpx;
		border-radius: 5%;
	}

	.text-container {
		display: flex;
		flex-direction: column;
		/* 纵向排列 */
		align-items: flex-start;
		/* 左对齐 */
	}

	.book-title {
		font-size: 20px;
		margin-top: 10rpx;

	}

	.discrip {
		color: gray;
		margin-top: 10rpx;

	}

	.num {
		margin-top: 35rpx;

	}
</style>