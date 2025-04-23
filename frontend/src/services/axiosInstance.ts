import axios from 'axios';

const API_URL = 'http://localhost:8080'; // Sesuaikan dengan base URL backend Anda

// Buat instance Axios dengan konfigurasi global
const axiosInstance = axios.create({
    baseURL: API_URL,
    timeout: 5000, // Timeout dalam milidetik
    headers: {
        'Content-Type': 'application/json',
    },
});

// Request interceptor (opsional)
axiosInstance.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token'); // Ambil token dari localStorage
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);

// Response interceptor (opsional)
axiosInstance.interceptors.response.use(
    (response) => {
        return response;
    },
    (error) => {
        if (error.response && error.response.status === 401) {
            console.warn('Unauthorized: Redirecting to login...');
            // Tambahkan logika redirect ke halaman login jika token tidak valid
        }
        return Promise.reject(error);
    }
);

export default axiosInstance;