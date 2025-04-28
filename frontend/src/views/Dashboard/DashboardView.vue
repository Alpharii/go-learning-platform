<script setup lang="ts">
import { onMounted, ref, computed } from 'vue'
import { useAuthStore } from '@/stores/authStores'
import { fetchCourses, type Course } from '@/services/courseService'
import {
  Card,
  CardHeader,
  CardTitle,
  CardDescription,
  CardContent,
} from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
} from '@/components/ui/dropdown-menu'
import { Avatar, AvatarImage, AvatarFallback } from '@/components/ui/avatar'
import { Skeleton } from '@/components/ui/skeleton'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { LogOut } from 'lucide-vue-next'

const auth = useAuthStore()
const courses = ref<Course[]>([])
const isLoading = ref(true)
const error = ref<string | null>(null)

const profileName = computed(() => 
  auth.user?.profile.name || 'Pengguna'
)

const profileInitials = computed(() => 
  profileName.value.split(' ').map(n => n[0]).join('').toUpperCase()
)

const loadCourses = async () => {
  try {
    courses.value = await fetchCourses()
    error.value = null
  } catch (err) {
    error.value = 'Gagal memuat data kursus. Silakan coba lagi nanti.'
    console.error('Failed to load courses:', err)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  loadCourses()
})
</script>

<template>
  <div class="min-h-screen bg-background p-6">
    <!-- Header -->
    <header class="max-w-7xl mx-auto flex justify-between items-center mb-8">
      <div class="space-y-1">
        <h1 class="text-3xl font-bold text-foreground">Dashboard</h1>
        <p class="text-sm text-muted-foreground">
          Selamat datang di platform pembelajaran kami
        </p>
      </div>

      <!-- Profile Dropdown -->
      <DropdownMenu>
        <DropdownMenuTrigger as-child>
          <Button variant="ghost" class="h-10 w-10 rounded-full p-0">
            <Avatar class="h-9 w-9">
              <AvatarImage 
                :src="auth.user?.profile.image" 
                alt="Profile picture" 
              />
              <AvatarFallback class="bg-primary text-primary-foreground">
                {{ profileInitials }}
              </AvatarFallback>
            </Avatar>
          </Button>
        </DropdownMenuTrigger>
        
        <DropdownMenuContent align="end" class="w-56">
          <DropdownMenuItem 
            @click="auth.logout()"
            class="text-destructive cursor-pointer"
          >
            <LogOut class="mr-2 h-4 w-4" />
            Keluar Akun
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>
    </header>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto space-y-8">
      <!-- Welcome Section -->
      <Card>
        <CardHeader>
          <CardTitle class="text-lg">
            👋 Halo, {{ profileName }}!
          </CardTitle>
          <CardDescription>
            Anda memiliki {{ courses.length }} kursus tersedia
          </CardDescription>
        </CardHeader>
      </Card>

      <!-- Courses Grid -->
      <div v-if="isLoading" class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        <Card v-for="n in 3" :key="n">
          <CardContent class="p-6 space-y-4">
            <Skeleton class="h-6 w-3/4" />
            <Skeleton class="h-4 w-full" />
            <Skeleton class="h-4 w-2/3" />
          </CardContent>
        </Card>
      </div>

      <Alert v-else-if="error" variant="destructive">
        <AlertDescription>
          {{ error }}
        </AlertDescription>
      </Alert>

      <div v-else class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
  <Card
    v-for="course in courses"
    :key="course.ID"
    class="hover:shadow-lg transition-shadow"
  >
    <CardHeader class="p-0 -mt-6">
      <!-- Gambar Course (mentok ke atas) -->
      <div class="relative">
        <img
          v-if="course.Image"
          :src="course.Image"
          alt="Course image"
          class="w-full h-40 object-cover rounded-t-md"
        />
      </div>
    </CardHeader>
    <CardContent class="flex flex-col justify-between p-6 pt-0">
      <!-- Title and Description -->
      <div class="space-y-2">
        <CardTitle class="text-base font-semibold">{{ course.Title }}</CardTitle>
        <CardDescription class="line-clamp-3 text-sm text-muted-foreground">
          {{ course.Description || 'Tidak ada deskripsi' }}
        </CardDescription>
      </div>

      <!-- Profile Pembuat -->
      <div class="flex items-center space-x-3 mt-4">
        <Avatar class="h-8 w-8">
          <AvatarImage
            :src="course.User?.Profile?.Image"
            alt="Author image"
          />
          <AvatarFallback class="bg-primary text-primary-foreground">
            {{ course.User?.Profile?.Name?.[0]?.toUpperCase() || '?' }}
          </AvatarFallback>
        </Avatar>
        <span class="text-sm font-medium">
          {{ course.User?.Profile?.Name || 'Unknown' }}
        </span>
      </div>
    </CardContent>

    <CardFooter class="p-6 pt-0">
      <Button variant="outline" size="sm" class="w-full">
        Lihat Detail
      </Button>
    </CardFooter>
  </Card>
</div>


    </main>
  </div>
</template>