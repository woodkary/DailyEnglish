<template>
	<view>
		<view class="head">
			<image class="back-icon" src="../../static/back.svg" @click="handleBack"></image>
			<span class="title">作文提交</span>
			<view class="composition-container">
				<span class="line1">题目：<span style="font-weight: normal;">{{ title }}</span> </span>
				<span class="line2">字数：<span style="font-weight: normal;">{{ word_cnt }}</span> </span>
				<span class="line3">要求： <span
						class="req">请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文请根据题目要求完成作文</span></span>
			</view>
		</view>
		<view class="container">
		<uni-file-picker v-show="showOnload" fileMediatype="image" mode="grid" @select="onFileSelected" style="margin-top:40px;"></uni-file-picker>
		<button v-show="showInput" @click="setShow(false)" class="upload">上传图片</button>
		<button v-show="showOnload" type="primary" class="submit" @click="setShow(true)">直接写作</button>
		<textarea v-show="showInput" auto-height maxlength="-1" placeholder="请输入" class="edit"/>
		<button type="primary" class="submit" @click="submit">提交</button>
	</view>
	</view>
</template>

<script>
	import UniFilePicker from "../../uni_modules/uni-file-picker/components/uni-file-picker/uni-file-picker.vue";

	export default {
		components: {
			UniFilePicker
		},
		data() {
			return {
				titleId: 1, // 文章ID
				base64Data: '' ,// 存储base64数据
				showInput: false,
				showOnload: true,
			};
		},
		onLoad(event) {
			this.titleId = event.titleId; // 获取文章ID
		},
		methods: {
      handleBack() {
        uni.navigateBack();
      },
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
					header: {
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

			},
			setShow(option){
				this.showInput = option;
				this.showOnload = !option;
			},
			submit() {
        if(this.base64Data==''){
          uni.showToast({
            title: '请先上传图片',
            icon: 'none'
          })
          return;
        }
				uni.showModal({
					title: '提示',
					content: '确定要提交作文吗？',
					success: (res)=> {
						if (res.confirm) {
							console.log('用户点击确定');
              this.uploadBase64ToServer(this.base64Data); // 上传base64数据
						} else if (res.cancel) {
							console.log('用户点击取消');
						}
					}
				});
			},
		}
	}
</script>

<style>
	.head {
		width: 100%;
		position: relative;
		top: 0;
		left: 0;
		z-index: 100;
		height: 280px;
	}

	.back-icon {
		width: 2rem;
		height: 2rem;
		position: absolute;
		top: 0.8rem;
		left: 0.5rem;
		cursor: pointer;
	}

	.title {
		position: absolute;
		font-size: 1.2rem;
		font-weight: bold;
		text-align: center;
		top: 0.8rem;
		left: 40%;
	}

	.composition-container {
		position: absolute;
		top: 4rem;
		left: 5%;
		width: 90%;
		height: 190px;
		display: flex;
		flex-direction: column;
		border: 1.6px solid #1b6ef3;
		border-radius: 1.5rem;
	}

	.line1,
	.line2,
	.line3 {
		font-size: 1rem;
		margin-left: 1rem;
		font-weight: bold;
		margin-right: 1rem;
	}

	.line1 {
		margin-top: 0.7rem;
	}

	.line2,
	.line3 {
		margin-top: 0.2rem;
	}

	.line3 {
		max-height: 120px;
		min-height: 120px;
	}

	.req {
		font-weight: normal;
		display: block;
		line-height: 20px;
		max-height: 80px;
		/* 20px * 3 */
		overflow-y: auto;
	}
	.container{
		display: flex;
		flex-direction: column;
		align-items: center;
	}
	.upload{
		margin-top: 20px;
		width: 80%;
		height: 40px;
		border-radius: 5%;
		background-color: white;
		border:1px solid #bababa;
		text-align: center;
		line-height: 40px;
	}
	.edit{
		margin-top: 20px;
		width: 90%;
		height: fit-content;
		min-height: 200px;
		border-radius: 5%;
		background-color: white;
		border:1px solid #bababa;
		text-align: left;
		line-height: 30px;
		padding: 10px;
		font-size:20px;
	}

	.submit{
		margin-top: 20px;
		width: 80%;
		height: 40px;
		border-radius: 5%;
		background-color: #1b6ef3;
		color: white;
		text-align: center;
		line-height: 40px;
	}
</style>