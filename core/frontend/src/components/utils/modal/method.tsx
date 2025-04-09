import { createVNode, render, VNodeChild } from 'vue'
import { render as renderContent } from 'naive-ui/es/_utils'
import Modal from './index.vue'

interface ModalConfig {
	title?: string
	content?: string | (() => VNodeChild)
	component?: () => VNodeChild
}

/**
 * @description 弹出模态框
 * @param option
 */
export const openModal = (option: ModalConfig) => {
	const container = document.createElement('div')
	container.className = 'bt-modal-container'

	const vNode = createVNode(
		Modal,
		{
			show: true,
			title: option.title,
			width: 384,
			onAfterLeave: () => {
				render(null, container)
				document.body.removeChild(container)
			},
		},
		{
			default: () => <>{renderContent(option.content)}</>,
		}
	)

	render(vNode, container)
	document.body.appendChild(container)
}
