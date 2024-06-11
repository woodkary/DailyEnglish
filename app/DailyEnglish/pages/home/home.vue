<template>
	<view class="homepage">
		<view class="search-container">
			<view class="search-head" style="display: flex;">
				<view class="search" :class="{active:isHistoryVisible}" @click="handleSearchShow()">
					<image class="search-icon" src="/static/search.svg"></image>
					<input placeholder="搜索" v-model="searchInput" @input="handleSearchInput(searchInput)" @confirm="handleSearchRouter">
				</view>
				<button class="cancel" v-if="isHistoryVisible" @click="cancelSearch">取消</button>
				<image class="Msg-icon" v-else src="/static/email.png"></image>
			</view>
			<view class="history" v-show="isHistoryVisible">
				<view class="history-header">
					<text class="title">搜索</text>
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
		<view class="daka-container" v-show="!isHistoryVisible">
			<image src="../../static/lihua.png" v-show="isDaka"
				style="position: absolute;width:330px;height:330px;top:140px;left:120px;"></image>
			<view class="daka-head">
				<view class="column">
					<image class="vocabook-img" src="../../static/book.png"></image>
				</view>
				<view class="column">
					<view class="row">
						<view class="daka-title">{{ daka_book }}</view>
						<view class="daka-subtitle">修改</view>
					</view>
					<view class="row">
						<progress percent="10" active-color="#10aeff" backgroundColor="#c8c8c8"
							stroke-width="7"></progress>
					</view>
					<view class="row">
						<view class="progress1">{{ wordNumLearned }}/{{ wordNumTotal }}</view>
						<view class="progress2">剩余{{ daysLeft }}天</view>
					</view>
				</view>

			</view>
			<view class="daka-line">
				<view class="daka-title" v-if="isDaka" style="font-size: 30rpx;font-weight: normal;">恭喜你！<br>完成今日打卡
				</view>
				<view class="daka-title" v-else>今日计划</view>
				<view class="daka-slogan">随时随地，单词猛记</view>
			</view>
			<view class="daka-plan">
				<view class="row">
					<view class="plan-title1" v-if="isDaka">今日已新学</view>
					<view class="plan-title1" v-else>需新学</view>
					<view class="plan-title2" v-if="isReview">今日已复习</view>
					<view class="plan-title2" v-else>需复习</view>
				</view>
				<view class="row">
					<view class="plan-num">
						<view class="number">{{ isDaka?wordNumPunched:wordNumToPunch }}</view>
						<text>词</text>
					</view>
					<view class="plan-num" style="margin-left:50px">
						<view class="number">{{ isReview?wordNumReviewed:wordNumToReview }}</view>
						<text>词</text>
					</view>

				</view>
				<view class="row">
					<button class="plan-btn1" v-show="!isDaka " @click="handleDaka">开始学习</button>
					<button class="plan-btn1" v-show="isDaka && !isReview" @click="handleReview">开始复习</button>
					<button class="plan-btn1" v-show="isDaka && isReview " @click="handleDaka">继续学习</button>
				</view>
			</view>
		</view>
		<view class="content-container" v-show="!isHistoryVisible">
			<view class="button-list">
				<view class="btn-item">
					<image src="/static/word-exercise.png"></image>
					<text>单词训练</text>
				</view>
				<view class="btn-item">
					<image src="/static/biji.svg"></image>
					<text>单词自检</text>
				</view>
				<view class="btn-item">
					<image src="/static/write.svg"></image>
					<text>爱写作</text>
				</view>
				<view class="btn-item">
					<image src="/static/read.svg"></image>
					<text>爱阅读</text>
				</view>
			</view>
		</view>
		<view class="ad-container" v-show="!isHistoryVisible">
			<view class="ad-list">
				<view class="ad-item">
					<view class="text-container">
						<text class="title">30分钟，拿下英语阅读</text>
						<text class="content">每日一读，提高英语阅读能力</text>
					</view>
					<image class="image" src="/static/ad1.svg"></image>
				</view>
				<view class="ad-item">
					<view class="text-container">
						<text class="title">30分钟，拿下英语阅读</text>
						<text class="content">每日一读，提高英语阅读能力</text>
					</view>
					<image class="image" src="/static/ad1.svg"></image>
				</view>
			</view>
		</view>
	</view>
