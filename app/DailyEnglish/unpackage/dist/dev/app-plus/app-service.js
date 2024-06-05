if (typeof Promise !== "undefined" && !Promise.prototype.finally) {
  Promise.prototype.finally = function(callback) {
    const promise = this.constructor;
    return this.then(
      (value) => promise.resolve(callback()).then(() => value),
      (reason) => promise.resolve(callback()).then(() => {
        throw reason;
      })
    );
  };
}
;
if (typeof uni !== "undefined" && uni && uni.requireGlobal) {
  const global2 = uni.requireGlobal();
  ArrayBuffer = global2.ArrayBuffer;
  Int8Array = global2.Int8Array;
  Uint8Array = global2.Uint8Array;
  Uint8ClampedArray = global2.Uint8ClampedArray;
  Int16Array = global2.Int16Array;
  Uint16Array = global2.Uint16Array;
  Int32Array = global2.Int32Array;
  Uint32Array = global2.Uint32Array;
  Float32Array = global2.Float32Array;
  Float64Array = global2.Float64Array;
  BigInt64Array = global2.BigInt64Array;
  BigUint64Array = global2.BigUint64Array;
}
;
if (uni.restoreGlobal) {
  uni.restoreGlobal(Vue, weex, plus, setTimeout, clearTimeout, setInterval, clearInterval);
}
(function(vue) {
  "use strict";
  function formatAppLog(type, filename, ...args) {
    if (uni.__log__) {
      uni.__log__(type, filename, ...args);
    } else {
      console[type].apply(console, [...args, filename]);
    }
  }
  function resolveEasycom(component, easycom) {
    return typeof component === "string" ? easycom : component;
  }
  const _export_sfc = (sfc, props) => {
    const target = sfc.__vccOpts || sfc;
    for (const [key, val] of props) {
      target[key] = val;
    }
    return target;
  };
  const _sfc_main$y = {
    data() {
      return {
        isHistoryVisible: false,
        //查询单词
        isDaka: false,
        //是否打卡
        isReview: false,
        //是否复习
        searchInput: "",
        daka_book: "",
        wordNumLearned: 123,
        wordNumTotal: 2345,
        daysLeft: 30,
        wordNumToPunch: 5,
        wordNumPunched: 15,
        wordNumToReview: 10,
        wordNumReviewed: 5,
        items: [{
          word: "apple",
          phonetic: "/ˈæpl/",
          meaning: "苹果111111111111111111111111111111111111111111111111"
        }, {
          word: "banana",
          phonetic: "/bəˈnɑː.nə/",
          meaning: "香蕉"
        }]
      };
    },
    methods: {
      handleDaka() {
        uni.navigateTo({
          url: "/pages/Examination/Examination?operation=0"
        });
      },
      handleReview() {
        uni.navigateTo({
          url: "/pages/Examination/Examination?operation=1"
        });
      },
      fetchData() {
        uni.request({
          url: "/api/punch/main_menu",
          method: "GET",
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
            if (res.statusCode === 200) {
              this.daka_book = res.data.task_doday.book_learning;
              this.wordNumLearned = res.data.task_doday.word_num_learned;
              this.wordNumTotal = res.data.task_doday.word_num_total;
              this.daysLeft = res.data.task_doday.days_left;
              this.wordNumToPunch = res.data.task_doday.word_num_to_punch;
              if (this.wordNumToPunch == 0) {
                this.isDaka = true;
              }
              this.wordNumPunched = res.data.task_doday.word_num_punched;
              this.wordNumToReview = res.data.task_doday.word_num_to_review;
              if (this.wordNumToReview == 0) {
                this.isReview = true;
              }
              this.wordNumReviewed = res.data.task_doday.word_num_reviewed;
            } else {
              formatAppLog("error", "at pages/home/home.vue:587", "请求失败", res);
              this.daka_book = "词汇书123";
            }
          },
          fail: (err) => {
            formatAppLog("error", "at pages/home/home.vue:592", "请求失败", err);
            this.daka_book = "词汇书123";
          }
        });
      },
      onLoad() {
        this.fetchData();
        formatAppLog("log", "at pages/home/home.vue:599", "hi");
      },
      handleSearchShow() {
        this.isHistoryVisible = true;
      },
      handleSearchRouter() {
        uni.navigateTo({
          // url: `/pages/word_details/word_details?word=${this.searchInput}`
          url: `/pages/word_details/word_details`
        });
        uni.showToast({
          title: "搜索成功",
          icon: "none"
        });
        formatAppLog("log", "at pages/home/home.vue:614", "本次搜索内容是", this.searchInput);
      },
      handleSearchInput(input) {
        this.searchInput = input;
      },
      cancelSearch() {
        this.isHistoryVisible = false;
        this.searchInput = "";
      }
    }
  };
  function _sfc_render$y(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", { class: "homepage" }, [
      vue.createElementVNode("view", { class: "search-container" }, [
        vue.createElementVNode("view", {
          class: "search-head",
          style: { "display": "flex" }
        }, [
          vue.createElementVNode(
            "view",
            {
              class: vue.normalizeClass(["search", { active: $data.isHistoryVisible }]),
              onClick: _cache[2] || (_cache[2] = ($event) => $options.handleSearchShow())
            },
            [
              vue.createElementVNode("image", {
                class: "search-icon",
                src: "/static/search.svg"
              }),
              vue.withDirectives(vue.createElementVNode(
                "input",
                {
                  placeholder: "搜索",
                  "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => $data.searchInput = $event),
                  onConfirm: _cache[1] || (_cache[1] = (...args) => $options.handleSearchRouter && $options.handleSearchRouter(...args))
                },
                null,
                544
                /* NEED_HYDRATION, NEED_PATCH */
              ), [
                [vue.vModelText, $data.searchInput]
              ])
            ],
            2
            /* CLASS */
          ),
          $data.isHistoryVisible ? (vue.openBlock(), vue.createElementBlock("button", {
            key: 0,
            class: "cancel",
            onClick: _cache[3] || (_cache[3] = (...args) => $options.cancelSearch && $options.cancelSearch(...args))
          }, "取消")) : (vue.openBlock(), vue.createElementBlock("image", {
            key: 1,
            class: "Msg-icon",
            src: "/static/email.png"
          }))
        ]),
        vue.withDirectives(vue.createElementVNode(
          "view",
          { class: "history" },
          [
            vue.createElementVNode("view", { class: "history-header" }, [
              vue.createElementVNode("text", { class: "title" }, "历史搜索"),
              vue.createElementVNode("text", { class: "clean" }, "清空")
            ]),
            vue.createElementVNode("view", { class: "list" }, [
              (vue.openBlock(true), vue.createElementBlock(
                vue.Fragment,
                null,
                vue.renderList($data.items, (item, index) => {
                  return vue.openBlock(), vue.createElementBlock("view", {
                    class: "item",
                    key: index,
                    onClick: ($event) => $options.handleSearchInput(item.word)
                  }, [
                    vue.createElementVNode("view", { class: "top-row" }, [
                      vue.createElementVNode(
                        "view",
                        { class: "word" },
                        vue.toDisplayString(item.word),
                        1
                        /* TEXT */
                      ),
                      vue.createElementVNode(
                        "view",
                        { class: "phonetic" },
                        vue.toDisplayString(item.phonetic),
                        1
                        /* TEXT */
                      )
                    ]),
                    vue.createElementVNode(
                      "view",
                      { class: "meaning" },
                      vue.toDisplayString(item.meaning),
                      1
                      /* TEXT */
                    )
                  ], 8, ["onClick"]);
                }),
                128
                /* KEYED_FRAGMENT */
              ))
            ])
          ],
          512
          /* NEED_PATCH */
        ), [
          [vue.vShow, $data.isHistoryVisible]
        ])
      ]),
      vue.withDirectives(vue.createElementVNode(
        "view",
        { class: "daka-container" },
        [
          vue.withDirectives(vue.createElementVNode(
            "image",
            {
              src: "/static/lihua.png",
              style: { "position": "absolute", "width": "330px", "height": "330px", "top": "140px", "left": "120px" }
            },
            null,
            512
            /* NEED_PATCH */
          ), [
            [vue.vShow, $data.isDaka]
          ]),
          vue.createElementVNode("view", { class: "daka-head" }, [
            vue.createElementVNode("view", { class: "column" }, [
              vue.createElementVNode("image", {
                class: "vocabook-img",
                src: "/static/book.png"
              })
            ]),
            vue.createElementVNode("view", { class: "column" }, [
              vue.createElementVNode("view", { class: "row" }, [
                vue.createElementVNode(
                  "view",
                  { class: "daka-title" },
                  vue.toDisplayString($data.daka_book),
                  1
                  /* TEXT */
                ),
                vue.createElementVNode("view", { class: "daka-subtitle" }, "修改")
              ]),
              vue.createElementVNode("view", { class: "row" }, [
                vue.createElementVNode("progress", {
                  percent: "10",
                  "active-color": "#10aeff",
                  backgroundColor: "#c8c8c8",
                  "stroke-width": "7"
                })
              ]),
              vue.createElementVNode("view", { class: "row" }, [
                vue.createElementVNode(
                  "view",
                  { class: "progress1" },
                  vue.toDisplayString($data.wordNumLearned) + "/" + vue.toDisplayString($data.wordNumTotal),
                  1
                  /* TEXT */
                ),
                vue.createElementVNode(
                  "view",
                  { class: "progress2" },
                  "剩余" + vue.toDisplayString($data.daysLeft) + "天",
                  1
                  /* TEXT */
                )
              ])
            ])
          ]),
          vue.createElementVNode("view", { class: "daka-line" }, [
            $data.isDaka ? (vue.openBlock(), vue.createElementBlock("view", {
              key: 0,
              class: "daka-title",
              style: { "font-size": "30rpx", "font-weight": "normal" }
            }, [
              vue.createTextVNode("恭喜你！"),
              vue.createElementVNode("br"),
              vue.createTextVNode("完成今日打卡 ")
            ])) : (vue.openBlock(), vue.createElementBlock("view", {
              key: 1,
              class: "daka-title"
            }, "今日计划")),
            vue.createElementVNode("view", { class: "daka-slogan" }, "随时随地，单词猛记")
          ]),
          vue.createElementVNode("view", { class: "daka-plan" }, [
            vue.createElementVNode("view", { class: "row" }, [
              $data.isDaka ? (vue.openBlock(), vue.createElementBlock("view", {
                key: 0,
                class: "plan-title1"
              }, "今日已新学")) : (vue.openBlock(), vue.createElementBlock("view", {
                key: 1,
                class: "plan-title1"
              }, "需新学")),
              $data.isReview ? (vue.openBlock(), vue.createElementBlock("view", {
                key: 2,
                class: "plan-title2"
              }, "今日已复习")) : (vue.openBlock(), vue.createElementBlock("view", {
                key: 3,
                class: "plan-title2"
              }, "需复习"))
            ]),
            vue.createElementVNode("view", { class: "row" }, [
              vue.createElementVNode("view", { class: "plan-num" }, [
                vue.createElementVNode(
                  "view",
                  { class: "number" },
                  vue.toDisplayString($data.isDaka ? $data.wordNumPunched : $data.wordNumToPunch),
                  1
                  /* TEXT */
                ),
                vue.createElementVNode("text", null, "词")
              ]),
              vue.createElementVNode("view", {
                class: "plan-num",
                style: { "margin-left": "100px" }
              }, [
                vue.createElementVNode(
                  "view",
                  { class: "number" },
                  vue.toDisplayString($data.isReview ? $data.wordNumReviewed : $data.wordNumToReview),
                  1
                  /* TEXT */
                ),
                vue.createElementVNode("text", null, "词")
              ])
            ]),
            vue.createElementVNode("view", { class: "row" }, [
              vue.withDirectives(vue.createElementVNode(
                "button",
                {
                  class: "plan-btn1",
                  onClick: _cache[4] || (_cache[4] = (...args) => $options.handleDaka && $options.handleDaka(...args))
                },
                "开始学习",
                512
                /* NEED_PATCH */
              ), [
                [vue.vShow, !$data.isDaka]
              ]),
              vue.withDirectives(vue.createElementVNode(
                "button",
                {
                  class: "plan-btn1",
                  onClick: _cache[5] || (_cache[5] = (...args) => $options.handleReview && $options.handleReview(...args))
                },
                "开始复习",
                512
                /* NEED_PATCH */
              ), [
                [vue.vShow, $data.isDaka && !$data.isReview]
              ]),
              vue.withDirectives(vue.createElementVNode(
                "button",
                {
                  class: "plan-btn1",
                  onClick: _cache[6] || (_cache[6] = (...args) => $options.handleDaka && $options.handleDaka(...args))
                },
                "继续学习",
                512
                /* NEED_PATCH */
              ), [
                [vue.vShow, $data.isDaka && $data.isReview]
              ])
            ])
          ])
        ],
        512
        /* NEED_PATCH */
      ), [
        [vue.vShow, !$data.isHistoryVisible]
      ]),
      vue.withDirectives(vue.createElementVNode(
        "view",
        { class: "content-container" },
        [
          vue.createElementVNode("view", { class: "button-list" }, [
            vue.createElementVNode("view", { class: "btn-item" }, [
              vue.createElementVNode("image", { src: "/static/word-exercise.png" }),
              vue.createElementVNode("text", null, "单词训练")
            ]),
            vue.createElementVNode("view", { class: "btn-item" }, [
              vue.createElementVNode("image", { src: "/static/biji.svg" }),
              vue.createElementVNode("text", null, "单词自检")
            ]),
            vue.createElementVNode("view", { class: "btn-item" }, [
              vue.createElementVNode("image", { src: "/static/write.svg" }),
              vue.createElementVNode("text", null, "爱写作")
            ]),
            vue.createElementVNode("view", { class: "btn-item" }, [
              vue.createElementVNode("image", { src: "/static/read.svg" }),
              vue.createElementVNode("text", null, "爱阅读")
            ])
          ])
        ],
        512
        /* NEED_PATCH */
      ), [
        [vue.vShow, !$data.isHistoryVisible]
      ]),
      vue.withDirectives(vue.createElementVNode(
        "view",
        { class: "ad-container" },
        [
          vue.createElementVNode("view", { class: "ad-list" }, [
            vue.createElementVNode("view", { class: "ad-item" }, [
              vue.createElementVNode("view", { class: "text-container" }, [
                vue.createElementVNode("text", { class: "title" }, "30分钟，拿下英语阅读"),
                vue.createElementVNode("text", { class: "content" }, "每日一读，提高英语阅读能力")
              ]),
              vue.createElementVNode("image", {
                class: "image",
                src: "/static/ad1.svg"
              })
            ]),
            vue.createElementVNode("view", { class: "ad-item" }, [
              vue.createElementVNode("view", { class: "text-container" }, [
                vue.createElementVNode("text", { class: "title" }, "30分钟，拿下英语阅读"),
                vue.createElementVNode("text", { class: "content" }, "每日一读，提高英语阅读能力")
              ]),
              vue.createElementVNode("image", {
                class: "image",
                src: "/static/ad1.svg"
              })
            ])
          ])
        ],
        512
        /* NEED_PATCH */
      ), [
        [vue.vShow, !$data.isHistoryVisible]
      ])
    ]);
  }
  const PagesHomeHome = /* @__PURE__ */ _export_sfc(_sfc_main$y, [["render", _sfc_render$y], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/home/home.vue"]]);
  const _sfc_main$x = {
    name: "CollapsibleView",
    props: {
      word: {
        type: Object,
        required: true
      }
    },
    data() {
      return {
        word: {},
        isExpanded: true
        // 初始状态为展开  
      };
    },
    methods: {
      toggleContent() {
        this.isExpanded = !this.isExpanded;
      }
    }
  };
  function _sfc_render$x(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock(
      "view",
      {
        class: vue.normalizeClass(["collapsible-view", { "collapsed": !$data.isExpanded }])
      },
      [
        vue.createElementVNode("header", {
          onClick: _cache[0] || (_cache[0] = (...args) => $options.toggleContent && $options.toggleContent(...args))
        }, [
          vue.renderSlot(_ctx.$slots, "header", {}, void 0, true)
        ]),
        vue.createElementVNode(
          "view",
          {
            class: "content",
            style: vue.normalizeStyle({ maxHeight: $data.isExpanded ? "100rem" : "0", transition: "max-height 0.3s ease" })
          },
          [
            vue.renderSlot(_ctx.$slots, "default", {}, void 0, true),
            vue.createCommentVNode(" 内容插槽 ")
          ],
          4
          /* STYLE */
        )
      ],
      2
      /* CLASS */
    );
  }
  const CollapsibleView = /* @__PURE__ */ _export_sfc(_sfc_main$x, [["render", _sfc_render$x], ["__scopeId", "data-v-4b4aa2b3"], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/components/CollapsibleView.vue"]]);
  const _sfc_main$w = {
    components: {
      CollapsibleView
    },
    data() {
      return {
        //词性简写转换表
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
        word: {
          name: "abandon",
          pronunciation: "/ əˈbændən /",
          meanings: ["v.    抛弃；放弃；沉湎于（某种情感）；舍弃，废弃", "n.    尽情，放纵"],
          book: "高中/ CET4 / CET6 / 考研 / IELTS / TOEFL / GRE"
        },
        //这是一个数组，元素是结构体，有word、details两个属性，分别代表单词和详细释义
        //其中details是一个数组，元素是结构体，有partOfSpeech、chineseMeaning、exampleSentence、sentenceMeaning四个属性，分别代表词性、中文释义、例句、英文释义
        wordDetail: {
          word: "abandon",
          details: [
            {
              partOfSpeech: "v.",
              chineseMeaning: "抛弃",
              exampleSentence: "Should I tell you to abandon me and save yourself, you must to do so. ",
              sentenceMeaning: "我若是让你别管我，救自己，你也必须照做。"
            },
            {
              partOfSpeech: "v.",
              chineseMeaning: "放弃",
              exampleSentence: "The girl has totally abandoned the use of computer for her homework.",
              sentenceMeaning: "这个女生彻底放弃使用电脑做作业了。"
            },
            {
              partOfSpeech: "v.",
              chineseMeaning: "沉湎于（某种情感）"
            },
            {
              partOfSpeech: "v.",
              chineseMeaning: "舍弃，废弃",
              exampleSentence: "The mining factory was abandoned a long time ago.",
              sentenceMeaning: "这个采矿工厂早已被放弃。"
            },
            {
              partOfSpeech: "n.",
              chineseMeaning: "尽情，放纵"
            }
          ]
        },
        //wordAndPhrases是一个数组，元素是结构体，有word、phraseAndMeanings两个属性，分别代表单词和短语释义
        //其中phraseAndMeanings是一个数组，元素是结构体，有phrase、meaning两个属性，分别代表短语和释义
        wordAndPhrase: {
          word: "abandon",
          phraseAndMeanings: [
            {
              phrase: "abandon oneself to",
              meaning: "沉溺于"
            },
            {
              phrase: "with abandon",
              meaning: "恣意地"
            },
            {
              phrase: "abandon doing sth.",
              meaning: "放弃做某事"
            },
            {
              phrase: "abandon ones belief",
              meaning: "放弃信仰"
            },
            {
              phrase: "abandon to",
              meaning: "离弃，遗弃，抛弃"
            }
          ]
        }
      };
    },
    onLoad(event) {
      let word = event["word"];
      let word_id = event["word_id"];
      this.word.name = word;
      this.wordDetail.word = word;
      this.wordAndPhrase.word = word;
      let localDetails = uni.getStorageSync(word_id);
      if (localDetails) {
        this.word.pronunciation = localDetails.pronunciation;
        this.word.meanings = this.transformMeaningsToText(localDetails.meanings);
      }
      uni.request({
        url: "/api/words/word_details",
        method: "POST",
        header: {
          "content-type": "application/json",
          // 默认值
          "Authorization": "Bearer " + uni.getStorageSync("token")
          // 登录后获取的token
        },
        data: {
          word_id
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
          let detailedMeanings = res.data.detailed_meanings;
          this.wordDetail.details = this.transformDetailedMeaningsToDetails(detailedMeanings);
          this.wordAndPhrase.phraseAndMeanings = res.data.phrases;
          this.word.book = res.data.word_book;
        },
        fail: (res) => {
          formatAppLog("log", "at pages/word_details/word_details.vue:236", res);
        }
      });
    },
    methods: {
      transformMeaningsToText(meanings) {
        const transformedMeanings = [];
        for (const speech in meanings) {
          if (meanings[speech] && meanings[speech].length > 0) {
            const speechAbbreviation = this.simplifiedSpeech[speech];
            const meaningText = `${speechAbbreviation}		${meanings[speech].join("; ")}
`;
            transformedMeanings.push(meaningText);
          }
        }
        return transformedMeanings;
      },
      transformDetailedMeaningsToDetails(detailedMeanings) {
        const details = [];
        for (const speech in detailedMeanings) {
          if (detailedMeanings[speech] && detailedMeanings[speech].length > 0) {
            const partOfSpeech = this.simplifiedSpeech[speech];
            detailedMeanings[speech].forEach((meaning) => {
              const detail = {
                partOfSpeech,
                chineseMeaning: meaning.chinese_meaning,
                exampleSentence: meaning.example_sentence,
                sentenceMeaning: meaning.sentence_meaning
              };
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
  };
  function _sfc_render$w(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_CollapsibleView = vue.resolveComponent("CollapsibleView");
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("view", { class: "word-list-container" }, [
        vue.createElementVNode("view", { class: "word-item" }, [
          vue.createElementVNode("view", { class: "word-card" }, [
            vue.createElementVNode(
              "h1",
              { class: "word" },
              vue.toDisplayString($data.word.name),
              1
              /* TEXT */
            ),
            vue.createElementVNode("image", {
              src: "/static/collection.png",
              class: "collection"
            })
          ]),
          vue.createElementVNode("view", { class: "pronunciation-card" }, [
            vue.createElementVNode(
              "p",
              null,
              vue.toDisplayString($data.word.pronunciation),
              1
              /* TEXT */
            ),
            vue.createElementVNode("image", {
              src: "/static/pronunciation.png",
              class: "pronunciation"
            })
          ]),
          (vue.openBlock(true), vue.createElementBlock(
            vue.Fragment,
            null,
            vue.renderList($data.word.meanings, (meaning) => {
              return vue.openBlock(), vue.createElementBlock(
                "p",
                null,
                vue.toDisplayString(meaning),
                1
                /* TEXT */
              );
            }),
            256
            /* UNKEYED_FRAGMENT */
          )),
          vue.createElementVNode(
            "p",
            { class: "book" },
            vue.toDisplayString($data.word.book),
            1
            /* TEXT */
          )
        ])
      ]),
      vue.createVNode(
        _component_CollapsibleView,
        { ref: "collapsibleView1" },
        {
          header: vue.withCtx(() => [
            vue.createElementVNode("view", { class: "header" }, [
              vue.createElementVNode("h3", null, "单词变形"),
              vue.createElementVNode("image", {
                onClick: _cache[0] || (_cache[0] = (...args) => $options.toggleCollapse1 && $options.toggleCollapse1(...args)),
                src: "/static/up.png",
                class: "up-arrow"
              })
            ])
          ]),
          default: vue.withCtx(() => [
            vue.createElementVNode("view", { class: "content" }, "这里是第一个可展开/缩起的内容")
          ]),
          _: 1
          /* STABLE */
        },
        512
        /* NEED_PATCH */
      ),
      vue.createVNode(
        _component_CollapsibleView,
        { ref: "collapsibleView2" },
        {
          header: vue.withCtx(() => [
            vue.createElementVNode("view", { class: "header" }, [
              vue.createElementVNode("h3", null, "详细释义"),
              vue.createElementVNode("image", {
                onClick: _cache[1] || (_cache[1] = (...args) => $options.toggleCollapse2 && $options.toggleCollapse2(...args)),
                src: "/static/up.png",
                class: "up-arrow"
              })
            ])
          ]),
          default: vue.withCtx(() => [
            vue.createElementVNode("view", null, [
              (vue.openBlock(true), vue.createElementBlock(
                vue.Fragment,
                null,
                vue.renderList($data.wordDetail.details, (detail) => {
                  return vue.openBlock(), vue.createElementBlock("view", {
                    key: detail.partOfSpeech
                  }, [
                    vue.createElementVNode("view", { style: { "display": "flex", "justify-content": "flex-start" } }, [
                      vue.createElementVNode(
                        "p",
                        { class: "partofspeech" },
                        vue.toDisplayString(detail.partOfSpeech),
                        1
                        /* TEXT */
                      ),
                      vue.createElementVNode(
                        "p",
                        null,
                        vue.toDisplayString(detail.chineseMeaning),
                        1
                        /* TEXT */
                      )
                    ]),
                    vue.createElementVNode("view", { style: { "display": "flex", "flex-direction": "row", "align-items": "center" } }, [
                      vue.createElementVNode(
                        "p",
                        { class: "example-text" },
                        vue.toDisplayString(detail.exampleSentence),
                        1
                        /* TEXT */
                      ),
                      vue.createElementVNode("image", {
                        class: "example-image",
                        src: "/static/pronunciation.png",
                        alt: "Example Image"
                      })
                    ]),
                    vue.createElementVNode(
                      "p",
                      { class: "meaning" },
                      vue.toDisplayString(detail.sentenceMeaning),
                      1
                      /* TEXT */
                    )
                  ]);
                }),
                128
                /* KEYED_FRAGMENT */
              ))
            ])
          ]),
          _: 1
          /* STABLE */
        },
        512
        /* NEED_PATCH */
      ),
      vue.createVNode(
        _component_CollapsibleView,
        { ref: "collapsibleView3" },
        {
          header: vue.withCtx(() => [
            vue.createElementVNode("view", { class: "header" }, [
              vue.createElementVNode("h3", null, "短语"),
              vue.createElementVNode("image", {
                onClick: _cache[2] || (_cache[2] = (...args) => $options.toggleCollapse3 && $options.toggleCollapse3(...args)),
                src: "/static/up.png",
                class: "up-arrow"
              })
            ])
          ]),
          default: vue.withCtx(() => [
            vue.createElementVNode("view", { style: { "text-align": "left" } }, [
              (vue.openBlock(true), vue.createElementBlock(
                vue.Fragment,
                null,
                vue.renderList($data.wordAndPhrase.phraseAndMeanings, (phraseAndMeaning) => {
                  return vue.openBlock(), vue.createElementBlock("view", {
                    key: phraseAndMeaning.phrase
                  }, [
                    vue.createElementVNode(
                      "p",
                      { class: "phrase" },
                      vue.toDisplayString(phraseAndMeaning.phrase),
                      1
                      /* TEXT */
                    ),
                    vue.createElementVNode(
                      "p",
                      { class: "phrase-meaning" },
                      vue.toDisplayString(phraseAndMeaning.meaning),
                      1
                      /* TEXT */
                    )
                  ]);
                }),
                128
                /* KEYED_FRAGMENT */
              ))
            ])
          ]),
          _: 1
          /* STABLE */
        },
        512
        /* NEED_PATCH */
      )
    ]);
  }
  const PagesWord_detailsWord_details = /* @__PURE__ */ _export_sfc(_sfc_main$w, [["render", _sfc_render$w], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/word_details/word_details.vue"]]);
  const _sfc_main$v = {
    data() {
      return {
        operation: 0,
        //0打卡，1复习
        swiperOptions: {
          // 其他配置...
          allowTouchMove: true,
          // 允许触摸滑动
          preventClicksPropagation: true
          // 阻止点击事件冒泡
          // 其他 Swiper 配置...
        },
        progress: 1,
        // 进度条的初始值
        current: 1,
        // 当前进度
        currentQuestionIndex: 0,
        //所有题目正确情况
        isCorrects: {
          1: false,
          2: false,
          3: false
        },
        questions: [
          // 题目和选项
          {
            word_id: 1,
            word: "abandon",
            phonetic: "[ə'bændən]",
            choices: ["1", "2", "2", "放弃"]
          },
          {
            word_id: 2,
            word: "abandon",
            phonetic: "[ə'bændən]",
            choices: ["1", "选项B", "选项C", "选项D"]
          },
          {
            word_id: 3,
            word: "abandon2",
            phonetic: "[ə'bændən]",
            choices: ["1", "选项B", "选项C", "选项D"]
          }
          // ...更多题目
        ],
        // 这里可以根据需要修改选项内容
        selectedChoice: "",
        // 用于存储用户选择的答案
        realAnswer: [
          "放弃",
          "选项B",
          "选项C"
          // 正确答案
        ]
      };
    },
    onLoad(event) {
      let operation = parseInt(event["operation"]);
      this.operation = operation;
      this.isCorrects = {};
      uni.request({
        //判断操作类型并发送请求
        url: !operation ? "/api/main/take_punch" : "/api/main/take_review",
        method: "GET",
        header: {
          "Authorization": `Bearer ${uni.getStorageSync("token")}`
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
          formatAppLog("log", "at pages/Examination/Examination.vue:114", res);
          if (res.data.code == 200) {
            let word_list = res.data.word_list;
            word_list.forEach((word, index) => {
              let question = {
                word_id: word.word_id,
                word: word.word,
                phonetic: word.phonetic_us,
                choices: Object.values(word.word_question)
              };
              this.isCorrects[word.word_id] = false;
              let realAnswer = word.word_question[word.answer];
              this.questions.push(question);
              this.realAnswer.push(realAnswer);
            });
          }
        },
        fail: (err) => {
          formatAppLog("log", "at pages/Examination/Examination.vue:143", err);
        }
      });
    },
    methods: {
      handleBack() {
        this.$router.back();
      },
      handleJump() {
        uni.switchTab({
          url: "../Vocab/Vocab"
        });
        uni.request({
          url: "xxvcav",
          method: "post",
          data: {
            //data
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
          }
        });
      },
      swiperChange(event) {
        const current = event.detail.current;
        const source = event.detail.source;
        if (source === "touch") {
          if (current > this.currentQuestionIndex) {
            this.currentQuestionIndex = current;
          } else if (current < this.currentQuestionIndex) {
            this.$refs.swiper.scrollTo(this.currentQuestionIndex, 0, false);
          }
        }
      },
      updateProgressBar() {
        this.updateProgress(this.progress + 1);
      },
      updateProgress(value) {
        if (value >= 0 && value <= 100) {
          this.progress = value;
          this.current = value;
        } else {
          formatAppLog("error", "at pages/Examination/Examination.vue:199", "进度值必须在 0 到 100 之间");
        }
      },
      selectChoice(index) {
        formatAppLog("log", "at pages/Examination/Examination.vue:203", index);
        let selectedChoice = this.questions[this.currentQuestionIndex].choices[index];
        let word_id = this.questions[this.currentQuestionIndex].word_id;
        formatAppLog("log", "at pages/Examination/Examination.vue:207", word_id);
        let isCorrect = selectedChoice === this.realAnswer[this.currentQuestionIndex];
        formatAppLog("log", "at pages/Examination/Examination.vue:210", isCorrect);
        this.isCorrects[word_id] = isCorrect;
        if (isCorrect) {
          let nextIndex = this.currentQuestionIndex;
          this.updateProgressBar();
          this.$nextTick(() => {
            this.showCorrectAnswer(this.realAnswer[nextIndex], nextIndex);
          });
          if (++this.currentQuestionIndex == this.questions.length) {
            uni.request({
              //判断操作类型并发送请求
              url: !this.operation ? "/api/main/punched" : "/api/main/reviewed",
              method: "POST",
              header: {
                "Authorization": `Bearer ${uni.getStorageSync("token")}`
              },
              data: {
                punch_result: this.isCorrects
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
                formatAppLog("log", "at pages/Examination/Examination.vue:237", res);
                if (res.data.code == 200) {
                  uni.showToast({
                    title: this.operation ? "复习结束" : "打卡结束",
                    icon: "none",
                    duration: 2e3,
                    success: () => {
                      uni.switchTab({
                        url: "../home/home"
                      });
                    }
                  });
                }
              },
              fail: (err) => {
                formatAppLog("log", "at pages/Examination/Examination.vue:252", err);
              }
            });
          }
        } else {
          let currIndex = this.currentQuestionIndex;
          this.$nextTick(() => {
            this.showIncorrectAnswer(index);
            this.showCorrectAnswer(this.realAnswer[currIndex]);
          });
        }
      },
      showCorrectAnswer(answer, index) {
        const correctIndex = this.questions[index].choices.indexOf(answer);
        const correctButton = this.$refs[`option${correctIndex}`];
        if (correctButton) {
          correctButton.classList.add("correct");
        }
      },
      showIncorrectAnswer(index) {
        const incorrectButton = this.$refs[`option${index}`];
        if (incorrectButton) {
          incorrectButton.classList.add("incorrect");
        }
      },
      preventSelect(event) {
        event.preventDefault();
      },
      getClass(index) {
        if (this.selectedChoice) {
          formatAppLog("log", "at pages/Examination/Examination.vue:291", this.currentQuestionIndex);
          if (this.questions[this.currentQuestionIndex].choices[index] === this.selectedChoice) {
            return this.questions[this.currentQuestionIndex].choices[index] === this.realAnswer ? "correct" : "incorrect";
          }
        }
        return "";
      }
    }
  };
  function _sfc_render$v(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", { class: "container" }, [
      vue.createElementVNode(
        "text",
        { class: "progress-text" },
        vue.toDisplayString($data.current) + "/" + vue.toDisplayString($data.questions.length),
        1
        /* TEXT */
      ),
      vue.createElementVNode("view", { class: "progress-container" }, [
        vue.createElementVNode(
          "view",
          {
            class: "progress-bar",
            style: vue.normalizeStyle({ width: $data.progress + "%" })
          },
          null,
          4
          /* STYLE */
        )
      ]),
      vue.createElementVNode("image", {
        class: "back-icon",
        src: "/static/back.svg",
        onClick: _cache[0] || (_cache[0] = (...args) => $options.handleBack && $options.handleBack(...args))
      }),
      vue.createElementVNode("swiper", {
        class: "question-container",
        options: $data.swiperOptions,
        "easing-function": "linear",
        duration: 250,
        onBeforeChange: _cache[3] || (_cache[3] = (...args) => $options.swiperChange && $options.swiperChange(...args))
      }, [
        (vue.openBlock(true), vue.createElementBlock(
          vue.Fragment,
          null,
          vue.renderList($data.questions, (question, index) => {
            return vue.openBlock(), vue.createElementBlock("swiper-item", { key: index }, [
              vue.createElementVNode("view", { class: "text-info" }, [
                vue.createElementVNode(
                  "text",
                  { class: "word" },
                  vue.toDisplayString(question.word),
                  1
                  /* TEXT */
                ),
                vue.createElementVNode(
                  "text",
                  { class: "phonetic" },
                  vue.toDisplayString(question.phonetic),
                  1
                  /* TEXT */
                )
              ]),
              vue.createElementVNode("view", { class: "button-group" }, [
                (vue.openBlock(true), vue.createElementBlock(
                  vue.Fragment,
                  null,
                  vue.renderList(question.choices, (choice, choiceIndex) => {
                    return vue.openBlock(), vue.createElementBlock("button", {
                      class: vue.normalizeClass(["option", $options.getClass(choiceIndex)]),
                      key: choiceIndex,
                      onClick: ($event) => $options.selectChoice(choiceIndex)
                    }, vue.toDisplayString(choice), 11, ["onClick"]);
                  }),
                  128
                  /* KEYED_FRAGMENT */
                ))
              ]),
              vue.createElementVNode("view", {
                class: "jump-group",
                onClick: _cache[1] || (_cache[1] = (...args) => $options.handleJump && $options.handleJump(...args))
              }, [
                vue.createElementVNode("text", { class: "link" }, "加入生词本"),
                vue.createElementVNode("image", {
                  class: "jump-icon",
                  src: "/static/jump.svg"
                })
              ]),
              vue.createElementVNode("view", {
                class: "jump-group2",
                onClick: _cache[2] || (_cache[2] = (...args) => _ctx.handleJump2 && _ctx.handleJump2(...args))
              }, [
                vue.createElementVNode("text", { class: "link" }, "不认识，下一个")
              ])
            ]);
          }),
          128
          /* KEYED_FRAGMENT */
        )),
        vue.createCommentVNode(' <view class="text-info">\r\n				<text class="word">{{ questions[currentQuestionIndex].word }}</text>\r\n				<text class="phonetic">{{questions[currentQuestionIndex].phonetic}}</text>\r\n			</view>\r\n\r\n			<view class="button-group">\r\n				<button class="option" v-for="(choice, index) in \r\n				questions[currentQuestionIndex].choices" :key="index" :class="getClass(index)"\r\n					@click="selectChoice(index)">{{ choice }}</button>\r\n			</view> ')
      ], 40, ["options"])
    ]);
  }
  const PagesExaminationExamination = /* @__PURE__ */ _export_sfc(_sfc_main$v, [["render", _sfc_render$v], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/Examination/Examination.vue"]]);
  const _sfc_main$u = {};
  function _sfc_render$u(_ctx, _cache) {
    return vue.openBlock(), vue.createElementBlock("view", { class: "tab-bar" });
  }
  const TabBar = /* @__PURE__ */ _export_sfc(_sfc_main$u, [["render", _sfc_render$u], ["__scopeId", "data-v-89ca1f91"], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/components/TabBar.vue"]]);
  const _sfc_main$t = {
    components: {
      TabBar
    },
    data() {
      return {
        title: "Hello"
      };
    },
    onLoad() {
    },
    methods: {}
  };
  function _sfc_render$t(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_TabBar = vue.resolveComponent("TabBar");
    return vue.openBlock(), vue.createElementBlock("view", { class: "content" }, [
      vue.createVNode(_component_TabBar),
      vue.createElementVNode("view", { class: "text-area" }, [
        vue.createElementVNode(
          "text",
          { class: "title" },
          vue.toDisplayString($data.title),
          1
          /* TEXT */
        )
      ])
    ]);
  }
  const PagesIndexIndex = /* @__PURE__ */ _export_sfc(_sfc_main$t, [["render", _sfc_render$t], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/index/index.vue"]]);
  const _sfc_main$s = {};
  function _sfc_render$s(_ctx, _cache) {
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("text", { class: "hello-world" }, "Hello World!")
    ]);
  }
  const PagesIndexHelloWorld = /* @__PURE__ */ _export_sfc(_sfc_main$s, [["render", _sfc_render$s], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/index/HelloWorld.vue"]]);
  const _sfc_main$r = {
    data() {
      return {
        punch_word_num: 120,
        total_punch_day: 12,
        consecutive_punch_day: 7
      };
    },
    methods: {
      goToTeam() {
        uni.navigateTo({
          url: "../MyTeam/MyTeam"
          //Todo: 跳转到团队页面
        });
      },
      goToCalendar() {
        uni.navigateTo({
          url: "../Calendar/Calendar"
        });
      },
      GotoPersonal_information() {
        uni.navigateTo({
          url: "../personal-information/personal-information"
        });
      }
    }
  };
  function _sfc_render$r(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("view", {
        class: "personal-information",
        onClick: _cache[0] || (_cache[0] = (...args) => $options.GotoPersonal_information && $options.GotoPersonal_information(...args))
      }, [
        vue.createElementVNode("image", {
          class: "head",
          src: "/static/pikachu.jpg"
        }),
        vue.createElementVNode("view", { style: { "display": "flex", "flex-direction": "column" } }, [
          vue.createElementVNode("span", { class: "username" }, "user")
        ]),
        vue.createElementVNode("image", {
          class: "right1",
          src: "/static/back.svg"
        })
      ]),
      vue.createElementVNode("view", { class: "container1" }, [
        vue.createElementVNode("view", { class: "words-amount" }, [
          vue.createElementVNode(
            "span",
            { class: "number1" },
            vue.toDisplayString($data.punch_word_num),
            1
            /* TEXT */
          ),
          vue.createElementVNode("span", { class: "words1" }, "打卡单词数")
        ]),
        vue.createElementVNode("view", { class: "lianxudakatianshu" }, [
          vue.createElementVNode(
            "span",
            { class: "number3" },
            vue.toDisplayString($data.consecutive_punch_day),
            1
            /* TEXT */
          ),
          vue.createElementVNode("span", { class: "words3" }, "连续打卡天数")
        ]),
        vue.createElementVNode("view", { class: "dakatianshu" }, [
          vue.createElementVNode(
            "span",
            { class: "number2" },
            vue.toDisplayString($data.total_punch_day),
            1
            /* TEXT */
          ),
          vue.createElementVNode("span", { class: "words2" }, "总打卡天数")
        ])
      ]),
      vue.createElementVNode("view", {
        class: "container2",
        onClick: _cache[1] || (_cache[1] = (...args) => $options.goToCalendar && $options.goToCalendar(...args))
      }, [
        vue.createElementVNode("image", {
          class: "calendar",
          src: "/static/calendar.png"
        }),
        vue.createElementVNode("span", { class: "word2" }, "我的日历"),
        vue.createElementVNode("image", {
          class: "right",
          src: "/static/back.svg"
        })
      ]),
      vue.createElementVNode("view", { class: "container3" }, [
        vue.createElementVNode("image", {
          class: "wordbook",
          src: "/static/biji2.svg"
        }),
        vue.createElementVNode("span", { class: "word3" }, "我的单词本"),
        vue.createElementVNode("image", {
          class: "right",
          src: "/static/back.svg"
        })
      ]),
      vue.createElementVNode("view", {
        class: "container3",
        onClick: _cache[2] || (_cache[2] = (...args) => $options.goToTeam && $options.goToTeam(...args))
      }, [
        vue.createElementVNode("image", {
          class: "team",
          src: "/static/team.svg"
        }),
        vue.createElementVNode("span", { class: "word31" }, "我的团队"),
        vue.createElementVNode("image", {
          class: "right",
          src: "/static/back.svg"
        })
      ]),
      vue.createElementVNode("view", { class: "container4" }, [
        vue.createElementVNode("image", {
          class: "setting",
          src: "/static/setting.png"
        }),
        vue.createElementVNode("span", { class: "word4" }, "设置"),
        vue.createElementVNode("image", {
          class: "right",
          src: "/static/back.svg"
        })
      ]),
      vue.createElementVNode("view", { class: "container5" }, [
        vue.createElementVNode("image", {
          class: "feedback",
          src: "/static/feedback.png"
        }),
        vue.createElementVNode("span", { class: "word5" }, "反馈"),
        vue.createElementVNode("image", {
          class: "right",
          src: "/static/back.svg"
        })
      ]),
      vue.createElementVNode("view", { class: "container5" }, [
        vue.createElementVNode("image", {
          class: "feedback",
          src: "/static/like.svg"
        }),
        vue.createElementVNode("span", { class: "word5" }, "给我们好评"),
        vue.createElementVNode("image", {
          class: "right",
          src: "/static/back.svg"
        })
      ])
    ]);
  }
  const PagesUserUser = /* @__PURE__ */ _export_sfc(_sfc_main$r, [["render", _sfc_render$r], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/user/user.vue"]]);
  const _sfc_main$q = {
    props: {
      word: String,
      pronunciation: String,
      meaning: String,
      reviewCount: Number,
      details: Object,
      id: Number,
      sound: String
    },
    // 组件的其他逻辑
    data() {
      return {
        word: this.word,
        pronunciation: this.pronunciation,
        meaning: this.meaning,
        showPronunciationMeaning: false,
        reviewCount: this.reviewCount,
        details: this.details,
        id: this.id,
        sound: this.sound
      };
    },
    methods: {
      togglePronunciationAndMeaning() {
        this.showPronunciationMeaning = !this.showPronunciationMeaning;
      },
      navigation() {
        let localDetails = uni.getStorageSync(this.id);
        if (!localDetails) {
          uni.setStorageSync(this.id, this.details);
        }
        uni.navigateTo({
          url: "/pages/word_details/word_details?word_id=" + this.id + "&word=" + this.word
        });
      },
      readWord() {
        const innerAudioContext = uni.createInnerAudioContext();
        innerAudioContext.src = this.sound;
        innerAudioContext.play();
      }
    }
  };
  function _sfc_render$q(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", { class: "word-card" }, [
      vue.createElementVNode("view", { class: "word-container" }, [
        vue.createElementVNode(
          "view",
          {
            class: "word",
            onClick: _cache[0] || (_cache[0] = (...args) => $options.navigation && $options.navigation(...args))
          },
          vue.toDisplayString($data.word),
          1
          /* TEXT */
        ),
        vue.createElementVNode(
          "view",
          {
            class: "pronunciation-meaning",
            onClick: _cache[1] || (_cache[1] = (...args) => $options.togglePronunciationAndMeaning && $options.togglePronunciationAndMeaning(...args))
          },
          vue.toDisplayString($data.showPronunciationMeaning ? `${$data.pronunciation} - ${$data.meaning}` : "点击显示音标和释义"),
          1
          /* TEXT */
        )
      ]),
      vue.createElementVNode("view", { class: "action-container" }, [
        vue.createElementVNode("image", {
          class: "read-button",
          src: "/static/read-icon.svg",
          onClick: _cache[2] || (_cache[2] = (...args) => $options.readWord && $options.readWord(...args))
        }),
        vue.createElementVNode("view", { class: "review-count" }, [
          vue.createElementVNode("p", null, "复习次数:"),
          vue.createElementVNode(
            "p",
            { style: { "color": "#e74c3c" } },
            vue.toDisplayString($data.reviewCount),
            1
            /* TEXT */
          )
        ])
      ])
    ]);
  }
  const PagesWordBlockWordBlock = /* @__PURE__ */ _export_sfc(_sfc_main$q, [["render", _sfc_render$q], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/WordBlock/WordBlock.vue"]]);
  const _sfc_main$p = {
    name: "backTop",
    methods: {
      backToTop() {
        uni.pageScrollTo({
          scrollTop: 0,
          duration: 300
          //滚动动画的时长
        });
      }
    }
  };
  function _sfc_render$p(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", {
      class: "backTop",
      onClick: _cache[0] || (_cache[0] = (...args) => $options.backToTop && $options.backToTop(...args))
    }, [
      vue.createElementVNode("image", { src: "/static/backTop.svg" })
    ]);
  }
  const backTop = /* @__PURE__ */ _export_sfc(_sfc_main$p, [["render", _sfc_render$p], ["__scopeId", "data-v-de36d26c"], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/components/backTop.vue"]]);
  const _sfc_main$o = {
    components: {
      WordBlock: PagesWordBlockWordBlock,
      backTop
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
        words: [
          {
            word_id: 1,
            spelling: "moral",
            pronunciation: "/ˈmɔːrəl/",
            meanings: {
              verb: null,
              adjective: ["道德的", "品行端正的", "伦理的", " 精神上的"],
              noun: ["道德教训", "寓意", "品德", "品行"],
              pronoun: null,
              adverb: null,
              conjunction: null,
              preposition: null,
              interjection: null
            },
            sound: "https://ssl.gstatic.com/dictionary/static/sounds/oxford/moral--_gb_1.mp3"
          },
          {
            word_id: 2,
            spelling: "abandon",
            pronunciation: "/əˈbændən/",
            meanings: {
              verb: ["抛弃", "放弃", "弃置", "放弃治疗"],
              noun: ["放弃物", "放弃的事物", "放弃的念头", "放弃的决定"],
              pronoun: null,
              adverb: null,
              conjunction: null,
              preposition: null,
              interjection: null
            },
            sound: "https://ssl.gstatic.com/dictionary/static/sounds/oxford/abandon--_gb_1.mp3"
          }
        ],
        // 单词列表
        cnt: 0,
        book: "cet4",
        startIndex: 0,
        // 开始加载的索引
        endIndex: 20,
        // 结束加载的索引
        showBackTop: false
        //是否显示返回顶部按钮
      };
    },
    /*  onLoad() {
        uni.request({
          url: "/api/words/get_starbk",
          method: "POST",
          header: {
            'Authorization': 'Bearer ' + uni.getStorageSync('token')
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
            this.words = res.data.words;
          },
          fail: (res) => {
            __f__('log','at pages/Vocab/Vocab.vue:98',"请求失败");
          }
        });
      },*/
    onPageScroll(e) {
      let scrollTop = e.scrollTop;
      if (scrollTop > 100) {
        this.showBackTop = true;
      } else {
        this.showBackTop = false;
      }
    },
    methods: {
      getMeaningStr(meanings) {
        let meaningStr = "";
        let foundFirst = false;
        for (let key in meanings) {
          if (meanings[key] && meanings[key].length > 0) {
            if (!foundFirst) {
              meaningStr += this.simplifiedSpeech[key];
              foundFirst = true;
            } else {
              meaningStr += "、";
            }
            meaningStr += meanings[key].slice(0, 2).join("、");
            if (meanings[key].length > 2) {
              meaningStr += "...";
              break;
            } else if (meanings[key].length === 2) {
              meaningStr += "\n";
              break;
            }
          }
        }
        return meaningStr;
      },
      handleBack() {
        uni.navigateBack();
      },
      Review() {
      },
      Export() {
      },
      handleTouchEnd(e) {
        const scrollTop = e.target.scrollTop;
        const scrollHeight = e.target.scrollHeight;
        const clientHeight = e.target.clientHeight;
        if (scrollTop + clientHeight >= scrollHeight) {
          formatAppLog("log", "at pages/Vocab/Vocab.vue:165", "滑动到底部");
        } else {
          wx.showToast({
            title: "Hello, World!",
            icon: "none"
            // 设置为'none'可以避免出现加载图标
          });
        }
      }
    }
  };
  function _sfc_render$o(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_word_block = vue.resolveComponent("word-block");
    const _component_backTop = vue.resolveComponent("backTop");
    return vue.openBlock(), vue.createElementBlock("view", { class: "container" }, [
      vue.createElementVNode("image", {
        class: "back-icon",
        src: "/static/back.svg",
        onClick: _cache[0] || (_cache[0] = (...args) => $options.handleBack && $options.handleBack(...args))
      }),
      vue.createElementVNode("view", { class: "vocabook" }, [
        vue.createElementVNode("image", {
          class: "vocabook-img",
          src: "/static/book.png"
        }),
        vue.createElementVNode(
          "view",
          { class: "vocabook-title" },
          "单词书:" + vue.toDisplayString($data.book),
          1
          /* TEXT */
        ),
        vue.createElementVNode(
          "view",
          { class: "vocabook-cnt" },
          "生词数：" + vue.toDisplayString($data.cnt),
          1
          /* TEXT */
        ),
        vue.createElementVNode("view", { class: "button-container" }, [
          vue.createElementVNode("button", {
            class: "review",
            onClick: _cache[1] || (_cache[1] = (...args) => $options.Review && $options.Review(...args))
          }, "复习")
        ]),
        vue.createElementVNode("view", { class: "button-container" }, [
          vue.createElementVNode("button", {
            class: "export",
            onClick: _cache[2] || (_cache[2] = (...args) => $options.Export && $options.Export(...args))
          }, "导出")
        ])
      ]),
      vue.createElementVNode(
        "view",
        {
          class: "word-blocks",
          onTouchend: _cache[3] || (_cache[3] = ($event) => $options.handleTouchEnd())
        },
        [
          (vue.openBlock(true), vue.createElementBlock(
            vue.Fragment,
            null,
            vue.renderList($data.words, (word) => {
              return vue.openBlock(), vue.createBlock(_component_word_block, {
                key: word.id,
                word: word.spelling,
                id: word.word_id,
                pronunciation: word.pronunciation,
                meaning: $options.getMeaningStr(word.meanings),
                details: word,
                "review-count": 5,
                sound: word.sound
              }, null, 8, ["word", "id", "pronunciation", "meaning", "details", "sound"]);
            }),
            128
            /* KEYED_FRAGMENT */
          ))
        ],
        32
        /* NEED_HYDRATION */
      ),
      vue.createElementVNode("view", null, [
        $data.showBackTop ? (vue.openBlock(), vue.createBlock(_component_backTop, { key: 0 })) : vue.createCommentVNode("v-if", true)
      ])
    ]);
  }
  const PagesVocabVocab = /* @__PURE__ */ _export_sfc(_sfc_main$o, [["render", _sfc_render$o], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/Vocab/Vocab.vue"]]);
  const _sfc_main$n = {
    data() {
      return {
        title: "",
        //划过之后的标题
        width: 0,
        //滑到多宽
        reactWidth: 0,
        //整个矩形的宽度
        sliderWidth: 0,
        //滑块宽度
        startX: 0,
        //开始触摸距离屏幕左面的位置
        sendFlag: false,
        //是否发送
        finishFlag: false,
        //是否允许滑动 判断是否滑动完成
        moveFlag: false
        //是否执行滑动函数
        //   isLongPress: false,//是否长按
        //   longPressTimeout: null,//长按定时器
        // longPressThreshold: 500, // 长按阈值，单位为毫秒
      };
    },
    mounted() {
      let selectFc = uni.createSelectorQuery().in(this);
      selectFc.select("#react").boundingClientRect((data2) => {
        this.reactWidth = data2.width - 2;
      }).exec();
      selectFc.select("#slider").boundingClientRect((data2) => {
        this.sliderWidth = data2.width;
      }).exec();
    },
    methods: {
      start(e) {
        let {
          clientX,
          clientY
        } = e.touches[0];
        this.startX = clientX;
        this.moveFlag = true;
      },
      reset() {
        this.sendFlag = false;
        this.finishFlag = false;
        this.width = 0;
        this.title = "";
      },
      finish() {
        this.finishFlag = true;
        this.title = "已提交";
      },
      move(e) {
        if (!this.moveFlag)
          return;
        if (this.width >= this.reactWidth - this.sliderWidth) {
          if (!this.sendFlag) {
            this.moveFlag = false;
            this.sendFlag = true;
            this.$emit("change", {
              finish: this.finish.bind(this),
              reset: this.reset.bind(this)
            });
          }
        } else {
          let {
            clientX,
            clientY
          } = e.touches[0];
          var width = clientX - this.startX;
          if (width >= this.reactWidth - this.sliderWidth) {
            width = this.reactWidth - this.sliderWidth;
          } else if (width <= 0) {
            width = 0;
          }
          this.width = width;
        }
      },
      end(e) {
        this.moveFlag = true;
        if (this.finishFlag) {
          if (this.width < this.reactWidth - this.sliderWidth) {
            this.width = 0;
          }
        } else {
          this.reset();
        }
      }
    }
  };
  function _sfc_render$n(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", {
      class: "react",
      id: "react"
    }, [
      vue.createElementVNode(
        "view",
        {
          style: vue.normalizeStyle({ width: $data.width + "px", backgroundColor: "#65B58A" }),
          class: "left kong"
        },
        vue.toDisplayString($data.title),
        5
        /* TEXT, STYLE */
      ),
      vue.createElementVNode(
        "view",
        {
          style: vue.normalizeStyle({ left: $data.width + "px" }),
          onTouchstart: _cache[0] || (_cache[0] = (...args) => $options.start && $options.start(...args)),
          onMousedown: _cache[1] || (_cache[1] = (...args) => $options.start && $options.start(...args)),
          onMouseup: _cache[2] || (_cache[2] = (...args) => $options.end && $options.end(...args)),
          onTouchmove: _cache[3] || (_cache[3] = (...args) => $options.move && $options.move(...args)),
          onMousemove: _cache[4] || (_cache[4] = (...args) => $options.move && $options.move(...args)),
          onTouchend: _cache[5] || (_cache[5] = (...args) => $options.end && $options.end(...args)),
          id: "slider",
          class: vue.normalizeClass({ slider: true, select: $data.title })
        },
        [
          !$data.title ? (vue.openBlock(), vue.createElementBlock("image", {
            key: 0,
            src: "https://img-blog.csdnimg.cn/a5bf3043a7d344cb88f762186c1dfc90.png#pic_center",
            mode: "widthFix"
          })) : (vue.openBlock(), vue.createElementBlock(
            vue.Fragment,
            { key: 1 },
            [
              vue.createCommentVNode(" 这是两张图片一个是大于号一个是对号"),
              vue.createElementVNode("image", {
                src: "https://img-blog.csdnimg.cn/e7b7798beb3a442f8a67a073b23c698a.png#pic_center",
                mode: "widthFix"
              })
            ],
            2112
            /* STABLE_FRAGMENT, DEV_ROOT_FRAGMENT */
          ))
        ],
        38
        /* CLASS, STYLE, NEED_HYDRATION */
      ),
      vue.createElementVNode("view", { class: "right kong" }, " 右滑提交 ")
    ]);
  }
  const sliderzz = /* @__PURE__ */ _export_sfc(_sfc_main$n, [["render", _sfc_render$n], ["__scopeId", "data-v-12311208"], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/components/sliderzz.vue"]]);
  const _sfc_main$m = {
    data() {
      return {
        visible: false,
        message: "",
        timeoutId: null
      };
    },
    methods: {
      showToast(message, duration = 2e3) {
        this.message = message;
        this.visible = true;
        clearTimeout(this.timeoutId);
        this.timeoutId = setTimeout(() => {
          this.visible = false;
        }, duration);
      },
      show() {
        this.showToast();
      }
    }
  };
  function _sfc_render$m(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.withDirectives((vue.openBlock(), vue.createElementBlock(
      "div",
      { class: "toast" },
      vue.toDisplayString($data.message),
      513
      /* TEXT, NEED_PATCH */
    )), [
      [vue.vShow, $data.visible]
    ]);
  }
  const toast = /* @__PURE__ */ _export_sfc(_sfc_main$m, [["render", _sfc_render$m], ["__scopeId", "data-v-c319bb4c"], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/components/toast.vue"]]);
  const _sfc_main$l = {
    data() {
      return {
        showBackTop: true,
        genderList: ["请选择性别", "男", "女"],
        genderIndex: 0
      };
    },
    components: {
      sliderzz,
      toast
    },
    methods: {
      change({ finish }) {
        finish();
      },
      showToast(message) {
        this.$refs.toast.showToast(message);
      },
      onGenderChange(e) {
        if (e.detail.value !== void 0) {
          this.genderIndex = e.detail.value;
        }
      }
    }
  };
  function _sfc_render$l(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_sliderzz = vue.resolveComponent("sliderzz");
    const _component_toast = vue.resolveComponent("toast");
    return vue.openBlock(), vue.createElementBlock("view", { class: "container" }, [
      vue.createElementVNode("view", { class: "scroll-view-container" }, [
        vue.createElementVNode("text", { class: "title" }, "以下是scroll-view的样例"),
        vue.createElementVNode("scroll-view", {
          class: "scroll-view",
          "scroll-x": "true"
        }, [
          vue.createCommentVNode(' scroll-x="true"表示横向滚动 '),
          vue.createElementVNode("view", { class: "list-container" }, [
            vue.createElementVNode("view", { class: "list" }, [
              vue.createElementVNode("view", { class: "item" }, "列表1 项目1")
            ]),
            vue.createElementVNode("view", { class: "list" }, [
              vue.createElementVNode("view", { class: "item" }, "列表2 项目1")
            ]),
            vue.createElementVNode("view", { class: "list" }, [
              vue.createElementVNode("view", { class: "item" }, "列表3 项目1")
            ]),
            vue.createElementVNode("view", { class: "list" }, [
              vue.createElementVNode("view", { class: "item" }, "列表4 项目1")
            ])
          ])
        ])
      ]),
      vue.createElementVNode("view", { class: "swiper-container" }, [
        vue.createElementVNode("text", { class: "title" }, "以下是swiper的样例"),
        vue.createElementVNode("swiper", {
          class: "swiper",
          "indicator-dots": true,
          autoplay: true,
          interval: 3e3,
          duration: 500,
          circular: true
        }, [
          vue.createCommentVNode(' indicator-dots="true"表示显示指示点 '),
          vue.createCommentVNode(' autoplay="true"表示自动播放 '),
          vue.createCommentVNode(' interval="3000"表示自动播放间隔时间 '),
          vue.createCommentVNode(' duration="500"表示切换动画时间 '),
          vue.createCommentVNode(' circular="true"表示循环播放 '),
          vue.createElementVNode("swiper-item", null, [
            vue.createElementVNode("view", { class: "swiper-item" }, "轮播1")
          ]),
          vue.createElementVNode("swiper-item", null, [
            vue.createElementVNode("view", { class: "swiper-item" }, "轮播2")
          ]),
          vue.createElementVNode("swiper-item", null, [
            vue.createElementVNode("view", { class: "swiper-item" }, "轮播3")
          ])
        ])
      ]),
      vue.createElementVNode("view", { style: { "height": "auto" } }, [
        vue.createVNode(_component_sliderzz, { onChange: $options.change }, null, 8, ["onChange"])
      ]),
      vue.createElementVNode("view", { class: "toast-container" }, [
        vue.createElementVNode("button", {
          onClick: _cache[0] || (_cache[0] = ($event) => $options.showToast("Hello World"))
        }, "点击我"),
        vue.createVNode(
          _component_toast,
          { ref: "toast" },
          null,
          512
          /* NEED_PATCH */
        )
      ]),
      vue.createElementVNode("view", { class: "jiliancontainer" }, [
        vue.createElementVNode("view", { class: "picker" }, [
          vue.createElementVNode("picker", {
            mode: "selector",
            range: $data.genderList,
            value: $data.genderIndex,
            onChange: _cache[1] || (_cache[1] = (...args) => $options.onGenderChange && $options.onGenderChange(...args))
          }, [
            vue.createElementVNode("view", { class: "picker-box" }, [
              vue.createElementVNode(
                "text",
                { class: "picker-text" },
                vue.toDisplayString($data.genderList[$data.genderIndex]),
                1
                /* TEXT */
              )
            ])
          ], 40, ["range", "value"])
        ])
      ])
    ]);
  }
  const PagesExampleExample = /* @__PURE__ */ _export_sfc(_sfc_main$l, [["render", _sfc_render$l], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/example/example.vue"]]);
  const _sfc_main$k = {
    data() {
      return {
        username: "",
        password: "",
        remember: false
      };
    },
    beforeMount() {
      let username = uni.getStorageSync("username");
      let password = uni.getStorageSync("password");
      let remember = uni.getStorageSync("remember");
      if (username && password && remember) {
        this.username = username;
        this.password = password;
        this.remember = remember;
      }
    },
    methods: {
      autoLogin() {
        this.remember = !this.remember;
        formatAppLog("log", "at pages/login/login.vue:57", this.remember);
        formatAppLog("log", "at pages/login/login.vue:58", this.username);
        formatAppLog("log", "at pages/login/login.vue:59", this.password);
      },
      login() {
        let flag = true;
        let username = this.username;
        if (!username) {
          this.$nextTick(() => {
            let usernameInput = document.getElementById("username");
            usernameInput.classList.add("inputActive");
            setTimeout(() => {
              usernameInput.classList.remove("inputActive");
            }, 2e3);
          });
          flag = false;
        }
        let password = this.password;
        if (!password) {
          this.$nextTick(() => {
            let passwordInput = document.getElementById("password");
            passwordInput.classList.add("inputActive");
            setTimeout(() => {
              passwordInput.classList.remove("inputActive");
            }, 2e3);
          });
          flag = false;
        }
        if (!flag) {
          return;
        }
        let remember = this.remember;
        uni.request({
          url: "/api/user/login",
          data: {
            username,
            password,
            remember
          },
          method: "POST",
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
            if (res.statusCode == 200) {
              if (remember) {
                uni.setStorageSync("username");
                uni.setStorageSync("password");
                uni.setStorageSync("remember");
              }
              let token = res.data.token;
              uni.setStorageSync("token", token);
              uni.navigateTo({
                //TODO: 跳转到首页，或处理其他逻辑
                url: res.data.have_word_book ? "../home/home" : `../Welcome/Welcome?operation=${res.data.have_word_book ? 1 : 0}`
              });
            } else if (res.statusCode == 400) {
              let usernameInput = document.getElementById("username");
              usernameInput.classList.add("inputActive");
              setTimeout(() => {
                usernameInput.classList.remove("inputActive");
              }, 2e3);
              let passwordInput = document.getElementById("password");
              passwordInput.classList.add("inputActive");
              setTimeout(() => {
                passwordInput.classList.remove("inputActive");
              }, 2e3);
              uni.showToast({
                title: "用户名或密码错误",
                icon: "none"
              });
            }
          },
          fail: (res) => {
            uni.showToast({
              title: "登录失败",
              icon: "none"
            });
          }
        });
      }
    }
  };
  function _sfc_render$k(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_span1 = vue.resolveComponent("span1");
    const _component_router_link = vue.resolveComponent("router-link");
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("view", { class: "all-container" }, [
        vue.createElementVNode("image", {
          class: "background",
          src: "/static/login1.svg"
        }),
        vue.createElementVNode("view", { class: "container" }, [
          vue.createVNode(_component_span1, null, {
            default: vue.withCtx(() => [
              vue.createTextVNode("Sign In")
            ]),
            _: 1
            /* STABLE */
          }),
          vue.createElementVNode("view", { class: "white-container1" }, [
            vue.createElementVNode("span", null, "账号"),
            vue.withDirectives(vue.createElementVNode(
              "input",
              {
                id: "username",
                class: "search-box",
                type: "text",
                "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => $data.username = $event),
                placeholder: "请输入账号"
              },
              null,
              512
              /* NEED_PATCH */
            ), [
              [vue.vModelText, $data.username]
            ])
          ]),
          vue.createElementVNode("view", { class: "white-container2" }, [
            vue.createElementVNode("span", null, "密码"),
            vue.createElementVNode("view", { class: "password-container" }, [
              vue.withDirectives(vue.createElementVNode(
                "input",
                {
                  id: "password",
                  class: "search-box",
                  type: "password",
                  "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => $data.password = $event),
                  placeholder: "请输入密码"
                },
                null,
                512
                /* NEED_PATCH */
              ), [
                [vue.vModelText, $data.password]
              ]),
              vue.createElementVNode(
                "img",
                {
                  ref: "errorIcon",
                  class: "error-icon",
                  src: "/static/errorCross.svg"
                },
                null,
                512
                /* NEED_PATCH */
              )
            ]),
            vue.createElementVNode("view", { class: "forgot-password-link" }, "忘记密码?")
          ]),
          vue.createElementVNode("button", {
            class: "login-button",
            onClick: _cache[2] || (_cache[2] = (...args) => $options.login && $options.login(...args))
          }, "登录"),
          vue.createCommentVNode(' 	<button class="register-button">注册</button><button class="forget">忘记密码？</button>\r\n				<button class="button1"></button>\r\n				<span class="text">登录代表你同意用户协议、隐私政策和儿童隐私政策</span> '),
          vue.createElementVNode("span", { class: "text" }, [
            vue.createTextVNode("have no account? "),
            vue.createVNode(_component_router_link, { to: "../register/register" }, {
              default: vue.withCtx(() => [
                vue.createTextVNode("click here")
              ]),
              _: 1
              /* STABLE */
            })
          ])
        ])
      ])
    ]);
  }
  const PagesLoginLogin = /* @__PURE__ */ _export_sfc(_sfc_main$k, [["render", _sfc_render$k], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/login/login.vue"]]);
  const _sfc_main$j = {
    data() {
      return {
        username: "",
        email: "",
        password: "",
        password2: "",
        initialVerifyCodeInput: "",
        verifyCode: ""
      };
    },
    methods: {
      checkEmail() {
        let regex = /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/;
        let emailInput = document.getElementById("emailInput");
        if (!regex.test(this.email)) {
          emailInput.classList.add("inputActive");
          setTimeout(() => {
            emailInput.classList.remove("inputActive");
          }, 2e3);
          return false;
        }
        return true;
      },
      sendCode() {
        if (!this.checkEmail()) {
          return;
        }
        let btn = document.getElementById("sendCodeBtn");
        uni.request({
          url: "/api/register/sendCode",
          data: {
            email: this.email
          },
          withCredentials: false,
          method: "POST",
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
            formatAppLog("log", "at pages/register/register.vue:73", res);
            if (res.statusCode === 200) {
              let vCode = res.data.data;
              let codeAndExpiry = {
                vCode,
                expiry: (/* @__PURE__ */ new Date()).getTime() + 1e3 * 60 * 5
                //5分钟有效
              };
              uni.setStorageSync("codeAndExpiry", codeAndExpiry);
              let timeLeft = 60;
              btn.disabled = true;
              let timer = setInterval(() => {
                timeLeft--;
                btn.innerText = `${timeLeft}秒后请重试`;
                if (timeLeft <= 0) {
                  clearInterval(timer);
                  btn.innerText = "发送验证码";
                  btn.disabled = false;
                }
              }, 1e3);
            } else if (res.statusCode == 409) {
              uni.showToast({
                title: "邮箱已注册",
                icon: "error"
              });
            } else if (res.statusCode == 400) {
              uni.showToast({
                title: "请求参数错误",
                icon: "error"
              });
            } else {
              uni.showToast({
                title: "发送失败",
                icon: "error"
              });
            }
          },
          fail: (res) => {
            formatAppLog("log", "at pages/register/register.vue:113", res);
            uni.showToast({
              title: "发送失败",
              icon: "error"
            });
          }
        });
      },
      checkInput() {
        let verifyCode = this.verifyCode;
        let verifyCodeInput = document.getElementById("verifyCodeInput");
        if (verifyCode !== "" && !this.checkVerifyCode(verifyCode)) {
          verifyCodeInput.classList.add("inputActive");
          setTimeout(() => {
            verifyCodeInput.classList.remove("inputActive");
          }, 2e3);
          return false;
        }
        return true;
      },
      checkVerifyCode(verifyCode) {
        let codeAndExpiry = uni.getStorageSync("codeAndExpiry");
        if (codeAndExpiry) {
          let now = (/* @__PURE__ */ new Date()).getTime();
          if (now < codeAndExpiry.expiry) {
            if (verifyCode == codeAndExpiry.vCode) {
              return true;
            }
          }
        }
        return false;
      },
      //密码输入框失去焦点时检查密码是否一致
      checkPasswordInput() {
        let password = this.password;
        let password2 = this.password2;
        let passwordInput = document.getElementById("passwordInput");
        if (password !== "" && password2 !== "" && password !== password2) {
          passwordInput.classList.add("inputActive");
          setTimeout(() => {
            passwordInput.classList.remove("inputActive");
          }, 2e3);
          uni.showToast({
            title: "两次输入的密码不一致",
            icon: "error"
          });
          return false;
        }
        return true;
      },
      //注册
      register() {
        let username = this.username;
        let email = this.email;
        let password = this.password;
        uni.request({
          url: "/api/user/register",
          data: {
            username,
            email,
            password
          },
          method: "POST",
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
            formatAppLog("log", "at pages/register/register.vue:181", res);
            if (res.statusCode === 200) {
              uni.showToast({
                title: "注册成功",
                icon: "success"
              });
              setTimeout(() => {
                uni.navigateBack();
              }, 1e3);
            } else if (data.code == 409) {
              uni.showToast({
                title: "用户名已注册",
                icon: "error"
              });
            } else if (data.code == 400) {
              uni.showToast({
                title: "请求参数错误",
                icon: "error"
              });
            } else {
              uni.showToast({
                title: "注册失败",
                icon: "error"
              });
            }
          },
          fail: (res) => {
            formatAppLog("log", "at pages/register/register.vue:208", res);
            uni.showToast({
              title: "注册失败",
              icon: "error"
            });
          }
        });
      }
    }
  };
  function _sfc_render$j(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("view", { class: "background" }, [
        vue.createElementVNode("span", { class: "span1" }, [
          vue.createElementVNode("span", { class: "sign" }, "Sign"),
          vue.createElementVNode("br"),
          vue.createTextVNode("Up")
        ]),
        vue.createElementVNode("image", {
          class: "pic",
          src: "/static/register1.svg"
        })
      ]),
      vue.createElementVNode("view", {
        class: "input-container",
        id: "1"
      }, [
        vue.createElementVNode("span", null, "昵称"),
        vue.withDirectives(vue.createElementVNode(
          "input",
          {
            class: "input",
            "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => $data.username = $event),
            type: "text",
            placeholder: "请输入昵称"
          },
          null,
          512
          /* NEED_PATCH */
        ), [
          [vue.vModelText, $data.username]
        ])
      ]),
      vue.createElementVNode("view", {
        class: "input-container",
        id: "2"
      }, [
        vue.createElementVNode("span", null, "邮箱"),
        vue.withDirectives(vue.createElementVNode(
          "input",
          {
            class: "input",
            id: "emailInput",
            ref: "email",
            "onUpdate:modelValue": _cache[1] || (_cache[1] = ($event) => $data.email = $event),
            type: "text",
            placeholder: "请输入邮箱"
          },
          null,
          512
          /* NEED_PATCH */
        ), [
          [vue.vModelText, $data.email]
        ]),
        vue.createElementVNode(
          "button",
          {
            class: "vtBtn",
            ref: "sendCodeBtn",
            id: "sendCodeBtn",
            onClick: _cache[2] || (_cache[2] = (...args) => $options.sendCode && $options.sendCode(...args))
          },
          "发送验证码",
          512
          /* NEED_PATCH */
        )
      ]),
      vue.createElementVNode("view", {
        class: "input-container",
        id: "3"
      }, [
        vue.createElementVNode("span", { style: { "left": "2rem" } }, "验证码"),
        vue.withDirectives(vue.createElementVNode(
          "input",
          {
            "onUpdate:modelValue": _cache[3] || (_cache[3] = ($event) => $data.verifyCode = $event),
            id: "verifyCodeInput",
            class: "input",
            onBlur: _cache[4] || (_cache[4] = (...args) => $options.checkInput && $options.checkInput(...args)),
            type: "text",
            placeholder: "请输入验证码"
          },
          null,
          544
          /* NEED_HYDRATION, NEED_PATCH */
        ), [
          [vue.vModelText, $data.verifyCode]
        ])
      ]),
      vue.createElementVNode("view", {
        class: "input-container",
        id: "4"
      }, [
        vue.createElementVNode("span", null, "密码"),
        vue.withDirectives(vue.createElementVNode(
          "input",
          {
            class: "input",
            "onUpdate:modelValue": _cache[5] || (_cache[5] = ($event) => $data.password = $event),
            type: "password",
            placeholder: "请输入密码"
          },
          null,
          512
          /* NEED_PATCH */
        ), [
          [vue.vModelText, $data.password]
        ])
      ]),
      vue.createElementVNode("view", {
        class: "input-container",
        id: "5"
      }, [
        vue.withDirectives(vue.createElementVNode(
          "input",
          {
            class: "input",
            id: "passwordInput",
            "onUpdate:modelValue": _cache[6] || (_cache[6] = ($event) => $data.password2 = $event),
            onBlur: _cache[7] || (_cache[7] = (...args) => $options.checkPasswordInput && $options.checkPasswordInput(...args)),
            type: "password",
            placeholder: "再次输入密码"
          },
          null,
          544
          /* NEED_HYDRATION, NEED_PATCH */
        ), [
          [vue.vModelText, $data.password2]
        ])
      ]),
      vue.createElementVNode("button", {
        class: "login-button",
        onClick: _cache[8] || (_cache[8] = (...args) => $options.register && $options.register(...args))
      }, "注册"),
      vue.createElementVNode("span", { class: "text" }, [
        vue.createTextVNode("Already have account?"),
        vue.createElementVNode("a", null, "click here to login")
      ])
    ]);
  }
  const PagesRegisterRegister = /* @__PURE__ */ _export_sfc(_sfc_main$j, [["render", _sfc_render$j], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/register/register.vue"]]);
  const _sfc_main$i = {
    data() {
      return {
        // 按钮id映射表
        buttonIds: [
          "cet4",
          "cet6"
        ],
        gradeDescriptions: {
          1: "小学一年级",
          2: "小学二年级",
          3: "小学三年级",
          4: "小学四年级",
          5: "小学五年级",
          6: "小学六年级",
          7: "初中一年级",
          8: "初中二年级",
          9: "初中三年级",
          10: "高中一年级",
          11: "高中二年级",
          12: "高中三年级",
          13: "四级",
          14: "六级"
        },
        defaultGradeDescription: {
          1: "小学",
          2: "初中",
          3: "高中",
          4: "四级",
          5: "六级",
          6: "其他"
        },
        books: [
          {
            book_id: 1,
            title: "四级词汇大全",
            decsrip: "四级最新考纲单词全收录，时候所有备考四级的同学",
            grade: "四级",
            gradeId: 1,
            num: "共4440词"
          },
          {
            book_id: 2,
            title: "四级高频",
            decsrip: "精选四级真题超高频词",
            grade: "四级",
            gradeId: 1,
            num: "共739词"
          },
          {
            book_id: 3,
            title: "四级高频",
            decsrip: "精选四级真题超高频词",
            grade: "四级",
            gradeId: 1,
            num: "共739词"
          },
          {
            book_id: 4,
            title: "四级高频",
            decsrip: "精选四级真题超高频词",
            grade: "四级",
            gradeId: 1,
            num: "共739词"
          },
          {
            book_id: 5,
            title: "四级高频",
            decsrip: "精选四级真题超高频词",
            grade: "四级",
            gradeId: 1,
            num: "共739词"
          },
          {
            book_id: 6,
            title: "四级高频",
            decsrip: "精选四级真题超高频词",
            grade: "四级",
            gradeId: 1,
            num: "共739词"
          },
          {
            book_id: 7,
            title: "六级词汇大全",
            decsrip: "六级最新考纲单词全收录，时候所有备考六级的同学",
            grade: "六级",
            gradeId: 2,
            num: "共6204词"
          },
          {
            book_id: 8,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 9,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 10,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 11,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 12,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 13,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 14,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 15,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 16,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 17,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 18,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          },
          {
            book_id: 19,
            title: "六级核心（过考版）",
            decsrip: "精选六级真题超高频词",
            grade: "六级",
            gradeId: 2,
            num: "共2551词"
          }
        ],
        // 词书列表
        activeTab: 0,
        // 默认选中第一个选项卡
        tabs: ["全部", "大学", "高中", "初中", "小学", "留学", "其他"],
        // 选项卡数组
        activeButton: null,
        // 记录当前活跃的按钮,
        operation: 0
        // 记录下一个应该跳转的页面
      };
    },
    onLoad(event) {
      this.operation = parseInt(event.operation);
      this.books = [];
      this.buttonIds = [];
      uni.request({
        url: "/api/users/navigate_books",
        method: "GET",
        header: {
          "Authorization": "Bearer " + uni.getStorageSync("token")
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
          formatAppLog("log", "at pages/Welcome/Welcome.vue:250", res.data);
          let books = res.data.books;
          books.forEach((book) => {
            this.books.push({
              book_id: book.book_id,
              title: book.book_name,
              decsrip: book.description,
              grade: book.grade_description,
              gradeId: book.grade,
              num: book.word_num
            });
            this.buttonIds.push(book.grade);
          });
        },
        fail: (err) => {
          formatAppLog("log", "at pages/Welcome/Welcome.vue:267", err);
        }
      });
    },
    methods: {
      getImageUrl(id) {
        return `../../static/${id}.jpg`;
      },
      filteredBooks(defaultId) {
        defaultId = parseInt(defaultId);
        let returnBookIds;
        switch (defaultId) {
          case 1:
            returnBookIds = [1, 2, 3, 4, 5, 6];
            break;
          case 2:
            returnBookIds = [7, 8, 9];
            break;
          case 3:
            returnBookIds = [10, 11, 12];
            break;
          case 4:
            returnBookIds = [13];
            break;
          case 5:
            returnBookIds = [14];
            break;
          default:
            returnBookIds = [0];
            break;
        }
        let returnBooks = [];
        this.books.forEach((book) => {
          if (returnBookIds.includes(book.gradeId)) {
            returnBooks.push(book);
          }
        });
        return returnBooks;
      },
      changeTab(tabNumber) {
        this.activeTab = tabNumber;
      },
      scrollToIdButton(id) {
        this.scrollToElement(id);
        this.activeButton = id;
      },
      scrollToElement(id) {
        const el = document.getElementById(id);
        if (el) {
          const offset = 110;
          const scrollPosition = el.offsetTop - offset;
          window.scrollTo({
            top: scrollPosition,
            behavior: "smooth"
          });
        }
      },
      bookConfirm(book_id, title) {
        formatAppLog("log", "at pages/Welcome/Welcome.vue:330", book_id, title);
        uni.showModal({
          title: "提示",
          content: "确定要选择《" + title + "》吗？",
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
            if (res.confirm) {
              uni.request({
                url: "/api/users/navigate_books",
                method: "POST",
                header: {
                  "Authorization": "Bearer " + uni.getStorageSync("token")
                },
                data: {
                  book_id
                },
                success: (res2) => {
                  formatAppLog("log", "at pages/Welcome/Welcome.vue:346", res2.data);
                  if (res2.data.code === 200 || res2.data.code === "200") {
                    uni.showToast({
                      title: "选择词库成功",
                      icon: "none"
                    });
                    setTimeout(() => {
                      if (this.operation === 0) {
                        uni.switchTab({
                          url: "../home/home"
                        });
                        return;
                      }
                      uni.navigateBack();
                    });
                  } else {
                    uni.showToast({
                      title: "选择词库失败",
                      icon: "none"
                    });
                  }
                },
                fail: (err) => {
                  formatAppLog("log", "at pages/Welcome/Welcome.vue:369", err);
                }
              });
            } else if (res.cancel) {
              formatAppLog("log", "at pages/Welcome/Welcome.vue:373", "取消");
            }
          }
        });
      }
    }
  };
  function _sfc_render$i(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("view", { class: "topbar" }, [
        vue.createElementVNode("span", { class: "host-title" }, "选择您的打卡计划"),
        vue.createElementVNode("span", { class: "skip-container" }, [
          vue.createElementVNode("a", {
            href: "",
            class: "skip"
          }, "跳过")
        ])
      ]),
      vue.createCommentVNode(" 添加选项卡 "),
      vue.createElementVNode("view", { class: "tab-bar-container" }, [
        vue.createElementVNode("view", { class: "tab-bar" }, [
          (vue.openBlock(true), vue.createElementBlock(
            vue.Fragment,
            null,
            vue.renderList($data.tabs, (tab, index) => {
              return vue.openBlock(), vue.createElementBlock("view", {
                class: vue.normalizeClass(["tab-item", { "active": $data.activeTab === index }]),
                key: index,
                onClick: ($event) => $options.changeTab(index)
              }, vue.toDisplayString(tab), 11, ["onClick"]);
            }),
            128
            /* KEYED_FRAGMENT */
          ))
        ])
      ]),
      vue.createCommentVNode(" 添加按钮行 "),
      vue.createElementVNode("view", { class: "button-row" }, [
        (vue.openBlock(true), vue.createElementBlock(
          vue.Fragment,
          null,
          vue.renderList($data.defaultGradeDescription, (defaultGrade, defaultGradeId) => {
            return vue.openBlock(), vue.createElementBlock("view", {
              class: vue.normalizeClass(["button", { "active": $data.activeButton === defaultGradeId }]),
              key: defaultGradeId,
              onClick: ($event) => $options.scrollToIdButton(defaultGradeId)
            }, vue.toDisplayString($data.defaultGradeDescription[defaultGradeId]), 11, ["onClick"]);
          }),
          128
          /* KEYED_FRAGMENT */
        )),
        vue.createCommentVNode(" 添加其他按钮 ")
      ]),
      vue.createElementVNode("view", { class: "book-list" }, [
        (vue.openBlock(true), vue.createElementBlock(
          vue.Fragment,
          null,
          vue.renderList($data.defaultGradeDescription, (defaultDesc, defaultId) => {
            return vue.openBlock(), vue.createElementBlock("view", {
              class: "book-type",
              key: defaultId,
              id: defaultId
            }, [
              (vue.openBlock(true), vue.createElementBlock(
                vue.Fragment,
                null,
                vue.renderList($options.filteredBooks(defaultId), (book, index) => {
                  return vue.openBlock(), vue.createElementBlock("view", {
                    class: "book-container",
                    key: index,
                    onClick: ($event) => $options.bookConfirm(book.book_id, book.title)
                  }, [
                    vue.createElementVNode("image", {
                      src: $options.getImageUrl(defaultId)
                    }, null, 8, ["src"]),
                    vue.createElementVNode("view", { class: "text-container" }, [
                      vue.createElementVNode(
                        "span",
                        { class: "book-title" },
                        vue.toDisplayString(book.title),
                        1
                        /* TEXT */
                      ),
                      vue.createElementVNode(
                        "span",
                        { class: "discrip" },
                        vue.toDisplayString(book.decsrip),
                        1
                        /* TEXT */
                      ),
                      vue.createElementVNode(
                        "span",
                        { class: "num" },
                        vue.toDisplayString(book.num),
                        1
                        /* TEXT */
                      )
                    ])
                  ], 8, ["onClick"]);
                }),
                128
                /* KEYED_FRAGMENT */
              ))
            ], 8, ["id"]);
          }),
          128
          /* KEYED_FRAGMENT */
        ))
      ])
    ]);
  }
  const PagesWelcomeWelcome = /* @__PURE__ */ _export_sfc(_sfc_main$i, [["render", _sfc_render$i], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/Welcome/Welcome.vue"]]);
  const _sfc_main$h = {
    data() {
      return {};
    },
    methods: {}
  };
  function _sfc_render$h(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", { class: "loading-container" }, [
      vue.createElementVNode("view", { class: "loading-text" }, "页面加载中..."),
      vue.createElementVNode("view", { class: "loading-arrow" })
    ]);
  }
  const PagesLoadingLoading = /* @__PURE__ */ _export_sfc(_sfc_main$h, [["render", _sfc_render$h], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/Loading/Loading.vue"]]);
  const _sfc_main$g = {
    data() {
      return {
        showPopup: true,
        popupMessage: ""
      };
    },
    onLoad() {
      this.openPopup("确定要删除吗？");
    },
    methods: {
      openPopup(message) {
        this.popupMessage = message;
        this.showPopup = true;
      },
      closePopup() {
        this.showPopup = false;
      },
      handleConfirm() {
        formatAppLog("log", "at pages/PopUp/PopUp.vue:33", "确认删除");
        this.closePopup();
      },
      handleCancel() {
        formatAppLog("log", "at pages/PopUp/PopUp.vue:37", "取消删除");
        this.closePopup();
      }
    }
  };
  function _sfc_render$g(_ctx, _cache, $props, $setup, $data, $options) {
    return $data.showPopup ? (vue.openBlock(), vue.createElementBlock("view", {
      key: 0,
      class: "popup-container"
    }, [
      vue.createElementVNode("view", { class: "popup-content" }, [
        vue.createElementVNode(
          "view",
          null,
          vue.toDisplayString($data.popupMessage),
          1
          /* TEXT */
        ),
        vue.createElementVNode("view", { class: "close-button-container" }, [
          vue.createElementVNode("view", {
            class: "button",
            onClick: _cache[0] || (_cache[0] = (...args) => $options.handleConfirm && $options.handleConfirm(...args))
          }, "取消"),
          vue.createElementVNode("view", {
            class: "button",
            onClick: _cache[1] || (_cache[1] = (...args) => $options.handleCancel && $options.handleCancel(...args))
          }, "删除")
        ])
      ])
    ])) : vue.createCommentVNode("v-if", true);
  }
  const PagesPopUpPopUp = /* @__PURE__ */ _export_sfc(_sfc_main$g, [["render", _sfc_render$g], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/PopUp/PopUp.vue"]]);
  const _sfc_main$f = {
    props: {
      label: String,
      value: String,
      placeholder: String
    },
    data() {
      return {
        showModal: false,
        inputValue: this.value
      };
    },
    methods: {
      handleOutsideClick(event) {
        if (!this.$el.contains(event.target)) {
          this.showModal = false;
        }
      },
      confirmEdit() {
        this.$emit("update:value", this.inputValue);
        this.showModal = false;
      },
      cancelEdit() {
        this.showModal = false;
        this.inputValue = this.value;
      }
    }
  };
  function _sfc_render$f(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", {
      class: "item",
      onClick: _cache[4] || (_cache[4] = ($event) => $data.showModal = true)
    }, [
      vue.createElementVNode(
        "text",
        null,
        vue.toDisplayString($props.label),
        1
        /* TEXT */
      ),
      !$data.showModal ? (vue.openBlock(), vue.createElementBlock("input", {
        key: 0,
        class: "1",
        value: $props.value,
        disabled: "",
        placeholder: "点击编辑"
      }, null, 8, ["value"])) : vue.createCommentVNode("v-if", true),
      $data.showModal ? (vue.openBlock(), vue.createElementBlock("view", {
        key: 1,
        class: "modal",
        onClick: _cache[3] || (_cache[3] = vue.withModifiers(($event) => $options.handleOutsideClick($event), ["stop"]))
      }, [
        vue.createElementVNode("view", { class: "modal-content" }, [
          vue.createElementVNode("view", { class: "coolinput" }, [
            vue.createElementVNode("label", {
              for: "teamNameInput",
              class: "text"
            }, "团队名*:"),
            vue.withDirectives(vue.createElementVNode("input", {
              "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => $data.inputValue = $event),
              placeholder: $props.placeholder,
              class: "inputfield"
            }, null, 8, ["placeholder"]), [
              [vue.vModelText, $data.inputValue]
            ])
          ]),
          vue.createElementVNode("view", { class: "buttons" }, [
            vue.createElementVNode("button", {
              onClick: _cache[1] || (_cache[1] = (...args) => $options.cancelEdit && $options.cancelEdit(...args)),
              class: "button-left"
            }, "取消"),
            vue.createElementVNode("button", {
              onClick: _cache[2] || (_cache[2] = (...args) => $options.confirmEdit && $options.confirmEdit(...args)),
              class: "button-right"
            }, "确定")
          ])
        ])
      ])) : vue.createCommentVNode("v-if", true)
    ]);
  }
  const EditableItemView = /* @__PURE__ */ _export_sfc(_sfc_main$f, [["render", _sfc_render$f], ["__scopeId", "data-v-7226b2f7"], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/components/EditableItemView.vue"]]);
  const _sfc_main$e = {
    components: {
      EditableItemView
    },
    data() {
      return {
        username: "wwswad",
        email: "3242412@xy.com",
        phone: "12312312312",
        birthday: "1999-01-01",
        team: "无"
        // 其他数据...  
      };
    },
    methods: {
      GotoUser() {
        uni.switchTab({
          url: "../user/user"
        });
      }
      // 如果还有其他方法，可以在这里继续添加  
    }
    // 注意这里没有多余的闭合大括号 "}"  
  };
  function _sfc_render$e(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_EditableItemView = vue.resolveComponent("EditableItemView");
    return vue.openBlock(), vue.createElementBlock("view", { class: "container" }, [
      vue.createElementVNode("image", {
        class: "back-icon",
        src: "/static/back.svg",
        onClick: _cache[0] || (_cache[0] = (...args) => $options.GotoUser && $options.GotoUser(...args))
      }),
      vue.createElementVNode("view", { class: "edit" }, [
        vue.createElementVNode("image", {
          class: "photo",
          src: "/static/logo.png"
        }),
        vue.createElementVNode("view", { class: "items" }, [
          vue.createVNode(_component_EditableItemView, {
            id: "username",
            label: "用户名",
            value: $data.username,
            placeholder: "wwswad",
            "onUpdate:value": _cache[1] || (_cache[1] = ($event) => $data.username = $event)
          }, null, 8, ["value"]),
          vue.createVNode(_component_EditableItemView, {
            id: "email",
            label: "邮箱",
            value: $data.email,
            placeholder: "3242412@xy.com",
            "onUpdate:value": _cache[2] || (_cache[2] = ($event) => $data.email = $event)
          }, null, 8, ["value"]),
          vue.createVNode(_component_EditableItemView, {
            id: "phone",
            label: "手机号",
            value: $data.phone,
            placeholder: "12312312312",
            "onUpdate:value": _cache[3] || (_cache[3] = ($event) => $data.phone = $event)
          }, null, 8, ["value"]),
          vue.createVNode(_component_EditableItemView, {
            id: "birthday",
            label: "生日",
            value: $data.birthday,
            placeholder: "1999-01-01",
            "onUpdate:value": _cache[4] || (_cache[4] = ($event) => $data.birthday = $event)
          }, null, 8, ["value"]),
          vue.createVNode(_component_EditableItemView, {
            id: "team",
            label: "团队",
            value: $data.team,
            placeholder: "无",
            "onUpdate:value": _cache[5] || (_cache[5] = ($event) => $data.team = $event)
          }, null, 8, ["value"])
        ])
      ])
    ]);
  }
  const PagesPersonalInformationPersonalInformation = /* @__PURE__ */ _export_sfc(_sfc_main$e, [["render", _sfc_render$e], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/personal-information/personal-information.vue"]]);
  const _sfc_main$d = {
    data() {
      return {
        //最后一次打卡的日期
        lastPunchDate: new Date(2024, 4, 4),
        year: 2024,
        month: 5,
        dates: [],
        // 存储当前月份的日期
        selectedDate: null,
        //当前选择的日期
        // 考试信息
        punchMsg: 63398853,
        //0000 0011 1100 0111 0110 0011 1100 0101,
        //存储前32天打卡信息，
        //每一位表示一天，0表示未打卡，1表示已打卡
        chosenYear: 2024,
        // 选中的年份
        chosenMonth: 5,
        // 选中的月份
        chosenDay: 4,
        // 选中的日期
        punchWords: [
          {
            word: "refuse",
            meanings: {
              verb: "拒绝,谢绝",
              noun: "废物",
              adj: "扔掉的，无用的",
              adv: null,
              prep: null
            }
          },
          {
            word: "objective",
            meanings: {
              verb: null,
              noun: "目的，目标，<语法>宾格，物镜",
              adj: "客观的，<语法>宾格的，真实的，目标的",
              adv: null,
              prep: null
            }
          }
        ]
      };
    },
    beforeMount() {
      this.generateDates();
    },
    methods: {
      //TODO 发送网络请求获取考试信息
      requestExamMsg() {
      },
      //todo 使用该函数计算某一月份需要占据日历表多少行
      getCalendarRows(year, month) {
        const firstDay = new Date(year, month - 1, 1);
        const firstDayOfWeek = firstDay.getDay();
        const daysInMonth = new Date(year, month, 0).getDate();
        return firstDayOfWeek === 0 || daysInMonth + firstDayOfWeek > 28 ? 5 : 4;
      },
      //判断是否有未完成的打卡计划
      getChosenDateFromDates() {
        let dateIndex = (new Date(this.chosenYear, this.chosenMonth - 1, this.chosenDay) - this.dates[0].date) / (1e3 * 60 * 60 * 24);
        dateIndex = Math.floor(dateIndex);
        if (dateIndex < 0 || dateIndex >= this.dates.length) {
          return -1;
        }
        if (this.dates[dateIndex].afterLastPunchDay) {
          return 1;
        }
        if (this.dates[dateIndex].expired) {
          return -1;
        }
        return this.dates[dateIndex].hasExam;
      },
      subMonth() {
        if ((/* @__PURE__ */ new Date()).getMonth() - this.month > 0) {
          return;
        }
        this.month--;
        if (this.month < 1) {
          this.month = 12;
          this.year--;
        }
        this.generateDates();
      },
      addMonth() {
        if (this.month - (/* @__PURE__ */ new Date()).getMonth() > 2) {
          return;
        }
        this.month++;
        if (this.month > 12) {
          this.month = 1;
          this.year++;
        }
        this.generateDates();
      },
      handleClick(date) {
        let year = date.date.getFullYear();
        let month = date.date.getMonth() + 1;
        let day = date.date.getDate();
        formatAppLog("log", "at pages/Calendar/Calendar.vue:187", year, month, day);
        this.selectedDate = date;
        this.chosenYear = year;
        this.chosenMonth = month;
        this.chosenDay = day;
        if (date.hasExam) {
          formatAppLog("log", "at pages/Calendar/Calendar.vue:193", "未完成打卡计划日期：", date.value);
        } else {
          formatAppLog("log", "at pages/Calendar/Calendar.vue:196", "已完成打卡计划日期：", date.value);
        }
      },
      generateDates() {
        const firstDay = new Date(this.year, this.month - 1, 1);
        const firstDayOfWeek = firstDay.getDay();
        const totalDays = new Date(this.year, this.month, 0).getDate();
        this.dates = [];
        for (let i = firstDayOfWeek - 1; i >= 0; i--) {
          let date = new Date(firstDay - (i + 1) * 24 * 60 * 60 * 1e3);
          let day = date.getDate();
          this.dates.push({
            date,
            value: day,
            dayOfWeek: "",
            hasExam: -1
          });
        }
        for (let i = 1; i <= totalDays; i++) {
          let dayOfWeek = (firstDayOfWeek + i - 1) % 7;
          let date = new Date(this.year, this.month - 1, i);
          let today = this.lastPunchDate;
          let diffDays = Math.floor((today - date) / (24 * 60 * 60 * 1e3));
          let d = /* @__PURE__ */ new Date();
          let expired = diffDays >= 32 || diffDays < 0;
          let hasExam = !expired ? this.punchMsg >> diffDays & 1 : false;
          this.dates.push({
            date,
            value: i,
            dayOfWeek,
            hasExam,
            // 判断当前日期是否有考试
            expired,
            //判断当前日期是否过期
            afterLastPunchDay: date > today && date < d
          });
        }
        for (let i = 0; i < 7 - (totalDays + firstDayOfWeek) % 7; i++) {
          let tempDate = new Date(this.year, this.month, i + 1);
          this.dates.push({
            date: tempDate,
            value: tempDate.getDate(),
            dayOfWeek: "",
            hasExam: -1
          });
        }
      },
      sign_again() {
        uni.switchTab({
          url: "../Examination/Examination"
        });
      }
    }
  };
  function _sfc_render$d(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", { class: "body" }, [
      vue.createElementVNode("view", { class: "color" }),
      vue.createElementVNode("view", null, [
        vue.createElementVNode("view", { class: "calendar" }, [
          vue.createElementVNode("view", { class: "head" }, [
            vue.createElementVNode(
              "text",
              { class: "date" },
              vue.toDisplayString($data.year) + "年" + vue.toDisplayString($data.month) + "月",
              1
              /* TEXT */
            ),
            vue.createElementVNode("button", {
              class: "last-or-next",
              onClick: _cache[0] || (_cache[0] = (...args) => $options.subMonth && $options.subMonth(...args)),
              style: { "margin-left": "180rpx" }
            }, [
              vue.createElementVNode("image", {
                class: "icon",
                src: "/static/last.svg"
              })
            ]),
            vue.createElementVNode("button", {
              class: "last-or-next",
              onClick: _cache[1] || (_cache[1] = (...args) => $options.addMonth && $options.addMonth(...args))
            }, [
              vue.createElementVNode("image", {
                class: "icon",
                src: "/static/next.svg"
              })
            ])
          ]),
          vue.createElementVNode("view", { class: "week" }, [
            vue.createElementVNode("text", { class: "week1" }, "日"),
            vue.createElementVNode("text", { class: "week2" }, "一"),
            vue.createElementVNode("text", { class: "week3" }, "二"),
            vue.createElementVNode("text", { class: "week4" }, "三"),
            vue.createElementVNode("text", { class: "week5" }, "四"),
            vue.createElementVNode("text", { class: "week6" }, "五"),
            vue.createElementVNode("text", { class: "week7" }, "六")
          ]),
          vue.createElementVNode("view", { class: "day" }, [
            (vue.openBlock(true), vue.createElementBlock(
              vue.Fragment,
              null,
              vue.renderList($data.dates, (date, index) => {
                return vue.openBlock(), vue.createElementBlock("view", {
                  class: vue.normalizeClass(["date-item", {
                    "not_finish": date.hasExam === 1,
                    "finish": date.hasExam === 0,
                    "not-this-month": date.hasExam === -1,
                    "sunday": date.dayOfWeek === 0,
                    "saturday": date.dayOfWeek === 6,
                    "selected": date === $data.selectedDate
                  }]),
                  key: index,
                  onClick: ($event) => $options.handleClick(date)
                }, [
                  vue.createTextVNode(
                    vue.toDisplayString(date.value) + " ",
                    1
                    /* TEXT */
                  ),
                  vue.createCommentVNode(' <span class="badge" v-if="date.hasExam"></span> ')
                ], 10, ["onClick"]);
              }),
              128
              /* KEYED_FRAGMENT */
            ))
          ])
        ])
      ]),
      vue.createElementVNode("view", { class: "examMsg" }, [
        vue.createElementVNode(
          "text",
          { class: "title" },
          vue.toDisplayString($data.chosenMonth) + "月" + vue.toDisplayString($data.chosenDay) + "日",
          1
          /* TEXT */
        ),
        vue.createElementVNode("view", { class: "card-container" }, [
          $options.getChosenDateFromDates() == 1 ? (vue.openBlock(), vue.createElementBlock("view", {
            key: 0,
            class: "card",
            id: "daka"
          }, [
            vue.createElementVNode("image", {
              src: "/static/not-done.svg",
              style: { "margin-left": "48rpx" }
            }),
            vue.createElementVNode("text", { class: "title" }, "打卡计划:"),
            vue.createElementVNode("text", { class: "state" }, "未完成"),
            vue.createCommentVNode('						<button @click="sign_again">补签</button>')
          ])) : $options.getChosenDateFromDates() == 0 ? (vue.openBlock(), vue.createElementBlock("view", { key: 1 }, [
            vue.createElementVNode("view", {
              class: "card",
              id: "daka"
            }, [
              vue.createElementVNode("image", { src: "/static/done.svg" }),
              vue.createElementVNode("text", { class: "title" }, "打卡计划:"),
              vue.createElementVNode("text", { class: "state" }, "已完成")
            ]),
            vue.createElementVNode("view", { class: "words" }, [
              vue.createCommentVNode("todo 为这个span添加样式"),
              (vue.openBlock(true), vue.createElementBlock(
                vue.Fragment,
                null,
                vue.renderList($data.punchWords, (word, index) => {
                  return vue.openBlock(), vue.createElementBlock(
                    "span",
                    { key: index },
                    vue.toDisplayString(word.word) + "： " + vue.toDisplayString(word.meanings.verb != null ? "v." : "") + vue.toDisplayString(word.meanings.verb) + " " + vue.toDisplayString(word.meanings.noun != null ? "n." : "") + " " + vue.toDisplayString(word.meanings.noun) + " " + vue.toDisplayString(word.meanings.adj != null ? "adj." : "") + " " + vue.toDisplayString(word.meanings.adj) + " " + vue.toDisplayString(word.meanings.adv != null ? "adv." : "") + " " + vue.toDisplayString(word.meanings.prep) + " " + vue.toDisplayString(word.meanings.prep != null ? "prep." : "") + " " + vue.toDisplayString(word.meanings.adv),
                    1
                    /* TEXT */
                  );
                }),
                128
                /* KEYED_FRAGMENT */
              ))
            ])
          ])) : (vue.openBlock(), vue.createElementBlock("view", {
            key: 2,
            class: "card",
            id: "daka"
          }, [
            vue.createElementVNode("image", { src: "/static/not-done.svg" }),
            vue.createElementVNode("text", { class: "title" }, "打卡计划:"),
            vue.createElementVNode("text", { class: "state" }, "已过期或未设置")
          ])),
          vue.createCommentVNode(' <view class="card" id="exam">\r\n					<image src="../../static/todo.svg"></image>\r\n					<text class="time">9:40</text>\r\n					<text class="course">语文</text>\r\n					<text class="score">得分：90</text>\r\n				</view> ')
        ])
      ])
    ]);
  }
  const PagesCalendarCalendar = /* @__PURE__ */ _export_sfc(_sfc_main$d, [["render", _sfc_render$d], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/Calendar/Calendar.vue"]]);
  const _sfc_main$c = {
    data() {
      return {
        examId: 1,
        examName: "第一单元第一次小测",
        examDate: "1970-01-01",
        questionNum: 20,
        correctNum: 19,
        score: 95,
        searchQuery: "",
        questions: [
          // ...题目数组
          {
            id: 1,
            sentences: [
              "-is your brother?",
              "-He is a doctor."
            ],
            options: {
              A: "What",
              B: "Who",
              C: "Where",
              D: "How"
            },
            correctAnswer: "B",
            correctPoints: 0,
            totalPoints: 5,
            userAnswer: "A"
          },
          {
            id: 2,
            sentences: [
              "-______is the nearest post office,please?",
              "-It's about half an hour's walk from here？"
            ],
            options: {
              A: "How far",
              B: "How long",
              C: "How often",
              D: "How soon"
            },
            correctAnswer: "A",
            correctPoints: 5,
            totalPoints: 5,
            userAnswer: "A"
          },
          {
            id: 3,
            sentences: [
              "-is your brother?",
              "-He is a doctor."
            ],
            options: {
              A: "What",
              B: "Who",
              C: "Where",
              D: "How"
            },
            correctAnswer: "B",
            correctPoints: 0,
            totalPoints: 5,
            userAnswer: "A"
          },
          {
            id: 4,
            sentences: [
              "-______is the nearest post office,please?",
              "-It's about half an hour's walk from here？"
            ],
            options: {
              A: "How far",
              B: "How long",
              C: "How often",
              D: "How soon"
            },
            correctAnswer: "A",
            correctPoints: 5,
            totalPoints: 5,
            userAnswer: "A"
          }
        ]
      };
    },
    computed: {
      filteredQuestions() {
        if (this.searchQuery) {
          return this.questions.filter((question) => question.sentence[1].includes(this.searchQuery));
        }
        return this.questions;
      }
    },
    onLoad(options) {
      this.examId = options.exam_id;
      this.examName = options.exam_name;
      uni.request({
        url: "/api/exams/examination_details",
        method: "POST",
        data: {
          exam_id: this.examId
        },
        header: {
          "content-type": "application/json",
          "Authorization": `Bearer ${uni.getStorageSync("token")}`
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
          this.examDate = res.data.exam_date;
          this.questionNum = res.data.question_num;
          this.correctNum = res.data.correct_num;
          this.score = res.data.score;
          this.questions = this.transformQuestions(res.data.questions);
        },
        fail: (res) => {
          uni.showToast({
            title: "获取考试详情失败",
            icon: "none"
          });
        }
      });
    },
    methods: {
      // 转换函数
      transformQuestions(questions) {
        return questions.map((question) => {
          return {
            id: question.question_id,
            // 使用原始的question_id作为id
            index: question.question_index,
            // 添加index属性
            sentences: question.question_decription.split("\n"),
            // 将题目描述按换行符分割成句子数组
            options: {
              A: question.choices["A"],
              B: question.choices["B"],
              C: question.choices["C"],
              D: question.choices["D"]
            },
            correctAnswer: question.correct_answer,
            // 正确答案
            correctPoints: question.score,
            // 正确得分
            totalPoints: question.full_score,
            // 总分
            userAnswer: question.my_answer
            // 用户答案
          };
        });
      },
      formatDate(date) {
        return date.replace("-", "年").replace("-", "月").replace("-", "日");
      },
      sortQuestions(by) {
        this.questions.sort((a, b) => {
          if (by === "order") {
            return a.id - b.id;
          } else if (by === "reverse") {
            return b.id - a.id;
          } else if (by === "score") {
            return b.correctPoints - a.correctPoints;
          } else if (by === "scoreReverse") {
            return a.correctPoints - b.correctPoints;
          }
        });
      },
      showAnalysis(questionId, questionIndex) {
        uni.navigateTo({
          url: "/pages/questionDetail/questionDetail?questionId=" + questionId + "&questionIndex=" + questionIndex + "&questionNum=" + this.questionNum
        });
      }
    }
  };
  function _sfc_render$c(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("view", { class: "container" }, [
        vue.createElementVNode("view", { class: "left-container" }, [
          vue.createElementVNode(
            "h3",
            { class: "exam-number" },
            vue.toDisplayString($data.examName),
            1
            /* TEXT */
          ),
          vue.createElementVNode(
            "p",
            { class: "date" },
            vue.toDisplayString($options.formatDate($data.examDate)),
            1
            /* TEXT */
          ),
          vue.createElementVNode("view", { style: { "display": "flex", "justify-content": "flex-start", "margin-top": "25px", "margin-bottom": "20px" } }, [
            vue.createElementVNode(
              "p",
              { class: "question-number" },
              "共" + vue.toDisplayString($data.questionNum) + "题",
              1
              /* TEXT */
            ),
            vue.createElementVNode(
              "p",
              { class: "correct-number" },
              "答对" + vue.toDisplayString($data.correctNum) + "/" + vue.toDisplayString($data.questionNum) + "题",
              1
              /* TEXT */
            )
          ])
        ]),
        vue.createElementVNode("view", { class: "right-container" }, [
          vue.createElementVNode(
            "p",
            { class: "point" },
            vue.toDisplayString($data.score),
            1
            /* TEXT */
          ),
          vue.createElementVNode("p", { class: "small-text" }, "分")
        ])
      ]),
      vue.createElementVNode("view", { class: "container2" }, [
        vue.createElementVNode("view", { class: "header-container" }, [
          vue.createElementVNode("h3", { style: { "font-size": "24px" } }, "考试题目"),
          vue.createElementVNode("view", { class: "search" }, [
            vue.createElementVNode("img", {
              class: "search-icon",
              src: "/static/search.svg"
            }),
            vue.withDirectives(vue.createElementVNode(
              "input",
              {
                "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => $data.searchQuery = $event),
                placeholder: "搜索"
              },
              null,
              512
              /* NEED_PATCH */
            ), [
              [vue.vModelText, $data.searchQuery]
            ])
          ])
        ]),
        vue.createElementVNode("view", { class: "button-group" }, [
          vue.createElementVNode("button", {
            class: "button1",
            onClick: _cache[1] || (_cache[1] = ($event) => $options.sortQuestions("order"))
          }, "题目顺序"),
          vue.createElementVNode("button", {
            class: "button2",
            onClick: _cache[2] || (_cache[2] = ($event) => $options.sortQuestions("reverse"))
          }, "题目逆序"),
          vue.createElementVNode("button", {
            class: "button3",
            onClick: _cache[3] || (_cache[3] = ($event) => $options.sortQuestions("score"))
          }, "分数顺序"),
          vue.createElementVNode("button", {
            class: "button4",
            onClick: _cache[4] || (_cache[4] = ($event) => $options.sortQuestions("scoreReverse"))
          }, "分数逆序")
        ])
      ]),
      (vue.openBlock(true), vue.createElementBlock(
        vue.Fragment,
        null,
        vue.renderList($options.filteredQuestions, (question, index) => {
          return vue.openBlock(), vue.createElementBlock("view", {
            key: question.id,
            class: "title-container"
          }, [
            vue.createElementVNode("view", { style: { "display": "flex", "justify-content": "space-between", "width": "90%" } }, [
              vue.createElementVNode(
                "p",
                { class: "title-number" },
                vue.toDisplayString(index + 1) + ".",
                1
                /* TEXT */
              ),
              vue.createElementVNode(
                "p",
                {
                  class: vue.normalizeClass(question.correctPoints >= question.totalPoints ? "correct-title-point" : "wrong-title-point")
                },
                vue.toDisplayString(question.correctPoints) + "/" + vue.toDisplayString(question.totalPoints),
                3
                /* TEXT, CLASS */
              )
            ]),
            (vue.openBlock(true), vue.createElementBlock(
              vue.Fragment,
              null,
              vue.renderList(question.sentences, (sentence, sentenceIndex) => {
                return vue.openBlock(), vue.createElementBlock("view", {
                  key: sentenceIndex,
                  class: "sentence-group"
                }, [
                  vue.createElementVNode(
                    "p",
                    {
                      class: vue.normalizeClass(sentence)
                    },
                    vue.toDisplayString(sentence),
                    3
                    /* TEXT, CLASS */
                  )
                ]);
              }),
              128
              /* KEYED_FRAGMENT */
            )),
            vue.createCommentVNode(" 使用v-for遍历options对象 "),
            vue.createElementVNode("view", { class: "options-group" }, [
              (vue.openBlock(true), vue.createElementBlock(
                vue.Fragment,
                null,
                vue.renderList(question.options, (optionContent, optionLabel) => {
                  return vue.openBlock(), vue.createElementBlock(
                    "p",
                    {
                      key: optionLabel,
                      class: "option"
                    },
                    vue.toDisplayString(optionLabel) + ": " + vue.toDisplayString(optionContent),
                    1
                    /* TEXT */
                  );
                }),
                128
                /* KEYED_FRAGMENT */
              ))
            ]),
            vue.createElementVNode("view", { style: { "display": "flex", "justify-content": "space-between", "width": "90%", "margin-top": "20px" } }, [
              vue.createElementVNode(
                "p",
                {
                  class: vue.normalizeClass(question.correctPoints >= question.totalPoints ? "correct-answer" : "wrong-answer")
                },
                "我的答案:" + vue.toDisplayString(question.userAnswer),
                3
                /* TEXT, CLASS */
              ),
              vue.createElementVNode(
                "p",
                { class: "true-answer" },
                "正确答案:" + vue.toDisplayString(question.correctAnswer),
                1
                /* TEXT */
              )
            ]),
            vue.createElementVNode("button", {
              class: "big-button",
              onClick: ($event) => $options.showAnalysis(question.id, index)
            }, "查看解析", 8, ["onClick"])
          ]);
        }),
        128
        /* KEYED_FRAGMENT */
      ))
    ]);
  }
  const PagesExam_detailsExam_details = /* @__PURE__ */ _export_sfc(_sfc_main$c, [["render", _sfc_render$c], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/exam_details/exam_details.vue"]]);
  const isObject = (val) => val !== null && typeof val === "object";
  const defaultDelimiters = ["{", "}"];
  class BaseFormatter {
    constructor() {
      this._caches = /* @__PURE__ */ Object.create(null);
    }
    interpolate(message, values, delimiters = defaultDelimiters) {
      if (!values) {
        return [message];
      }
      let tokens = this._caches[message];
      if (!tokens) {
        tokens = parse(message, delimiters);
        this._caches[message] = tokens;
      }
      return compile(tokens, values);
    }
  }
  const RE_TOKEN_LIST_VALUE = /^(?:\d)+/;
  const RE_TOKEN_NAMED_VALUE = /^(?:\w)+/;
  function parse(format, [startDelimiter, endDelimiter]) {
    const tokens = [];
    let position = 0;
    let text = "";
    while (position < format.length) {
      let char = format[position++];
      if (char === startDelimiter) {
        if (text) {
          tokens.push({ type: "text", value: text });
        }
        text = "";
        let sub = "";
        char = format[position++];
        while (char !== void 0 && char !== endDelimiter) {
          sub += char;
          char = format[position++];
        }
        const isClosed = char === endDelimiter;
        const type = RE_TOKEN_LIST_VALUE.test(sub) ? "list" : isClosed && RE_TOKEN_NAMED_VALUE.test(sub) ? "named" : "unknown";
        tokens.push({ value: sub, type });
      } else {
        text += char;
      }
    }
    text && tokens.push({ type: "text", value: text });
    return tokens;
  }
  function compile(tokens, values) {
    const compiled = [];
    let index = 0;
    const mode = Array.isArray(values) ? "list" : isObject(values) ? "named" : "unknown";
    if (mode === "unknown") {
      return compiled;
    }
    while (index < tokens.length) {
      const token = tokens[index];
      switch (token.type) {
        case "text":
          compiled.push(token.value);
          break;
        case "list":
          compiled.push(values[parseInt(token.value, 10)]);
          break;
        case "named":
          if (mode === "named") {
            compiled.push(values[token.value]);
          } else {
            {
              console.warn(`Type of token '${token.type}' and format of value '${mode}' don't match!`);
            }
          }
          break;
        case "unknown":
          {
            console.warn(`Detect 'unknown' type of token!`);
          }
          break;
      }
      index++;
    }
    return compiled;
  }
  const LOCALE_ZH_HANS = "zh-Hans";
  const LOCALE_ZH_HANT = "zh-Hant";
  const LOCALE_EN = "en";
  const LOCALE_FR = "fr";
  const LOCALE_ES = "es";
  const hasOwnProperty = Object.prototype.hasOwnProperty;
  const hasOwn = (val, key) => hasOwnProperty.call(val, key);
  const defaultFormatter = new BaseFormatter();
  function include(str, parts) {
    return !!parts.find((part) => str.indexOf(part) !== -1);
  }
  function startsWith(str, parts) {
    return parts.find((part) => str.indexOf(part) === 0);
  }
  function normalizeLocale(locale, messages2) {
    if (!locale) {
      return;
    }
    locale = locale.trim().replace(/_/g, "-");
    if (messages2 && messages2[locale]) {
      return locale;
    }
    locale = locale.toLowerCase();
    if (locale === "chinese") {
      return LOCALE_ZH_HANS;
    }
    if (locale.indexOf("zh") === 0) {
      if (locale.indexOf("-hans") > -1) {
        return LOCALE_ZH_HANS;
      }
      if (locale.indexOf("-hant") > -1) {
        return LOCALE_ZH_HANT;
      }
      if (include(locale, ["-tw", "-hk", "-mo", "-cht"])) {
        return LOCALE_ZH_HANT;
      }
      return LOCALE_ZH_HANS;
    }
    let locales = [LOCALE_EN, LOCALE_FR, LOCALE_ES];
    if (messages2 && Object.keys(messages2).length > 0) {
      locales = Object.keys(messages2);
    }
    const lang = startsWith(locale, locales);
    if (lang) {
      return lang;
    }
  }
  class I18n {
    constructor({ locale, fallbackLocale, messages: messages2, watcher, formater }) {
      this.locale = LOCALE_EN;
      this.fallbackLocale = LOCALE_EN;
      this.message = {};
      this.messages = {};
      this.watchers = [];
      if (fallbackLocale) {
        this.fallbackLocale = fallbackLocale;
      }
      this.formater = formater || defaultFormatter;
      this.messages = messages2 || {};
      this.setLocale(locale || LOCALE_EN);
      if (watcher) {
        this.watchLocale(watcher);
      }
    }
    setLocale(locale) {
      const oldLocale = this.locale;
      this.locale = normalizeLocale(locale, this.messages) || this.fallbackLocale;
      if (!this.messages[this.locale]) {
        this.messages[this.locale] = {};
      }
      this.message = this.messages[this.locale];
      if (oldLocale !== this.locale) {
        this.watchers.forEach((watcher) => {
          watcher(this.locale, oldLocale);
        });
      }
    }
    getLocale() {
      return this.locale;
    }
    watchLocale(fn) {
      const index = this.watchers.push(fn) - 1;
      return () => {
        this.watchers.splice(index, 1);
      };
    }
    add(locale, message, override = true) {
      const curMessages = this.messages[locale];
      if (curMessages) {
        if (override) {
          Object.assign(curMessages, message);
        } else {
          Object.keys(message).forEach((key) => {
            if (!hasOwn(curMessages, key)) {
              curMessages[key] = message[key];
            }
          });
        }
      } else {
        this.messages[locale] = message;
      }
    }
    f(message, values, delimiters) {
      return this.formater.interpolate(message, values, delimiters).join("");
    }
    t(key, locale, values) {
      let message = this.message;
      if (typeof locale === "string") {
        locale = normalizeLocale(locale, this.messages);
        locale && (message = this.messages[locale]);
      } else {
        values = locale;
      }
      if (!hasOwn(message, key)) {
        console.warn(`Cannot translate the value of keypath ${key}. Use the value of keypath as default.`);
        return key;
      }
      return this.formater.interpolate(message[key], values).join("");
    }
  }
  function watchAppLocale(appVm, i18n) {
    if (appVm.$watchLocale) {
      appVm.$watchLocale((newLocale) => {
        i18n.setLocale(newLocale);
      });
    } else {
      appVm.$watch(() => appVm.$locale, (newLocale) => {
        i18n.setLocale(newLocale);
      });
    }
  }
  function getDefaultLocale() {
    if (typeof uni !== "undefined" && uni.getLocale) {
      return uni.getLocale();
    }
    if (typeof global !== "undefined" && global.getLocale) {
      return global.getLocale();
    }
    return LOCALE_EN;
  }
  function initVueI18n(locale, messages2 = {}, fallbackLocale, watcher) {
    if (typeof locale !== "string") {
      [locale, messages2] = [
        messages2,
        locale
      ];
    }
    if (typeof locale !== "string") {
      locale = getDefaultLocale();
    }
    if (typeof fallbackLocale !== "string") {
      fallbackLocale = typeof __uniConfig !== "undefined" && __uniConfig.fallbackLocale || LOCALE_EN;
    }
    const i18n = new I18n({
      locale,
      fallbackLocale,
      messages: messages2,
      watcher
    });
    let t2 = (key, values) => {
      if (typeof getApp !== "function") {
        t2 = function(key2, values2) {
          return i18n.t(key2, values2);
        };
      } else {
        let isWatchedAppLocale = false;
        t2 = function(key2, values2) {
          const appVm = getApp().$vm;
          if (appVm) {
            appVm.$locale;
            if (!isWatchedAppLocale) {
              isWatchedAppLocale = true;
              watchAppLocale(appVm, i18n);
            }
          }
          return i18n.t(key2, values2);
        };
      }
      return t2(key, values);
    };
    return {
      i18n,
      f(message, values, delimiters) {
        return i18n.f(message, values, delimiters);
      },
      t(key, values) {
        return t2(key, values);
      },
      add(locale2, message, override = true) {
        return i18n.add(locale2, message, override);
      },
      watch(fn) {
        return i18n.watchLocale(fn);
      },
      getLocale() {
        return i18n.getLocale();
      },
      setLocale(newLocale) {
        return i18n.setLocale(newLocale);
      }
    };
  }
  const en = {
    "uni-countdown.day": "day",
    "uni-countdown.h": "h",
    "uni-countdown.m": "m",
    "uni-countdown.s": "s"
  };
  const zhHans = {
    "uni-countdown.day": "天",
    "uni-countdown.h": "时",
    "uni-countdown.m": "分",
    "uni-countdown.s": "秒"
  };
  const zhHant = {
    "uni-countdown.day": "天",
    "uni-countdown.h": "時",
    "uni-countdown.m": "分",
    "uni-countdown.s": "秒"
  };
  const messages = {
    en,
    "zh-Hans": zhHans,
    "zh-Hant": zhHant
  };
  const {
    t
  } = initVueI18n(messages);
  const _sfc_main$b = {
    name: "UniCountdown",
    emits: ["timeup"],
    props: {
      showDay: {
        type: Boolean,
        default: true
      },
      showHour: {
        type: Boolean,
        default: true
      },
      showMinute: {
        type: Boolean,
        default: true
      },
      showColon: {
        type: Boolean,
        default: true
      },
      start: {
        type: Boolean,
        default: true
      },
      backgroundColor: {
        type: String,
        default: ""
      },
      color: {
        type: String,
        default: "#333"
      },
      fontSize: {
        type: Number,
        default: 14
      },
      splitorColor: {
        type: String,
        default: "#333"
      },
      day: {
        type: Number,
        default: 0
      },
      hour: {
        type: Number,
        default: 0
      },
      minute: {
        type: Number,
        default: 0
      },
      second: {
        type: Number,
        default: 0
      },
      timestamp: {
        type: Number,
        default: 0
      }
    },
    data() {
      return {
        timer: null,
        syncFlag: false,
        d: "00",
        h: "00",
        i: "00",
        s: "00",
        leftTime: 0,
        seconds: 0
      };
    },
    computed: {
      dayText() {
        return t("uni-countdown.day");
      },
      hourText(val) {
        return t("uni-countdown.h");
      },
      minuteText(val) {
        return t("uni-countdown.m");
      },
      secondText(val) {
        return t("uni-countdown.s");
      },
      timeStyle() {
        const {
          color,
          backgroundColor,
          fontSize
        } = this;
        return {
          color,
          backgroundColor,
          fontSize: `${fontSize}px`,
          width: `${fontSize * 22 / 14}px`,
          // 按字体大小为 14px 时的比例缩放
          lineHeight: `${fontSize * 20 / 14}px`,
          borderRadius: `${fontSize * 3 / 14}px`
        };
      },
      splitorStyle() {
        const { splitorColor, fontSize, backgroundColor } = this;
        return {
          color: splitorColor,
          fontSize: `${fontSize * 12 / 14}px`,
          margin: backgroundColor ? `${fontSize * 4 / 14}px` : ""
        };
      }
    },
    watch: {
      day(val) {
        this.changeFlag();
      },
      hour(val) {
        this.changeFlag();
      },
      minute(val) {
        this.changeFlag();
      },
      second(val) {
        this.changeFlag();
      },
      start: {
        immediate: true,
        handler(newVal, oldVal) {
          if (newVal) {
            this.startData();
          } else {
            if (!oldVal)
              return;
            clearInterval(this.timer);
          }
        }
      }
    },
    created: function(e) {
      this.seconds = this.toSeconds(this.timestamp, this.day, this.hour, this.minute, this.second);
      this.countDown();
    },
    unmounted() {
      clearInterval(this.timer);
    },
    methods: {
      toSeconds(timestamp, day, hours, minutes, seconds) {
        if (timestamp) {
          return timestamp - parseInt((/* @__PURE__ */ new Date()).getTime() / 1e3, 10);
        }
        return day * 60 * 60 * 24 + hours * 60 * 60 + minutes * 60 + seconds;
      },
      timeUp() {
        clearInterval(this.timer);
        this.$emit("timeup");
      },
      countDown() {
        let seconds = this.seconds;
        let [day, hour, minute, second] = [0, 0, 0, 0];
        if (seconds > 0) {
          day = Math.floor(seconds / (60 * 60 * 24));
          hour = Math.floor(seconds / (60 * 60)) - day * 24;
          minute = Math.floor(seconds / 60) - day * 24 * 60 - hour * 60;
          second = Math.floor(seconds) - day * 24 * 60 * 60 - hour * 60 * 60 - minute * 60;
        } else {
          this.timeUp();
        }
        if (day < 10) {
          day = "0" + day;
        }
        if (hour < 10) {
          hour = "0" + hour;
        }
        if (minute < 10) {
          minute = "0" + minute;
        }
        if (second < 10) {
          second = "0" + second;
        }
        this.d = day;
        this.h = hour;
        this.i = minute;
        this.s = second;
      },
      startData() {
        this.seconds = this.toSeconds(this.timestamp, this.day, this.hour, this.minute, this.second);
        if (this.seconds <= 0) {
          this.seconds = this.toSeconds(0, 0, 0, 0, 0);
          this.countDown();
          return;
        }
        clearInterval(this.timer);
        this.countDown();
        this.timer = setInterval(() => {
          this.seconds--;
          if (this.seconds < 0) {
            this.timeUp();
            return;
          }
          this.countDown();
        }, 1e3);
      },
      update() {
        this.startData();
      },
      changeFlag() {
        if (!this.syncFlag) {
          this.seconds = this.toSeconds(this.timestamp, this.day, this.hour, this.minute, this.second);
          this.startData();
          this.syncFlag = true;
        }
      }
    }
  };
  function _sfc_render$b(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", { class: "uni-countdown" }, [
      $props.showDay ? (vue.openBlock(), vue.createElementBlock(
        "text",
        {
          key: 0,
          style: vue.normalizeStyle([$options.timeStyle]),
          class: "uni-countdown__number"
        },
        vue.toDisplayString($data.d),
        5
        /* TEXT, STYLE */
      )) : vue.createCommentVNode("v-if", true),
      $props.showDay ? (vue.openBlock(), vue.createElementBlock(
        "text",
        {
          key: 1,
          style: vue.normalizeStyle([$options.splitorStyle]),
          class: "uni-countdown__splitor"
        },
        vue.toDisplayString($options.dayText),
        5
        /* TEXT, STYLE */
      )) : vue.createCommentVNode("v-if", true),
      $props.showHour ? (vue.openBlock(), vue.createElementBlock(
        "text",
        {
          key: 2,
          style: vue.normalizeStyle([$options.timeStyle]),
          class: "uni-countdown__number"
        },
        vue.toDisplayString($data.h),
        5
        /* TEXT, STYLE */
      )) : vue.createCommentVNode("v-if", true),
      $props.showHour ? (vue.openBlock(), vue.createElementBlock(
        "text",
        {
          key: 3,
          style: vue.normalizeStyle([$options.splitorStyle]),
          class: "uni-countdown__splitor"
        },
        vue.toDisplayString($props.showColon ? ":" : $options.hourText),
        5
        /* TEXT, STYLE */
      )) : vue.createCommentVNode("v-if", true),
      $props.showMinute ? (vue.openBlock(), vue.createElementBlock(
        "text",
        {
          key: 4,
          style: vue.normalizeStyle([$options.timeStyle]),
          class: "uni-countdown__number"
        },
        vue.toDisplayString($data.i),
        5
        /* TEXT, STYLE */
      )) : vue.createCommentVNode("v-if", true),
      $props.showMinute ? (vue.openBlock(), vue.createElementBlock(
        "text",
        {
          key: 5,
          style: vue.normalizeStyle([$options.splitorStyle]),
          class: "uni-countdown__splitor"
        },
        vue.toDisplayString($props.showColon ? ":" : $options.minuteText),
        5
        /* TEXT, STYLE */
      )) : vue.createCommentVNode("v-if", true),
      vue.createElementVNode(
        "text",
        {
          style: vue.normalizeStyle([$options.timeStyle]),
          class: "uni-countdown__number"
        },
        vue.toDisplayString($data.s),
        5
        /* TEXT, STYLE */
      ),
      !$props.showColon ? (vue.openBlock(), vue.createElementBlock(
        "text",
        {
          key: 6,
          style: vue.normalizeStyle([$options.splitorStyle]),
          class: "uni-countdown__splitor"
        },
        vue.toDisplayString($options.secondText),
        5
        /* TEXT, STYLE */
      )) : vue.createCommentVNode("v-if", true)
    ]);
  }
  const __easycom_0 = /* @__PURE__ */ _export_sfc(_sfc_main$b, [["render", _sfc_render$b], ["__scopeId", "data-v-c592f7f2"], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/uni_modules/uni-countdown/components/uni-countdown/uni-countdown.vue"]]);
  const _sfc_main$a = {
    data() {
      return {
        swiperOptions: {
          // 其他配置...
          allowTouchMove: true,
          // 允许触摸滑动
          preventClicksPropagation: true
          // 阻止点击事件冒泡
          // 其他 Swiper 配置...
        },
        exam_id: null,
        exam_name: null,
        progress: 1,
        // 进度条的初始值
        current: 0,
        // 当前进度
        currentQuestionIndex: 0,
        //当前正在做的题目序号
        selectedIndex: -1,
        // 当前题目的选中按钮序号
        questionButtonIndex: 0,
        // 当前题目的按钮序号
        isShow: false,
        //是否显示全部题目
        currentFillAnswer: null,
        //当前填空题输入的答案
        questions: [
          // 题目和选项
          {
            question_id: 1,
            question_type: 1,
            //1单选2填空
            question: `__ is your brother?
									-He is a doctor.`,
            activeButtonIndex: null,
            // 用于存储当前激活的按钮索引
            choices: ["1", "2", "2", "放弃"],
            fullScore: 5
            // 题目满分
          },
          {
            question_id: 2,
            question_type: 1,
            //1单选2填空
            question: "abandon",
            activeButtonIndex: null,
            // 用于存储当前激活的按钮索引
            choices: ["1", "选项B", "选项C", "选项D"],
            fullScore: 5
            // 题目满分
          },
          {
            question_id: 3,
            question_type: 1,
            //1单选2填空
            question: "abandon2",
            activeButtonIndex: null,
            // 用于存储当前激活的按钮索引
            choices: ["1", "选项B", "选项C", "选项D"],
            fullScore: 5
            // 题目满分
          }
          // ...更多题目
        ],
        // 这里可以根据需要修改选项内容
        realAnswer: [
          "放弃",
          "选项B",
          "选项C"
          // 正确答案
        ],
        maxButtonsPerRow: 5,
        // 每行的最大元素个数
        buttonMargin: 35,
        // 元素间隔
        selectedChoiceAndScore: {
          /*//key为question_id
          1: {
                 selectedChoice: null, // 用于存储当前选择的选项，或者是输入的答案
                 score: 0 // 用于存储当前题目的分数
               },
          2: {
                 selectedChoice: null, // 用于存储当前选择的选项，或者是输入的答案
                 score: 0 // 用于存储当前题目的分数
               },
          3: {
                 selectedChoice: null, // 用于存储当前选择的选项，或者是输入的答案
                 score: 0 // 用于存储当前题目的分数
               },*/
        },
        isFinished: {
          1: true,
          2: true,
          3: true
        },
        // 是否完成答题
        hasShownSubmitPrompt: false,
        // 是否已显示提交提示
        correctAnswers: 0
        // 正确答案数
      };
    },
    onLoad(event) {
      let exam_id = parseInt(event.exam_id);
      this.exam_name = event.name;
      this.exam_id = exam_id;
      uni.request({
        url: "/api/exams/take_examination",
        method: "POST",
        data: {
          exam_id
        },
        header: {
          "Authorization": `Bearer ${uni.getStorageSync("token")}`
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
          let questionAndAnswer = this.transformQuestions(res.data.question_list);
          this.questions = questionAndAnswer.questions;
          this.realAnswer = questionAndAnswer.realAnswer;
          questionAndAnswer.questions.forEach((question, index) => {
            this.isFinished[question.question_id] = false;
            this.selectedChoiceAndScore[question.question_id] = {
              selectedChoice: null,
              // 用于存储当前选择的选项
              score: 0
              // 用于存储当前题目的分数
            };
          });
        },
        fail: (res) => {
          uni.showToast({
            title: "获取题目失败",
            icon: "none"
          });
        }
      });
    },
    computed: {
      //这是每一行的按钮，其中最多有maxButtonsPerRow个
      rows() {
        const rows = [];
        for (let i = 0; i < this.questions.length; i += this.maxButtonsPerRow) {
          let thisRowQuestions = [];
          for (let j = i; j < i + this.maxButtonsPerRow && j < this.questions.length; j++) {
            thisRowQuestions.push({
              //题目序号
              index: j,
              //题目
              question: this.questions[j]
            });
          }
          rows.push(thisRowQuestions);
        }
        return rows;
      }
    },
    methods: {
      //输入填空题的逻辑
      inputFillAnswer(currentQuestionIndex) {
        let correct = this.currentFillAnswer === this.realAnswer[currentQuestionIndex];
        this.selectedChoiceAndScore[this.questions[currentQuestionIndex].question_id].selectedChoice = this.currentFillAnswer;
        if (correct) {
          this.selectedChoiceAndScore[this.questions[currentQuestionIndex].question_id].score = this.questions[currentQuestionIndex].fullScore;
          this.correctAnswers++;
        } else {
          this.selectedChoiceAndScore[this.questions[currentQuestionIndex].question_id].score = 0;
        }
      },
      nextQuestionForFills(currentQuestionIndex) {
        if (!this.isFinished[this.questions[currentQuestionIndex].question_id]) {
          this.isFinished[this.questions[currentQuestionIndex].question_id] = true;
          this.currentQuestionIndex++;
          this.current++;
        }
        this.selectedIndex = -1;
        this.currentFillAnswer = null;
        if (this.currentQuestionIndex === this.questions.length && !this.hasShownSubmitPrompt) {
          this.currentQuestionIndex--;
          this.hasShownSubmitPrompt = true;
          this.submitExam();
        } else {
          this.swiperChange({
            detail: {
              current: this.currentQuestionIndex,
              source: "touch"
            }
          });
        }
      },
      transformQuestions(questionList) {
        let questions = [];
        let realAnswer = [];
        questionList.forEach((item, index) => {
          questions.push({
            question_id: item.question_id,
            question_type: item.question_type,
            question: item.question_content,
            activeButtonIndex: null,
            // 初始化激活按钮索引
            choices: item.question_choices,
            fullScore: item.full_score
            // 题目满分
          });
          realAnswer.push(item.question_answer);
        });
        return {
          questions,
          realAnswer
        };
      },
      setCurrentQuestionIndexByQuestionId(question_id) {
        formatAppLog("log", "at pages/exam/exam.vue:251", "setCurrentQuestionIndexByQuestionId:" + question_id);
        for (let i = 0; i < this.questions.length; i++) {
          if (this.questions[i].question_id === question_id) {
            this.currentQuestionIndex = i;
            break;
          }
        }
      },
      isAllFinished() {
        let allFinished = true;
        for (let key in this.isFinished) {
          if (!this.isFinished[key]) {
            allFinished = false;
            break;
          }
        }
        return allFinished;
      },
      getProgress() {
        let progress = 0;
        for (let key in this.isFinished) {
          if (this.isFinished[key]) {
            progress++;
          }
        }
        return progress / this.questions.length * 100;
      },
      finishQuestion(index) {
        if (this.selectedIndex == -1) {
          return;
        }
        let selectedChoice = this.questions[index].choices[this.selectedIndex];
        formatAppLog("log", "at pages/exam/exam.vue:284", "第" + index + "题你选择了" + selectedChoice);
        let question_id = this.questions[index].question_id;
        this.selectedChoiceAndScore[question_id].selectedChoice = this.selectedIndex;
        if (this.selectedIndex === this.realAnswer[index]) {
          this.selectedChoiceAndScore[question_id].score = this.questions[index].fullScore;
          this.correctAnswers++;
        } else {
          this.selectedChoiceAndScore[question_id].score = 0;
        }
        if (!this.isFinished[this.questions[index].question_id]) {
          this.isFinished[this.questions[index].question_id] = true;
          this.currentQuestionIndex++;
          this.current++;
        }
        this.selectedIndex = -1;
        this.currentFillAnswer = null;
        if (this.currentQuestionIndex === this.questions.length && !this.hasShownSubmitPrompt) {
          this.currentQuestionIndex--;
          this.hasShownSubmitPrompt = true;
          this.submitExam();
        } else {
          this.swiperChange({
            detail: {
              current: this.currentQuestionIndex,
              source: "touch"
            }
          });
        }
      },
      getTotalScore() {
        let totalScore = 0;
        for (let key in this.selectedChoiceAndScore) {
          totalScore += this.selectedChoiceAndScore[key].score;
        }
        return totalScore;
      },
      submitExam() {
        uni.showModal({
          title: "提示",
          content: this.isAllFinished() ? "您已完成全部题目，是否确认提交" : "您还有题目未完成，是否确认提交",
          showCancel: true,
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
            if (res.confirm) {
              let examResult = {
                exam_id: this.exam_id,
                examTitle: this.exam_name,
                score: this.getTotalScore(),
                //考试总分
                totalQuestions: this.questions.length,
                //总题目数
                correctAnswers: this.correctAnswers
                //正确答案数
              };
              formatAppLog("log", "at pages/exam/exam.vue:348", examResult);
              uni.setStorageSync("examResult", examResult);
              uni.request({
                url: `/api/exams/submitExamResult`,
                method: "POST",
                data: {
                  selectedChoiceAndScore: this.selectedChoiceAndScore,
                  exam_id: this.exam_id
                },
                header: {
                  "Authorization": `Bearer ${uni.getStorageSync("token")}`
                },
                success: (res2) => {
                  uni.showToast({
                    title: "提交成功",
                    icon: "none"
                  });
                  this.handleJump();
                },
                fail: (res2) => {
                  uni.showToast({
                    title: "提交失败",
                    icon: "none"
                  });
                }
              });
            }
          }
        });
      },
      handleJump() {
        uni.navigateTo({
          url: "/pages/finishexam/finishexam?progress=" + this.getProgress()
        });
      },
      swiperChange(event) {
        const current = event.detail.current;
        const source = event.detail.source;
        if (source === "touch") {
          this.currentQuestionIndex = current;
          this.selectedIndex = -1;
        }
      },
      selectChoice(index, currentQuestionIndex) {
        this.selectedIndex = index;
        this.questions[currentQuestionIndex].question_id;
        if (this.questions[currentQuestionIndex].activeButtonIndex === index) {
          this.questions[currentQuestionIndex].activeButtonIndex = null;
        } else {
          this.questions[currentQuestionIndex].activeButtonIndex = index;
        }
        this.finishQuestion(currentQuestionIndex);
      },
      getLabel(choiceIndex) {
        const labels = ["A", "B", "C", "D"];
        return labels[choiceIndex];
      },
      showQuestions() {
        this.isShow = !this.isShow;
      }
    }
  };
  function _sfc_render$a(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_uni_countdown = resolveEasycom(vue.resolveDynamicComponent("uni-countdown"), __easycom_0);
    return vue.openBlock(), vue.createElementBlock("view", { class: "container" }, [
      vue.createElementVNode(
        "text",
        { class: "progress-text" },
        vue.toDisplayString($data.current) + "/" + vue.toDisplayString($data.questions.length),
        1
        /* TEXT */
      ),
      vue.createElementVNode("swiper", {
        class: "question-container",
        options: $data.swiperOptions,
        "easing-function": "linear",
        duration: 500,
        current: $data.currentQuestionIndex,
        onBeforeChange: _cache[3] || (_cache[3] = (...args) => $options.swiperChange && $options.swiperChange(...args))
      }, [
        (vue.openBlock(true), vue.createElementBlock(
          vue.Fragment,
          null,
          vue.renderList($data.questions, (question, index) => {
            return vue.openBlock(), vue.createElementBlock("swiper-item", { key: index }, [
              vue.createElementVNode("view", { class: "text-info" }, [
                vue.createElementVNode(
                  "text",
                  { class: "number" },
                  vue.toDisplayString(index + 1),
                  1
                  /* TEXT */
                ),
                vue.createCommentVNode(" 以上是题目序号 "),
                vue.createElementVNode(
                  "text",
                  { class: "question" },
                  vue.toDisplayString(question.question),
                  1
                  /* TEXT */
                )
              ]),
              vue.createCommentVNode("        如果是单选题，则显示选项按钮，否则显示输入框"),
              vue.withDirectives(vue.createElementVNode(
                "view",
                { class: "button-group" },
                [
                  (vue.openBlock(true), vue.createElementBlock(
                    vue.Fragment,
                    null,
                    vue.renderList(question.choices, (choice, choiceIndex) => {
                      return vue.openBlock(), vue.createElementBlock("div", {
                        key: choiceIndex,
                        class: "choice-container"
                      }, [
                        vue.createElementVNode("button", {
                          class: vue.normalizeClass(["option", { "active": choiceIndex === question.activeButtonIndex }]),
                          onClick: ($event) => $options.selectChoice(choiceIndex, $data.currentQuestionIndex)
                        }, vue.toDisplayString($options.getLabel(choiceIndex)), 11, ["onClick"]),
                        vue.createElementVNode(
                          "span",
                          { class: "choice-content" },
                          vue.toDisplayString(choice),
                          1
                          /* TEXT */
                        )
                      ]);
                    }),
                    128
                    /* KEYED_FRAGMENT */
                  )),
                  vue.createCommentVNode('					<button class="confirm" @click="finishQuestion(index)">确认答案</button>')
                ],
                512
                /* NEED_PATCH */
              ), [
                [vue.vShow, question.question_type == 1]
              ]),
              vue.createCommentVNode("        TODO 填空题的输入框的样式"),
              vue.withDirectives(vue.createElementVNode(
                "view",
                { class: "input-container" },
                [
                  vue.withDirectives(vue.createElementVNode(
                    "input",
                    {
                      type: "text",
                      placeholder: "请输入答案",
                      "onUpdate:modelValue": _cache[0] || (_cache[0] = ($event) => $data.currentFillAnswer = $event),
                      onBlur: _cache[1] || (_cache[1] = ($event) => $options.inputFillAnswer($data.currentQuestionIndex)),
                      onConfirm: _cache[2] || (_cache[2] = ($event) => $options.nextQuestionForFills($data.currentQuestionIndex))
                    },
                    null,
                    544
                    /* NEED_HYDRATION, NEED_PATCH */
                  ), [
                    [vue.vModelText, $data.currentFillAnswer]
                  ])
                ],
                512
                /* NEED_PATCH */
              ), [
                [vue.vShow, question.question_type == 2]
              ])
            ]);
          }),
          128
          /* KEYED_FRAGMENT */
        ))
      ], 40, ["options", "current"]),
      vue.createElementVNode("view", { class: "footer" }, [
        vue.createElementVNode("view", { style: { "display": "flex", "white-space": "nowrap" } }, [
          vue.createVNode(_component_uni_countdown, {
            class: "daojishi",
            "show-day": false,
            hour: 12,
            minute: 12,
            second: 12,
            "font-size": 20
          }),
          vue.createElementVNode("image", {
            src: "/static/xuanxiang.svg",
            class: "xuanxiangbtn",
            onClick: _cache[4] || (_cache[4] = (...args) => $options.showQuestions && $options.showQuestions(...args))
          })
        ]),
        vue.withDirectives(vue.createElementVNode(
          "view",
          { class: "xuanxiang-container" },
          [
            (vue.openBlock(true), vue.createElementBlock(
              vue.Fragment,
              null,
              vue.renderList($options.rows, (thisRowQuestions, rowIndex) => {
                return vue.openBlock(), vue.createElementBlock("view", {
                  key: rowIndex,
                  class: "row"
                }, [
                  (vue.openBlock(true), vue.createElementBlock(
                    vue.Fragment,
                    null,
                    vue.renderList(thisRowQuestions, (thisRowQuestion, index) => {
                      return vue.openBlock(), vue.createElementBlock("button", {
                        key: index,
                        class: vue.normalizeClass(["option", { "finished": $data.isFinished[thisRowQuestion.question.question_id] }]),
                        onClick: ($event) => $options.setCurrentQuestionIndexByQuestionId(thisRowQuestion.question.question_id),
                        style: vue.normalizeStyle({ margin: $data.buttonMargin + "rpx" })
                      }, vue.toDisplayString(thisRowQuestion.index + 1), 15, ["onClick"]);
                    }),
                    128
                    /* KEYED_FRAGMENT */
                  ))
                ]);
              }),
              128
              /* KEYED_FRAGMENT */
            )),
            vue.createElementVNode("button", {
              class: "submit",
              onClick: _cache[5] || (_cache[5] = (...args) => $options.submitExam && $options.submitExam(...args))
            }, "直接交卷")
          ],
          512
          /* NEED_PATCH */
        ), [
          [vue.vShow, $data.isShow]
        ])
      ])
    ]);
  }
  const PagesExamExam = /* @__PURE__ */ _export_sfc(_sfc_main$a, [["render", _sfc_render$a], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/exam/exam.vue"]]);
  const _imports_0 = "/static/score1.svg";
  const _imports_1 = "/static/score2.svg";
  const _imports_2 = "/static/score3.png";
  const _sfc_main$9 = {
    data() {
      return {
        exams: [
          {
            exam_id: 1,
            name: "第一单元第一次小测",
            start_time: "20:00",
            duration: 60,
            info: "共20题",
            questionNum: 20
          },
          {
            exam_id: 2,
            name: "第一单元第一次小测",
            start_time: "20:00",
            duration: 60,
            info: "共20题",
            questionNum: 20
          }
          // ...更多的考试对象
        ],
        finishedExams: [
          {
            exam_id: 3,
            name: "第一单元第一次小测",
            date: "2023年1月1日",
            info: "共20题",
            questionNum: 20,
            score: 95
          },
          {
            exam_id: 4,
            name: "第一单元第二次小测",
            date: "2023年1月1日",
            info: "共20题",
            questionNum: 20,
            score: 35
          },
          {
            exam_id: 5,
            name: "第二单元第一次小测",
            date: "2023年1月1日",
            info: "共20题",
            questionNum: 20,
            score: 70
          }
          // ...更多的考试结果对象
        ]
      };
    },
    onLoad() {
      this.getTodayExams();
      this.getPreviousExams();
    },
    methods: {
      getDateByString(dateStr) {
        const regex = /(\d+)年(\d+)月(\d+)日/;
        const match = dateStr.match(regex);
        if (!match) {
          return null;
        }
        const year = parseInt(match[1], 10);
        const month = parseInt(match[2], 10) - 1;
        const day = parseInt(match[3], 10);
        return new Date(year, month, day);
      },
      sortExamsBy(param) {
        this.finishedExams.sort((a, b) => {
          if (param === "date")
            return this.getDateByString(a.date) - this.getDateByString(b.date);
          else if (param === "score")
            return b.score - a.score;
        });
      },
      sortExamsReverse(param) {
        this.finishedExams.sort((a, b) => {
          if (param === "date")
            return this.getDateByString(b.date) - this.getDateByString(a.date);
          else if (param === "score")
            return a.score - b.score;
        });
      },
      formatTimeRange(start_time, duration) {
        const [startHour, startMinute] = start_time.split(":").map(Number);
        const endMinute = (startMinute + duration) % 60;
        const endHour = startHour + Math.floor((startMinute + duration) / 60);
        const endTime = `${endHour.toString().padStart(2, "0")}:${endMinute.toString().padStart(2, "0")}`;
        return `${start_time} ~${endTime}`;
      },
      takeExam(exam) {
        uni.setStorageSync("startExam", JSON.stringify(exam));
        uni.navigateTo({
          url: `../startexam/startexam`
        });
      },
      //由date类型转为类似于'2022-01-03'字符串类型
      getExamDate(date) {
        const year = date.getFullYear();
        const month = date.getMonth() + 1;
        const day = date.getDate();
        const formattedMonth = month < 10 ? "0" + month : month;
        const formattedDay = day < 10 ? "0" + day : day;
        return `${year}-${formattedMonth}-${formattedDay}`;
      },
      getTodayExams() {
        uni.request({
          url: "/api/exams/exam_date",
          method: "POST",
          header: {
            "Authorization": `Bearer ${uni.getStorageSync("token")}`
          },
          data: {
            date: this.getExamDate(/* @__PURE__ */ new Date())
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
            if (res.data.code == 200) {
              this.exams = this.transformExams(res.data.exams);
            }
          },
          fail: (err) => {
            formatAppLog("log", "at pages/ExamHistory/ExamHistory.vue:191", err);
          }
        });
      },
      getPreviousExams() {
        uni.request({
          url: "/api/exams/previous_examinations",
          method: "GET",
          header: {
            "Authorization": `Bearer ${uni.getStorageSync("token")}`
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
            this.finishedExams = this.transformExams(res.data.exams);
          }
        });
      },
      viewDetails(exam) {
        uni.navigateTo({
          url: `../exam_details/exam_details?exam_id=${exam.exam_id}&exam_name=${exam.name}`
        });
      },
      transformExams(exams) {
        return exams.map((exam) => {
          const dateParts = exam.exam_date.split("-");
          const dateInChineseFormat = `${dateParts[0]}年${dateParts[1]}月${dateParts[2]}日`;
          return {
            exam_id: exam.exam_id,
            name: exam.exam_name,
            date: dateInChineseFormat,
            info: `共${exam.question_num}题`,
            questionNum: exam.question_num,
            score: exam.exam_score
          };
        });
      }
    }
  };
  function _sfc_render$9(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("view", { class: "today-container" }, [
        vue.createElementVNode("span", { class: "title" }, "今日考试"),
        $data.exams.length === 0 ? (vue.openBlock(), vue.createElementBlock("span", {
          key: 0,
          class: "no-exam"
        }, [
          vue.createTextVNode("今日暂无考试，"),
          vue.createElementVNode("navigator", null, "前往复习")
        ])) : vue.createCommentVNode("v-if", true),
        (vue.openBlock(true), vue.createElementBlock(
          vue.Fragment,
          null,
          vue.renderList($data.exams, (exam) => {
            return vue.openBlock(), vue.createElementBlock("view", {
              class: "todo-exam",
              key: exam.name
            }, [
              vue.createElementVNode("view", { class: "row1" }, [
                vue.createElementVNode(
                  "text",
                  { class: "exam-name" },
                  vue.toDisplayString(exam.name),
                  1
                  /* TEXT */
                ),
                vue.createElementVNode(
                  "text",
                  { class: "exam-time" },
                  vue.toDisplayString($options.formatTimeRange(exam.start_time, exam.duration)),
                  1
                  /* TEXT */
                ),
                vue.createElementVNode(
                  "text",
                  { class: "exam-info" },
                  vue.toDisplayString(exam.info),
                  1
                  /* TEXT */
                )
              ]),
              vue.createElementVNode("view", { class: "row2" }, [
                vue.createElementVNode("button", {
                  class: "todo-btn",
                  onClick: ($event) => _ctx.reminder(exam)
                }, "提醒我", 8, ["onClick"]),
                vue.createElementVNode("button", {
                  class: "todo-btn",
                  onClick: ($event) => $options.takeExam(exam)
                }, "去考试", 8, ["onClick"])
              ])
            ]);
          }),
          128
          /* KEYED_FRAGMENT */
        ))
      ]),
      vue.createElementVNode("view", { class: "history-container" }, [
        vue.createElementVNode("view", { class: "_row1" }, [
          vue.createElementVNode("span", { class: "title" }, "所有考试"),
          vue.createElementVNode("input")
        ]),
        vue.createElementVNode("view", { class: "_row2" }, [
          vue.createElementVNode("button", {
            class: "choice selected",
            onClick: _cache[0] || (_cache[0] = ($event) => $options.sortExamsBy("date"))
          }, "时间顺序"),
          vue.createElementVNode("button", {
            class: "choice",
            onClick: _cache[1] || (_cache[1] = ($event) => $options.sortExamsReverse("date"))
          }, "时间逆序"),
          vue.createElementVNode("button", {
            class: "choice",
            onClick: _cache[2] || (_cache[2] = ($event) => $options.sortExamsBy("score"))
          }, "成绩顺序"),
          vue.createElementVNode("button", {
            class: "choice",
            onClick: _cache[3] || (_cache[3] = ($event) => $options.sortExamsReverse("score"))
          }, "成绩逆序")
        ]),
        (vue.openBlock(true), vue.createElementBlock(
          vue.Fragment,
          null,
          vue.renderList($data.finishedExams, (exam) => {
            return vue.openBlock(), vue.createElementBlock("view", {
              class: "finished-exam",
              key: exam.date
            }, [
              exam.score >= 80 ? (vue.openBlock(), vue.createElementBlock("image", {
                key: 0,
                class: "level",
                src: _imports_0
              })) : exam.score >= 60 && exam.score < 80 ? (vue.openBlock(), vue.createElementBlock("image", {
                key: 1,
                class: "level",
                src: _imports_1
              })) : (vue.openBlock(), vue.createElementBlock("image", {
                key: 2,
                class: "level",
                src: _imports_2
              })),
              vue.createCommentVNode(' <image class="level" src="@/static/score1.svg" v-else></image> '),
              vue.createElementVNode("view", { class: "row1" }, [
                vue.createElementVNode(
                  "text",
                  { class: "exam-name" },
                  vue.toDisplayString(exam.name),
                  1
                  /* TEXT */
                ),
                vue.createElementVNode(
                  "text",
                  { class: "exam-date" },
                  vue.toDisplayString(exam.date),
                  1
                  /* TEXT */
                ),
                vue.createElementVNode(
                  "text",
                  { class: "exam-info" },
                  vue.toDisplayString(exam.info),
                  1
                  /* TEXT */
                )
              ]),
              vue.createElementVNode("view", { class: "row22" }, [
                vue.createElementVNode("view", null, [
                  vue.createElementVNode(
                    "span",
                    { class: "score" },
                    vue.toDisplayString(exam.score),
                    1
                    /* TEXT */
                  ),
                  vue.createElementVNode("span", { style: { "margin-left": "8rpx", "font-size": "24px" } }, "分")
                ]),
                vue.createElementVNode("button", {
                  class: "todetail-btn",
                  onClick: ($event) => $options.viewDetails(exam)
                }, "考试详情", 8, ["onClick"])
              ])
            ]);
          }),
          128
          /* KEYED_FRAGMENT */
        ))
      ])
    ]);
  }
  const PagesExamHistoryExamHistory = /* @__PURE__ */ _export_sfc(_sfc_main$9, [["render", _sfc_render$9], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/ExamHistory/ExamHistory.vue"]]);
  const _sfc_main$8 = {
    data() {
      return {
        screenWidth: 0
        // 屏幕宽度
      };
    },
    props: {
      progress: {
        type: Number,
        required: true,
        validator(value) {
          return value >= 0 && value <= 100;
        }
      },
      backgroundColor: {
        type: String,
        default: "#EFEFF4"
      },
      progressBackgroundColor: {
        type: String,
        default: "#07C160"
      },
      showText: {
        type: Boolean,
        default: false
      },
      textColor: {
        type: String,
        default: "#000000"
      },
      textSize: {
        type: Number,
        default: 28
      },
      height: {
        type: Number,
        default: 20
      },
      isCircular: {
        type: Boolean,
        default: false
      },
      diameter: {
        type: Number,
        default: 100
      },
      canvasId: {
        type: String,
        default: "canvasId"
      }
    },
    mounted() {
      uni.getSystemInfo({
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
          this.screenWidth = res.screenWidth;
          if (this.isCircular) {
            this.drawCircularProgress();
          }
        }
      });
    },
    watch: {
      progress: function(val) {
        formatAppLog("log", "at uni_modules/piaoyi-progress-bar/components/piaoyi-progress-bar/piaoyi-progress-bar.vue:87", val);
        if (this.isCircular) {
          this.drawCircularProgress();
        }
      }
    },
    methods: {
      drawCircularProgress() {
        const canvas = uni.createCanvasContext(this.canvasId, this);
        const radius = (this.rpxToPx(this.diameter) - this.rpxToPx(this.height)) / 2;
        const startAngle = -Math.PI / 2;
        const endAngle = 2 * Math.PI * this.progress / 100 + startAngle;
        canvas.setLineWidth(this.rpxToPx(this.height));
        canvas.setStrokeStyle(this.backgroundColor);
        canvas.beginPath();
        canvas.arc(this.rpxToPx(this.diameter) / 2, this.rpxToPx(this.diameter) / 2, radius, 0, 2 * Math.PI);
        canvas.stroke();
        canvas.setLineWidth(this.rpxToPx(this.height));
        canvas.setStrokeStyle(this.progressBackgroundColor);
        canvas.beginPath();
        canvas.arc(
          this.rpxToPx(this.diameter) / 2,
          this.rpxToPx(this.diameter) / 2,
          radius,
          startAngle,
          endAngle,
          false
        );
        canvas.stroke();
        canvas.draw();
      },
      rpxToPx(rpx) {
        const px = rpx / 750 * this.screenWidth;
        return px;
      }
    }
  };
  function _sfc_render$8(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock(
      "view",
      {
        class: "progress-bar",
        style: vue.normalizeStyle("min-height:" + $props.height + "rpx")
      },
      [
        !$props.isCircular ? (vue.openBlock(), vue.createElementBlock(
          "view",
          {
            key: 0,
            class: "progress-bar-bg",
            style: vue.normalizeStyle({ backgroundColor: $props.backgroundColor, height: $props.height + "rpx" })
          },
          null,
          4
          /* STYLE */
        )) : vue.createCommentVNode("v-if", true),
        !$props.isCircular ? (vue.openBlock(), vue.createElementBlock(
          "view",
          {
            key: 1,
            class: "progress-bar-inner",
            style: vue.normalizeStyle({ width: $props.progress + "%", backgroundColor: $props.progressBackgroundColor, height: $props.height + "rpx" })
          },
          [
            $props.showText ? (vue.openBlock(), vue.createElementBlock(
              "view",
              {
                key: 0,
                class: "progress-bar-text",
                style: vue.normalizeStyle({ color: $props.textColor, fontSize: $props.textSize + "rpx" })
              },
              vue.toDisplayString($props.progress + "%"),
              5
              /* TEXT, STYLE */
            )) : vue.createCommentVNode("v-if", true)
          ],
          4
          /* STYLE */
        )) : vue.createCommentVNode("v-if", true),
        $props.isCircular ? (vue.openBlock(), vue.createElementBlock(
          "view",
          {
            key: 2,
            class: "progress-bar-circular",
            style: vue.normalizeStyle({ width: $props.diameter + "rpx", height: $props.diameter + "rpx" })
          },
          [
            vue.createElementVNode("canvas", {
              "canvas-id": $props.canvasId,
              style: vue.normalizeStyle({ width: $props.diameter + "rpx", height: $props.diameter + "rpx" })
            }, null, 12, ["canvas-id"]),
            $props.showText ? (vue.openBlock(), vue.createElementBlock(
              "view",
              {
                key: 0,
                class: "progress-bar-text",
                style: vue.normalizeStyle({ color: $props.textColor, fontSize: $props.textSize + "rpx" })
              },
              vue.toDisplayString($props.progress + "%"),
              5
              /* TEXT, STYLE */
            )) : vue.createCommentVNode("v-if", true)
          ],
          4
          /* STYLE */
        )) : vue.createCommentVNode("v-if", true)
      ],
      4
      /* STYLE */
    );
  }
  const piaoyiProgressBar = /* @__PURE__ */ _export_sfc(_sfc_main$8, [["render", _sfc_render$8], ["__scopeId", "data-v-1dec2857"], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/uni_modules/piaoyi-progress-bar/components/piaoyi-progress-bar/piaoyi-progress-bar.vue"]]);
  const _sfc_main$7 = {
    components: {
      piaoyiProgressBar
    },
    data() {
      return {
        exam_id: 1,
        examTitle: "第一单元第一次小测",
        score: 95,
        totalQuestions: 20,
        correctAnswers: 19,
        //完成度百分比
        progress: 95
      };
    },
    onLoad(event) {
      this.progress = parseInt(event.progress);
      this.fetchData();
    },
    methods: {
      progressBarColor(progress) {
        if (progress >= 90) {
          return "#3FC681";
        } else if (progress >= 80) {
          return "#FFC107";
        } else if (progress >= 60) {
          return "#FF4949";
        } else {
          return "#EFEFF4";
        }
      },
      //todo:exam_score的值不需要从后端获取，而是从本地缓存中获取
      fetchData() {
        let examResult = uni.getStorageSync("examResult");
        if (examResult) {
          this.exam_id = examResult.exam_id;
          this.examTitle = examResult.examTitle;
          this.score = examResult.score;
          this.totalQuestions = examResult.totalQuestions;
          this.correctAnswers = examResult.correctAnswers;
          this.progress = parseInt(examResult.correctAnswers / examResult.totalQuestions * 100);
        }
      },
      toDetail() {
        uni.navigateTo({
          url: "/pages/exam_details/exam_details"
        });
      },
      toHome() {
        uni.reLaunch({
          url: "/pages/home/home"
        });
      }
    }
  };
  function _sfc_render$7(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_piaoyiProgressBar = vue.resolveComponent("piaoyiProgressBar");
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("view", { class: "background" }, [
        vue.createElementVNode("span", null, "试卷已提交"),
        vue.createElementVNode("image", {
          src: "/static/horse.svg",
          class: "horse"
        })
      ]),
      vue.createElementVNode("view", { class: "center-container" }, [
        vue.createElementVNode("view", { class: "exam-result" }, [
          vue.createElementVNode(
            "h3",
            { class: "exam-title" },
            vue.toDisplayString($data.examTitle),
            1
            /* TEXT */
          ),
          vue.createElementVNode(
            "span",
            { class: "exam-score" },
            vue.toDisplayString($data.score),
            1
            /* TEXT */
          ),
          vue.createElementVNode("span", { style: { "color": "#3FC681", "font-size": "25px" } }, "分"),
          vue.createElementVNode("br"),
          vue.createElementVNode(
            "span",
            { class: "exam-num" },
            "共" + vue.toDisplayString($data.totalQuestions) + "题",
            1
            /* TEXT */
          ),
          vue.createElementVNode("span", { class: "true-num" }, [
            vue.createTextVNode("答对"),
            vue.createElementVNode(
              "span",
              {
                style: vue.normalizeStyle({ color: $data.correctAnswers === $data.totalQuestions ? "#3FC681" : "#FF4949" })
              },
              vue.toDisplayString($data.correctAnswers),
              5
              /* TEXT, STYLE */
            ),
            vue.createTextVNode(
              "/" + vue.toDisplayString($data.totalQuestions) + "题",
              1
              /* TEXT */
            )
          ]),
          vue.createElementVNode("view", { style: { "margin-left": "100px", "margin-top": "20px" } }, [
            vue.createVNode(_component_piaoyiProgressBar, {
              canvasId: "progressCanvas4",
              progress: $data.progress,
              backgroundColor: "#EFEFF4",
              progressBackgroundColor: "#07C160",
              showText: true,
              textColor: "#456DE7",
              textSize: 48,
              height: 20,
              isCircular: true,
              diameter: 200
            }, null, 8, ["progress"]),
            vue.createElementVNode("view", { class: "bg" })
          ]),
          vue.createElementVNode("view", {
            class: "btn",
            onClick: _cache[0] || (_cache[0] = (...args) => $options.toDetail && $options.toDetail(...args))
          }, "考试详情")
        ])
      ]),
      vue.createElementVNode("view", {
        class: "btn",
        onClick: _cache[1] || (_cache[1] = (...args) => $options.toHome && $options.toHome(...args))
      }, "完成")
    ]);
  }
  const PagesFinishexamFinishexam = /* @__PURE__ */ _export_sfc(_sfc_main$7, [["render", _sfc_render$7], ["__scopeId", "data-v-6238ca85"], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/finishexam/finishexam.vue"]]);
  const _sfc_main$6 = {
    components: {
      CollapsibleView
    },
    data() {
      return {
        isExpanded: false
        // 初始状态为折叠
      };
    },
    computed: {
      imageSrc() {
        return this.isExpanded ? "../../static/up.png" : "../../static/up.png";
      }
    },
    methods: {
      toggleCollapse1() {
        this.isExpanded = !this.isExpanded;
      }
    }
  };
  function _sfc_render$6(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_CollapsibleView = vue.resolveComponent("CollapsibleView");
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("view", { class: "title-container" }, [
        vue.createElementVNode("image", {
          class: "back-icon",
          src: "/static/back.svg"
        }),
        vue.createElementVNode("span", null, "第1/20题")
      ]),
      vue.createElementVNode("view", { class: "question-container" }, [
        vue.createElementVNode("view", { style: { "width": "max-content" } }, [
          vue.createElementVNode("span", { style: { "font-size": "24px" } }, "题目"),
          vue.createElementVNode("hr", { style: { "border": "0", "border-top": "5px solid #456DE7", "height": "0", "border-radius": "2.5px" } })
        ]),
        vue.createElementVNode("view", { class: "question" }, [
          vue.createElementVNode("span", null, "—is your brother?"),
          vue.createElementVNode("span", null, "—He is a doctor."),
          vue.createElementVNode("span", { class: "answer" }, "A. What"),
          vue.createElementVNode("span", { class: "answer" }, "B. Who"),
          vue.createElementVNode("span", { class: "answer" }, "C. Where"),
          vue.createElementVNode("span", { class: "answer" }, "D. How")
        ])
      ]),
      vue.createElementVNode("hr", { style: { "border": "0", "border-bottom": "8px solid #e3e3e3", "height": "0" } }),
      vue.createElementVNode("view", { class: "answer-container" }, [
        vue.createElementVNode("view", { style: { "width": "max-content" } }, [
          vue.createElementVNode("span", { style: { "font-size": "24px" } }, "答案"),
          vue.createElementVNode("hr", { style: { "border": "0", "border-top": "5px solid #456DE7", "height": "0", "border-radius": "2.5px" } })
        ]),
        vue.createElementVNode("span", { class: "true-answer" }, "正确答案：A"),
        vue.createElementVNode("span", { class: "your-answer" }, "您的答案：B")
      ]),
      vue.createElementVNode("hr", { style: { "border": "0", "border-bottom": "8px solid #e3e3e3", "height": "0" } }),
      vue.createElementVNode("view", { class: "explain-container" }, [
        vue.createVNode(_component_CollapsibleView, {
          ref: "collapsibleView1",
          "is-expanded": $data.isExpanded
        }, {
          header: vue.withCtx(() => [
            vue.createElementVNode("view", {
              class: "header",
              onClick: _cache[0] || (_cache[0] = (...args) => $options.toggleCollapse1 && $options.toggleCollapse1(...args))
            }, [
              vue.createElementVNode("view", { style: { "width": "max-content" } }, [
                vue.createElementVNode("span", { style: { "font-size": "24px" } }, "解析"),
                vue.createElementVNode("hr", { style: { "border": "0", "border-top": "5px solid #456DE7", "height": "0", "border-radius": "2.5px" } })
              ]),
              vue.createElementVNode("image", {
                src: $options.imageSrc,
                class: vue.normalizeClass(["up-arrow", { "rotated": $data.isExpanded }])
              }, null, 10, ["src"])
            ])
          ]),
          default: vue.withCtx(() => [
            vue.createElementVNode("span", null, '在回答“What"开头的问句时，要注意回答的内容应该是描述事物的性质或身份，如职业、颜色、形状等。因此该题选A。')
          ]),
          _: 1
          /* STABLE */
        }, 8, ["is-expanded"])
      ])
    ]);
  }
  const PagesQuestionDetailQuestionDetail = /* @__PURE__ */ _export_sfc(_sfc_main$6, [["render", _sfc_render$6], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/questionDetail/questionDetail.vue"]]);
  const _sfc_main$5 = {
    data() {
      return {
        // 考试信息的数据属性
        exam_id: 1,
        name: "第一单元第一次小测",
        time: "20:00 ~ 21:00",
        examDuration: 60,
        questionNum: 20
      };
    },
    onLoad() {
      let startExam = JSON.parse(uni.getStorageSync("startExam"));
      this.exam_id = startExam.exam_id;
      this.name = startExam.name;
      this.time = this.formatTimeRange(startExam.start_time, startExam.duration);
      this.examDuration = startExam.duration;
      this.questionNum = startExam.question_num;
    },
    methods: {
      // 格式化时间范围字符串(20:00),60分钟=(20:00 ~ 20:59)
      formatTimeRange(start_time, duration) {
        const [startHour, startMinute] = start_time.split(":").map(Number);
        const endMinute = (startMinute + duration) % 60;
        const endHour = startHour + Math.floor((startMinute + duration) / 60);
        const endTime = `${endHour.toString().padStart(2, "0")}:${endMinute.toString().padStart(2, "0")}`;
        return `${start_time} ~${endTime}`;
      },
      startExam() {
        uni.navigateTo({
          url: `/pages/exam/exam?exam_id=${this.exam_id}&name=${this.name}`
        });
      }
    }
  };
  function _sfc_render$5(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createCommentVNode(" 头部标题栏 "),
      vue.createElementVNode("view", { class: "title-container" }, [
        vue.createElementVNode("image", {
          class: "back-icon",
          src: "/static/back.svg",
          onClick: _cache[0] || (_cache[0] = (...args) => _ctx.back && _ctx.back(...args))
        }),
        vue.createElementVNode("span", null, "考试")
      ]),
      vue.createCommentVNode(" 考试信息部分 "),
      vue.createElementVNode("view", { class: "exam-info" }, [
        vue.createElementVNode("h2", { title: $data.name }, vue.toDisplayString($data.name), 9, ["title"]),
        vue.createElementVNode(
          "h3",
          null,
          vue.toDisplayString($data.time),
          1
          /* TEXT */
        ),
        vue.createElementVNode("view", { class: "circle" }, [
          vue.createElementVNode("span", {
            class: "exam-time",
            title: $data.examDuration
          }, "考试时间" + vue.toDisplayString($data.examDuration) + "分钟", 9, ["title"]),
          vue.createElementVNode("span", {
            class: "exam-num",
            title: $data.questionNum
          }, "共" + vue.toDisplayString($data.questionNum) + "题", 9, ["title"])
        ]),
        vue.createCommentVNode(" 点击去考试按钮，调用 startExam 方法 "),
        vue.createElementVNode("view", {
          class: "start-exam-btn",
          onClick: _cache[1] || (_cache[1] = (...args) => $options.startExam && $options.startExam(...args))
        }, "去考试")
      ])
    ]);
  }
  const PagesStartexamStartexam = /* @__PURE__ */ _export_sfc(_sfc_main$5, [["render", _sfc_render$5], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/startexam/startexam.vue"]]);
  const _sfc_main$4 = {
    data() {
      return {
        haveExam: true,
        examcnt: 1,
        exams: [
          { id: 1, title: "第一单元第一次小测", time: "20:00 " },
          { id: 2, title: "第二单元第一次小测", time: "21:00 " }
        ]
      };
    },
    methods: {
      toexam() {
        this.$router.push("/exam");
      }
    }
  };
  function _sfc_render$4(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("view", { class: "background" }, [
        vue.createElementVNode("span", null, "Go for it！"),
        vue.createElementVNode("span", { style: { "color": "grey", "font-weight": "lighter", "margin-top": "20rpx", "font-size": "20px" } }, [
          vue.createTextVNode("今日已学"),
          vue.createElementVNode("span", { class: "study-num" }, "10"),
          vue.createTextVNode("个单词")
        ]),
        vue.createElementVNode("image", {
          src: "/static/sanyueqi.png",
          class: "sanyueqi"
        })
      ]),
      vue.createElementVNode("view", { class: "center-container" }, [
        vue.createElementVNode("view", { class: "border1" }, [
          vue.createElementVNode("span", { style: { "font-size": "20px", "margin-left": "-160px" } }, "你已连续学习"),
          vue.createElementVNode("br"),
          vue.createElementVNode("span", { class: "study-day" }, "10"),
          vue.createElementVNode("span", { style: { "color": "#6200EE", "font-size": "20px", "margin-left": "-10px" } }, "天"),
          vue.createElementVNode("image", {
            src: "/static/flash.png",
            class: "flash"
          }),
          vue.createElementVNode("view", { class: "btn" }, "查看我的足迹")
        ])
      ]),
      vue.withDirectives(vue.createElementVNode(
        "h3",
        { style: { "margin-left": "40px", "margin-top": "20px" } },
        [
          vue.createTextVNode("你今天有"),
          vue.createElementVNode(
            "span",
            null,
            vue.toDisplayString($data.examcnt),
            1
            /* TEXT */
          ),
          vue.createTextVNode("场考试")
        ],
        512
        /* NEED_PATCH */
      ), [
        [vue.vShow, $data.haveExam]
      ]),
      vue.withDirectives(vue.createElementVNode(
        "view",
        { class: "center-container2" },
        [
          (vue.openBlock(true), vue.createElementBlock(
            vue.Fragment,
            null,
            vue.renderList($data.exams, (exam) => {
              return vue.openBlock(), vue.createElementBlock("view", {
                class: "border2",
                key: exam.id
              }, [
                vue.createElementVNode("view", { class: "exam-info" }, [
                  vue.createElementVNode(
                    "h3",
                    null,
                    vue.toDisplayString(exam.title),
                    1
                    /* TEXT */
                  ),
                  vue.createCommentVNode(" 使用考试标题 "),
                  vue.createElementVNode(
                    "h4",
                    null,
                    "开始时间：" + vue.toDisplayString(exam.time),
                    1
                    /* TEXT */
                  ),
                  vue.createCommentVNode(" 使用考试时间 "),
                  vue.createElementVNode("view", {
                    class: "small-btn",
                    onClick: _cache[0] || (_cache[0] = (...args) => $options.toexam && $options.toexam(...args))
                  }, "去考试")
                ]),
                vue.createElementVNode("image", {
                  src: "/static/gotoexam.png",
                  style: { "width": "150px", "height": "150px" }
                })
              ]);
            }),
            128
            /* KEYED_FRAGMENT */
          ))
        ],
        512
        /* NEED_PATCH */
      ), [
        [vue.vShow, $data.haveExam]
      ]),
      vue.createElementVNode("view", {
        class: "btn",
        style: { "background-color": "#456de7", "font-size": "26px" }
      }, "完成")
    ]);
  }
  const PagesFinishClockinFinishClockin = /* @__PURE__ */ _export_sfc(_sfc_main$4, [["render", _sfc_render$4], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/finishClockin/finishClockin.vue"]]);
  const _sfc_main$3 = {
    data() {
      return {
        inputClearValue: ""
      };
    },
    methods: {
      handleBack() {
        uni.navigateBack({
          delta: 1
        });
      },
      clearIcon() {
        this.inputClearValue = "";
      },
      joinTeam() {
        uni.request({
          url: "/api/users/my_team/join_team",
          method: "POST",
          data: {
            invitation_code: this.inputClearValue
          },
          header: {
            "content-type": "application/json",
            // 默认值
            "Authorization": `Bearer ${uni.getStorageSync("token")}`
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
            if (res.data.code === 200 || res.data.code === "200") {
              uni.showToast({
                title: "加入成功",
                icon: "success",
                duration: 2e3
              });
              uni.navigateBack({
                delta: 1
              });
            } else {
              uni.showToast({
                title: "加入失败",
                icon: "none",
                duration: 2e3
              });
            }
          },
          fail: (res) => {
            uni.showToast({
              title: "加入失败",
              icon: "none",
              duration: 2e3
            });
          }
        });
      }
    }
  };
  function _sfc_render$3(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", { class: "container" }, [
      vue.createElementVNode("view", { class: "head" }, [
        vue.createElementVNode("image", {
          class: "back-icon",
          src: "/static/back.svg",
          onClick: _cache[0] || (_cache[0] = (...args) => $options.handleBack && $options.handleBack(...args))
        }),
        vue.createElementVNode("span", null, "加入团队"),
        vue.createElementVNode("button", {
          onClick: _cache[1] || (_cache[1] = (...args) => $options.joinTeam && $options.joinTeam(...args))
        }, "确认")
      ]),
      vue.createElementVNode("view", { class: "input-wrapper" }, [
        vue.withDirectives(vue.createElementVNode(
          "input",
          {
            class: "uni-input",
            placeholder: "请输入团队码",
            "onUpdate:modelValue": _cache[2] || (_cache[2] = ($event) => $data.inputClearValue = $event)
          },
          null,
          512
          /* NEED_PATCH */
        ), [
          [vue.vModelText, $data.inputClearValue]
        ]),
        $data.inputClearValue.length > 0 ? (vue.openBlock(), vue.createElementBlock("image", {
          key: 0,
          class: "uni-icon",
          src: "/static/not-done2.svg",
          onClick: _cache[3] || (_cache[3] = (...args) => $options.clearIcon && $options.clearIcon(...args))
        })) : vue.createCommentVNode("v-if", true)
      ])
    ]);
  }
  const PagesJoinJoin = /* @__PURE__ */ _export_sfc(_sfc_main$3, [["render", _sfc_render$3], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/Join/Join.vue"]]);
  const _sfc_main$2 = {
    data() {
      return {
        teamName: "春田花花幼稚园",
        managerName: "佐藤太郎",
        memberNum: 50,
        members: [
          {
            userName: "张三",
            userSex: 1
            //1:男 0:女
          },
          {
            userName: "李四",
            userSex: 0
            //1:男 0:女
          }
        ]
      };
    },
    onLoad() {
      uni.request({
        url: "/api/users/my_team",
        method: "GET",
        header: {
          "Authorization": `Bearer ${uni.getStorageSync("token")}`
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
          if (res.data.code == 200) {
            let teamInfo = res.data.team;
            this.teamName = teamInfo.team_name;
            this.managerName = teamInfo.manager_name;
            this.memberNum = teamInfo.member_num;
            this.members = [];
            teamInfo.member_list.forEach((member) => {
              this.members.push({
                userName: member.user_name,
                userSex: member.user_sex
              });
            });
            this.setData({
              firstLetter: this.data.managerName[0]
            });
          }
        },
        fail: (error) => {
          formatAppLog("log", "at pages/MyTeam/MyTeam.vue:94", error);
        }
      });
    },
    methods: {
      goToJoin() {
        uni.navigateTo({
          url: "../Join/Join"
        });
      }
    }
  };
  function _sfc_render$2(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", { class: "container" }, [
      vue.createElementVNode("view", { class: "head" }, [
        vue.createElementVNode("image", {
          class: "back-icon",
          src: "/static/back.svg",
          onClick: _cache[0] || (_cache[0] = (...args) => _ctx.handleBack && _ctx.handleBack(...args))
        }),
        vue.createElementVNode("span", { class: "title" }, "我的团队"),
        vue.createElementVNode("button", {
          onClick: _cache[1] || (_cache[1] = (...args) => $options.goToJoin && $options.goToJoin(...args))
        }, "加入团队")
      ]),
      vue.createElementVNode("view", { class: "container1" }, [
        vue.createElementVNode("view", { class: "team-icon-container" }, [
          vue.createElementVNode("span", { class: "team-initial" }, "M")
        ]),
        vue.createElementVNode("view", { class: "team-info" }, [
          vue.createElementVNode("view", { class: "team-name" }, [
            vue.createElementVNode(
              "text",
              { class: "team-name-text" },
              vue.toDisplayString($data.teamName),
              1
              /* TEXT */
            )
          ]),
          vue.createElementVNode("view", { class: "people" }, [
            vue.createElementVNode(
              "text",
              { class: "people-text" },
              "团队人数：" + vue.toDisplayString($data.memberNum),
              1
              /* TEXT */
            )
          ])
        ])
      ]),
      vue.createElementVNode("view", { class: "container2" }, [
        vue.createElementVNode("view", { class: "team-captain-row" }, [
          vue.createElementVNode("span", null, "团队队长")
        ]),
        vue.createElementVNode("view", { class: "name-row" }, [
          vue.createElementVNode("view", { class: "people-icon-container" }, [
            vue.createElementVNode(
              "span",
              { class: "people-initial" },
              vue.toDisplayString(_ctx.firstLetter),
              1
              /* TEXT */
            )
          ]),
          vue.createElementVNode(
            "span",
            { class: "manager-name" },
            vue.toDisplayString($data.managerName),
            1
            /* TEXT */
          )
        ])
      ]),
      vue.createElementVNode("view", { class: "container3" }, [
        vue.createElementVNode("view", { class: "team-captain-row" }, [
          vue.createElementVNode("span", null, "团队成员")
        ]),
        (vue.openBlock(true), vue.createElementBlock(
          vue.Fragment,
          null,
          vue.renderList($data.members, (item, index) => {
            return vue.openBlock(), vue.createElementBlock("view", {
              key: index,
              class: "name-row"
            }, [
              vue.createElementVNode("view", { class: "people-icon-container" }, [
                vue.createElementVNode(
                  "span",
                  { class: "people-initial" },
                  vue.toDisplayString(item.userName[0]),
                  1
                  /* TEXT */
                )
              ]),
              vue.createElementVNode(
                "span",
                { class: "manager-name" },
                vue.toDisplayString(item.userName),
                1
                /* TEXT */
              )
            ]);
          }),
          128
          /* KEYED_FRAGMENT */
        ))
      ])
    ]);
  }
  const PagesMyTeamMyTeam = /* @__PURE__ */ _export_sfc(_sfc_main$2, [["render", _sfc_render$2], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/MyTeam/MyTeam.vue"]]);
  const dbName = "dailyenglish";
  const dbPath = "_doc/dailyEnglish.db";
  function open() {
    plus.sqlite.openDatabase({
      name: dbName,
      path: dbPath,
      success(e) {
        formatAppLog("log", "at sqlite/sqlite.js:8", "打开数据库成功");
        formatAppLog("log", "at sqlite/sqlite.js:9", e);
        resolve(e);
      },
      fail(e) {
        formatAppLog("log", "at sqlite/sqlite.js:13", "打开数据库失败");
        reject(e);
      }
    });
  }
  function close() {
    plus.sqlite.closeDatabase({
      name: dbName,
      success(e) {
        formatAppLog("log", "at sqlite/sqlite.js:30", "关闭数据库成功");
        resolve(e);
      },
      fail(e) {
        formatAppLog("log", "at sqlite/sqlite.js:34", "关闭数据库失败");
        reject(e);
      }
    });
  }
  function createWordsTable() {
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
  );`;
    plus.sqlite.executeSql({
      name: dbName,
      sql,
      success(e) {
        formatAppLog("log", "at sqlite/sqlite.js:63", "创建表成功");
        resolve(e);
      },
      fail(e) {
        formatAppLog("log", "at sqlite/sqlite.js:67", "创建表失败");
        formatAppLog("log", "at sqlite/sqlite.js:68", e);
        reject(e);
      }
    });
  }
  function insertword(words) {
    const sql = `INSERT INTO word (word_id, word, phonetic_us, describe, morpheme, example_sentence, other, word_quetion, answer, learn_times, interval_history, feedback_history, interval_days, difficulty, is_memory) VALUES (${words.word_id}, '${words.word}', '${words.phonetic_us}', '${words.describe}', '${words.morpheme}', '${words.example_sentence}', '${words.other}', '${words.word_quetion}', '${words.answer}', ${words.learn_times}, '${words.interval_history}', '${words.feedback_history}', ${words.interval_days}, ${words.difficulty}, ${words.is_memory})`;
    plus.sqlite.executeSql({
      name: dbName,
      sql,
      success(e) {
        formatAppLog("log", "at sqlite/sqlite.js:80", "插入数据成功");
        resolve(e);
      },
      fail(e) {
        formatAppLog("log", "at sqlite/sqlite.js:84", "插入数据失败");
        formatAppLog("log", "at sqlite/sqlite.js:85", e);
        reject(e);
      }
    });
  }
  function selectword() {
    const sql = `SELECT * FROM word`;
    plus.sqlite.selectSql({
      name: dbName,
      sql,
      success(data2) {
        formatAppLog("log", "at sqlite/sqlite.js:151", "查询数据成功");
        formatAppLog("log", "at sqlite/sqlite.js:152", data2);
        resolve(data2);
        return data2;
      },
      fail(e) {
        formatAppLog("log", "at sqlite/sqlite.js:157", "查询数据失败");
        reject(e);
      }
    });
  }
  const _sfc_main$1 = {
    data() {
      return {
        words: []
      };
    },
    methods: {
      async openDatabase() {
        try {
          await open();
          formatAppLog("log", "at pages/sql/sql.vue:50", "数据库已打开");
        } catch (error) {
          formatAppLog("error", "at pages/sql/sql.vue:52", `打开数据库失败: ${error.message}`);
        }
      },
      async closeDatabase() {
        try {
          await close();
          formatAppLog("log", "at pages/sql/sql.vue:58", "数据库已关闭");
        } catch (error) {
          formatAppLog("error", "at pages/sql/sql.vue:60", `关闭数据库失败: ${error.message}`);
        }
      },
      async createTable() {
        try {
          await createWordsTable();
          formatAppLog("log", "at pages/sql/sql.vue:66", "表已创建");
        } catch (error) {
          formatAppLog("error", "at pages/sql/sql.vue:68", `创建表失败: ${error.message}`);
        }
      },
      async insertWord() {
        const word = {
          word_id: 1,
          word: "apple",
          phonetic_us: "/",
          describe: "苹果",
          morpheme: "ap+ple",
          example_sentence: "I like apple",
          other: "n. 苹果；苹果树；苹果似的东西；[美俚]头；[美俚]脑袋；[美俚]人；[美俚]家伙；[美俚]家庭；[美俚]公司；[美俚]事情；[美俚]事业；[美俚]目标；[美俚]目的；[美俚]目的地",
          word_quetion: "apple",
          answer: "苹果",
          learn_times: 0,
          interval_history: "0",
          feedback_history: "0",
          interval_days: 1,
          difficulty: 1,
          is_memory: 0
        };
        try {
          await insertword(word);
          formatAppLog("log", "at pages/sql/sql.vue:92", "单词已插入");
        } catch (error) {
          formatAppLog("error", "at pages/sql/sql.vue:94", `插入单词失败: ${error.message}`);
        }
      },
      async selectWords() {
        try {
          this.words = await selectword();
          formatAppLog("log", "at pages/sql/sql.vue:100", this.words);
          formatAppLog("log", "at pages/sql/sql.vue:101", "查询成功");
        } catch (error) {
          formatAppLog("error", "at pages/sql/sql.vue:103", `查询失败: ${error.message}`);
        }
      }
    }
  };
  function _sfc_render$1(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("div", null, [
      vue.createElementVNode("button", {
        onClick: _cache[0] || (_cache[0] = (...args) => $options.openDatabase && $options.openDatabase(...args))
      }, "打开数据库"),
      vue.createElementVNode("button", {
        onClick: _cache[1] || (_cache[1] = (...args) => $options.closeDatabase && $options.closeDatabase(...args))
      }, "关闭数据库"),
      vue.createElementVNode("button", {
        onClick: _cache[2] || (_cache[2] = (...args) => $options.createTable && $options.createTable(...args))
      }, "创建表"),
      vue.createElementVNode("button", {
        onClick: _cache[3] || (_cache[3] = (...args) => $options.insertWord && $options.insertWord(...args))
      }, "插入单词"),
      vue.createElementVNode("button", {
        onClick: _cache[4] || (_cache[4] = (...args) => $options.selectWords && $options.selectWords(...args))
      }, "查询单词"),
      vue.createElementVNode("div", null, [
        vue.createElementVNode("h2", null, "单词列表："),
        vue.createElementVNode("table", null, [
          vue.createElementVNode("thead", null, [
            vue.createElementVNode("tr", null, [
              vue.createElementVNode("th", null, "new_id"),
              vue.createElementVNode("th", null, "word_id"),
              vue.createElementVNode("th", null, "word"),
              vue.createElementVNode("th", null, "phonetic_us"),
              vue.createElementVNode("th", null, "describe"),
              vue.createCommentVNode(" 添加其他列的表头 ")
            ])
          ]),
          vue.createElementVNode("tbody", null, [
            (vue.openBlock(true), vue.createElementBlock(
              vue.Fragment,
              null,
              vue.renderList($data.words, (word) => {
                return vue.openBlock(), vue.createElementBlock("tr", {
                  key: word.new_id
                }, [
                  vue.createElementVNode(
                    "td",
                    null,
                    vue.toDisplayString(word.new_id),
                    1
                    /* TEXT */
                  ),
                  vue.createElementVNode(
                    "td",
                    null,
                    vue.toDisplayString(word.word_id),
                    1
                    /* TEXT */
                  ),
                  vue.createElementVNode(
                    "td",
                    null,
                    vue.toDisplayString(word.word),
                    1
                    /* TEXT */
                  ),
                  vue.createElementVNode(
                    "td",
                    null,
                    vue.toDisplayString(word.phonetic_us),
                    1
                    /* TEXT */
                  ),
                  vue.createElementVNode(
                    "td",
                    null,
                    vue.toDisplayString(word.describe),
                    1
                    /* TEXT */
                  ),
                  vue.createCommentVNode(" 显示其他列的数据 ")
                ]);
              }),
              128
              /* KEYED_FRAGMENT */
            ))
          ])
        ])
      ])
    ]);
  }
  const PagesSqlSql = /* @__PURE__ */ _export_sfc(_sfc_main$1, [["render", _sfc_render$1], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/pages/sql/sql.vue"]]);
  __definePage("pages/home/home", PagesHomeHome);
  __definePage("pages/word_details/word_details", PagesWord_detailsWord_details);
  __definePage("pages/Examination/Examination", PagesExaminationExamination);
  __definePage("pages/index/index", PagesIndexIndex);
  __definePage("pages/index/HelloWorld", PagesIndexHelloWorld);
  __definePage("pages/user/user", PagesUserUser);
  __definePage("pages/Vocab/Vocab", PagesVocabVocab);
  __definePage("pages/example/example", PagesExampleExample);
  __definePage("pages/login/login", PagesLoginLogin);
  __definePage("pages/register/register", PagesRegisterRegister);
  __definePage("pages/Welcome/Welcome", PagesWelcomeWelcome);
  __definePage("pages/Loading/Loading", PagesLoadingLoading);
  __definePage("pages/PopUp/PopUp", PagesPopUpPopUp);
  __definePage("pages/WordBlock/WordBlock", PagesWordBlockWordBlock);
  __definePage("pages/personal-information/personal-information", PagesPersonalInformationPersonalInformation);
  __definePage("pages/Calendar/Calendar", PagesCalendarCalendar);
  __definePage("pages/exam_details/exam_details", PagesExam_detailsExam_details);
  __definePage("pages/exam/exam", PagesExamExam);
  __definePage("pages/ExamHistory/ExamHistory", PagesExamHistoryExamHistory);
  __definePage("pages/finishexam/finishexam", PagesFinishexamFinishexam);
  __definePage("pages/questionDetail/questionDetail", PagesQuestionDetailQuestionDetail);
  __definePage("pages/startexam/startexam", PagesStartexamStartexam);
  __definePage("pages/finishClockin/finishClockin", PagesFinishClockinFinishClockin);
  __definePage("pages/Join/Join", PagesJoinJoin);
  __definePage("pages/MyTeam/MyTeam", PagesMyTeamMyTeam);
  __definePage("pages/sql/sql", PagesSqlSql);
  const _sfc_main = {
    components: {
      TabBar
    }
  };
  function _sfc_render(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_router_link = vue.resolveComponent("router-link");
    const _component_router_view = vue.resolveComponent("router-view");
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createCommentVNode(" 应用的其他内容 "),
      vue.createVNode(_component_router_link, {
        class: "uni-nav",
        to: "/pages/HelloWorld/HelloWorld"
      }, {
        default: vue.withCtx(() => [
          vue.createTextVNode(" Go to HelloWorld ")
        ]),
        _: 1
        /* STABLE */
      }),
      vue.createCommentVNode(" 渲染当前路由对应的组件 "),
      vue.createVNode(_component_router_view)
    ]);
  }
  const App = /* @__PURE__ */ _export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "D:/code/DailyEnglish/app/DailyEnglish/App.vue"]]);
  function createApp() {
    const app = vue.createVueApp(App);
    return {
      app
    };
  }
  const { app: __app__, Vuex: __Vuex__, Pinia: __Pinia__ } = createApp();
  uni.Vuex = __Vuex__;
  uni.Pinia = __Pinia__;
  __app__.provide("__globalStyles", __uniConfig.styles);
  __app__._component.mpType = "app";
  __app__._component.render = () => {
  };
  __app__.mount("#app");
})(Vue);
