<!--
 * @Date: 2024-04-02 18:56:15
-->
<template>
	<view class="container">
		<view class="question-container">

			<text class="progress-text">{{ current }}/{{questions.length}}</text>
			<view class="progress-container">
				<view class="progress-bar" :style="{ width:progress + '%' }"></view>
				<view class="progress-bar" :style="{ width:progress + '%'}"></view>
			</view>

			<image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
			<view class="text-info">
				<text class="word">{{ questions[currentQuestionIndex].word }}</text>
				<text class="phonetic">{{questions[currentQuestionIndex].phonetic}}</text>
			</view>

			<view class="button-group">
				<button class="option" v-for="(choice, index) in 
				questions[currentQuestionIndex].choices" 
				:key="index"	:class="getClass(index)" 
				@click="selectChoice(index)">{{ choice }}</button>
			</view>
			 <view class="button-group">
			    <button class="option" v-for="(choice, index) in questions[currentQuestionIndex].choices" :key="index" :ref="'option' + index" :class="getClass(index)" @click="selectChoice(index)">
			      {{ choice }}
			    </button>
			  </view>


				<text class="link">加入生词本</text>
				<image class="jump-icon" src="../../static/jump.svg" />

			</view>

		</view>
	</view>
</template>


<script>
	export default {
		data() {
			return {
				progress: 1, // 进度条的初始值
				current: 1, // 当前进度
				currentQuestionIndex:0,
				currentQuestionIndex: 0,

				questions: [
					// 题目和选项
					{
						word: 'abandon',
						phonetic: '[ə\'bændən]',
					},
					{
						word: 'abandon',
						phonetic: '[ə\'bændən]',
						choices: ['1', '选项B', '选项C', '选项D']
					},
					{
						word: 'abandon2',
						phonetic: '[ə\'bændən]',
						choices: ['1', '选项B', '选项C', '选项D']
					},
					// ...更多题目
				], // 这里可以根据需要修改选项内容
				selectedChoice: '', // 用于存储用户选择的答案
				realAnswer:[
				realAnswer: [
					'放弃', '选项B', '选项C' // 正确答案
				],
				

			}
		},
		methods: {
			handleBack() {
				// 处理返回按钮点击事件
				this.$router.back();
				// 例如：uni.navigateBack();
			},
			},
			updateProgressBar() {
				// 使得进度条增加1
				this.updateProgress(this.progress + 1);
			},
			updateProgress(value) {
				// 更新进度条的方法，value 是 0 到 100 之间的数值
				} else {
				}
			},
			selectChoice(index) {
				let selectedChoice = this.questions[this.currentQuestionIndex].choices[index];
				// 检查选中的答案是否正确
				if (selectedChoice === this.realAnswer[this.currentQuestionIndex]) {
					// 正确答案的逻辑
          let index=this.currentQuestionIndex++; // 切换到下一题
					this.$nextTick(() => {
						this.showCorrectAnswer(this.realAnswer[index]);
					});
          this.unlockSlice();
				} else {
          let currIndex=this.currentQuestionIndex;
					// 错误答案的逻辑
					this.$nextTick(() => {
						this.showIncorrectAnswer(index);
						this.showCorrectAnswer(this.realAnswer[currIndex]);
					});
				}
			    const selectedChoice = this.questions[this.currentQuestionIndex].choices[index];
			    // Check if the selected choice is correct
			    if (selectedChoice === this.realAnswer[this.currentQuestionIndex]) {
			        // Correct answer logic
			        this.showCorrectAnswer(index);
			        this.unlockSlice();
			    } else {
			        // Incorrect answer logic
			        this.showIncorrectAnswer(index);
			    }
			},
			showCorrectAnswer(answer) {
				// 找到正确答案的索引
				const correctIndex = this.questions[this.currentQuestionIndex].choices.indexOf(answer);
				// 应用正确答案的样式
				const correctButton = this.$refs[`option${correctIndex}`];
				if (correctButton) {
					correctButton.classList.add('correct');
				}

			showCorrectAnswer(index) {
			    // Apply correct answer style
			    const correctButton = this.$refs[`option${index}`];
			    if (correctButton) {
			        correctButton.classList.add('correct');
			    }
			},
			showIncorrectAnswer(index) {
				// 应用错误答案的样式
				const incorrectButton = this.$refs[`option${index}`];
				if (incorrectButton) {
					incorrectButton.classList.add('incorrect');
				}
			    // Apply incorrect answer style
			    const incorrectButton = this.$refs[`option${index}`];
			    if (incorrectButton) {
			        incorrectButton.classList.add('incorrect');
			    }
			},

			preventSelect(event) {
				// 阻止长按事件的默认行为
				event.preventDefault();
				// 在这里可以添加长按的额外逻辑，比如显示一个提示框
			},
			getClass(index) {
				// 根据选中状态和答案正确与否返回相应的样式类
				if (this.selectedChoice) {
					if (this.questions[this.currentQuestionIndex].choices[index] === this.selectedChoice) {
						return this.questions[this.currentQuestionIndex].choices[index] === this.realAnswer ? 'correct' : 'incorrect';
						return this.questions[this.currentQuestionIndex].choices[index] === this.realAnswer ? 'correct' :
							'incorrect';
					}
				}
				return '';
			},
			unlockSlice() {
				// 解锁下一题
				this.updateProgressBar();
				/*this.selectedChoice = '';*/
				this.$nextTick(() => {
					/*this.resetButtonStyle(); //重置按钮样式*/
				});

			},
		}
	}
