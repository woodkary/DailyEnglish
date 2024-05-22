<template>
	<view class="container">
		<text class="progress-text">{{ current }}/{{questions.length}}</text>

		<swiper class="question-container" :options="swiperOptions" :easing-function="'linear'" :duration="250"
			@before-change="swiperChange">
			<swiper-item v-for="(question, index) in questions" :key="index">
				<view class="text-info">
					<text class="number">{{ index + 1 }}</text>
					<!-- 以上是题目序号 -->
					<text class="question">{{ question.question }}</text>
				</view>
				<view class="button-group">
					<div v-for="(choice, choiceIndex) in question.choices" :key="choiceIndex" class="choice-container">
						<button class="option" :class="{ 'active': choiceIndex === question.activeButtonIndex }" @click="selectChoice(choiceIndex,index)">
						    {{ getLabel(choiceIndex) }}
						</button>

						<span class="choice-content">{{ choice }}</span>
					</div>
<!--					<button class="confirm" @click="finishQuestion(index)">确认答案</button>-->
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
						:class="{ 'finished': isFinished[thisRowQuestion.question.question_id], 'selected': thisRowQuestion.index === current }"
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
				current: 0, // 当前进度
				currentQuestionIndex: 0,
				selectedIndex: -1, // 当前题目的选中按钮序号
				questionButtonIndex: 0, // 当前题目的按钮序号
				isShow: false, //是否显示全部题目
				questions: [
					// 题目和选项
					{
						question_id: 1,
						question: `__ is your brother?
									-He is a doctor.`,
            activeButtonIndex: null, // 用于存储当前激活的按钮索引
						choices: ['1', '2', '2', '放弃']
					},
					{
						question_id: 2,
						question: 'abandon',
            activeButtonIndex: null, // 用于存储当前激活的按钮索引
						choices: ['1', '选项B', '选项C', '选项D']
					},
					{
						question_id: 3,
						question: 'abandon2',
            activeButtonIndex: null, // 用于存储当前激活的按钮索引
						choices: ['1', '选项B', '选项C', '选项D']
					},
					// ...更多题目
				], // 这里可以根据需要修改选项内容
				realAnswer: [
					'放弃', '选项B', '选项C' // 正确答案
				],
				maxButtonsPerRow: 6, // 每行的最大元素个数
				buttonMargin: 35, // 元素间隔
				isCorrects: {
					1: false,
					2: false,
					3: false,
				},
				isFinished: {
					1: false,
					2: false,
					3: false
				}, // 是否完成答题
				hasShownSubmitPrompt: false, // 是否已显示提交提示
			}
		},
    onLoad(event){
      let exam_id=parseInt(event.exam_id);
      uni.request({
        url: '/api/exams/getExamQuestions',
        method: 'POST',
        data: {
          exam_id: exam_id
        },
        header: {
          'Authorization': `Bearer ${uni.getStorageSync('token')}`
        },
        success: (res) => {
          //todo 获取所有题目信息
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
				if(this.selectedIndex==-1){
					return;
				}
				// 记录用户的答案
				let selectedChoice = this.questions[index].choices[this.selectedIndex];
				console.log("第"+index+"题你选择了" + selectedChoice);

				if (!this.isFinished[this.questions[index].question_id]) {
					// 保存是否完成到 map 中
					this.isFinished[this.questions[index].question_id] = true;

					// 更新当前题目索引
					this.currentQuestionIndex++;
					this.current++;
				}
				this.selectedIndex=-1;

				// 检查是否完成所有题目
				if (this.currentQuestionIndex === this.questions.length && !this.hasShownSubmitPrompt) {
					this.hasShownSubmitPrompt = true; // 设置为已显示提交提示
					uni.showModal({
						title: '提示',
						content: this.isAllFinished() ? '您已完成全部题目，是否确认提交' : '您还有题目未完成，是否确认提交',
						showCancel: true,
						success: (res) => {
							if (res.confirm) {
								this.handleJump();
							}
						}
					});
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

			submitExam() {
				uni.showModal({
					title: '提示',
					content: this.isAllFinished() ? '您已完成全部题目，是否确认提交' : '您还有题目未完成，是否确认提交',
					showCancel: true,
					success: (res) => {
						if (res.confirm) {
							this.handleJump();
						}
					}
				})

			},
			handleJump() {
				uni.navigateTo({
					url: '/pages/finishexam/finishexam?progress=' + this.getProgress()
				});
				//todo 提交考试结果到服务器
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
					// 更新当前题目索引
					this.currentQuestionIndex = current;
					this.selectedIndex = -1; // 重置选中的选项，确保选项渲染正确
				}
			},
			
			selectChoice(index,currentQuestionIndex) {
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
        if(this.questions[currentQuestionIndex].activeButtonIndex===index){
          this.questions[currentQuestionIndex].activeButtonIndex=null;
        }else{
          // 否则，设置当前点击的按钮为激活状态
          this.questions[currentQuestionIndex].activeButtonIndex=index;
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
			}

		}
	}
</script>


<style>
	@import '@/pages/exam/exam.css';

	@font-face {
		font-family: "pingfang";
		src: url('@/static/PingFang Medium_downcc.otf');
	}
  .active{
    background-color: #e74c3c;
  }
</style>