import { useAuthStore } from '@/stores/authStores'
import CreateCourse from '@/views/Course/Create/CreateCourse.vue'
import DetailCourse from '@/views/Course/Detail/DetailCourse.vue'
import EditCourse from '@/views/Course/Edit/EditCourse.vue'
import CourseManagement from '@/views/Course/Management/CourseManagement.vue'
import MyCourse from '@/views/Course/MyCourse/MyCourse.vue'
import DashboardView from '@/views/Dashboard/DashboardView.vue'
import HomeView from '@/views/HomeView/HomeView.vue'
import AddLesson from '@/views/Lesson/Create/AddLesson.vue'
import LessonManagement from '@/views/Lesson/Create/Management/LessonManagement.vue'
import LoginView from '@/views/LoginView/LoginView.vue'
import CreateProfile from '@/views/Profile/Create/CreateProfile.vue'
import MyProfile from '@/views/Profile/Get/MyProfile.vue'
import UpdateProfile from '@/views/Profile/Update/UpdateProfile.vue'
import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: DashboardView,
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
    {
      path: '/profile/create',
      name: 'CreateProfile',
      component: CreateProfile,
      meta: { requiresAuth: true }, // Pastikan hanya pengguna autentikasi yang bisa mengakses
    },
    {
      path: '/profile/update',
      name: 'UpdateProfile',
      component: UpdateProfile,
      meta: { requiresAuth: true }, // Pastikan hanya pengguna autentikasi yang bisa mengakses
    },
    {
      path: '/myProfile',
      name: 'MyProfile',
      component: MyProfile,
      meta: { requiresAuth: true }, // Pastikan hanya pengguna autentikasi yang bisa mengakses
    },
    {
      path: '/my-course',
      name: 'MyCourse',
      component: MyCourse,
      meta: { requiresAuth: true }, // Pastikan hanya pengguna autentikasi yang bisa mengakses
    },
    {
      path: '/course/create',
      name: 'CreateCourse',
      component: CreateCourse,
      meta: { requiresAuth: true }, // Pastikan hanya pengguna autentikasi yang bisa mengakses
    },
    {
      path: '/courses/:id',
      name: 'CourseDetail',
      component: () => DetailCourse,
      meta: { requiresAuth: true },
    },
    {
      path: '/courses/edit/:id',
      name: 'EditCourse',
      component: () => EditCourse,
      meta: { requiresAuth: true }
    },
    {
      path: '/courses/management',
      name: 'CourseManagement',
      component: () => CourseManagement,
      meta: { requiresAuth: true }
    },
    {
      path: '/courses/:id/lessons/new',
      name: 'AddLesson',
      component: AddLesson,
      meta: { requiresAuth: true }
    },
    {
      path: '/courses/:id/lessons',
      name: 'LessonManagement',
      component: () => LessonManagement,
      meta: { requiresAuth: true }
    },
    {
      path: '/courses/:id/lessons/:lessonId/edit',
      name: 'EditLesson',
      component: () => CourseManagement,
      meta: { requiresAuth: true }
    },
  ],
})

router.beforeEach((to, _, next) => {
  const auth = useAuthStore()

  // Jika pengguna sudah di halaman CreateProfile, jangan redirect lagi
  if (to.name === 'CreateProfile') {
    console.log('Already on CreateProfile route, skipping redirection')
    return next()
  }

  if (to.meta.requiresAuth && !auth.token) {
    console.log('Redirecting to Login because token is missing')
    next({ name: 'Login' }) // Redirect ke halaman login jika tidak ada token
  } else if (to.meta.requiresAuth && (!auth.user || !auth.user.profile?.name)) {
    console.log('Redirecting to CreateProfile because user or name is missing')
    next({ name: 'CreateProfile' }) // Redirect ke halaman pembuatan profil jika nama kosong
  } else {
    console.log('Proceeding to route:', to.name)
    next() // Lanjutkan ke rute yang diminta
  }
})
export default router
