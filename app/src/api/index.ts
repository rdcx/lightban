import axios from 'axios';
import type { CodeFile } from '../types';
import { useUserStore } from '../stores/user';
import { useAlertStore } from '../stores/alert';
import router from '../router';

const api = axios.create({
    baseURL: import.meta.env.VITE_API_URL as string,
});

api.interceptors.request.use((config) => {
    const userStore = useUserStore();
    if (userStore.token) {
        config.headers.Authorization = `Bearer ${userStore.token}`;
    }

    return config;
});

api.interceptors.response.use(res => res, (res) => {
    if (res.status >= 200 && res.status < 300) {
        return res
    }
    if (res.response.status === 401) {
        const userStore = useUserStore();
        const alertStore = useAlertStore();
        userStore.logout();
        router.push('/login');
        alertStore.setAlert('You are not logged in', 'error');
    }

    return Promise.reject(res)
})

export default {
    auth: {
        login(username: String, password: String) {
            return api.post('/login', { username, password })
        },
        register(username: String, email: String, password: String) {
            return api.post('/register', { username, email, password })
        },
        user() {
            return api.get('/user')
        }
    },
}