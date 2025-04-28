<template>
	<modal :title="$t('contacts.subscribers.edit.title')" width="520">
		<div class="pt-12px">
			<bt-form ref="formRef" :model="form" :rules="rules">
				<n-form-item :label="$t('contacts.subscribers.edit.email')" path="email">
					<div class="w-280px">
						<n-input v-model:value="form.email" :disabled="true"></n-input>
					</div>
				</n-form-item>
				<n-form-item :label="$t('contacts.subscribers.edit.group')" path="group_ids">
					<div class="w-280px">
						<group-select v-model:value="form.group_ids"></group-select>
					</div>
				</n-form-item>
			</bt-form>
		</div>
	</modal>
</template>

<script lang="ts" setup>
import { FormRules } from 'naive-ui'
import { useModal } from '@/hooks/modal/useModal'
import { updateSubscriberGroup } from '@/api/modules/contacts/subscribers'
import type { Subscriber } from '../interface'

import GroupSelect from './GroupMultipleSelect.vue'

const { t } = useI18n()

const formRef = useTemplateRef('formRef')

const form = reactive({
	email: '',
	status: 1,
	group_ids: [] as number[],
})

const rules: FormRules = {
	email: {
		required: true,
		message: t('contacts.subscribers.edit.validation.emailRequired'),
		trigger: 'blur',
	},
	group_ids: {
		trigger: 'change',
		validator: () => {
			if (form.group_ids.length === 0) {
				return new Error(t('contacts.subscribers.edit.validation.groupRequired'))
			}
			return true
		},
	},
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const state = modalApi.getState<{ row: Subscriber | null }>()
			const { row } = state
			if (row) {
				form.email = row.email
				form.status = row.active
				form.group_ids = row.groups.map(item => item.id)
			}
		} else {
			form.email = ''
			form.status = 1
			form.group_ids = []
		}
	},
	onConfirm: async () => {
		await formRef.value?.validate()
		await updateSubscriberGroup({
			emails: [form.email],
			group_ids: form.group_ids,
			status: form.status,
		})
		const state = modalApi.getState<{ refresh: () => void }>()
		state.refresh()
	},
})
</script>
