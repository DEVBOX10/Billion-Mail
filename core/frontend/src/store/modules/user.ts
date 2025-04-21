import { defineStore } from 'pinia'
import { confirm } from '@/utils'
import { logout as logoutApi } from '@/api/modules/user'
import i18n from '@/i18n'
import router from '@/router'

export default defineStore(
	'UserStore',
	() => {
		const { t } = i18n.global

		const login = ref({
			token: '', // Token
			refresh_token: '', // Refresh Token
			ttl: 0, // Token valid time/second
			expire: 0, // Token expiration time
		})

		/**
		 * @description Determine if the user is logged in
		 */
		const isLogin = computed(() => {
			return login.value.token && login.value.expire > Date.now()
		})

		/**
		 * @description Set user login information
		 * @param userVal
		 */
		const setLoginInfo = (userVal: { token: string; refresh_token: string; ttl: number }) => {
			login.value.token = userVal.token
			login.value.refresh_token = userVal.refresh_token
			login.value.ttl = userVal.ttl
			login.value.expire = userVal.ttl * 1000 + Date.now()
		}

		const resetLoginInfo = () => {
			login.value.token = ''
			login.value.refresh_token = ''
			login.value.ttl = 0
			login.value.expire = 0
		}

		const logout = () => {
			confirm({
				title: t('user.logout.title'),
				content: t('user.logout.content'),
				onConfirm: async () => {
					await logoutApi()
					resetLoginInfo()
					router.push('/login')
				},
			})
		}

		return {
			login,
			isLogin,
			logout,
			setLoginInfo,
			resetLoginInfo,
		}
	},
	{
		persist: {
			pick: ['login'],
		},
	}
)
