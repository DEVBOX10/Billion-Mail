<template>
	<n-tabs v-model:value="menu" class="bt-tabs h-full" type="card" placement="top">
		<slot></slot>
	</n-tabs>
</template>

<script lang="ts" setup>
const menu = defineModel<string | number>('value')
</script>

<style lang="scss" scoped>
.bt-tabs {
	--n-tab-gap: 4px;
	--n-tab-border-radius: 0;
	--n-tab-text-color-active: #333;
	--n-tab-border-color: #cacad9;
	--n-tab-padding-vertical: 0 16px;
	--n-pane-padding-top: 16px;
	--n-pane-padding-left: 0;
	--n-pane-padding-right: 0;
	--n-pane-padding-bottom: 0;
	:deep(> .n-tabs-nav) {
		.n-tabs-wrapper {
			.n-tabs-tab-wrapper {
				.n-tabs-tab {
					height: 32px;
					border-bottom: none;
					background: -webkit-gradient(linear, 0% 0, 0% 100%, from(#f6f6f6), to(#ddd));
					&--active {
						background: #0000;
					}
				}
			}
		}
	}
	:deep(> .n-tab-pane) {
		flex: 1;
		overflow: auto;
	}
}
</style>
