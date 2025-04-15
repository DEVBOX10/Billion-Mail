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
// Check if <slot name="toolsLeft"/> has a value
const slotLeft = !!useSlots().toolsLeft
// Check if <slot name="toolsRight"/> has a value
const slotRight = !!useSlots().toolsRight
// Check if <slot name="pageRight"/> has a value
const slotPage = !!useSlots().pageLeft || !!useSlots().pageRight
</script>
