<template>
	<view>

		<view class="topbar">
			<span class="host-title">开始您的打卡计划</span>
			<span class="skip-container">
				<a href="" class="skip">跳过</a>
			</span>
		</view>
		<!-- 添加选项卡 -->
		<view class="tab-bar-container">
			<view class="tab-bar">
				<view class="tab-item" :class="{ 'active': activeTab === index }" v-for="(tab, index) in tabs"
					:key="index" @click="changeTab(index)">
					{{ tab }}
				</view>
			</view>
		</view>

		<!-- 添加按钮行 -->
		<view class="button-row">
			<view class="button" :class="{ 'active': activeButton === 'cet4' }" @click="scrollToCET4">
				四级
			</view>
			<view class="button" :class="{ 'active': activeButton === 'cet6' }" @click="scrollToCET6">
				六级
			</view>
			<!-- 添加其他按钮 -->
		</view>

		<view class="book-list">

			<view class="book-type" id="cet4">
				<view class="book-container">
					<image src="../../static/cet4.jpg"></image>
					<view class="text-container">
						<span class="book-title">四级词汇大全</span>
						<span class="discrip">四级最新考纲单词全收录，时候所有备考四级的同学</span>
						<span class="num">共4440词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet4.jpg"></image>
					<view class="text-container">
						<span class="book-title">四级高频</span>
						<span class="discrip">精选四级真题超高频词</span>
						<span class="num">共739词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet4.jpg"></image>
					<view class="text-container">
						<span class="book-title">四级高频</span>
						<span class="discrip">精选四级真题超高频词</span>
						<span class="num">共739词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet4.jpg"></image>
					<view class="text-container">
						<span class="book-title">四级高频</span>
						<span class="discrip">精选四级真题超高频词</span>
						<span class="num">共739词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet4.jpg"></image>
					<view class="text-container">
						<span class="book-title">四级高频</span>
						<span class="discrip">精选四级真题超高频词</span>
						<span class="num">共739词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet4.jpg"></image>
					<view class="text-container">
						<span class="book-title">四级高频</span>
						<span class="discrip">精选四级真题超高频词</span>
						<span class="num">共739词</span>
					</view>
				</view>

			</view>


			<view class="book-type" id="cet6">

				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级词汇大全</span>
						<span class="discrip">六级最新考纲单词全收录，时候所有备考六级的同学</span>
						<span class="num">共6204词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级核心（过考版）</span>
						<span class="discrip">精选六级真题超高频词</span>
						<span class="num">共2551词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级核心（过考版）</span>
						<span class="discrip">精选六级真题超高频词</span>
						<span class="num">共2551词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级核心（过考版）</span>
						<span class="discrip">精选六级真题超高频词</span>
						<span class="num">共2551词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级核心（过考版）</span>
						<span class="discrip">精选六级真题超高频词</span>
						<span class="num">共2551词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级核心（过考版）</span>
						<span class="discrip">精选六级真题超高频词</span>
						<span class="num">共2551词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级核心（过考版）</span>
						<span class="discrip">精选六级真题超高频词</span>
						<span class="num">共2551词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级核心（过考版）</span>
						<span class="discrip">精选六级真题超高频词</span>
						<span class="num">共2551词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级核心（过考版）</span>
						<span class="discrip">精选六级真题超高频词</span>
						<span class="num">共2551词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级核心（过考版）</span>
						<span class="discrip">精选六级真题超高频词</span>
						<span class="num">共2551词</span>
					</view>
				</view>
				<view class="book-container">
					<image src="../../static/cet6.jpg"></image>
					<view class="text-container">
						<span class="book-title">六级核心（过考版）</span>
						<span class="discrip">精选六级真题超高频词</span>
						<span class="num">共2551词</span>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				activeTab: 0, // 默认选中第一个选项卡
				tabs: ["全部", "大学", "高中", "初中", "小学", "留学", "其他"], // 选项卡数组
				activeButton: null // 记录当前活跃的按钮
			};
		},
		methods: {
			changeTab(tabNumber) {
				this.activeTab = tabNumber;
			},
			scrollToCET4() {
				// 点击按钮，滚动到四级部分
				this.scrollToElement('cet4');
				this.activeButton = 'cet4'; // 设置活跃的按钮为四级按钮
			},
			scrollToCET6() {
				// 点击按钮，滚动到六级部分
				this.scrollToElement('cet6');
				this.activeButton = 'cet6'; // 设置活跃的按钮为六级按钮
			},
			scrollToElement(id) {
				const el = document.getElementById(id);
				if (el) {
					const offset = 110; // 与顶部保持的距离
					const scrollPosition = el.offsetTop - offset;
					window.scrollTo({
						top: scrollPosition,
						behavior: 'smooth'
					});
				}
			}
		}
	};
