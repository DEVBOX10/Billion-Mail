<template>
	<div class="py-16px px-24px">
		<n-card :bordered="false">
			<settings-service />
		</n-card>
	</div>
</template>

<script lang="ts" setup>
import SettingsService from './components/SettingsService.vue'
</script>
