<template>
	<div class="drag-icon">
		<i class="i-mdi-file"></i>
	</div>
	<div class="drag-text">{{ fileName }}</div>
</template>

<script lang="ts" setup>
defineProps({
	fileName: {
		type: String,
		default: '',
	},
})
</script>