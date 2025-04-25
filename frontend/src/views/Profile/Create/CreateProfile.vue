<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStores'
import axiosInstance from '@/services/axiosInstance'

const auth = useAuthStore()
const router = useRouter()

const profileForm = ref({
  name: '',
  image: '',
})

const handleSubmit = async () => {
  try {
    const res = await axiosInstance.put('/profile/update', profileForm.value, {
      headers: {
        Authorization: `Bearer ${auth.token}`,
      },
    })

    console.log('Profile updated:', res.data)

    // Update state di Pinia
    await auth.fetchUser(auth.token)

    // Redirect ke dashboard setelah berhasil membuat profil
    router.push('/')
  } catch (error) {
    console.error('Error updating profile:', error)
  }
}
</script>

<template>
  <div class="min-h-screen bg-gray-50 py-10 px-4">
    <div class="max-w-md mx-auto space-y-6">
      <h1 class="text-2xl font-bold text-slate-800">Lengkapi Profil Anda</h1>
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
          <label for="image" class="block text-sm font-medium text-slate-700">URL Gambar Profil</label>
          <input
            v-model="profileForm.image"
            type="url"
            id="image"
            placeholder="Masukkan URL gambar"
            class="mt-1 block w-full px-3 py-2 border border-slate-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
          />
        </div>
        <button
          type="submit"
          class="w-full flex justify-center py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Simpan Profil
        </button>
      </form>
    </div>
  </div>
</template>