export default {
	dns: {
		title: 'DNS 设置',
		step1: {
			title: '第一步：设置 MX 记录',
			description: '请在您的域名解析服务商处添加以下 MX 记录',
		},
		step2: {
			title: '第二步：设置 TXT 记录',
			description: '请在您的域名解析服务商处添加以下 TXT 记录',
		},
		step3: {
			title: '第三步：等待生效',
			description1: '域名解析记录通常会在几分钟内生效，最长可能需要 48 小时',
			description2: '您可以点击下方按钮验证域名解析是否生效',
		},
		table: {
			recordType: '记录类型',
			host: '主机记录',
			value: '记录值',
			mxPriority: '优先级',
		},
		status: {
			ok: '已设置',
			notSet: '未设置',
		},
		columns: {
			domain: '域名',
		},
		verify: '已设置，验证域名解析',
	},
	api: {
		loading: {
			creating: '正在创建域名...',
			updating: '正在更新域名...',
			deleting: '正在删除域名...',
			settingSSL: '正在设置 SSL 证书...',
			applyingCert: '正在申请证书...',
			refreshingDns: '正在刷新DNS记录，请稍候...',
		},
	},
}
