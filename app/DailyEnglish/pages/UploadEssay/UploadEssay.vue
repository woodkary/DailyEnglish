<template>
	<view>
    <uni-file-picker fileMediatype="image" mode="grid" @select="onFileSelected(event)"></uni-file-picker>
		<button type="primary" @click="uploadImage">上传图片</button>
	</view>
</template>

<script>
	import UniFilePicker from "../../uni_modules/uni-file-picker/components/uni-file-picker/uni-file-picker.vue";

  export default {
    components: {UniFilePicker},
		data() {
			return {
				
			}
		},
		onLoad() {

		},

		methods: {
      onFileSelected(e) {
        // 获取文件对象
        const file = e.tempFiles[0];
        // 获取文件的临时路径
        const tempFilePath = file.path;
        // 转换为 Base64
        this.getBase64(tempFilePath);
      },
      async getBase64(tempFilePath) {
        // 读取文件内容并转换为 Base64 编码
        const base64 = await new Promise((resolve, reject) => {
          uni.getFileSystemManager().readFile({
            filePath: tempFilePath,
            encoding: 'base64',
            success: res => {
              resolve(res.data);
            },
            fail: err => {
              reject(err);
            }
          });
        });
        // 发送到后端
        this.uploadBase64ToServer(base64);
      },
      uploadBase64ToServer(base64) {
        // 构建请求体
        const data = {
          // 根据后端要求，这里可能是 'image'，'file'，或者其他字段
          image: base64
        };
        // 发送请求
        uni.request({
          url: 'http://localhost:8080/api/users/upload',
          method: 'POST',
          data,
          header: {
            'content-type': 'application/json'
          },
          success: res => {
            console.log('上传成功', res);
          },
          fail: err => {
            console.error('上传失败', err);
          }
        });
      },
			  uploadImage() {
          uni.chooseImage({
            count: 9, //默认9
            sizeType: ['original', 'compressed'], //可以指定是原图还是压缩图，默认二者都有
            sourceType: ['album','camera'], //从相册选择，或拍照
            success: function (res) {
              console.log(JSON.stringify(res.tempFilePaths));
            }
          });
        }
		}
	}
</script>

<style>

</style>
