import { instance } from '@/api'
import i18n from '@/i18n'

const { t } = i18n.global

export const getServiceList = () => {
	return instance.get('/docker_api/list')
}

export const restartService = (params: { container_id: string }) => {
	return instance.post('/docker_api/restart', params, {
		fetchOptions: {
			loading: t('settings.api.loading.restarting'),
			successMessage: true,
		},
	})
}
