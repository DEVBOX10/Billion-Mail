<template>
	<n-select
		v-model:value="value"
		max-tag-count="responsive"
		:multiple="true"
		:filterable="true"
		:loading="loading"
		:options="groupOptions">
	</n-select>
</template>

<script lang="ts" setup>
import { SelectOption } from 'naive-ui'
import { isObject } from '@/utils'
import { getGroupAll } from '@/api/modules/contacts/group'
import type { Group } from '../../group/interface'

const value = defineModel<number[]>('value')

const groupOptions = ref<SelectOption[]>([])

const loading = ref(false)

const getGroupOptions = async () => {
	try {
		loading.value = true
		const res = await getGroupAll()
		if (isObject<{ list: Group[] }>(res)) {
			groupOptions.value = res.list.map(item => ({
				label: item.name,
				value: item.id,
			}))
		}
	} finally {
		loading.value = false
	}
}

getGroupOptions()
</script>
