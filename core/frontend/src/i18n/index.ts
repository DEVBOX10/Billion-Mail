import { createI18n } from 'vue-i18n'
import en from './lang/en.json'
import zh from './lang/zh.json'

const i18n = createI18n({
	legacy: false,
	globalInjection: true,
	locale: 'en',
	fallbackLocale: 'en',
	messages: {
		en,
		zh,
	},
})

export default i18n
