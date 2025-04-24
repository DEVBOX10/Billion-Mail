import { instance } from '@/api'
import type { TemplateParams } from '@/views/market/template/interface'

export function getTemplateList(params: TemplateParams) {
	return instance.get('/email_template/list', { params })
}

export function getTemplateAll() {
	return instance.get('/email_template/all')
}

export function getTemplateDetails(params: { id: string }) {
	return instance.get('/email_template/get', { params })
}

interface TemplateAddForm {
	temp_name: string
	add_type: number
	file_data: string
	content: string
	render: string
}

export function addTemplate(params: TemplateAddForm) {
	return instance.post('/email_template/create', params, {
		fetchOptions: {
			loading: 'Creating template, please wait...',
			successMessage: true,
		},
	})
}

interface TemplateEditForm {
	id: number
	temp_name: string
	content: string
}

export function updateTemplate(params: TemplateEditForm) {
	return instance.post('/email_template/update', params, {
		fetchOptions: {
			loading: 'Updating template, please wait...',
			successMessage: true,
		},
	})
}

export function duplicateTemplate(params: { id: number }) {
	return instance.post('/email_template/copy', params, {
		fetchOptions: {
			loading: 'Copying template, please wait...',
			successMessage: true,
		},
	})
}

export function deleteTemplate(params: { id: number }) {
	return instance.post('/email_template/delete', params, {
		fetchOptions: {
			loading: 'Deleting template, please wait...',
			successMessage: true,
		},
	})
}
