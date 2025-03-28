import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/settings',
	component: Layout,
	meta: { sort: 7, key: 'settings', title: 'Settings', icon: 'settings-outline' },
	children: [
		{
			path: '/settings',
			name: 'Settings',
			component: () => import('@/views/settings/index.vue'),
		},
	],
}

export default route
