import { isObject } from '@/utils'

interface TableParams {
	[key: string]: unknown
	page: number
	page_size: number
	keyword: string
}

interface UseTableDataOptions {
	fetchFn: (params: TableParams) => Promise<unknown>
	params?: TableParams
	immediate?: boolean
}

export const useTableData = <T>(options: UseTableDataOptions) => {
	const { fetchFn, immediate = false, params = { page: 1, page_size: 10, keyword: '' } } = options

	const loading = ref(false)

	const tableList = ref<T[]>([])

	const tableTotal = ref(0)

	const tableParams = ref(params)

	// 获取表格数据
	const getTableData = async (resetPage = false) => {
		if (resetPage) {
			params.page = 1
		}

		loading.value = true
		try {
			const res = await fetchFn(tableParams.value)
			if (isObject<{ list: T[]; total: number }>(res)) {
				tableList.value = res.list
				tableTotal.value = res.total
			}
		} finally {
			loading.value = false
		}
	}

	// 如果设置了立即加载，则立即获取数据
	if (immediate) {
		getTableData()
	}

	return {
		loading,
		tableList,
		tableTotal,
		tableParams,
		getTableData,
	}
}
