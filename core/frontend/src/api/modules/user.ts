import { instance } from '@/api'
import i18n from '@/i18n'

const { t } = i18n.global

/**
 * @description 登录
 * @param params
 * @returns
 */
export const login = (params: { username: string; password: string }) => {
	return instance.post('/login', params, {
		fetchOptions: {
			loading: t('user.api.loading.login'),
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
				loading: t('user.api.loading.logout'),
				successMessage: true,
			},
		}
	)
}
