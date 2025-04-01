<template>
	<n-layout-sider
		:width="264"
		:content-style="{
			display: 'flex',
			flexDirection: 'column',
			height: '100%',
			overflow: 'hidden',
		}">
		<!-- 应用标志和名称 -->
		<div class="app-logo">
			<a href="#">
				<img class="icon" src="@images/logo.png" alt="" />
				<span class="app-name">Billion Mail</span>
			</a>
		</div>

		<!-- 导航菜单 -->
		<div class="nav-section">
			<n-menu :value="activeMenuKey" :options="menuOptions" :root-indent="24"> </n-menu>
		</div>
	</n-layout-sider>
</template>

<script lang="tsx" setup>
import { RouterLink, RouteRecordRaw } from 'vue-router'
import { useMenuStore } from '@/store'
import { menuList } from '@/router/router'

const route = useRoute()

const menuStore = useMenuStore()

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
	}))
})

const renderLabel = (route: RouteRecordRaw) => {
	const name = route.children?.[0]?.name || ''
	return (
		<RouterLink class="flex items-center" to={{ name }}>
			<span class={`${route.meta?.icon || ''} text-22px`}></span>
			<span class="ml-14px">{route.meta?.title || ''}</span>
		</RouterLink>
	)
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
	--n-item-color-active: rgba(0, 0, 0, 0.1);
	--n-item-text-color-active: #ffffff;
	--n-item-color-hover: rgba(0, 0, 0, 0.1);
	--n-item-text-color-hover: #ffffff;
	--n-item-color-active-hover: rgba(0, 0, 0, 0.1);
	--n-item-text-color-active-hover: #ffffff;
	:deep(.n-menu-item) {
		margin-top: 0;
		.n-menu-item-content {
			padding-right: 24px;
			line-height: inherit;
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
