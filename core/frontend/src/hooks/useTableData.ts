import { isObject } from '@/utils'

interface TableParams {
	page: number
	page_size: number
}

interface UseTableDataOptions<K> {
	params: K
	fetchFn?: (params: K) => Promise<unknown>
	loading?: boolean
	immediate?: boolean
}

export const useTableData = <T, K extends TableParams>(options: UseTableDataOptions<K>) => {
	const { fetchFn, params, loading = false, immediate = false } = options

	const loadingRef = ref(loading)

	const tableList = ref<T[]>([])

	const tableTotal = ref(0)

	const tableParams = ref<K>(params)

	// 获取表格数据
	const getTableData = async (resetPage = false) => {
		if (resetPage) {
			params.page = 1
		}

		if (fetchFn) {
			loadingRef.value = true
			try {
				const res = await fetchFn(tableParams.value)
				if (isObject<{ list: T[]; total: number }>(res)) {
					tableList.value = res.list
					tableTotal.value = res.total
				}
			} finally {
				loadingRef.value = false
			}
		}
	}

	// 如果设置了立即加载，则立即获取数据
	if (immediate) {
		getTableData()
	}

	return {
		loading: loadingRef,
		tableList,
		tableTotal,
		tableParams,
		getTableData,
	}
}
