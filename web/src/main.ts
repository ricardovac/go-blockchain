import './assets/main.css'

import { createApp } from 'vue'

import { VueQueryPlugin } from '@tanstack/vue-query'
import naive from 'naive-ui'
import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(VueQueryPlugin)
app.use(naive)
app.use(router)

app.mount('#app')
