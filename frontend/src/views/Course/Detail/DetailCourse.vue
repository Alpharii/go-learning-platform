<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { fetchCourseById, type CourseDetail } from '@/services/courseService'
import { Card, CardHeader, CardContent, CardTitle, CardDescription } from '@/components/ui/card'
import { Skeleton } from '@/components/ui/skeleton'
import { Alert, AlertDescription } from '@/components/ui/alert'
import { Button } from '@/components/ui/button'

const route = useRoute()
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
  <div class="min-h-screen bg-orange-50">
    <!-- LOADING STATE -->
    <div v-if="isLoading" class="max-w-7xl mx-auto py-12 px-4">
      <Skeleton class="h-8 w-1/2 mb-4" />
      <Skeleton class="h-6 w-full mb-2" />
      <Skeleton class="h-6 w-3/4" />
    </div>

    <!-- ERROR STATE -->
    <div v-else-if="error" class="max-w-2xl mx-auto py-12 px-4">
      <Alert>
        <AlertDescription>{{ error }}</AlertDescription>
      </Alert>
    </div>

    <!-- COURSE DETAIL -->
    <div v-else>
      <div class="max-w-7xl mx-auto px-4 py-12 grid grid-cols-1 md:grid-cols-2 gap-10 items-center">
        <!-- LEFT SIDE -->
        <div class="space-y-6">
          <h1 class="text-4xl font-bold leading-tight">{{ course?.title }}</h1>
          <p class="text-muted-foreground">{{ course?.description }}</p>

          <div class="flex items-center gap-6 text-sm text-muted-foreground">
            <div>⭐ <strong>{{ course?.rating || 4.7 }}</strong> Rating</div>
            <div><strong>{{ course?.lessons.length || 0 }}</strong> Modul</div>
          </div>

          <div class="flex items-center gap-4">
            <Button size="lg">Mulai Sekarang</Button>
            <div class="text-lg font-semibold text-foreground">Rp{{ new Intl.NumberFormat().format(1607000) }}</div>
          </div>

          <p class="text-sm text-muted-foreground">
            {{ course?.enrollments || "1,481,897" }} peserta telah mendaftar
          </p>

          <!-- INSTRUCTOR INFO -->
          <div class="flex items-center gap-3 mt-2">
            <img
              v-if="course?.user.profile.image"
              :src="course.user.profile.image"
              class="w-10 h-10 rounded-full object-cover"
              alt="Instructor"
            />
            <p class="text-sm text-muted-foreground">
              Dibuat oleh <span class="font-semibold">{{ course.user.profile.name }}</span>
            </p>
          </div>
        </div>

        <!-- RIGHT SIDE (IMAGE) -->
        <div class="flex justify-center md:justify-end">
          <img
            v-if="course?.image"
            :src="course.image"
            class="rounded-2xl shadow-lg w-full max-w-md object-cover"
            alt="Course"
          />
        </div>
      </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 bg-yellow-50 py-10">
      <Card
        v-for="(topic, index) in course?.lessons"
        :key="index"
        class="overflow-hidden transition-shadow hover:shadow-md ml-36"
      >
        <!-- Image -->
        <div v-if="topic.image" class="relative h-60 w-full -mb-15">
          <img
            :src="topic.image"
            alt="Lesson Image"
            class="h-full w-full object-cover -mt-10"
          />
        </div>

        <CardContent class="p-6">
          <!-- Title -->
          <h3 class="text-lg font-semibold mb-2">{{ topic.title }}</h3>

          <!-- Content Snippet -->
          <p class="text-sm text-muted-foreground line-clamp-3">
            {{ topic.content }}
          </p>

          <!-- Action Button -->
          <Button variant="link" size="sm" class="mt-4 p-0">
            Lihat konten →
          </Button>
        </CardContent>
      </Card>
    </div>
    </div>
  </div>
</template>
