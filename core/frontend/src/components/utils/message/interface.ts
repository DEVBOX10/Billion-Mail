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
	// Show information prompt
	info: (content: RenderContent | MessageConfig, config?: Partial<MessageConfig>) => MessageReturn
	// Show success prompt
	success: (
		content: RenderContent | MessageConfig,
		config?: Partial<MessageConfig>
	) => MessageReturn
	// Show warning prompt
	warning: (
		content: RenderContent | MessageConfig,
		config?: Partial<MessageConfig>
	) => MessageReturn
	// Show error prompt
	error: (content: RenderContent | MessageConfig, config?: Partial<MessageConfig>) => MessageReturn
	// Show loading prompt
	loading: (
		content: RenderContent | MessageConfig,
		config?: Partial<MessageConfig>
	) => MessageReturn
	// Clear all prompts
	clear: () => void
}

export type MessageReturn = {
	// Close the current message
	close: () => void
}

export type MessageConfig = {
	// Content
	content: RenderContent
	// Unique id
	id?: string
	// Whether to have a background mask
	mask?: boolean
	// Whether to close
	close?: boolean
	// Message display duration
	duration?: number
	// Message position
	position?: MessagePosition
	// Whether to be html
	isHtml?: boolean
	maxWidth?: number | string
	// Callback function when closing
	onClose?: (id: number | string) => void
}
