<template>
	<div class="flex-1">
		<div class="flex items-center">
			<div class="w-260px">
				<n-select
					v-model:value="group"
					:options="groupOptions"
					:multiple="true"
					:filterable="true"
					:max-tag-count="'responsive'"
					:loading="loading"
					:placeholder="$t('market.task.edit.recipientsPlaceholder')"
					@update:value="handleUpdateGroup">
				</n-select>
			</div>
			<div class="ml-12px">
				<n-button text type="primary" @click="handleShowCreate">
					{{ $t('common.actions.create') }}
				</n-button>
			</div>
		</div>
		<i18n-t class="mt-8px" tag="div" scope="global" keypath="market.task.edit.recipientsCount">
			<template #count>
				<b>{{ recipientsCount }}</b>
			</template>
		</i18n-t>

		<create-modal />
	</div>
</template>

<script lang="ts" setup>
import { SelectOption } from 'naive-ui'
import { isObject } from '@/utils'
import { useModal } from '@/hooks/modal/useModal'
import { getContactCount, getGroupAll } from '@/api/modules/contacts/group'
import type { Group } from '@/views/contacts/group/interface'

import CreateGroup from './CreateGroup.vue'

const group = defineModel<number[]>('value', {
	default: () => [],
})

const labels = defineModel<string[]>('label', {
	default: () => [],
})

const loading = ref(false)

const groupOptions = ref<SelectOption[]>([])

const recipientsCount = ref(0)

const [CreateModal, createModalApi] = useModal({
	component: CreateGroup,
	state: {
		refresh: () => {
			getGroupOptions()
		},
	},
})

const handleShowCreate = () => {
	createModalApi.open()
}

const handleUpdateGroup = async (val: number[]) => {
	labels.value = val.map(item => {
		const groupOption = groupOptions.value.find(option => option.value === item)
		return `${groupOption?.label || ''}`
	})
	if (val.length === 0) {
		recipientsCount.value = 0
		return
	}
	const res = await getContactCount({ group_ids: val })
	if (isObject<{ total: number }>(res)) {
		recipientsCount.value = res.total
	}
}

const getGroupOptions = async () => {
	try {
		loading.value = true
		const res = await getGroupAll()
		if (isObject<{ list: Group[] }>(res)) {
			groupOptions.value = res.list.map(item => {
				return {
					label: item.name,
					value: item.id,
				}
			})
		}
	} finally {
		loading.value = false
	}
}

getGroupOptions()
</script>

<style lang="scss" scoped></style>
