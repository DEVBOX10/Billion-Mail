<template>
	<div ref="chartRef" :style="{ width: '100%', height: chartHeight }"></div>
</template>

<script setup>
// 按需引入echarts核心模块
import * as echarts from 'echarts/core'
// 按需引入图表类型，以下是常用图表示例，请根据实际需要修改
import { BarChart, LineChart } from 'echarts/charts'
// 引入提示框、标题、直角坐标系、数据集等组件
import {
	TitleComponent,
	TooltipComponent,
	GridComponent,
	DatasetComponent,
	TransformComponent,
	LegendComponent,
} from 'echarts/components'
// 引入Canvas渲染器
import { CanvasRenderer } from 'echarts/renderers'

// 注册必须的组件
echarts.use([
	TitleComponent,
	TooltipComponent,
	GridComponent,
	DatasetComponent,
	TransformComponent,
	LegendComponent,
	BarChart,
	LineChart,
	CanvasRenderer,
])

const props = defineProps({
	options: {
		type: Object,
		required: true,
	},
	height: {
		type: [String, Number],
		default: '100%',
	},
	theme: {
		type: String,
		default: 'default',
	},
	// Control whether to merge new and old configurations
	notMerge: {
		type: Boolean,
		default: false,
	},
	// Whether to not display the transition animation when updating
	lazyUpdate: {
		type: Boolean,
		default: false,
	},
})

const chartRef = ref(null)

const chartHeight = computed(() => {
	return typeof props.height === 'number' ? `${props.height}px` : props.height
})

let chartInstance = null
let resizeObserver = null

const initChart = () => {
	if (chartRef.value) {
		// Use theme and add renderer options to improve rendering performance
		chartInstance = echarts.init(chartRef.value, props.theme, {
			renderer: 'canvas',
			useDirtyRect: true, // Use dirty rectangle rendering to improve performance
		})

		// Initial set options
		updateChart()

		// Use ResizeObserver instead of window event to improve performance
		resizeObserver = new ResizeObserver(() => {
			requestAnimationFrame(() => {
				chartInstance?.resize()
			})
		})

		resizeObserver.observe(chartRef.value)
	}
}

// Update chart method, separated from initialization
const updateChart = () => {
	if (!chartInstance || !props.options) return

	// Use notMerge and lazyUpdate to improve update performance
	chartInstance.setOption(props.options, {
		notMerge: props.notMerge,
		lazyUpdate: props.lazyUpdate,
	})
}

onMounted(() => {
	initChart()
})

onUnmounted(() => {
	// Clean up all resources
	resizeObserver?.disconnect()
	if (chartInstance) {
		chartInstance.dispose()
		chartInstance = null
	}
})

// Use deep watch to optimize chart updates
watch(
	() => props.options,
	() => {
		updateChart()
	},
	{ deep: true }
)

// When the theme changes, reinitialize the chart
watch(
	() => props.theme,
	() => {
		if (chartInstance) {
			chartInstance.dispose()
			initChart()
		}
	}
)

// Export chart instance, so parent components can access more methods
defineExpose({
	getChartInstance: () => chartInstance,
	resize: () => chartInstance?.resize(),
})
</script>
