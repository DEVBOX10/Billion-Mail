<template>
	<n-layout-header ref="headerRef" :style="{ top: `${top}px` }">
		<div class="header-left">
			<n-button class="icon-btn" :bordered="false" @click="handleCollapse">
				<i class="icon" :class="isCollapse ? 'i-mdi-menu-close' : 'i-mdi-menu-open'"></i>
			</n-button>
			<div class="text-16px">{{ routeTitle }}</div>
		</div>

		<div class="header-right">
			<n-button class="icon-btn" :bordered="false">
				<i class="icon i-mdi-language"></i>
			</n-button>
			<n-dropdown :options="userOptions" @select="handleUserAction">
				<n-button class="icon-btn" :bordered="false">
					<i class="icon i-mdi-user-outline"></i>
				</n-button>
			</n-dropdown>
		</div>
	</n-layout-header>
</template>

<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { useUserStore, useGlobalStore } from '@/store'

defineProps({
	top: {
		type: Number,
		default: 0,
	},
})

const route = useRoute()

const userStore = useUserStore()
const globalStore = useGlobalStore()

const { isCollapse } = storeToRefs(globalStore)

const routeTitle = computed(() => {
	return `${route.meta.title}`
})

const handleCollapse = () => {
	globalStore.setCollapse()
}

const userOptions = [
	{
		label: 'Logout',
		key: 'logout',
	},
]

const handleUserAction = (key: string) => {
	switch (key) {
		case 'logout':
			userStore.logout()
			break
	}
}
</script>

<style lang="scss" scoped>
.n-layout-header {
	position: absolute;
	top: 0;
	left: 0;
	right: 0;
	display: flex;
	justify-content: space-between;
	align-items: center;
	height: 72px;
	padding: 0 24px;
	backdrop-filter: blur(7px);
	z-index: 1000;
}

.header-left,
.header-right {
	display: flex;
	align-items: center;
	gap: 8px;
}

.icon-btn {
	--n-width: 44px;
	--n-height: 44px;
	--n-padding: 0;
	--n-font-size: 24px;
	--n-text-color: #5b6b79;
	--n-ripple-color: none;
}
</style>
