<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { Input } from '@/components/ui/input'
import { Textarea } from '@/components/ui/textarea'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Loader2, Upload } from 'lucide-vue-next'
import { createCourse } from '@/services/courseService'

const title = ref('')
const description = ref('')
const image = ref<File | null>(null)
const imagePreview = ref<string | null>(null)
const loading = ref(false)
const router = useRouter()

const handleSubmit = async () => {
  if (!title.value || !description.value) {
    alert('Title and Description are required!')
    return
  }

  loading.value = true

  const formData = new FormData()
  formData.append('title', title.value)
  formData.append('description', description.value)
  if (image.value) {
    formData.append('image', image.value)
  }

  try {
    await createCourse(formData)
    alert('Course created successfully!')
    router.push('/') // atau route ke /course-management
  } catch (error: any) {
    console.error(error)
    alert(error.response?.data?.error || 'Failed to create course.')
  } finally {
    loading.value = false
  }
}

const handleImageChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  if (target.files && target.files.length > 0) {
    const file = target.files[0]
    image.value = file
    imagePreview.value = URL.createObjectURL(file)
  }
}
</script>

<template>
  <div class="flex justify-center items-center min-h-screen bg-muted/50 p-6">
    <Card class="w-full max-w-2xl">
      <CardHeader>
        <CardTitle>Create Course</CardTitle>
      </CardHeader>

      <CardContent>
        <form @submit.prevent="handleSubmit" class="space-y-6">
          <div class="space-y-2">
            <label for="title" class="block font-medium">Title</label>
            <Input id="title" v-model="title" placeholder="Enter course title" />
          </div>

          <div class="space-y-2">
            <label for="description" class="block font-medium">Description</label>
            <Textarea id="description" v-model="description" placeholder="Enter course description" rows="5" />
          </div>

          <div class="space-y-2">
            <label for="image" class="block font-medium">Course Image</label>
            <Input id="image" type="file" accept="image/*" @change="handleImageChange" />
          </div>

          <div v-if="imagePreview" class="mt-4">
            <p class="font-medium mb-2">Image Preview:</p>
            <img :src="imagePreview" alt="Image Preview" class="w-full rounded-lg object-cover" />
          </div>

          <Button type="submit" class="w-full" :disabled="loading">
            <Loader2 v-if="loading" class="animate-spin mr-2 h-5 w-5" />
            <Upload v-else class="mr-2 h-5 w-5" />
            {{ loading ? 'Submitting...' : 'Create Course' }}
          </Button>
        </form>
      </CardContent>
    </Card>
  </div>
</template>
