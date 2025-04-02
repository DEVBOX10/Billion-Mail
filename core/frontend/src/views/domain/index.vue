<template>
	<div class="py-16px px-24px">
		<n-card :bordered="false">
			<bt-table-layout>
				<template #toolsLeft>
					<n-button type="primary" @click="handleAddDomain">添加域名</n-button>
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
					<catch-modal />
					<ssl-modal />
					<dns-modal />
				</template>
			</bt-table-layout>
		</n-card>
	</div>
</template>

<script lang="tsx" setup>
import { DataTableColumns, NFlex, NButton } from 'naive-ui'
import { getByteUnit } from '@/utils'
import { useModal } from '@/hooks/modal/useModal'
import { useTableData } from '@/hooks/useTableData'
import { deleteDomain, getDomainList } from '@/api/modules/domain'
import type { MailDomain, MailDomainParams } from './interface'

import DomainForm from './components/DomainForm.vue'
import DomainCatch from './components/DomainCatch.vue'
import DomainSsl from './components/DomainSsl.vue'
import DomainDns from './components/DomainDns.vue'

const { tableParams, tableList, loading, tableTotal, getTableData } = useTableData<
	MailDomain,
	MailDomainParams
>({
	immediate: true,
	params: {
		page: 1,
		page_size: 10,
		keyword: '',
	},
	fetchFn: getDomainList,
})

// Table columns
const columns = ref<DataTableColumns<MailDomain>>([
	{
		key: 'domain',
		title: '域名',
		minWidth: 130,
		ellipsis: {
			tooltip: true,
		},
	},
	{
		title: 'quota',
		key: 'quota',
		render: row => getByteUnit(row.quota),
	},
	{
		key: 'mailboxes',
		title: 'Mailboxes',
	},
	{
		key: 'mailbox_quota',
		title: 'Default mailbox size',
		render: row => getByteUnit(row.mailbox_quota),
	},
	// {
	// 	key: 'catch_all',
	// 	title: 'Catch All',
	// 	render: row => (
	// 		<NSwitch
	// 			value={row.catch_all}
	// 			size="small"
	// 			onUpdateValue={(val: boolean) => {
	// 				if (val) {
	// 					handleOpenCatch(row)
	// 				} else {
	// 					handleCloseCatch(row)
	// 				}
	// 			}}
	// 		/>
	// 	),
	// },
	// {
	// 	key: 'ssl',
	// 	title: 'SSL',
	// 	render: row => {
	// 		if (row.ssl_status) {
	// 			const ssl = row.ssl_info
	// 			return (
	// 				<NButton
	// 					type={ssl.endtime < 0 ? 'error' : 'primary'}
	// 					text
	// 					onClick={() => {
	// 						handleShowSsl(row)
	// 					}}>
	// 					{ssl.endtime < 0 ? '已过期' : `剩余${ssl.endtime}天`}
	// 				</NButton>
	// 			)
	// 		}
	// 		return (
	// 			<NButton
	// 				type="warning"
	// 				text
	// 				onClick={() => {
	// 					handleShowSsl(row)
	// 				}}>
	// 				未设置
	// 			</NButton>
	// 		)
	// 	},
	// },
	{
		title: '操作',
		key: 'actions',
		align: 'right',
		render: row => (
			<NFlex inline={true}>
				<NButton
					type="primary"
					text={true}
					onClick={() => {
						handleDNSRecord(row)
					}}>
					DNS记录
				</NButton>
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
	component: DomainForm,
	state: {
		isEdit: false,
		refresh: getTableData,
	},
})

// Handle add domain
const handleAddDomain = () => {
	formModalApi.setState({
		row: null,
		isEdit: false,
	})
	formModalApi.open()
}

const [CatchModal] = useModal({
	component: DomainCatch,
})

// Handle open catch all
// const handleOpenCatch = (row: MailDomain) => {
// 	catchModalApi.setState({ row })
// 	catchModalApi.open()
// }

// Handle close catch all
// const handleCloseCatch = (row: MailDomain) => {
// 	console.log(row)
// }

const [SslModal] = useModal({
	component: DomainSsl,
})

// Handle show ssl
// const handleShowSsl = (row: MailDomain) => {
// 	sslModalApi.setState({ row })
// 	sslModalApi.open()
// }

const [DnsModal, dnsModalApi] = useModal({
	component: DomainDns,
	state: {},
})

// Handle DNS records
const handleDNSRecord = (row: MailDomain) => {
	dnsModalApi.setState({ row })
	dnsModalApi.open()
}

// Handle edit
const handleEdit = (row: MailDomain) => {
	formModalApi.setState({
		row,
		isEdit: true,
	})
	formModalApi.open()
}

// Handle delete
const handleDelete = async (row: MailDomain) => {
	await deleteDomain({ domain: row.domain })
	getTableData()
}
</script>

<style lang="scss" scoped>
.domain-container {
	padding: 16px;
	background-color: #fff;
	border-radius: 3px;

	.action-bar {
		display: flex;
		justify-content: space-between;
		margin-bottom: 16px;
	}
}
</style>
