<template>
	<div class="h-180px">
		<bt-charts :options="chartOptions" />
	</div>
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
		title: {
			left: 'left',
		},
		grid: {
			top: '8%',
			left: '2%',
			right: '2%',
			bottom: '0',
			containLabel: true,
		},
		xAxis: {
			type: 'category',
			data: props.chartData.labels,
			boundaryGap: false,
		},
		yAxis: {
			type: 'value',
		},
		series: [
			{
				data: props.chartData.values,
				type: 'line',
				smooth: true,
				areaStyle: {
					color: {
						type: 'linear',
						x: 0,
						y: 0,
						x2: 0,
						y2: 1,
						colorStops: [
							{
								offset: 0,
								color: 'rgba(105, 121, 248, 0.3)',
							},
							{
								offset: 1,
								color: 'rgba(105, 121, 248, 0.1)',
							},
						],
					},
				},
				itemStyle: {
					color: '#6979F8',
				},
				lineStyle: {
					color: '#6979F8',
				},
			},
		],
	}
})
</script>
