export interface Group {
	id: number
	name: string
	description: string
	create_time: number
	update_time: number
	total_count: number
	active_count: number
	unsubscribe_count: number
}

export interface GroupParams {
	page: number
	page_size: number
	keyword: string
}