</script>


<style>
	.container {

		display: flex;
		/*flex布局 */
		flex-direction: column;
		/*垂直布局 */
		align-items: center;
		/*水平居中 */
		justify-content: center;
		/*垂直居中 */
		height: 100vh;
		/*占满整个屏幕 */
		overflow: hidden;
		/*隐藏溢出部分 不能滚动*/
		background-image: linear-gradient(-190deg, #fff669 0%, #ecf1f1 50%, #d6f8f7 100%);
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
		/* 垂直居中 */

		z-index: 100;
	}

	.progress-bar {
		height: 100%;
		background-color: #00ff00;
	}

	.progress-text {
		position: absolute;
		color: #333;
		/* 文本颜色 */
		font-size: 0.8rem;
		/* 文本大小 */
		top: 1.3rem;
		left: 3rem;
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

	.text-info {
		position: absolute;
		top: 10rem;
		left: 50%;
		transform: translateX(-50%);
		text-align: center;
		overflow: auto;
		height: auto;
	}

	.word {
		position: relative;
		font-size: 3rem;
		/*字体大小 */
		font-weight: bold;
		/*加粗 */
		margin-bottom: 1rem;
		/*调整与phonetic之间的距离 */
	}

	.phonetic {
		display: block;
		/* 在新行显示 */
		font-size: 1.7rem;
		/*字体大小 */
		color: #666;
		/*字体颜色 */
	}

	.button-group {
		display: flex;
		flex-direction: column;
		gap: 2rem;
		margin-top: 9rem;
	}

	.option {
		border: 0.01rem solid #000;
		width: 17rem;
		height: 3.7rem;
		border-radius: 2rem;
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: rgba(255, 255, 255, 0.2);
		cursor: pointer;
	}

	.option.selected {
		background-color: #5c90ff;
		color: white;
	}

	/* 正确答案的样式 */
	.correct {
		border-color: green;
		color: green;
	}

	/* 错误答案的样式 */
	.incorrect {
		border-color: red;
		color: red;
	}

	.jump-group {
		position: fixed;
		/*固定定位 */
		bottom: 4rem;
		right: 2rem;
		display: flex;
		/* 使用 flexbox 布局 */
		font-size: 1rem;
	}

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
</style>