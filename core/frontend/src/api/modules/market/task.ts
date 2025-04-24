import { instance } from '@/api'
import type { TaskParams } from '@/views/market/task/interface'

/**
 * 获取任务列表
 */
export function getTaskList(params: TaskParams) {
	return instance.get('/batch_mail/task/list', { params })
}

interface TaskAddParams {
	track_open: number
	track_click: number
	addresser: string
	full_name: string
	subject: string
	group_ids: number[]
	template_id: number
	is_record: number
	unsubscribe: number
	threads: number
	start_time: number
	remark: string
}

/**
 * 获取任务列表
 */
export function addTask(params: TaskAddParams) {
	return instance.post('/batch_mail/task/create', params, {
		fetchOptions: {
			loading: 'Creating task, please wait...',
			successMessage: true,
		},
	})
}

/**
 * 删除任务
 */
export function deleteTask(params: { id: number }) {
	return instance.post('/batch_mail/task/delete', params, {
		fetchOptions: {
			loading: 'Deleting task, please wait...',
			successMessage: true,
		},
	})
}

/**
 * 暂停任务
 */
export function pauseTask(params: { task_id: number }) {
	return instance.post('/batch_mail/task/pause', params, {
		fetchOptions: {
			loading: 'Pausing task, please wait...',
			successMessage: true,
		},
	})
}

/**
 * 恢复任务
 */
export function resumeTask(params: { task_id: number }) {
	return instance.post('/batch_mail/task/resume', params, {
		fetchOptions: {
			loading: 'Resuming task, please wait...',
			successMessage: true,
		},
	})
}

/**
 * 恢复任务
 */
export function sendTestEmail(params: {
	addresser: string
	subject: string
	recipient: string
	template_id: number
}) {
	return instance.post('/batch_mail/task/send_test', params, {
		fetchOptions: {
			loading: 'Sending test email, please wait...',
			successMessage: true,
		},
	})
}
