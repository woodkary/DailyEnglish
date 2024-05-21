<template>
<<<<<<< Updated upstream
	<view>
		<view class="">
			
=======
	<view class="container">
		<image class="back-icon" src="../../static/back.svg" @click="GotoUser"></image>
		<view class="edit">
			<image class="photo" src="../../static/logo.png"></image>
			<view class="items">
				<EditableItemView 
						id="username"
				      label="用户名"  
				      :value="username"  
				      placeholder="wwswad"  
				      @update:value="username = $event"  
				    />
				<EditableItemView
				id="email"
				      label="邮箱"  
				      :value="email"  
				      placeholder="3242412@xy.com"  
				      @update:value="email = $event"  
				    />
				<EditableItemView
				id="phone"
				      label="手机号"  
				      :value="phone"  
				      placeholder="12312312312"  
				      @update:value="phone = $event"  
				    />
				<EditableItemView
				 id="birthday"
				  label="生日"  
				  :value="birthday"  
				  placeholder="1999-01-01"  
				  @update:value="birthday = $event"  
				/>  
				<EditableItemView
				id="team"
				      label="团队"  
				      :value="team"  
				      placeholder="无"  
				      @update:value="team = $event"  
				    />
			</view>
>>>>>>> Stashed changes
		</view>
	</view>
</template>

<script>
<<<<<<< Updated upstream
	export default {
		data() {
			return {
				
			}
		},
		methods: {
			
		}
	}
=======
	import EditableItemView from '../../components/EditableItemView.vue';   
		export default {  
		  components: {  
		    EditableItemView
		  },  
		  data() {  
		    return {  
		      username: 'wwswad',  
		      email: '3242412@xy.com',  
			  phone: '12312312312',
			  birthday: '1999-01-01',
			  team: '无'
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
>>>>>>> Stashed changes
</script>

<style>

</style>
