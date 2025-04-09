<template>
	<div class="p-20px">
		<n-button @click="handleClick">aaaaa</n-button>
	</div>
</template>

<script lang="ts" setup>
import { confirm } from '@/utils'

const handleClick = () => {
	confirm({
		title: 'xxxx',
		content: 'bbbb',
	})
}
</script>
