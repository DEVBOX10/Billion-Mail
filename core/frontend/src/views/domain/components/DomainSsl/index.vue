<template>
	<modal :title="$t('domain.ssl.title')" :width="650" :footer="false">
		<bt-tabs v-model:value="menu">
			<n-tab-pane name="other" :tab="$t('domain.ssl.tabs.otherCertificate')">
				<other-cert v-if="domainInfo" :domain="domainInfo" :refresh="refresh" />
			</n-tab-pane>
		</bt-tabs>
	</modal>
</template>

<script setup lang="ts">
import { useModal } from '@/hooks/modal/useModal'
import { MailDomain } from '../../interface'

import OtherCert from './OtherCert.vue'

const menu = ref('other')

const domainInfo = ref<MailDomain | null>(null)

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const state = modalApi.getState<{ row: MailDomain | null }>()
			const { row } = state
			if (row) {
				domainInfo.value = row
			}
		}
	},
})

const refresh = () => {
	const state = modalApi.getState<{ refresh: () => void }>()
	state.refresh()
}
</script>
