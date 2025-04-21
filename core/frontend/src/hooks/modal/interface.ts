export interface ModalApiOptions {
	/**
	 * @description Shared state data
	 */
	state?: Record<string, unknown>
	/**
	 * @description Independent modal component
	 */
	component?: Component
	/**
	 * @description Modal state change callback
	 */
	onChangeState?: (isOpen: boolean) => void
	/**
	 * @description Click cancel button callback
	 */
	onCancel?: () => void
	/**
	 * @description Click confirm button callback
	 */
	onConfirm?: () => void | Promise<unknown>
	/**
	 * @description Modal close end callback
	 */
	onClose?: () => void
}

export interface ModalApi {
	open: () => void
	close: () => void
	setState: (state: Record<string, unknown>) => void
	getState: <T = Record<string, unknown>>() => T
}

export interface ModalInjectData {
	state: Record<string, unknown>
	extendApi: (api: ModalApi) => void
	setState: (state: Record<string, unknown>) => void
}
