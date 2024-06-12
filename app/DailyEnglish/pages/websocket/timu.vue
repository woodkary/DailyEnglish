<template>
	<view class="container">
		<image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
		<view class="progress-container">
			<view class="progress">
				{{ currentQuestionIndex + 1 }} / {{ questions.length }}
			</view>
			<view class="progress-bar" :style="{ width: progress + '%' }"></view>
		</view>
		<view class="question">
			{{ questions[currentQuestionIndex].word }}
		</view>
		<view class="phonetic">
			{{ questions[currentQuestionIndex].phonetic_us }}
		</view>
		<view class="options">
			<button v-for="(choice, choiceIndex) in questions[currentQuestionIndex].choices" :key="choiceIndex"
				@click="selectOption(choiceIndex)" :class="getOptionClass(choiceIndex)"
				:disabled="selectedOptions[currentQuestionIndex] !== null">
				{{ choice }}
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
		<view class="jump-group" @click="handleJump(questions[currentQuestionIndex])">
			<text class="link">加入生词本</text>
			<image class="jump-icon" src="../../static/jump.svg" />
		</view>
		<!-- <view class="jump-group2" @click="handleJump2">
			<text class="link">不认识，下一个</text>
		</view> -->
	</view>
</template>

<script>
	export default {
		data() {
			return {
				currentQuestionIndex: 0,
				selectedOptions: [], // 保存每道题选中选项的数组，初始化为null
				isCorrect: [], // 保存每道题是否答对的数组
				questions: [{
						word_id: 1,
						word: 'abandon',
						phonetic_us: '[ə\'bændən]',
						choices: ['1', '2', '2', '放弃'],
						answer: '放弃'
					},
					{
						word_id: 2,
						word: 'abandon',
						phonetic_us: '[ə\'bændən]',
						choices: ['1', '选项B', '选项C', '选项D'],
						answer: '1'
					},
					{
						word_id: 3,
						word: 'abandon2',
						phonetic_us: '[ə\'bændən]',
						choices: ['1', '选项B', '选项C', '选项D'],
						answer: '1'
					},
				],
			};
		},
		computed: {
			progress() {
				return ((this.currentQuestionIndex + 1) / this.questions.length) * 100;
			}
		},
		onLoad() {
			uni.request({
				url: 'http://localhost:8080/api/main/take_punch',
				method: 'get',
				header: {
					'Authorization': `Bearer ${uni.getStorageSync('token')}`
				},
				success: (res) => {
					this.questions = [];
					res.data.word_list.forEach((question) => {
						this.questions.push({
							word_id: question.word_id,
							word: question.word,
							phonetic_us: question.phonetic_us,
							choices: Object.values(question.word_question),
							answer: question.word_question[question.answer],
						});
					});
					this.selectedOptions = Array(this.questions.length).fill(null);
					this.isCorrect = Array(this.questions.length).fill(null);
					console.log(this.questions);
				},
			})
		},
		methods: {
			handleBack() {
				this.$router.back();
			},
			handleJump(question) {
				uni.request({
					url: 'http://localhost:8080/api/words/add_new_word',
					method: 'post',
					header: {
						'Authorization': `Bearer ${uni.getStorageSync('token')}`
					},
					data: {
						word_id: question.word_id,
					},
					success: (res) => {
						uni.showToast({
							title: '加入生词本成功',
							icon: 'none',
							duration: 2000,
						});
						uni.setStorageSync(question.word, true);
					},
				})
			},
			selectOption(choiceIndex) {
				const correctIndex = this.getCorrectOptionIndex();
				const isCorrect = choiceIndex === correctIndex;
				this.$set(this.selectedOptions, this.currentQuestionIndex, choiceIndex);
				this.$set(this.isCorrect, this.currentQuestionIndex, isCorrect);
			},
			getCorrectOptionIndex() {
				const correctAnswer = this.questions[this.currentQuestionIndex].answer;
				const choices = this.questions[this.currentQuestionIndex].choices;
				return choices.indexOf(correctAnswer);
			},
			getOptionClass(choiceIndex) {
				if (this.selectedOptions[this.currentQuestionIndex] !== null) {
					const correctIndex = this.getCorrectOptionIndex();
					if (choiceIndex === correctIndex) {
						return 'correct';
					} else if (choiceIndex === this.selectedOptions[this.currentQuestionIndex]) {
						return this.isCorrect[this.currentQuestionIndex] ? 'correct' : 'incorrect';
					}
				}
				return '';
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
			handleJump2() {
				if (this.currentQuestionIndex < this.questions.length - 1) {
					this.currentQuestionIndex++;
				}
			},

		},
		created() {
			// 初始化selectedOptions和isCorrect数组
			const length = this.questions.length;
			this.selectedOptions = Array(length).fill(null);
			this.isCorrect = Array(length).fill(null);
		}
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
		margin-top: 7rem;
	}

	.phonetic {
		font-size: 1.7rem;
		margin-bottom: 40px;
		text-decoration: underline;
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
		font-size: 20px;
	}

	.options button.correct {
		background-color: #4caf50;
		color: white;
	}

	.options button.incorrect {
		background-color: #f44336;
		color: white;
	}

	.navigation {
		position: absolute;
		display: flex;
		justify-content: space-between;
		gap: 19rem;
		background-color: transparent;
		top: 21rem;

	}

	.navigation button {
		background-color: transparent;
		border: none;
		height: 3rem;
		width: 3rem;
	}

	uni-button:after {
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
		right: 2rem;
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