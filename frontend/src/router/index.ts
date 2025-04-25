import { useAuthStore } from '@/stores/authStores'
import HomeView from '@/views/HomeView/HomeView.vue'
import LoginView from '@/views/LoginView/LoginView.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: HomeView,
      meta: { requiresAuth: true },
    },
    {
      path: '/about',
      name: 'About',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('../views/AboutView.vue'),
      meta: { requiresAuth: true },
    },
    {
      path: '/login',
      name: 'Login',
      component: LoginView,
    },
  ],
})

router.beforeEach((to, _, next) => {
  const auth = useAuthStore()
  console.log(auth.token)
  if (to.meta.requiresAuth && !auth.token) {
    next({ name: 'Login' })
  } else {
    next()
  }
})

export default router
