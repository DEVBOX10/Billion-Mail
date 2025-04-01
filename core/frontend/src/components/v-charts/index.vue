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
	// 控制是否合并新旧配置
	notMerge: {
		type: Boolean,
		default: false,
	},
	// 更新时是否不显示过渡动画
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
		// 使用主题并添加渲染器选项，提高渲染性能
		chartInstance = echarts.init(chartRef.value, props.theme, {
			renderer: 'canvas',
			useDirtyRect: true, // 使用脏矩形渲染提高性能
		})

		// 初始设置选项
		updateChart()

		// 使用ResizeObserver代替window事件，提高性能
		resizeObserver = new ResizeObserver(() => {
			requestAnimationFrame(() => {
				chartInstance?.resize()
			})
		})

		resizeObserver.observe(chartRef.value)
	}
}

// 更新图表的方法，与初始化分离
const updateChart = () => {
	if (!chartInstance || !props.options) return

	// 使用notMerge和lazyUpdate提高更新性能
	chartInstance.setOption(props.options, {
		notMerge: props.notMerge,
		lazyUpdate: props.lazyUpdate,
	})
}

onMounted(() => {
	initChart()
})

onUnmounted(() => {
	// 清理所有资源
	resizeObserver?.disconnect()
	if (chartInstance) {
		chartInstance.dispose()
		chartInstance = null
	}
})

// 使用深度监听优化图表更新
watch(
	() => props.options,
	() => {
		updateChart()
	},
	{ deep: true }
)

// 主题变化时重新初始化图表
watch(
	() => props.theme,
	() => {
		if (chartInstance) {
			chartInstance.dispose()
			initChart()
		}
	}
)

// 导出图表实例，使父组件可以访问更多方法
defineExpose({
	getChartInstance: () => chartInstance,
	resize: () => chartInstance?.resize(),
})
</script>
