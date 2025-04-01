<template>
	<n-card title="打开率">
		<line-chart :chart-data="openRateData" />
	</n-card>
	<n-card title="点击率">
		<line-chart :chart-data="clickRateData" />
	</n-card>
	<n-card title="退回率">
		<line-chart :chart-data="bounceRateData" />
	</n-card>
</template>

<script setup lang="ts">
import LineChart from './LineChart.vue'

// 示例数据
const timeLabels = ['00:00', '01:00', '02:00', '03:00', '04:00']
const chartValues = [100, 150, 220, 100, 130]

const openRateData = {
	labels: timeLabels,
	values: chartValues,
}

const clickRateData = {
	labels: timeLabels,
	values: chartValues,
}

const bounceRateData = {
	labels: timeLabels,
	values: chartValues,
}
</script>
