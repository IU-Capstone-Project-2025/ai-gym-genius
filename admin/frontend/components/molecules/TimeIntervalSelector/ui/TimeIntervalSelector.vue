<template>
  <UButtonGroup size="sm" orientation="horizontal">
    <TimeIntervalButton
        v-for="interval in intervals"
        :key="interval.value"
        :label="interval.label"
        :is-active="selectedInterval === interval.value"
        @select="$emit('update:selectedInterval', interval.value)"
    />
  </UButtonGroup>
</template>

<script setup lang="ts">

import {TimeIntervalButton} from "~/components/atoms/TimeIntervalButton";

export type TimeInterval = '24h' | '7d' | '30d' | '90d'

interface TimeIntervalOption {
  label: string
  value: TimeInterval
}

interface Props {
  selectedInterval: TimeInterval
}

defineProps<Props>()
defineEmits<{
  'update:selectedInterval': [value: TimeInterval]
}>()

const intervals: TimeIntervalOption[] = [
  {label: '24 Hours', value: '24h'},
  {label: '7 Days', value: '7d'},
  {label: '30 Days', value: '30d'},
  {label: '90 Days', value: '90d'}
]
</script>