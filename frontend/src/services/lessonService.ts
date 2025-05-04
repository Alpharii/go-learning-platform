import axiosInstance from './axiosInstance'
import type { Lesson } from './courseService'

/**
 * Fetch semua lesson berdasarkan course ID
 * @param courseId - ID kursus
 * @returns {Promise<Lesson[]>}
 */
export const fetchLessonsByCourse = async (courseId: number): Promise<Lesson[]> => {
  try {
    const response = await axiosInstance.get(`/lessons/${courseId}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('authToken')}`,
      },
    })
    return response.data.data.map((lesson: any) => ({
      id: lesson.ID,
      title: lesson.Title,
      content: lesson.Content,
      order: lesson.Order,
      image: lesson.Image || undefined,
      quizzes: (lesson.Quizzes || []).map((q: any) => ({
        id: q.ID,
        lessonId: q.LessonID,
        question: q.Question,
        options: q.Options,
      })),
    }))
  } catch (error) {
    console.error('Error fetching lessons:', error)
    throw error
  }
}

/**
 * Fetch detail satu lesson berdasarkan ID
 * @param lessonId - ID lesson
 * @returns {Promise<Lesson>}
 */
export const fetchLessonById = async (lessonId: number): Promise<Lesson> => {
  try {
    const response = await axiosInstance.get(`/lesson/${lessonId}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('authToken')}`,
      },
    })
    const data = response.data.data
    return {
      id: data.ID,
      title: data.Title,
      content: data.Content,
      order: data.Order,
      image: data.Image || undefined,
      quizzes: (data.Quizzes || []).map((q: any) => ({
        id: q.ID,
        lessonId: q.LessonID,
        question: q.Question,
        options: q.Options,
      })),
    }
  } catch (error) {
    console.error('Error fetching lesson:', error)
    throw error
  }
}

/**
 * Create lesson baru
 * @param formData - FormData berisi title, content, order, course_id, dan image (opsional)
 * @returns {Promise<Lesson>}
 */
export const createLesson = async (formData: FormData): Promise<Lesson> => {
  try {
    const response = await axiosInstance.post('/lessons', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
        Authorization: `Bearer ${localStorage.getItem('authToken')}`,
      },
    })
    return response.data.data
  } catch (error) {
    console.error('Error creating lesson:', error)
    throw error
  }
}

/**
 * Update lesson berdasarkan ID
 * @param lessonId - ID lesson
 * @param formData - FormData berisi title, content, order, course_id, dan image (opsional)
 * @returns {Promise<Lesson>}
 */
export const updateLesson = async (lessonId: number, formData: FormData): Promise<Lesson> => {
  try {
    const response = await axiosInstance.put(`/lesson/${lessonId}`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
        Authorization: `Bearer ${localStorage.getItem('authToken')}`,
      },
    })
    return response.data.data
  } catch (error) {
    console.error('Error updating lesson:', error)
    throw error
  }
}

/**
 * Delete lesson berdasarkan ID
 * @param lessonId - ID lesson
 * @returns {Promise<void>}
 */
export const deleteLesson = async (lessonId: number): Promise<void> => {
  try {
    await axiosInstance.delete(`/lesson/${lessonId}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('authToken')}`,
      },
    })
  } catch (error) {
    console.error('Error deleting lesson:', error)
    throw error
  }
}