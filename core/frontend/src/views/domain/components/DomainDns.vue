<template>
	<modal title="DNS Record" :width="750" :footer="false">
		<div class="pt-20px">
			<n-data-table class="mb-20px" :data="statusData" :columns="statusColumns"></n-data-table>
			<div class="record-title">步骤1：添加MX记录</div>
			<div class="record-desc">
				登录到域名服务提供商，添加记录类型为MX的记录用于邮箱服务（请直接复制以下参数）
			</div>
			<n-table class="mb-30px">
				<thead>
					<tr>
						<th>记录类型</th>
						<th>主机记录</th>
						<th>记录值</th>
						<th>MX优先级</th>
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
			<div class="record-title">步骤2：添加TXT记录</div>
			<div class="record-desc">添加记录类型为TXT的记录用于邮箱反垃圾邮件（请直接复制以下参数）</div>
			<n-table class="mb-30px">
				<thead>
					<tr>
						<th>记录类型</th>
						<th>主机记录</th>
						<th>记录值</th>
					</tr>
				</thead>
				<tbody>
					<tr>
						<td>TXT</td>
						<td>@</td>
						<td>
							<n-ellipsis class="mr-4px max-w-360px!">{{ command }}</n-ellipsis>
							<n-button text type="primary" @click="onCopy(command)">(Copy)</n-button>
						</td>
					</tr>
					<tr>
						<td>TXT</td>
						<td>default._domainkey</td>
						<td>
							<n-ellipsis class="mr-4px max-w-360px!">
								{{ dkimValue || '--' }}
							</n-ellipsis>
							<n-button v-if="dkimValue" text type="primary" @click="onCopy(dkimValue)">
								(Copy)
							</n-button>
						</td>
					</tr>
					<tr>
						<td>TXT</td>
						<td>_dmarc</td>
						<td>
							<n-ellipsis class="mr-4px max-w-360px!">
								{{ mxValue }}
							</n-ellipsis>
							<n-button text type="primary" @click="onCopy(mxValue)"> (Copy) </n-button>
						</td>
					</tr>
				</tbody>
			</n-table>
			<div class="record-title">步骤3：添加PTR记录</div>
			<div class="record-desc">（可选）PTR记录用于反向 DNS 查找</div>
			<div class="record-desc">联系您的IP提供商创建 PTR 记录</div>
		</div>
	</modal>
</template>

<script lang="tsx" setup>
import { DataTableColumns, NIcon } from 'naive-ui'
import { SuccessIcon, ErrorIcon } from 'naive-ui/es/_internal/icons'
import { copyText } from '@/utils/copy'
import { useModal } from '@/hooks/modal/useModal'
import { MailDomain } from '../interface'

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

const command = computed(() => {
	let ip = ''
	// ip += row.value.ip_address.ipv4.map(item => `+ip4:${item}`).join(' ')
	// ip += row.value.ip_address.ipv6.map(item => `+ip6:${item}`).join(' ')
	return `v=spf1 +a +mx ${ip} -all`
})

const onCopy = (value: string) => {
	copyText(value)
}

const StatusCol = ({ status }: { status: boolean }) => {
	return (
		<div class={['inline-flex items-center', status ? 'text-primary' : 'text-error']}>
			<NIcon size="16">{status ? <SuccessIcon /> : <ErrorIcon />}</NIcon>
			<span class="ml-4px">{status ? 'OK' : 'Not Set'}</span>
		</div>
	)
}

const statusColumns = ref<DataTableColumns<MailDomain>>([
	{
		key: 'domain',
		title: '域名',
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
