<template>
	<view class="container">
		<image class="back-icon" src="../../static/back.svg" @click="GotoUser"></image>
		<span class="title">个人信息</span>
		<view class="edit">
			<image class="photo" src="../../static/logo.png"></image>
			<view class="items">
				<EditableItemView 
						id="username"
				      label="用户名"  
				      :value="username"  
				      placeholder="kary"
				      @update:value="username = $event"  
				    />
				<EditableItemView
				id="email"
				      label="邮箱"  
				      :value="email"  
				      placeholder="834479572@qq.com"
				      @update:value="email = $event"  
				    />
				<EditableItemView
				id="phone"
				      label="手机号"  
				      :value="phone"  
				      placeholder="18923208090"
				      @update:value="phone = $event"  
				    />
				<EditableItemView
				 id="birthday"
				  label="生日"  
				  :value="birthday"  
				  placeholder="2003-05-28"
				  @update:value="birthday = $event"  
				/>  
				<EditableItemView
				id="team"
				      label="团队"  
				      :value="team"  
				      placeholder="英语每日练小组"
				      @update:value="team = $event"  
				    />
			</view>
		</view>
	</view>
</template>

<script>
	import EditableItemView from '../../components/EditableItemView.vue';   
		export default {  
		  components: {  
		    EditableItemView
		  },  
		  data() {  
		    return {  
		      username: 'kary',
		      email: '834479572@qq.com',
			  phone: '18923208090',
			  birthday: '2003-05-28',
			  team: '英语每日练小组'
		      // 其他数据...  
		    };  
		  },  
		  methods: {  
		    GotoUser() {  
		      uni.switchTab({  
		        url: '../user/user'  
		      });  
		    }  
		    // 如果还有其他方法，可以在这里继续添加  
		  }  
		  // 注意这里没有多余的闭合大括号 "}"  
		};  
</script>

<style>
	body{
		background-color: white;
	}
	.container {
		display: flex;
		width: 100%;
		height: 100%;
		
	}

	.back-icon {
		width: 2rem;
		/*图标宽度 */
		height: 2rem;
		/*图标高度 */
		position: absolute;
		/*绝对定位 */
		top: 0.8rem;
		/*距离顶部20px */
		left: 0.5rem;
		/*距离左侧20px */
		cursor: pointer;
		/*鼠标移上去显示小手 */
	}
	
	.title{
		position: absolute;
		top: 0.8rem;
		left: 40%;
		font-size: 1.2rem;
		font-weight: 500;
		
	}
	.edit {
		position: absolute;
		top: 80px;
		/* border: 1rpx solid black; */
		width: 100%;
		height:1400rpx;
	}

	.photo {
		width: 280rpx;
		height:280rpx;
		top:80rpx;
		left: 240rpx;
		border-radius: 50%;
		box-shadow: 0px 10px 25px rgb(0, 0, 0, 0.2);
	}
	.items {
		width: 100%;
		margin-top: 180rpx;
		/* border: 1rpx solid black; */
		display: flex;
		flex-direction: column;
	}
	.item {
		width: 100%;
		display: flex; 
		justify-content: space-between; /* 内容左右分布 */
		padding: 45rpx 20rpx ; /* 上下左右边距 */
		font-size:40rpx;
		border-bottom: 1rpx solid #e3e3e3;
		/* gap: 20rpx; /* 间距 */
	}
	
	.item text {
		flex-shrink: 0; /* 防止text收缩 */
		margin-left: 50rpx;
		color:#b2b2b2;
		font-weight:500;
	}
	.item input ,
	.item span{
		width: 200rpx;
		margin-right: 70rpx;
		flex-grow: 1; /* 让input或span占据剩余空间 */
		text-align: right; /* 文本右对齐 */
		font-weight: 500;
		color:black;
	}
	.item input {
		border: none;
		font-size: 40rpx;
	}
	.item :last-child{
		border-bottom: none;
	}
</style>