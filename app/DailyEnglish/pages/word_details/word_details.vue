<template>
	<view>
		<view class="word-list-container">
			<view class="word-item" v-for="(word, index) in words" :key="index">
				<view class="word-card">
					<h1 class="word">{{ word.name }}</h1>
					<image src="../../static/collection.png" class="collection"></image>
				</view>
				<view class="pronunciation-card">
					<p>{{ word.pronunciation }}</p>
					<image src="../../static/pronunciation.png" class="pronunciation"></image>
				</view>
				<p>{{ word.meaning }}</p>
				<p>{{ word.meaning2 }}</p>
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
		<CollapsibleView ref="collapsibleView2" v-for="(word, index) in words2" :key="index">
			<template #header>
				<view class="header">
					<h3>详细释义</h3>
					<image @click="toggleCollapse2" src="../../static/up.png" class="up-arrow"></image>
				</view>
			</template>
			<view>
				<view>
					<view style="display: flex; justify-content: flex-start;">
						<p class="partofspeech">{{ word.partofspeech1 }}</p>
						<p>{{ word.meaning1 }}</p>
					</view>
					<view style="display: flex; flex-direction: row; align-items: center;">
						<p class="example-text">{{ word.example1 }}</p>
						<image class="example-image" src="../../static/pronunciation.png" alt="Example Image"></image>
					</view>
					<p class="meaning">{{word.example1meaning}}</p>
				</view>
				<view>
					<view style="display: flex; justify-content: flex-start;">
						<p class="partofspeech">{{ word.partofspeech2 }}</p>
						<p>{{ word.meaning2 }}</p>
					</view>
					<view style="display: flex; flex-direction: row; align-items: center;">
						<p class="example-text">{{ word.example2 }}</p>
						<image class="example-image" src="../../static/pronunciation.png" alt="Example Image"></image>
					</view>
					<p class="meaning">{{word.example2meaning}}</p>
				</view>
				<view>
					<view style="display: flex; justify-content: flex-start;">
						<p class="partofspeech">{{ word.partofspeech3 }}</p>
						<p>{{ word.meaning3 }}</p>
					</view>
				</view>
				<view>
					<view style="display: flex; justify-content: flex-start;">
						<p class="partofspeech">{{ word.partofspeech4 }}</p>
						<p>{{ word.meaning4 }}</p>
					</view>
					<view style="display: flex; flex-direction: row; align-items: center;">
						<p class="example-text">{{ word.example4 }}</p>
						<image class="example-image" src="../../static/pronunciation.png" alt="Example Image"></image>
					</view>
					<p class="meaning">{{word.example4meaning}}</p>
				</view>
				<view>
					<view style="display: flex; justify-content: flex-start;">
						<p class="partofspeech">{{ word.partofspeech5 }}</p>
						<p>{{ word.meaning5 }}</p>
					</view>
				</view>
			</view>
		</CollapsibleView>
		<CollapsibleView ref="collapsibleView3" v-for="(word, index) in words3" :key="index">
			<template #header>
				<view class="header">
					<h3>短语</h3>
					<image @click="toggleCollapse3" src="../../static/up.png" class="up-arrow"></image>
				</view>
			</template>
			<view style="text-align: left;">
				<view>
					<p class="phrase">{{ word.phrase1 }}</p>
					<p class="phrase-meaning">{{ word.meaning1 }}</p>
				</view>
				<view>
					<p class="phrase">{{ word.phrase2 }}</p>
					<p class="phrase-meaning">{{ word.meaning2 }}</p>
				</view>
				<view>
					<p class="phrase">{{ word.phrase3 }}</p>
					<p class="phrase-meaning">{{ word.meaning3 }}</p>
				</view>
				<view>
					<p class="phrase">{{ word.phrase4 }}</p>
					<p class="phrase-meaning">{{ word.meaning4 }}</p>
				</view>
				<view>
					<p class="phrase">{{ word.phrase5 }}</p>
					<p class="phrase-meaning">{{ word.meaning5 }}</p>
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
				words: [{
						name: 'abandon',
						pronunciation: '/ əˈbændən /',
						meaning: 'v.    抛弃；放弃；沉湎于（某种情感）；舍弃，废弃',
						meaning2: 'n.    尽情，放纵',
						book: '高中/ CET4 / CET6 / 考研 / IELTS / TOEFL / GRE'
					},
					// 更多单词...  
				],
				words2: [{
						partofspeech1: 'v.',
						meaning1: '抛弃',
						example1: 'Should I tell you to abandon me and save yourself, you must to do so. ',
						example1meaning: '我若是让你别管我，救自己，你也必须照做。',
						partofspeech2: 'v.',
						meaning2: '放弃',
						example2: 'The girl has totally abandoned the use of computer for her homework.',
						example2meaning: '这个女生彻底放弃使用电脑做作业了。',
						partofspeech3: 'v.',
						meaning3: '沉湎于（某种情感）',
						partofspeech4: 'v.',
						meaning4: '舍弃，废弃',
						example4: 'The mining factory was abandoned a long time ago.',
						example4meaning: '这个采矿工厂早已被放弃。',
						partofspeech5: 'n.',
						meaning5: '尽情，放纵',
					},

					// 更多单词...  
				],
				words3: [{
						phrase1: 'abandon oneself to',
						meaning1: '沉溺于',
						phrase2: 'with abandon',
						meaning2: '恣意地',
						phrase3: 'abandon doing sth.',
						meaning3: '放弃做某事',
						phrase4: 'abandon ones belief',
						meaning4: '放弃信仰',
						phrase5: 'abandon to',
						meaning5: '离弃，遗弃，抛弃',
					},
					// 更多单词...  
				],
			}
		},
		methods: {
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