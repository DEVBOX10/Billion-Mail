<template>
	<n-select
		v-model:value="domain"
		:loading="loading"
		:filterable="true"
		:options="domainOptions"
		placeholder="请选择域名">
	</n-select>
</template>

<script lang="ts" setup>
import { SelectOption } from 'naive-ui'
import { isObject } from '@/utils'
import { getDomainList } from '@/api/modules/domain'
import { MailDomain } from '@/views/domain/interface'

const domain = defineModel<string | null>('value', {
	default: () => null,
})

const loading = ref(false)

const domainOptions = ref<SelectOption[]>([])

const getDomainSelect = async () => {
	try {
		loading.value = true
		const res = await getDomainList({ page: 1, page_size: 10000, keyword: '' })
		if (isObject<{ list: MailDomain[] }>(res)) {
			domainOptions.value = res.list.map(item => ({
				label: item.domain,
				value: item.domain,
			}))
			if (res.list.length > 0 && !domain.value) {
				domain.value = res.list[0].domain
			}
		}
	} finally {
		loading.value = false
	}
}

getDomainSelect()
</script>

<style scoped>
.domain-select {
	width: 100%;
}
</style>
