import { instance } from '@/api'

export const getDomainList = (params: { page: number; page_size: number; keyword: string }) => {
	return instance.get('/domains/list', { params })
}

type DomainParams = {
	domain: string
	quota: number
	mailboxes: number
}

export const createDomain = (params: DomainParams) => {
	return instance.post('/domains/create', params, {
		fetchOptions: {
			successMessage: true,
		},
	})
}

export const updateDomain = (params: DomainParams) => {
	return instance.post('/domains/update', params, {
		fetchOptions: {
			successMessage: true,
		},
	})
}

export const deleteDomain = (params: { domain: string }) => {
	return instance.post('/domains/delete', params, {
		fetchOptions: {
			successMessage: true,
		},
	})
}
