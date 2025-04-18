<template>
	<modal :title="title" width="520">
		<bt-form ref="formRef" class="pt-20px" :model="form" :rules="rules">
			<n-form-item :label="t('mailbox.form.emailAddress')" path="full_name">
				<div class="w-280px">
					<n-input-group>
						<n-input v-model:value="form.full_name" class="flex-1" :disabled="isEdit"> </n-input>
						<domain-select
							v-model:value="form.domain"
							class="flex-1"
							:is-all="false"
							:disabled="isEdit">
						</domain-select>
					</n-input-group>
				</div>
			</n-form-item>
			<n-form-item :label="t('mailbox.form.password')" path="password">
				<div class="w-280px">
					<n-input
						v-model:value="form.password"
						:placeholder="t('mailbox.form.passwordPlaceholder')">
					</n-input>
				</div>
			</n-form-item>
			<n-form-item :label="t('mailbox.form.quota')" path="quota">
				<div class="w-280px">
					<n-input-group>
						<n-input-number v-model:value="form.quota" class="flex-1" :min="1" :show-button="false">
						</n-input-number>
						<n-select v-model:value="form.unit" class="flex-1" :options="unitOptions"> </n-select>
					</n-input-group>
				</div>
			</n-form-item>
			<n-form-item :label="t('mailbox.form.userType')" path="userType">
				<div class="w-280px">
					<n-select v-model:value="form.isAdmin" :options="userTypeOptions" />
				</div>
			</n-form-item>
			<n-form-item :label="t('mailbox.form.status')">
				<n-switch v-model:value="form.active" :checked-value="1" :unchecked-value="0" />
			</n-form-item>
		</bt-form>
	</modal>
</template>

<script lang="ts" setup>
import { FormRules } from 'naive-ui'
import { getByteUnit, getNumber } from '@/utils'
import { useModal } from '@/hooks/modal/useModal'
import { createMailbox, updateMailbox } from '@/api/modules/mailbox'
import type { MailBox } from '../interface'

import DomainSelect from './DomainSelect.vue'

const { t } = useI18n()

const isEdit = ref(false)

const title = computed(() =>
	isEdit.value ? t('mailbox.form.editTitle') : t('mailbox.form.addTitle')
)

const formRef = useTemplateRef('formRef')

const form = reactive({
	quota: 5,
	unit: 'GB',
	isAdmin: 0,
	full_name: '',
	domain: null as string | null,
	password: '',
	active: 1,
})

const unitOptions = [
	{ label: 'GB', value: 'GB' },
	{ label: 'MB', value: 'MB' },
]

const userTypeOptions = [
	{ label: t('mailbox.userType.general'), value: 0 },
	{ label: t('mailbox.userType.admin'), value: 1 },
]

const rules: FormRules = {
	full_name: {
		trigger: 'blur',
		validator: () => {
			if (form.full_name.trim() === '' || !form.domain) {
				return new Error(t('mailbox.validation.emailRequired'))
			}
			return true
		},
	},
	password: {
		trigger: 'blur',
		validator: () => {
			if (!isEdit.value) {
				if (form.password.trim().length < 8) {
					return new Error(t('mailbox.validation.passwordLength'))
				}
				const pwdReg = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).*$/
				if (!pwdReg.test(form.password)) {
					return new Error(t('mailbox.validation.passwordFormat'))
				}
			} else {
				if (form.password && form.password.trim().length < 8) {
					return new Error(t('mailbox.validation.passwordLength'))
				}
				const pwdReg = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).*$/
				if (form.password && !pwdReg.test(form.password)) {
					return new Error(t('mailbox.validation.passwordFormat'))
				}
			}
			return true
		},
	},
}

/**
 * @description Calculate the byte number based on the domain quota and unit
 * @param quota
 * @param quota_unit
 */
const getQuotaByte = (quota: number, quota_unit: string) => {
	switch (quota_unit) {
		case 'GB':
			quota *= 1024
		case 'MB':
			quota *= 1024 * 1024
	}
	return quota
}

const getParams = () => {
	return {
		full_name: form.full_name,
		domain: form.domain || '',
		password: form.password,
		quota: getQuotaByte(form.quota, form.unit),
		isAdmin: form.isAdmin,
		active: form.active,
	}
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const state = modalApi.getState<{ isEdit: boolean; row: MailBox | null }>()
			const { row } = state
			isEdit.value = state.isEdit
			if (row) {
				form.full_name = row.full_name
				form.domain = row.domain
				form.isAdmin = row.is_admin
				form.active = row.active

				const quota = getByteUnit(row.quota)
				const [quotaNum, quotaUnit] = quota.split(' ')
				form.quota = getNumber(quotaNum)
				form.unit = quotaUnit
			}
		} else {
			form.full_name = ''
			form.domain = null
			form.password = ''
			form.quota = 5
			form.unit = 'GB'
			form.isAdmin = 0
			form.active = 1
		}
	},
	onConfirm: async () => {
		await formRef.value?.validate()
		const params = getParams()
		if (isEdit.value) {
			await updateMailbox(params)
		} else {
			await createMailbox(params)
		}
		const state = modalApi.getState<{ refresh: Function }>()
		state.refresh()
	},
})
</script>
