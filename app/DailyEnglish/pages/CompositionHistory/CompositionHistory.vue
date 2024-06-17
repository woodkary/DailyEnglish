<template>
	<view>
		<view class="head">
			<image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
			<span class="title">我的写作</span>
		</view>
		<view class="task-list-container">
			<span class="task-list-title">写作任务</span>
			<view class="task-items" v-for="(task, index) in writingTasks" :key="index">
				<view class="task-item" v-if="writingTasks.length > 0">
					<view style="padding-bottom: 0.2rem; border-bottom: 2px solid #69c0ff">
						<view class="task-item-title"><span style="color: rgb(69, 109, 231)">[题目]</span><span
								style="margin-left: 1rem">{{ task.title }}</span></view>
					</view>
					<view class="word-num">字数：{{ task.word_num }}词</view>
					<view class="requirement">要求：{{ task.requirement }}</view>
					<view class="line">
						<view class="manager-name">{{ task.manager_name }}</view>
						<view class="publish-date">发布日期：{{ task.publish_date }}</view>
						<button class="submit-btn" @click="handleSubmit(task.title_id)">
							提交
						</button>
					</view>
				</view>
				<view v-else>暂无写作任务</view><!--todo:暂无任务-->
			</view>
		</view>
		<view class="composition-container">
			<view class="composition-head">作文广场</view>
			<view class="composition-tabs">
				<Tabs>
					<template v-slot:tab1-content>
						<view class="history-items" v-for="(task, index) in writingCompleted" :key="index">
							<view class="history-item" v-if="writingCompleted.length > 0">
								<view style="
                    padding-bottom: 0.2rem;
                    border-bottom: 2px solid #69c0ff;
                  ">
									<view class="history-item-title"><span
											style="color: rgb(69, 109, 231)">[题目]</span><span
											style="margin-left: 1rem">{{
                      task.title
                    }}</span></view>
								</view>
								<view class="word-num">字数：{{ task.word_num }}词</view>
								<view class="requirement">要求：{{ task.requirement }}</view>
								<view class="line">
									<view class="manager-name">{{ task.manager_name }}</view>
									<view class="publish-date">提交日期：{{ task.submit_date }}</view>
									<span class="tag">{{ task.tag }}</span>
								</view>
								<div class="progress-ring">
									<svg class="progress-ring__svg" viewBox="0 0 120 120">
										<circle class="progress-ring__circle-bg" stroke="#ddd" stroke-width="13"
											fill="transparent" r="50" cx="60" cy="60" />
										<circle class="progress-ring__circle-fg" stroke="#44a0fb" stroke-width="13"
											fill="transparent" r="50" cx="60" cy="60" :style="{
                        strokeDasharray: circumference,
                        strokeDashoffset: offset,
                      }" />
									</svg>
									<span class="score">{{ score }}</span>
								</div>
							</view>
						</view>
					</template>
					<template v-slot:tab2-content>
						<view class="task-items-2" v-for="(task, index) in writingTraining" :key="index" >
							<view class="task-item-2">
								<view style="padding-bottom: 0.2rem; border-bottom: 2px solid #69c0ff">
									<view class="task-item-title"><span
											style="color: rgb(69, 109, 231)">[题目]</span><span
											style="margin-left: 1rem">{{ task.title }}</span></view>
								</view>
								<view class="word-num">字数：{{ task.word_num }}词</view>
								<view class="requirement">要求：{{ task.requirement }}</view>
								<view class="line">
									<view class="manager-name">{{ task.manager_name }}</view>
									<button class="submit-btn" @click="handleSubmit(task.title_id)">
										提交
									</button>
								</view>
							</view>
						</view>
					</template>
				</Tabs>
			</view>
		</view>
	</view>
</template>

<script>
	import Tabs from "../../components/Tabs.vue";
	export default {
		name: "App",
		components: {
			Tabs,
		},
		data() {
			return {
				score: 50,
				radius: 50, //半径
				circumference: 2 * Math.PI * 50, //周长
				//写作任务列表
				writingTasks: [{
						title_id: 6,
						title: "小作文1",
						manager_name: "qwerty",
						word_num: "50~100",
						requirement: "请不要用真实姓名，使用“李明”asdfafaswetregrtiyujnhdasfwerwedgsft54ewawds代替",
						publish_date: "2024-06-15",
						grade: "小学",
					},
				],
				//写作训练列表
				writingTraining:[
					{
						title_id: 6,
						title: "作文训练1",
						manager_name: "qwerty",
						word_num: "50~100",
						requirement: "请不要用真实姓名，使用“李明”asdfafaswetregrtiyujnhdasfwerwedgsft54ewawds代替",
						publish_date: "2024-06-15",
						grade: "小学",
					},
				],
				//作文广场
				//已提交的写作任务
				writingCompleted: [{
						tag: "训练",
						title_id: 6,
						title: "小作文1",
						manager_name: "qwerty",
						word_num: "50~100",
						requirement: "请不要用真实姓名，使用“李明”asdfafaswetregrtiyujnhdasfwerwedgsft54ewawds代替",
						submit_date: "2024-06-15",
						grade: "小学",
					},
				],
			};
		},
		onLoad() {
			//获取写作数据
			this.getWirtingData();
		},
		methods: {
      handleSubmit(titleId){
        //跳转到提交页面
        uni.navigateTo({
          url: `../UploadEssay/UploadEssay?titleId=${titleId}`,
        });
      },
			getWirtingData() {
				uni.request({
					url: "http://localhost:8080/api/users/composition_mission",
					method: "GET",
					header: {
						Authorization: `Bearer ${uni.getStorageSync("token")}`,
					},
					success: (res) => {
						console.log(res.data);
						if (res.statusCode === 200) {
							this.writingTasks = res.data.tasks;
							this.writingCompleted = res.data.finished_writings;
							this.writingTraining = res.data.trainings;
						}
					},
					fail: (err) => {
						console.log(err);
					},
				});
			},
		},
		computed: {
			offset() {
				let progress = this.score / 100;
				return this.circumference * (1 - progress);
			},
		},
	};
