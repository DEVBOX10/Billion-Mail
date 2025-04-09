<template>
	<div class="ssl-cert-container">
		<n-alert v-if="domain.cert_info.endtime > 0" class="mb-16px" type="success" :show-icon="false">
			<div class="brand-info">
				<div class="info-label">品牌：</div>
				<div class="info-value">{{ domain.cert_info.issuer }}</div>
				<div class="info-label">证书域名：</div>
				<div class="info-value font-normal!">
					<n-tag v-for="(domain, index) in domains" :key="index" size="small">
						{{ domain }}
					</n-tag>
				</div>
			</div>
		</n-alert>

		<div class="cert-content">
			<div class="cert-section">
				<div class="section-title">私钥 (KEY)</div>
				<n-input
					v-model:value="privateKey"
					type="textarea"
					placeholder=""
					:rows="14"
					:input-props="{ spellcheck: false }">
				</n-input>
			</div>

			<div class="cert-section">
				<div class="section-title">证书 (CRT/PEM)</div>
				<n-input
					v-model:value="certificate"
					type="textarea"
					placeholder=""
					:input-props="{ spellcheck: false }"
					:rows="14">
				</n-input>
			</div>
		</div>

		<div class="cert-actions">
			<n-button type="primary" @click="saveCertificate">保存</n-button>
			<n-button type="primary" @click="applyCertificate">申请免费证书</n-button>
		</div>

		<bt-tips>
			<li>粘贴您的 KEY 和 CRT 内容，然后保存</li>
			<li>如果浏览器提示证书链不完整，请检查 PEM 证书是否正确拼接</li>
			<li>PEM 格式证书 = 域名证书 .crt + 根证书 (root_bundle).crt</li>
			<li>当未指定 SSL 默认站点时，未配置 SSL 的站点将使用 HTTPS 访问启用了 SSL 的站点</li>
		</bt-tips>
	</div>
</template>

<script lang="ts" setup>
import { applyCert, setSsl } from '@/api/modules/domain'
import { MailDomain } from '../../interface'

const { domain, refresh } = defineProps({
	domain: {
		type: Object as PropType<MailDomain>,
		required: true,
	},
	refresh: {
		type: Function as PropType<() => void>,
		required: true,
	},
})

// 数据
const domains = computed(() => {
	return domain.cert_info.dns
})

const privateKey = ref<string>('')

const certificate = ref<string>('')

// 保存证书
const saveCertificate = async () => {
	await setSsl({ domain: domain.domain, key: privateKey.value, certificate: certificate.value })
	refresh()
}

const applyCertificate = async () => {
	await applyCert({ domain: domain.domain })
	refresh()
}

const init = () => {
	privateKey.value = domain.cert_info.key_pem
	certificate.value = domain.cert_info.cert_pem
}

init()
</script>

<style lang="scss" scoped>
.ssl-cert-container {
	.brand-info {
		display: flex;
		align-items: center;
		gap: 8px;

		.info-label {
			font-weight: 500;
			color: #606266;
		}

		.info-value {
			margin-right: 16px;
			font-weight: bold;
			color: #67c23a;
			.n-tag + .n-tag {
				margin-left: 8px;
			}
		}
	}

	.cert-content {
		display: flex;
		gap: 20px;
		margin-bottom: 20px;

		.cert-section {
			flex: 1;

			.section-title {
				font-weight: 500;
				margin-bottom: 8px;
				color: #303133;
			}
		}
	}

	.cert-actions {
		display: flex;
		gap: 12px;
		margin-bottom: 24px;
	}
}
</style>
