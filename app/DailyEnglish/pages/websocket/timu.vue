<template>
	<view class="container">
		<image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
		<view class="progress-container">
			<view class="progress">
				{{ currentQuestionIndex + 1 }} / {{ questions.length }}
			</view>
			<view class="progress-bar" :style="{ width:progress + '%' }"></view>
		</view>
		<view class="question">
			{{ questions[currentQuestionIndex].word }}
		</view>
		<view class="phonetic">
			{{ questions[currentQuestionIndex].phonetic_us }}
		</view>
		<view class="options">
			<button v-for="(option, index) in questions[currentQuestionIndex].options" :key="index"
				@click="selectOption(index)" :class="{ selected: selectedOptions[currentQuestionIndex] === index }">
				{{ option }}
			</button>
		</view>
		<view class="navigation">
			<button @click="prevQuestion" :disabled="currentQuestionIndex === 0">
				<image class="back-icon" src="../../static/back.svg"></image>
			</button>
			<button @click="nextQuestion" :disabled="selectedOptions[currentQuestionIndex] === null">
				<image class="back-icon" src="../../static/back.svg" style="transform: rotateY(180deg);"></image>
			</button>
		</view>
		<view class="jump-group" @click="handleJump(question)">
			<text class="link">加入生词本</text>
			<image class="jump-icon" src="../../static/jump.svg" />
		
		</view>
		<view class="jump-group2" @click="handleJump2">
			<text class="link">不认识，下一个</text>				
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				currentQuestionIndex: 0,
				selectedOptions: Array(10).fill(null), // 保存每道题选中选项的数组，初始化为null
				questions: [{
						word_id: 1,
						word: "题目1",
						phonetic_us: "音标",
						options: ["选项1", "选项2", "选项3", "选项4"],
					},
					{
						word_id: 2,
						word: "题目2",
						phonetic_us: "音标",
						options: ["选项1", "选项2", "选项3", "选项4"],
					},
					// 继续添加其他题目，总共10题
				],
			};
		},
		methods: {
			handleBack() {
				this.$router.back();
			},
			// handleJump(questions[currentQuestionIndex]) {
			// 	uni.request({
			// 		url: 'http://localhost:8080/api/words/add_new_word',
			// 		method: 'post',
			// 		header: {
			// 			'Authorization': `Bearer ${uni.getStorageSync('token')}`
			// 		},
			// 		data: {
			// 			word_id: questions[currentQuestionIndex].word_id,
			// 		},
			// 		success: (res) => {
			// 			uni.showToast({
			// 				title: '加入生词本成功',
			// 				icon: 'none',
			// 				duration: 2000,
			// 			});
			// 			uni.setStorageSync(word, true);
			// 		},
			// 	})
			// },
			selectOption(index) {
				this.$set(this.selectedOptions, this.currentQuestionIndex, index);
			},
			nextQuestion() {
				if (this.selectedOptions[this.currentQuestionIndex] !== null && this.currentQuestionIndex < this.questions
					.length - 1) {
					this.currentQuestionIndex++;
				}
			},
			prevQuestion() {
				if (this.currentQuestionIndex > 0) {
					this.currentQuestionIndex--;
				}
			},
		},
	};
</script>

<style>
	.container {
		padding: 20px;
		display: flex;
		flex-direction: column;
		align-items: center;
		
		height: 100vh;
		overflow: hidden;
		background-image: linear-gradient(-190deg, #fff669 0%, #ecf1f1 50%, #d6f8f7 100%);
	}

	.back-icon {
		width: 2rem;
		height: 2rem;
		position: absolute;
		top: 0.8rem;
		left: 0.5rem;
		cursor: pointer;
	}
	
	.progress-container {
		position: absolute;
		width: 70%;
		height: 0.5rem;
		top: 1.5rem;
		margin-left: 5rem;
		left: 1rem;
		background: cadetblue;
		border: 0.1rem solid #000;
		border-radius: 0.5rem;
		display: flex;
		align-items: center;
		z-index: 100;
	}
	
	.progress-bar {
		height: 100%;
		background-color: #00ff00;
	}

	.progress {
		font-size: 18px;
		margin-bottom: 10px;
		margin-top: 0.4rem;
		margin-left: -2.5rem;	
	}

	.question {
		position: relative;
		font-size: 3rem;
		font-weight: bold;
		margin-bottom: 20px;
		margin-top:7rem;
	}
	.phonetic {
		font-size: 1.7rem;
		margin-bottom: 40px;
	}

	.options {
		margin-bottom: 0px;
		gap: 2rem;
		display: flex;
		flex-direction: column;
	}

	.options button {
		border: 0.01rem solid #000;
		width: 17rem;
		height: 3.7rem;
		border-radius: 2rem;
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: rgba(255, 255, 255, 0.2);
		cursor: pointer;
		font-size:20px;
	}

	.options button.selected {
		background-color: #4caf50;
		color: white;
	}

	.navigation {
		position: absolute;
		display: flex;
		justify-content: space-between;
		gap: 19rem;
		background-color: transparent;
		top:21rem;
		
	}

	.navigation button {
		background-color: transparent;
		border: none;
		height: 3rem;
		width: 3rem;
	}
	uni-button:after{
		content: "";
		display: block;
		clear: both;
		background-color: transparent;
		border: none;
	}
	
	.jump-group {
		position: fixed;
		/*固定定位 */
		bottom: 6rem;
		left: 2rem;
		display: flex;
		/* 使用 flexbox 布局 */
		font-size: 1rem;
		.link {
			width: 5rem;
			height: 2rem;
			cursor: pointer;
		}
		
		.jump-icon {
			width: 1rem;
			height: 1rem !important;
			margin-left: 0.5rem;
			margin-top: 0.2rem;
		}
		
	}
	
	.jump-group2 {
		position: fixed;
		/*固定定位 */
		bottom: 6rem;
		right: 3rem;
		/*不换行 */
		maxlines: 1;
		display: flex;
		/* 使用 flexbox 布局 */
		font-size: 1rem;
		.link {
			width: 5rem;
			height: 2rem;
			cursor: pointer;
			white-space: nowrap;
			}
	}
</style>