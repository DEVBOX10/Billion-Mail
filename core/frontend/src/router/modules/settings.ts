import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/settings',
	component: Layout,
	meta: { sort: 8, key: 'settings', title: 'Settings', titleKey: 'layout.menu.settings' },
	children: [
		{
			path: '/settings',
			name: 'Settings',
			component: () => import('@/views/settings/index.vue'),
		},
	],
}

export default route
