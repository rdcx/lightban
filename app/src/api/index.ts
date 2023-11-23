import axios from 'axios';
import type { CodeFile } from '../types';
import { useUserStore } from '../stores/user';
import { useAlertStore } from '../stores/alert';
import router from '../router';

const api = axios.create({
    baseURL: 'http://localhost:8080',
});

api.interceptors.request.use((config) => {
    const userStore = useUserStore();
    if (userStore.token) {
        config.headers.Authorization = `Bearer ${userStore.token}`;
    }

    return config;
});

api.interceptors.response.use(res => res, (err) => {
    if (err.response.status === 401) {
        const userStore = useUserStore();
        const alertStore = useAlertStore();
        userStore.logout();
        router.push('/login');
        alertStore.setAlert('You are not logged in', 'error');
    }

    return err;
})

export default {
    auth: {
        login(username: String, password: String) {
            return api.post('/login', { username, password })
                .then(res => {

                    if (res.status == 200, res.data.token) {
                        const userStore = useUserStore();
                        const alertStore = useAlertStore();
                        alertStore.setAlert('Successfully logged in', 'info');
                        userStore.setToken(res.data.token);
                        this.user();
                    }

                    return res;
                })
                .catch(err => {
                    const alertStore = useAlertStore();
                    alertStore.setAlert(err.response.data.message, 'error');

                    return err;
                })
        },
        register(username: String, email: String, password: String) {
            return api.post('/register', { username, email, password })
                .then(res => {
                    const alertStore = useAlertStore();
                    alertStore.setAlert('Successfully registered', 'success');

                    return res;
                })
                .catch(err => {
                    const alertStore = useAlertStore();
                    alertStore.setAlert(err.response.data.message, 'error');

                    return err;
                })
        },
        user() {
            return api.get('/user')
                .then(res => {
                    const userStore = useUserStore();
                    userStore.setUser(res.data);
                })
        }
    },
}