<script setup lang="ts">
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '@/stores/authStores'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle, CardDescription } from '@/components/ui/card'
import { LogIn } from 'lucide-vue-next'

const authStore = useAuthStore()
const router = useRouter()

const loginWithGoogle = () => {
  authStore.loginWithGoogle()
}

onMounted(() => {
  const urlParams = new URLSearchParams(window.location.search)
  const token = urlParams.get('token')
  console.log('TOKEN DARI URL:', token)
  if (token) {
    authStore.fetchUser(token).then(() => {
      router.push({ name: 'Home' })
    })
  }
})
</script>

<template>
  <div
    class="flex min-h-screen items-center justify-center bg-gradient-to-br from-gray-100 to-slate-200 px-4"
  >
    <Card class="w-full max-w-md shadow-xl border-none animate-fade-in rounded-2xl">
      <CardHeader class="text-center space-y-1">
        <CardTitle class="text-3xl font-extrabold text-slate-800">Welcome Back!</CardTitle>
        <CardDescription class="text-gray-500"
          >Silakan login menggunakan akun Google Anda.</CardDescription
        >
      </CardHeader>

      <CardContent>
        <Button
          variant="default"
          size="lg"
          class="w-full gap-2 bg-blue-600 hover:bg-blue-700 text-white transition-all"
          @click="loginWithGoogle"
        >
          <LogIn class="w-5 h-5" />
          Login dengan Google
        </Button>
      </CardContent>
    </Card>
  </div>
</template>

<style scoped>
@keyframes fade-in {
  from {
    opacity: 0;
    transform: translateY(16px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
.animate-fade-in {
  animation: fade-in 0.5s ease-out;
}
</style>
