<template>
  <div class="toast" v-show="visible">
    {{ message }}
  </div>
</template>

<script>
export default {
  data() {
    return {
      visible: false,
      message: '',
      timeoutId: null,
    };
  },
  methods: {
    showToast(message, duration = 2000) {
      this.message = message;
            this.visible = true;
            clearTimeout(this.timeoutId);
            this.timeoutId = setTimeout(() => {
              this.visible = false;
            }, duration);
    },
	 show() {
	      // 这里可以定义其他的显示逻辑
	      this.showToast();
	    }
  },
};
</script>

<style>
.toast {
  position: fixed;/*固定定位，建议别改*/
  top: 20%;/*距离顶部，使用中修改*/
  left: 50%;/*居中，使用中修改*/
  transform: translate(-50%, -50%);/*可以修改*/
  background-color: black;/*背景色，使用中修改*/
  color: white;/*字体颜色，使用中修改*/
  padding: 10px 20px;/*内边距，使用中修改*/
  border-radius: 5px;/*圆角，使用中修改*/
  z-index: 9999;/*层级，不建议更改*/
  text-align: center;/*文字居中，不建议更改*/
}
</style>
