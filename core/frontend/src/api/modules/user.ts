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

/**
 * @description 退出登录
 * @param params
 * @returns
 */
export const logout = () => {
	return instance.post(
		'/logout',
		{},
		{
			fetchOptions: {
				loading: 'Logging out, please wait...',
				successMessage: true,
			},
		}
	)
}
