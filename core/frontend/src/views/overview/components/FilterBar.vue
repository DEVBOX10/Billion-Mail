<template>
	<div class="filter-bar">
		<div class="w-260px">
			<n-select v-model:value="selectedDomain" :options="domainOptions"></n-select>
		</div>

		<div class="date-filters">
			<n-radio-group v-model:value="activeFilter" @update:value="setFilter">
				<n-radio-button value="today">Today</n-radio-button>
				<n-radio-button value="yesterday">Yesterday</n-radio-button>
				<n-radio-button value="last7days">Last 7 Days</n-radio-button>
			</n-radio-group>
			<div class="w-240px">
				<n-date-picker
					:value="dateRange"
					type="daterange"
					:is-date-disabled="dateDisabled"
					@update:value="onUpdateTime">
				</n-date-picker>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import { addDays, startOfDay, endOfDay } from 'date-fns'

const selectedDomain = ref('all')

const domainOptions = ref([{ label: 'All Domains', value: 'all' }])

const activeFilter = ref('today')

const dateRange = ref(getTimeData())

function getTimeData(date = new Date()): [number, number] {
	const start = startOfDay(date)
	const end = endOfDay(date)
	return [start.getTime(), end.getTime()]
}

const dateDisabled = (ts: number) => {
	const today = new Date()
	const before = addDays(today, -30)
	return startOfDay(before).getTime() > ts || endOfDay(today).getTime() < ts
}

const setFilter = (filter: string) => {
	switch (filter) {
		case 'today':
			dateRange.value = getTimeData()
			break
		case 'yesterday':
			dateRange.value = getTimeData(addDays(new Date(), -1))
			break
		case 'last7days':
			dateRange.value = getSevenTime()
			break
	}
}

const getSevenTime = (): [number, number] => {
	const today = new Date()
	const end = endOfDay(today)
	const start = startOfDay(addDays(today, -6))

	return [start.getTime(), end.getTime()]
}

const onUpdateTime = (val: [number, number]) => {
	activeFilter.value = 'custom'
	dateRange.value = [startOfDay(val[0]).getTime(), endOfDay(val[1]).getTime()]
}
</script>

<style scoped>
.filter-bar {
	display: flex;
	justify-content: space-between;
	margin-bottom: 20px;
}

.date-filters {
	display: flex;
	gap: 10px;
}
</style>
