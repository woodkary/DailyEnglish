<template>
  <view class="container">
    <input v-model="userId" placeholder="请输入用户ID" />
    <button @click="connectWebSocket">连接</button>
    
    <view v-if="taskToday">
      <text>书本学习: {{ taskToday.book_learning }}</text>
      <text>已学单词数: {{ taskToday.word_num_learned }}</text>
      <text>总单词数: {{ taskToday.word_num_total }}</text>
      <text>剩余天数: {{ taskToday.days_left }}</text>
      <text>打卡数: {{ taskToday.punch_num }}</text>
      <text>复习数: {{ taskToday.review_num }}</text>
      <text>是否打卡: {{ taskToday.ispunched ? '是' : '否' }}</text>
    </view>
  </view>
</template>

<script>
export default {
  data() {
    return {
      userId: '',
      taskToday: null,
      socket: null
    };
  },
  methods: {
    connectWebSocket() {
      if (this.userId === '') {
        uni.showToast({ title: '请输入用户ID', icon: 'none' });
        return;
      }

      this.socket = uni.connectSocket({
        url: 'ws://localhost:8080/api/punch/main_menu',
        success: () => {
          console.log('WebSocket连接成功');
        },
        fail: (err) => {
          console.error('WebSocket连接失败', err);
        }
      });

      this.socket.onOpen(() => {
        console.log('WebSocket已打开');
        this.socket.send({
          data: JSON.stringify({ userId: this.userId })
        });
      });

      this.socket.onMessage((res) => {
        const data = JSON.parse(res.data);
        if (data.code === 200) {
          this.taskToday = data.task_today;
        } else {
          uni.showToast({ title: data.msg, icon: 'none' });
        }
      });

      this.socket.onClose(() => {
        console.log('WebSocket已关闭');
      });

      this.socket.onError((err) => {
        console.error('WebSocket错误', err);
      });
    }
  }
};
</script>

<style scoped>
.container {
  padding: 20px;
}

input {
  width: 100%;
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
}

button {
  padding: 10px;
  background-color: #007aff;
  color: white;
  border: none;
  cursor: pointer;
}

button:disabled {
  background-color: #ccc;
}

text {
  display: block;
  margin: 5px 0;
}
</style>
