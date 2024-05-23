<template>
	<view>
		<view class="title-container">
			<image class="back-icon" src="../../static/back.svg"></image>
			<span>第1/20题</span>
		</view>
		<view class="question-container">
			<view style="width: max-content;">
				<span style="font-size: 24px;">题目</span>
				<hr style="border: 0;border-top: 5px solid #456DE7;height: 0;border-radius: 2.5px;">
			</view>
			<view class="question">
				<span>—is your brother?</span>
				<span>—He is a doctor.</span>
				<span class="answer">A. What</span>
				<span class="answer">B. Who</span>
				<span class="answer">C. Where</span>
				<span class="answer">D. How</span>
			</view>
		</view>
		<hr style="border: 0;border-bottom: 8px solid #e3e3e3;height: 0;">
		<view class="answer-container">
			<view style="width: max-content;">
				<span style="font-size: 24px;">答案</span>
				<hr style="border: 0;border-top: 5px solid #456DE7;height: 0;border-radius: 2.5px;">
			</view>
			<span class="true-answer">正确答案：A</span>
			<span class="your-answer">您的答案：B</span>
		</view>
		<hr style="border: 0;border-bottom: 8px solid #e3e3e3;height: 0;">
		<view class="explain-container">
			<CollapsibleView ref="collapsibleView1" :is-expanded="isExpanded">
				<template #header>
					<view class="header" @click="toggleCollapse1">
						<view style="width: max-content;">
							<span style="font-size: 24px;">解析</span>
							<hr style="border: 0;border-top: 5px solid #456DE7;height: 0;border-radius: 2.5px;">
						</view>
						<image :src="imageSrc" :class="['up-arrow', { 'rotated': isExpanded }]"></image>
					</view>
				</template>
				<span>在回答“What"开头的问句时，要注意回答的内容应该是描述事物的性质或身份，如职业、颜色、形状等。因此该题选A。</span>
			</CollapsibleView>
		</view>
	</view>
</template>

<script>
	import CollapsibleView from '../../components/CollapsibleView.vue';

	export default {
		components: {
			CollapsibleView
		},
		data() {
			return {
				isExpanded: false // 初始状态为折叠
			};
		},
		computed: {
			imageSrc() {
				return this.isExpanded ? '../../static/up.png' : '../../static/up.png'; // 可根据状态切换不同的图片
			}
		},
		methods: {
			toggleCollapse1() {
				this.isExpanded = !this.isExpanded; // 切换展开/折叠状态
			}
		}
	}
</script>

<style>
	.title-container {
		display: flex;
		line-height: 80rpx;
	}

	.back-icon {
		width: 60rpx;
		margin-right: 80rpx;
		height: 80rpx;
		margin-top: 20rpx;
		margin-bottom: 20rpx;
	}

	.title-container span {
		font-family: 'Source Han Sans', 'Microsoft YaHei', Arial, sans-serif;
		font-weight: bold;
		font-size: 18px;
		margin-top: 20rpx;
		margin-left: 160rpx;
	}

	.question-container {
		display: flex;
		flex-direction: column;
		/* 确保题目和答案分为两行 */
		gap: 10px;
		/* 添加元素之间的间距 */
		margin: 20px 40px;
		/* 为每个问题容器添加底部间距 */
	}

	.separator {
		border: 0;
		border-top: 5px solid #456DE7;
		height: 0;
		flex-grow: 1;
		/* 让水平线占据剩余的空间 */
		margin: 0 10px;
		/* 在水平线两侧添加间距 */
	}

	.question {
		font-size: 22px;
		margin: 5px 10px;
	}

	.answer {
		display: block;
		/* 每个答案占据一行 */
		margin: 5px 10px;
		/* 添加上下间距 */
		font-size: 22px;
	}

	.answer-container {
		gap: 10px;
		/* 添加元素之间的间距 */
		margin: 10px 40px;
		/* 为每个问题容器添加底部间距 */
		padding: 20px 0;
	}

	.true-answer {
		font-size: 20px;
		color: #456DE7;
		margin-right: 30px;
	}

	.your-answer {
		font-size: 20px;
		color: #EE4141;
	}

	.explain-container {
		/* 添加元素之间的间距 */
		margin: 10px 20px;
		/* 为每个问题容器添加底部间距 */
		padding: 20px 0;
	}

	.explain-container span {
		font-size: 20px;
	}

	.up-arrow {
		width: 1.5rem;
		height: 1rem;
		transition: transform 0.3s ease;
	}

	.header {
		display: flex;
		justify-content: space-between;
		/* 将元素分别放置在左右两侧 */
		align-items: center;
		/* 垂直居中 */
	}

	.rotated {
		transform: rotate(180deg);
	}
</style>