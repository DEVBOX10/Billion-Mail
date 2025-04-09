import { defineStore } from 'pinia'
import { confirm } from '@/utils'
import { logout as logoutApi } from '@/api/modules/user'
import router from '@/router'

export default defineStore(
	'UserStore',
	() => {
		const login = ref({
			token: '', // Token
			refresh_token: '', // 刷新Token
			ttl: 0, // Token有效时间/秒
			expire: 0, // Token到期时间
		})

		/**
		 * @description 判断用户是否登录
		 */
		const isLogin = computed(() => {
			return login.value.token && login.value.expire > Date.now()
		})

		/**
		 * @description 设置用户登录信息
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
				title: '提示',
				content: '是否确认退出登录？',
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
