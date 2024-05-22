<template>
  <view>
    <view class="background">
      <span>试卷已提交</span>
      <image src="../../static/horse.svg" class="horse"></image>
    </view>
    <view class="center-container">
      <view class="exam-result">
        <h3 class="exam-title">{{ examTitle }}</h3>
        <span class="exam-score">{{ score }}</span><span style="color:#3FC681;font-size: 25px;">分</span>
        <br>
        <span class="exam-num">共{{ totalQuestions }}题</span><span class="true-num">答对<span :style="{color: correctAnswers === totalQuestions? '#3FC681' : '#FF4949'}">{{ correctAnswers }}</span>/{{ totalQuestions }}题</span>

        <view style="margin-left: 100px;margin-top: 20px;">
          <piaoyiProgressBar canvasId="progressCanvas4" :progress="progress" backgroundColor="#EFEFF4"
                             progressBackgroundColor="#07C160" :showText="true" textColor="#456DE7" :textSize="48" :height="20"
                             :isCircular="true" :diameter="200"></piaoyiProgressBar>
          <view class="bg"></view>
        </view>
        <view class="btn " @click="toDetail">考试详情</view>
      </view>

    </view>
    <view class="btn" @click="toHome">完成</view>
  </view>
</template>

<script>
import piaoyiProgressBar from '@/uni_modules/piaoyi-progress-bar/components/piaoyi-progress-bar/piaoyi-progress-bar.vue';
export default {
  components: {
    piaoyiProgressBar
  },
  data() {
    return {
      exam_id: 1,
      examTitle: '第一单元第一次小测',
      score: 95,
      totalQuestions: 20,
      correctAnswers: 19,
      //完成度百分比
      progress: 95
    };
  },
  onLoad(event) {
    this.progress=parseInt(event.progress);
    this.fetchData();
  },
  methods: {
    progressBarColor(progress){
      if(progress >= 90){
        return '#3FC681';
      }else if(progress >= 80) {
        return '#FFC107';
      }else if(progress >= 60){
        return '#FF4949';
      }else {
        return '#EFEFF4';
      }
    },
    //todo:exam_score的值不需要从后端获取，而是从本地缓存中获取
    fetchData() {
      let examResult=uni.getStorageSync("examResult");
      if(examResult){
        this.exam_id=examResult.exam_id;
        this.examTitle=examResult.examTitle;
        this.score=examResult.score;
        this.totalQuestions=examResult.totalQuestions;
        this.correctAnswers=examResult.correctAnswers;
        this.progress=parseInt(examResult.score/examResult.totalQuestions*100);
      }
    },
    toDetail() {
      uni.navigateTo({
        url: '/pages/exam_details/exam_details'
      });
    },
    toHome(){
      uni.reLaunch({
        url:'/pages/home/home'
      })
    }
  }
}
</script>

<style scoped>
.background {
  background-color: rgba(79, 252, 120, 0.57);
  display: flex;
  /* 使用 Flexbox 布局 */
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
  flex-direction: column;
  /* 垂直布局 */
  border: 0 solid #000;
  /* 添加边框，2px宽度，黑色 */
  border-radius: 0 0 30px 30px;
  /* 左下、右下圆角30度 */
  padding: 20px;
  /* 添加内边距 */
  z-index: 1;
}

.background span {
  margin-top: 60rpx;
  font-size: 24px;
  font-weight: bold;
}

.horse {
  width: 180px;
  height: 180px;
}

.center-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: -70px;
  z-index: 2;
}

.exam-result {
  width: 80%;
  text-align: center;
  background: #ffffff;
  /* 文本居中 */
  border: 2px solid #ccc;
  /* 添加边框，2px宽度，灰色 */
  box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.2);
  /* 添加阴影，水平和垂直偏移均为2px，模糊半径为5px，透明度为0.2 */
  padding: 20px;
  /* 添加内边距 */
  border-radius: 18px;
  z-index: 2;
}

.exam-title {
  margin: 35rpx 0;
}

.exam-score {
  color: #3FC681;
  font-size: 50px;
  margin-bottom: 35rpx;
}

.exam-num {
  color: #a7a7a7;
  font-size: 18px;
  margin-right: 20rpx;
}

.true-num {
  color: #a7a7a7;
  font-size: 18px;
}

.btn {
  background-color: #456DE7;
  /* 蓝色背景 */
  color: #fff;
  /* 白色文字 */
  padding: 10px 10px;
  /* 添加适当的内边距 */
  border-radius: 5px;
  /* 圆角 */
  margin: 20px 40px;
  /* 与 .circle 之间的间距 */
  text-align: center;
  /* 文本居中 */
  cursor: pointer;
  /* 鼠标指针样式 */
  font-size: 26px;
}
</style>