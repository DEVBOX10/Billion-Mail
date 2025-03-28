import { get } from 'lodash-es'
import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import { is } from '@/utils/is'
import { Layout } from '@/router/constant'

// 从模块中导入路由
const modules = import.meta.webpackContext('./modules', {
	// 是否搜索子目录
	recursive: false,
	regExp: /\.ts$/,
})

// 模块路由
export const menuList: RouteRecordRaw[] = []

// 遍历模块列表生成模块路由
for (const path of modules.keys()) {
	const mod = modules(path)
	if (is<{ default: RouteRecordRaw }>(mod, 'Module')) {
		menuList.push(mod.default)
	}
}

// 模块路由排序
menuList.sort((a, b) => {
	const aSort = get(a, 'meta.sort', 0) as number
	const bSort = get(b, 'meta.sort', 0) as number
	return aSort - bSort
})

const dashboardRouter: RouteRecordRaw = {
	path: '/',
	component: Layout,
	redirect: '/overview',
	meta: {
		keepAlive: true,
		key: 'overview',
		title: 'Overview',
		icon: 'home-outline',
	},
	children: [
		{
			path: 'overview',
			name: 'Overview',
			component: () => import('@/views/overview/index.vue'),
		},
	],
}

menuList.unshift(dashboardRouter)

export const routes: RouteRecordRaw[] = [...menuList]

const router = createRouter({
	history: createWebHistory('/'),
	routes,
	strict: false,
	scrollBehavior: () => ({ left: 0, top: 0 }),
})

export default router
