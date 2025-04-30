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

export interface Quiz {
  id: number
  lessonId: number
  question: string
  options: string[]
}

export interface Lesson {
  id: number
  title: string
  content: string
  order: number
  image?: string
  quizzes?: Quiz[]
}

/**
 * Represents detailed information about a course, including relations.
 */
export interface CourseDetail {
  id: number
  title: string
  description: string
  image?: string
  user: {
    id: number
    email: string
    profile: {
      id: number
      name: string
      image?: string
    }
  }
  lessons: Lesson[]
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

export const fetchCourseById = async (courseId: number): Promise<CourseDetail> => {
  try {
    const { data } = await axiosInstance.get(`/courses/${courseId}`, {
      headers: { Authorization: `Bearer ${localStorage.getItem('authToken')}` },
    })

    // raw API payload
    const raw = data.course

    // map to your CourseDetail interface
    const mapped: CourseDetail = {
      id: raw.ID,
      title: raw.Title,
      description: raw.Description,
      image: raw.Image || undefined,
      user: {
        id: raw.User.ID,
        email: raw.User.Email,
        profile: {
          id: raw.User.Profile?.ID ?? 0,
          name: raw.User.Profile?.Name ?? 'Unknown',
          image: raw.User.Profile?.Image,
        },
      },
      lessons: (raw.Lessons || []).map((l: any) => ({
        id: l.ID,
        title: l.Title,
        content: l.Content,
        order: l.Order,
        image: l.Image,
        quizzes: (l.Quizzes || []).map((q: any) => ({
          id: q.ID,
          lessonId: q.LessonID,
          question: q.Question,
          options: q.Options,
        })),
      })),
    }

    return mapped
  } catch (error) {
    console.error('Error fetching course detail:', error)
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