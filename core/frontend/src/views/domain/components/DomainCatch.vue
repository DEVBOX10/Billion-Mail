<template>
	<modal title="捕获邮件" :width="540">
		<div class="py-20px">
			<bt-form :model="form">
				<n-form-item label="转发邮件" :show-feedback="false">
					<div class="w-290px">
						<n-input v-model:value="form.email" placeholder="捕获不存在的邮件，转发至此邮件">
						</n-input>
					</div>
				</n-form-item>
			</bt-form>
		</div>
	</modal>
</template>

<script lang="ts" setup>
import { useModal } from '@/hooks/modal/useModal'

const form = reactive({
	email: '',
})

const [Modal] = useModal({
	onConfirm: async () => {},
})
</script>
