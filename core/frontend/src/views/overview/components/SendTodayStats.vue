<template>
	<div class="send-today-stats">
		<div class="stats-header">
			<div class="stat-item">
				<div class="stat-label">成功率</div>
				<div class="stat-value">{{ successRate }}</div>
			</div>
			<div class="stat-item">
				<div class="stat-label">成功</div>
				<div class="stat-value">{{ successCount }}</div>
			</div>
			<div class="stat-item">
				<div class="stat-label">失败</div>
				<div class="stat-value">{{ failCount }}</div>
			</div>
		</div>
		<div class="h-180px">
			<bar-chart :chart-data="barChartData" />
		</div>
	</div>
</template>

<script setup lang="ts">
import BarChart from './BarChart.vue'

// 统计数据
const successRate = ref('99%')
const successCount = ref('9999')
const failCount = ref('9')

// 柱状图数据
const barChartData = {
	labels: ['00:00', '01:00', '02:00', '03:00', '04:00'],
	values: [100, 150, 220, 100, 150],
}
</script>

<style lang="scss" scoped>
.stats-header {
	display: flex;
	justify-content: space-around;
	margin-bottom: 20px;

	.stat-item {
		text-align: center;
	}

	.stat-label {
		font-size: 14px;
		color: #666;
		margin-bottom: 5px;
	}

	.stat-value {
		font-size: 18px;
		font-weight: bold;
	}
}
</style>
