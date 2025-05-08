import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/overview',
	component: Layout,
	meta: { sort: 0, key: 'overview', title: 'Overview', titleKey: 'layout.menu.overview' },
	children: [
		{
			path: '/overview',
			name: 'Overview',
			component: () => import('@/views/overview/index.vue'),
		},
	],
}

export default route
