import axiosInstance from './axiosInstance'

/**
 * Represents a course as returned within an enrollment.
 */
export interface Course {
  ID: number
  Title: string
  Description: string
  Image: string
  // include other fields if needed
}

/**
 * Represents an enrollment record.
 */
export interface Enrollment {
  ID: number
  UserID: number
  CourseID: number
  Progress: number
  Course: Course
}

/**
 * Enroll the current user in a course.
 * @param courseId - The ID of the course to enroll in.
 * @returns The created enrollment record.
 */
export const enrollUser = async (courseId: number): Promise<Enrollment> => {
  try {
    const response = await axiosInstance.post(
      '/enroll',
      { course_id: courseId },
      {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('authToken')}`,
        },
      }
    )
    return response.data.enrollment as Enrollment
  } catch (error) {
    console.error('Error enrolling user:', error)
    throw error
  }
}

/**
 * Fetch all enrollments for a given user.
 * @param userId - The ID of the user whose enrollments you want.
 * @returns A list of enrollment records.
 */
export const getEnrollments = async (userId: number): Promise<Enrollment[]> => {
  try {
    const response = await axiosInstance.get(
      `/enrollments/${userId}`,
      {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('authToken')}`,
        },
      }
    )
    return response.data.enrollments as Enrollment[]
  } catch (error) {
    console.error('Error fetching enrollments:', error)
    throw error
  }
}

/**
 * Cancel an existing enrollment.
 * @param enrollmentId - The ID of the enrollment to cancel.
 */
export const cancelEnrollment = async (enrollmentId: number): Promise<void> => {
  try {
    await axiosInstance.delete(
      `/enroll/${enrollmentId}`,
      {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('authToken')}`,
        },
      }
    )
  } catch (error) {
    console.error('Error canceling enrollment:', error)
    throw error
  }
}
