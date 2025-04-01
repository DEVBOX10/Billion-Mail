<template>
	<div class="provider-table">
		<n-data-table :columns="columns" :data="tableData" :bordered="false"> </n-data-table>
	</div>
</template>

<script setup lang="ts">
import { NDataTable } from 'naive-ui'

const columns = [
	{
		title: '邮件提供商',
		key: 'provider',
	},
	{
		title: '送达率',
		key: 'delivered',
	},
	{
		title: '打开率',
		key: 'open',
	},
	{
		title: '点击率',
		key: 'click',
	},
	{
		title: '退回率',
		key: 'bounce',
	},
]

const tableData = [
	{
		provider: 'google',
		delivered: '99%',
		open: '98%',
		click: '20%',
		bounce: '1%',
	},
	{
		provider: 'outlook',
		delivered: '99%',
		open: '98%',
		click: '20%',
		bounce: '1%',
	},
	{
		provider: 'outlook',
		delivered: '99%',
		open: '98%',
		click: '20%',
		bounce: '1%',
	},
	{
		provider: 'outlook',
		delivered: '99%',
		open: '98%',
		click: '20%',
		bounce: '1%',
	},
	{
		provider: 'outlook',
		delivered: '99%',
		open: '98%',
		click: '20%',
		bounce: '1%',
	},
]
</script>
