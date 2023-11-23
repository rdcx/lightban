<script lang="ts" setup>
import api from '@/api';
import { ref } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()

const router = useRouter()
const username = ref('')
const password = ref('')

const error = ref('')

const login = () => {
    error.value = ''
    api.auth.login(username.value, password.value)
        .then((res) => {
            setTimeout(() => {
                router.push({ name: 'home' })
            }, 250)
        })
        .catch((err) => {
            error.value = t("login_error")
        })
}
</script>
<template>
    <div class="max-w-sm mx-auto">
        <div>
            <h1 class="text-2xl font-bold">Login</h1>
            <p class="text-gray-400">Login to your account</p>
        </div>
        <div v-if="error" class="mt-4 text-red-500">{{ error }}</div>
        <div class="mt-8">
            <input class="input input-bordered w-full" v-model="username" placeholder="Username" />
        </div>
        <div class="mt-4">
            <input class="input input-bordered w-full" type="password" v-model="password" placeholder="Password" />
        </div>

        <div class="mt-4 flex">
            <button @click="login" class="btn btn-primary">Login</button>
            <div class="grow"></div>
            <RouterLink :to="{ name: 'register' }">Don't have an account?</RouterLink>
        </div>
    </div>
</template>