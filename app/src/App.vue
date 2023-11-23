<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { useUserStore } from './stores/user';
import { useRoute } from 'vue-router'
import Alerts from './components/Alerts.vue';
const route = useRoute()
const userStore = useUserStore()
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
                <RouterLink to="/">Lightban</RouterLink>

                <div class="flex items-center gap-x-8">
                    <div class="flex-none">
                        <div v-if="userStore.user" class="dropdown dropdown-end">
                            <label tabindex="0" class="btn btn-ghost btn-circle avatar">
                                <div class="w-10 rounded-full">
                                    <img src="/avatar.png" />
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

        <div class="mt-8">
            <RouterView />
        </div>
    </div>
</template>