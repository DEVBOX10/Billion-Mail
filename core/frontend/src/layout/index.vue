<template>
	<config-provider>
		<n-layout class="w-full h-full" has-sider>
			<sidebar></sidebar>
			<n-layout>
				<app-header></app-header>
				<app-main></app-main>
			</n-layout>
		</n-layout>
	</config-provider>
</template>

<script lang="ts" setup>
import { Sidebar, AppHeader, AppMain } from './components'
</script>
