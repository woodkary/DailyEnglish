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
  const global = uni.requireGlobal();
  ArrayBuffer = global.ArrayBuffer;
  Int8Array = global.Int8Array;
  Uint8Array = global.Uint8Array;
  Uint8ClampedArray = global.Uint8ClampedArray;
  Int16Array = global.Int16Array;
  Uint16Array = global.Uint16Array;
  Int32Array = global.Int32Array;
  Uint32Array = global.Uint32Array;
  Float32Array = global.Float32Array;
  Float64Array = global.Float64Array;
  BigInt64Array = global.BigInt64Array;
  BigUint64Array = global.BigUint64Array;
}
;
if (uni.restoreGlobal) {
  uni.restoreGlobal(Vue, weex, plus, setTimeout, clearTimeout, setInterval, clearInterval);
}
(function(vue) {
  "use strict";
  const _export_sfc = (sfc, props) => {
    const target = sfc.__vccOpts || sfc;
    for (const [key, val] of props) {
      target[key] = val;
    }
    return target;
  };
  const _sfc_main$5 = {};
  function _sfc_render$5(_ctx, _cache) {
    const _component_router_link = vue.resolveComponent("router-link");
    return vue.openBlock(), vue.createElementBlock("view", { class: "tab-bar" }, [
      vue.createVNode(_component_router_link, {
        class: "tab-item",
        to: "/pages/Home/Home"
      }, {
        default: vue.withCtx(() => [
          vue.createTextVNode("Home")
        ]),
        _: 1
        /* STABLE */
      }),
      vue.createVNode(_component_router_link, {
        class: "tab-item",
        to: "/pages/About/About"
      }, {
        default: vue.withCtx(() => [
          vue.createTextVNode("About")
        ]),
        _: 1
        /* STABLE */
      }),
      vue.createCommentVNode(" 添加其他导航项 ")
    ]);
  }
  const TabBar = /* @__PURE__ */ _export_sfc(_sfc_main$5, [["render", _sfc_render$5], ["__scopeId", "data-v-89ca1f91"], ["__file", "C:/Users/fifi/Documents/HBuilderProjects/DailyEnglish/components/TabBar.vue"]]);
  const _sfc_main$4 = {
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
  function _sfc_render$4(_ctx, _cache, $props, $setup, $data, $options) {
    const _component_TabBar = vue.resolveComponent("TabBar");
    return vue.openBlock(), vue.createElementBlock("view", { class: "content" }, [
      vue.createVNode(_component_TabBar),
      vue.createElementVNode("image", {
        class: "logo",
        src: "/static/logo.png"
      }),
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
  const PagesIndexIndex = /* @__PURE__ */ _export_sfc(_sfc_main$4, [["render", _sfc_render$4], ["__file", "C:/Users/fifi/Documents/HBuilderProjects/DailyEnglish/pages/index/index.vue"]]);
  const _sfc_main$3 = {};
  function _sfc_render$3(_ctx, _cache) {
    return vue.openBlock(), vue.createElementBlock("view", null, [
      vue.createElementVNode("text", { class: "hello-world" }, "Hello World!")
    ]);
  }
  const PagesIndexHelloWorld = /* @__PURE__ */ _export_sfc(_sfc_main$3, [["render", _sfc_render$3], ["__file", "C:/Users/fifi/Documents/HBuilderProjects/DailyEnglish/pages/index/HelloWorld.vue"]]);
  const _sfc_main$2 = {
    data() {
      return {};
    },
    methods: {}
  };
  function _sfc_render$2(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view");
  }
  const PagesUserUser = /* @__PURE__ */ _export_sfc(_sfc_main$2, [["render", _sfc_render$2], ["__file", "C:/Users/fifi/Documents/HBuilderProjects/DailyEnglish/pages/user/user.vue"]]);
  const _sfc_main$1 = {
    methods: {
      handleBack() {
      },
      handleJump() {
      }
    }
  };
  function _sfc_render$1(_ctx, _cache, $props, $setup, $data, $options) {
    return vue.openBlock(), vue.createElementBlock("view", { class: "container" }, [
      vue.createElementVNode("image", {
        class: "back-icon",
        src: "/static/back.svg",
        onClick: _cache[0] || (_cache[0] = (...args) => $options.handleBack && $options.handleBack(...args))
      }),
      vue.createElementVNode("view", { class: "text-info" }, [
        vue.createElementVNode("text", { class: "word" }, "abandon"),
        vue.createElementVNode("text", { class: "phonetic" }, "[ə'bændən]")
      ]),
      vue.createElementVNode("view", { class: "button-group" }, [
        vue.createElementVNode("button", { class: "button" }, "A"),
        vue.createElementVNode("button", { class: "button" }, "B"),
        vue.createElementVNode("button", { class: "button" }, "C"),
        vue.createElementVNode("button", { class: "button" }, "D")
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
      ])
    ]);
  }
  const PagesExaminationExamination = /* @__PURE__ */ _export_sfc(_sfc_main$1, [["render", _sfc_render$1], ["__file", "C:/Users/fifi/Documents/HBuilderProjects/DailyEnglish/pages/Examination/Examination.vue"]]);
  __definePage("pages/index/index", PagesIndexIndex);
  __definePage("pages/index/HelloWorld", PagesIndexHelloWorld);
  __definePage("pages/user/user", PagesUserUser);
  __definePage("pages/Examination/Examination", PagesExaminationExamination);
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
  const App = /* @__PURE__ */ _export_sfc(_sfc_main, [["render", _sfc_render], ["__file", "C:/Users/fifi/Documents/HBuilderProjects/DailyEnglish/App.vue"]]);
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
