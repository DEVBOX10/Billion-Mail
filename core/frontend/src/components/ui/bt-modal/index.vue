<template>
	<n-modal
		v-model:show="show"
		preset="dialog"
		:z-index="zIndex"
		:draggable="true"
		:auto-focus="false"
		:show-icon="false"
		:mask-closable="false"
		:internal-appear="true"
		:internal-dialog="true"
		:transform-origin="transformOrigin"
		:style="{ width: formatLength(width) }"
		@after-leave="onAfterLeave">
		<template #header>{{ title }}</template>
		<div :style="{ maxHeight: maxScrollHeight, overflow: 'auto' }">
			<slot></slot>
		</div>
		<template v-if="footer" #action>
			<n-button class="cancel-btn" color="#cbcbcb" @click="onCancel">取消</n-button>
			<n-button :type="confirmType" @click="onConfirm">{{ confirmText || '确定' }}</n-button>
		</template>
	</n-modal>
</template>

<script lang="ts" setup>
import { useWindowSize } from '@vueuse/core'
import { formatLength } from 'naive-ui/es/_utils'
import { Type } from 'naive-ui/es/button/src/interface'

const props = defineProps({
	zIndex: Number,
	onAfterLeave: Function as PropType<() => void>,
	title: {
		type: String,
		default: '',
	},
	width: {
		type: [String, Number],
		default: 'auto',
	},
	minHeight: {
		type: [Number, String],
		default: 'auto',
	},
	footer: {
		type: Boolean,
		default: true,
	},
	transformOrigin: {
		type: String as PropType<'mouse' | 'center'>,
		default: 'center',
	},
	confirmText: {
		type: String,
		default: '',
	},
	confirmType: {
		type: String as PropType<Type>,
		default: 'primary',
	},
	onCancel: {
		type: Function as PropType<() => unknown>,
		default: () => true,
	},
	onConfirm: {
		type: Function as PropType<() => void | boolean | Promise<unknown>>,
		default: () => true,
	},
})

const show = defineModel<boolean>('show')

// 窗口高度
const { height } = useWindowSize()

// 最大滚动高度
const maxScrollHeight = computed(() => {
	if (height.value) return `${height.value - 32}px`
	return 'auto'
})

const closeModal = () => {
	show.value = false
}

const onCancel = () => {
	const result = props.onCancel()
	if (result === false) return
	show.value = false
}

const onConfirm = () => {
	void Promise.resolve(props.onConfirm()).then(value => {
		if (value === false) return
		closeModal()
	})
}

const onAfterLeave = () => {
	props.onAfterLeave?.()
}
</script>

<style lang="scss" scoped>
.cancel-btn {
	--n-color-hover: #c9302c;
}
</style>
