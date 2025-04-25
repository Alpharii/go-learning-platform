<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useAuthStore } from '@/stores/authStores'
import { LogOut } from 'lucide-vue-next'

interface Course {
  id: number
  title: string
  description: string
}

const auth = useAuthStore()
const courses = ref<Course[]>([])

const fetchCourses = async () => {
  const res = await fetch('http://localhost:8080/courses', {
    headers: {
      Authorization: `Bearer ${auth.token}`,
    },
  })
  const data = await res.json()
  courses.value = data
}

onMounted(() => {
  fetchCourses()
})
</script>

<template>
  <div class="min-h-screen bg-gray-50 py-10 px-4">
    <div class="max-w-4xl mx-auto space-y-6">
      <div class="flex items-center justify-between">
        <h1 class="text-3xl font-bold text-slate-800">Dashboard</h1>
        <button @click="auth.logout()" class="text-red-600 flex items-center gap-1 hover:underline">
          <LogOut class="w-4 h-4" /> Logout
        </button>
      </div>

      <div>
        <p class="text-muted-foreground">Selamat datang, {{ auth.user?.name || 'User' }}</p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div
          v-for="course in courses"
          :key="course.id"
          class="rounded-xl border border-slate-200 bg-white p-5 shadow-sm hover:shadow-md transition-shadow"
        >
          <h2 class="text-xl font-semibold text-slate-800">{{ course.title }}</h2>
          <p class="text-sm text-slate-500">{{ course.description }}</p>
        </div>
      </div>
    </div>
  </div>
</template>
