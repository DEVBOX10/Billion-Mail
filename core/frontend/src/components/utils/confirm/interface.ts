import { Type } from 'naive-ui/es/button/src/interface'

export interface ConfirmOptions {
	title?: string
	content?: string
	confirmText?: string
	confirmType?: Type
	onCancel?: () => unknown
	onConfirm?: () => void | boolean | Promise<unknown>
}
