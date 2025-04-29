import axiosInstance from './axiosInstance'

export interface Profile {
  ID: number
  Name: string
  Image: string
}

export interface User {
  ID: number
  Email: string
  Profile: Profile
}

export interface Course {
  ID: number
  CreatedAt: string
  UpdatedAt: string
  DeletedAt: string | null
  Title: string
  Description: string
  Image: string
  User: User
  Progress: number
}

/**
 * Fetch all courses from the backend API.
 * @returns {Promise<Course[]>} - A list of courses.
 */
export const fetchCourses = async (): Promise<Course[]> => {
  try {
    const response = await axiosInstance.get('/courses', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('authToken')}`, // Assuming token is stored in localStorage
      },
    })
    return response.data.courses
  } catch (error) {
    console.error('Error fetching courses:', error)
    throw error
  }
}

/**
 * Create a new course.
 * @param {FormData} formData - The course data with text and image.
 * @returns {Promise<Course>} - The newly created course.
 */
export const createCourse = async (formData: FormData): Promise<Course> => {
  try {
    const response = await axiosInstance.post('/courses', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
        Authorization: `Bearer ${localStorage.getItem('authToken')}`,
      },
    })
    return response.data.course
  } catch (error) {
    console.error('Error creating course:', error)
    throw error
  }
}

/**
 * Update an existing course.
 * @param {number} courseId - The ID of the course to update.
 * @param {Partial<Course>} courseData - The updated course data.
 * @returns {Promise<Course>} - The updated course.
 */
export const updateCourse = async (courseId: number, courseData: Partial<Course>): Promise<Course> => {
  try {
    const response = await axiosInstance.put(`/courses/${courseId}`, courseData, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('authToken')}`,
      },
    })
    return response.data
  } catch (error) {
    console.error('Error updating course:', error)
    throw error
  }
}

/**
 * Delete a course by its ID.
 * @param {number} courseId - The ID of the course to delete.
 * @returns {Promise<void>}
 */
export const deleteCourse = async (courseId: number): Promise<void> => {
  try {
    await axiosInstance.delete(`/courses/${courseId}`, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('authToken')}`,
      },
    })
  } catch (error) {
    console.error('Error deleting course:', error)
    throw error
  }
}