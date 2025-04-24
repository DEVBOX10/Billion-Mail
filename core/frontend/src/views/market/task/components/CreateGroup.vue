<template>
	<modal :title="t('market.task.createGroup.title')" width="580">
		<div class="pt-12px">
			<bt-form ref="formRef" :model="form" :rules="rules">
				<n-form-item :label="t('market.task.createGroup.groupName')" path="name">
					<div class="w-360px">
						<n-input
							v-model:value="form.name"
							:placeholder="t('market.task.createGroup.groupNamePlaceholder')">
						</n-input>
					</div>
				</n-form-item>
				<n-form-item :label="t('market.task.createGroup.email')" path="file_data">
					<div class="w-360px">
						<bt-file-upload :is-upload="false" :accept="['txt']" @change="handleChangeFile">
						</bt-file-upload>
						<div class="mt-16px text-12px text-desc">
							{{ $t('contacts.subscribers.import.fileTypeHint') }}
							<n-button text type="primary" @click="handleDownloadTemplate">
								{{ $t('common.actions.download') }}
							</n-button>
						</div>
					</div>
				</n-form-item>
			</bt-form>
		</div>
	</modal>
</template>

<script lang="ts" setup>
import { isEmpty } from 'lodash-es'
import { FormRules, UploadFileInfo } from 'naive-ui'
import { Message } from '@/utils'
import { useModal } from '@/hooks/modal/useModal'
import { downloadFile } from '@/api/modules/public'
import { createGroup } from '@/api/modules/contacts/group'

const { t } = useI18n()

const formRef = useTemplateRef('formRef')

const form = reactive({
	create_type: 2,
	name: '',
	file_data: '',
	file_type: '',
})

const rules: FormRules = {
	name: {
		required: true,
		message: t('market.task.createGroup.validation.groupNameRequired'),
		trigger: ['blur', 'input'],
	},
}

const handleChangeFile = (file: UploadFileInfo) => {
	const reader = new FileReader()
	reader.onload = e => {
		form.file_type = file.name.substring(file.name.lastIndexOf('.') + 1)
		form.file_data = e.target?.result as string
	}
	if (file.file) {
		reader.readAsText(file.file)
	}
}

const handleDownloadTemplate = async () => {
	await downloadFile({ file_path: '/opt/Billion-Mail/data/example_recipients.txt' })
}

const validateForm = async () => {
	await formRef.value?.validate()
	if (isEmpty(form.file_data)) {
		Message.error(t('contacts.subscribers.import.validation.fileRequired'))
		return Promise.reject()
	}
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			form.name = ''
			form.file_data = ''
			form.file_type = ''
		}
	},
	onConfirm: async () => {
		await validateForm()
		await createGroup(toRaw(form))
		const state = modalApi.getState<{ refresh: () => void }>()
		state?.refresh()
	},
})
</script>

<style lang="scss" scoped></style>
