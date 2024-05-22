<template>
  <view>
    <view class="title-container">
      <img class="back-icon" src="../../static/back.svg">
      <span>第{{ questionIndex + 1 }}/{{ questionNum }}题</span>
    </view>
    <view class="question-container">
      <view style="width: max-content;">
        <span style="font-size: 24px;">题目</span>
        <hr style="border: 0;border-top: 5px solid #456DE7;height: 0;">
      </view>
      <view class="question">
        <span v-for="(sentence, index) in currentQuestion.sentences" :key="index">{{ sentence }}</span>
        <view class="answers">
          <label v-for="(option, label) in currentQuestion.options" :key="label" class="answer">
<!--            <input type="radio" :value="label" v-model="currentQuestion.userAnswer">-->
            {{ label }}. {{ option }}
          </label>
        </view>
      </view>
    </view>
    <hr style="border: 0;border-bottom: 12px solid #BBBBBB;height: 0;">
    <view class="answer-container">
      <view style="width: max-content;">
        <span style="font-size: 24px;">答案</span>
        <hr style="border: 0;border-top: 5px solid #456DE7;height: 0;">
      </view>
      <span class="true-answer">正确答案：{{ currentQuestion.correctAnswer }}</span>
      <span class="your-answer">您的答案：{{ currentQuestion.userAnswer }}</span>
    </view>
    <hr style="border: 0;border-bottom: 12px solid #BBBBBB;height: 0;">
    <view class="explain-container">
      <view style="width: max-content;">
        <span style="font-size: 24px;">解析</span>
        <hr style="border: 0;border-top: 5px solid #456DE7;height: 0;">
      </view>
      <span>{{ currentQuestion.questionAnalysis }}</span>
    </view>
    <hr style="border: 0;border-bottom: 12px solid #BBBBBB;height: 0;">
  </view>
</template>

<script>
export default {
  data() {
    return {
      questionId: 1,
      questionIndex: 0,
      questionNum: 20,
      currentQuestion: {
        id: 1,
        sentences: [
          "-is your brother?",
          "-He is a doctor."
        ],
        options: {
          A: "What",
          B: "Who",
          C: "Where",
          D: "How"
        },
        correctAnswer: 'B',
        correctPoints: 0,
        totalPoints: 5,
        userAnswer: 'A',
        questionAnalysis: '在回答“What"开头的问句时，要注意回答的内容应该是描述事物的性质或身份，如职业、颜色、形状等。因此该题选A。'
      }
    };
  },
  onLoad(event){
    this.questionId = parseInt(event["questionId"]);
    this.questionIndex = parseInt(event["questionIndex"]);
    this.questionNum = parseInt(event["questionNum"]);
  }
};
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
		font-size: 24px;
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
		gap: 10px;
		/* 添加元素之间的间距 */
		margin: 10px 40px;
		/* 为每个问题容器添加底部间距 */
		padding: 20px 0;
	}

	.explain-container span {
		font-size: 20px;
	}
</style>