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
		}),
	],
	transformers: [transformerDirectives({ enforce: 'pre' })],
	theme: {
		colors: {
			primary: '#20a53a',
			primaryHover: '#1D9534',
			error: '#ef0808',
			warning: '#f0ad4e',
		},
	},
	shortcuts: {
		'flex-center': 'flex items-center justify-center',
	},
})
