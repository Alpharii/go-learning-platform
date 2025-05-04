<template>
  <div class="min-h-screen bg-background">
    <main class="max-w-4xl mx-auto py-10 px-4">
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-xl font-semibold text-foreground">Manajemen Lesson</h2>
        <Button variant="secondary" size="sm" class="flex items-center gap-1" @click="addLesson">
          <PlusCircle class="w-4 h-4" /> Tambah Lesson
        </Button>
      </div>

      <Alert v-if="error" variant="destructive" class="mb-4">
        <AlertDescription>{{ error }}</AlertDescription>
      </Alert>

      <div v-if="isLoading" class="mt-8 space-y-4">
        <Skeleton class="h-6 w-1/3" />
        <Skeleton v-for="n in 3" :key="n" class="h-12 rounded-md" />
      </div>

      <div v-else-if="lessons.length">
        <ul class="space-y-4">
          <li v-for="lesson in lessons" :key="lesson.id" class="flex items-center justify-between p-4 bg-white rounded-lg shadow">
            <div>
              <p class="font-semibold">{{ lesson.title }}</p>
              <p class="text-sm text-muted-foreground line-clamp-2">{{ lesson.content }}</p>
            </div>
            <div class="flex items-center gap-2">
              <Button variant="ghost" size="sm" @click="editLesson(lesson.id)">
                <Pencil class="w-4 h-4" />
              </Button>
              <Button variant="destructive" size="sm" @click="removeLesson(lesson.id)">
                <Trash2 class="w-4 h-4" />
              </Button>
            </div>
          </li>
        </ul>
      </div>

      <p v-else class="text-muted-foreground">Belum ada lesson untuk kursus ini.</p>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { Button } from '@/components/ui/button'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Skeleton } from '@/components/ui/skeleton'
import { PlusCircle, Pencil, Trash2 } from 'lucide-vue-next'
import { fetchLessonsByCourse, deleteLesson } from '@/services/lessonService'

type Lesson = {
  id: number
  title: string
  content: string
  order: number
}

const route = useRoute()
const router = useRouter()
const courseId = Number(route.params.id)

const lessons = ref<Lesson[]>([])
const isLoading = ref(true)
const error = ref<string | null>(null)

const loadLessons = async () => {
  try {
    lessons.value = await fetchLessonsByCourse(courseId)
    error.value = null
  } catch (err) {
    error.value = 'Gagal memuat daftar lesson.'
    console.error(err)
  } finally {
    isLoading.value = false
  }
}

onMounted(loadLessons)

const addLesson = () => router.push({ name: 'AddLesson', params: { id: courseId } })
const editLesson = (lessonId: number) => router.push({ name: 'EditLesson', params: { courseId, lessonId } })

const removeLesson = async (lessonId: number) => {
  if (!confirm('Yakin ingin menghapus lesson ini?')) return
  try {
    await deleteLesson(lessonId)
    lessons.value = lessons.value.filter(l => l.id !== lessonId)
  } catch (err) {
    alert('Gagal menghapus lesson.')
    console.error(err)
  }
}
</script>
