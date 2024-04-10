<template>
  <view v-if="showPopup" class="popup-container">
    <view class="popup-content">
		<view>{{ popupMessage }}</view>
      <view class="close-button-container">
        <view class="button" @click="handleConfirm">取消</view>
        <view class="button" @click="handleCancel">删除</view>
      </view>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      showPopup: true,
      popupMessage: ""
    };
  },
  onLoad() {
    this.openPopup("确定要删除吗？");
  },
  methods: {
    openPopup(message) {
      this.popupMessage = message;
      this.showPopup = true;
    },
    closePopup() {
      this.showPopup = false;
    },
    handleConfirm() {
      console.log("确认删除");
      this.closePopup();
    },
    handleCancel() {
      console.log("取消删除");
      this.closePopup();
    }
  }
};
</script>

<style>
.popup-container {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
}

.popup-content {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 80%;
}

.close-button-container {
  display: flex;
  width: 112%;
  justify-content: space-between;
  margin-top: 20px;
  margin-left: -6%;
  margin-bottom: -6%;
}
.button {
  width: 45%; /* 设置按钮宽度为弹出框宽度的一半 */
  padding: 10px;
  text-align: center;
  cursor: pointer;
  color: #222;
}

.button:nth-child(2) {
  color: #e74c3c;
}
</style>
