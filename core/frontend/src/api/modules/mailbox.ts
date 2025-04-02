import { instance } from '@/api'

export const getMailboxList = (params: {
	page: number
	page_size: number
	domain: string
	keyword: string
}) => {
	return instance.get('/mailbox/list', { params })
}

type MailboxParams = {
	domain: string
	username: string
	password: string
	status: number
	isAdmin: number
	quota: number
}

export const createMailbox = (params: MailboxParams) => {
	return instance.post('/mailbox/create', params, {
		fetchOptions: {
			successMessage: true,
		},
	})
}

export const updateMailbox = (params: MailboxParams) => {
	return instance.post('/mailbox/update', params, {
		fetchOptions: {
			successMessage: true,
		},
	})
}

export const deleteMailbox = (params: { email: string }) => {
	return instance.post('/mailbox/delete', params, {
		fetchOptions: {
			successMessage: true,
		},
	})
}
