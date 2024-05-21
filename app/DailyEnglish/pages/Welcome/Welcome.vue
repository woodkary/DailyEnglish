<template>
	<view>

		<view class="topbar">
			<span class="host-title">选择您的打卡计划</span>
			<span class="skip-container">
				<a href="" class="skip">跳过</a>
			</span>
		</view>
		<!-- 添加选项卡 -->
		<view class="tab-bar-container">
			<view class="tab-bar">
				<view class="tab-item" :class="{ 'active': activeTab === index }" v-for="(tab, index) in tabs"
					:key="index" @click="changeTab(index)">
					{{ tab }}
				</view>
			</view>
		</view>

		<!-- 添加按钮行 -->
		<view class="button-row">
			<view class="button" v-for="(desc, id) in buttonIdMap" :key="id" :class="{ 'active': activeButton === id }" @click="scrollToIdButton(id)">
        {{ desc }}
			</view>
			<!-- 添加其他按钮 -->
		</view>

		<view class="book-list">
			<view class="book-type" v-for="(desc, id) in buttonIdMap" :key="id" :id="id">
				<view class="book-container" v-for="(book, index) in filteredBooks(id)" :key="index" @click="bookConfirm(book.book_id,book.title)">
					<image :src="getImageUrl(id)"></image>
					<view class="text-container">
						<span class="book-title">{{book.title}}</span>
						<span class="discrip">{{book.decsrip}}</span>
						<span class="num">{{book.num}}</span>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
export default {
  data() {
    return {
      // 按钮id映射表
      buttonIdMap: {
        'cet4': '四级',
        'cet6': '六级'
      },
      books: [
        {
          book_id: 1,
          title: '四级词汇大全',
          decsrip: '四级最新考纲单词全收录，时候所有备考四级的同学',
          grade: '四级',
          num: '共4440词'
        },
        {
          book_id: 2,
          title: '四级高频',
          decsrip: '精选四级真题超高频词',
          grade: '四级',
          num: '共739词'
        },
        {
          book_id: 3,
          title: '四级高频',
          decsrip: '精选四级真题超高频词',
          grade: '四级',
          num: '共739词'
        },
        {
          book_id: 4,
          title: '四级高频',
          decsrip: '精选四级真题超高频词',
          grade: '四级',
          num: '共739词'
        },
        {
          book_id: 5,
          title: '四级高频',
          decsrip: '精选四级真题超高频词',
          grade: '四级',
          num: '共739词'
        },
        {
          book_id: 6,
          title: '四级高频',
          decsrip: '精选四级真题超高频词',
          grade: '四级',
          num: '共739词'
        },
        {
          book_id: 7,
          title: '六级词汇大全',
          decsrip: '六级最新考纲单词全收录，时候所有备考六级的同学',
          grade: '六级',
          num: '共6204词'
        },
        {
          book_id: 8,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 9,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 10,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 11,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 12,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 13,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 14,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 15,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 16,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 17,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 18,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        },
        {
          book_id: 19,
          title: '六级核心（过考版）',
          decsrip: '精选六级真题超高频词',
          grade: '六级',
          num: '共2551词'
        }
      ], // 词书列表
      activeTab: 0, // 默认选中第一个选项卡
      tabs: ["全部", "大学", "高中", "初中", "小学", "留学", "其他"], // 选项卡数组
      activeButton: null, // 记录当前活跃的按钮,
      operation: 0 // 记录下一个应该跳转的页面
    };
  },
  onLoad(event){
    this.operation = parseInt(event.operation);
    //清空books数组
    this.books = [];
    //清空按钮id映射表
    this.buttonIdMap = {};
    //获取所有词书
    uni.request({
      url: '/api/users/navigate_books',
      method: 'GET',
      header: {
        'Authorization': 'Bearer '+uni.getStorageSync('token')
      },
      success: (res) => {
        console.log(res.data);
        let books = res.data.data;
        //将词书信息添加到books数组中
        books.forEach(book => {
          this.books.push({
            book_id: book.book_id,
            title: book.book_name,
            decsrip: book.description,
            grade: book.grade_description,
            num: book.word_num,
          });
          // 按钮id映射表，书的等级对应按钮的id
          this.buttonIdMap[book.grade] = book.grade_description;
        });
      },
      fail: (err) => {
        console.log(err)
      }
    });
  },
  methods: {
    getImageUrl(id){
      return `../../static/${id}.jpg`;
    },
    filteredBooks(grade){
      return this.books.filter(book => book.grade === this.buttonIdMap[grade]);
    },
    changeTab(tabNumber) {
      this.activeTab = tabNumber;
    },
    scrollToIdButton(id) {
      // 点击按钮，滚动到四级部分
      this.scrollToElement(id);
      this.activeButton = id; // 设置活跃的按钮当前的id
    },
    scrollToElement(id) {
      const el = document.getElementById(id);
      if (el) {
        const offset = 110; // 与顶部保持的距离
        const scrollPosition = el.offsetTop - offset;
        window.scrollTo({
          top: scrollPosition,
          behavior: 'smooth'
        });
      }
    },
    bookConfirm(book_id,title){
      console.log(book_id,title);
      uni.showModal({
        title: '提示',
        content: '确定要选择《'+title+'》吗？',
        success: (res) => {
          if (res.confirm) {
            uni.request({
              url: '/api/users/confirm_book_to_learn',
              method: 'PUT',
              header: {
                'Authorization': 'Bearer '+uni.getStorageSync('token')
              },
              data: {
                book_id: book_id
              },
              success: (res) => {
                console.log(res.data);
                if (res.data.code === 200) {
                  uni.showToast({
                    title: '选择词库成功',
                    icon: 'none'
                  });
                  setTimeout(() => {
                    if(this.operation===0){
                      uni.switchTab({
                        url: '../home/home'
                      });
                      return;
                    }
                    uni.navigateBack();
                  });
                } else {
                  uni.showToast({
                    title: '选择词库失败',
                    icon: 'none'
                  });
                }
              },
              fail: (err) => {
                console.log(err)
              }
            });
          } else if (res.cancel) {
            console.log('取消')
          }
        }
      });
    }
  }
};
</script>

