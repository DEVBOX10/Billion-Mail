import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'
import useMenuStore from './modules/menu'

const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

export { useMenuStore }

export default pinia
