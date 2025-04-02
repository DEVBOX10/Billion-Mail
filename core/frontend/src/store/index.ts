import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import useMenuStore from './modules/menu'
import useUserStore from './modules/user'

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

export { useMenuStore, useUserStore }

export default pinia
