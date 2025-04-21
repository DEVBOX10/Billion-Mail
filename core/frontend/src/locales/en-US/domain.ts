export default {
	dns: {
		title: 'DNS Settings',
		step1: {
			title: 'Step 1: Set MX Record',
			description: 'Please add the following MX record at your domain registrar',
		},
		step2: {
			title: 'Step 2: Set TXT Records',
			description: 'Please add the following TXT records at your domain registrar',
		},
		step3: {
			title: 'Step 3: Wait for Propagation',
			description1:
				'DNS records typically take a few minutes to propagate, but may take up to 48 hours',
			description2: 'You can click the button below to verify if the DNS records are effective',
		},
		table: {
			recordType: 'Record Type',
			host: 'Host',
			value: 'Value',
			mxPriority: 'Priority',
		},
		status: {
			ok: 'Set',
			notSet: 'Not Set',
		},
		columns: {
			domain: 'Domain',
		},
		verify: 'Records set, verify DNS resolution',
	},
	api: {
		loading: {
			creating: 'Creating domain...',
			updating: 'Updating domain...',
			deleting: 'Deleting domain...',
			settingSSL: 'Setting SSL certificate...',
			applyingCert: 'Applying for certificate...',
			refreshingDns: 'Refreshing DNS records, please wait...',
		},
	},
}
