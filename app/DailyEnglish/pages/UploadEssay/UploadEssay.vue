<template>
  <view>
    <uni-file-picker fileMediatype="image" mode="grid" @select="onFileSelected"></uni-file-picker>
    <button type="primary" @click="uploadImage">上传图片</button>
  </view>
</template>

<script>
import UniFilePicker from "../../uni_modules/uni-file-picker/components/uni-file-picker/uni-file-picker.vue";

export default {
  components: { UniFilePicker },
  data() {
    return {
      titleId: 1, // 文章ID
      base64Data: '' // 存储base64数据
    };
  },
  onLoad(event){
    this.titleId = event.titleId; // 获取文章ID
  },
  methods: {
    onFileSelected(e) {
      console.log('File selected:', e);
      const file = e.tempFiles[0];
      if (file) {
        const tempFilePath = file.path;
        console.log('Temp file path:', tempFilePath);
        this.getBase64(tempFilePath);
      } else {
        console.error('No file selected');
      }
    },
    async getBase64(tempFilePath) {
      console.log('Getting base64 for:', tempFilePath);
      try {
        const base64 = await this.convertToBase64(tempFilePath);
        console.log('Base64:', base64);
        this.base64Data = base64; // 存储base64数据
      } catch (error) {
        console.error('Error getting base64:', error);
      }
    },
    convertToBase64(filePath) {
      return new Promise((resolve, reject) => {
        // #ifdef MP-WEIXIN
        uni.getFileSystemManager().readFile({
          filePath: filePath,
          encoding: 'base64',
          success: res => {
            resolve(res.data);
          },
          fail: err => {
            reject(err);
          }
        });
        // #endif

        // #ifdef H5
        const reader = new FileReader();
        fetch(filePath)
          .then(response => response.blob())
          .then(blob => {
            reader.readAsDataURL(blob);
            reader.onload = () => {
              resolve(reader.result.split(',')[1]);
            };
            reader.onerror = error => {
              reject(error);
            };
          })
          .catch(err => {
            reject(err);
          });
        // #endif

        // #ifdef APP-PLUS
        plus.io.resolveLocalFileSystemURL(filePath, function(entry) {
          entry.file(function(file) {
            const reader = new plus.io.FileReader();
            reader.onloadend = function(e) {
              const base64 = e.target.result.split(',')[1];
              resolve(base64);
            };
            reader.onerror = function(e) {
              reject(e);
            };
            reader.readAsDataURL(file);
          });
        }, function(err) {
          reject(err);
        });
        // #endif
      });
    },
    uploadBase64ToServer(base64) {
      console.log('Uploading base64 to server');
      const data = {
        title_id: this.titleId,
        image: base64
      };
      uni.request({
        url: 'http://localhost:8080/api/users/upload',
        method: 'POST',
        header:{
          'Authorization': `Bearer ${uni.getStorageSync('token')}`
        },
        data,
        success: res => {
          console.log('上传成功', res);
          uni.navigateTo({
            url: '/pages/EssayResult/EssayResult?titleId=' + this.titleId
          });
        },
        fail: err => {
          console.error('上传失败', err);
        }
      });
    },
    uploadImage() {
      this.uploadBase64ToServer(this.base64Data); // 上传base64数据
    }
  }
}
</script>

<style>
</style>
