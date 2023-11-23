import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import type { User, Alert, AlertType } from '../types'

export const useAlertStore = defineStore('alert', () => {
    const alerts = ref<Array<Alert>>([])

    const setAlert = (message: string, type: AlertType) => {
        const alert: Alert = {
            id: Date.now().toString(),
            message,
            type,
        }
        alerts.value.push(alert)
        setTimeout(() => {
            alerts.value = alerts.value.filter((a) => a.id !== alert.id)
        }, 5000)
    }

    const dismiss = (alert: Alert) => {
        alerts.value = alerts.value.filter((a) => a.id !== alert.id)
    }

    return { alerts, setAlert, dismiss }
})