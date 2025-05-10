import { defineStore } from 'pinia'

export const useCourseStore = defineStore('course', {
  state: () => ({
    currentCourseId: '' as string
  }),
  actions: {
    setCurrentCourse(id: string) {
      this.currentCourseId = id
    }
  }
})
