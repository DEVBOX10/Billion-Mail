import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/email_marketing',
	component: Layout,
	meta: {
		sort: 2,
		key: 'market',
		title: 'Email Marketing',
		icon: 'i-mdi-email-fast-outline',
		hidden: true,
	},
	children: [
		{
			path: '/email_marketing',
			name: 'Market',
			component: () => import('@/views/market/index.vue'),
		},
	],
}

export default route
