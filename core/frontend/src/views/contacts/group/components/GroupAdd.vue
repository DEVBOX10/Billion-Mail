<template>
	<modal :title="t('contacts.group.add.title')" width="580">
		<div class="pt-12px">
			<bt-form ref="formRef" :model="form" :rules="rules">
				<n-form-item :label="t('contacts.group.form.name')" path="name">
					<div class="w-240px">
						<n-input v-model:value="form.name"></n-input>
					</div>
				</n-form-item>
				<n-form-item :label="t('contacts.group.form.file')" path="file_data">
					<div class="w-360px">
						<bt-file-upload :is-upload="false" :accept="['txt']" @change="handleChangeFile">
						</bt-file-upload>
						<div class="mt-16px text-desc">
							<span class="mr-4px">{{ t('contacts.group.fileUpload.hint') }}</span>
							<n-button text type="primary" @click="handleDownloadTemplate">
								{{ $t('common.actions.download') }}
							</n-button>
						</div>
						<div class="mt-8px text-desc">{{ t('contacts.group.fileUpload.emptyHint') }}</div>
					</div>
				</n-form-item>
			</bt-form>
		</div>
	</modal>
</template>

<script lang="ts" setup>
import { UploadFileInfo } from 'naive-ui'
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

const rules = {
	name: {
		required: true,
		message: t('contacts.group.form.nameRequired'),
		trigger: 'blur',
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