</script>

<style>
	body {
		background-color: ;
		text-align: center;
	}

	.topbar {
		position: sticky;
		top: 0;
		z-index: 1000;
		/* 确保在最顶层 */
		display: flex;
		justify-content: space-between;
		z-index: 1000;
		/* 确保在最顶层 */
		/* 将元素分散对齐 */
		align-items: center;
		/* 垂直居中 */
		background-color: white;
		width: 100%;
		padding: 10px;

	}

	.host-title {
		font-family: Cambria, Cochin, Georgia, Times, 'Times New Roman', serif;
		font-size: 45rpx;
		width: 100%;
	}

	.skip-container {
		display: flex;

	}

	.skip {
		margin-right: 20px;
		margin-bottom: 0;
		text-decoration: none;
		white-space: nowrap;
		/* 取消下划线 */
		font-size: 16px;
		/* 适当调整字体大小 */
	}


	.tab-bar-container {
		background-color: white;
		position: sticky;
		top: 50px;
		/* 顶部导航栏的高度 */
		z-index: 1000;
		/* 确保在最顶层 */
		overflow-x: overlay;
		/* 隐藏滚动条，但保留滑动功能 */
		white-space: nowrap;
		/* 不换行 */
		border-bottom: 0.5rpx solid grey;
	}

	.tab-bar-container::-webkit-scrollbar {
		display: none;
		/* 隐藏滚动条 */
	}

	.tab-bar {
		display: inline-block;
		/* 内联块元素 */
	}

	.tab-item {
		display: inline-block;
		/* 内联块元素 */
		padding: 10px 20px;
		color: gray;
		/* 初始字体颜色 */
		cursor: pointer;
		font-size: 35rpx;
	}

	.tab-item.active {
		color: black;
		/* 选中时字体颜色 */
	}

	.tab-item.active::after {
		content: '';
		display: block;

		width: 50%;
		height: 3px;
		background-color: blue;
		/* 选中时底边框颜色 */
		margin: 0 auto;
		/* 将下划线水平居中 */
	}

	.button-row {
		position: sticky;
		background-color: white;
		top: 100px;
		/* 顶部导航栏和选项卡的高度 */
		z-index: 1000;
		/* 确保在最顶层 */
		display: flex;
		justify-content: center;
		margin-top: 20px;
	}

	.button {
		margin: 0 10px;
		padding: 5px 10px;
		cursor: pointer;
		background-color: #f8f8f8;
		color: gray;
		border: 1px solid #ccc;
		border-radius: 4px;
	}

	.button.active {
		background-color: blue;
		color: white;
	}

	.book-type-span {
		float: left;
		font-size: 25px;
		margin-left: 50rpx;
	}

	.book-container {
		display: flex;
		width: 100%;
		height: 220rpx;
		align-items: center;
		/* 垂直居中 */
	}

	.book-container image {

		width: 150rpx;
		height: 180rpx;
		margin: 10rpx 40rpx 5rpx 60rpx;
		border-radius: 5%;
	}

	.text-container {
		width: calc(100% - 260px);
		/* 调整宽度，确保留足够的空间给图片 */
		display: flex;
		flex-direction: column;
		align-items: flex-start;
	}

	.book-title {
		font-size: 20px;
		margin-top: 10rpx;
		white-space: nowrap;
	}

	.discrip {
		color: gray;
		margin-top: 10rpx;
		overflow: hidden;
		white-space: nowrap;

	}

	.num {
		margin-top: 35rpx;

	}
</style>