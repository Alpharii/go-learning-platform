<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStores'
import { Button } from '@/components/ui/button'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Skeleton } from '@/components/ui/skeleton'
import { Card, CardHeader, CardContent, CardDescription, CardTitle, CardFooter } from '@/components/ui/card'
import { Trash2 } from 'lucide-vue-next'
import { getMyProfile } from '@/services/profileServices'
import { cancelEnrollment } from '@/services/enrollServices'

// Interface untuk data kursus beserta enrollment ID
interface EnrolledCourse {
  enrollmentId: number
  courseId: number
  title: string
  description: string
  image: string
  progress: number
}

const router = useRouter()
const auth = useAuthStore()
const enrolledCourses = ref<EnrolledCourse[]>([])
const isLoading = ref(true)
const error = ref<string | null>(null)

const loadEnrolled = async () => {
  try {
    const data = await getMyProfile()
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
    error.value = 'Gagal memuat kursus. Silakan coba lagi.'
    console.error(err)
  } finally {
    isLoading.value = false
  }
}

onMounted(loadEnrolled)

async function handleUnenroll(enrollmentId: number) {
  if (!confirm('Yakin ingin berhenti mengikuti kursus ini?')) return
  try {
    await cancelEnrollment(enrollmentId)
    enrolledCourses.value = enrolledCourses.value.filter(c => c.enrollmentId !== enrollmentId)
  } catch (err) {
    alert('Gagal membatalkan pendaftaran kursus.')
    console.error(err)
  }
}

function viewDetail(id: number) {
  router.push(`/courses/${id}`)
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
        <div v-for="n in 3" :key="n" class="flex flex-col space-y-2">
          <Skeleton class="h-40 rounded-md" />
          <Skeleton class="h-6 w-3/4" />
          <Skeleton class="h-4 w-full" />
        </div>
      </div>

      <!-- Kursus yang Anda Ikuti -->
      <template v-else>
        <h2 class="text-xl font-semibold text-foreground mb-4">Kursus yang Anda Ikuti</h2>

        <div v-if="enrolledCourses.length" class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
          <Card
            v-for="course in enrolledCourses"
            :key="course.enrollmentId"
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
              <Button variant="outline" size="sm" class="w-full" @click="viewDetail(course.courseId)">
                Lihat Detail
              </Button>
              <Button variant="ghost" size="sm" class="w-full text-destructive flex items-center gap-1"
                      @click="handleUnenroll(course.enrollmentId)">
                <Trash2 class="w-4 h-4" /> Batalkan Pendaftaran
              </Button>
            </CardFooter>
          </Card>
        </div>

        <div v-else class="text-center py-10">
          <p class="text-muted-foreground">Anda belum mengikuti kursus apa pun.</p>
          <Button class="mt-4" @click="router.push('/courses')">Telusuri Kursus</Button>
        </div>
      </template>
    </main>
  </div>
</template>
