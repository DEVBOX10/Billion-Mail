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

/**
 * @description 获取验证码
 * @param params
 * @returns
 */
export const getValidateCode = () => {
	return instance.get('/get_validate_code')
}
