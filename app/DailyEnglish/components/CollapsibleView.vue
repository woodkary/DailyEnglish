<template>  
  <view class="collapsible-view" :class="{ 'collapsed': !isExpanded }">  
    <header @click="toggleContent">  
          <slot name="header"></slot>  
    </header> 
    <view class="content" :style="{ maxHeight: isExpanded ? '100rem' : '0', transition: 'max-height 0.3s ease' }">  
      <slot></slot> <!-- 内容插槽 -->  
    </view>  
  </view>  
</template>  
  
<script>  
export default {  
  name: 'CollapsibleView',
	props: {  
	    word: {  
	      type: Object,  
	      required: true  
	    }  
	  },
  data() {  
    return {
		 word: {}, 
      isExpanded: true // 初始状态为展开  
    };  
  },  
  methods: {  
    toggleContent() {  
      this.isExpanded = !this.isExpanded; // 切换展开/缩起状态  
    },    
  }  
};  
</script>  
  
<style>  
.collapsible-view {
	width: 90%;
	margin: 0 auto;
	display: flex;  
	flex-direction: column; /* 设置为列方向，实现从上往下排列 */  
	align-items: stretch; /* 子项默认拉伸以填充整个容器 */ 
}  
  .content {  
	  overflow: hidden; /* 隐藏超出内容区域的内容 */
	   border-bottom: 2px solid #e6e6e6; /* 固定的底部边框 */ 
}  
</style>