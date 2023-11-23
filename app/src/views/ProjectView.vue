<script lang="ts" setup>
import api from '@/api';
import { onMounted } from 'vue';
import type { List, Project } from '@/types';
import { ref } from 'vue';
import { useRoute } from 'vue-router'
import { computed } from 'vue';

const route = useRoute()
const project = ref<Project>()

const listsCount = computed(() => {
    if (!project.value) return 0;
    return project.value.lists.length;
})

const form = ref({
    name: '',
    description: ''
})

const addTask = (listId: number) => {
    api.tasks.create(form.value.name, form.value.description, listId)
        .then((res) => {
            project.value?.lists.forEach((list: List) => {
                if (list.id == listId) {
                    list.tasks.push(res.data)
                }
            })
        })
        .catch((err) => {
            console.log(err)
        })
}

const dateFormat = (date: string) => {
    const d = new Date(date)

    return `${d.getDate()}/${d.getMonth() + 1}/${d.getFullYear()}` +
        ` ${d.getHours()}:${d.getMinutes()}:${d.getSeconds()}`
}

const moveLeft = (task: any) => {

    // Find the list
    const list = project.value?.lists.find((list: List) => list.id == task.list_id)


    // Find the list to the left
    const leftList = project.value?.lists.find((list: List) => list.id == task.list_id - 1)

    if (!leftList) return

    // Find the index of the task
    const index = list?.tasks.findIndex((t: any) => t.id == task.id)

    // Remove the task from the list
    list?.tasks.splice(index as number, 1)

    // Update the task
    api.tasks.update(task.id, task.name, task.description, leftList?.id)
        .then((res) => {
            project.value?.lists.forEach((list: List) => {
                if (list.id == res.data.list_id) {
                    list.tasks.push(res.data)
                }
            })
        })
        .catch((err) => {
            console.log(err)
        })
}

const moveRight = (task: any) => {

    // Find the list
    const list = project.value?.lists.find((list: List) => list.id == task.list_id)

    // Find the list to the left
    const rightList = project.value?.lists.find((list: List) => list.id == task.list_id + 1)

    if (!rightList) return

    // Find the index of the task
    const index = list?.tasks.findIndex((t: any) => t.id == task.id)

    // Remove the task from the list
    list?.tasks.splice(index as number, 1)

    // Update the task
    api.tasks.update(task.id, task.name, task.description, rightList?.id)
        .then((res) => {
            project.value?.lists.forEach((list: List) => {
                if (list.id == res.data.list_id) {
                    list.tasks.push(res.data)
                }
            })
        })
        .catch((err) => {
            console.log(err)
        })
}

onMounted(() => {
    api.projects.get(route.params.id as unknown as number)
        .then((res) => {
            project.value = res.data
        })
        .catch((err) => {
            console.log(err)
        })
})
</script>
<template>
    <div class="flex">
        <div v-for="list in project?.lists.sort((a, b) => a.id - b.id)" class="w-full bordered shadow-lg m-8">
            <div class="card-body">
                <h2 class="card-title">{{ list.name }}</h2>
            </div>

            <div class="card-body">

                <input class="input input-bordered w-full" v-model="form.name" placeholder="Task name" />

                <textarea class="textarea textarea-bordered w-full" v-model="form.description"
                    placeholder="Task description"></textarea>

                <button class="btn btn-primary" @click="addTask(list.id)">Add task</button>
            </div>

            <div class="px-8">
                <div v-for="task in list.tasks" class="my-2 p-2 bg-gray-800 rounded-xl">
                    <div class="text-2xl font-bold">{{ task.name }}</div>
                    <div class="text-secondary">{{ task.description || 'No description' }}</div>
                    <div>{{ dateFormat(task.created_at) }}</div>

                    <div class="flex">
                        <div class="grow"></div>
                        <button @click="moveLeft(task)" class="hover:underline">Left</button>
                        &lt;>
                        <button @click="moveRight(task)" class="hover:underline">Right</button>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>