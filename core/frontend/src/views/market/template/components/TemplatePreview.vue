<template>
	<modal :title="$t('market.template.preview.title')" width="820" :footer="false">
		<bt-preview :height="580" :value="content"></bt-preview>
	</modal>
</template>

<script lang="ts" setup>
import { useModal } from '@/hooks/modal/useModal'
import type { Template } from '../interface'
import { isObject } from '@/utils'
import { getTemplateDetails } from '@/api/modules/market/template'

const content = ref('')

const getTemplateHtml = async (id: string) => {
	const res = await getTemplateDetails({ id })
	if (isObject<Template>(res)) {
		content.value = res.content
	}
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const state = modalApi.getState<{ row: Template }>()
			const { row } = state
			if (row) {
				getTemplateHtml(`${row.id}`)
			}
		}
	},
})
</script>
