<template>
  <div class="w-full h-full">
    <Line :data="chartData" :options="chartOptions" />
  </div>
</template>

<script setup lang="ts">
import { Line } from 'vue-chartjs'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
} from 'chart.js'
import { computed } from 'vue'
import { useAppConfig } from '#imports'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
)

interface Props {
  data: {
    labels: string[]
    values: number[]
  }
  title?: string
  label?: string
}

const props = defineProps<Props>()
const appConfig = useAppConfig()

const chartData = computed(() => ({
  labels: props.data.labels,
  datasets: [{
    label: props.label || 'Active Users',
    data: props.data.values,
    borderColor: 'rgb(59, 130, 246)', // blue-500
    backgroundColor: 'rgba(59, 130, 246, 0.1)',
    tension: 0.3,
    fill: true
  }]
}))

const chartOptions = computed(() => ({
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    title: {
      display: !!props.title,
      text: props.title,
      color: 'rgb(156, 163, 175)', // gray-400
      font: {
        size: 16
      }
    },
    legend: {
      display: true,
      labels: {
        color: 'rgb(156, 163, 175)', // gray-400
        font: {
          size: 12
        }
      }
    },
    tooltip: {
      backgroundColor: 'rgb(31, 41, 55)', // gray-800
      titleColor: 'rgb(243, 244, 246)', // gray-100
      bodyColor: 'rgb(209, 213, 219)', // gray-300
      borderColor: 'rgb(75, 85, 99)', // gray-600
      borderWidth: 1,
      callbacks: {
        label: function(context: any) {
          return `${context.dataset.label}: ${context.parsed.y}`
        }
      }
    }
  },
  scales: {
    x: {
      grid: {
        color: 'rgba(156, 163, 175, 0.1)', // gray-400 with opacity
        borderColor: 'rgb(75, 85, 99)' // gray-600
      },
      ticks: {
        color: 'rgb(156, 163, 175)' // gray-400
      }
    },
    y: {
      grid: {
        color: 'rgba(156, 163, 175, 0.1)', // gray-400 with opacity
        borderColor: 'rgb(75, 85, 99)' // gray-600
      },
      ticks: {
        color: 'rgb(156, 163, 175)' // gray-400
      }
    }
  }
}))
</script>