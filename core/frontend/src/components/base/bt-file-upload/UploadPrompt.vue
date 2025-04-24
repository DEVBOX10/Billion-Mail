<template>
	<div class="drag-icon">
		<i class="i-mdi-upload"></i>
	</div>
	<div class="drag-text">{{ $t('components.upload.prompt.dragText') }}</div>
	<div class="drag-suffix">{{ $t('components.upload.prompt.supportedFiles', { suffix }) }}</div>
</template>

<script lang="ts" setup>
defineProps({
	suffix: {
		type: String,
		default: '',
	},
})
</script>
