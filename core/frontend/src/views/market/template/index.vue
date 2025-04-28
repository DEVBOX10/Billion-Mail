<template>
	<div>
		<bt-table-layout>
			<template #toolsLeft>
				<n-button type="primary" @click="handleAdd">
					{{ t('market.template.actions.add') }}
				</n-button>
			</template>
			<template #toolsRight>
				<bt-search
					v-model:value="tableParams.keyword"
					:width="240"
					:placeholder="t('market.template.search.namePlaceholder')"
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
				<add-modal />
				<preview-modal />
			</template>
		</bt-table-layout>
	</div>
</template>

<script lang="tsx" setup>
import { DataTableColumns, NButton, NFlex } from 'naive-ui'
import { useModal } from '@/hooks/modal/useModal'
import { useTableData } from '@/hooks/useTableData'
import { confirm, formatTime } from '@/utils'
import { deleteTemplate, duplicateTemplate, getTemplateList } from '@/api/modules/market/template'
import type { Template, TemplateParams } from './interface'

import TemplateAdd from './components/TemplateAdd.vue'
import TemplatePreview from './components/TemplatePreview.vue'

const { t } = useI18n()

const { tableParams, tableList, loading, tableTotal, getTableData } = useTableData<
	Template,
	TemplateParams
>({
	loading: true,
	immediate: true,
	params: {
		page: 1,
		page_size: 10,
		keyword: '',
	},
	fetchFn: getTemplateList,
})

// Table columns
const columns = ref<DataTableColumns<Template>>([
	{
		key: 'temp_name',
		title: t('market.template.columns.name'),
		minWidth: 120,
		ellipsis: {
			tooltip: true,
		},
	},
	{
		key: 'create_time',
		title: t('market.template.columns.createdAt'),
		minWidth: 140,
		render: row => formatTime(row.create_time),
	},
	{
		title: t('common.columns.actions'),
		key: 'actions',
		align: 'right',
		width: 240,
		render: row => (
			<NFlex inline={true}>
				<NButton
					type="primary"
					text={true}
					onClick={() => {
						handlePreview(row)
					}}>
					{t('common.actions.preview')}
				</NButton>
				<NButton
					type="primary"
					text={true}
					onClick={() => {
						handleEdit(row)
					}}>
					{t('common.actions.edit')}
				</NButton>
				<NButton
					type="primary"
					text={true}
					onClick={() => {
						handleDuplicate(row)
					}}>
					{t('market.template.actions.duplicate')}
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

const [AddModal, addModalApi] = useModal({
	component: TemplateAdd,
	state: {
		refresh: getTableData,
	},
})

// Handle add template
const handleAdd = () => {
	addModalApi.open()
}

const router = useRouter()

// Handle edit
const handleEdit = (row: Template) => {
	router.push({
		path: `/market/template/edit/${row.id}`,
	})
}

const [PreviewModal, previewModalApi] = useModal({
	component: TemplatePreview,
})

// Handle preview
const handlePreview = (row: Template) => {
	previewModalApi.setState({ row })
	previewModalApi.open()
}

// Handle duplicate
const handleDuplicate = async (row: Template) => {
	await duplicateTemplate({ id: row.id })
	getTableData()
}

// Handle delete
const handleDelete = (row: Template) => {
	confirm({
		title: t('market.template.delete.title'),
		content: t('market.template.delete.confirm', { name: row.temp_name }),
		confirmText: t('common.actions.delete'),
		confirmType: 'error',
		onConfirm: async () => {
			await deleteTemplate({ id: row.id })
			getTableData()
		},
	})
}
</script>

<style lang="scss" scoped>
.template-container {
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
