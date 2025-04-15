import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/contacts',
	component: Layout,
	meta: {
		sort: 4,
		key: 'contacts',
		title: 'Contacts',
		hidden: true,
	},
	children: [
		{
			path: '/contacts',
			name: 'Contacts',
			component: () => import('@/views/contacts/index.vue'),
		},
	],
}

export default route
