<!--
 * @Date: 2024-04-07 17:52:57
-->
<template>
	<view class="container">
		<view class="scroll-view-container">
			<text class="title">以下是scroll-view的样例</text>
			<scroll-view class="scroll-view" scroll-x="true">
				<!-- scroll-x="true"表示横向滚动 -->
				<view class="list-container">
					<view class="list">
						<view class="item">列表1 项目1</view>
					</view>
					<view class="list">
						<view class="item">列表2 项目1</view>
					</view>
					<view class="list">
						<view class="item">列表3 项目1</view>
					</view>
					<view class="list">
						<view class="item">列表4 项目1</view>
					</view>
				</view>
			</scroll-view>
		</view>

		<view class="swiper-container">
			<text class="title">以下是swiper的样例</text>
			<swiper class="swiper" :indicator-dots="true" :autoplay="true" :interval="3000" :duration="500"
				:circular="true">
				<!-- indicator-dots="true"表示显示指示点 -->
				<!-- autoplay="true"表示自动播放 -->
				<!-- interval="3000"表示自动播放间隔时间 -->
				<!-- duration="500"表示切换动画时间 -->
				<!-- circular="true"表示循环播放 -->
				<swiper-item>
					<view class="swiper-item">轮播1</view>
				</swiper-item>
				<swiper-item>
					<view class="swiper-item">轮播2</view>
				</swiper-item>
				<swiper-item>
					<view class="swiper-item">轮播3</view>
				</swiper-item>
			</swiper>
		</view>

		<view style="height: auto">
			<sliderzz @change="change"></sliderzz>
		</view>

		<view class="toast-container">
			<button @click="showToast('Hello World')">点击我</button>
			<toast ref="toast" />
		</view>

		<view class="backTop-container">
			<backTop v-if="showBackTop"></backTop>
		</view>

	</view>
</template>

<script>
	import sliderzz from "@/components/sliderzz.vue";
	import toast from "@/components/toast.vue";
	import backTop from '@/components/backTop.vue';
	export default {
		data() {
			return {
				showBackTop:true
			};

		},
		components: {
			sliderzz,
			toast,
			backTop
		},
		methods: {
			change({
				finish,
				reset
			}) {
				//这里有两个函数一般滑动都是要校验之后才能提交的，所以第一个finish知行之后他会提交
				finish(); //保存提交状态
				// setTimeout(() => {
				// 	reset();//这个是提交之后过一会可能会重新提交进入重置状态
				// },2000)
			},
			showToast(message) {
				this.$refs.toast.showToast(message);
			},
			onPageScroll(e) {
			    // 获取滚动的距离
			    let scrollTop = e.scrollTop;
			
			    // 如果滚动距离超过100px，执行某些操作
			    this.showBackTop = true;
			}
		},
	};
</script>

<style lang="scss">
	.title {
		font-size: large;
	}

	// 容器样式
	.scroll-view-container {
		display: flex;
		flex-direction: column;
		align-items: center; //交叉轴居中
		justify-content: center;
	}

	// scroll-view样式
	.scroll-view {
		width: 100%;
		white-space: nowrap; //不换行
		max-height: 100 rpx; //这个单位可以保证在不同设备上显示效果一致，是比例单位
	}

	// 列表容器样式
	.list-container {
		display: flex;
	}

	// 列表样式
	.list {
		display: flex;
		flex-direction: column;
		margin: 10rpx;
		width: 300rpx; //这个本来4个list超出屏幕，现在只显示3个，scroll-view的属性设置了不换行，并且最高高度锁死，所以只能横向滚动
	}

	// 列表项目样式
	.item {
		margin: 5px;
		padding: 10px;
		background-color: #f0f0f0;
		border-radius: 5px;
	}

	// swiper容器样式，这是scss的写法，可以嵌套，在此容器外的swiper-item样式不会受影响，需要另外写样式，接下来为了美观，都使用scss写法
	.swiper-container {
		width: 100%;
		height: 200rpx;

		.swiper {
			width: 100%;
			height: 100%;
		}

		.swiper-item {
			width: 99%;
			height: 99%;
			border: 1rpx solid #f00000;
		}
	}
</style>