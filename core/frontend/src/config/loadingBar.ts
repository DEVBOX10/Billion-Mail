import { createDiscreteApi } from 'naive-ui'

const { loadingBar } = createDiscreteApi(['loadingBar'], {
	configProviderProps: {
		themeOverrides: {
			LoadingBar: {
				colorLoading: '#20a53a',
				colorError: '#ef0808',
			},
		},
	},
})

export default loadingBar
