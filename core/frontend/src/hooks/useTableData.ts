import { isNumber } from 'lodash-es'
import { isArray, isObject } from '@/utils'

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

	// Get table data
	const getTableData = async (resetPage = false) => {
		if (resetPage) {
			params.page = 1
		}

		if (fetchFn) {
			loadingRef.value = true
			try {
				const res = await fetchFn(tableParams.value)
				if (isObject<{ list: T[]; total: number }>(res)) {
					tableList.value = isArray(res.list) ? res.list : []
					tableTotal.value = isNumber(res.total) ? res.total : 0
				}
			} finally {
				loadingRef.value = false
			}
		}
	}

	// If immediate is set, get data immediately
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
