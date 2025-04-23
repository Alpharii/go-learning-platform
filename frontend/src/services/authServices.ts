import axiosInstance from './axiosInstance';

export const googleLogin = () => {
    window.location.href = `${axiosInstance.defaults.baseURL}/auth/google/login`;
};

export const handleGoogleCallback = async (token: string) => {
    try {
        const response = await axiosInstance.post('/auth/google/callback', { token });
        return response.data;
    } catch (error) {
        console.error('Error handling Google callback:', error);
        throw error;
    }
};