<template>
  <view>
    <view class="container">
      <view class="left-container">
        <h3 class="exam-number">{{ examName }}</h3>
        <p class="date">{{ formatDate(examDate) }}</p>
        <view style="display: flex; justify-content: flex-start; margin-top: 25px; margin-bottom: 20px;">
          <p class="question-number">共{{ questionNum }}题</p>
          <p class="correct-number">答对{{ correctNum }}/{{ questionNum }}题</p>
        </view>
      </view>
      <view class="right-container">
        <p class="point">{{ score }}</p>
        <p class="small-text">分</p>
      </view>
    </view>
    <view class="container2">
      <view class="header-container">
        <h3 style="font-size:24px">考试题目</h3>
        <view class="search">
          <img class="search-icon" src="/static/search.svg">
          <input v-model="searchQuery" placeholder="搜索">
        </view>
      </view>
      <view class="button-group">
        <button class="button1" @click="sortQuestions('order')">题目顺序</button>
        <button class="button2" @click="sortQuestions('reverse')">题目逆序</button>
        <button class="button3" @click="sortQuestions('score')">分数顺序</button>
        <button class="button4" @click="sortQuestions('scoreReverse')">分数逆序</button>
      </view>
    </view>
    <view v-for="(question, index) in filteredQuestions" :key="question.id" class="title-container">
      <view style="display: flex; justify-content: space-between; width: 90%;">
        <p class="title-number">{{ index + 1 }}.</p>
        <p :class="question.correctPoints>=question.totalPoints ? 'correct-title-point' : 'wrong-title-point'">{{ question.correctPoints }}/{{question.totalPoints}}</p>
      </view>
      <view v-for="(sentence, sentenceIndex) in question.sentences" :key="sentenceIndex" class="sentence-group">
        <p :class="sentence">{{ sentence }}</p>
      </view>
      <!-- 使用v-for遍历options对象 -->
      <view class="options-group">
        <p v-for="(optionContent, optionLabel) in question.options" :key="optionLabel" class="option">
          {{ optionLabel }}: {{ optionContent }}
        </p>
      </view>
      <view style="display: flex; justify-content: space-between; width: 90%; margin-top: 20px;">
        <p :class="question.correctPoints>=question.totalPoints ? 'correct-answer' : 'wrong-answer'">我的答案:{{ question.userAnswer }}</p>
        <p class="true-answer">正确答案:{{ question.correctAnswer }}</p>
      </view>
      <button class="big-button" @click="showAnalysis(question.id,index)">查看解析</button>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      examId: 1,
      examName: '第一单元第一次小测',
      examDate: '1970-01-01',
      questionNum: 20,
      correctNum: 19,
      score: 95,
      searchQuery: '',
      questions: [
        // ...题目数组
        {
          id: 1,
          sentences:[
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
        },
        {
          id: 2,
         sentences:[
            "-______is the nearest post office,please?",
            "-It's about half an hour's walk from here？"
          ],
          options: {
            A: "How far",
            B: "How long",
            C: "How often",
            D: "How soon"
          },
          correctAnswer: 'A',
          correctPoints: 5,
          totalPoints: 5,
          userAnswer: 'A',
          
        },
        {
          id: 3,
          sentences:[
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
        },
        {
          id: 4,
          sentences:[
            "-______is the nearest post office,please?",
            "-It's about half an hour's walk from here？"
          ],
          options: {
            A: "How far",
            B: "How long",
            C: "How often",
            D: "How soon"
          },
          correctAnswer: 'A',
          correctPoints: 5,
          totalPoints: 5,
          userAnswer: 'A',

        }
      ]
    };
  },
  computed: {
    filteredQuestions() {
      if (this.searchQuery) {
        // 根据搜索框内容过滤题目
        return this.questions.filter(question => question.sentence[1].includes(this.searchQuery));
      }
      return this.questions;
    }
  },
  onLoad(event){
    this.examId=event.exam_id;
    this.examName=event.exam_name;
    uni.request({
      url: '/api/exams/examination_details',
      data: {
        exam_id: this.examId
      },
      header: {
        'content-type': 'application/json',
        'Authorization': `Bearer ${uni.getStorageSync('token')}`
      },
      success: (res) => {
        this.examDate = res.data.exam_date;
        this.questionNum = res.data.question_num;
        this.correctNum = res.data.correct_num;
        this.score = res.data.score;
        this.questions = this.transformQuestions(res.data.questions);
      },
      fail: (res) => {
        uni.showToast({
          title: '获取考试详情失败',
          icon: 'none'
        });
        }
    });
  },
  methods: {
    // 转换函数
    transformQuestions(questions) {
      return questions.map(question => {
        return {
          id: question.question_id, // 使用原始的question_id作为id
          index: question.question_index, // 添加index属性
          sentences: question.question_decription.split("\n"), // 将题目描述按换行符分割成句子数组
          options: {
            A: question.choices["A"],
            B: question.choices["B"],
            C: question.choices["C"],
            D: question.choices["D"]
          },
          correctAnswer: question.correct_answer, // 正确答案
          correctPoints: question.score, // 正确得分
          totalPoints: question.full_score, // 总分
          userAnswer: question.my_answer // 用户答案
        };
      });
    },
    formatDate(date) {
      // 格式化日期
      return date.replace('-', '年').replace('-', '月').replace('-', '日');
    },
    sortQuestions(by) {
      // 根据不同的标准排序题目
      this.questions.sort((a, b) => {
        if (by === 'order') {
          return a.id - b.id;
        } else if (by ==='reverse') {
          return b.id - a.id;
        } else if (by ==='score') {
          return b.correctPoints - a.correctPoints;
        } else if (by ==='scoreReverse') {
          return a.correctPoints - b.correctPoints;
        }
      });
      // ...
    },
    showAnalysis(questionId,questionIndex) {
      // 显示题目解析
      uni.navigateTo({
        url: '/pages/questionDetail/questionDetail?questionId=' + questionId+'&questionIndex='+questionIndex+'&questionNum='+this.questionNum
      });
      // ...
    }
  }
};
</script>

