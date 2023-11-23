<script lang="ts" setup>
import api from '@/api';
import { onMounted } from 'vue';
import type { Project } from '@/types';
import { ref } from 'vue';
import { useRouter } from 'vue-router'

const router = useRouter()
const projects = ref<Project[]>([])
const name = ref('')

const createProject = () => {
    api.projects.create(name.value)
        .then((res) => {
            projects.value.push(res.data)
        })
        .catch((err) => {
            console.log(err)
        })
}

const openProject = (id: number) => {
    router.push({ name: 'project', params: { id } })
}

onMounted(() => {
    api.projects.all()
        .then((res) => {
            projects.value = res.data
        })
        .catch((err) => {
            console.log(err)
        })
})

</script>
<template>
    <div class="max-w-sm mx-auto mt-8">
        <div>
            <h1 class="text-2xl font-bold">Projects</h1>
            <p class="text-gray-400">Your projects</p>

            <div class="mt-4 flex">
                <input class="input input-bordered rounded-r-none w-full" v-model="name" placeholder="Name" />
                <div class="grow"></div>
                <button class="btn btn-primary rounded-l-none" @click="createProject">Create project</button>
            </div>
        </div>
        <div v-for="project in projects" class="card bordered shadow-lg cursor-pointer" @click="openProject(project.id)">
            <div class="card-body">
                <h2 class="card-title">{{ project.name }}</h2>
            </div>
        </div>
    </div>
</template>