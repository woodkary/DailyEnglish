<!--
 * @Date: 2024-04-02 18:56:15
-->
<template>
  <view class="container">
	  <text class="progress-text">{{ current }}/100</text>
	  <view class="progress-container">
	      <view class="progress-bar" :style="{ width: progress + '%' }"></view>
	    </view>

    <image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
    <view class="text-info">
      <text class="word">abandon</text>
      <text class="phonetic">[ə'bændən]</text>
    </view>
    <view class="button-group">
      <button class="button">A</button>
      <button class="button">B</button>
      <button class="button">C</button>
      <button class="button">D</button>
    </view>
    <view class="jump-group" @click="handleJump">
      <text class="link" >加入生词本</text>
      <image class="jump-icon" src="../../static/jump.svg" />
      
    </view>
  </view>
</template>


<script>
export default {
	 data() {
	    return {
	      progress: 70 ,// 进度条的初始值
		  current: 70
	    }
	  },
  methods: {
    handleBack() {
      // 处理返回按钮点击事件
      this.$router.back();
      // 例如：uni.navigateBack();
    },
    handleJump() {
      // 处理跳转链接点击事件
      this.$router.push('../index/index');
      // 例如：uni.navigateTo({ url: '/pages/OtherPage.vue' });
    },
	updateProgress(value) {
	      // 更新进度条的方法，value 是 0 到 100 之间的数值
	      if (value >= 0 && value <= 100) {
	        this.progress = value;
			this.current = value;
	      } else {
	        console.error('进度值必须在 0 到 100 之间');
	      }
	    }
  }
}
</script>


<style>
.container {
	
  display: flex; /*flex布局 */
  flex-direction: column; /*垂直布局 */
  align-items: center;  /*水平居中 */
  justify-content: center;  /*垂直居中 */
  height: 100vh;  /*占满整个屏幕 */
  overflow: hidden;  /*隐藏溢出部分 不能滚动*/
  background-image: linear-gradient(-190deg, #fff669 0%, #ecf1f1 50%, #d6f8f7 100%);
}

.progress-container {
position: absolute;
  width: 70%; 
  height: 0.5rem;
  top:1.5rem;
  margin-left:5rem;
  left:1rem;
  background: cadetblue;
  border:0.1rem solid #000;
  border-radius: 0.5rem;
  display: flex;
    align-items: center; /* 垂直居中 */

  z-index: 100;
}

.progress-bar {
  height: 100%; 
  background-color: #00ff00; 
}
.progress-text {
  position: absolute;
  color: #333; /* 文本颜色 */
  font-size: 0.8rem; /* 文本大小 */
   top: 1.3rem;	
   left: 3rem;
}

.back-icon {
  width: 2rem;  /*图标宽度 */
  height: 2rem; /*图标高度 */
  position: absolute; /*绝对定位 */
  top: 0.8rem;  /*距离顶部20px */
  left: 0.5rem; /*距离左侧20px */
  cursor: pointer;  /*鼠标移上去显示小手 */
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
  position:relative;
  font-size: 3rem;  /*字体大小 */
  font-weight: bold;  /*加粗 */
  margin-bottom: 1rem; /*调整与phonetic之间的距离 */
}

.phonetic {
  display: block; /* 在新行显示 */
  font-size: 1.7rem;  /*字体大小 */
  color: #666;  /*字体颜色 */
}

.button-group {
  display: flex;
  flex-direction: column;
  gap: 2rem;
  margin-top: 9rem;
}

.button {
  border:0.01rem solid #000;
  width: 17rem;
  height: 3.7rem;
  border-radius: 2rem;/*圆角 */
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(255, 255, 255, 0.2); ;/*背景透明度 */
}

.jump-group {
  position: fixed;/*固定定位 */
  bottom: 4rem;
  right: 2rem;
  display: flex; /* 使用 flexbox 布局 */
  font-size: 1rem;
}

.link {
	width: 5rem;
	height: 2rem;
  cursor: pointer;
}

.jump-icon {
  width: 1rem;
  height: 1rem!important;
  margin-left: 0.5rem;
  margin-top: 0.2rem;
}


</style>
