import { createApp } from 'vue'
import router from '@/router'
import pinia from '@/store'
import i18n from '@/i18n'
import App from '@/App.vue'

import '@unocss/reset/normalize.css'
import 'uno.css'
import '@/styles/index.scss'

const app = createApp(App)
app.use(router)
app.use(pinia)
app.use(i18n)
app.mount('#root')
