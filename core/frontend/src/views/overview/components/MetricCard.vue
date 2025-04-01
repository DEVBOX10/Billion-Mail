<template>
	<n-card class="metric-card" :bordered="false">
		<div class="title">{{ title }}</div>
		<div class="value">{{ value }}%</div>
	</n-card>
</template>

<script setup lang="ts">
defineProps({
	title: {
		type: String,
		default: '',
	},
	value: {
		type: Number,
		default: 0,
	},
})
</script>

<style scoped>
.metric-card {
	--n-padding-top: 24px;
	--n-padding-bottom: 24px;
	--n-text-color: #1d2630;
}

.title {
	margin-bottom: 15px;
	font-size: 18px;
	font-weight: 400;
}

.value {
	font-size: 20px;
	font-weight: 300;
}
</style>
