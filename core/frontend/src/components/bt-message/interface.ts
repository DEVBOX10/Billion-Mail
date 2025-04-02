import { VNodeChild } from 'vue'

export type RenderContent = string | (() => VNodeChild)

export const MESSAGE_TYPES = ['info', 'success', 'warning', 'error', 'loading'] as const

export const MESSAGE_POSITION = ['center'] as const

export type MessageType = (typeof MESSAGE_TYPES)[number]

export type MessagePosition = (typeof MESSAGE_POSITION)[number]

export type MessageItem = {
	id: number | string
	content: RenderContent
	type?: MessageType
	mask?: boolean
	close?: boolean
	isHtml?: boolean
	duration?: number
	maxWidth?: number | string
	onClose?: (id: number | string) => void
}

export type MessageMethod = {
	// 显示信息提示
	info: (content: RenderContent | MessageConfig, config?: Partial<MessageConfig>) => MessageReturn
	// 显示成功提示
	success: (
		content: RenderContent | MessageConfig,
		config?: Partial<MessageConfig>
	) => MessageReturn
	// 显示警告提示
	warning: (
		content: RenderContent | MessageConfig,
		config?: Partial<MessageConfig>
	) => MessageReturn
	// 显示错误提示
	error: (content: RenderContent | MessageConfig, config?: Partial<MessageConfig>) => MessageReturn
	// 显示加载中提示
	loading: (
		content: RenderContent | MessageConfig,
		config?: Partial<MessageConfig>
	) => MessageReturn
	// 清空全部提示
	clear: () => void
}

export type MessageReturn = {
	// 关闭当前消息
	close: () => void
}

export type MessageConfig = {
	// 内容
	content: RenderContent
	// 唯一id
	id?: string
	// 是否有背景遮罩
	mask?: boolean
	// 是否关闭
	close?: boolean
	// 消息显示的持续时间
	duration?: number
	// 消息的位置
	position?: MessagePosition
	// 是否为html
	isHtml?: boolean
	maxWidth?: number | string
	// 关闭时的回调函数
	onClose?: (id: number | string) => void
}
