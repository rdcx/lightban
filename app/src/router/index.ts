import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('../views/HomeView.vue'),
    },
    {
      name: 'login',
      path: '/login',
      component: () => import('../views/LoginView.vue'),
    },
    {
      name: 'register',
      path: '/register',
      component: () => import('../views/RegisterView.vue'),
    },
  ]
})

export default router
