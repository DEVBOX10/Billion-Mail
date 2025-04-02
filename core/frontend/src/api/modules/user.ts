import { instance } from '@/api'

/**
 * @description 登录
 * @param params
 * @returns
 */
export const login = (params: { username: string; password: string }) => {
	return instance.post('/login', params, {
		fetchOptions: {
			successMessage: true,
		},
	})
}
