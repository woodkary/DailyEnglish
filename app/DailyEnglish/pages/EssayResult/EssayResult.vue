<template>
	<view>
		<view class="head">
			<image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
			<span class="title">作文解析</span>
			<view class="composition-container">
				<span class="line1">题目：<span style="font-weight: normal;">{{ title }}</span> </span>
				<span class="line2">字数：<span style="font-weight: normal;">{{ word_cnt }}</span> </span>
				<span class="line3">要求： <span
						class="req">请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文</span></span>
			</view>
		</view>

		<view class="composition-tabs">
			<Tabs>
				<template v-slot:tab1-content>
					<view style="display: flex;">
						<div class="progress-ring">
							<svg class="progress-ring__svg" viewBox="0 0 120 120">
								<circle class="progress-ring__circle-bg" stroke="#ddd" stroke-width="13"
									fill="transparent" r="50" cx="60" cy="60" />
								<circle class="progress-ring__circle-fg" stroke="#44a0fb" stroke-width="13"
									fill="transparent" r="50" cx="60" cy="60" :style="{
					  strokeDasharray: circumference,
					  strokeDashoffset: offset,
					}" />
							</svg>
							<span class="score">{{ score }}</span>
						</div>
						<span class="evaluate">评价：{{ pingjia }}</span>
					</view>
				</template>
			</Tabs>
		</view>

		<view class="composition-content">
			<span class="article">{{ article }}</span>
		</view>

		<view class="evaluation">
			<span class="eva-title">逐句点评</span>
			<view class="eva-content">
				<view class="eva-item" v-for="(item, index) in evaluation" :key="index">
					<view class="eva-item-title">
						<span class="order">{{ item.order }}</span>
						<span>{{ item.sentence }}</span>
					</view>
					<view class="eva-item-content">{{ item.content }}</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	import Tabs from "../../components/Tabs.vue";
	export default {
		name: "App",
		components: {
			Tabs,
		},
		data() {
			return {
				title: '作文解析',
				word_cnt: "100~300",
				radius: 50, //半径
				circumference: 2 * Math.PI * 50, //周长
				score: 80,
				evaluation: [{
					order:1.1,
					sentence: 'fuck u kary',
					content: 'This essay is really well-written, with a clear structure and strong arguments. The author',
				}],

				pingjia: 'i哦部分的几天里你看加热别人叫我其他的v偶尔输入年份v看国足色加热别人叫我其他的v偶尔输入年份v看国足色弱问人家看望母亲吗',
				article: "I am happy to join with you today in what will go down in history as the greatest demonstration for freedom in the history of our nation.\nFive score years ago, a great American, in whose symbolic shadow we stand today, signed the Emancipation Proclamation. This momentous decree came as a great beacon light of hope to millions of Negro slaves who had been seared in the flames of withering injustice. It came as a joyous daybreak to end the long night of bad captivity.."
			}
		},
		methods: {
        handleBack() {
          uni.switchTab({
            url: '../user/user'
          });
        }
		},
		computed: {
			offset() {
				let progress = this.score / 100;
				return this.circumference * (1 - progress);
			},
		},
	}
</script>

<style>
	.head {
		width: 100%;
		position: relative;
		top: 0;
		left: 0;
		z-index: 100;
		height: 280px;
	}

	.back-icon {
		width: 2rem;
		height: 2rem;
		position: absolute;
		top: 0.8rem;
		left: 0.5rem;
		cursor: pointer;
	}

	.title {
		position: absolute;
		font-size: 1.2rem;
		font-weight: bold;
		text-align: center;
		top: 0.8rem;
		left: 40%;
	}

	.composition-container {
		position: absolute;
		top: 4rem;
		left: 5%;
		width: 90%;
		height: 170px;
		display: flex;
		flex-direction: column;
		border: 1.6px solid #1b6ef3;
		border-radius: 1.5rem;
	}

	.line1,
	.line2,
	.line3 {
		font-size: 1rem;
		margin-left: 1rem;
		font-weight: bold;
		margin-right: 1rem;
	}

	.line1 {
		margin-top: 0.7rem;
	}

	.line2,
	.line3 {
		margin-top: 0.2rem;
	}

	.line3 {
		max-height: 100px;
		min-height: 100px;
	}

	.req {
		font-weight: normal;
		display: block;
		line-height: 20px;
		max-height: 60px;
		/* 20px * 3 */
		overflow-y: auto;
	}

	.composition-tabs {
		position: relative;
		margin-top: 2px;
		border: 1px solid #44a0fb;
		border-radius: 1rem;
		width: 90%;
		margin-left: 5%;
		height: 150px;
	}

	.progress-ring {
		position: relative;
		min-width: 75px;
		min-height: 75px;
		background-color: white;
		margin-right: auto;
		margin-top: 10px;
	}

	.score {
		position: absolute;
		top: 45%;
		left: 50%;
		transform: translate(-50%, -50%);
		font-size: 1.7rem;
		color: #44a0fb;
	}

	.progress-ring__svg {
		transform: rotate(-90deg);
	}

	.progress-ring__circle-bg,
	.progress-ring__circle-fg {
		stroke-dasharray: 314.1592653589793;
	}

	.progress-ring__circle-fg {
		transition: stroke-dashoffset 0.35s;
	}

	.evaluate {
		position: relative;
		font-size: 1rem;
		margin-left: 1rem;
		margin-top: 10px;
		line-height: 1.2rem;
		max-height: 4.8rem;
		overflow-y: auto;
		scrollbar-width: none;
		/* Firefox */
		-ms-overflow-style: none;
		/* Internet Explorer 10+ */
	}

	.article,
	.evaluate::-webkit-scrollbar {
		/* WebKit */
		width: 0;
		height: 0;
	}

	@media screen and (max-width: 600px) {

		.article,
		.evaluate {
			scrollbar-width: none;
			/* Firefox */
			-ms-overflow-style: none;
			/* Internet Explorer 10+ */
		}

		.article,
		.evaluate::-webkit-scrollbar {
			/* WebKit */
			width: 0;
			height: 0;
		}
	}

	.composition-content {
		margin-top: 1rem;
		width: 90%;
		margin-left: 5%;
		border: 1px solid #44a0fb;
		border-radius: 1rem;
		padding: 10px;
		box-sizing: border-box;
		max-height: 250px;
		overflow-y: auto;
		scrollbar-width: none;
		/* Firefox */
		-ms-overflow-style: none;
		/* Internet Explorer 10+ */
	}

	.article {
		font-size: 1rem;
		white-space: pre-wrap;
		text-align: left;
		hyphens: auto;
		/*连字符*/
	}
	
	.evaluation {
		margin-top: 1rem;
		width: 90%;
		margin-left: 5%;
	}

	.eva-title{
		font-size: 1.2rem;
		margin-left: 0.5rem;
		/* border-bottom: 2px solid #44a0fb; */
	}

	.eva-item {
		margin-top: 0.4rem;
		width:100%;
		border:1px solid #44a0fb;
		display:flex;
		height: fit-content;
	}

	.eva-item-title {
		width:40%;
		border-right:1px solid #44a0fb;
	}
	
	.eva-item-content {
		width:60%;
	}
</style>