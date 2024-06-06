<template>
  <view>
    <view class="word-list-container">
      <view class="word-item">
        <view class="word-card">
          <h1 class="word">{{ word.name }}</h1>
          <image src="../../static/collection.png" class="collection"></image>
        </view>
        <view class="pronunciation-card">
          <p>{{ word.pronunciation }}</p>
          <image src="../../static/pronunciation.png" class="pronunciation"></image>
        </view>
        <p v-for="meaning in word.meanings">{{ meaning }}</p>
        <p class="book">{{ word.book}}</p>
      </view>
    </view>
    <CollapsibleView ref="collapsibleView1">
      <template #header>
        <view class="header">
          <h3>单词变形</h3>
          <image @click="toggleCollapse1" src="../../static/up.png" class="up-arrow"></image>
        </view>
      </template>
      <view class="content">这里是第一个可展开/缩起的内容</view>
    </CollapsibleView>
    <CollapsibleView ref="collapsibleView2">
      <template #header>
        <view class="header">
          <h3>详细释义</h3>
          <image @click="toggleCollapse2" src="../../static/up.png" class="up-arrow"></image>
        </view>
      </template>
      <view>
        <view v-for="detail in wordDetail.details" :key="detail.partOfSpeech">
          <view style="display: flex; justify-content: flex-start;">
            <p class="partofspeech">{{ detail.partOfSpeech }}</p>
            <p>{{ detail.chineseMeaning }}</p>
          </view>
          <view style="display: flex; flex-direction: row; align-items: center;">
            <p class="example-text">{{ detail.exampleSentence }}</p>
            <image class="example-image" src="../../static/pronunciation.png" alt="Example Image"></image>
          </view>
          <p class="meaning">{{detail.sentenceMeaning}}</p>
        </view>
      </view>
    </CollapsibleView>
    <CollapsibleView ref="collapsibleView3">
      <template #header>
        <view class="header">
          <h3>短语</h3>
          <image @click="toggleCollapse3" src="../../static/up.png" class="up-arrow"></image>
        </view>
      </template>
      <view style="text-align: left;">
        <view v-for="phraseAndMeaning in wordAndPhrase.phraseAndMeanings" :key="phraseAndMeaning.phrase">
          <p class="phrase">{{ phraseAndMeaning.phrase }}</p>
          <p class="phrase-meaning">{{ phraseAndMeaning.meaning }}</p>
        </view>
      </view>
    </CollapsibleView>
  </view>
</template>

<script>
import CollapsibleView from '../../components/CollapsibleView.vue';

