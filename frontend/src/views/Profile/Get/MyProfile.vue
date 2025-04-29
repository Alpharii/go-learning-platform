<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStores'
import { Avatar, AvatarImage, AvatarFallback } from '@/components/ui/avatar'
import { Button } from '@/components/ui/button'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Skeleton } from '@/components/ui/skeleton'
import { LogOut, Pencil } from 'lucide-vue-next'
import { getMyProfile } from '@/services/profileServices'

// Define the shape of courses returned by profileServices
interface ProfileCourse {
  id: number
  title: string
  description: string
  image: string
  progress?: number
}

const router = useRouter()
const auth = useAuthStore()

const createdCourses = ref<ProfileCourse[]>([])
const enrolledCourses = ref<ProfileCourse[]>([])
const isLoading = ref(true)
const error = ref<string | null>(null)

const profileName = computed(() => auth.user?.profile?.name || 'Pengguna')
const profileEmail = computed(() => auth.user?.email || 'user@example.com')
const profileImage = computed(() => auth.user?.profile?.image || '')
const profileInitials = computed(() =>
  profileName.value.split(' ').map(n => n[0]).join('').toUpperCase()
)

const loadProfileAndCourses = async () => {
  try {
    const data = await getMyProfile()
    console.log(data)
    createdCourses.value = data.created_courses || []
    enrolledCourses.value = data.enrolled_courses || []
    error.value = null
  } catch (err) {
    error.value = 'Gagal memuat profil. Silakan coba lagi nanti.'
    console.error('Failed to load profile:', err)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  loadProfileAndCourses()
})

function goToEdit() {
  router.push('/profile/update')
}

function logout() {
  auth.logout()
}
</script>

<template>
  <div class="min-h-screen bg-background">
    <main class="max-w-4xl mx-auto pt-6 pb-12">
      <!-- Error Alert -->
      <Alert v-if="error" variant="destructive" class="mb-4">
        <AlertDescription>{{ error }}</AlertDescription>
      </Alert>

      <!-- Profile Cover -->
      <div class="relative rounded-lg overflow-hidden bg-gray-200 h-40"></div>
      <div class="relative -mt-16 flex flex-col items-center">
        <Avatar class="w-32 h-32 border-4 border-white shadow-lg">
          <AvatarImage :src="profileImage" alt="Profile picture" />
          <AvatarFallback class="bg-primary text-primary-foreground">
            {{ profileInitials }}
          </AvatarFallback>
        </Avatar>
        <h1 class="mt-4 text-3xl font-bold text-foreground">{{ profileName }}</h1>
        <p class="text-muted-foreground mb-4">{{ profileEmail }}</p>
        <div class="flex space-x-2">
          <Button variant="outline" size="sm" class="flex items-center gap-1" @click="goToEdit">
            <Pencil class="w-4 h-4" />
            Edit Profil
          </Button>
          <Button variant="ghost" size="sm" class="flex items-center gap-1 text-destructive" @click="logout">
            <LogOut class="w-4 h-4" />
            Keluar
          </Button>
        </div>
      </div>

      <!-- Loading Skeleton -->
      <div v-if="isLoading" class="mt-8 grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        <div v-for="n in 6" :key="n">
          <Skeleton class="h-40 rounded-md mb-4" />
          <Skeleton class="h-6 w-3/4 mb-2" />
          <Skeleton class="h-4 w-full" />
        </div>
      </div>

      <!-- Courses Section -->
      <template v-else>
        <!-- Created Courses -->
        <section class="mt-10">
          <h2 class="text-xl font-semibold text-foreground mb-4">Kursus yang Anda Buat</h2>
          <div class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
            <div
              v-for="course in createdCourses"
              :key="course.id"
              class="bg-card rounded-lg overflow-hidden shadow hover:shadow-lg transition-shadow"
            >
              <img
                v-if="course.image"
                :src="course.image"
                alt="Course image"
                class="w-full h-40 object-cover"
              />
              <div class="p-4">
                <h3 class="text-lg font-medium text-foreground mb-1">{{ course.title }}</h3>
                <p class="text-sm text-muted-foreground line-clamp-2">
                  {{ course.description || 'Tidak ada deskripsi' }}
                </p>
              </div>
            </div>
          </div>
        </section>

        <!-- Enrolled Courses -->
        <section class="mt-10">
          <h2 class="text-xl font-semibold text-foreground mb-4">Kursus yang Anda Ikuti</h2>
          <div class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
            <div
              v-for="course in enrolledCourses"
              :key="course.id"
              class="bg-card rounded-lg overflow-hidden shadow hover:shadow-lg transition-shadow"
            >
              <img
                v-if="course.image"
                :src="course.image"
                alt="Course image"
                class="w-full h-40 object-cover"
              />
              <div class="p-4">
                <div class="flex justify-between items-center mb-1">
                  <h3 class="text-lg font-medium text-foreground">{{ course.title }}</h3>
                  <span class="text-sm text-muted-foreground">{{ course.progress ?? 0 }}%</span>
                </div>
                <p class="text-sm text-muted-foreground line-clamp-2">
                  {{ course.description || 'Tidak ada deskripsi' }}
                </p>
              </div>
            </div>
          </div>
        </section>
      </template>
    </main>
  </div>
</template>
