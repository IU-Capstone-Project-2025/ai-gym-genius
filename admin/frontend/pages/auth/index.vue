<template>
  <div class="auth__page-container">
    <UCard class="auth__login-form">
      <template #header>
        Log In
      </template>

      <div class="flex flex-col gap-y-4">
        <UInput placeholder="Login"/>
        <UInput placeholder="Password"/>
      </div>

      <template #footer>
        <div class="grid grid-cols-1 justify-items-end">
          <UButton>Sign in</UButton>
        </div>
      </template>
    </UCard>
  </div>
</template>

<script setup lang="ts">

</script>

<style scoped>


.auth__page-container {
  @apply w-full h-full flex items-center justify-center
}

.auth__login-form {
  @apply w-[500px]
}

</style>