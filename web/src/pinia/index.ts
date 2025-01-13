import { createPinia } from 'pinia'
import { useAppStore } from './modules/app'

const store = createPinia()

export { store, useAppStore }
