import ResizeObserver from 'resize-observer-polyfill'
import { withDirectives } from 'vue'
import { NIcon, zindexable } from 'naive-ui'
import { formatLength, render } from 'naive-ui/es/_utils'
import { NBaseLoading } from 'naive-ui/es/_internal'
import { SuccessIcon, ErrorIcon, WarningIcon, InfoIcon } from 'naive-ui/es/_internal/icons'
import type { MessageType, RenderContent } from './interface'
import styles from './index.module.scss'

import BtCloseBtn from '../../base/bt-close-btn/index.vue'

export default defineComponent({
	name: 'BtMessage',
	props: {
		type: {
			type: String as PropType<MessageType>,
			default: 'info',
		},
		content: {
			type: [String, Function] as PropType<RenderContent>,
			default: '',
		},
		duration: {
			type: Number,
			default: 2500,
		},
		close: {
			type: Boolean,
			default: false,
		},
		isHtml: {
			type: Boolean,
			default: false,
		},
		maxWidth: {
			type: [String, Number],
			default: 360,
		},
		zIndex: Number,
	},
	emits: ['close'],
	setup(props, { emit }) {
		const messageRef = ref<HTMLElement | null>(null)

		const typeMap = {
			success: {
				icon: () => <SuccessIcon></SuccessIcon>,
				color: '#20a53a',
			},
			error: {
				icon: () => <ErrorIcon></ErrorIcon>,
				color: '#E85445',
			},
			warning: {
				icon: () => <WarningIcon></WarningIcon>,
				color: '#e6a23c',
			},
			info: {
				icon: () => <InfoIcon></InfoIcon>,
				color: '#2080F0',
			},
			loading: {
				icon: () => <NBaseLoading clsPrefix="bt" strokeWidth={24} scale={0.85}></NBaseLoading>,
				color: '#20a53a',
			},
		}

		const iconInfo = computed(() => {
			return typeMap[props.type]
		})

		let timer = 0

		// Close message box
		const handleClose = () => {
			emit('close')
		}

		const startTimer = () => {
			const { duration } = props
			if (duration > 0) {
				timer = window.setTimeout(handleClose, duration)
			}
		}

		const clearTimer = () => {
			if (timer) {
				window.clearTimeout(timer)
				timer = 0
			}
		}

		// Set to the center
		const setMsgCenter = () => {
			const dom = messageRef.value
			if (!dom) return

			const bodyDom = window.document.body
			const left = (bodyDom.clientWidth - dom.clientWidth) / 2
			const top = (bodyDom.clientHeight - dom.clientHeight) / 2
			dom.style.left = left + 'px'
			dom.style.top = top + 'px'
		}

		let first = true

		// Add window change
		const robserver = new ResizeObserver(() => {
			if (first) {
				first = false
				return
			}
			setMsgCenter()
		})

		// Show message after centering
		const showMsg = () => {
			nextTick(() => {
				const dom = messageRef.value
				if (!dom) return
				dom.style.opacity = '1'
			})
		}

		onMounted(() => {
			setMsgCenter()
			showMsg()
			startTimer()
			robserver.observe(window.document.body)
		})

		onBeforeUnmount(() => {
			clearTimer()
			robserver.disconnect()
		})

		return {
			iconInfo,
			messageRef,
			handleClose,
		}
	},
	render() {
		return withDirectives(
			<div
				ref="messageRef"
				class={[styles['bt-message'], `bt-message-${this.type}`]}
				style={{ maxWidth: formatLength(this.maxWidth) }}>
				{this.close && (
					<BtCloseBtn
						onClick={() => {
							this.handleClose()
						}}
					/>
				)}
				<div class="flex-basis-30px">
					<NIcon size={30} color={this.iconInfo.color}>
						{this.iconInfo.icon()}
					</NIcon>
				</div>
				<div class="flex-1 ml-10px max-h-600px leading-20px text-14px overflow-auto">
					{!this.isHtml && <div>{render(this.content)}</div>}
					{this.isHtml && <div v-html={this.content}></div>}
				</div>
			</div>,
			[[zindexable, { zIndex: this.zIndex, enabled: true }]]
		)
	},
})
