import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/domain',
	component: Layout,
	meta: { sort: 5, key: 'domain', title: 'MailDomain', icon: 'i-mdi-web' },
	children: [
		{
			path: '/domain',
			name: 'Domain',
			component: () => import('@/views/domain/index.vue'),
		},
	],
}

export default route