export default {
  components: {
    CollapsibleView
  },
  data() {
    return {
      //词性简写转换表
      simplifiedSpeech:{
        verb: "v.",
        adjective: "adj.",
        noun: "n.",
        pronoun: "pron.",
        adverb: "adv.",
        conjunction: "conj.",
        preposition: "prep.",
        interjection: "int."
      },
      word: {
        name: 'abandon',
        pronunciation: '/ əˈbændən /',
        meanings: ['v.    抛弃；放弃；沉湎于（某种情感）；舍弃，废弃','n.    尽情，放纵'],
        book: '高中/ CET4 / CET6 / 考研 / IELTS / TOEFL / GRE'
      },
      //这是一个数组，元素是结构体，有word、details两个属性，分别代表单词和详细释义
      //其中details是一个数组，元素是结构体，有partOfSpeech、chineseMeaning、exampleSentence、sentenceMeaning四个属性，分别代表词性、中文释义、例句、英文释义
      wordDetail:
        {
          word: 'abandon',
          details: [{
            partOfSpeech: 'v.',
            chineseMeaning: '抛弃',
            exampleSentence: 'Should I tell you to abandon me and save yourself, you must to do so. ',
            sentenceMeaning: '我若是让你别管我，救自己，你也必须照做。',
          },
          {
            partOfSpeech: 'v.',
            chineseMeaning: '放弃',
            exampleSentence: 'The girl has totally abandoned the use of computer for her homework.',
            sentenceMeaning: '这个女生彻底放弃使用电脑做作业了。',
          },
          {
            partOfSpeech: 'v.',
            chineseMeaning: '沉湎于（某种情感）',
          },
          {
            partOfSpeech: 'v.',
            chineseMeaning: '舍弃，废弃',
            exampleSentence: 'The mining factory was abandoned a long time ago.',
            sentenceMeaning: '这个采矿工厂早已被放弃。',
          },
          {
            partOfSpeech: 'n.',
            chineseMeaning: '尽情，放纵',
          }
          ],
        },
      //wordAndPhrases是一个数组，元素是结构体，有word、phraseAndMeanings两个属性，分别代表单词和短语释义
      //其中phraseAndMeanings是一个数组，元素是结构体，有phrase、meaning两个属性，分别代表短语和释义
      wordAndPhrase: {
        word: 'abandon',
        phraseAndMeanings:[
          {
            phrase: 'abandon oneself to',
            meaning: '沉溺于',
          },
          {
            phrase: 'with abandon',
            meaning: '恣意地',
          },
          {
            phrase: 'abandon doing sth.',
            meaning: '放弃做某事',
          },
          {
            phrase: 'abandon ones belief',
            meaning: '放弃信仰',
          },
          {
            phrase: 'abandon to',
            meaning: '离弃，遗弃，抛弃',
          },
        ]
      },
    }
  },
  onLoad(event) {
    //获取请求参数中的word和word_id，从Vocab页面中传来
    let word=event["word"];
    let word_id=event["word_id"];
    this.word.name=word;
    this.wordDetail.word=word;
    this.wordAndPhrase.word=word;
    //从本地缓存中获取word的信息
    let localDetails=uni.getStorageSync(word_id);
    //localDetails结构如下
    /**
     * {
     *           word_id: 2,
     *           spelling: "abandon",
     *           pronunciation: "/əˈbændən/",
     *           meanings:{
     *             verb:["抛弃","放弃","弃置","放弃治疗"],
     *             noun:["放弃物","放弃的事物","放弃的念头","放弃的决定"],
     *             pronoun:null,
     *             adverb:null,
     *             conjunction:null,
     *             preposition:null,
     *             interjection:null
     *           },
     *           sound:"https://ssl.gstatic.com/dictionary/static/sounds/oxford/abandon--_gb_1.mp3"
     *         }
     */
    if(localDetails){
      //获取发音和释义
      this.word.pronunciation=localDetails.pronunciation;
      //转换为如下格式：
      //meanings: ['v.    抛弃；放弃；沉湎于（某种情感）；舍弃，废弃','n.    尽情，放纵']
      this.word.meanings=this.transformMeaningsToText(localDetails.meanings);
    }
    uni.request({
      url: 'http://localhost:8080/api/words/word_details',
      method: 'POST',
      header: {
        'content-type': 'application/json', // 默认值
        'Authorization': 'Bearer '+uni.getStorageSync('token') // 登录后获取的token
      },
      data: {
        word_id: word_id
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
        //发送请求，获取详细释义和短语释义
        let detailedMeanings=res.data.detailed_meanings;
        //原格式：
        /*
    "detailed_meanings": {
		"verb": [
			{
				"chinese_meaning": "抛弃",
				"example_sentence": "Should I tell you to abandon me and save yourself, you must to do so. ",
				"sentence_meaning": "我若是让你别管我，救自己，你也必须照做。",
				"sentence_sound": "https://example1.mp3"
			},
			{
				"chinese_meaning": "放弃",
				"example_sentence": "The girl has totally abandoned the use of computer for her homework.",
				"sentence_meaning": "这个女生彻底放弃使用电脑做作业了。",
				"sentence_sound": "https://example2.mp3"
			}
		],
		"noun": [
			{
				"chinese_meaning": "放任，放纵",
				"example_sentence": "she sings and sways with total abandon.",
				"sentence_meaning": "她纵情地边唱边晃。",
				"sentence_sound": "https://example3.mp3"
			}
		],
        "pronoun": null,
        "adjective": null,
        "adverb": null,
        "preposition": null,
        "conjunction": null,
        "interjection": null
	},*/
        this.wordDetail.details=this.transformDetailedMeaningsToDetails(detailedMeanings);
        this.wordAndPhrase.phraseAndMeanings=res.data.phrases;
        this.word.book=res.data.word_book;
      },
      fail: (res) => {
        console.log(res);
      }
    });
  },
  methods: {
      transformMeaningsToText(meanings) {
          const transformedMeanings = [];
          for (const speech in meanings) {
            if (meanings[speech] && meanings[speech].length > 0) {
              // 获取词性简写
              const speechAbbreviation = this.simplifiedSpeech[speech];
              // 生成词性及其意思的字符串，之间加入两个制表符
              const meaningText = `${speechAbbreviation}\t\t${meanings[speech].join("; ")}\n`;
              // 添加到数组中
              transformedMeanings.push(meaningText);
            }
          }
          return transformedMeanings;
    },
    transformDetailedMeaningsToDetails(detailedMeanings) {
      const details = [];
      for (const speech in detailedMeanings) {
        if (detailedMeanings[speech] && detailedMeanings[speech].length > 0) {
          // 获取词性简写
          const partOfSpeech = this.simplifiedSpeech[speech];
          detailedMeanings[speech].forEach(meaning => {
            // 创建一个对象，包含词性、中文意思、示例句子和句子意思
            const detail = {
              partOfSpeech: partOfSpeech,
              chineseMeaning: meaning.chinese_meaning,
              exampleSentence: meaning.example_sentence,
              sentenceMeaning: meaning.sentence_meaning
            };
            // 添加到details数组中
            details.push(detail);
          });
        }
      }
      return details;
    },
  toggleCollapse1() {

    },
    toggleCollapse2() {

    },
    toggleCollapse3() {

    }
  }
}
</script>

