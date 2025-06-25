<template>
  <div class="auth__page-container">
    <UCard class="auth__login-form">
      <template #header>
        {{ header }}
      </template>

      <div class="flex flex-col gap-y-4">
        <UInput placeholder="Login" size="xl"/>
        <UInput placeholder="Password" size="xl"/>
      </div>

      <template #footer>
        <div class="grid grid-cols-1 justify-items-end">
          <div class="flex gap-x-2">
            <UButton
                :variant="type === 'sign_in' ? 'outline' : 'solid'"
                :color="type === 'sign_in' ? 'neutral' : 'primary'"
                size="xl"
                @click="emit('switch', 'sign_up')"
            >
              Sign up
            </UButton>

            <UButton
                size="xl"
                :variant="type === 'sign_in' ? 'solid' : 'outline'"
                :color="type === 'sign_in' ? 'primary' : 'neutral'"
                @click="emit('switch', 'sign_in')"
            >
              Sign in
            </UButton>
          </div>
        </div>
      </template>
    </UCard>
  </div>
</template>

<script setup lang="ts">

const props = defineProps<{
  type: "sign_in" | "sign_up"
}>();

const emit = defineEmits<{
  (e : "switch", value : "sign_in" | "sign_up")
}>()

const header = computed(() => {
  return props.type === "sign_in" ? "Sign In" : "Sign up"
});


</script>

<style scoped>


.auth__page-container {
  @apply w-full h-full flex items-center justify-center
}

.auth__login-form {
  @apply w-[500px]
}

</style>