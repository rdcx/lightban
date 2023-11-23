<script lang="ts" setup>
import api from '@/api';
import { ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
import { useAlertStore } from '@/stores/alert';
const { t } = useI18n()

const router = useRouter()
const username = ref('')
const email = ref('')
const password = ref('')

const error = ref('')

const register = () => {
    error.value = ''
    api.auth.register(username.value, email.value, password.value)
        .then((res) => {
            const alertStore = useAlertStore();
            alertStore.setAlert('Successfully registered', 'success');

            router.push({ name: 'login' })
        })
        .catch((err) => {
            if (err.response.status == 400) {
                error.value = t(err.response.data.message)
                return
            }
            error.value = t("register_error")
        })
}
</script>
<template>
    <div class="max-w-sm mx-auto">
        <div>
            <h1 class="text-2xl font-bold">Register</h1>
            <p class="text-gray-400">Create an account</p>
        </div>
        <div v-if="error" class="mt-4 text-red-500">{{ error }}</div>
        <div class="mt-8">
            <input class="input input-bordered w-full" v-model="username" placeholder="Username" />
        </div>
        <div class="mt-4">
            <input class="input input-bordered w-full" v-model="email" placeholder="Email" />
        </div>
        <div class="mt-4">
            <input class="input input-bordered w-full" type="password" v-model="password" placeholder="Password" />
        </div>

        <div class="mt-4 flex">
            <button @click="register" class="btn btn-primary">Register</button>
            <div class="grow"></div>
            <RouterLink :to="{ name: 'login' }">Already have an account?</RouterLink>
        </div>
    </div>
</template>