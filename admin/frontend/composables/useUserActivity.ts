import {ref, computed} from 'vue'
import type {TimeInterval} from '~/components/molecules/TimeIntervalSelector.vue'
import {useUsersStats} from "~/composables/useUsersStats";

const {fetchUsersStats, usersStats} = useUsersStats();

interface ActivityData {
    labels: string[]
    values: number[]
}

export const useUserActivity = () => {
    const selectedInterval = ref<TimeInterval>('24h')
    const isLoading = ref(false)

    // Mock data generator
    const generateMockData = (interval: TimeInterval): ActivityData => {
        switch (interval) {
            case '24h':
                return {
                    labels: Array.from({length: 24}, (_, i) =>
                        `${i.toString().padStart(2, '0')}:00`
                    ),
                    values: Array.from({length: 24}, () =>
                        Math.floor(Math.random() * 100) + 20
                    )
                }
            case '7d':
                const days = ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
                return {
                    labels: days,
                    values: Array.from({length: 7}, () =>
                        Math.floor(Math.random() * 500) + 100
                    )
                }
            case '30d':
                return {
                    labels: Array.from({length: 30}, (_, i) =>
                        `Day ${i + 1}`
                    ),
                    values: Array.from({length: 30}, () =>
                        Math.floor(Math.random() * 500) + 100
                    )
                }
            case '90d':
                return {
                    labels: Array.from({length: 12}, (_, i) =>
                        `Week ${i + 1}`
                    ),
                    values: Array.from({length: 12}, () =>
                        Math.floor(Math.random() * 2000) + 500
                    )
                }
        }
    }

    const activityData = computed(() => generateMockData(selectedInterval.value))

    const fetchActivityData = async () => {
        isLoading.value = true
        // Simulate API call delay
        // await new Promise(resolve => setTimeout(resolve, 500));
        await fetchUsersStats();
        isLoading.value = false
    }

    return {
        selectedInterval,
        activityData,
        isLoading,
        fetchActivityData
    }
}