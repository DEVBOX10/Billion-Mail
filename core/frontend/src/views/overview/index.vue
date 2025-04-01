<template>
	<div class="email-analytics-container">
		<!-- 顶部筛选区域 -->
		<filter-bar />

		<!-- 主要指标卡片 -->
		<div class="metrics-cards">
			<metric-card title="已送达" :value="99" />
			<metric-card title="已打开" :value="99" />
			<metric-card title="已点击" :value="99" />
			<metric-card title="已退回" :value="99" />
		</div>

		<!-- 数据详情区域 -->
		<div class="detail-row">
			<!-- 左侧邮件提供商表格 -->
			<n-card class="provider-table-card" title="邮件提供商">
				<provider-table />
			</n-card>

			<!-- 右侧今日发送统计 -->
			<n-card class="send-today-card" title="发送统计">
				<send-today-stats />
			</n-card>
		</div>

		<!-- 底部比率图表区域 -->
		<div class="rate-charts-card">
			<rate-chart-panel />
		</div>
	</div>
</template>

<script lang="ts" setup>
import FilterBar from './components/FilterBar.vue'
import MetricCard from './components/MetricCard.vue'
import ProviderTable from './components/ProviderTable.vue'
import SendTodayStats from './components/SendTodayStats.vue'
import RateChartPanel from './components/RateChartPanel.vue'
</script>

<style lang="scss" scoped>
.email-analytics-container {
	padding: 20px;
}

.metrics-cards {
	display: flex;
	justify-content: space-between;
	gap: 20px;
	margin-bottom: 20px;
}

.detail-row {
	display: flex;
	gap: 20px;
	margin-bottom: 20px;
}

.provider-table-card {
	flex: 1;
}

.send-today-card {
	flex: 1;
}

.rate-charts-card {
	display: flex;
	gap: 20px;
}
</style>
