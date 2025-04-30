<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStores'
import {
  Avatar,
  AvatarImage,
  AvatarFallback
} from '@/components/ui/avatar'
import {
  Button
} from '@/components/ui/button'
import {
  Alert,
  AlertDescription
} from '@/components/ui/alert'
import {
  Skeleton
} from '@/components/ui/skeleton'
import {
  Card,
  CardHeader,
  CardContent,
  CardDescription,
  CardTitle,
  CardFooter
} from '@/components/ui/card'
import {
  LogOut,
  Pencil,
  Trash2
} from 'lucide-vue-next'
import { getMyProfile } from '@/services/profileServices'

// Interface untuk data kursus
interface EnrolledCourse {
  id: number
  title: string
  description: string
  image: string
  progress: number
  courseId: number
}

// Setup
const router = useRouter()
const auth = useAuthStore()
const enrolledCourses = ref<EnrolledCourse[]>([])
const isLoading = ref(true)
const error = ref<string | null>(null)

// Data Profil
const profileName = computed(() => auth.user?.profile?.name || 'Pengguna')
const profileEmail = computed(() => auth.user?.email || 'user@example.com')
const profileImage = computed(() => auth.user?.profile?.image || '')
const profileInitials = computed(() =>
  profileName.value.split(' ').map(n => n[0]).join('').toUpperCase()
)

// Muat data profil dan kursus
const loadEnrolledCourses = async () => {
  try {
    const data = await getMyProfile()
    enrolledCourses.value = data.enrolled_courses.map((course: any) => ({
      id: course.id,
      courseId: course.course_id,
      title: course.title,
      description: course.description,
      image: course.image,
      progress: course.progress || 0
    }))
    error.value = null
  } catch (err) {
    error.value = 'Gagal memuat kursus yang diikuti. Silakan coba lagi.'
    console.error('Error fetching enrolled courses:', err)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  loadEnrolledCourses()
})

// Aksi
function goToEdit() {
  router.push('/profile/update')
}

function logout() {
  auth.logout()
}

async function handleUnenroll(courseId: number) {
  if (confirm('Yakin ingin berhenti mengikuti kursus ini?')) {
    try {
    //   await cancelEnrollment(courseId)
      enrolledCourses.value = enrolledCourses.value.filter(c => c.courseId !== courseId)
    } catch (err) {
      alert('Gagal membatalkan pendaftaran kursus.')
      console.error('Error canceling enrollment:', err)
    }
  }
}
</script>

<template>
  <div class="min-h-screen bg-background">
    <main class="max-w-4xl mx-auto pt-6 pb-12 px-4">
      <!-- Error Alert -->
      <Alert v-if="error" variant="destructive" class="mb-4">
        <AlertDescription>{{ error }}</AlertDescription>
      </Alert>

      <!-- Profile Header -->
      <div class="relative">
        <!-- Cover Image -->
        <div class="rounded-lg overflow-hidden bg-gray-200 h-40"></div>

        <!-- Profile Info -->
        <div class="-mt-16 flex flex-col items-center text-center">
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
      </div>

      <!-- Loading Skeleton -->
      <div v-if="isLoading" class="mt-8 grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        <div v-for="n in 3" :key="`skeleton-${n}`" class="flex flex-col space-y-2">
          <Skeleton class="h-40 rounded-md" />
          <Skeleton class="h-6 w-3/4" />
          <Skeleton class="h-4 w-full" />
        </div>
      </div>

      <!-- Konten Utama -->
      <template v-else>
        <!-- Judul -->
        <h2 class="text-xl font-semibold text-foreground mt-10 mb-4">Kursus yang Anda Ikuti</h2>

        <!-- Daftar Kursus -->
        <div class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
          <Card
            v-for="course in enrolledCourses"
            :key="course.id"
            class="hover:shadow-lg transition-shadow flex flex-col h-full"
          >
            <!-- Gambar Kursus -->
            <CardHeader class="p-0 -mt-6">
              <div class="relative">
                <img
                  :src="course.image || 'https://via.placeholder.com/400x160'"
                  alt="Course image"
                  class="w-full h-40 object-cover rounded-t-md"
                />
              </div>
            </CardHeader>

            <!-- Konten Kursus -->
            <CardContent class="flex-1 p-6 pt-4 flex flex-col justify-between">
              <div>
                <CardTitle class="text-base font-semibold mb-2">{{ course.title }}</CardTitle>
                <CardDescription class="line-clamp-3 text-sm text-muted-foreground">
                  {{ course.description || 'Tidak ada deskripsi' }}
                </CardDescription>
              </div>

              <!-- Progres -->
              <div class="mt-4">
                <div class="flex justify-between text-xs mb-1">
                  <span>Progres</span>
                  <span>{{ course.progress }}%</span>
                </div>
                <div class="w-full bg-gray-200 rounded-full h-2">
                  <div
                    class="bg-primary h-2 rounded-full transition-all duration-500"
                    :style="{ width: `${course.progress}%` }"
                  ></div>
                </div>
              </div>
            </CardContent>

            <!-- Aksi -->
            <CardFooter class="p-6 pt-0 flex flex-col space-y-2">
              <Button variant="outline" size="sm" class="w-full" @click="router.push(`/courses/${course.courseId}`)">
                Lihat Detail
              </Button>
              <Button variant="ghost" size="sm" class="w-full text-destructive flex items-center gap-1" @click="handleUnenroll(course.courseId)">
                <Trash2 class="w-4 h-4" />
                Batalkan Pendaftaran
              </Button>
            </CardFooter>
          </Card>
        </div>

        <!-- Jika tidak ada kursus -->
        <div v-if="enrolledCourses.length === 0 && !isLoading" class="text-center py-10">
          <p class="text-muted-foreground">Anda belum mengikuti kursus apa pun.</p>
          <Button class="mt-4" @click="router.push('/courses')">Telusuri Kursus</Button>
        </div>
      </template>
    </main>
  </div>
</template>