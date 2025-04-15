import { createVNode, render } from 'vue'
import { app } from '@/index'
import type { ConfirmOptions } from './interface'

import Modal from '@/components/ui/bt-modal/index.vue'
import ConfigProvider from '@/components/ui/bt-config-provider/index.vue'

/**
 * @description 命令式调用确认对话框
 * @param options 确认对话框选项
 */
export function confirm(options: ConfirmOptions) {
	const container = document.createElement('div')

	const removeModalElement = () => {
		render(null, container)
		if (container.parentNode) {
			container.parentNode.removeChild(container)
		}
	}

	// 创建虚拟节点
	const modalNode = createVNode(
		Modal,
		{
			show: true,
			width: 420,
			title: options.title,
			onCancel: options.onCancel,
			onConfirm: options.onConfirm,
			confirmText: options.confirmText,
			confirmType: options.confirmType,
			onAfterLeave: () => {
				removeModalElement()
			},
		},
		{
			default: () => <div class="text-14px">{options.content}</div>,
		}
	)

	// 创建配置提供者节点
	const configNode = createVNode(
		ConfigProvider,
		{},
		{
			default: () => [modalNode],
		}
	)

	configNode.appContext = app._context

	// 渲染到 Modal
	render(configNode, container)
}
