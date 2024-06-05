<template>
	<view class="container">
		<text class="progress-text">{{ current }}/{{questions.length}}</text>

		<swiper class="question-container" :options="swiperOptions" :easing-function="'linear'" :duration="500"
			:current="currentQuestionIndex" @before-change="swiperChange">
			<swiper-item v-for="(question, index) in questions" :key="index">
				<view class="text-info">
					<text class="number">{{ index + 1 }}</text>
					<!-- 以上是题目序号 -->
					    <text class="question" v-html="highlightUnderline(question.question)"></text>

				</view>
				<!--        如果是单选题，则显示选项按钮，否则显示输入框-->
				<view class="button-group" v-show="question.question_type==1">
					<div v-for="(choice, choiceIndex) in question.choices" :key="choiceIndex" class="choice-container">
						<button class="option" :class="{ 'active': choiceIndex === question.activeButtonIndex }"
							@click="selectChoice(choiceIndex,currentQuestionIndex)">
							{{ getLabel(choiceIndex) }}
						</button>

						<span class="choice-content">{{ choice }}</span>
					</div>
					<!--					<button class="confirm" @click="finishQuestion(index)">确认答案</button>-->
				</view>
				<!--        TODO 填空题的输入框的样式-->
				<view class="input-group" v-show="question.question_type==2">
					<span class="prompt">请注意大小写</span>
					<!-- <input type="text" placeholder="请输入答案" v-model="currentFillAnswer"
						@blur="inputFillAnswer(currentQuestionIndex)"
						@confirm="nextQuestionForFills(currentQuestionIndex)" /> -->
					<view class="input-items">
						<span class="item-index">1</span>
						<InputContainer class="input-item" placeholder="请输入答案" v-model="currentFillAnswer"
							@blur="inputFillAnswer(currentQuestionIndex)"
							@confirm="nextQuestionForFills(currentQuestionIndex)" />
					</view>
				</view>
			</swiper-item>
		</swiper>
		<view class="footer">

			<view style="display: flex;white-space: nowrap;">
				<uni-countdown class="daojishi" :show-day="false" :hour="12" :minute="12" :second="12"
					:font-size="20" />
				<image src="/static/xuanxiang.svg" class="xuanxiangbtn" @click="showQuestions"></image>
			</view>


			<view class="xuanxiang-container" v-show="isShow">
				<view v-for="(thisRowQuestions,rowIndex) in rows" :key="rowIndex" class="row">
					<button v-for="(thisRowQuestion,index) in thisRowQuestions" :key="index" class="option"
						@click="setCurrentQuestionIndexByQuestionId(thisRowQuestion.question.question_id)"
						:class="{ 'finished': isFinished[thisRowQuestion.question.question_id]}"
						:style="{margin:buttonMargin+'rpx'}">
						{{thisRowQuestion.index+1}}
					</button>
				</view>
				<button class="submit" @click="submitExam">直接交卷</button>
			</view>

		</view>
	</view>
</template>


