import { useUserStore } from '@/store'
import { routes } from '@/router/router'
import router from '@/router/router'
import loadingBar from '@/config/loadingBar'

// 路由白名单
const whitePathList = ['/login']

router.beforeEach(async (to, from, next) => {
	loadingBar.start()

	// 判断访问的路由是否存在于注册的路由中
	const routeExists = routes.some(route => route.path === to.path)
	// 如果路由不存在，则直接通过
	if (!routeExists) {
		next()
		return
	}

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
