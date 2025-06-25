<template>
  <div class="organisms__navbar">
    <div class="flex-grow"></div>
    <NuxtLink to="/auth">
      <UButton>Login</UButton>
    </NuxtLink>
  </div>
</template>

<script setup lang="ts">

</script>

<style scoped>

@reference "@/assets/css/main.css";

.organisms__navbar {
  @apply w-screen p-5 flex items-center gap-x-10
}

</style>