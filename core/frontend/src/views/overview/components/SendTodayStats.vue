<template>
	<div class="send-today-stats">
		<div class="stats-header">
			<div class="stat-item">
				<div class="stat-label">{{ $t('overview.send.successRate') }}</div>
				<div class="stat-value text-primary">{{ successRate }}%</div>
			</div>
			<div class="stat-item">
				<div class="stat-label">{{ $t('overview.send.success') }}</div>
				<div class="stat-value text-primary">{{ successCount }}</div>
			</div>
			<div class="stat-item">
				<div class="stat-label">{{ $t('overview.send.fail') }}</div>
				<div class="stat-value text-error">{{ failCount }}</div>
			</div>
		</div>
		<div class="h-180px">
			<bar-chart :chart-data="barChartData" />
		</div>
	</div>
</template>

<script setup lang="ts">
import { PropType } from 'vue'
import { formatTime } from '@/utils'
import type { SendMail } from '../interface'

import BarChart from './BarChart.vue'

const { data } = defineProps({
	data: {
		type: Object as PropType<SendMail>,
		required: true,
	},
})

const { t } = useI18n()

const successRate = computed(() => {
	return data.dashboard.delivery_rate
})

const successCount = computed(() => {
	return data.dashboard.delivered
})

const failCount = computed(() => {
	return data.dashboard.failed
})

const getChartTime = (type: string, x: number) => {
	let date = new Date()
	if (type === 'hourly') {
		date.setMinutes(0)
		date.setSeconds(0)
		date.setHours(x)
	} else if (type === 'daily') {
		date = new Date(x * 1000)
	}
	return formatTime(date)
}

// 柱状图数据
const barChartData = computed(() => {
	const failList: [string, number, string][] = []
	const successList: [string, number, string][] = []

	data.data.forEach(item => {
		failList.push([getChartTime(data.column_type, item.x), item.failed, data.column_type])
		successList.push([getChartTime(data.column_type, item.x), item.delivered, data.column_type])
	})

	return {
		fail: {
			label: t('overview.send.sendFail'),
			value: failList,
		},
		success: {
			label: t('overview.send.sendSuccess'),
			value: successList,
		},
	}
})
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
