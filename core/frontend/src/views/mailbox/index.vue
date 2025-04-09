<template>
	<div class="py-16px px-24px">
		<n-card :bordered="false">
			<bt-table-layout>
				<template #toolsLeft>
					<n-button type="primary" @click="handleAdd">添加邮箱</n-button>
				</template>
				<template #toolsRight>
					<div class="w-220px">
						<domain-select
							v-model:value="tableParams.domain"
							@update:value="() => getTableData(true)">
						</domain-select>
					</div>
					<bt-search
						v-model:value="tableParams.keyword"
						width="280"
						placeholder="搜索用户名"
						@search="() => getTableData(true)">
					</bt-search>
				</template>
				<template #table>
					<n-data-table :loading="loading" :columns="columns" :data="tableList"> </n-data-table>
				</template>
				<template #pageRight>
					<bt-table-page
						v-model:page="tableParams.page"
						v-model:page-size="tableParams.page_size"
						:item-count="tableTotal"
						@refresh="() => getTableData(true)">
					</bt-table-page>
				</template>
				<template #modal>
					<form-modal />
				</template>
			</bt-table-layout>
		</n-card>
	</div>
</template>

<script lang="tsx" setup>
import { DataTableColumns, NButton, NFlex, NSwitch } from 'naive-ui'
import { getByteUnit } from '@/utils'
import { useModal } from '@/hooks/modal/useModal'
import { useTableData } from '@/hooks/useTableData'
import { deleteMailbox, getMailboxList, updateMailbox } from '@/api/modules/mailbox'
import { MailBox, MailBoxParams } from './interface'

import TablePassword from '@/components/base/bt-table-password/index.vue'
import DomainSelect from './components/DomainSelect.vue'
import MailboxForm from './components/MailboxForm.vue'

const { tableParams, tableList, loading, tableTotal, getTableData } = useTableData<
	MailBox,
	MailBoxParams
>({
	loading: true,
	immediate: false,
	params: {
		page: 1,
		page_size: 10,
		domain: '',
		keyword: '',
	},
	fetchFn: getMailboxList,
})

// Table columns
const columns = ref<DataTableColumns<MailBox>>([
	{
		key: 'username',
		title: 'Username',
		minWidth: 120,
		ellipsis: {
			tooltip: true,
		},
	},
	{
		key: 'password',
		title: 'Password',
		width: '18%',
		minWidth: 120,
		render: (row, index) => <TablePassword value={row.password || `--${index + 1}`} />,
	},
	{
		key: 'quota',
		title: 'Quota',
		width: '14%',
		minWidth: 120,
		render: row => getByteUnit(row.quota),
	},
	{
		key: 'is_admin',
		title: 'Type',
		width: '14%',
		minWidth: 100,
		render: row => {
			return row.is_admin === 1 ? 'Admin' : 'General user'
		},
	},
	{
		key: 'status',
		title: 'Status',
		width: '12%',
		minWidth: 80,
		render: row => {
			return (
				<NSwitch
					value={row.active}
					checked-value={1}
					unchecked-value={0}
					size="small"
					onUpdateValue={val => {
						handleStatusChange(row, val)
					}}
				/>
			)
		},
	},
	{
		title: '操作',
		key: 'actions',
		align: 'right',
		width: 120,
		render: row => (
			<NFlex inline={true}>
				<NButton
					type="primary"
					text={true}
					onClick={() => {
						handleEdit(row)
					}}>
					编辑
				</NButton>
				<NButton
					type="error"
					text={true}
					onClick={() => {
						handleDelete(row)
					}}>
					删除
				</NButton>
			</NFlex>
		),
	},
])

const [FormModal, formModalApi] = useModal({
	component: MailboxForm,
	state: {
		isEdit: false,
		refresh: getTableData,
	},
})

const handleAdd = () => {
	formModalApi.setState({ isEdit: false, row: null })
	formModalApi.open()
}

const handleStatusChange = async (row: MailBox, val: number) => {
	await updateMailbox({
		full_name: row.full_name,
		domain: row.domain,
		password: row.password,
		quota: row.quota,
		isAdmin: row.is_admin,
		active: val,
	})
	row.active = val
}

const handleEdit = (row: MailBox) => {
	formModalApi.setState({ isEdit: true, row })
	formModalApi.open()
}

const handleDelete = async (row: MailBox) => {
	await deleteMailbox({ email: row.username })
	getTableData()
}
</script>
