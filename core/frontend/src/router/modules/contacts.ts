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
			children: [
				{
					path: 'subscribers',
					name: 'ContactsSubscribers',
					meta: { title: 'Subscribers' },
					component: () => import('@/views/contacts/subscribers/index.vue'),
				},
				{
					path: 'group',
					name: 'ContactsGroup',
					meta: { title: 'Group' },
					component: () => import('@/views/contacts/group/index.vue'),
				},
			],
		},
	],
}

export default route
