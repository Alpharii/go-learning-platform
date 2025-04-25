import axiosInstance from '@/services/axiosInstance'
import { googleLogin, handleGoogleCallback } from '../services/authServices'
import { defineStore } from 'pinia'

interface User {
  id: string
  name: string
  email: string
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as any | null,
    token: null as string | null,
  }),
  persist: true,
  actions: {
    async loginWithGoogle() {
      googleLogin()
    },
    async fetchUser(token: string | null) {
      if (!token) {
        console.warn('Token kosong, tidak bisa fetch user')
        return
      }
      this.token = token
      localStorage.setItem('token', token)

      try {
        const res = await axiosInstance.get('/profile/me') // token otomatis di-inject via interceptor
        this.user = res.data
      } catch (err) {
        console.error('Gagal ambil user dari token:', err)
      }
    },
    logout() {
      this.user = null
      this.token = null
      localStorage.removeItem('token')
    },
  },
})
