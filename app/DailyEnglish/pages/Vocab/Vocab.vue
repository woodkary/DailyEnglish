<template>
	<view class="container">
		<image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
		<view class="vocabook">
			<image class="vocabook-img" src="../../static/book.png"></image>
			<h3 class="vocabook-title">单词书:{{ book }}</h3>
			<view class="vocabook-cnt">生词数：{{ cnt }}</view>
			<view class="button-container"><button class="review" @click="Review">复习</button></view>
			<view class="button-container"><button class="export" @click="Export">导出</button></view>
		</view>
		<view class="word-blocks" @touchend="handleTouchEnd()">
			<word-block v-for="word in words" :key="word.id" :word="word.spelling" :id="word.word_id"
				:pronunciation="word.pronunciation" :meaning="getMeaningStr(word.meanings)" :details="word"
				:review-count="5" :sound="word.sound">
			</word-block>
		</view>
		<view>
			<backTop v-if="showBackTop"></backTop>
		</view>
	</view>
</template>

<script>
	import WordBlock from "../WordBlock/WordBlock.vue";
	import backTop from "@/components/backTop.vue";
	export default {
		components: {
			WordBlock,
			backTop,
		},
		// mounted() {
		// 	this.fetchWords(); //加载单词,mounted()是在页面加载完成后执行的
		// },

		data() {
			return {
				//词性简写
				simplifiedSpeech: {
					verb: "v.",
					adjective: "adj.",
					noun: "n.",
					pronoun: "pron.",
					adverb: "adv.",
					conjunction: "conj.",
					preposition: "prep.",
					interjection: "int."
				},
				words: [], // 单词列表
				cnt: 0,
				book: "cet4",
				startIndex: 0, // 开始加载的索引
				endIndex: 20, // 结束加载的索引
				showBackTop: false, //是否显示返回顶部按钮
			};
		},
    onLoad() {
      this.cnt = this.words.length;
      this.book = "cet4";
      // 显示加载进度条
      uni.showLoading({
        title: '加载中'
      });
      uni.request({
        url: "http://localhost:8080/api/words/get_starbk",
        method: "GET",
        header: {
          'Authorization': 'Bearer ' + uni.getStorageSync('token')
        },
        success: (res) => {
          // 隐藏加载进度条
          uni.hideLoading();
          // token失效
          if (res.statusCode === 401) {
            uni.removeStorageSync('token');
            uni.showToast({
              title: '登录已过期，请重新登录',
              icon: 'none',
              duration: 2000
            });
            uni.navigateTo({
              url: '../logins/login'
            });
          }
          if(res.data.code === 200||res.data.code === "200"){
            res.data.words.forEach(word => {
              let res=word;
              res["sound"]=this.getSoundUrl(res);
              this.words.push(res);
            });
          }
        },
        fail: (res) => {
          // 隐藏加载进度条
          uni.hideLoading();
          console.log("请求失败");
        }
      });
    },

		onPageScroll(e) {
			// 获取滚动的距离
			let scrollTop = e.scrollTop;

			// 如果滚动距离超过100px，显示返回顶部按钮
			if (scrollTop > 100) {
				this.showBackTop = true;
			} else {
				this.showBackTop = false;
			}
		},
		methods: {
      transformWordDescription(simpleDescription) {
        try {
          // 解析输入的简略描述
          const word = JSON.parse(this.fixJsonString(simpleDescription));

          // 创建详细描述的基础结构
          const detailedDescription = {
            word_id: word.word_id,
            spelling: word.spelling,
            pronunciation: word.pronunciation,
            meanings: {
              verb: [],
              adjective: null,
              noun: null,
              pronoun: null,
              adverb: null,
              conjunction: null,
              preposition: null,
              interjection: null
            },
            sound: null // 假设没有提供发音链接
          };

          // 处理meanings部分
          const meaningsArray = word.meanings.split('，');
          meaningsArray.forEach(meaning => {
            const [partOfSpeech, definition] = meaning.split('.');
            if (partOfSpeech && definition) {
              const pos = partOfSpeech.trim();
              const def = definition.trim();
              if (pos === this.simplifiedSpeech.verb) {
                detailedDescription.meanings.verb.push(def);
              } else if (pos === this.simplifiedSpeech.adjective) {
                detailedDescription.meanings.adjective = [def];
              } else if (pos === this.simplifiedSpeech.noun) {
                detailedDescription.meanings.noun = [def];
              }
              // 添加其他词性判断
            }
          });

          return detailedDescription;
        } catch (error) {
          console.error("Error parsing JSON: ", error);
          return null;
        }
      },
      fixJsonString(str) {
        return str
            .replace(/'/g, '"') // 替换单引号为双引号
            .replace(/(\w+):/g, '"$1":') // 给键添加双引号
            .replace(/:([a-zA-Z]+)/g, ':"$1"') // 给值添加双引号（简单情况下，非数组、对象等）
            .replace(/,\s*}/g, '}') // 移除尾随逗号
            .replace(/https":/, 'https:'); // 修正URL的冒号问题
      },
      isSimplifiedDescription(wordString) {
        // 检查meanings是否是字符串类型，判断是否为简略描述
        const fixedWordString = this.fixJsonString(wordString);
        try {
          const parsedWord = JSON.parse(fixedWordString);
          return typeof parsedWord.meanings === 'string';
        } catch (error) {
          console.error("Error parsing word description: ", error);
          return false;
        }
      },
			getSoundUrl(word){
				return `https://ssl.gstatic.com/dictionary/static/sounds/oxford/${word.spelling}--_gb_1.mp3`;
			},
			getMeaningStr(meanings) {
				let meaningStr = "";
				let foundFirst = false; // 标志是否找到了第一个非空词性的意思

				for (let key in meanings) {
					if (meanings[key] && meanings[key].length > 1) {
						if (!foundFirst) {
							// 如果是第一个非空词性，添加词性前缀
							meaningStr += this.simplifiedSpeech[key];
							foundFirst = true;
						} else {
							// 如果不是第一个非空词性，则不再添加词性前缀
							meaningStr += "、";
						}

						// 只添加前两个意思
						meaningStr += meanings[key].slice(0, 2).join("、");

						// 如果意思多于两个，添加 "..."
						if (meanings[key].length > 2) {
							meaningStr += "..."
							break; // 找到了第一个非空词性的前两个意思，结束循环
						} else if (meanings[key].length === 2) {
							meaningStr += "\n"; // 添加换行符，但只有当添加了两个意思时
							break; // 找到了第一个非空词性的前两个意思，结束循环
						}
					}
				}

				return meaningStr;
			},

			handleBack() {
				uni.navigateBack();
			},
			Review() {
				//预计跳转到Review页面，后续再写，内容是复习生词
			},
			Export() {
				//预计导出生词，后续再写
			},
			handleTouchEnd(e) {
				// 获取scroll-view的滚动高度
				const scrollTop = e.target.scrollTop;
				const scrollHeight = e.target.scrollHeight;
				const clientHeight = e.target.clientHeight;

				// 判断是否滑动到了页面底部
				if (scrollTop + clientHeight >= scrollHeight) {
					// 用户已经滑动到了页面底部
					console.log("滑动到底部");
				} else {
					// 这里可以执行其他逻辑，比如弹出提示
					wx.showToast({
						title: "Hello, World!",
						icon: "none", // 设置为'none'可以避免出现加载图标
					});
				}
			},
		},
	};
</script>

<style>
	.container {
		display: grid;
		flex-direction: column;
		/*垂直布局 */
		align-items: center;
		/*水平居中 */
		justify-content: center;
		/*垂直居中 */
		height: 100vh;
		background-image: white;
	}

	.back-icon {
		width: 2rem;
		/*图标宽度 */
		height: 2rem;
		/*图标高度 */
		position: absolute;
		/*绝对定位 */
		top: 0.8rem;
		/*距离顶部20px */
		left: 0.5rem;
		/*距离左侧20px */
		cursor: pointer;
		/*鼠标移上去显示小手 */
	}

	.vocabook {
		position: absolute;
		width: 100%;
		/*宽度100%*/
		display: flex;
		/*flex布局 */
		border-bottom: thick groove #ffff00;
		/*改一下颜色*/
		height: 10rem;
		top: 1rem;
	}

	.vocabook-img {
		width: 6rem;
		height: 7rem;
		margin-left: 2rem;
		/* 在父容器中垂直居中 */
		align-self: center;
	}

	.vocabook-title {
		font-size: 1rem;
		margin-left: 0.5rem;
		margin-top: 1.7rem;
		margin-bottom: 17rem;
	}

	.vocabook-cnt {
		font-size: 0.9rem;
		margin-left: -5.3rem;
		margin-top: 3.8rem;
		color: #a7a7a7;
	}

	.button-container {
		display: flex;
		justify-content: center;
		/*水平居中*/
		align-items: center;

	}

	.review {
		margin-left: -4rem;
		margin-top: 5rem;
		border-radius: 5px;
		width: 4rem;
		height: 2rem;
		font-size: 1rem;
		/*文字居中*/
		text-align: center;
		line-height: 2rem;
		background-color: #456de7;
		color: white;
	}

	.export {
		margin-left: 1rem;
		margin-top: 5rem;
		border: 1px solid #000;
		border-radius: 5px;
		width: 4rem;
		height: 2rem;
		font-size: 1rem;
		/*文字居中*/
		text-align: center;
		line-height: 2rem;
	}

	.word-blocks {
		position: absolute;
		display: flex;
		flex-wrap: wrap;
		justify-content: flex-start;
		/* 或使用 'center', 'flex-end' 等 */
		align-content: flex-start;
		width: 100%;
		/* height: auto; 移除这一行，或者使用 min-height */
		/* min-height: 50vh; */
		/* 根据需要设置最小高度 */
		top: 11.1rem;
	}
</style>