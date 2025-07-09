<template>
  <div class="organisms__navbar">
    <!-- Navigation Links -->
    <div class="flex items-center gap-6">
      <NuxtLink to="/" class="text-lg font-bold">
        Admin Panel
      </NuxtLink>
      <nav class="flex items-center gap-4">
        <NuxtLink
          to="/"
          class="nav-link"
          :class="{ 'nav-link-active': $route.name === 'index' }"
        >
          Dashboard
        </NuxtLink>
        <NuxtLink
          to="/users"
          class="nav-link"
          :class="{ 'nav-link-active': $route.name === 'users' }"
        >
          Users
        </NuxtLink>
      </nav>
    </div>
    
    <div class="flex-grow"></div>
    
    <NuxtLink to="/auth" v-if="authButtonVisible">
      <UButton>Login</UButton>
    </NuxtLink>
  </div>
</template>

<script setup lang="ts">

const authButtonVisible = computed(() => {
  return useRoute().name !== 'auth';
})

</script>

<style scoped>

@reference "@/assets/css/main.css";

.organisms__navbar {
  @apply w-screen p-5 flex items-center gap-x-10 border-b border-gray-200 dark:border-gray-700
}

.nav-link {
  @apply px-3 py-2 text-sm font-medium text-gray-600 dark:text-gray-300 hover:text-gray-900 dark:hover:text-gray-100 transition-colors duration-200
}

.nav-link-active {
  @apply text-blue-600 dark:text-blue-400
}

</style>