import { RouteRecordRaw } from 'vue-router'
import { Layout } from '@/router/constant'

const route: RouteRecordRaw = {
	path: '/mailbox',
	component: Layout,
	meta: { sort: 6, key: 'mailbox', title: 'MailBoxes', icon: 'at' },
	children: [
		{
			path: '/mailbox',
			name: 'Mailbox',
			component: () => import('@/views/mailbox/index.vue'),
		},
	],
}

export default route
