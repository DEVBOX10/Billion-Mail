<template>
	<modal title="Domain SSL" :width="650" :footer="false">
		<bt-tabs v-model:value="menu">
			<n-tab-pane name="other" tab="Other certificate">
				<other-cert />
			</n-tab-pane>
		</bt-tabs>
	</modal>
</template>

<script setup lang="ts">
import { useModal } from '@/hooks/modal/useModal'
import { MailDomain } from '../../interface'

import OtherCert from './OtherCert.vue'

const menu = ref('other')

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const state = modalApi.getState<{ row: MailDomain | null }>()
			const { row } = state
			if (row) {
				// console.log(row)
			}
		}
	},
})
</script>
