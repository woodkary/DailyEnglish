<template>
  <view class="container">
    <image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
	<view class="vocabook" >
		<image class="vocabook-img" src="../../static/书.png"></image>
		<view class="vocabook-title">单词书1</view>
		<view class="vocabook-cnt">生词数：{{cnt}}</view>
		<view class="button-container"><button class="review" @click="Review">复习</button></view>
		<view class="button-container"><button class="export" @click="Export">导出</button></view>
	</view>
</view>
</template>

<script>
export default {
  data() {
    return {
      words: ['moral', 'abandon', 'banana', 'apple', 'cat', 'dog'] ,// 单词列表
	  cnt: 0,
    }
  },
  methods: {
    handleBack() {
      uni.navigateBack()
    },
	Review(){
		;//预计跳转到Review页面，后续再写，内容是复习生词
	},
	Export(){
		;//预计导出生词，后续再写
		
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
  width: 2rem;  /*图标宽度 */
  height: 2rem; /*图标高度 */
  position: absolute; /*绝对定位 */
  top: 0.8rem;  /*距离顶部20px */
  left: 0.5rem; /*距离左侧20px */
  cursor: pointer;  /*鼠标移上去显示小手 */
}

.vocabook {
	width:100%;/*宽度100%*/
	display:flex;/*flex布局 */
	border-bottom: thick groove #ffff00;/*改一下颜色*/
	height: 10rem;
	margin-top: -35rem;
	
}
.vocabook-img{
	width: 6rem;
	height: 7rem;
	margin-left: 2rem;
	/* 在父容器中垂直居中 */
	align-self: center;
}

.vocabook-title{
	font-size:1rem;
	margin-left: 2rem;
	margin-top: 1.7rem;
	margin-bottom: 17rem;
}
.vocabook-cnt{
	font-size:0.9rem;
	margin-left: -3.7rem;
	margin-top: 3.8rem;
	margin-bottom: 17rem;
}
.button-container{
	display: flex;
	  justify-content: center;/*水平居中*/
	  align-items: center;/*垂直居中*/
	}
.review{
	margin-left: -4rem;
	margin-top: 5rem;
	border: 1px solid #000;
	border-radius: 5px;
	width: 4rem;
	height: 2rem;
	font-size: 1rem;
	/*文字居中*/
	text-align: center;
	line-height: 2rem;
}
.export{
	margin-left: 1rem;
	margin-top: 5rem;
	border: 1px solid #000;
	border-radius: 5px;
	width: 4rem;
	height: 2rem;
	font-size: 1rem;
	/*文字居中*/
	text-align: center;
	line-height: 2rem;
}
</style>

