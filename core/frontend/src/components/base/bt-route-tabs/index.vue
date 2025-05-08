<template>
	<div class="mb-4px text-24px font-bold">{{ title }}</div>
	<n-tabs :value="activeTab" type="line" class="route-tabs" @update:value="handleUpdateTab">
		<n-tab-pane v-for="tab in tabs" :key="tab.name" :name="tab.name" :tab="$t(tab.title)">
		</n-tab-pane>
	</n-tabs>
	<router-view />
</template>

<script lang="ts" setup>
import { get } from 'lodash-es'
import { RouteLocationNormalized } from 'vue-router'

interface TabItem {
	title: string
	name: string
}

const { t } = useI18n()

const route = useRoute()

const router = useRouter()

const tabs = ref<TabItem[]>([])

const activeTab = computed(() => {
	return String(route.name)
})

const title = computed(() => {
	return t(`${get(route.matched[0], 'meta.titleKey')}`)
})

const handleUpdateTab = (name: string) => {
	router.push({ name })
}

// Add tab when route changes
const addTab = (route: RouteLocationNormalized) => {
	const existTab = tabs.value.find(tab => tab.name === route.name)
	if (existTab) return

	if (route.matched[1].children) {
		tabs.value = route.matched[1].children.map(routeItem => {
			return {
				title: `${get(routeItem, 'meta.titleKey')}`,
				name: String(routeItem.name),
			}
		})
	}
}

addTab(route)
</script>

<style lang="scss" scoped>
.route-tabs {
	--n-tab-gap: 32px;
	--n-tab-padding: 12px 4px;
	--n-pane-padding-top: 16px;
	--n-tab-border-color: rgba(229, 231, 235, 0.75);
}
</style>
