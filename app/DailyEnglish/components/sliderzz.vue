<template>
	<view class="react" id="react">
		<view :style="{ width: width + 'px', backgroundColor: '#65B58A' }" class="left kong">{{ title }}</view>
		<view :style="{ left: width + 'px' }" @touchstart="start" @mousedown="start" @mouseup="end" @touchmove="move"
			@mousemove="move" @touchend="end" id="slider" :class="{ slider: true, select: title }">
			<image v-if="!title" src="https://img-blog.csdnimg.cn/a5bf3043a7d344cb88f762186c1dfc90.png#pic_center"
				mode="widthFix"></image>
			<!-- 这是两张图片一个是大于号一个是对号-->
			<image v-else src="https://img-blog.csdnimg.cn/e7b7798beb3a442f8a67a073b23c698a.png#pic_center"
				mode="widthFix"></image>
		</view>
		<view class="right kong"> 右滑提交 </view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				title: "", //划过之后的标题
				width: 0, //滑到多宽
				reactWidth: 0, //整个矩形的宽度
				sliderWidth: 0, //滑块宽度
				startX: 0, //开始触摸距离屏幕左面的位置
				sendFlag: false, //是否发送
				finishFlag: false, //是否允许滑动 判断是否滑动完成
				moveFlag: false, //是否执行滑动函数
				//   isLongPress: false,//是否长按
				//   longPressTimeout: null,//长按定时器
				// longPressThreshold: 500, // 长按阈值，单位为毫秒
			};
		},
		mounted() {
			let selectFc = uni.createSelectorQuery().in(this);
			selectFc
				.select("#react")
				.boundingClientRect((data) => {
					//获取总宽度
					this.reactWidth = data.width - 2; //矩形宽度去掉边框宽度
				})
				.exec();
			selectFc
				.select("#slider")
				.boundingClientRect((data) => {
					//获取滑块的宽度
					this.sliderWidth = data.width;
				})
				.exec();
		},
		methods: {
			start(e) {
				//开始的触摸
				let {
					clientX,
					clientY
				} = e.touches[0];
				this.startX = clientX; //记录按下时刻距离屏幕左侧的距离
				this.moveFlag = true; //允许滑动
				// if(e.type==='mousedown'){

				//  this.longPressTimeout = setTimeout(() => {
				//   this.isLongPress = true;
				// }, this.longPressThreshold);
				// }
			},
			reset() {
				//重置划款状态
				this.sendFlag = false;
				this.finishFlag = false;
				this.width = 0;
				this.title = "";
				// this.isLongPress = false;
				//      clearTimeout(this.longPressTimeout);
			},
			finish() {
				//划款完成状态
				this.finishFlag = true;
				this.title = "已提交";
			},
			move(e) {
				//划款移动中
				if (!this.moveFlag) return;
				// if (!this.isLongPress) return;
				if (this.width >= this.reactWidth - this.sliderWidth) {
					if (!this.sendFlag) {
						//到达最后面后就不允许他在滑动了，不然他会跳动体验比较差，所以加了限制
						this.moveFlag = false;
						this.sendFlag = true;
						this.$emit("change", {
							finish: this.finish.bind(this),
							reset: this.reset.bind(this),
						}); //此时划款正好完成达到最右侧
					}
				} else {
					let {
						clientX,
						clientY
					} = e.touches[0];
					var width = clientX - this.startX; //下面判断要是小0就不能在滑动，要是大于最大长度也要停止
					if (width >= this.reactWidth - this.sliderWidth) {
						width = this.reactWidth - this.sliderWidth;
					} else if (width <= 0) {
						width = 0;
					}
					this.width = width;
				}
			},
			end(e) {
				//滑块结束时刻
				this.moveFlag = true;
				if (this.finishFlag) {
					//完成状态
					if (this.width < this.reactWidth - this.sliderWidth) {
						this.width = 0;
					}
				} else {
					//没有完成每次都要重置
					this.reset();
				}
				// clearTimeout(this.longPressTimeout);
				//       if (!this.isLongPress) {
				//         this.reset();
				//       }
			},
		},
	};
</script>

<style lang="scss" scoped>
	.react {
		margin: 0 auto;
		margin-top: 100rpx;
		width: 287rpx;
		height: 60rpx;
		box-sizing: content-box;
		border: 1px solid #0e57a1;
		display: flex;
		align-items: center;
		position: relative;

		.sliderYuan {
			position: absolute;
			width: 100%;
			left: 0;
			top: 50%;
			margin: 0;
			transform: translate(0, -50%);
		}

		.slider {
			width: 60rpx;
			height: 60rpx;
			background: #0e57a1;
			display: flex;
			align-items: center;
			justify-content: center;
			position: relative;

			&.select {
				background: #ffffff;
			}

			image {
				display: block;
				width: 29rpx;
				height: auto;
			}
		}

		.kong {
			text-align: center;
			line-height: 60rpx;
			font-size: 22rpx;
			font-family: PingFangSC-Medium, PingFang SC;
			font-weight: 500;
			letter-spacing: 1rpx;

			&.right {
				flex: 1;
				color: #0e57a1;
			}

			&.left {
				position: absolute;
				left: 1;
				top: 0;
				z-index: 10;
				height: 60rpx;
				color: #ffffff;
			}
		}
	}
</style>
<!-- 模板部分：
<view :style="{left:width+'px'}" @touchstart="start" @touchmove="move" @touchend="end" id="slider" :class="{slider:true,select:title}">：定义了一个滑动区域，其初始位置由 width 决定。当滑动到右侧时，它会显示对号图标；当未滑动到右侧时，它会显示大于号图标。
<view class="right kong">右滑提交</view>：定义了一个右侧区域，显示右滑提交的提示。
脚本部分：
mounted() {：当组件挂载后，获取 react 和 slider 的宽度和总宽度。
start(e)：当触摸开始时，记录起始位置和设置允许滑动。
move(e)：在触摸移动时，更新宽度并检查是否到达终点。
end(e)：在触摸结束时，重置或完成状态。
reset()：重置滑动状态。
finish()：设置完成状态。
-->