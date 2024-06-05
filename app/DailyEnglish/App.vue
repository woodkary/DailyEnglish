<template>
  <view>
    <!-- 应用的其他内容 -->
    <router-link class="uni-nav" to="/pages/HelloWorld/HelloWorld">
      Go to HelloWorld
    </router-link>
    <!-- 渲染当前路由对应的组件 -->
    <router-view></router-view>
  </view>
</template>

<script>
import TabBar from './components/TabBar.vue' // 导入底部导航栏组件路径，根据实际路径修改

export default {
	components:{
		TabBar
	},
}
</script>

<style>
/* 每个页面公共样式 */
/* #ifdef H5 */ 
  uni-page-head { display: none; } 
  /* #endif */
  .uni-tabbar-bottom .uni-tabbar{
	      height: 70px;
  }
  .uni-tabbar__item {
    img {
      width: 27px;
      height: 27px;
	  margin-left:-2px;
    }
	
  }
</style>
