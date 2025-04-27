<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStores'
import axiosInstance from '@/services/axiosInstance'

const auth = useAuthStore()
const router = useRouter()

const profileForm = ref({
  name: auth.user?.profile?.name || '',
  image: null as File | null, // ganti dari string jadi File
})

const errorMessages = ref<string[]>([])
const isLoading = ref(false)

const handleSubmit = async () => {
  errorMessages.value = []
  isLoading.value = true

  try {
    if (!profileForm.value.name.trim()) {
      errorMessages.value.push('Nama lengkap tidak boleh kosong.')
      return
    }

    if (!profileForm.value.image) {
      errorMessages.value.push('Gambar profil wajib diupload.')
      return
    }

    const formData = new FormData()
    formData.append('name', profileForm.value.name)
    formData.append('image', profileForm.value.image)

    console.log('sending',formData)

    const res = await axiosInstance.put('/profile/update', formData, {
      headers: {
        Authorization: `Bearer ${auth.token}`,
        'Content-Type': 'multipart/form-data',
      },
    })

    console.log('Profile updated:', res.data)
    await auth.fetchUser(auth.token)
    router.push('/')
  } catch (error: any) {
    console.error('Error updating profile:', error)
    const message = error.response?.data?.message || 'Terjadi kesalahan saat memperbarui profil.'
    errorMessages.value.push(message)
  } finally {
    isLoading.value = false
  }
}

// Untuk handle perubahan file input
const handleFileChange = (event: Event) => {
  const target = event.target as HTMLInputElement
  const file = target.files?.[0] || null
  if (file && (file.type === 'image/jpeg' || file.type === 'image/png')) {
    profileForm.value.image = file
  } else {
    profileForm.value.image = null
    errorMessages.value.push('File harus berupa gambar JPG atau PNG.')
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-50 py-10 px-4">
    <div class="max-w-md mx-auto space-y-6">
      <h1 class="text-2xl font-bold text-slate-800">Perbarui Profil Anda</h1>

      <div v-if="errorMessages.length" class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
        <ul class="list-disc pl-5">
          <li v-for="(error, index) in errorMessages" :key="index">{{ error }}</li>
        </ul>
      </div>

      <form @submit.prevent="handleSubmit" class="space-y-4">
        <div>
          <label for="name" class="block text-sm font-medium text-slate-700">Nama Lengkap</label>
          <input
            v-model="profileForm.name"
            type="text"
            id="name"
            placeholder="Masukkan nama lengkap"
            required
            class="mt-1 block w-full px-3 py-2 border border-slate-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>

        <div>
          <label for="image" class="block text-sm font-medium text-slate-700">Upload Gambar Profil</label>
          <input
            type="file"
            id="image"
            accept="image/jpeg, image/png, image/jpg"
            @change="handleFileChange"
            class="mt-1 block w-full text-sm text-slate-700 file:mr-4 file:py-2 file:px-4 file:rounded-md file:border-0 file:text-sm file:font-semibold file:bg-indigo-50 file:text-indigo-700 hover:file:bg-indigo-100"
          />
        </div>

        <button
          type="submit"
          :disabled="isLoading"
          class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 disabled:bg-indigo-400"
        >
          <span v-if="isLoading">Memperbarui...</span>
          <span v-else>Simpan Perubahan</span>
        </button>
      </form>
    </div>
  </div>
</template>
