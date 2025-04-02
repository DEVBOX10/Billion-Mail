<template>
	<modal title="Domain SSL" :width="650" :footer="false">
		<div>aaaaa</div>
	</modal>
</template>

<script setup lang="ts">
import { useModal } from '@/hooks/modal/useModal'
import { MailDomain } from '../interface'

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