<style>
.word-list-container {
  width: 90%;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  /* 设置为列方向，实现从上往下排列 */
  align-items: stretch;
  /* 子项默认拉伸以填充整个容器 */
  border-bottom: 2px solid #e6e6e6;
}

.word-card {
  display: flex;
  /* 设置为 flex 容器，以便子元素可以水平排列 */
  justify-content: space-between;
  /* 子元素在主轴上均匀分布，左边和右边有空间 */
  align-items: flex-start;
  /* 子元素在交叉轴上对齐到开始位置 */
}

.pronunciation-card {
  display: flex;
  /* 设置为 flex 容器，以便子元素可以水平排列 */
  align-items: flex-start;
  /* 子元素在交叉轴上对齐到开始位置 */
}

.collection {
  margin-top: 3.5rem;
  margin-right: 1rem;
  width: 1.3rem;
  height: 1.3rem;
}

.word {
  margin-top: 3rem;
  margin-bottom: 10px;
  font-size: 2rem;
}

.pronunciation {
  margin-top: -0.2rem;
  margin-left: 0.5rem;
  width: 1.6rem;
  height: 1.6rem;
}

p {
  margin-bottom: 10px;
}


.book {
  font-size: 14px;
  color: #A7A7A7;
}

.container {
  width: 90%;
  margin: 0 auto;
  height: 5rem;
  padding: 10px;
}

.header {
  display: flex;
  justify-content: space-between;
  /* 让子元素在主轴上均匀分布，第一个在左边，最后一个在右边 */
  align-items: center;
  /* 交叉轴上居中，使得文本和图片垂直对齐 */
  margin-bottom: 10px;
}

.header-text {
  /* 文本样式 */
}

.up-arrow {
  height: 1rem;
  width: 1rem;
}

.content {
  /* 你可以添加更多样式来控制内容的外观 */
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  /* 设置为列方向，实现从上往下排列 */
  align-items: stretch;
  /* 子项默认拉伸以填充整个容器 */
  overflow: hidden;/* 隐藏超出内容区域的内容 */
  width:100%;
  border-bottom: 1px solid #e6e6e6;
}

.partofspeech {
  margin-right: 15px;
  background-color: #00ffff;
  color: #456de7;
  padding: 0 5px 2px 5px;
  border-radius: 5px;
}

.meaning {
  color: #A7A7A7;
}

.example-text {
  flex: 1 1 80%;
}

.example-image {
  margin-top: -1.5rem;
  height: 1.8rem;
  width: 2.3rem;
}

.phrase {
  color: #456de7;
}

.phrase-meaning {
  color: #A7A7A7;
}
</style>