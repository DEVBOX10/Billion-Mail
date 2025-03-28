import router from '@/router/router'
import loadingBar from '@/config/loadingBar'

router.beforeEach(async (to, from, next) => {
	loadingBar.start()

	next()
})

router.afterEach(() => {
	loadingBar.finish()
})

export default router
