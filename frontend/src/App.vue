<template>
  <div class="flex min-h-screen">
    <Sidebar v-if="showSidebar" />
    <div class="flex-1">
      <RouterView />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import Sidebar from './components/Sidebar.vue'

const route = useRoute()

const hiddenSidebarRoutes = ['Login', 'Register', 'ForgotPassword']
const showSidebar = computed(() => !hiddenSidebarRoutes.includes(route.name as string))
</script>