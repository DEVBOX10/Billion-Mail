<template>
	<modal :title="t('contacts.group.add.title')" width="460">
		<div class="pt-12px">
			<bt-form ref="formRef" :model="form" :rules="rules">
				<n-form-item :label="t('contacts.group.form.name')" path="name">
					<div class="w-240px">
						<n-input v-model:value="form.name"></n-input>
					</div>
				</n-form-item>
			</bt-form>
		</div>
	</modal>
</template>

<script lang="ts" setup>
import { useModal } from '@/hooks/modal/useModal'
import { createGroup } from '@/api/modules/contacts/group'

const { t } = useI18n()

const formRef = useTemplateRef('formRef')

const form = reactive({
	create_type: 1,
	name: '',
})

const rules = {
	name: {
		required: true,
		message: t('contacts.group.form.nameRequired'),
		trigger: 'blur',
	},
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			form.name = ''
		}
	},
	onConfirm: async () => {
		await formRef.value?.validate()
		await createGroup(toRaw(form))
		const state = modalApi.getState<{ refresh: () => void }>()
		state?.refresh()
	},
})
</script>
