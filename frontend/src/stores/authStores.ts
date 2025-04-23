import { googleLogin, handleGoogleCallback } from '../services/authServices';
import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        user: null as any | null,
        token: null as string | null,
    }),
    actions: {
        async loginWithGoogle() {
            googleLogin();
        },
        async fetchUser(token: string | null) {
            if (!token) {
                console.error('Token is null or undefined');
                return;
            }

            try {
                const data = await handleGoogleCallback(token);
                this.user = data.user;
                this.token = data.token;
                if (this.token !== null) {
                    localStorage.setItem('token', this.token);
                }
            } catch (error) {
                console.error('Failed to fetch user:', error);
            }
        },
        logout() {
            this.user = null;
            this.token = null;
            localStorage.removeItem('token');
        },
    },
});