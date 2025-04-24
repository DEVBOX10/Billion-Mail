<template>
	<modal :title="title" :width="600" :footer="false">
		<bt-table-layout>
			<template #toolsLeft>
				<!-- <n-radio-group>
					<n-radio-button>Domain</n-radio-button>
					<n-radio-button>Status</n-radio-button>
				</n-radio-group> -->
			</template>
			<template #toolsRight> </template>
			<template #table>
				<n-data-table :bordered="false" :loading="loading" :columns="columns" :data="tableList">
				</n-data-table>
			</template>
			<template #modal>
				<details-modal />
			</template>
		</bt-table-layout>
	</modal>
</template>

<script lang="tsx" setup>
import { DataTableColumns, NButton } from 'naive-ui'
import { isArray } from '@/utils'
import { useModal } from '@/hooks/modal/useModal'
import { getMailProvider } from '@/api/modules/market/task'
import type { Task, TaskStatus } from '../interface'

import TaskStatusDetails from './TaskStatusDetails.vue'

const title = ref('Error tasks')

const loading = ref(false)

const search = reactive({
	status: 0,
	task_id: 0,
})

const tableList = ref<TaskStatus[]>([])

const columns = ref<DataTableColumns<TaskStatus>>([
	{
		key: 'domain',
		title: 'Domain',
	},
	{
		key: 'count',
		title: 'Count',
	},
	{
		key: 'active',
		title: 'Operator',
		align: 'right',
		render: row => {
			return (
				<NButton
					text
					type="primary"
					onClick={() => {
						handleShowDetails(row)
					}}>
					Details
				</NButton>
			)
		},
	},
])

const [DetailsModal, detailsModalApi] = useModal({
	component: TaskStatusDetails,
})

const handleShowDetails = (row: TaskStatus) => {
	detailsModalApi.setState({ taskId: search.task_id, status: search.status, row })
	detailsModalApi.open()
}

const getTableData = async () => {
	const res = await getMailProvider(toRaw(search))
	if (isArray<TaskStatus>(res)) {
		tableList.value = res
	}
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const state = modalApi.getState<{ row: Task }>()
			const { row } = state
			if (row) {
				tableList.value = []
				search.task_id = row.id
				title.value = `Error tasks [${row.subject}]`
				getTableData()
			}
		} else {
			tableList.value = []
		}
	},
})
</script>
