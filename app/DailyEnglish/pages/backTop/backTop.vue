<template>
  <view>
    <!-- 页面内容 -->
    <backtop v-if="showBackTop"></backtop>
  </view>
</template>

<script>
  import backTop from '@/components/backTop.vue'
  export default {
    components: {
      backTop
    },
    data() {
      return {
        showBackTop: false
      }
    },
    onPageScroll(e) {
      // 当页面滚动到100px时显示返回顶部按钮
      this.showBackTop=true;
    }
  }
</script>
