<template>
	<view class="container">
		<view class="head">
			<image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
			<span>加入团队</span>
			<button @click="joinTeam">确认</button>
		</view>

		<view class="input-wrapper">
			<input class="uni-input" placeholder="请输入团队码" v-model="inputClearValue" />
			<image class="uni-icon" src="../../static/not-done2.svg" v-if="inputClearValue.length>0" @click="clearIcon">
			</image>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				inputClearValue: '',
			}
		},
		methods: {
			handleBack() {
				uni.navigateBack({
					delta: 1
				})
			},
			clearIcon() {
				this.inputClearValue = '';
			},
      joinTeam() {
        uni.request({
          url: '/api/users/my_team/join_team',
          method: 'POST',
          data: {
            Invitation_Code: this.inputClearValue
          },
          header: {
            'content-type': 'application/json', // 默认值
            'Authorization': `Bearer ${uni.getStorageSync('token')}`
          },
          success: (res) => {
            if (res.data.code === 200) {
              uni.showToast({
                title: '加入成功',
                icon: 'success',
                duration: 2000
              });
              uni.navigateBack({
                delta: 1
              });
            } else {
              uni.showToast({
                title: '加入失败',
                icon: 'none',
                duration: 2000
              });
            }
          },
          fail: (res) => {
            uni.showToast({
              title: '加入失败',
              icon: 'none',
              duration: 2000
            });
          }
        });
      },
		}
	}
</script>

<style>
	.container {
		display: flex;
		flex-direction: column;
		align-items: center;
		height: 100vh;
		width: 100vw;
	}

	.head {
		display: flex;
		height: 2rem;
		width: 100vw;
		background-color: transparent;
	}

	.back-icon {
		width: 2rem;
		/*图标宽度 */
		height: 2rem;

	}

	span {
		margin-left: 7rem;
		font-size: 40rpx;
		font-weight: bold;
		margin-top: 2rpx;
	}

	button {
		margin-left: 5rem;
		background-color: transparent;
		line-height: 2rem;
		height: 2rem;
		font-size: 21px;
		align-items: center;
		border: none;
		color:#456DE7;
	}

	button::after {
		border: none;
	}

	/* 
	input {
		width: 90%;
		background-color: rgb(227, 227, 227);
		margin-top: 2rem;
		height: 2.6rem;
	} */
	.input-wrapper {
		display: flex;
		padding-left: 16px;
		padding-right: 13px;
		padding-top: 16px;
		padding-bottom: 8px;
		margin-top: 2rem ;
		flex-direction: row;
		flex-wrap: nowrap;
		background-color: #e3e3e3;
		width: 80%;
		height: 2.2rem;
		border-radius: 5px;
	}

	.uni-input {
		height: 28px;
		font-size: 22px;
		padding: 0px;
		flex: 1;
		background-color: #e3e3e3;
		color: #000;
	}

	.uni-icon {
		width: 24px;
		height: 24px;
	}
</style>