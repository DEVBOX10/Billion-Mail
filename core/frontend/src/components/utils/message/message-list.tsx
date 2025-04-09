import { TransitionGroup, withDirectives } from 'vue'
import { zindexable } from 'naive-ui'
import type { MessageItem, MessagePosition, MessageType } from './interface'

import Message from './index'
import './animation.scss'

export default defineComponent({
	name: 'MessageList',
	props: {
		messages: {
			type: Array as PropType<MessageItem[]>,
			default: () => [],
		},
		type: {
			type: String as PropType<MessageType>,
			default: 'success',
		},
		position: {
			type: String as PropType<MessagePosition>,
			default: 'center',
		},
	},
	emits: ['close', 'afterClose'],
	setup() {},
	render() {
		return (
			<>
				{this.messages.map(item => {
					if (!item.mask) return
					return withDirectives(<div class="n-modal-mask"></div>, [
						[zindexable, { zIndex: undefined, enabled: true }],
					])
				})}
				<TransitionGroup
					name="bounce"
					onAfterLeave={() => {
						this.$emit('afterClose')
					}}>
					{this.messages.map(item => {
						return (
							<Message
								key={item.id}
								type={item.type}
								close={item.close}
								isHtml={item.isHtml}
								content={item.content}
								duration={item.duration}
								maxWidth={item.maxWidth}
								onClose={() => {
									this.$emit('close', item.id)
								}}
							/>
						)
					})}
				</TransitionGroup>
			</>
		)
	},
})
