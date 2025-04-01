<template>
	<div class="py-16px px-24px">
		<n-card :bordered="false">
			<n-flex class="mb-16px" justify="space-between">
				<div>
					<n-button type="primary" @click="handleAddDomain">添加域名</n-button>
				</div>
				<div>
					<n-button class="refresh-btn" @click="refreshDomainRecord">刷新域名记录</n-button>
				</div>
			</n-flex>
			<n-data-table :columns="columns" :data="domainData"> </n-data-table>
		</n-card>
	</div>
</template>

<script lang="tsx" setup>
import { DataTableColumns, NFlex, NButton, NSwitch } from 'naive-ui'

// 域名数据
const domainData = ref([
	{
		name: 'lotkfc.cn',
		quota: '10.00 GB',
		mailboxes: 50,
		defaultSize: '5.00 GB',
		catchAll: false,
		ssl: { status: '过期于82天', type: 'warning' },
		id: 1,
	},
	{
		name: 'aaa.text',
		quota: '5.00 GB',
		mailboxes: 50,
		defaultSize: '5.00 GB',
		catchAll: false,
		ssl: { status: '未设置', type: 'error' },
		id: 2,
	},
	{
		name: 'aapanel.store',
		quota: '5.00 GB',
		mailboxes: 50000,
		defaultSize: '5.00 GB',
		catchAll: false,
		ssl: { status: '未设置', type: 'error' },
		id: 3,
	},
	{
		name: 'lootk.cn',
		quota: '20.00 MB',
		mailboxes: 50000,
		defaultSize: '5.00 GB',
		catchAll: false,
		ssl: { status: '过期于82天', type: 'warning' },
		id: 4,
	},
])

// 表格列定义
const columns = ref<DataTableColumns>([
	{
		key: 'name',
		title: '域名',
	},
	{
		title: '配额',
		key: 'quota',
	},
	{
		title: '邮箱数',
		key: 'mailboxes',
	},
	{
		title: '默认邮箱大小',
		key: 'defaultSize',
	},
	{
		title: 'Catch All',
		key: 'catchAll',
		render: () => <NSwitch size="small"></NSwitch>,
	},
	{
		key: 'ssl',
		title: 'SSL',
	},
	{
		title: '操作',
		key: 'actions',
		align: 'right',
		render: () => (
			<NFlex inline={true}>
				<NButton
					type="primary"
					text={true}
					onClick={() => {
						handleDNSRecord()
					}}>
					DNS记录
				</NButton>
				<NButton
					type="primary"
					text={true}
					onClick={() => {
						handleEdit()
					}}>
					编辑
				</NButton>
				<NButton
					type="error"
					text={true}
					onClick={() => {
						handleDelete()
					}}>
					删除
				</NButton>
			</NFlex>
		),
	},
])

// 处理添加域名
const handleAddDomain = () => {}

// 刷新域名记录
const refreshDomainRecord = () => {}

// 处理DNS记录
const handleDNSRecord = () => {}

// 处理编辑
const handleEdit = () => {}

// 处理删除
const handleDelete = () => {}
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
