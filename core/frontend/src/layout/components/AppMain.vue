<template>
	<router-view :key="key" v-slot="{ Component, route }">
		<component :is="Component" :key="route.path" />
	</router-view>
</template>

<script lang="ts" setup>
const key = computed(() => {
	return Math.random()
})
</script>
