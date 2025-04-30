<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { fetchCourseById, type CourseDetail } from '@/services/courseService'
import { Card, CardHeader, CardContent, CardTitle, CardDescription } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'

const route = useRoute()
const router = useRouter()
const courseId = Number(route.params.id)

const course = ref<CourseDetail | null>(null)
const isLoading = ref(true)
const error = ref<string | null>(null)

const loadCourse = async () => {
  try {
    course.value = await fetchCourseById(courseId)
  } catch (err) {
    error.value = 'Gagal memuat detail kursus.'
    console.error(err)
  } finally {
    isLoading.value = false
  }
}

onMounted(loadCourse)
</script>

<template>
  <div class="min-h-screen bg-background p-6 space-y-8">
    <Button variant="ghost" @click="router.back()">Kembali</Button>

    <template v-if="isLoading">
      <div class="space-y-4">
        <Skeleton class="h-8 w-1/3" />
        <Skeleton class="h-64 w-full" />
        <Skeleton class="h-4 w-3/4" />
        <Skeleton class="h-4 w-2/3" />
      </div>
    </template>

    <Alert v-else-if="error" variant="destructive">
      <AlertDescription>{{ error }}</AlertDescription>
    </Alert>

    <template v-else>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6 items-start">
        <!-- Course Info -->
        <div class="md:col-span-2 space-y-4">
          <h1 class="text-3xl font-bold">{{ course?.title }}</h1>
          <p class="text-sm text-muted-foreground">Dibuat oleh: {{ course?.user.profile.name }}</p>
          <Card>
            <CardHeader>
              <CardTitle>Deskripsi</CardTitle>
            </CardHeader>
            <CardContent>
              <CardDescription>{{ course?.description }}</CardDescription>
            </CardContent>
          </Card>

          <!-- Lessons Grid -->
          <section>
            <h2 class="text-xl font-semibold mb-4">Daftar Lesson</h2>
            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
              <template v-for="lesson in course?.lessons" :key="lesson.id">
                <Card class="h-full flex flex-col">
                  <CardHeader>
                    <CardTitle>{{ lesson.title }}</CardTitle>
                  </CardHeader>
                  <CardContent class="flex-1 flex flex-col justify-between">
                    <CardDescription class="line-clamp-3 mb-3">{{ lesson.content }}</CardDescription>
                    <Button variant="outline" size="sm" @click="router.push(`/lessons/${lesson.id}`)">Lihat Lesson</Button>
                  </CardContent>
                </Card>
              </template>
            </div>
            <p v-if="course?.lessons.length === 0" class="text-center text-muted-foreground mt-6">
              Belum ada lesson untuk kursus ini.
            </p>
          </section>
        </div>

        <!-- Sidebar: Image & Quiz -->
        <div class="space-y-6">
          <div v-if="course?.image" class="overflow-hidden rounded-lg shadow-lg">
            <img :src="course.image" alt="Course image" class="object-cover w-full h-48" />
          </div>

          <section v-if="course?.lessons.some(l => l.quizzes?.length)">
            <h2 class="text-lg font-semibold mb-3">Quiz Tersedia</h2>
            <ul class="space-y-2">
              <li v-for="lesson in course.lessons" :key="lesson.id">
                <div v-if="lesson.quizzes?.length">
                  <p class="font-medium mb-1">{{ lesson.title }}</p>
                  <ul class="ml-4 list-disc space-y-1">
                    <li v-for="quiz in lesson.quizzes" :key="quiz.id">
                      <Button variant="link" size="sm" @click="router.push(`/quizzes/${quiz.id}`)">Quiz: {{ quiz.question }}</Button>
                    </li>
                  </ul>
                </div>
              </li>
            </ul>
          </section>
        </div>
      </div>
    </template>
  </div>
</template>