</template>

<style>
	@font-face {
		font-family: "SF-UI-Text";
		src: url('@/static/SF-UI-Text-Regular.otf');
	}

	.homepage {
		background-color: #f8f8f8;
		height: 100vh;
		width: 100vw;
	}

	.search-container {
		width: 100%;
		padding-top: 20rpx;
	}

	.search {
		background-color: #ffffff;
		border-radius: 50rpx;
		display: flex;
		height: 72rpx;
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
		border: 2px solid rgba(255, 115, 0, 0.4);
		height: 65rpx;
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

	.Msg-icon {
		width: 75rpx;
		height: 75rpx;
		margin-left: 40rpx;
	}

	.search.active input {
		width: 80%;
		text-align: left;
		max-width: 80%;
		color: #000000;
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

	.daka-container {
		width: 90%;
		height: 690rpx;
		margin-left: 5%;
		margin-top: 40rpx;
		background-color: white;
		border-radius: 10px;
		box-shadow: 0px 1px 10px rgba(0, 0, 0, 0.1);

		.daka-head {
			display: flex;
			flex-direction: row;

			.column {
				display: flex;
				flex-direction: column;
			}

			.row {
				display: flex;
				flex-direction: row;
				align-items: center;
			}

			.vocabook-img {
				width: 160rpx;
				height: 180rpx;
				margin-top: 40rpx;
				margin-left: 40rpx;
			}

			.daka-title {
				margin-top: 40rpx;
				font-weight: 550;
				font-size: 40rpx;
				margin-left: 40rpx;
			}

			.daka-subtitle {
				margin-top: 42rpx;
				margin-left: 20rpx;
				color: #9e9e9e;
			}

			progress {
				margin-left: 40rpx;
				height: 60rpx;
				width: 350rpx;
			}

			.progress1 {
				margin-left: 30rpx;
				color: #9e9e9e;
			}

			.progress2 {
				margin-left: 100rpx;
				color: #9e9e9e;
			}
		}

		.daka-line {
			display: flex;
			flex-direction: row;
			align-items: center;
			justify-content: space-between;
			margin-top: 40rpx;
			margin-left: 40rpx;
			margin-right: 40rpx;
			/* 	border-top: 1px solid #e4e4e4;
			border-bottom: 1px solid #e4e4e4; */

			.daka-title {
				font-weight: 550;
				font-size: 40rpx;
			}

			.daka-slogan {
				color: #F55F4A;
				margin-top: 10rpx;
			}
		}

		.daka-plan {
			display: flex;
			flex-direction: column;
			margin-top: 40rpx;
			margin-left: 20rpx;

			.row {
				display: flex;
				flex-direction: row;

				.plan-title1 {
					fontsize: 45rpx;
					margin-left: 40rpx;
				}

				.plan-title2 {
					fontsize: 45rpx;
					margin-left: 200rpx;
				}

				.plan-num {
					margin-left: 40rpx;
					display: flex;

					text {
						margin-left: 20rpx;
						margin-top: 90rpx;
						font-size: 45rpx;
						font-weight: 550;
					}
				}

				.number {
					font-size: 150rpx;
					font-weight: 600;
					font-family: 'SF-UI-Text';
					/*斜体*/
					font-style: italic;
				}

			}

			.plan-btn1 {
				width: 90%;
				height: 90rpx;
				background-color: #2c8af5;
				color: white;
				font-size: 25px;
				display: flex;
				justify-content: center;
				align-items: center;
				/* 垂直居中 */
				margin-left: 21rpx;
			}
		}

	}

	.content-container {
		display: flex;
		justify-content: center;
		align-items: center;
		width: 100%;
	}

	.button-list {
		display: flex;
		white-space: nowrap;
		justify-content: center;
		width: 100%;
		margin-top: 40rpx;
		background-color: transparent;

	}

	.btn-item {
		display: flex;
		flex-direction: column;
		align-items: center;
		width: 144rpx;
		height: 144rpx;
		border: 1px solid #e4e4e4;
		margin-left: 30rpx;
		border-radius: 10rpx;
		background-color: #fff;
		box-shadow: 0px 4px 10px rgba(0, 0, 0, 0.1);
		/* 渐变阴影 */

	}

	.btn-item:first-child {
		margin-left: 0rpx;
	}

	.btn-item image {
		margin-top: 10rpx;
		width: 102rpx;
		height: 102rpx;
	}

	.ad-container {
		width: 100%;
		margin-top: 40rpx;
	}

	.ad-list {
		display: flex;
		flex-direction: column;
		/* background-color: white; */
		width: 90%;
		margin-left: 5%;
	}

	.ad-item {
		display: flex;
		border: 1px solid #e4e4e4;
		border-radius: 10rpx;
		margin-bottom: 40rpx;
		height: 86px;
		background-color: #fff;
		box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1);
	}

	.text-container {
		display: flex;
		/* 将文本容器设置为Flex容器 */
		flex-direction: column;
		/* 垂直排列 */
		margin-right: 30px;
		/* 右边距 */
		/*居中*/
		justify-content: center;
		margin-left: 20px;
	}

	.title {
		font-size: 20px;
		font-weight: bold;
	}

	.content {
		font-size: 16px;
		color: #666;
		margin-top: 5px;
	}

	.image {
		width: 110px;
		height: 110px;
		margin-top: -15px;
	}
</style>
<script>
	export default {
		data() {
			return {
				isHistoryVisible: false, //查询单词
				isDaka: false, //是否打卡
				isReview: false, //是否复习
				searchInput: '',
				daka_book: '',
				wordNumLearned: 123,
				wordNumTotal: 2345,
				daysLeft: 30,
				wordNumToPunch: 5,
				wordNumPunched: 5,
				wordNumToReview: 5,
				wordNumReviewed: 5,
        searchHistory: {}, //历史搜索记录
        //历史搜索记录
				items: [/*{
					word: 'apple',
					phonetic: '/ˈæpl/',
					meaning: '苹果'
				}, {
					word: 'banana',
					phonetic: '/bəˈnɑː.nə/',
					meaning: '香蕉'
				}*/]
			}
		},
    //switchTab后调用的函数
    onShow() {
      // 从本地缓存中获取搜索记录
      this.items=uni.getStorageSync('searchHistory')||[];
      let toReview = uni.getStorageSync('toReview');
      if(toReview){
        this.isDaka=true;
        this.isReview = false;
        this.wordNumPunched=this.wordNumToPunch;
      }
      let reviewed = uni.getStorageSync('reviewed');
      if(reviewed){
        this.isReview=true;
        this.isDaka = true;
        this.wordNumReviewed = this.wordNumToReview;
      }
      this.fetchData();
    },
    onLoad() {
      console.log("hi");
    },
		methods: {
      transformWordObject(original) {
        // 创建一个新的对象，用于存储转换后的结构
        let transformed = {
          word: original.spelling,
          phonetic: original.pronunciation
        };

        // 遍历原始对象的meanings属性，找到第一个非空的词性数组
        for (let partOfSpeech in original.meanings) {
          if (original.meanings[partOfSpeech].length > 0) {
            // 使用第一个词义作为meaning属性的值
            transformed.meaning = original.meanings[partOfSpeech][0];
            break; // 找到第一个非空的词性后，跳出循环
          }
        }

        // 返回转换后的对象
        return transformed;
      },
			handleDaka() {
				//operation 0：打卡，1：复习
				uni.navigateTo({
					url: "/pages/Examination/Examination?operation=" + 0
				});
			},
			handleReview() {
				uni.navigateTo({
					url: "/pages/Examination/Examination?operation=" + 1
				});
			},
			fetchData() {
				uni.request({
					url: "http://localhost:8080/api/punch/main_menu",
					header: {
						'Authorization': `Bearer ${uni.getStorageSync('token')}`
					},
					method: 'POST',
          data:{
            times: this.isDaka+this.isReview
          },
					success: (res) => {
            //token失效
            if(res.statusCode === 401){
              uni.removeStorageSync('token');
              uni.showToast({
                title: '登录已过期，请重新登录',
                icon: 'none',
                duration: 2000
              });
              uni.navigateTo({
                url: '../login/login'
              });
            }
						if (res.statusCode === 200||res.statusCode === 404) {
							this.daka_book = res.data.task_today.book_learning;
							this.wordNumLearned = res.data.task_today.word_num_learned;
							this.wordNumTotal = res.data.task_today.word_num_total;
							this.daysLeft = res.data.task_today.days_left;
							this.wordNumToPunch = 10;
							this.wordNumToReview = res.data.task_today.review_num;
              if(res.data.task_today.punch_num==0){
                //说明今天已经打卡完，转到复习模式
                uni.setStorageSync('toReview', true);
                this.isDaka=true;
                this.isReview = false;
                this.wordNumPunched=this.wordNumToPunch;
              }
            } else {
							console.error("请求失败", res);
							this.daka_book = "词汇书123"
						}
					},
					fail: (err) => {
						console.error("请求失败", err);
						this.daka_book = "词汇书123"
					}
				});
			},
			handleSearchShow() {
				this.isHistoryVisible = true;
			},
			handleSearchRouter() {
        // 保存搜索记录
        let searchHistory = uni.getStorageSync('searchHistory');
        if(searchHistory){
          //搜索记录是this.searchInput+"_words"对应的结构体
          searchHistory.push(this.searchHistory[this.searchInput+"_words"]);
          uni.setStorageSync('searchHistory', searchHistory);
        }else{
          uni.setStorageSync('searchHistory', [this.searchHistory[this.searchInput+"_words"]]);
        }
			// 跳转到搜索结果页
			uni.navigateTo({
				// url: `/pages/word_details/word_details?word=${this.searchInput}`
				url: `/pages/word_details/word_details`
			});
			uni.showToast({
				title: '搜索成功',
				icon: 'none'
			});
			console.log("本次搜索内容是", this.searchInput);
		},
		handleSearchInput(input) {
        if(input.length==0){
          this.searchHistory = uni.getStorageSync('searchHistory');
		  return;
        }
		    this.searchInput = input;
        let input_words = this.searchHistory[input+"_words"];
        // 从本地缓存中获取上次的搜索记录
        if(input_words){
          this.items = input_words;
          return;
        }
        //向后端传递搜索结果searchInput
        uni.request({
          url: "http://localhost:8080/api/users/search_words",
          data: {
            input: this.searchInput
          },
          method: 'POST',
          success: (res) => {
            if (res.statusCode === 200) {
              let words = res.data.words;
			  if(words==null) words = [];
              this.items = [];//清空原有记录
              // 遍历words数组，将每一个对象转换为新的结构，并添加到items数组中
              words.forEach(word => {
                this.items.push(this.transformWordObject(word));
              });
              // 缓存搜索结果
              this.searchHistory[input+"_words"]=Array.from(this.items);
            } else {
              console.error("请求失败", res);
            }
          },
          fail: (err) => {
            console.error("请求失败", err);
          }
        });
			},
			cancelSearch() {
				this.isHistoryVisible = false;
				this.searchInput = '';
			}

		}
	}
</script>