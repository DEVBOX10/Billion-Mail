export interface Template {
	id: number
	temp_name: string
	add_type: number
	content: string
	render: string
	create_time: number
	update_time: number
}

export interface TemplateParams {
	page: number
	page_size: number
	keyword: string
}
