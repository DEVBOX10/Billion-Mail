export interface MailDomainParams {
	page: number
	page_size: number
	keyword: string
}

export interface MailDomain {
	domain: string
	a_record: string
	mailboxes: number
	mailbox_quota: number
	quota: number
	rate_limit: number
	create_time: number
	active: number
	dns_records: DomainDnsRecords
}

export type DomainDnsRecords = {
	[key: string]: {
		type: string
		host: string
		value: string
		valid: boolean
	}
}
