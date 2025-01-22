import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/home.vue'
import WhatIsBlockchainView from '../views/what-is-blockchain.vue'
const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/what-is-blockchain',
      name: 'what-is-blockchain',
      component: WhatIsBlockchainView,
    },
  ],
})

export default router
