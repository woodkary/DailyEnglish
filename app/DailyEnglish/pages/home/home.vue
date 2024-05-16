<template>
	<view class="homepage">
		<view class="search-container">
			<view class="search-head" style="display: flex;">
				<view class="search" :class="{active:isHistoryVisible}" @click="handleSearch">
					<image class="search-icon" src="/static/search.svg"></image>
					<input placeholder="搜索" v-model:value="searchInput">
				</view>
				<button class="cancel" v-show="isHistoryVisible" @click="cancelSearch">取消</button>
			</view>
			<view class="history" v-show="isHistoryVisible">
				<view class="history-header">
					<text class="title">历史搜索</text>
					<text class="clean">清空</text>
				</view>
				<view class="list">
					<view class="item" v-for="(item, index) in items" :key="index"
						@click="handleSearchInput(item.word)">
						<view class="top-row">
							<view class="word">{{ item.word }}</view>
							<view class="phonetic">{{ item.phonetic }}</view>
						</view>
						<view class="meaning">{{ item.meaning }}</view>
					</view>
				</view>

			</view>
		</view>
	</view>
</template>

<style>
	.homepage {
		background-color: #f5f5f5;
		height: 100vh;
		width: 100vw;
	}

	.search-container {
		width: 100%;
		padding-top: 20rpx;
	}

	.search {
		background-color: #fff;
		border-radius: 50rpx;
		display: flex;
		height: 65rpx;
		padding: 5rpx;
		width: 70%;
		align-items: center;
		margin-left: 45rpx;
		/* transition: all 0.3s ease; */
		/* 为所有属性添加过渡效果 */
	}

	.search.active {
		align-items: flex-start;
		/* 从 left 改为 flex-start */
		border-radius: 15rpx;
		width: 75%;
		margin-left: 40rpx;
		padding: 5rpx;
    border-color: rgba(247, 127, 0, 0.4);
    background-color: #fff;
    box-shadow: 0 0 0 4px rgb(247 127 0 / 10%);
	}

	.search-icon {
		width: 25px;
		height: 25px;
		margin-left: 200rpx;
		margin-right: 20rpx;
		/* transition: all 0.3s ease; */
		/* 为所有属性添加过渡效果 */
	}

	.search.active .search-icon {
		margin-left: 20rpx;
		margin-right: 0;
		width: 25px;
		height: 25px;
		margin-top: 10rpx;
	}

	input {
		flex: 1;
		border: none;
		outline: none;
		text-align: center;
		width: 10rpx;
		font-size: 30rpx;
		max-width: 60rpx;
		/* transition: all 0.3s ease; */
		/* 为所有属性添加过渡效果 */
	}

	.cancel {
		font-size: 40rpx;
		margin-top: -20rpx;
		font-weight: 530;
		background-color: transparent;
		color: #000000;
		border: none;

		&::after {
			border: none;
		}

	}


	.search.active input {
		width: 80%;
		text-align: left;
		max-width: 80%;
		height: 100%;
		font-size: 38rpx;
		margin-left: 10rpx;
		/* margin-top:10rpx; */
	}

	.history {
		margin-top: 30rpx;
		width: 100%;
		background-color: #fff;
		height: calc(100vh - 60rpx);
	}

	.history-header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 30rpx;
	}

	.title {
		font-size: 35rpx;
		color: #767676;
	}

	.clean {
		font-size: 35rpx;
		color: #767676;
		cursor: pointer;
	}

	.list {
		display: flex;
		flex-direction: column;
		width: 90%;
		margin-left: 5%;
		height: auto;
	}

	.item {
		margin-bottom: 10px;
		border-bottom: 1px solid #cecece;
		height: 130rpx;
	}

	.top-row {
		display: flex;
	}

	.word {
		margin-right: 30px;
		font-size: 40rpx;
		font-weight: 600;
	}

	.phonetic {
		font-size: 30rpx;
		margin-top: 5px;
		color: #767676;
		font-weight: 600;
	}

	.meaning {
		margin-top: 10px;
		font-size: 30rpx;
		color: #767676;
		overflow: hidden;
		white-space: nowrap;
		font-weight: 500;
	}
</style>
<script>
	export default {
		data() {
			return {
				isHistoryVisible: false,
				searchInput: '',
				items: [{
						word: 'apple',
						phonetic: '/ˈæpl/',
						meaning: '苹果111111111111111111111111111111111111111111111111'
					}, {
						word: 'banana',
						phonetic: '/bəˈnɑː.nə/',
						meaning: '香蕉'
					}

					// 其他列表项
				]
			}
		},
		methods: {
			handleSearch() {
				this.isHistoryVisible = true;
			},
			handleSearchInput(input) {
				this.searchInput = input;
			},
			cancelSearch() {
				this.isHistoryVisible = false;
				this.searchInput = '';
			}
		}
	}
</script>