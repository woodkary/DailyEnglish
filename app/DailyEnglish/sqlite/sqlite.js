const dbName = 'dailyenglish';
const dbPath = '_doc/dailyEnglish.db';
export function open() {
  plus.sqlite.openDatabase({
    name: dbName,
    path: dbPath,
    success(e) {
      console.log('打开数据库成功');
      console.log(e)
      resolve(e);
    },
    fail(e) {
      console.log('打开数据库失败');
      reject(e);
    }
  })
}
export function isOpen() {
  // 由于箭头函数没有自己的this，这里直接使用dbName和dbPath
  var open = plus.sqlite.isOpenDatabase({
    name: dbName,
    path: dbPath
  });
  return open;
}
export function close() {
  plus.sqlite.closeDatabase({
    name: dbName,
    success(e) {
      console.log('关闭数据库成功');
      resolve(e);
    },
    fail(e) {
      console.log('关闭数据库失败');
      reject(e);
    }
  })
}
export function deleteTable() {
  const sql = `DROP TABLE word`;
  plus.sqlite.executeSql({
    name: dbName,
    sql: sql,
    success(e) {
      console.log('删除表成功');
      resolve(e);
    },
    fail(e) {
      console.log('删除表失败');
      console.log(e)
      reject(e);
    }
  })
}
export function createWordsTable() {
  //sql语句
  const sql = `CREATE TABLE word (
    new_id INTEGER PRIMARY KEY AUTOINCREMENT,
    word_id INTEGER NOT NULL,
    word TEXT NOT NULL,
    phonetic_us TEXT NOT NULL,
    describe TEXT NOT NULL,
    morpheme TEXT,
    example_sentence TEXT,
    other TEXT,
    word_quetion TEXT,
    answer TEXT,
    learn_times INTEGER NOT NULL,
    interval_history TEXT NOT NULL,
    feedback_history TEXT NOT NULL,
    review_date DATE NOT NULL,
    difficulty INTEGER NOT NULL,
    is_memory INTEGER NOT NULL
  )`;
  plus.sqlite.executeSql({
    name: dbName,
    sql: sql,
    success(e) {
      console.log('创建表成功');
      resolve(e);
    },
    fail(e) {
      console.log('创建表失败');
      console.log(e)
      reject(e);
    }
  })
}
export function insertword(words) {
  //接收word对象
  const sql = `INSERT INTO word (word_id, word, phonetic_us, describe, morpheme, example_sentence, other, word_quetion, answer, learn_times, interval_history, feedback_history, interval_days, difficulty, is_memory) VALUES (${words.word_id}, '${words.word}', '${words.phonetic_us}', '${words.describe}', '${words.morpheme}', '${words.example_sentence}', '${words.other}', '${words.word_quetion}', '${words.answer}', ${words.learn_times}, '${words.interval_history}', '${words.feedback_history}', ${words.interval_days}, ${words.difficulty}, ${words.is_memory})`;
  plus.sqlite.executeSql({
    name: dbName,
    sql: sql,
    success(e) {
      console.log('插入数据成功');
      resolve(e);
    },
    fail(e) {
      console.log('插入数据失败');
      console.log(e)
      reject(e);
    }
  })
}
//插入一本书
export function insertBook(book) {
  for (var i = 0; i < book.length; i++) {
    const sql = `INSERT INTO word (word_id, word, phonetic_us, describe, morpheme, example_sentence, other, word_quetion, answer, learn_times, interval_history, feedback_history, interval_days, difficulty, is_memory) VALUES (${book[i].word_id}, '${book[i].word}', '${book[i].phonetic_us}', '${book[i].describe}', '${book[i].morpheme}', '${book[i].example_sentence}', '${book[i].other}', '${book[i].word_quetion}', '${book[i].answer}', ${book[i].learn_times}, '${book[i].interval_history}', '${book[i].feedback_history}', ${book[i].interval_days}, ${book[i].difficulty}, ${book[i].is_memory})`;
    plus.sqlite.executeSql({
      name: dbName,
      sql: sql,
      success(e) {
        console.log('插入数据成功');
        resolve(e);
      },
      fail(e) {
        console.log('插入数据失败');
        console.log(e)
        reject(e);
      }
    })
  }
}
//打卡单词--查询new_id在learn_index-(learn_index+plan_words)之间的单词
export function selectWords(learn_index, plan_words) {
  const sql = `SELECT * FROM word WHERE new_id BETWEEN ${learn_index} AND ${learn_index + plan_words}`;
  plus.sqlite.selectSql({
    name: dbName,
    sql: sql,
    success(data) {
      console.log('查询数据成功');
      console.log(data)
      resolve(data);
      return data;
    },
    fail(e) {
      console.log('查询数据失败');
      reject(e);
    }
  })
}
//复习单词，查询review_date在当前日期之前的单词
export function selectReviewWords() {
  const sql = `SELECT * FROM word WHERE review_date < ${new Date().getTime()}`;
  plus.sqlite.selectSql({
    name: dbName,
    sql: sql,
    success(data) {
      console.log('查询数据成功');
      console.log(data)
      resolve(data);
      return data;
    },
    fail(e) {
      console.log('查询数据失败');
      reject(e);
    }
  })
}
export function selectword() {
  const sql = `SELECT * FROM word`;
  plus.sqlite.selectSql({
    name: dbName,
    sql: sql,
    success(data) {
      console.log('查询数据成功');
      console.log(data)
      resolve(data);
      return data;
    },
    fail(e) {
      console.log('查询数据失败');
      reject(e);
    }
  })
}