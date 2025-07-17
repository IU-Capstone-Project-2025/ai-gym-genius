<template>
  <UCard class="auth__login-form">
    <template #header>
      Log In
    </template>

    <div class="flex flex-col gap-y-4">
      <UFormGroup
          label="Login"
          :error="errors.login"
          class="space-y-1"
      >
        <UInput
            class="w-full"
            placeholder="Enter your login"
            size="xl"
            autocomplete="off"
            v-model="login"
            :disabled="isLoading"
            @blur="validateField('login')"
        />
      </UFormGroup>

      <UFormGroup
          label="Password"
          :error="errors.password"
          class="space-y-1"
      >
        <UInput
            class="w-full"
            placeholder="Enter your password"
            size="xl"
            type="password"
            autocomplete="new-password"
            v-model="password"
            :disabled="isLoading"
            @blur="validateField('password')"
            @keyup.enter="handleLogin"
        />
      </UFormGroup>

      <UAlert
          v-if="generalError"
          color="red"
          variant="solid"
          :title="generalError"
          class="mb-4"
      />
    </div>

    <template #footer>
      <div class="flex justify-end">
        <UButton
            size="xl"
            color="primary"
            :loading="isLoading"
            :disabled="isLoading"
            @click="handleLogin"
        >
          {{ isLoading ? 'Logging in...' : 'Log in' }}
        </UButton>
      </div>
    </template>
  </UCard>
</template>

<script setup lang="ts">
import {loginSchema} from '~/utils/validation/auth'
import type {ZodError} from 'zod'

const props = defineProps<{
  isLoading?: boolean
}>();

const emit = defineEmits<{
  (e: "submit", login: string, password: string): void
}>()

const login = ref("");
const password = ref("");
const errors = ref({
  login: '',
  password: ''
});
const generalError = ref('');

const validateField = (field: 'login' | 'password') => {
  try {
    const data = {login: login.value, password: password.value};
    loginSchema.parse(data);
    errors.value[field] = '';
  } catch (error) {
    if (error instanceof Error && 'errors' in error) {
      const zodError = error as ZodError;
      const fieldError = zodError.errors.find(e => e.path[0] === field);
      if (fieldError) {
        errors.value[field] = fieldError.message;
      }
    }
  }
};

const validateForm = (): boolean => {
  // Clear previous errors
  errors.value.login = '';
  errors.value.password = '';
  generalError.value = '';

  return login.value && password.value;
};

const handleLogin = () => {
  console.log('handleLogin in AuthForm called')
  generalError.value = '';

  if (!validateForm()) {
    generalError.value = "All fields must be filled";
    console.log('Validation failed');
    return;
  }

  console.log('Emitting submit event with:', login.value, password.value)
  emit('submit', login.value, password.value);
}

// Clear general error when user starts typing
watch([login, password], () => {
  generalError.value = '';
});

// Show error passed from parent
defineExpose({
  setError: (error: string) => {
    generalError.value = error;
  }
});

</script>

<style scoped>


.auth__login-form {
  @apply w-[500px]
}

</style>