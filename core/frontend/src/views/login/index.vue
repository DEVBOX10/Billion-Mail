<template>
	<div class="login-container">
		<div class="login-card">
			<div class="logo-container">
				<div class="logo">
					<svg viewBox="0 0 24 24" class="logo-icon">
						<path
							d="M16,12V4H17V2H7V4H8V12L6,14V16H11.2V22H12.8V16H18V14L16,12Z"
							fill="currentColor" />
					</svg>
				</div>
			</div>

			<h2 class="login-title">{{ t('login.title') }}</h2>

			<n-form ref="formRef" size="large" :model="form" :rules="rules">
				<n-form-item :show-label="false" path="username">
					<n-input v-model:value="form.username" :placeholder="t('login.form.usernamePlaceholder')">
					</n-input>
				</n-form-item>
				<n-form-item :show-label="false" path="password">
					<n-input
						v-model:value="form.password"
						class="password-input"
						type="password"
						show-password-on="click"
						:placeholder="t('login.form.passwordPlaceholder')"
						@keyup.enter="handleLogin">
					</n-input>
				</n-form-item>
				<n-form-item v-if="isCode" :show-label="false" path="validate_code">
					<n-input
						v-model:value="form.validate_code"
						class="flex-1"
						:placeholder="t('login.form.captcha')"
						:input-props="{ spellcheck: false }"
						@keydown.enter="handleLogin">
					</n-input>
					<n-spin size="small" :show="codeLoading">
						<div class="code" @click="getCode()">
							<img
								class="w-full h-full"
								:src="codeUrl"
								:title="t('login.form.changeCaptcha')"
								:alt="t('login.form.captcha')" />
						</div>
					</n-spin>
				</n-form-item>
				<n-form-item :show-label="false" :show-feedback="false">
					<n-button
						type="primary"
						size="large"
						class="font-bold"
						:loading="loading"
						:disabled="loading"
						:block="true"
						@click="handleLogin">
						{{ t('login.form.loginButton') }}
					</n-button>
				</n-form-item>
			</n-form>
		</div>
	</div>
</template>

<script lang="ts" setup>
import { useUserStore } from '@/store'
import { isObject } from '@/utils'
import { getValidateCode, login } from '@/api/modules/user'

const { t } = useI18n()

const router = useRouter()
const userStore = useUserStore()

const formRef = useTemplateRef('formRef')

const isCode = ref(false)

const codeUrl = ref('')

const codeLoading = ref(false)

const form = reactive({
	username: '',
	password: '',
	validate_code: '',
	validate_code_id: '',
})

const rules = {
	username: {
		required: true,
		message: t('login.validation.usernameRequired'),
		trigger: ['blur', 'input'],
	},
	password: {
		required: true,
		message: t('login.validation.passwordRequired'),
		trigger: ['blur', 'input'],
	},
	validate_code: {
		required: true,
		trigger: ['blur', 'input'],
		message: t('login.validation.captchaRequired'),
	},
}

interface CodeResponse {
	mustValidateCode: boolean
	validateCodeBase64: string
	validateCodeId: string
}

const getCode = async () => {
	try {
		codeLoading.value = true
		const res = await getValidateCode()
		if (isObject<CodeResponse>(res)) {
			isCode.value = res.mustValidateCode
			if (res.mustValidateCode) {
				codeUrl.value = res.validateCodeBase64
				form.validate_code_id = res.validateCodeId
			}
		}
	} finally {
		codeLoading.value = false
	}
}

const loading = ref(false)

interface LoginResponse {
	token: string
	refresh_token: string
	ttl: number
}

const handleLogin = async () => {
	try {
		await formRef.value?.validate()
		loading.value = true
		const res = await login(toRaw(form))
		if (isObject<LoginResponse>(res)) {
			userStore.setLoginInfo({
				token: res.token,
				refresh_token: res.refresh_token,
				ttl: res.ttl,
			})
			setTimeout(() => {
				router.push('/')
			}, 1000)
		}
	} catch {
		getCode()
	} finally {
		loading.value = false
	}
}

getCode()
</script>

<style lang="scss" scoped>
.login-container {
	--primary-color: #00b9f5;
	--primary-hover: #0095c8;
	--text-dark: #333;
	--text-light: #666;
	--accent-green: #25cdb1;
	--accent-purple: #7e6ed5;
	--border-light: #ddd;
}

.login-container {
	display: flex;
	position: relative;
	justify-content: center;
	align-items: center;
	min-height: 100%;
	background: linear-gradient(135deg, #f0f3ff 0%, #e5f8f7 100%);
	overflow: hidden;
	&::before {
		content: '';
		position: absolute;
		top: -15%;
		right: -15%;
		width: 40vw;
		height: 40vw;
		border-radius: 50%;
		background-color: var(--accent-green);
		z-index: 0;
	}
	&::after {
		content: '';
		position: absolute;
		left: -15%;
		bottom: -15%;
		width: 40vw;
		height: 40vw;
		border-radius: 50%;
		background-color: var(--accent-purple);
		z-index: 0;
	}
}

.login-card {
	width: 100%;
	max-width: 400px;
	background-color: #fff;
	padding: 52px 32px 62px;
	border-radius: 8px;
	box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
	z-index: 1;
}

.logo {
	display: flex;
	align-items: center;
	justify-content: center;
	width: 40px;
	height: 40px;
	background-color: var(--primary-color);
	border-radius: 5px;

	&-container {
		display: flex;
		align-items: center;
		justify-content: center;
		margin-bottom: 24px;
	}

	&-icon {
		width: 24px;
		height: 24px;
		color: white;
	}

	&-text {
		font-size: 20px;
		font-weight: 600;
		color: var(--text-dark);
	}
}

.login-title {
	margin-bottom: 32px;
	text-align: center;
	font-weight: 500;
	color: var(--text-dark);
}

.code {
	width: 120px;
	height: 40px;
	margin-left: 12px;
	border-radius: 4px;
	border: 1px solid #dcdfe6;
	overflow: hidden;
	cursor: pointer;
}
</style>
