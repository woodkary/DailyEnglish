<template>
  <view class="word-card">
    <view class="word-container">
      <view class="word">{{ word }}</view>
      <view class="pronunciation-meaning" @click="togglePronunciationAndMeaning">
        {{ showPronunciationMeaning ? `${pronunciation} - ${meaning}` : '点击显示音标和释义' }}
      </view>
    </view>
    <view class="action-container">
      <image class="read-button" src="../../static/read-icon.svg" @click="readWord"></image>
      <view class="review-count">
        <p>复习次数:</p>
        <p style="color: #e74c3c">{{reviewCount }}</p>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  props: {
    word: String,
    pronunciation: String,
    meaning: String,
    reviewCount: Number
  },
  // 组件的其他逻辑
  data() {
    return {
      word: this.word,
      pronunciation: this.pronunciation,
      meaning: this.meaning,
      showPronunciationMeaning: false,
      reviewCount: this.reviewCount
    };
  },
  methods: {
    togglePronunciationAndMeaning() {
      this.showPronunciationMeaning =!this.showPronunciationMeaning;
    },
    readWord() {
      // 实现系统朗读的逻辑
      // 例如：uni.$audio.play('path_to_audio/apple.mp3');
    }
  }
};
</script>

<style>
.word-card {
  display: flex;
  justify-content: space-between;
  border: 3rpx solid #ccc;
  padding: 20rpx;
  width: 100%;
  background: #fff;
}

.word-container {
  display: flex;
  flex-direction: column;
}

.action-container {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
}

.word {
  font-size: 45rpx;
  font-weight: bold;
  margin-bottom: 10rpx;
  margin-left: 30rpx;
}

.pronunciation-meaning {
  background-color: lightgrey;
  padding: 10rpx;
  font-size: 40rpx;
  font-family: Calibri, sans-serif;
  cursor: pointer;
  margin-left: 30rpx;
}

.read-button {
  width: 48rpx;
  height: 48rpx;
  cursor: pointer;
  margin-right: 20rpx;
}

.review-count {
  margin-top: 30rpx;
  display: flex;
  font-size: 35rpx;
  justify-content: space-between;
  align-items: center;
  margin-right: 20rpx;
}
</style>
