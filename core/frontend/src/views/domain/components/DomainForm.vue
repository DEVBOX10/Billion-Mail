<template>
	<modal :title="title" :width="580">
		<div class="pt-20px">
			<bt-form ref="formRef" :model="form" :rules="rules">
				<n-form-item label="域名" path="domain">
					<div class="w-320px">
						<n-input
							v-model:value="form.domain"
							:disabled="isEdit"
							placeholder="请输入域名，例如：aapanel.com">
						</n-input>
					</div>
				</n-form-item>
				<n-form-item label="A记录">
					<div class="w-320px">
						<n-input
							v-model:value="form.a_record"
							:disabled="isEdit"
							placeholder="请输入A记录，例如：mail.aapanel.com">
						</n-input>
					</div>
				</n-form-item>
				<n-form-item label="域名配额">
					<n-input-number v-model:value="form.quota" class="w-200px" :min="0" :show-button="false">
					</n-input-number>
					<div class="w-100px ml-20px">
						<n-select v-model:value="form.quota_unit" :options="unitOptions"></n-select>
					</div>
				</n-form-item>
				<n-form-item label="邮箱数量">
					<n-input-number
						v-model:value="form.mailboxes"
						class="w-320px"
						:min="0"
						:show-button="false">
					</n-input-number>
				</n-form-item>
				<n-form-item v-if="false" label="全局捕获">
					<div class="w-320px">
						<n-input v-model:value="form.email" placeholder="捕获不存在的邮件，转发至此邮件">
						</n-input>
					</div>
				</n-form-item>
			</bt-form>
			<bt-tips>
				<li class="text-error">提示A记录解析失败。请检查填写的A记录域名是否已解析到服务器 IP</li>
				<li>需要在您的DNS服务提供商控制台添加A记录</li>
				<li>如果使用CloudFlare，请在添加记录时选择【仅DNS】</li>
			</bt-tips>
		</div>
	</modal>
</template>

<script lang="ts" setup>
import { FormRules } from 'naive-ui'
import { getByteUnit, getNumber } from '@/utils'
import { useModal } from '@/hooks/modal/useModal'
import { createDomain, updateDomain } from '@/api/modules/domain'
import { MailDomain } from '../interface'

const isEdit = ref(false)

const title = computed(() => {
	return isEdit.value ? '编辑域名' : '添加域名'
})

const formRef = useTemplateRef('formRef')

const form = reactive({
	domain: '',
	a_record: '',
	quota: 5,
	quota_unit: 'GB',
	mailboxes: 50,
	email: '',
})

const unitOptions = [
	{ label: 'GB', value: 'GB' },
	{ label: 'MB', value: 'MB' },
]

const rules: FormRules = {
	domain: {
		trigger: ['blur', 'input'],
		validator: () => {
			if (form.domain.trim() === '') {
				return new Error('邮箱域名不能为空!')
			}
			return true
		},
	},
}

/**
 * @description 根据域名配额和单位，计算出字节数
 * @param quota
 * @param quota_unit
 */
const getQuotaByte = (quota: number, quota_unit: string) => {
	switch (quota_unit) {
		case 'GB':
			return quota * 1024 * 1024 * 1024
		case 'MB':
			return quota * 1024 * 1024
		default:
			return quota
	}
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const state = modalApi.getState<{ isEdit: boolean; row: MailDomain | null }>()
			isEdit.value = state.isEdit
			if (state.row) {
				const { row } = state
				form.domain = row.domain
				form.a_record = row.a_record
				const quota = getByteUnit(row.quota)
				form.quota = getNumber(quota.split(' ')[0])
				form.quota_unit = quota.split(' ')[1]
				form.mailboxes = row.mailboxes
				// form.email = row.email
			}
		} else {
			form.domain = ''
			form.a_record = ''
			form.quota = 5
			form.quota_unit = 'GB'
			form.mailboxes = 50
			// form.email = ''
		}
	},
	onConfirm: async () => {
		await formRef.value?.validate()
		if (isEdit.value) {
			await updateDomain({
				domain: form.domain,
				quota: getQuotaByte(form.quota, form.quota_unit),
				mailboxes: form.mailboxes,
			})
		} else {
			await createDomain({
				domain: form.domain,
				quota: getQuotaByte(form.quota, form.quota_unit),
				mailboxes: form.mailboxes,
			})
		}
		const state = modalApi.getState<{ refresh: Function }>()
		state.refresh()
	},
})
</script>
