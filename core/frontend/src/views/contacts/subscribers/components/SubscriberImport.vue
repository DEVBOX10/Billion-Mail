<template>
	<modal :title="$t('contacts.subscribers.import.title')" width="640">
		<div class="pt-12px">
			<bt-form ref="formRef" :model="form" :rules="rules">
				<n-form-item :label="$t('contacts.subscribers.import.group')" path="group_id">
					<div class="w-320px">
						<group-select v-model:value="form.group_ids"></group-select>
					</div>
				</n-form-item>
				<n-form-item :label="$t('contacts.subscribers.import.data')" path="contacts">
					<div class="w-440px">
						<n-radio-group v-model:value="form.import_type">
							<n-radio-button :value="2">
								{{ $t('contacts.subscribers.import.pasteData') }}
							</n-radio-button>
							<n-radio-button :value="1">
								{{ $t('contacts.subscribers.import.uploadFile') }}
							</n-radio-button>
						</n-radio-group>
						<div class="mt-12px">
							<template v-if="form.import_type === 2">
								<n-input
									v-model:value="form.contacts"
									type="textarea"
									:rows="8"
									:placeholder="$t('contacts.subscribers.import.pastePlaceholder')">
								</n-input>
							</template>
							<template v-else-if="form.import_type === 1">
								<bt-file-upload :is-upload="false" :accept="['txt']" @change="handleChangeFile">
								</bt-file-upload>
								<div class="mt-16px text-12px text-desc">
									{{ $t('contacts.subscribers.import.fileTypeHint') }}
									<n-button text type="primary" @click="handleDownloadTemplate">
										{{ $t('common.actions.download') }}
									</n-button>
								</div>
							</template>
						</div>
					</div>
				</n-form-item>
			</bt-form>
		</div>
	</modal>
</template>

<script lang="ts" setup>
import { FormRules, UploadFileInfo } from 'naive-ui'
import { isEmpty } from 'lodash-es'
import { Message } from '@/utils'
import { useModal } from '@/hooks/modal/useModal'
import { importSubscribers } from '@/api/modules/contacts/subscribers'

import GroupSelect from './GroupMultipleSelect.vue'
import { downloadFile } from '@/api/modules/public'

const { t } = useI18n()

const formRef = useTemplateRef('formRef')

const dataType = ref('paste')

const form = reactive({
	group_ids: [] as number[],
	import_type: 2,
	contacts: '',
	file_data: '',
	file_type: '',
})

const rules: FormRules = {
	group_ids: {
		trigger: 'change',
		validator: () => {
			if (form.group_ids.length === 0) {
				return new Error(t('contacts.subscribers.import.validation.groupRequired'))
			}
			return true
		},
	},
	contacts: {
		trigger: ['input', 'blur'],
		validator: () => {
			if (form.import_type === 2 && isEmpty(form.contacts)) {
				return new Error(t('contacts.subscribers.import.validation.dataRequired'))
			}
			return true
		},
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
	if (dataType.value === 'upload' && isEmpty(form.file_data)) {
		Message.error(t('contacts.subscribers.import.validation.fileRequired'))
		return Promise.reject()
	}
}

const [Modal, modalApi] = useModal({
	onChangeState: isOpen => {
		if (isOpen) {
			const state = modalApi.getState<{ group_id: number }>()
			if (state?.group_id) {
				form.group_ids = [state.group_id]
			} else {
				form.group_ids = []
			}
			form.import_type = 2
			form.contacts = ''
			form.file_data = ''
			form.file_type = ''
		}
	},
	onConfirm: async () => {
		await validateForm()
		await importSubscribers(toRaw(form))
		const state = modalApi.getState<{ refresh: () => void }>()
		state?.refresh()
	},
})
</script>

<style lang="scss" scoped>
.upload-trigger {
	display: flex;
	flex-direction: column;
	align-items: center;
	justify-content: center;
	height: 120px;
	border: 1px dashed #d9d9d9;
	border-radius: 2px;
	cursor: pointer;
	transition: border-color 0.3s;

	&:hover {
		border-color: #1890ff;
	}
}
</style>
