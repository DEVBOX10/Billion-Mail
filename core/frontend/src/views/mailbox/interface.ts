export interface MailBoxParams {
	page: number
	page_size: number
	keyword: string
	domain: string | null
}

export interface MailBox {
	username: string
	password: string
	password_encode: string
	full_name: string
	is_admin: number
	maildir: string
	quota: number
	local_part: string
	domain: string
	create_time: number
	update_time: number
	active: number
}
