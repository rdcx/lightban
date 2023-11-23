<script lang="ts" setup>
import { RouterLink, useRouter } from 'vue-router'
import { useI18n } from 'vue-i18n'
const { t } = useI18n()
const router = useRouter()
</script>
<template>
    <div class="max-w-sm mx-auto">
        <div>
            <h1 class="text-2xl font-bold">Lightweight Kanban</h1>
            <p class="text-gray-400">Stupid simple.</p>

            <button class="btn btn-primary mt-4" @click="router.push({ name: 'projects' })">Get started</button>
        </div>
    </div>
</template>