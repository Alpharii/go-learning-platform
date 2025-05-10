<template>
  <div class="min-h-screen bg-background">
    <main class="max-w-4xl mx-auto py-10 px-4">
      <h2 class="text-xl font-semibold text-foreground mb-4">Manajemen Kursus Saya</h2>

      <Alert v-if="error" variant="destructive" class="mb-4">
        <AlertDescription>{{ error }}</AlertDescription>
      </Alert>

      <div v-if="isLoading" class="mt-8 grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        <div v-for="n in 3" :key="n" class="flex flex-col space-y-2">
          <Skeleton class="h-40 rounded-md" />
          <Skeleton class="h-6 w-3/4" />
          <Skeleton class="h-4 w-full" />
        </div>
      </div>

      <div v-else-if="createdCourses.length" class="grid gap-6 grid-cols-1 md:grid-cols-2 lg:grid-cols-3">
        <Card
          v-for="course in createdCourses"
          :key="course.courseId"
          class="hover:shadow-lg transition-shadow flex flex-col h-full"
        >
          <CardHeader class="p-0 -mt-6">
            <img
              :src="course.image || 'https://via.placeholder.com/400x160'"
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
            <Button variant="outline" size="sm" class="w-full" @click="viewDetail(course.courseId)">
              Lihat Detail
            </Button>
            <Button variant="ghost" size="sm" class="w-full flex items-center gap-1" @click="editCourse(course.courseId)">
              <Pencil class="w-4 h-4" /> Edit
            </Button>
            <Button variant="secondary" size="sm" class="w-full" @click="manageLessons(course.courseId)">
              Kelola Lesson
            </Button>
            <Button variant="destructive" size="sm" class="w-full flex items-center gap-1" @click="handleDelete(course.courseId)">
              <Trash2 class="w-4 h-4" /> Hapus
            </Button>
          </CardFooter>
        </Card>
      </div>

      <p v-else class="text-muted-foreground">Anda belum membuat kursus apa pun.</p>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { Button } from '@/components/ui/button'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Skeleton } from '@/components/ui/skeleton'
import { Card, CardHeader, CardContent, CardDescription, CardTitle, CardFooter } from '@/components/ui/card'
import { Pencil, Trash2 } from 'lucide-vue-next'
import { getMyProfile } from '@/services/profileServices'
import { deleteCourse } from '@/services/courseService'
import { useCourseStore } from '@/stores/course'


interface CreatedCourse {
  courseId: number
  title: string
  description: string
  image: string
}

const router = useRouter()
const createdCourses = ref<CreatedCourse[]>([])
const isLoading = ref(true)
const error = ref<string | null>(null)

const loadCreated = async () => {
  try {
    const data = await getMyProfile()
    createdCourses.value = data.created_courses.map((c: any) => ({
      courseId: c.id,
      title: c.title,
      description: c.description,
      image: c.image
    }))
    error.value = null
  } catch (err) {
    error.value = 'Gagal memuat kursus buatan.'
    console.error(err)
  } finally {
    isLoading.value = false
  }
}

onMounted(loadCreated)

const viewDetail = (id: number) => router.push(`/courses/${id}`)
const editCourse = (id: number) => router.push(`/courses/${id}/edit`)
const courseStore = useCourseStore()

const manageLessons = (id: number) => {
  console.log('manage lesson')
  courseStore.setCurrentCourse(String(id))
  console.log('succes store')
  router.push({ name: 'LessonManagement', params: { id } })
}

const handleDelete = async (id: number) => {
  if (!confirm('Yakin ingin menghapus kursus ini?')) return
  try {
    await deleteCourse(id)
    createdCourses.value = createdCourses.value.filter(c => c.courseId !== id)
  } catch (err) {
    alert('Gagal menghapus kursus.')
    console.error(err)
  }
}
</script>

<style scoped>
/* tambahkan styling jika perlu */
</style>
