<!--
 * @Date: 2024-04-02 18:56:15
-->
<template>
  <view class="container">
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

.back-icon {
  width: 30px;  /*图标宽度 */
  height: 30px; /*图标高度 */
  position: absolute; /*绝对定位 */
  top: 20px;  /*距离顶部20px */
  left: 20px; /*距离左侧20px */
  cursor: pointer;  /*鼠标移上去显示小手 */
}

.text-info {
  position: absolute;
  top: 140px;
  left: 50%;
  transform: translateX(-50%);
  text-align: center;
  overflow: auto;
  height: auto;
}

.word {
  position:relative;
  font-size: 50px;  /*字体大小 */
  font-weight: bold;  /*加粗 */
  margin-bottom: 1px; /*调整与phonetic之间的距离 */
}

.phonetic {
  display: block; /* 在新行显示 */
  font-size: 30px;  /*字体大小 */
  color: #666;  /*字体颜色 */
}

.button-group {
  display: flex;
  flex-direction: column;
  gap: 40px;
  margin-top: 120px;
}

.button {
  border:1px solid #000;
  width: 280px;
  height: 60px;
  border-radius: 200px;/*圆角 */
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: rgba(255, 255, 255, 0.2); ;/*背景透明度 */
}

.jump-group {
  position: fixed;/*固定定位 */
  bottom: 40px;
  right: 20px;
  display: flex; /* 使用 flexbox 布局 */
  font-size: 20px;
}

.link {
  cursor: pointer;
}

.jump-icon {
  width: 22px;
  height: 30px!important;
  margin-left: 5px;
  
}


</style>
