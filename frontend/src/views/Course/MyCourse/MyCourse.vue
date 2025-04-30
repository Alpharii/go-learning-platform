<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStores'
import { Button } from '@/components/ui/button'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Skeleton } from '@/components/ui/skeleton'
import { Card, CardHeader, CardContent, CardDescription, CardTitle, CardFooter } from '@/components/ui/card'
import { LogOut, Pencil, Trash2 } from 'lucide-vue-next'
import { getMyProfile } from '@/services/profileServices'
import { cancelEnrollment } from '@/services/enrollServices'
import { deleteCourse } from '@/services/courseService'

// Interface untuk data kursus beserta enrollment ID
interface EnrolledCourse {
  enrollmentId: number
  courseId: number
  title: string
  description: string
  image: string
  progress: number
}

// Setup
const router = useRouter()
const auth = useAuthStore()
const createdCourses = ref<EnrolledCourse[]>([])
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
const loadCourses = async () => {
  try {
    const data = await getMyProfile()
    createdCourses.value = data.created_courses.map((c: any) => ({
      enrollmentId: 0,
      courseId: c.id,
      title: c.title,
      description: c.description,
      image: c.image,
      progress: 0
    }))
    enrolledCourses.value = data.enrolled_courses.map((e: any) => ({
      enrollmentId: e.id,
      courseId: e.course_id,
      title: e.title,
      description: e.description,
      image: e.image,
      progress: e.progress || 0
    }))
    error.value = null
  } catch (err) {
    error.value = 'Gagal memuat data kursus. Silakan coba lagi.'
    console.error('Error fetching courses:', err)
  } finally {
    isLoading.value = false
  }
}

onMounted(() => {
  loadCourses()
})

// Aksi
function goToEditProfile() {
  router.push('/profile/update')
}

function logout() {
  auth.logout()
}

async function handleUnenroll(enrollmentId: number) {
  if (!confirm('Yakin ingin berhenti mengikuti kursus ini?')) return
  try {
    console.log(enrollmentId)
    await cancelEnrollment(enrollmentId)
    enrolledCourses.value = enrolledCourses.value.filter(c => c.enrollmentId !== enrollmentId)
  } catch (err) {
    alert('Gagal membatalkan pendaftaran kursus.')
    console.error('Error canceling enrollment:', err)
  }
}

async function handleDeleteCourse(courseId: number) {
  if (!confirm('Yakin ingin menghapus kursus ini?')) return
  try {
    await deleteCourse(courseId)
    createdCourses.value = createdCourses.value.filter(c => c.courseId !== courseId)
  } catch (err) {
    alert('Gagal menghapus kursus.')
    console.error('Delete course error:', err)
  }
}

function viewDetail(courseId: number) {
  router.push(`/courses/${courseId}`)
}

function editCourse(courseId: number) {
  router.push(`/courses/edit/${courseId}`)
}
</script>

<template>
  <div class="min-h-screen bg-background">
    <main class="max-w-4xl mx-auto pt-6 pb-12 px-4">
      <!-- Error Alert -->
      <Alert v-if="error" variant="destructive" class="mb-4">
        <AlertDescription>{{ error }}</AlertDescription>
      </Alert>

      <!-- Loading Skeleton -->
      <div v-if="isLoading" class="mt-8 grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        <div v-for="n in 3" :key="`skeleton-${n}`" class="flex flex-col space-y-2">
          <Skeleton class="h-40 rounded-md" />
          <Skeleton class="h-6 w-3/4" />
          <Skeleton class="h-4 w-full" />
        </div>
      </div>

      <!-- Konten -->
      <template v-else>
        <!-- Kursus Buatan Saya -->
        <section class="mt-10">
          <h2 class="text-xl font-semibold text-foreground mb-4">Kursus yang Anda Buat</h2>
          <div v-if="createdCourses.length" class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
            <Card
              v-for="course in createdCourses"
              :key="`created-${course.courseId}`"
              class="hover:shadow-lg transition-shadow flex flex-col h-full"
            >
              <CardHeader class="p-0 -mt-6">
                <img
                  v-if="course.image"
                  :src="course.image"
                  alt="Course image"
                  class="w-full h-40 object-cover rounded-t-md"
                />
              </CardHeader>

              <CardContent class="flex-1 p-6 pt-4">
                <CardTitle class="text-base font-semibold mb-2">{{ course.title }}</CardTitle>
                <CardDescription class="line-clamp-3 text-sm text-muted-foreground">
                  {{ course.description || 'Tidak ada deskripsi' }}
                </CardDescription>
              </CardContent>

              <CardFooter class="p-6 pt-0 flex flex-col space-y-2">
                <Button variant="outline" size="sm" class="w-full" @click="viewDetail(course.courseId)">Lihat Detail</Button>
                <Button variant="ghost" size="sm" class="w-full flex items-center gap-1" @click="editCourse(course.courseId)">
                  <Pencil class="w-4 h-4" />
                  Edit
                </Button>
                <Button variant="destructive" size="sm" class="w-full flex items-center gap-1" @click="handleDeleteCourse(course.courseId)">
                  <Trash2 class="w-4 h-4" />
                  Hapus
                </Button>
              </CardFooter>
            </Card>
          </div>
          <p v-else class="text-muted-foreground">Anda belum membuat kursus apa pun.</p>
        </section>

        <!-- Kursus Saya Ikuti -->
        <section class="mt-10">
          <h2 class="text-xl font-semibold text-foreground mb-4">Kursus yang Anda Ikuti</h2>
          <div v-if="enrolledCourses.length" class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
            <Card
              v-for="course in enrolledCourses"
              :key="`enrolled-${course.enrollmentId}`"
              class="hover:shadow-lg transition-shadow flex flex-col h-full"
            >
              <CardHeader class="p-0 -mt-6">
                <img
                  :src="course.image || 'https://via.placeholder.com/400x160'"
                  alt="Course image"
                  class="w-full h-40 object-cover rounded-t-md"
                />
              </CardHeader>

              <CardContent class="flex-1 p-6 pt-4 flex flex-col justify-between">
                <div>
                  <CardTitle class="text-base font-semibold mb-2">{{ course.title }}</CardTitle>
                  <CardDescription class="line-clamp-3 text-sm text-muted-foreground">
                    {{ course.description || 'Tidak ada deskripsi' }}
                  </CardDescription>
                </div>

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

              <CardFooter class="p-6 pt-0 flex flex-col space-y-2">
                <Button variant="outline" size="sm" class="w-full" @click="viewDetail(course.courseId)">Lihat Detail</Button>
                <Button variant="ghost" size="sm" class="w-full text-destructive flex items-center gap-1" @click="handleUnenroll(course.enrollmentId)">
                  <Trash2 class="w-4 h-4" />
                  Batalkan Pendaftaran
                </Button>
              </CardFooter>
            </Card>
          </div>
          <div v-else class="text-center py-10">
            <p class="text-muted-foreground">Anda belum mengikuti kursus apa pun.</p>
            <Button class="mt-4" @click="router.push('/courses')">Telusuri Kursus</Button>
          </div>
        </section>
      </template>
    </main>
  </div>
</template>
