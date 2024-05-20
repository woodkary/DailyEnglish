<template>
	<view class="container">
		<text class="progress-text">{{ current }}/{{questions.length}}</text>

		<swiper class="question-container" :options="swiperOptions" :easing-function="'linear'" :duration="250"
			@before-change="swiperChange">
			<swiper-item v-for="(question, index) in questions" :key="index">
				<view class="text-info">
					<text class="number">1</text>
					<!-- 以上是题目序号 -->
					<text class="question">{{ question.question }}</text>
				</view>
				<view class="button-group">
					<div v-for="(choice, choiceIndex) in question.choices" :key="choiceIndex" class="choice-container">
						<button class="option" :class="getClass(choiceIndex)" @click="selectChoice(choiceIndex)">
							{{ getLabel(choiceIndex) }}
						</button>
						<span class="choice-content">{{ choice }}</span>
					</div>
					<button class="confirm">确认答案</button>
				</view>


			</swiper-item>
		</swiper>
		<view class="footer">
      <view v-for="(thisRowQuestions,index) in rows" :key="index" class="row">
        <button v-for="(thisRowQuestion,index) in thisRowQuestions" :key="index" class="option" :style="{margin:buttonMargin+'rpx'}">
          {{thisRowQuestion.index+1}}
        </button>
      </view>
			<view style="display: flex;white-space: nowrap;">
				<uni-countdown class="daojishi" :show-day="false" :hour="12" :minute="12" :second="12"
					:font-size="20" />
				<image src="/static/xuanxiang.svg" class="xuanxiangbtn"></image>
			</view>
			<view class="xuanxiang-container" >
				
				<button class="submit">直接交卷</button>
			</view>
		</view>
	</view>
</template>


<script>
	export default {
		data() {
			return {
				swiperOptions: {
					// 其他配置...
					allowTouchMove: true, // 允许触摸滑动
					preventClicksPropagation: true, // 阻止点击事件冒泡
					// 其他 Swiper 配置...
				},
				progress: 1, // 进度条的初始值
				current: 1, // 当前进度
				currentQuestionIndex: 0,
        questionButtonIndex: 0, // 当前题目的按钮序号

				questions: [
					// 题目和选项
					{
						question: `__ is your brother?
									-He is a doctor.`,

						choices: ['1', '2', '2', '放弃']
					},
					{
						question: 'abandon',

						choices: ['1', '选项B', '选项C', '选项D']
					},
					{
						question: 'abandon2',

						choices: ['1', '选项B', '选项C', '选项D']
					},
					// ...更多题目
				], // 这里可以根据需要修改选项内容
				selectedChoice: '', // 用于存储用户选择的答案
				realAnswer: [
					'放弃', '选项B', '选项C' // 正确答案
				],
        maxButtonsPerRow: 6, // 每行的最大元素个数
        buttonMargin: 35, // 元素间隔
			}
		},
    computed: {
      //这是每一行的按钮，其中最多有maxButtonsPerRow个
      rows() {
        const rows = [];
        for (let i = 0; i < this.questions.length; i += this.maxButtonsPerRow) {
          let thisRowQuestions=[];
          for(let j=i;j<i+this.maxButtonsPerRow;j++){
            thisRowQuestions.push({
              index:j,
              question:this.questions[j]
            });
          }
          rows.push(thisRowQuestions);
        }
        return rows;
      },
    },
		methods: {

			handleJump() {

				uni.request({
					url: 'xxvcav',
					method: 'post',
					data: {
						//data
					},
					success: (res) => {
						//success
					},
				})
			},
			swiperChange(event) {
				const current = event.detail.current;
				const source = event.detail.source; // "touch" 或 "autoplay" 或 "pagination"

				// 仅当用户通过触摸滑动时处理
				if (source === 'touch') {
					// 判断滑动方向
					if (current > this.currentQuestionIndex) {
						// 左滑
						this.currentQuestionIndex = current;
					} else if (current < this.currentQuestionIndex) {
						// 右滑，防止切换
						// 可以使用swiper的scrollTo方法回到原来的位置
						this.$refs.swiper.scrollTo(this.currentQuestionIndex, 0, false);
					}
				}
			},
			selectChoice(index) {
				console.log(index);
				// 选择答案
			},

			preventSelect(event) {
				// 阻止长按事件的默认行为
				event.preventDefault();
				// 在这里可以添加长按的额外逻辑，比如显示一个提示框
			},
			getClass(index) {
				// 根据选中状态和答案正确与否返回相应的样式类
				if (this.selectedChoice) {
					console.log(this.currentQuestionIndex);
					if (this.questions[this.currentQuestionIndex].choices[index] === this.selectedChoice) {
						return this.questions[this.currentQuestionIndex].choices[index] === this.realAnswer ? 'correct' :
							'incorrect';
					}
				}
				return '';
			},
			getClass(choiceIndex) {
				// Your logic to return class based on choiceIndex
				return 'some-class-based-on-index'; // Placeholder
			},
			selectChoice(choiceIndex) {
				// Your logic to handle choice selection
				console.log(`Choice ${choiceIndex + 1} selected`);
			},
			getLabel(choiceIndex) {
				const labels = ['A', 'B', 'C', 'D'];
				return labels[choiceIndex];
			}

		}
	}
</script>


<style>
	@font-face {
		font-family: "pingfang";
		src: url('@/static/PingFang Medium_downcc.otf');
	}
  .row {
    display: flex;
    flex-wrap: wrap;
  }

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
		background-color: white;
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


	.text-info {
		position: absolute;
		top: 2rem;
		left: 7%;
		text-align: center;
		overflow: auto;
		height: auto;
		display: flex;
	}

	.number {
		position: relative;
		font-size: 1rem;
	}

	.question {
		position: relative;
		font-size: 1.5rem;
		margin-top: 1rem;
		white-space: pre-line;
		font-family: "pingfang";
	}

	.button-group {
		display: flex;
		flex-direction: column;
		gap: 2rem;
		margin-top: 10rem;
	}

	.choice-container {
		display: flex;
		align-items: center;
		margin-bottom: 10px;
	}

	.choice-content {
		flex-grow: 1;
		margin-left: 12px;
		font-size: 26px;
		font-family: "pingfang";
	}

	.option {
		box-shadow: 0 0 0 1px #aaa39b;
		width: 2rem;
		height: 2rem;
		border-radius: 2rem;
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: rgba(255, 255, 255, 0.2);
		cursor: pointer;
		font-size: 27px;
		margin-left: 10%;
		background-color: white;

	}

	.option.selected {
		background-color: #597dea;
		color: white;

	}

	.confirm {
		width: 80%;
		background-color: #2f5eed;
		font-size: 1.2rem;
		font-family: "pingfang";
		color: white;
		border-radius: 2rem;
		height: 3rem;
	}

	.question-container {
		width: 100%;
		height: 90%;
	}

	.footer {
		bottom: 0;
		position: fixed;
		width: 100%;
		display: flex;
		flex-direction: column;
		.xuanxiangbtn {
			width: 35px;
			height: 35px;
			margin-left: 260px;
		}

		.daojishi {
			margin-left: 10px;
		}
	}
	
	.xuanxiang-container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		border-top: 1px solid #e6e6e6;
		margin-top: 10px;
		.submit{
			margin-top:20px;
			width: 80%;
			background-color: #2f5eed;
			font-size: 1.2rem;
			font-family: "pingfang";
			color: white;
			border-radius: 0.5rem;
			height: 3rem;
		}
	}
</style>