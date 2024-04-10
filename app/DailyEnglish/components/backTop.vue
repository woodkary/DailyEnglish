<!--
 * @Date: 2024-04-10 16:25:26
 * 建议学习vocab界面的方法，那边写好了达到某个滑动距离显示返回顶部按钮的功能
-->
<template>
  <view class="backTop" @click="backToTop">
    <image src="../static/backTop.svg"></image>
  </view>
</template>

<script>
  export default {
    name: 'backTop',
    methods: {
      backToTop() {
        uni.pageScrollTo({
          scrollTop: 0,
          duration: 300//滚动动画的时长
        });
      }
    }
  }
</script>

<style scoped>
  .backTop {
    position: fixed;
    bottom:80px;
    right: 20px;
    width: 40px;
    height: 40px;
    border-radius: 50%;
    background-color: white;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
	  z-index: 999;
  }
</style>
