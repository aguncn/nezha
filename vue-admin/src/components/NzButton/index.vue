<template>
	<el-button :size="size" :type="type" :icon="icon" :label="label"
		:loading="loading" :plain="plain" :circle="circle" @click="handleClick">
    <span v-if="isLabel"> {{ label }} </span>
	</el-button>
</template>

<script>
	export default {
		name: 'NzButton',
		props: {
      isLabel: {
      	type: Boolean,
      	default: true
      },
			label: {
				type: String,
				default: 'Button'
			},
			icon: {
				type: String,
				default: ''
			},
			size: {
				type: String,
				default: 'mini'
			},
			type: {
				type: String,
				default: null
			},
			loading: {
				type: Boolean,
				default: false
			},
      plain: {
      	type: Boolean,
      	default: false
      },
      circle: {
      	type: Boolean,
      	default: false
      },
			disabled: {
				type: Boolean,
				default: false
			}
		},
		data() {
			return {

			}
		},
		methods: {
			handleClick: function() {
				this.$emit('click', {})
			}
		},
		mounted() {

		}
	}
</script>

<style>
</style>