<script>
	import InputContainer from '@/components/examinput.vue';
	export default {
		components: {
			InputContainer
		},
		data() {
			return {
				swiperOptions: {
					// 其他配置...
					allowTouchMove: true, // 允许触摸滑动
					preventClicksPropagation: true, // 阻止点击事件冒泡
					// 其他 Swiper 配置...
				},
				exam_id: null,
				exam_name: null,
				progress: 1, // 进度条的初始值
				current: 0, // 当前进度
				currentQuestionIndex: 0, //当前正在做的题目序号
				selectedIndex: -1, // 当前题目的选中按钮序号
				questionButtonIndex: 0, // 当前题目的按钮序号
				isShow: false, //是否显示全部题目
				currentFillAnswer: null, //当前填空题输入的答案
				questions: [
					// 题目和选项
					{
						question_id: 1,
						question_type: 1, //1单选2填空
						question: `__ is your brother?
									-He is a doctor.`,
						activeButtonIndex: null, // 用于存储当前激活的按钮索引
						choices: ['1', '2', '2', '放弃'],
						fullScore: 5 // 题目满分
					},
					{
						question_id: 2,
						question_type: 1, //1单选2填空
						question: 'abandon',
						activeButtonIndex: null, // 用于存储当前激活的按钮索引
						choices: ['1', '选项B', '选项C', '选项D'],
						fullScore: 5 // 题目满分
					},
					{
						question_id: 3,
						question_type: 1, //1单选2填空
						question: 'abandon2',
						activeButtonIndex: null, // 用于存储当前激活的按钮索引
						choices: ['1', '选项B', '选项C', '选项D'],
						fullScore: 5 // 题目满分
					},
					{
						question_id: 4,
						question_type: 2, //1单选2填空
						question: '__ is your brother?',
						activeButtonIndex: null, // 用于存储当前激活的按钮索引
						choices: null,
						fullScore: 5 // 题目满分
					}
					// ...更多题目
				], // 这里可以根据需要修改选项内容
				realAnswer: [
					'放弃', '选项B', '选项C', 'Who' // 正确答案
				],
				maxButtonsPerRow: 5, // 每行的最大元素个数
				buttonMargin: 35, // 元素间隔
				selectedChoiceAndScore: {
					/*//key为question_id
					1: {
            selectedChoice: null, // 用于存储当前选择的选项，或者是输入的答案
            score: 0 // 用于存储当前题目的分数
          },
					2: {
            selectedChoice: null, // 用于存储当前选择的选项，或者是输入的答案
            score: 0 // 用于存储当前题目的分数
          },
					3: {
            selectedChoice: null, // 用于存储当前选择的选项，或者是输入的答案
            score: 0 // 用于存储当前题目的分数
          },*/
				},
				isFinished: {
					1: true,
					2: true,
					3: true
				}, // 是否完成答题
				hasShownSubmitPrompt: false, // 是否已显示提交提示
				correctAnswers: 0, // 正确答案数
			}
		},
		onLoad(event) {
			let exam_id = parseInt(event.exam_id);
			this.exam_name = event.name;
			this.exam_id = exam_id;
			uni.request({
				url: 'http://localhost:8080/api/exams/take_examination',
				method: 'POST',
				data: {
					exam_id: exam_id
				},
				header: {
					'Authorization': `Bearer ${uni.getStorageSync('token')}`
				},
				success: (res) => {
            //token失效
            if(res.statusCode === 401){
              uni.removeStorageSync('token');
              uni.showToast({
                title: '登录已过期，请重新登录',
                icon: 'none',
                duration: 2000
              });
              uni.navigateTo({
                url: '../login/login'
              });
            }
					//todo 获取所有题目信息
					let questionAndAnswer = this.transformQuestions(res.data.question_list);
					this.questions = questionAndAnswer.questions;
					this.realAnswer = questionAndAnswer.realAnswer;
					questionAndAnswer.questions.forEach((question, index) => {
						this.isFinished[question.question_id] = false;
						this.selectedChoiceAndScore[question.question_id] = {
							selectedChoice: null, // 用于存储当前选择的选项
							score: 0 // 用于存储当前题目的分数
						};
					});
				},
				fail: (res) => {
					uni.showToast({
						title: '获取题目失败',
						icon: 'none'
					});
				}
			});
		},
		computed: {
			//这是每一行的按钮，其中最多有maxButtonsPerRow个
			rows() {
				const rows = [];
				for (let i = 0; i < this.questions.length; i += this.maxButtonsPerRow) {
					let thisRowQuestions = [];
					for (let j = i; j < i + this.maxButtonsPerRow && j < this.questions.length; j++) {
						thisRowQuestions.push({
							//题目序号
							index: j,
							//题目
							question: this.questions[j],
						});
					}
					rows.push(thisRowQuestions);
				}
				return rows;
			},

		},
		methods: {
			//输入填空题的逻辑
			inputFillAnswer(currentQuestionIndex) {
				//判断当前输入的答案是否正确
				let correct = this.currentFillAnswer === this.realAnswer[currentQuestionIndex];
				// 记录用户的答案
				this.selectedChoiceAndScore[this.questions[currentQuestionIndex].question_id].selectedChoice = this
					.currentFillAnswer;
				if (correct) {
					this.selectedChoiceAndScore[this.questions[currentQuestionIndex].question_id].score = this.questions[
						currentQuestionIndex].fullScore;
					this.correctAnswers++;
				} else {
					this.selectedChoiceAndScore[this.questions[currentQuestionIndex].question_id].score = 0;
				}
			},
			nextQuestionForFills(currentQuestionIndex) {
				if (!this.isFinished[this.questions[currentQuestionIndex].question_id]) {
					// 保存是否完成到 map 中
					this.isFinished[this.questions[currentQuestionIndex].question_id] = true;

					// 更新当前题目索引
					this.currentQuestionIndex++;
					this.current++;
				}
				this.selectedIndex = -1;
				this.currentFillAnswer = null;
				// 检查是否完成所有题目
				if (this.currentQuestionIndex === this.questions.length && !this.hasShownSubmitPrompt) {
					this.currentQuestionIndex--; // 回退到上一题
					this.hasShownSubmitPrompt = true; // 设置为已显示提交提示
					this.submitExam();
				} else {
					// 切换到下一题
					this.swiperChange({
						detail: {
							current: this.currentQuestionIndex,
							source: 'touch'
						}
					});
				}
			},
			transformQuestions(questionList) {
				let questions = [];
				let realAnswer = [];

				questionList.forEach((item, index) => {
					// 为每个问题创建一个新的对象，并添加到 questions 数组中
					questions.push({
						question_id: item.question_id,
						question_type: item.question_type,
						question: item.question_content,
						activeButtonIndex: null, // 初始化激活按钮索引
						choices: item.question_choices,
						fullScore: item.full_score // 题目满分
					});

					// 将正确答案添加到 realAnswer 数组中
					realAnswer.push(item.question_answer);
				});

				return {
					questions: questions,
					realAnswer: realAnswer
				};
			},
			setCurrentQuestionIndexByQuestionId(question_id) {
				console.log("setCurrentQuestionIndexByQuestionId:" + question_id);
				for (let i = 0; i < this.questions.length; i++) {
					if (this.questions[i].question_id === question_id) {
						this.currentQuestionIndex = i;
						break;
					}
				}
			},
			isAllFinished() {
				let allFinished = true;
				for (let key in this.isFinished) {
					if (!this.isFinished[key]) {
						allFinished = false;
						break;
					}
				}
				return allFinished;
			},
			getProgress() {
				let progress = 0;
				for (let key in this.isFinished) {
					if (this.isFinished[key]) {
						progress++;
					}
				}
				return progress / this.questions.length * 100;
			},
			finishQuestion(index) {
				if (this.selectedIndex == -1) {
					return;
				}
				// 记录用户的答案
				let selectedChoice = this.questions[index].choices[this.selectedIndex];
				console.log("第" + index + "题你选择了" + selectedChoice);

				// 更新当前题目的选择和分数
				let question_id = this.questions[index].question_id; //获取当前题目的id
				//将当前题目的选择和分数保存到selectedChoiceAndScore中
				this.selectedChoiceAndScore[question_id].selectedChoice = this.selectedIndex;
				if (this.selectedIndex === this.realAnswer[index]) {

					// 如果选择正确，则加满分
					this.selectedChoiceAndScore[question_id].score = this.questions[index].fullScore;
					this.correctAnswers++;
				} else {
					// 如果选择错误，则扣除分数
					this.selectedChoiceAndScore[question_id].score = 0;
				}

				if (!this.isFinished[this.questions[index].question_id]) {
					// 保存是否完成到 map 中
					this.isFinished[this.questions[index].question_id] = true;

					// 更新当前题目索引
					this.currentQuestionIndex++;
					this.current++;
				}
				this.selectedIndex = -1;
				this.currentFillAnswer = null;

				// 检查是否完成所有题目
				if (this.currentQuestionIndex === this.questions.length && !this.hasShownSubmitPrompt) {
					this.currentQuestionIndex--; // 回退到上一题
					this.hasShownSubmitPrompt = true; // 设置为已显示提交提示
					this.submitExam();
				} else {
					// 切换到下一题
					this.swiperChange({
						detail: {
							current: this.currentQuestionIndex,
							source: 'touch'
						}
					});
				}
			},
			getTotalScore() {
				let totalScore = 0;
				for (let key in this.selectedChoiceAndScore) {
					totalScore += this.selectedChoiceAndScore[key].score;
				}
				return totalScore;
			},
			submitExam() {
				uni.showModal({
					title: '提示',
					content: this.isAllFinished() ? '您已完成全部题目，是否确认提交' : '您还有题目未完成，是否确认提交',
					showCancel: true,
					success: (res) => {
            //token失效
            if(res.statusCode === 401){
              uni.removeStorageSync('token');
              uni.showToast({
                title: '登录已过期，请重新登录',
                icon: 'none',
                duration: 2000
              });
              uni.navigateTo({
                url: '../login/login'
              });
            }
						if (res.confirm) {
							//计算并保存考试结果到本地
							let examResult = {
								exam_id: this.exam_id,
								examTitle: this.exam_name,
								score: this.getTotalScore(), //考试总分
								totalQuestions: this.questions.length, //总题目数
								correctAnswers: this.correctAnswers, //正确答案数
							};
							console.log(examResult);
							uni.setStorageSync('examResult', examResult);
							//todo 提交考试结果到服务器
							uni.request({
								url: `/api/exams/submitExamResult`,
								method: 'POST',
								data: {
									selectedChoiceAndScore: this.selectedChoiceAndScore,
									exam_id: this.exam_id,
								},
								header: {
									'Authorization': `Bearer ${uni.getStorageSync('token')}`
								},
								success: (res) => {
            //token失效
            if(res.statusCode === 401){
              uni.removeStorageSync('token');
              uni.showToast({
                title: '登录已过期，请重新登录',
                icon: 'none',
                duration: 2000
              });
              uni.navigateTo({
                url: '../login/login'
              });
            }
									uni.showToast({
										title: '提交成功',
										icon: 'none'
									});
									this.handleJump();
								},
								fail: (res) => {
									uni.showToast({
										title: '提交失败',
										icon: 'none'
									});
								}
							});
						}
					}
				})

			},
			handleJump() {
				uni.navigateTo({
					url: '/pages/finishexam/finishexam?progress=' + this.getProgress()
				});
			},
			swiperChange(event) {
				const current = event.detail.current;
				const source = event.detail.source; // "touch" 或 "autoplay" 或 "pagination"

				// 仅当用户通过触摸滑动时处理
				if (source === 'touch') {
					// 更新当前题目索引
					this.currentQuestionIndex = current;
					this.selectedIndex = -1; // 重置选中的选项，确保选项渲染正确
				}
			},

			selectChoice(index, currentQuestionIndex) {
				/*// 如果当前点击的按钮已经是激活状态，则移除激活状态
				if (this.activeButtonIndex === index) {
				  this.activeButtonIndex = null;
				} else {
				  // 否则，设置当前点击的按钮为激活状态
				  this.activeButtonIndex = index;
				}*/
				// 记录用户的答案
				this.selectedIndex = index;
				//获取当前题目的word_id
				let question_id = this.questions[currentQuestionIndex].question_id;
				// 如果当前点击的按钮已经是激活状态，则移除激活状态
				if (this.questions[currentQuestionIndex].activeButtonIndex === index) {
					this.questions[currentQuestionIndex].activeButtonIndex = null;
				} else {
					// 否则，设置当前点击的按钮为激活状态
					this.questions[currentQuestionIndex].activeButtonIndex = index;
				}
				//直接提交当前题目的选择
				this.finishQuestion(currentQuestionIndex)

			},
			getLabel(choiceIndex) {
				const labels = ['A', 'B', 'C', 'D'];
				return labels[choiceIndex];
			},
			showQuestions() {
				this.isShow = !this.isShow;
			},
			 highlightUnderline(text) {
			      // 使用正则表达式匹配下划线，并将其替换为红色下划线
			      return text.replace(/__/g, '<span class="red-underline">__</span>');
			    },

		}
	}
</script>


<style>
	@import '@/pages/exam/exam.css';

	@font-face {
		font-family: "pingfang";
		src: url('@/static/PingFang Medium_downcc.otf');
	}

	.active {
		background-color: #e74c3c;
		color: white;
	}
	.red-underline {
	  color: red; /* 设置红色 */
	}
</style>