import axiosInstance from '@/services/axiosInstance'

export const getMyProfile = async () => {
  const response = await axiosInstance.get('/profile/me') // atau endpoint yang sesuai
  return response.data
}
