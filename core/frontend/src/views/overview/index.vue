<template>
	<div class="email-analytics-container">
		<!-- Top Filter Area -->
		<filter-bar
			v-model:domain="domain"
			v-model:date-range="dateRange"
			@update:domain="handleDataUpdate"
			@update:date-range="handleDataUpdate">
		</filter-bar>

		<!-- Main Metric Cards -->
		<div class="metrics-cards">
			<template v-for="(item, key) in rateData" :key="key">
				<metric-card :title="item.label" :value="item.value" />
			</template>
		</div>

		<!-- Data Details Area -->
		<div class="detail-row">
			<!-- Left: Mail Provider Table -->
			<n-card class="provider-table-card" :title="$t('overview.mailProviders')">
				<provider-table v-model:value="providers" />
			</n-card>

			<!-- Right: Today's Sending Statistics -->
			<n-card class="send-today-card" :title="$t('overview.sendStats')">
				<send-today-stats :data="sendMail" />
			</n-card>
		</div>

		<!-- Bottom Rate Charts Area -->
		<div class="rate-charts-card">
			<rate-chart-panel :bounce="bounceRate" :click="clickRate" :open="openRate" />
		</div>
	</div>
</template>

<script lang="ts" setup>
import { useDebounceFn } from '@vueuse/core'
import { getDayTimeRange, isObject } from '@/utils'
import { getOverviewInfo } from '@/api/modules/overview'
import type { MailOverview, MailProvider, RateData, RateKey } from './interface'

import FilterBar from './components/FilterBar.vue'
import MetricCard from './components/MetricCard.vue'
import ProviderTable from './components/ProviderTable.vue'
import SendTodayStats from './components/SendTodayStats.vue'
import RateChartPanel from './components/RateChartPanel.vue'

const { t } = useI18n()

const domain = ref('')

const dateRange = ref(getDayTimeRange())

const providers = ref<MailProvider[]>([])

const rateData = reactive<RateData>({
	delivery: { label: t('overview.delivered'), value: 0 },
	open: { label: t('overview.opened'), value: 0 },
	click: { label: t('overview.clicked'), value: 0 },
	bounce: { label: t('overview.bounced'), value: 0 },
})

const sendMail = ref<MailOverview['send_mail_chart']>({
	column_type: 'hourly',
	dashboard: {
		delivered: 0,
		delivery_rate: 0,
		failed: 0,
		failure_rate: 0,
		sends: 0,
	},
	data: [],
})

const bounceRate = ref<MailOverview['bounce_rate_chart']>({
	column_type: 'hourly',
	data: [],
})

const clickRate = ref<MailOverview['click_rate_chart']>({
	column_type: 'hourly',
	data: [],
})

const openRate = ref<MailOverview['open_rate_chart']>({
	column_type: 'hourly',
	data: [],
})

// Function to handle data update
const handleDataUpdate = useDebounceFn(fetchOverviewData, 300)

// Function to update rate data
const updateRateData = (dashboard: MailOverview['dashboard']) => {
	Object.entries(dashboard).forEach(([key, value]) => {
		const rateKey = key.replace('_rate', '') as RateKey
		if (rateKey in rateData) {
			rateData[rateKey].value = value
		}
	})
}

// Function to fetch overview data
async function fetchOverviewData() {
	const res = await getOverviewInfo({
		domain: domain.value,
		start_time: Math.floor(dateRange.value[0] / 1000),
		end_time: Math.floor(dateRange.value[1] / 1000),
	})

	if (isObject<MailOverview>(res)) {
		updateRateData(res.dashboard)

		providers.value = res.mail_providers
		sendMail.value = res.send_mail_chart
		bounceRate.value = res.bounce_rate_chart
		clickRate.value = res.click_rate_chart
		openRate.value = res.open_rate_chart
	}
}

onMounted(() => {
	fetchOverviewData()
})
</script>

<style lang="scss" scoped>
.email-analytics-container {
	padding: 20px;
}

.metrics-cards {
	display: flex;
	justify-content: space-between;
	gap: 20px;
	margin-bottom: 20px;
}

.detail-row {
	display: flex;
	gap: 20px;
	margin-bottom: 20px;

	.provider-table-card,
	.send-today-card {
		flex: 1;
	}
}

.rate-charts-card {
	display: flex;
	gap: 20px;
}
</style>
