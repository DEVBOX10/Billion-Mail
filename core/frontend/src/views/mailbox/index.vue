<template>
	<div class="p-24px">
		<div class="mb-20px text-24px font-bold">{{ t('layout.menu.mailboxes') }}</div>
		<bt-table-layout>
			<template #toolsLeft>
				<n-button type="primary" @click="handleAdd">{{ t('mailbox.actions.add') }}</n-button>
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
					:placeholder="t('mailbox.search.usernamePlaceholder')"
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
					@refresh="getTableData">
				</bt-table-page>
			</template>
			<template #modal>
				<form-modal />
			</template>
		</bt-table-layout>
	</div>
</template>

<script lang="tsx" setup>
import { DataTableColumns, NButton, NFlex, NSwitch } from 'naive-ui'
import { confirm, getByteUnit } from '@/utils'
import { useModal } from '@/hooks/modal/useModal'
import { useCopy } from '@/hooks/useCopy'
import { useTableData } from '@/hooks/useTableData'
import { deleteMailbox, getMailboxList, updateMailbox } from '@/api/modules/mailbox'
import { MailBox, MailBoxParams } from './interface'

import TablePassword from '@/components/base/bt-table-password/index.vue'
import DomainSelect from './components/DomainSelect.vue'
import MailboxForm from './components/MailboxForm.vue'

const { t } = useI18n()

const { copyText } = useCopy()

const { tableParams, tableList, loading, tableTotal, getTableData } = useTableData<
	MailBox,
	MailBoxParams
>({
	loading: true,
	immediate: false,
	params: {
		page: 1,
		page_size: 10,
		domain: null,
		keyword: '',
	},
	fetchFn: getMailboxList,
})

// Table columns
const columns = ref<DataTableColumns<MailBox>>([
	{
		key: 'username',
		title: t('mailbox.columns.username'),
		minWidth: 120,
		ellipsis: {
			tooltip: true,
		},
	},
	{
		key: 'password',
		title: t('mailbox.columns.password'),
		width: '18%',
		minWidth: 120,
		render: row => <TablePassword value={row.password || `--`} />,
	},
	{
		key: 'login',
		title: t('mailbox.columns.loginInfo'),
		render: row => {
			return (
				<NButton
					text
					type="primary"
					onClick={() => {
						copyText(
							t('mailbox.loginInfo.template', {
								webmail: window.location.origin + '/roundcube',
								username: row.username,
								password: row.password,
							})
						)
					}}>
					{t('common.actions.copy')}
				</NButton>
			)
		},
	},
	{
		key: 'quota',
		title: t('mailbox.columns.quota'),
		width: '14%',
		minWidth: 120,
		render: row => getByteUnit(row.quota),
	},
	{
		key: 'is_admin',
		title: t('mailbox.columns.type'),
		width: '14%',
		minWidth: 100,
		render: row => {
			return row.is_admin === 1 ? t('mailbox.userType.admin') : t('mailbox.userType.general')
		},
	},
	{
		key: 'status',
		title: t('mailbox.columns.status'),
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
		title: t('common.columns.actions'),
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
					{t('common.actions.edit')}
				</NButton>
				<NButton
					type="error"
					text={true}
					onClick={() => {
						handleDelete(row)
					}}>
					{t('common.actions.delete')}
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

const handleDelete = (row: MailBox) => {
	confirm({
		title: t('mailbox.delete.title'),
		content: t('mailbox.delete.confirm', { name: row.username }),
		confirmText: t('common.actions.delete'),
		confirmType: 'error',
		onConfirm: async () => {
			await deleteMailbox({ email: row.username })
			getTableData()
		},
	})
}
</script>
