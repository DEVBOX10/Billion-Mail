<template>
	<modal :title="t('domain.dns.title')" :width="750" :footer="false">
		<div>
			<n-data-table class="mb-20px" :data="statusData" :columns="statusColumns"></n-data-table>
			<div class="record-title">{{ t('domain.dns.step1.title') }}</div>
			<div class="record-desc">{{ t('domain.dns.step1.description') }}</div>
			<n-table class="mb-30px">
				<thead>
					<tr>
						<th>{{ t('domain.dns.table.recordType') }}</th>
						<th>{{ t('domain.dns.table.host') }}</th>
						<th>{{ t('domain.dns.table.value') }}</th>
						<th>{{ t('domain.dns.table.mxPriority') }}</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>MX</td>
						<td>@</td>
						<td>{{ mxValue }}</td>
						<td>10</td>
					</tr>
				</tbody>
			</n-table>
			<div class="record-title">{{ t('domain.dns.step2.title') }}</div>
			<div class="record-desc">{{ t('domain.dns.step2.description') }}</div>
			<n-table class="mb-30px">
				<thead>
					<tr>
						<th>{{ t('domain.dns.table.recordType') }}</th>
						<th>{{ t('domain.dns.table.host') }}</th>
						<th>{{ t('domain.dns.table.value') }}</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>TXT</td>
						<td>@</td>
						<td>
							<n-ellipsis class="mr-4px max-w-360px!">{{ command }}</n-ellipsis>
							<n-button text type="primary" @click="() => copyText(command)">
								{{ t('common.actions.copy') }}
							</n-button>
						</td>
					</tr>
					<tr>
						<td>TXT</td>
						<td>default._domainkey</td>
						<td>
							<n-ellipsis class="mr-4px max-w-360px!">
								{{ dkimValue || '--' }}
							</n-ellipsis>
							<n-button v-if="dkimValue" text type="primary" @click="() => copyText(dkimValue)">
								{{ t('common.actions.copy') }}
							</n-button>
						</td>
					</tr>
					<tr>
						<td>TXT</td>
						<td>_dmarc</td>
						<td>
							<n-ellipsis class="mr-4px max-w-360px!">
								{{ dmarcValue }}
							</n-ellipsis>
							<n-button v-if="dmarcValue" text type="primary" @click="() => copyText(dmarcValue)">
								{{ t('common.actions.copy') }}
							</n-button>
						</td>
					</tr>
				</tbody>
			</n-table>
			<div class="record-title">{{ t('domain.dns.step3.title') }}</div>
			<div class="record-desc">{{ t('domain.dns.step3.description1') }}</div>
			<div class="record-desc">{{ t('domain.dns.step3.description2') }}</div>
		</div>
		<div class="flex justify-center mt-24px">
			<n-button type="primary" @click="onVerify">{{ t('domain.dns.verify') }}</n-button>
		</div>
	</modal>
</template>

<script lang="tsx" setup>
import { DataTableColumns, NIcon } from 'naive-ui'
import { SuccessIcon, ErrorIcon } from 'naive-ui/es/_internal/icons'
import { isObject } from '@/utils'
import { useCopy } from '@/hooks/useCopy'
import { useModal } from '@/hooks/modal/useModal'
import { freshDnsRecord } from '@/api/modules/domain'
import type { DomainDnsRecords, MailDomain } from '../interface'

const { t } = useI18n()

const { copyText } = useCopy()

const statusData = ref<MailDomain[]>([])

const getRecordValue = (record: string) => {
	if (statusData.value.length > 0) {
		return statusData.value[0].dns_records[record].value
	}
	return ''
}

const dkimValue = computed(() => {
	return getRecordValue('dkim')
})

const mxValue = computed(() => {
	return getRecordValue('mx')
})

const dmarcValue = computed(() => {
	return getRecordValue('dmarc')
})

const command = computed(() => {
	return getRecordValue('spf')
})

const StatusCol = ({ status }: { status: boolean }) => {
	return (
		<div class={['inline-flex items-center', status ? 'text-primary' : 'text-error']}>
			<NIcon size="16">{status ? <SuccessIcon /> : <ErrorIcon />}</NIcon>
			<span class="ml-4px">
				{status ? t('domain.dns.status.ok') : t('domain.dns.status.notSet')}
			</span>
		</div>
	)
}

const statusColumns = ref<DataTableColumns<MailDomain>>([
	{
		key: 'domain',
		title: t('domain.dns.columns.domain'),
		width: '20%',
		ellipsis: {
			tooltip: true,
		},
	},
	{
		key: 'mx',
		title: 'MX',
		render: row => {
			const mx = row.dns_records.mx
			return <StatusCol status={mx?.valid || false}></StatusCol>
		},
	},
	{
		key: 'a',
		title: 'A record',
		render: row => {
			const a = row.dns_records.a
			return <StatusCol status={a?.valid || false}></StatusCol>
		},
	},
	{
		key: 'spf',
		title: 'SPF',
		render: row => {
			const spf = row.dns_records.spf
			return <StatusCol status={spf?.valid || false}></StatusCol>
		},
	},
	{
		key: 'dkim',
		title: 'DKIM',
		render: row => {
			const dkim = row.dns_records.dkim
			return <StatusCol status={dkim?.valid || false}></StatusCol>
		},
	},
	{
		key: 'dmarc',
		title: 'DMARC',
		render: row => {
			const dmarc = row.dns_records.dmarc
			return <StatusCol status={dmarc?.valid || false}></StatusCol>
		},
	},
	{
		key: 'ptr',
		title: 'PTR',
		render: row => {
			const ptr = row.dns_records.ptr
			return <StatusCol status={ptr?.valid || false}></StatusCol>
		},
	},
])

const onVerify = async () => {
	const res = await freshDnsRecord({ domain: statusData.value[0].domain })
	if (isObject<DomainDnsRecords>(res)) {
		statusData.value[0].dns_records = res
	}
	const { refresh } = modalApi.getState<{ refresh: () => void }>()
	refresh()
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const state = modalApi.getState<{ row: MailDomain | null }>()
			const { row } = state
			if (row) {
				statusData.value = [row]
			}
		}
	},
})
</script>

<style lang="scss" scoped>
.record-title {
	font-size: 16px;
	color: #444;
	font-weight: 600;
	margin-bottom: 10px;
}

.record-desc {
	font-size: 14px;
	color: #666;
	margin-bottom: 10px;
}
</style>
