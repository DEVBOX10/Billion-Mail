<template>
	<div>
		<n-flex v-if="slotLeft || slotRight" class="items-center mb-16px" justify="space-between">
			<n-flex class="items-center" :size="8">
				<slot name="toolsLeft"></slot>
			</n-flex>
			<n-flex class="items-center" :size="8">
				<slot name="toolsRight"></slot>
			</n-flex>
		</n-flex>
		<slot name="table"></slot>
		<n-flex v-if="slotPage" class="items-center mt-16px" justify="space-between">
			<div>
				<slot name="pageLeft"></slot>
			</div>
			<n-flex>
				<slot name="pageRight"></slot>
			</n-flex>
		</n-flex>
		<slot name="modal"></slot>
	</div>
</template>

<script lang="ts" setup>
// 判断<slot name="toolsLeft"/>是否有传值
const slotLeft = !!useSlots().toolsLeft
// 判断<slot name="toolsRight"/>是否有传值
const slotRight = !!useSlots().toolsRight
// 判断<slot name="pageRight"/>是否有传值
const slotPage = !!useSlots().pageLeft || !!useSlots().pageRight
</script>
