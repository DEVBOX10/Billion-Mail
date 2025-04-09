<template>
	<n-layout class="w-full h-full" has-sider>
		<sidebar></sidebar>
		<n-layout :content-style="{ position: 'relative', paddingTop: '48px' }" @scroll="handleScroll">
			<app-header :top="scrollTop"></app-header>
			<app-main></app-main>
		</n-layout>
	</n-layout>
</template>

<script lang="ts" setup>
import { Sidebar, AppHeader, AppMain } from './components'

const scrollTop = ref(0)

const handleScroll = (e: Event) => {
	scrollTop.value = (e.target as HTMLElement).scrollTop || 0
}
</script>
