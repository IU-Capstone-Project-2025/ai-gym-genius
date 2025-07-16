<template>
  <div class="h-full flex items-center justify-center">
    <AuthForm
        :type="type"
        @switch="value => type = value"
        @action="(login, password) => action(login, password)"
    />
  </div>
</template>
<script setup lang="ts">
import { AuthForm } from '~/components/organisms/AuthForm'

// Disable auth middleware for this page and use auth layout
definePageMeta({
  middleware: [],
  layout: 'auth'
})

const { login: authLogin, isLoading } = useAuth()
const type: Ref<"sign_in" | "sign_up"> = ref("sign_in")

const action = async (email: string, password: string) => {
  if (type.value === 'sign_in') {
    try {
      await authLogin({ email, password })
    } catch (error) {
      console.error('Login failed:', error)
      // TODO: Show error message to user
    }
  } else {
    // TODO: Implement sign up functionality
    console.warn('Sign up not implemented yet')
  }
}
</script>