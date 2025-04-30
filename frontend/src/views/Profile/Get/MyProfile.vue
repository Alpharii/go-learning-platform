<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
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
  Pencil 
} from 'lucide-vue-next'

import { getMyProfile } from '@/services/profileServices'

// Define types
interface ProfileCourse {
  id: number
  title: string
  description: string
  image: string
  progress?: number
}

// Hooks & Stores
const router = useRouter()
const auth = useAuthStore()

// Reactive states
const createdCourses = ref<ProfileCourse[]>([])
const enrolledCourses = ref<ProfileCourse[]>([])
const isLoading = ref(true)
const error = ref<string | null>(null)

// Computed properties
const profileName = computed(() => auth.user?.profile?.name || 'Pengguna')
const profileEmail = computed(() => auth.user?.email || 'user@example.com')
const profileImage = computed(() => auth.user?.profile?.image || '')
const profileInitials = computed(() =>
  profileName.value.split(' ').map(n => n[0]).join('').toUpperCase()
)

// Methods
const loadProfileAndCourses = async () => {
  try {
    const data = await getMyProfile()
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

function goToEdit() {
  router.push('/profile/update')
}

function logout() {
  auth.logout()
}

// Lifecycle
onMounted(() => {
  loadProfileAndCourses()
})
</script>

<template>
  <div class="min-h-screen bg-background">
    <main class="max-w-4xl mx-auto pt-6 pb-12">
      <!-- Error Notification -->
      <Alert v-if="error" variant="destructive" class="mb-4">
        <AlertDescription>{{ error }}</AlertDescription>
      </Alert>

      <!-- Profile Header -->
      <div class="relative">
        <!-- Cover Image -->
        <div class="rounded-lg overflow-hidden bg-gray-200 h-40"></div>
        
        <!-- Profile Info -->
        <div class="-mt-16 flex flex-col items-center">
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
            <Button 
              variant="ghost" 
              size="sm" 
              class="flex items-center gap-1 text-destructive" 
              @click="logout"
            >
              <LogOut class="w-4 h-4" />
              Keluar
            </Button>
          </div>
        </div>
      </div>

      <!-- Loading Skeleton -->
      <div v-if="isLoading" class="mt-8 grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        <div v-for="n in 6" :key="`skeleton-${n}`" class="flex flex-col space-y-2">
          <Skeleton class="h-40 rounded-md" />
          <Skeleton class="h-6 w-3/4" />
          <Skeleton class="h-4 w-full" />
        </div>
      </div>

      <!-- Content -->
      <template v-else>
        <!-- Created Courses Section -->
        <section class="mt-10">
          <h2 class="text-xl font-semibold text-foreground mb-4">Kursus yang Anda Buat</h2>
          <div class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
            <Card 
              v-for="course in createdCourses" 
              :key="`created-${course.id}`"
              class="hover:shadow-lg transition-shadow flex flex-col"
            >
              <!-- Course Image -->
              <CardHeader class="p-0 -mt-6">
                <div class="relative">
                  <img
                    v-if="course.image"
                    :src="course.image"
                    alt="Course image"
                    class="w-full h-40 object-cover rounded-t-md"
                  />
                </div>
              </CardHeader>

              <!-- Course Content -->
              <CardContent class="flex-1 p-6 pt-0">
                <div class="space-y-2">
                  <CardTitle class="text-base font-semibold">{{ course.title }}</CardTitle>
                  <CardDescription class="line-clamp-3 text-sm text-muted-foreground">
                    {{ course.description || 'Tidak ada deskripsi' }}
                  </CardDescription>
                </div>
                
                <!-- Author Info -->
                <div class="flex items-center space-x-3 mt-4 pt-4 border-t">
                  <Avatar class="h-8 w-8">
                    <AvatarImage :src="profileImage" alt="Author image" />
                    <AvatarFallback class="bg-primary text-primary-foreground">
                      {{ profileInitials }}
                    </AvatarFallback>
                  </Avatar>
                  <span class="text-sm font-medium">{{ profileName }}</span>
                </div>
              </CardContent>

              <!-- Action Button -->
              <CardFooter class="p-6 pt-0">
                <Button variant="outline" size="sm" class="w-full">Lihat Detail</Button>
              </CardFooter>
            </Card>
          </div>
        </section>

        <!-- Enrolled Courses Section -->
        <section class="mt-10">
          <h2 class="text-xl font-semibold text-foreground mb-4">Kursus yang Anda Ikuti</h2>
          <div class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
            <Card 
              v-for="course in enrolledCourses" 
              :key="`enrolled-${course.id}`"
              class="hover:shadow-lg transition-shadow flex flex-col"
            >
              <!-- Course Image -->
              <CardHeader class="p-0 -mt-6">
                <div class="relative">
                  <img
                    v-if="course.image"
                    :src="course.image"
                    alt="Course image"
                    class="w-full h-40 object-cover rounded-t-md"
                  />
                </div>
              </CardHeader>

              <!-- Course Content -->
              <CardContent class="flex-1 p-6 pt-0">
                <div class="space-y-2">
                  <CardTitle class="text-base font-semibold">{{ course.title }}</CardTitle>
                  <div class="flex justify-between items-start">
                    <CardDescription class="line-clamp-3 text-sm text-muted-foreground">
                      {{ course.description || 'Tidak ada deskripsi' }}
                    </CardDescription>
                    <span class="text-sm font-medium text-primary">
                      {{ course.progress ?? 0 }}%
                    </span>
                  </div>
                </div>
              </CardContent>

              <!-- Action Button -->
              <CardFooter class="p-6 pt-0">
                <Button variant="outline" size="sm" class="w-full">Lihat Detail</Button>
              </CardFooter>
            </Card>
          </div>
        </section>
      </template>
    </main>
  </div>
</template>