</script>

<style>
	.head {
		width: 100%;
		height: 3rem;
		position: relative;
		top: 0;
		left: 0;
		z-index: 100;
	}

	.back-icon {
		width: 2rem;
		height: 2rem;
		position: absolute;
		top: 0.8rem;
		left: 0.5rem;
		cursor: pointer;
	}

	.title {
		position: absolute;
		font-size: 1.2rem;
		font-weight: bold;
		text-align: center;
		top: 0.8rem;
		left: 40%;
	}

	.task-list-container {
		position: relative;
		margin-top: 2rem;
		width: 100%;
	}

	.task-list-title {
		font-size: 1.2rem;
		font-weight: bold;
		margin-left: 0.5rem;
	}

	.task-items {
		margin-top: 1rem;
		width: 100%;
	}

	.task-item {
		width: 90%;
		margin-left: 5%;
		background-color: #ffffff;
		margin-top: 1rem;
		border-radius: 0.5rem;
		border: 1px solid #000000;
	}

	.task-item-title {
		font-size: 1.1rem;
		margin-left: 0.5rem;
	}

	.word-num {
		font-size: 0.9rem;
		margin-left: 0.5rem;
		margin-top: 0.5rem;
	}

	.requirement {
		font-size: 0.9rem;
		margin-left: 0.5rem;
		overflow: hidden;
		text-overflow: ellipsis;
		/* white-space: nowrap; */
		margin-bottom: 0.5rem;
    max-width: 300px;
    max-height: 40px;
    min-height:40px;
	}

	.line {
		display: flex;
		margin-top: 1rem;
		border-top: 2px solid #69c0ff;
		padding-top: 0.3rem;
		padding-bottom: 0.3rem;
	}

	.manager-name {
		font-size: 0.9rem;
		margin-left: 0.5rem;
	}

	.publish-date {
		font-size: 0.9rem;
		margin-left: 1rem;
	}

	.submit-btn {
		width: 4rem;
		height: 1.5rem;
		background-color: #69c0ff;
		color: #ffffff;
		font-size: 0.8rem;
		cursor: pointer;
		text-align: center;
		line-height: 1.5rem;
	}

	.composition-container {
		margin-top: 2rem;
		width: 100%;
	}

	.composition-head {
		font-size: 1.2rem;
		font-weight: bold;
		margin-left: 0.5rem;
	}

	.composition-tabs {
		margin-top: 1rem;
	}

	.history-items {
		margin-top: 1rem;
		width: 100%;
		position: relative;
	}

	.history-item {
		background-color: #ffffff;
		margin-top: 1rem;
		border-radius: 0.5rem;
		border: 1px solid #000000;
	}

	.history-item-title {
		font-size: 1.1rem;
		margin-left: 0.5rem;
	}

	.tag {
		font-size: 0.9rem;
		margin-right: 2rem;
		margin-left: auto;
		color: #0089ea;
		border: 1px solid #0e7bf8;
		height: 1.3rem;
		line-height: 1.3rem;
		width: 2rem;
		text-align: center;
		border-radius: 15%;
	}

	.progress-ring {
		position: absolute;
		width: 75px;
		height: 75px;
		right: 0;
		top: 20px;
		z-index: 1;
		background-color: white;
	}

	.score {
		position: absolute;
		top: 50%;
		left: 50%;
		transform: translate(-50%, -50%);
		font-size: 1.7rem;
		color: #44a0fb;
	}

	.progress-ring__svg {
		transform: rotate(-90deg);
	}

	.progress-ring__circle-bg,
	.progress-ring__circle-fg {
		stroke-dasharray: 314.1592653589793;
	}

	.progress-ring__circle-fg {
		transition: stroke-dashoffset 0.35s;
	}
  .task-items-2 {
    margin-top: 1rem;
    width: 100%;
    border:1px solid #000000;
    border-radius: 0.5rem;
    .submit-btn{
      margin-right:2rem;
    
    }
  }
</style>