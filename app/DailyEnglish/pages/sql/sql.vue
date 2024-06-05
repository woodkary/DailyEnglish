<template>
  <div>
    <button @click="openDatabase">打开数据库</button>
    <button @click="closeDatabase">关闭数据库</button>
    <button @click="createTable">创建表</button>
    <button @click="insertWord">插入单词</button>
    <button @click="selectWords">查询单词</button>

    <div>
      <h2>单词列表：</h2>
      <table>
        <thead>
          <tr>
            <th>new_id</th>
            <th>word_id</th>
            <th>word</th>
            <th>phonetic_us</th>
            <th>describe</th>
            <!-- 添加其他列的表头 -->
          </tr>
        </thead>
        <tbody>
          <tr v-for="word in words" :key="word.new_id">
            <td>{{ word.new_id }}</td>
            <td>{{ word.word_id }}</td>
            <td>{{ word.word }}</td>
            <td>{{ word.phonetic_us }}</td>
            <td>{{ word.describe }}</td>
            <!-- 显示其他列的数据 -->
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import { open, close, createWordsTable, insertword, selectword } from '../../sqlite/sqlite.js'; // 替换为您的 SQL.js 文件路径

export default {
  data() {
    return {
      words: []
    };
  },
  methods: {
    async openDatabase() {
      try {
        await open();
        console.log('数据库已打开');
      } catch (error) {
        console.error(`打开数据库失败: ${error.message}`);
      }
    },
    async closeDatabase() {
      try {
        await close();
        console.log('数据库已关闭');
      } catch (error) {
        console.error(`关闭数据库失败: ${error.message}`);
      }
    },
    async createTable() {
      try {
        await createWordsTable();
        console.log('表已创建');
      } catch (error) {
        console.error(`创建表失败: ${error.message}`);
      }
    },
    async insertWord() {
      const word = {
        word_id: 1,
        word: 'apple',
        phonetic_us: '/',
        describe: '苹果',
        morpheme: 'ap+ple',
		example_sentence: 'I like apple',
		other: 'n. 苹果；苹果树；苹果似的东西；[美俚]头；[美俚]脑袋；[美俚]人；[美俚]家伙；[美俚]家庭；[美俚]公司；[美俚]事情；[美俚]事业；[美俚]目标；[美俚]目的；[美俚]目的地',
		word_quetion: 'apple',
		answer: '苹果',
		learn_times: 0,
		interval_history: '0',
		feedback_history: '0',
		interval_days: 1,
		difficulty:1,
		is_memory: 0,

      };
      try {
        await insertword(word);
        console.log('单词已插入');
      } catch (error) {
        console.error(`插入单词失败: ${error.message}`);
      }
    },
    async selectWords() {
      try {
        this.words = await selectword();
		console.log(this.words);
        console.log('查询成功');
      } catch (error) {
        console.error(`查询失败: ${error.message}`);
      }
    }
  }
};
</script>

<style>
/* 添加一些样式（如果需要） */
table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  border: 1px solid #ddd;
  padding: 8px;
  text-align: left;
}

th {
  background-color: #f2f2f2;
}
</style>
