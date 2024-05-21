### piaoyi-progress-bar 进度条

**使用方法：**

正常进度条
```
<piaoyiProgressBar :progress="50" backgroundColor="#EFEFF4" progressBackgroundColor="#07C160" :showText="true"
	textColor="#fff" :textSize="24" :height="20"></piaoyiProgressBar>
<view class="bg">
```

环状进度条
```
<piaoyiProgressBar canvasId="progressCanvas4" :progress="75" backgroundColor="#EFEFF4"
	progressBackgroundColor="#07C160" :showText="true" textColor="#000000" :textSize="24" :height="20"
	:isCircular="true" :diameter="200"></piaoyiProgressBar>
<view class="bg">
```

```
import piaoyiProgressBar from '@/uni_modules/piaoyi-progress-bar/components/piaoyi-progress-bar/piaoyi-progress-bar.vue';
export default {
    components: {
        piaoyiProgressBar
    },
    data() {
        return {
        }
    },
    methods: {
    }
}
```

#### 事件说明

无

#### Prop

| 参数名称 | 描述                           | 默认值                         |
| -------- | ------------------------------ | ------------------------------ |
| progress | 进度值 |       空（范围：0-100）        |
| backgroundColor | 背景色 |   #EFEFF4   |
| progressBackgroundColor | 进度背景色 |   #07C160   |
| showText | 是否显示文本 |   true   |
| textColor | 文本颜色 |   #000000   |
| textSize | 文本大小 |       24     |
| height | 进度条线条宽度 |       20     |
| diameter | 进度条整体大小 |       200     |
| isCircular | 是否显示环状 |       false     |
| canvasId | isCircular为true时必传,一个页面使用多个需要定义不同的canvasId |       canvasId     |

### 可接定制化组件开发
### 右侧有本人代表作小程序二维码，可以扫码体验
### 如使用过程中有问题或有一些好的建议，欢迎加QQ群互相学习交流：120594820