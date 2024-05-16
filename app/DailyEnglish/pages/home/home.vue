<template>
	<view class="homepage">
		<view class="search-container" >
			<view class="search" @click="handleSearch">
				<image class="search-icon" src="/static/search.svg"></image>
				<input placeholder="搜索" v-model:value="searchInput">
			</view>
			<view class="history" v-show="isHistoryVisible">
				<view class="history-header">
					<text class="title">历史搜索</text>
					<text class="clean">清空</text>
				</view>
				<view class="list">
					<view class="item" v-for="(item, index) in items" :key="index" @click="handleSearchInput(item.word)">
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
		height: 60rpx;
		padding: 5rpx;
		width: 70%;
		align-items: center;
		margin-left: 40rpx;
	}

	.search-icon {
		width: 25px;
		height: 25px;
		margin-left: 200rpx;
		margin-right: 20rpx;
	}

	.search input {
		flex: 1;
		border: none;
		outline: none;
		text-align: center;
		width: 10rpx;
		font-size: 30rpx;
		max-width: 60rpx;
	}

	.history {
		margin-top: 30rpx;
		width: 100%;
		background-color: #fff;
		height: calc(100vh - 60rpx);
		/* display:none; */
	}


	.history-header {
		display: flex;
		justify-content: space-between;
		/* 使.title和.clean分别位于顶部两边 */
		align-items: center;
		/* padding-top: 20rpx;
		padding-left: 30rpx;
		padding-right: 30rpx; */
		padding: 30rpx;
	}

	.title {
		font-size: 35rpx;
		color: #767676;
	}

	.clean {
		font-size: 30rpx;
		color: #767676;
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
		height:130rpx;
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
					},{
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
				this.isHistoryVisible = !this.isHistoryVisible;
			},
      handleSearchInput(input) {
        this.searchInput = input;
      }
		}
	}
	
</script>