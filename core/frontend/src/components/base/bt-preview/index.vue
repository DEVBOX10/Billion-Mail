<template>
	<iframe width="100%" :height="height" :srcdoc="value" :style="{ border: 'none' }"> </iframe>
</template>

<script lang="ts" setup>
defineProps({
	value: {
		type: String,
		default: '',
	},
	height: {
		type: [String, Number],
		default: '100%',
	},
})
</script>

<style lang="scss" scoped></style>
