import { defineConfig } from 'unocss'
import presetUno from '@unocss/preset-uno'
import presetIcons from '@unocss/preset-icons'
import presetAttributify from '@unocss/preset-attributify'
import transformerDirectives from '@unocss/transformer-directives'

export default defineConfig({
	presets: [
		presetUno({ preflight: false }),
		presetAttributify(),
		presetIcons({
			prefix: 'i-',
			extraProperties: {
				display: 'inline-flex',
			},
			collections: {
				mdi: () => import('@iconify-json/mdi/icons.json').then(i => i.default),
			},
		}),
	],
	transformers: [transformerDirectives({ enforce: 'pre' })],
	theme: {
		colors: {
			primary: '#20a53a',
			primaryHover: '#1D9534',
			error: '#ef0808',
			warning: '#f0ad4e',
			desc: '#999',
		},
	},
	shortcuts: {
		'flex-center': 'flex items-center justify-center',
	},
})
