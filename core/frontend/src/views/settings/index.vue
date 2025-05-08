<template>
	<div class="p-24px">
		<div class="mb-20px text-24px font-bold">{{ $t('layout.menu.settings') }}</div>
		<n-card :bordered="false">
			<settings-service />
		</n-card>
	</div>
</template>

<script lang="ts" setup>
import SettingsService from './components/SettingsService.vue'
</script>
