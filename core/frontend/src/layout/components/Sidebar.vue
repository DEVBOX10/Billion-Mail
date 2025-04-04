<template>
	<n-layout-sider
		:width="isCollapse ? 64 : 240"
		:content-style="{
			display: 'flex',
			flexDirection: 'column',
			height: '100%',
			overflow: 'hidden',
		}">
		<!-- 应用标志和名称 -->
		<div class="app-logo" :class="{ collapse: isCollapse }">
			<a href="#">
				<img class="icon" src="@images/logo.png" alt="" />
				<span v-show="!isCollapse" class="app-name">Billion Mail</span>
			</a>
		</div>

		<!-- 导航菜单 -->
		<div class="nav-section">
			<n-menu
				:value="activeMenuKey"
				:collapsed="isCollapse"
				:collapsed-width="64"
				:options="menuOptions"
				:root-indent="24">
			</n-menu>
		</div>
	</n-layout-sider>
</template>

<script lang="tsx" setup>
import { RouterLink, RouteRecordRaw } from 'vue-router'
import { useMenuStore, useGlobalStore } from '@/store'
import { menuList } from '@/router/router'
import { storeToRefs } from 'pinia'
import { VNodeChild } from 'vue'

const route = useRoute()

const menuStore = useMenuStore()
const globalStore = useGlobalStore()

const { isCollapse } = storeToRefs(globalStore)

// 当前菜单名称
const activeMenuKey = computed(() => {
	return String(route.meta?.key || '')
})

// 路由菜单
const routerMenus = computed(() => {
	return menuStore.menuList.filter(route => route.meta && !route.meta.hidden)
})

// 导航菜单选项
const menuOptions = computed(() => {
	return routerMenus.value.map(route => ({
		key: `${route.meta?.key}`,
		label: () => renderLabel(route),
		icon: () => renderIcon(route),
	}))
})

const renderLabel = (route: RouteRecordRaw) => {
	const name = route.children?.[0]?.name || ''
	return (
		<RouterLink class="flex items-center" to={{ name }}>
			<span class="ml-10px">{route.meta?.title || ''}</span>
		</RouterLink>
	)
}

const iconMap: Record<string, VNodeChild> = {
	domain: <i class="i-mdi-web"></i>,
	mailbox: <i class="i-mdi-email"></i>,
	settings: <i class="i-mdi-settings-outline"></i>,
}

const renderIcon = (route: RouteRecordRaw) => {
	if (!route.meta?.key) return null
	return iconMap[route.meta.key as string]
}

onMounted(() => {
	menuStore.setMenuList(menuList)
})
</script>

<style lang="scss" scoped>
.n-layout-sider {
	--n-color: #3f4d67;
	--n-text-color: #fff;
	box-shadow: 1px 0 20px 0 #3f4d67;
	z-index: 1010;
}

.app-logo {
	display: flex;
	padding: 16px 24px;
	&.collapse {
		justify-content: center;
		padding: 16px 0;
	}
	a {
		display: flex;
		align-items: center;
	}
	.icon {
		height: 36px;
	}
	.app-name {
		margin-left: 10px;
		font-size: 18px;
		font-weight: 500;
		color: #e8edf7;
	}
}

.nav-section {
	flex: 1;
	overflow: auto;
	padding-top: 10px;
}

.n-menu {
	--n-item-height: 44px;
	--n-item-text-color: #b7c0cd;
	--n-item-icon-color: #b7c0cd;
	--n-item-color-active: rgba(0, 0, 0, 0.1);
	--n-item-text-color-active: #ffffff;
	--n-item-icon-color-active: #ffffff;
	--n-item-color-hover: rgba(0, 0, 0, 0.1);
	--n-item-text-color-hover: #ffffff;
	--n-item-icon-color-hover: #ffffff;
	--n-item-color-active-hover: rgba(0, 0, 0, 0.1);
	--n-item-text-color-active-hover: #ffffff;
	--n-item-icon-color-active-hover: #ffffff;
	--n-item-icon-color-collapsed: #b7c0cd;
	--n-item-color-active-collapsed: rgba(0, 0, 0, 0.1);
	:deep(.n-menu-item) {
		margin-top: 0;
		.n-menu-item-content {
			padding-right: 24px;
			line-height: 1.4;
			&::before {
				left: 0;
				right: 0;
			}
			&::after {
				content: '';
				position: absolute;
				top: 0;
				left: 0;
				bottom: 0;
				width: 3px;
				transition: background-color 0.3s var(--n-bezier);
			}
			&:hover,
			&.n-menu-item-content--selected {
				font-weight: 500;
				&::after {
					background-color: #20a53a;
				}
			}
		}
	}
}
</style>
