import { instance } from '@/api'

export const getServiceList = () => {
	return instance.get('/docker_api/list')
}

export const restartService = (params: { container_id: string }) => {
	return instance.post('/docker_api/restart', params, {
		fetchOptions: {
			successMessage: true,
		},
	})
}
