import axiosInstance from '@/services/axiosInstance'
import { googleLogin, handleGoogleCallback } from '../services/authServices'
import { defineStore } from 'pinia'

interface Profile {
  name: string
  image: string
}

interface User {
  id: number
  email: string
  profile: Profile
}

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as User | null,
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
        const res = await axiosInstance.get('/profile/me')
        const userData = res.data

        console.log(res.data)
    
        // Pastikan data sesuai dengan interface User
        this.user = {
          id: userData.id,
          email: userData.email,
          profile: {
            name: userData.profile?.name || '',
            image: userData.profile?.image || '',
          },
        }
    
        console.log('Fetched user data:', this.user) // Debugging
    
        // Redirect ke halaman pembuatan profil jika nama kosong
        if (!this.user.profile.name) {
          window.location.href = '/profile/create'
        }
      } catch (err) {
        console.error('Gagal ambil user dari token:', err)
      }
    },
    logout() {
      this.user = null
      this.token = null
      localStorage.removeItem('token')
      window.location.href = '/login'
    },
  },
})
