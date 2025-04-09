import { useUserStore } from '@/store'
import router from '@/router/router'
import loadingBar from '@/config/loadingBar'

// 路由白名单
const whitePathList = ['/login']

router.beforeEach(async (to, from, next) => {
	loadingBar.start()

	const userStore = useUserStore()
	// 用户已登录
	if (userStore.isLogin) {
		// 如果访问的是白名单中的路由，则跳转到主页
		if (whitePathList.includes(to.path)) {
			next('/')
		} else {
			next()
		}
	} else if (whitePathList.includes(to.path)) {
		// 如果访问的是白名单中的路由，则直接跳转
		next()
	} else {
		next('/login')
	}
})

router.afterEach(() => {
	loadingBar.finish()
})

export default router
