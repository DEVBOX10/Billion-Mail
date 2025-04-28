<template>
	<div class="h-180px">
		<bt-charts :options="chartOptions" />
	</div>
</template>

<script setup lang="ts">
import { formatTime } from '@/utils'

const { chartName, chartColor, dateType, chartData } = defineProps({
	chartName: {
		type: String,
		required: true,
	},
	chartColor: {
		type: String,
		required: true,
	},
	dateType: {
		type: String,
		required: true,
	},
	chartData: {
		type: Array as PropType<[string, number][]>,
		required: true,
	},
})

const chartOptions = computed(() => {
	return {
		tooltip: {
			trigger: 'axis',
			order: 'seriesDesc',
			axisPointer: {
				type: 'shadow',
			},
		},
		grid: {
			top: '16%',
			left: '2%',
			right: '2%',
			bottom: '2%',
			containLabel: true,
		},
		xAxis: {
			type: 'category',
			axisLabel: {
				formatter: (val: string) => {
					if (dateType === 'hourly') {
						return formatTime(val, 'HH:mm')
					}
					return formatTime(val, 'yyyy-MM-dd')
				},
			},
		},
		yAxis: {
			name: '%',
			type: 'value',
			splitLine: {
				show: true,
				lineStyle: {
					type: 'dashed',
					width: 1,
					color: '#ebeef5',
				},
			},
			max: ({ max }: { max: number }) => {
				return max > 100 ? 100 : max
			},
		},
		series: [
			{
				name: chartName,
				type: 'line',
				data: chartData,
				itemStyle: {
					color: chartColor,
				},
				smooth: false,
				showSymbol: false,
				sampling: 'average',
			},
		],
	}
})
</script>
