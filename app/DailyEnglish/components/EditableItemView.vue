<template>  
  <view class="item" @click="showModal = true">  
    <text>{{ label }}</text>  
    <input v-if="!showModal" :value="value" disabled placeholder="点击编辑" />  
    <view class="modal" v-if="showModal" @click.stop="handleOutsideClick($event)">  
	<view class="modal-content">
      <input v-model="inputValue" :placeholder="placeholder" />  
	  <view class="buttons">
		<button @click="cancelEdit" class="button-left">取消</button>  
		<button @click="confirmEdit" class="button-right">确定</button>
	  </view>
    </view>  
	</view>
  </view>  
</template>  
  
<script>  
export default {  
  props: {  
    label: String,  
    value: String,  
    placeholder: String  
  },  
  data() {  
    return {  
      showModal: false,  
      inputValue: this.value  
    };  
  },  
  methods: {  
    handleOutsideClick(event) {  
      // 如果点击的是模态框外部，则关闭模态框  
      if (!this.$el.contains(event.target)) {  
        this.showModal = false;  
      }  
    },  
    confirmEdit() {  
      this.$emit('update:value', this.inputValue);  
      this.showModal = false;  
    },  
    cancelEdit() {  
      this.showModal = false;  
      this.inputValue = this.value;  
    }  
  }  
};  
</script>  
  
<style scoped>  
.item {
		width: 100%;
		display: flex; 
		justify-content: space-between; /* 内容左右分布 */
		padding: 45rpx 20rpx ; /* 上下左右边距 */
		font-size:40rpx;
		border-bottom: 1rpx solid #e3e3e3;
		/* gap: 20rpx; /* 间距 */
	}
	
	.item text {
		flex-shrink: 0; /* 防止text收缩 */
		margin-left: 50rpx;
		color:#b2b2b2;
		font-weight:500;
	}
	.item input ,
	.item span{
		width: 200rpx;
		margin-right: 70rpx;
		flex-grow: 1; /* 让input或span占据剩余空间 */
		text-align: right; /* 文本右对齐 */
		font-weight: 5000;
		color:black;
	}
	.item input {
		border: none;
		font-size: 40rpx;
	}
	.item :last-child{
		border-bottom: none;
	}
.modal {  
  /* modal 的样式，例如位置、背景等 */  
  position: fixed;  
  top: 0;  
  left: 0;  
  width: 100%;  
  height: 100%;  
  background: rgba(0, 0, 0, 0.5);  
  display: flex;  
  justify-content: center;  
  align-items: center;  
  z-index: 1;
}  
.modal input{
	text-align: left; /* 文本左对齐 */
	width: 90%;
	margin-left: 10rpx;
	background-color: #f6f6f6;
	border: bone;
	font-size: 40rpx;
	padding: 10px;
}
.modal-content {  
  /* 使用 Flexbox 布局 */  
  width: 80%;
  display: flex;  
  flex-direction: column;  
  align-items: flex-start; /* 垂直对齐方式，这里设为顶部对齐 */  
  background-color: #fff; /* 背景色 */  
  padding: 16px; /* 内边距 */  
  /* 其他样式... */  
}  
.buttons {  
  /* 按钮的容器也使用 Flexbox 布局 */  
  width: 100%;
  display: flex;  
  justify-content: space-between; /* 两端对齐 */  
  margin-top: 16px; /* 与输入框之间的间距 */  
}  
.button-left{
	margin-left: 0;
}

.button-right {  
  /* 可选的，用于给右边的按钮添加额外的样式（如果需要） */
	margin-right: 0;
}

/* 其他样式 */  
</style>