<template>
  <div class="flex justify-center items-start p-6">
    <Card class="w-full max-w-2xl">
      <CardHeader>
        <CardTitle>Edit Kursus</CardTitle>
      </CardHeader>

      <CardContent>
        <div v-if="isLoading" class="text-center py-6">
          <p>Memuat data...</p>
        </div>

        <div v-else-if="errorCourse" class="text-red-500 mb-4 text-center">
          {{ errorCourse }}
        </div>

        <form v-else @submit.prevent="handleCourseSubmit" class="space-y-6">
          <div class="space-y-2">
            <label for="title" class="block font-medium">Judul Kursus</label>
            <Input v-model="title" id="title" placeholder="Masukkan judul kursus" />
          </div>

          <div class="space-y-2">
            <label for="description" class="block font-medium">Deskripsi</label>
            <Textarea v-model="description" id="description" placeholder="Masukkan deskripsi kursus" rows="5" />
          </div>

          <div class="space-y-2">
            <label for="image" class="block font-medium">Gambar Kursus</label>
            <Input id="image" type="file" accept="image/*" @change="handleImageUpload" />
          </div>

          <div v-if="imagePreview" class="mt-4">
            <p class="font-medium mb-2">Pratinjau Gambar:</p>
            <img :src="imagePreview" alt="Preview Gambar" class="w-full rounded-lg object-cover" />
          </div>

          <div class="flex justify-end gap-4">
            <Button variant="outline" type="button" @click="router.back()" :disabled="isSubmittingCourse">Batal</Button>
            <Button type="submit" :disabled="isSubmittingCourse" class="flex items-center gap-2">
              <Loader2 v-if="isSubmittingCourse" class="animate-spin h-5 w-5" />
              <Upload v-else class="h-5 w-5" />
              <span>{{ isSubmittingCourse ? 'Menyimpan...' : 'Simpan Perubahan' }}</span>
            </Button>
          </div>
        </form>
      </CardContent>
    </Card>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { fetchCourseById, updateCourseWithFormData, type CourseDetail } from '@/services/courseService'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Loader2, Upload } from 'lucide-vue-next'

const route = useRoute()
const router = useRouter()
const courseId = Number(route.params.id)

const course = ref<CourseDetail | null>(null)
const isLoading = ref(true)
const errorCourse = ref<string | null>(null)
const isSubmittingCourse = ref(false)

const title = ref('')
const description = ref('')
const imageFile = ref<File | null>(null)
const imagePreview = ref<string | null>(null)

const loadCourse = async () => {
  try {
    const data = await fetchCourseById(courseId)
    course.value = data
    title.value = data.title
    description.value = data.description
    if (data.image) imagePreview.value = data.image
  } catch (err) {
    errorCourse.value = 'Gagal memuat data kursus.'
    console.error(err)
  } finally {
    isLoading.value = false
  }
}

onMounted(loadCourse)

const handleImageUpload = (e: Event) => {
  const files = (e.target as HTMLInputElement).files
  if (files?.length) {
    imageFile.value = files[0]
    imagePreview.value = URL.createObjectURL(files[0])
  }
}

const handleCourseSubmit = async () => {
  if (!title.value || !description.value) {
    alert('Judul dan Deskripsi wajib diisi!')
    return
  }
  isSubmittingCourse.value = true
  const formData = new FormData()
  formData.append('Title', title.value)
  formData.append('Description', description.value)
  if (imageFile.value) formData.append('Image', imageFile.value)
  try {
    await updateCourseWithFormData(courseId, formData)
    router.push({ name: 'CourseDetail', params: { id: courseId } })
  } catch (err) {
    errorCourse.value = 'Gagal menyimpan perubahan kursus.'
    console.error(err)
  } finally {
    isSubmittingCourse.value = false
  }
}
</script>

<style scoped>
/* tambahkan styling jika perlu */
</style>