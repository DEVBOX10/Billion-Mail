import { get } from 'lodash-es'
import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import { is } from '@/utils'
import { Layout } from '@/router/constant'

// Import routes from modules
const modules = import.meta.webpackContext('./modules', {
	// Whether to search for subdirectories
	recursive: false,
	regExp: /\.ts$/,
})

// Module routes
export const menuList: RouteRecordRaw[] = []

// Iterate through the module list to generate module routes
for (const path of modules.keys()) {
	const mod = modules(path)
	if (is<{ default: RouteRecordRaw }>(mod, 'Module')) {
		menuList.push(mod.default)
	}
}

// Sort module routes
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
		key: 'overview',
		title: 'Overview',
		hidden: false,
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

export const routes: RouteRecordRaw[] = [
	{
		path: '/login',
		name: 'Login',
		component: () => import('@/views/login/index.vue'),
	},
	...menuList,
]

const router = createRouter({
	history: createWebHistory('/'),
	routes,
	strict: false,
	scrollBehavior: () => ({ left: 0, top: 0 }),
})

export default router
