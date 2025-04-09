import { ModalApiOptions, ModalApi, ModalInjectData } from './interface'
import BtModal from '@/components/ui/bt-modal/index.vue'

const MODAL_INJECT_KEY = Symbol('MODAL_INJECT')

export const useModal = (options: ModalApiOptions = {}) => {
	const { state = {}, component, onCancel, onConfirm } = options

	const extendedApi = reactive<Partial<ModalApi>>({})

	if (component) {
		const Modal = defineComponent({
			setup(props, { attrs, slots }) {
				provide(MODAL_INJECT_KEY, {
					state,
					options,
					extendApi: (api: ModalApi) => {
						Object.setPrototypeOf(extendedApi, api)
					},
					setState: (val: Record<string, unknown>) => {
						Object.assign(state, val)
					},
				})
				return () => h(component, { ...props, ...attrs }, slots)
			},
		})
		return [Modal, extendedApi as ModalApi] as const
	}

	const Modal = defineComponent({
		setup(props, { attrs, slots }) {
			const injectData = inject<ModalInjectData>(MODAL_INJECT_KEY)

			const show = ref(false)

			extendedApi.open = () => {
				show.value = true
				options.onChangeState?.(show.value)
			}

			extendedApi.close = () => {
				show.value = false
			}

			extendedApi.setState = (val: Record<string, unknown>) => {
				injectData?.setState(val)
				options.onChangeState?.(show.value)
			}

			extendedApi.getState = <T = Record<string, unknown>,>() => {
				return (injectData?.state ?? {}) as T
			}

			injectData?.extendApi(extendedApi as ModalApi)

			return () => (
				<BtModal
					show={show.value}
					{...props}
					{...attrs}
					onCancel={onCancel}
					onConfirm={onConfirm}
					onUpdate:show={val => {
						show.value = val ?? false
					}}>
					{slots.default?.()}
				</BtModal>
			)
		},
	})

	return [Modal, extendedApi as ModalApi] as const
}