<style>
	body {
		background-color: ;
		text-align: center;
	}

	.topbar {
		position: sticky;
		top: 0;
		z-index: 1000;
		/* 确保在最顶层 */
		display: flex;
		justify-content: space-between;
		z-index: 1000;
		/* 确保在最顶层 */
		/* 将元素分散对齐 */
		align-items: center;
		/* 垂直居中 */
		background-color: white;
		width: 100%;
		padding: 10px;

	}

	.host-title {
		font-family: Cambria, Cochin, Georgia, Times, 'Times New Roman', serif;
		font-size: 45rpx;
		width: 100%;
	}

	.skip-container {
		display: flex;

	}

	.skip {
		margin-right: 20px;
		margin-bottom: 0;
		text-decoration: none;
		white-space: nowrap;
		/* 取消下划线 */
		font-size: 16px;
		/* 适当调整字体大小 */
	}


	.tab-bar-container {
		background-color: white;
		position: sticky;
		top: 47px;
		/* 顶部导航栏的高度 */
		z-index: 1000;
		/* 确保在最顶层 */
		overflow-x: overlay;
		/* 隐藏滚动条，但保留滑动功能 */
		white-space: nowrap;
		/* 不换行 */
		border-bottom: 0.5rpx solid grey;
	}

	.tab-bar-container::-webkit-scrollbar {
		display: none;
		/* 隐藏滚动条 */
	}

	.tab-bar {
		display: inline-block;
		/* 内联块元素 */
	}

	.tab-item {
		display: inline-block;
		/* 内联块元素 */
		padding: 10px 20px;
		color: gray;
		/* 初始字体颜色 */
		cursor: pointer;
		font-size: 35rpx;
	}

	.tab-item.active {
		color: black;
		/* 选中时字体颜色 */
	}

	.tab-item.active::after {
		content: '';
		display: block;

		width: 50%;
		height: 3px;
		background-color: blue;
		/* 选中时底边框颜色 */
		margin: 0 auto;
		/* 将下划线水平居中 */
	}

	.button-row {
		position: sticky;
		background-color: white;
		top: 90px;
		/* 顶部导航栏和选项卡的高度 */
		z-index: 1000;
		/* 确保在最顶层 */
		display: flex;
		justify-content: center;
		margin-top: 20px;
	}

	.button {
		margin: 0 10px;
		padding: 5px 10px;
		cursor: pointer;
		background-color: #f8f8f8;
		color: gray;
		border: 1px solid #ccc;
		border-radius: 4px;
	}

	.button.active {
		background-color: blue;
		color: white;
	}

	.book-type-span {
		float: left;
		font-size: 25px;
		margin-left: 50rpx;
	}

	.book-container {
		display: flex;
		width: 100%;
		height: 220rpx;
		align-items: center;
		/* 垂直居中 */
	}

	.book-container image {

		width: 150rpx;
		height: 180rpx;
		margin: 10rpx 40rpx 5rpx 60rpx;
		border-radius: 5%;
	}

	.text-container {
		width: calc(100% - 260px);
		/* 调整宽度，确保留足够的空间给图片 */
		display: flex;
		flex-direction: column;
		align-items: flex-start;
	}

	.book-title {
		font-size: 20px;
		margin-top: 10rpx;
		white-space: nowrap;
	}

	.discrip {
		color: gray;
		margin-top: 10rpx;
		overflow: hidden;
		white-space: nowrap;

	}

	.num {
		margin-top: 35rpx;

	}
</style>