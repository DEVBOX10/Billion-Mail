<template>
	<n-input
		ref="inputRef"
		v-model:value="search"
		v-bind="$attrs"
		class="bt-search"
		:style="{ width: `${width}px` }"
		@keyup.enter="onSearch">
		<template v-if="prefix" #prefix>
			<div class="flex items-center cursor-pointer" @click="onSearch">
				<i class="i-mdi-search text-16px"></i>
			</div>
		</template>
		<template v-if="!prefix" #suffix>
			<div class="flex items-center cursor-pointer" @click="onSearch">
				<i class="i-mdi-search text-16px"></i>
			</div>
		</template>
	</n-input>
</template>

<script lang="ts" setup>
import { InputInst } from 'naive-ui'

interface Props {
	width?: number | string
	prefix?: boolean
}

withDefaults(defineProps<Props>(), {
	width: 240,
	prefix: false,
})

const inputRef = ref<InputInst | null>(null)

const emit = defineEmits<{
	search: [val: string]
}>()

const search = defineModel<string>('value', {
	default: '',
})

const onSearch = () => {
	emit('search', search.value)
}

defineExpose({
	focus: () => {
		inputRef.value?.focus()
	},
})
</script>
