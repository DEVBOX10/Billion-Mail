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
			loading: 'Creating Domain, please wait...',
			successMessage: true,
		},
	})
}

export const updateDomain = (params: DomainParams) => {
	return instance.post('/domains/update', params, {
		fetchOptions: {
			loading: 'Updating Domain, please wait...',
			successMessage: true,
		},
	})
}

export const deleteDomain = (params: { domain: string }) => {
	return instance.post('/domains/delete', params, {
		fetchOptions: {
			loading: 'Deleting Domain, please wait...',
			successMessage: true,
		},
	})
}

export const setSsl = (params: { domain: string; certificate: string; key: string }) => {
	return instance.post('/domains/set_ssl', params, {
		fetchOptions: {
			loading: 'Setting SSL, please wait...',
			successMessage: true,
		},
	})
}

export const getSsl = (params: { domain: string }) => {
	return instance.get('/domains/get_ssl', { params })
}

export const applyCert = (params: { domain: string }) => {
	return instance.post('/ssl/apply_cert', params, {
		fetchOptions: {
			loading: 'Applying SSL Certificate, please wait...',
			successMessage: true,
		},
	})
}
