<template>
	<n-select
		v-model:value="domain"
		:loading="loading"
		:filterable="true"
		:options="domainOptions"
		:placeholder="t('mailbox.domain.selectPlaceholder')">
	</n-select>
</template>

<script lang="ts" setup>
import { SelectOption } from 'naive-ui'
import { isArray } from '@/utils'
import { getDomainAll } from '@/api/modules/domain'
import type { MailDomain } from '@/views/domain/interface'

const { t } = useI18n()

const { isAll } = defineProps({
	isAll: {
		type: Boolean,
		default: true,
	},
})

const domain = defineModel<string | null>('value', {
	default: () => null,
})

const loading = ref(false)

const domainOptions = ref<SelectOption[]>([])

const getDomainSelect = async () => {
	try {
		loading.value = true
		const res = await getDomainAll()
		if (isArray<MailDomain>(res)) {
			domainOptions.value = res.map(item => ({
				label: item.domain,
				value: item.domain,
			}))
			// Add all options
			if (isAll) {
				domainOptions.value.unshift({
					label: t('common.all.text'),
					value: '',
				})
				domain.value = ''
			} else {
				if (res.length > 0 && !domain.value) {
					domain.value = res[0].domain
				}
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
