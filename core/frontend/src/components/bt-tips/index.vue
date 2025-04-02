<template>
	<ul class="bt-tips flex flex-col gap-4px list-outside list-disc leading-20px text-[#777]">
		<slot></slot>
	</ul>
</template>

<style lang="scss" scoped>
.bt-tips {
	:deep(li) {
		list-style: disc;
		margin-left: 1.5em;
	}
}
</style>
