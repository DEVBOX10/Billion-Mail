<template>
	<n-tabs :value="activeTab" class="route-tabs">
		<n-tab-pane v-for="tab in tabs" :key="tab.path" :name="tab.path" :tab="tab.title"> </n-tab-pane>
	</n-tabs>
</template>

<script lang="ts" setup>
import { RouteLocationNormalized } from 'vue-router'

interface TabItem {
	title: string
	path: string
}

const route = useRoute()

const activeTab = ref('')

const tabs = ref<TabItem[]>([])

// 添加标签页
const addTab = (route: RouteLocationNormalized) => {
	const existTab = tabs.value.find(tab => tab.path === route.path)
	if (!existTab) {
		const newTab: TabItem = {
			title: (route.meta.title as string) || '未命名',
			path: route.path,
		}

		tabs.value.push(newTab)
	}
}

addTab(route)
</script>

<style lang="scss" scoped></style>
