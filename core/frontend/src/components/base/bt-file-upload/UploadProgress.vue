<template>
	<div class="w-80%">
		<n-progress type="line" status="success" :percentage="progress" :height="4" :show-indicator="false">
		</n-progress>
	</div>
	<div class="drag-text">{{ progress }}%</div>
</template>

<script lang="ts" setup>
defineProps({
	progress: {
		type: Number,
		default: 0,
	},
})
</script>