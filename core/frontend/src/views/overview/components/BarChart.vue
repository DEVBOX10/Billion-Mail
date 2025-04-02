<template>
	<bt-charts :options="chartOptions" />
</template>

<script setup lang="ts">
const props = defineProps({
	chartData: {
		type: Object,
		required: true,
	},
})

const chartOptions = computed(() => {
	return {
		grid: {
			top: '4%',
			left: '4%',
			right: '4%',
			bottom: '0',
			containLabel: true,
		},
		xAxis: {
			type: 'category',
			data: props.chartData.labels,
			axisTick: {
				alignWithLabel: true,
			},
		},
		yAxis: {
			type: 'value',
		},
		series: [
			{
				type: 'bar',
				barWidth: '60%',
				data: props.chartData.values,
				itemStyle: {
					color: '#6979F8',
				},
			},
		],
	}
})
</script>
