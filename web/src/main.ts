import './assets/main.css'

import { createApp } from 'vue'

import { VueQueryPlugin } from '@tanstack/vue-query'
import naive from 'naive-ui'
import App from './App.vue'
import { store } from './pinia'
import router from './router'

const app = createApp(App)

app.use(VueQueryPlugin)
app.use(naive)
app.use(store)
app.use(router)

app.mount('#app')
