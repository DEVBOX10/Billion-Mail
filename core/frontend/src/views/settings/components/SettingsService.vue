<template>
	<div class="settings-title">Service Status</div>
	<n-data-table :loading="loading" :data="serviceList" :columns="columns"></n-data-table>
</template>

<script lang="tsx" setup>
import { NButton, NTag, DataTableColumns } from 'naive-ui'
import { isArray } from '@/utils'
import { getServiceList, restartService } from '@/api/modules/settings'
import { DockerService } from '../types/common'

const loading = ref(false)

const serviceList = ref<DockerService[]>([])

const columns = ref<DataTableColumns<DockerService>>([
	{
		key: 'Names',
		title: 'Service Name',
		render: row => {
			return row.Names[0]
		},
	},
	{
		key: 'State',
		title: 'Status',
		render: row => {
			return (
				<NTag size="small" bordered={false} type={row.State === 'running' ? 'success' : 'error'}>
					{row.State === 'running' ? 'Running' : 'Stopped'}
				</NTag>
			)
		},
	},
	{
		key: 'action',
		title: '操作',
		align: 'right',
		render: row => (
			<>
				<NButton
					type="primary"
					text
					onClick={() => {
						handleRestart(row)
					}}>
					重启
				</NButton>
			</>
		),
	},
])

const handleRestart = async (row: DockerService) => {
	await restartService({ container_id: row.Id })
	getList()
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
