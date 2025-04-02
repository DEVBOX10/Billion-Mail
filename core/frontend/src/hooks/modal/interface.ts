export interface ModalApiOptions {
	/**
	 * @description 共享状态数据
	 */
	state?: Record<string, unknown>
	/**
	 * @description 独立的弹窗组件
	 */
	component?: Component
	/**
	 * @description 弹窗状态变化回调
	 */
	onChangeState?: (isOpen: boolean) => void
	/**
	 * @description 点击取消按钮的回调
	 */
	onCancel?: () => void
	/**
	 * @description 点击确定按钮的回调
	 */
	onConfirm?: () => void | Promise<unknown>
	/**
	 * @description 弹窗关闭结束的回调
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
