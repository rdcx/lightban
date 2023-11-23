<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useUserStore } from './stores/user';
import { useContextStore } from './stores/context';
import { useRoute } from 'vue-router'
import Alerts from './components/Alerts.vue';
import { computed } from 'vue';
const route = useRoute()
const userStore = useUserStore()
const contextStore = useContextStore()
const logout = () => {
    userStore.logout()
    window.location.href = '/'
}

</script>

<template>
    <div class="flex min-h-full flex-col">
        <Alerts />
        <header class="shrink-0 shadow-md">
            <div class="mx-auto flex h-16 max-w-[90rem] items-center justify-between px-4 sm:px-6 lg:px-8">
                <RouterLink to="/"><img alt="Vue logo" src="/logo/default-monochrome.svg" class="h-6" /></RouterLink>

                <div class="flex items-center gap-x-8">
                    <div class="flex-none">
                        <div v-if="userStore.user" class="dropdown dropdown-end">
                            <label tabindex="0" class="btn btn-ghost btn-circle avatar">
                                <div class="w-10 rounded-full">
                                    <img src="/avatars/unisex-avatar.png" />
                                </div>
                            </label>
                            <ul tabindex="0"
                                class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52">
                                <li>
                                    <a class="justify-between">
                                        Profile
                                        <span class="badge">New</span>
                                    </a>
                                </li>
                                <li><a>Settings</a></li>
                                <li @click="logout"><a>Logout</a></li>
                            </ul>
                        </div>

                        <div v-else>
                            <RouterLink to="/login" class="btn btn-sm"
                                :class="route.name == 'login' ? 'btn-primary' : 'btn-ghost'">Login</RouterLink>
                            <RouterLink to="/register" class="ml-2 btn btn-sm"
                                :class="route.name == 'register' ? 'btn-primary' : 'btn-ghost'">Register</RouterLink>
                        </div>
                    </div>
                </div>
            </div>
        </header>

        <div class="mx-auto flex w-full max-w-[90rem] items-start gap-x-8 px-4 py-10 sm:px-6 lg:px-8">
            <RouterView />
        </div>
    </div>
</template>