import { createVNode, render } from 'vue'
import { isFunction, isString, isUndefined } from '@/utils'
import {
	MESSAGE_TYPES,
	MessageType,
	MessageItem,
	MessageConfig,
	MessageMethod,
	MessagePosition,
} from './interface'
import MessageList from './message-list'

type _MessageConfig = MessageConfig & {
	type: MessageType
}

class MessageManger {
	private readonly messageIds: Set<number | string>

	private readonly messages: Ref<MessageItem[]>

	private readonly position: MessagePosition

	private container: HTMLElement | null

	private messageCount = 0

	constructor(config: _MessageConfig) {
		const { position = 'center' } = config
		const container = document.createElement('div')
		container.classList.add('bt-message-container')

		this.container = container
		this.messages = ref([])
		this.messageIds = new Set()
		this.position = position

		const vm = createVNode(MessageList, {
			position,
			messages: this.messages.value,
			type: config.type,
			onClose: this.remove,
			onAfterClose: this.destroy,
		})

		render(vm, this.container)
		document.body.appendChild(this.container)
	}

	add = (config: MessageConfig) => {
		this.messageCount++
		const id = config.id ?? `__bt_message_${this.messageCount}`
		if (this.messageIds.has(id)) {
			return this.update(id, config)
		}
		const message = reactive<MessageItem>({ id, ...config })
		this.messages.value.push(message)
		this.messageIds.add(id)
		return {
			close: () => this.remove(id),
		}
	}

	update = (id: number | string, config: MessageConfig) => {
		for (let i = 0; i < this.messages.value.length; i++) {
			if (this.messages.value[i].id === id) {
				const resetOnUpdate = !isUndefined(config.duration)
				Object.assign(this.messages.value[i], { ...config, id, resetOnUpdate })
				break
			}
		}
		return {
			close: () => this.remove(id),
		}
	}

	remove = (id: number | string) => {
		for (let i = 0; i < this.messages.value.length; i++) {
			const item = this.messages.value[i]
			if (item.id === id) {
				if (isFunction(item.onClose)) {
					item.onClose(id)
				}

				this.messages.value.splice(i, 1)
				this.messageIds.delete(id)
				break
			}
		}
	}

	clear = () => {
		this.messages.value.splice(0)
	}

	destroy = () => {
		if (this.messages.value.length === 0 && this.container) {
			render(null, this.container)
			document.body.removeChild(this.container)
			this.container = null
			messageInstance[this.position] = undefined
		}
	}
}

const messageInstance: { center?: MessageManger } = {}

export const Message = MESSAGE_TYPES.reduce((pre, value) => {
	pre[value] = (content, config) => {
		if (isString(content) || isFunction(content)) {
			if (config) {
				config.content = content
			} else {
				config = { content }
			}
		} else {
			config = { ...content }
		}

		if (value === 'loading') {
			config.mask = config.mask !== undefined ? config.mask : true
			config.duration = config.duration ? config.duration : 0
		} else if (config.close) {
			config.mask = config.mask !== undefined ? config.mask : true
			config.duration = config.duration ? config.duration : 0
		}

		const _config = { type: value, ...config } as _MessageConfig
		const { position = 'center' } = _config
		if (!messageInstance[position]) {
			messageInstance[position] = new MessageManger(_config)
		}
		return messageInstance[position].add(_config)
	}
	return pre
}, {} as MessageMethod)

Message.clear = (position?: MessagePosition) => {
	if (position) {
		messageInstance[position]?.clear()
	} else {
		Object.values(messageInstance).forEach(item => item?.clear())
	}
}
