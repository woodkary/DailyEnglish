<template>
	<view>
		
	</view>
</template>

<script>
	export default {
		data() {
			return {
        //写作任务列表
				writingTasks:[
          {
            title_id: 6,
            title: "小作文1",
            manager_name: "qwerty",
            word_num: "50~100",
            requirement: "请不要用真实姓名，使用“李明”代替",
            publish_date: "2024-06-15",
            grade: "小学"
          },
          {
            title_id: 7,
            title: "小作文2",
            manager_name: "asdfgh",
            word_num: "50~100",
            requirement: "无",
            publish_date: "2024-06-15",
            grade: "小学"
          },
          {
            title_id: 8,
            title: "小作文3",
            manager_name: "zxcvbn",
            word_num: "50~100",
            requirement: "无",
            publish_date: "2024-06-15",
            grade: "小学"
          }
				]
			}
		},
    onLoad() {
      //获取写作任务列表
      this.getWritingTasks()
    },
		methods: {
			getWritingTasks() {
        uni.request({
          url: 'http://localhost:8080/api/users/composition_mission',
          method: 'GET',
          header: {
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          },
          success: (res) => {
            console.log(res.data)
            if(res.statusCode === 200) {
              this.writingTasks = res.data.tasks
              console.log(this.writingTasks)
            }
          },
          fail: (err) => {
            console.log(err)
          }
        });
      }
		}
	}
</script>

<style>

</style>
