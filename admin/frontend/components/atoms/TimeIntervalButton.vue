<template>
  <UButton
    :color="isActive ? 'primary' : 'gray'"
    :variant="isActive ? 'solid' : 'ghost'"
    size="sm"
    @click="$emit('select')"
  >
    {{ label }}
  </UButton>
</template>

<script setup lang="ts">
interface Props {
  label: string
  isActive: boolean
}

defineProps<Props>()
defineEmits<{
  select: []
}>()
</script>