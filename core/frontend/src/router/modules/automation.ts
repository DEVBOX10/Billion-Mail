import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/automation',
	component: Layout,
	meta: {
		sort: 3,
		key: 'automation',
		title: 'Automation',
		icon: 'i-mdi-refresh-auto',
		hidden: true,
	},
	children: [
		{
			path: '/automation',
			name: 'Automation',
			component: () => import('@/views/automation/index.vue'),
		},
	],
}

export default route
