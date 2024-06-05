<template>
	<input :placeholder="placeholderText" type="text" class="inner-input" :class="{ 'has-focus': focused }"
		@focus="handleFocus" @blur="handleBlur" required />
</template>

<script>
	export default {
		name: 'InputContainer',
		props: {
			placeholderText: {
				type: String,
				default: 'Username'
			}
		},
		data() {
			return {
				focused: false
			};
		},
		methods: {
			handleFocus() {
				this.focused = true;
			},
			handleBlur() {
				this.focused = false;
			}
		}
	};
</script>

<style scoped>
	.inner-input {
		line-height: 30px;
		border: 2px solid transparent;
		border-bottom-color: #777;
		padding: .2rem 0;
		outline: none;
		background-color: transparent;
		color: #0d0c22;
		transition: .3s cubic-bezier(0.645, 0.045, 0.355, 1);
	}

	.has-focus,
	.inner-input:hover {
		outline: none;
		padding: .2rem 1rem;
		border-radius: 1rem;
		border-color: #7a9cc6;
	}
	
	
</style>