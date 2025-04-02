<template>
	<n-form
		v-bind="$attrs"
		ref="formRef"
		:model="model"
		:rules="rules"
		:label-width="labelWidth"
		:label-placement="labelPlacement"
		:show-require-mark="false">
		<slot></slot>
	</n-form>
</template>

<script lang="ts" setup>
import { FormInst, FormRules, LabelPlacement } from 'naive-ui/es/form/src/interface'

interface Props {
	model?: Record<string, unknown>
	rules?: FormRules
	labelWidth?: string | number
	labelPlacement?: LabelPlacement
}

withDefaults(defineProps<Props>(), {
	labelWidth: '130',
	labelPlacement: 'left',
})

const formRef = ref<FormInst | null>(null)

defineExpose({
	validate: () => {
		return formRef.value?.validate()
	},
	restoreValidation: () => {
		return formRef.value?.restoreValidation()
	},
})
</script>
