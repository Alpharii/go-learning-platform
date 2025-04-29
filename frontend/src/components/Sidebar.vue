<template>
  <aside
    class="fixed top-0 left-0 h-screen bg-background border-r flex flex-col transition-all duration-300"
    :class="isCollapsed ? 'w-20 p-2' : 'w-64 p-6'"
  >
    <div class="flex items-center justify-between mb-6">
      <h1 v-if="!isCollapsed" class="text-xl font-bold">Logo</h1>
      <Button variant="ghost" size="icon" @click="toggleSidebar">
        <component :is="isCollapsed ? Menu : X" class="h-5 w-5" />
      </Button>
    </div>

    <nav class="flex-1 space-y-2">
      <RouterLink
        v-for="item in menuItemsTop"
        :key="item.path"
        :to="item.path"
        class="flex items-center gap-3 text-foreground hover:bg-muted px-3 py-2 rounded-lg transition"
        :class="{ 'bg-muted font-semibold': route.path === item.path }"
      >
        <component :is="item.icon" class="h-5 w-5 shrink-0" />
        <span v-if="!isCollapsed" class="truncate">{{ item.label }}</span>
      </RouterLink>

      <div v-if="!isCollapsed" class="pt-4 border-t my-4"></div>

      <RouterLink
        v-for="item in menuItemsBottom"
        :key="item.path"
        :to="item.path"
        class="flex items-center gap-3 text-foreground hover:bg-muted px-3 py-2 rounded-lg transition"
        :class="{ 'bg-muted font-semibold': route.path === item.path }"
      >
        <component :is="item.icon" class="h-5 w-5 shrink-0" />
        <span v-if="!isCollapsed" class="truncate">{{ item.label }}</span>
      </RouterLink>
    </nav>

    <div class="pt-6 mt-auto border-t">
      <Button
        variant="ghost"
        class="w-full flex items-center gap-3 text-destructive"
        @click="logout"
      >
        <LogOut class="h-5 w-5" />
        <span v-if="!isCollapsed">Logout</span>
      </Button>
    </div>
  </aside>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '@/stores/authStores'
import { useRouter, useRoute, RouterLink } from 'vue-router'
import {
  Home,
  BookOpen,
  User,
  PlusCircle,
  Settings,
  LogOut,
  Menu,
  X,
} from 'lucide-vue-next'
import { Button } from '@/components/ui/button'

const auth = useAuthStore()
const router = useRouter()
const route = useRoute()

const isCollapsed = ref(false)

const toggleSidebar = () => {
  isCollapsed.value = !isCollapsed.value
}

const logout = () => {
  auth.logout()
  router.push('/login')
}

const menuItemsTop = [
  { label: 'Home', icon: Home, path: '/' },
  { label: 'My Course', icon: BookOpen, path: '/my-course' },
  { label: 'Profile', icon: User, path: '/myProfile' },
]

const menuItemsBottom = [
  { label: 'Create Course', icon: PlusCircle, path: '/course/create' },
  { label: 'Course Management', icon: Settings, path: '/course/management' },
]
</script>

<style scoped>
/* Sidebar fixed position */
.fixed {
  position: fixed;
  top: 0;
  left: 0;
  height: 100vh;
  z-index: 1000;
}

/* Additional spacing for the content area */
.ml-20 {
  margin-left: 5rem;
}

.ml-64 {
  margin-left: 16rem;
}

/* Shrink adjustment for icons */
.shrink-0 {
  flex-shrink: 0;
}
</style>
