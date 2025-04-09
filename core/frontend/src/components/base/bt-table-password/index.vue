<template>
	<div class="flex">
		<div class="w-0 flex-1 flex items-center">
			<div class="min-w-0">
				<n-ellipsis class="max-w-full" :tooltip="{ contentStyle: { maxWidth: '200px' } }">
					{{ password }}
				</n-ellipsis>
			</div>
			<div class="ml-6px text-15px cursor-pointer" title="Show" @click="onToggle">
				<i :class="show ? 'i-mdi-eye-outline' : 'i-mdi-eye-off-outline'"></i>
			</div>
			<div class="reset ml-6px text-15px cursor-pointer text-#333" title="Copy" @click="onCopy">
				<i class="i-mdi-content-copy"></i>
			</div>
		</div>
	</div>
</template>

<script lang="ts" setup>
import { copyText } from '@/utils'

interface Props {
	value: string | number
}

const props = withDefaults(defineProps<Props>(), {
	value: '',
})

const show = ref(false)

const password = computed(() => {
	return show.value ? props.value : '**********'
})

const onToggle = () => {
	show.value = !show.value
}

const onCopy = async () => {
	copyText(`${props.value}`)
}
</script>