<style>
	.container{
		width: 90%;
		margin: 0 auto;
		margin-top: 1rem;
		display: flex;  
		align-items: stretch; /* 拉伸子元素以匹配最高的子元素 */  
		justify-content: space-between; /* 水平对齐 */  
		border: 2px solid #e6e6e6;
		border-radius: 10px;
		padding: 15px;
		box-shadow: 0 0 0 2px rgb(247 127 0 / 10%);
	}
	.left-container{
		flex: 0 0 60%;
		display: flex;
		flex-direction: column;
		/* 设置为列方向，实现从上往下排列 */
		align-items: stretch;
		/* 子项默认拉伸以填充整个容器 */
	}
	.right-container {  
	  flex: 0 0 40%; /* 右边占40%宽度 */  
	  display: flex;  
	    flex-direction: column; /* 垂直方向 */  
	    justify-content: center; /* 垂直居中 */  
	    align-items: center; /* 水平居中 */  
	    text-align: center; /* 水平文本居中（对于单行文本）*/ 
	  /* 其他样式... */  
	} 
	.exam-number{
		margin-top: 10px;
	}
	.date{
		margin-top: 25px;
		color: #456de7;
		font-size: 20px;
	}
	.question-number{
		color: #a7a7a7;
		margin-right: 10px;
		font-size: 18px;
	}
	.correct-number{
		color: #a7a7a7;
		font-size: 18px;
	}
	.point{
		  font-size: 80px; /* 大字号的数字字体大小 */  
		  font-style: italic; /* 设置为斜体 */
			margin-left: -40px;
			color: #3FC681;
	}
	.small-text{
			font-style: italic; /* 设置为斜体 */ 
			 font-size: 26px;
			 margin-top: -40px;
			 margin-left: 80px;
			 color: #3FC681;
	}
	.container2 {  
	    width: 95%;  
	    margin: 0 auto;  
	    margin-top: 1rem;
		  margin-bottom: 1rem;
	    display: flex;  
	    flex-direction: column; /* 如果需要垂直堆叠 */  
	}  
	  
	.header-container {  
	    display: flex;  
	    align-items: center;  
	    justify-content: space-between; /* 根据需要，如果h3和search之间需要间隔 */
		  margin-bottom: 10px;
	}  
	  
	.search {  
	    flex: 1; /* 让.search占据剩余空间 */  
	    background-color: #ffffff;  
	    border-radius: 50rpx;  
	    display: flex;  
	    padding: 5rpx;  
	    align-items: center;  
	    /* 移除margin-left，或根据需要设置 */  
	    border: 1px solid #e6e6e6;
		box-shadow: 0 0 0 2px rgb(247 127 0 / 10%);
		margin-left: 30px;
	}  
	  
	.search-icon {  
	    width: 25px;  
	    height: 25px;  
	    margin-right: 20rpx;  
	}  
	  
	.search-input {  
	    flex: 1; /* 让输入框占据.search的剩余空间 */  
	    border: none;  
	    padding: 0 10rpx; /* 根据需要添加内边距 */  
	}
	.button-group{
		    display: flex;  
		    justify-content: flex-start; /* 让按钮从左边开始排列 */ 
			 width: 80%;
	}
	.button-group button {  
	    margin-right: 10px; /* 按钮之间的间距 */  
	    padding: 0 10px; /* 内边距，控制按钮的大小 */  
	    border: none; /* 移除默认的边框 */  
	    border-radius: 5px; /* 圆角 */  
	    background-color: #ededed; /* 背景色 */  
	    color: gray; /* 文字颜色 */  
	    font-size: 12px; /* 文字大小 */  
	    cursor: pointer; /* 鼠标悬停时变为小手形状 */  
	}
	.button-group .button1{
		color: #2e47e7;
	 }
	 .title-container{
		width: 90%;
		margin: 0 auto;
		 border: 2px solid #e6e6e6;
		 border-radius: 10px;
		 padding: 15px;
		 box-shadow: 0 0 0 2px rgb(247 127 0 / 10%);
		 margin-bottom: 20px;
	 }
	 .wrong-title-point{
		 font-size: 36px;
		font-style: italic; /* 设置为斜体 */ 
		 color: #ee4e41;
	 }
	 .correct-title-point{
		 font-size: 36px;
		 font-style: italic; /* 设置为斜体 */ 
		  color: #3fc681;
	 }
	 .sentence{
		 margin-left: 30px;
		 margin-bottom: 5px;
	 }
	 .answer{
	 		 margin-left: 30px;
	 }
	 .wrong-answer{
		 color: red;
	 }
	 .correct-answer{
	 		 color: #3fc681;
	 }
	 .true-answer{
		 color: #2e47e7;
	 }
	 .big-button{
		 width: 90%;
		 background-color: #2e47e7;
		 color: white;
		 margin-top: 20px;
		 font-size: 14px;
	 }
</style>
