<template>
	<div class="settings-title">{{ t('settings.service.title') }}</div>
	<n-data-table :loading="loading" :data="serviceList" :columns="columns"></n-data-table>
</template>

<script lang="tsx" setup>
import { NButton, NTag, DataTableColumns } from 'naive-ui'
import { confirm, isArray } from '@/utils'
import { getServiceList, restartService } from '@/api/modules/settings'
import { DockerService } from '../types/common'

const { t } = useI18n()
const loading = ref(false)
const serviceList = ref<DockerService[]>([])

const columns = ref<DataTableColumns<DockerService>>([
	{
		key: 'Names',
		title: t('settings.service.columns.name'),
		render: row => {
			return row.Names[0]
		},
	},
	{
		key: 'State',
		title: t('settings.service.columns.status'),
		render: row => {
			return (
				<NTag size="small" bordered={false} type={row.State === 'running' ? 'success' : 'error'}>
					{row.State === 'running'
						? t('settings.service.status.running')
						: t('settings.service.status.stopped')}
				</NTag>
			)
		},
	},
	{
		key: 'action',
		title: t('common.columns.actions'),
		align: 'right',
		render: row => (
			<>
				<NButton
					type="primary"
					text
					onClick={() => {
						handleRestart(row)
					}}>
					{t('common.actions.restart')}
				</NButton>
			</>
		),
	},
])

const handleRestart = async (row: DockerService) => {
	confirm({
		title: t('settings.service.restart.title'),
		content: t('settings.service.restart.confirm', { name: row.Names[0] }),
		confirmText: t('common.actions.confirm'),
		onConfirm: async () => {
			await restartService({ container_id: row.Id })
			getList()
		},
	})
}

const getList = async () => {
	try {
		loading.value = true
		const res = await getServiceList()
		if (isArray<DockerService>(res)) {
			serviceList.value = res
		}
	} finally {
		loading.value = false
	}
}

getList()
</script>

<style lang="scss" scoped>
.settings-title {
	margin-bottom: 16px;
	font-size: 16px;
	font-weight: 600;
	color: #666666;
}
</style>
