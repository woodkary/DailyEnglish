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
    interval_days INTEGER NOT NULL,
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
export function insertword(word) {
  //接收word对象
  //除了new_id其他字段都要插入数据,没有就空着
  const sql = `INSERT INTO word (word_id, word, phonetic_us, describe, morpheme, example_sentence, other, word_quetion, answer, learn_times, interval_history, feedback_history, interval_days,review_date, difficulty, is_memory) VALUES (${word.word_id}, '${word.word}', '${word.phonetic_us}', '${word.describe}', '${word.morpheme}', '${word.example_sentence}', '${word.other}', '${word.word_quetion}', '${word.answer}', ${word.learn_times}, '${word.interval_history}', '${word.feedback_history}', ${word.interval_days},${word.review_date} ${word.difficulty}, ${word.is_memory})`;
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
    const sql = `INSERT INTO word (word_id, word, phonetic_us, describe, morpheme, example_sentence, other, word_quetion, answer, learn_times, interval_history, feedback_history, interval_days,review_date, difficulty, is_memory) VALUES (${book[i].word_id}, '${book[i].word}', '${book[i].phonetic_us}', '${book[i].describe}', '${book[i].morpheme}', '${book[i].example_sentence}', '${book[i].other}', '${book[i].word_quetion}', '${book[i].answer}', ${book[i].learn_times}, '${book[i].interval_history}', '${book[i].feedback_history}', ${book[i].interval_days},${book[i].review_date}, ${book[i].difficulty}, ${book[i].is_memory})`;
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
//打卡或复习完处理单词
export function updateWord(results) {
  for (var i = 0; i < results.length; i++) {
    //根据results[i].new_id查询数据库获得word信息
    const sql_1 = `SELECT * FROM word WHERE new_id=${results[i].new_id}`;
    plus.sqlite.selectSql({
      name: dbName,
      sql: sql_1,
      success(data) {
        console.log('查询数据成功');
        console.log(data)
        //先调用calculateBestInterval计算复习间隔
        const best_interval = calculateBestInterval(data.difficulty, data.interval_history, data.feedback_history);
        //更新复习日期
        const review_date = new Date().getTime() + best_interval * 24 * 60 * 60 * 1000;
        //更新复习次数
        const learn_times = data.learn_times + 1;
        //更新复习间隔
        const interval_days = best_interval;
        //更新复习历史
        const interval_history = data.interval_history + ',' + data.interval_days;
        //更新反馈历史
        const feedback_history = data.feedback_history + ',' + results[i].is_memory;
        const sql = `UPDATE word SET learn_times=${learn_times}, interval_history='${interval_history}', feedback_history='${feedback_history}', interval_days=${interval_days}, review_date=${review_date} WHERE new_id=${results[i].new_id}`;
        plus.sqlite.executeSql({
          name: dbName,
          sql: sql,
          success(e) {
            console.log('更新数据成功');
            resolve(e);
          },
          fail(e) {
            console.log('更新数据失败');
            console.log(e)
            reject(e);
          }
        })
      },
      fail(e) {
        console.log('查询数据失败');
        reject(e);
      }
    })
  }
}
//计算复习间隔
function calculateBestInterval(d, interval_history, feedback_history) {
  const iterations = 200000;
  const max_index = 122;
  const min_index = -30;
  const base = 1.05;
  const recall_cost = 3.0;
  const forget_cost = 9.0;
  const d_limit = 18;
  const d_offset = 2;

  function cal_start_halflife(difficulty) {
    let p = Math.max(0.925 - 0.05 * difficulty, 0.025);
    return -1 / Math.log2(p);
  }

  function cal_next_recall_halflife(h, p, d, recall) {
    if (recall === 1) {
      return h * (1 + Math.exp(3.81140723) * Math.pow(d, -0.5345194) * Math.pow(h, -0.12641492) * Math.pow(1 - p, 0.97043354));
    } else {
      return Math.exp(-0.04141891) * Math.pow(d, -0.04074844) * Math.pow(h, 0.37749318) * Math.pow(1 - p, -0.22722912);
    }
  }

  function cal_halflife_index(h) {
    return Math.max(Math.round(Math.log(h) / Math.log(base)) - min_index, 0);
  }

  // Convert string histories to arrays
  let interval_history_array = interval_history.split(',').map(Number);
  let feedback_history_array = feedback_history.split(',').map(Number);

  // Calculate best interval using the provided functions and histories
  // This is a simplified version of the original C++ code and may not work exactly the same way
  let best_interval = 0;
  for (let i = 0; i < iterations; i++) {
    for (let j = 0; j < interval_history_array.length; j++) {
      let interval = interval_history_array[j];
      let feedback = feedback_history_array[j];
      let h = cal_start_halflife(d);
      let p_recall = Math.exp2(-interval / h);
      let recall_h = cal_next_recall_halflife(h, p_recall, d, feedback);
      let h_index = cal_halflife_index(recall_h);
      if (h_index > best_interval) {
        best_interval = h_index;
      }
    }
  }

  return best_interval;
}