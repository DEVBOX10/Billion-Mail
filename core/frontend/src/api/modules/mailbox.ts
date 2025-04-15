import { instance } from '@/api'
import i18n from '@/i18n'

const { t } = i18n.global

export const getMailboxList = (params: {
	page: number
	page_size: number
	domain: string | null
	keyword: string
}) => {
	return instance.get('/mailbox/list', { params })
}

type MailboxParams = {
	full_name: string
	domain: string
	password: string
	active: number
	isAdmin: number
	quota: number
}

export const createMailbox = (params: MailboxParams) => {
	return instance.post('/mailbox/create', params, {
		fetchOptions: {
			loading: t('mailbox.api.loading.creating'),
			successMessage: true,
		},
	})
}

export const updateMailbox = (params: MailboxParams) => {
	return instance.post('/mailbox/update', params, {
		fetchOptions: {
			loading: t('mailbox.api.loading.updating'),
			successMessage: true,
		},
	})
}

export const deleteMailbox = (params: { email: string }) => {
	return instance.post('/mailbox/delete', params, {
		fetchOptions: {
			loading: t('mailbox.api.loading.deleting'),
			successMessage: true,
		},
	})
}
