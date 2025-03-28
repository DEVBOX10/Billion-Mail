import { defineStore } from 'pinia'
import { RouteRecordRaw } from 'vue-router'

export default defineStore('MenuStore', () => {
	const isCollapse = ref(false)
	const menuList = ref<RouteRecordRaw[]>([])

	const setCollapse = () => {
		isCollapse.value = !isCollapse.value
	}

	const setMenuList = (menus: RouteRecordRaw[]) => {
		menuList.value = menus
	}

	const closeSidebar = () => {
		isCollapse.value = false
	}

	return {
		isCollapse,
		menuList,
		setCollapse,
		setMenuList,
		closeSidebar,
	}
})
