import { ref, computed } from 'vue'
import type { Ref } from 'vue'
import { authService } from '~/services/auth.service'
import { validateLoginForm } from '~/utils/validation/auth'

export interface AuthUser {
  id: string
  email: string
  name: string
  role: 'admin' | 'super_admin'
}

export interface LoginCredentials {
  login: string
  password: string
}

export interface AuthState {
  user: Ref<AuthUser | null>
  token: Ref<string | null>
  isAuthenticated: Ref<boolean>
  isLoading: Ref<boolean>
}

export const useAuth = () => {
  const user = useCookie<AuthUser | null>('auth-user', {
    httpOnly: true,
    sameSite: 'strict',
    secure: true
  })
  
  const token = useCookie<string | null>('auth-token', {
    httpOnly: true,
    sameSite: 'strict',
    secure: true
  })

  const isLoading = ref(false)
  const isAuthenticated = computed(() => !!token.value && !!user.value)

  const login = async (credentials: LoginCredentials): Promise<void> => {
    isLoading.value = true
    console.log('Login attempt with credentials:', credentials)
    
    try {
      console.log('Calling API...')
      // Call the backend API
      const response = await authService.login(credentials)
      console.log('API response:', response)

      // Store the token
      token.value = response.token

      // If user data is provided in response, use it; otherwise, create minimal user object
      if (response.user) {
        user.value = {
          id: response.user.id,
          email: response.user.email,
          name: response.user.name,
          role: response.user.role as 'admin' | 'super_admin'
        }
      } else {
        // Create minimal user object from login credentials
        user.value = {
          id: '1', // This should ideally come from the backend
          email: credentials.login, // Using login as email for now
          name: 'Admin User',
          role: 'admin'
        }
      }

      // Navigate to dashboard after successful login
      await navigateTo('/')
    } catch (error: any) {
      console.error('Login error:', error)
      // Re-throw the error so it can be handled by the component
      throw error
    } finally {
      isLoading.value = false
    }
  }

  const logout = async (): Promise<void> => {
    isLoading.value = true
    
    try {
      // TODO: Call logout API endpoint if needed
      
      // Clear auth data
      token.value = null
      user.value = null

      // Navigate to auth page
      await navigateTo('/auth')
    } catch (error) {
      console.error('Logout error:', error)
      throw error
    } finally {
      isLoading.value = false
    }
  }

  const checkAuth = async (): Promise<boolean> => {
    // If we have a token but no user data, try to fetch user info
    if (token.value && !user.value) {
      try {
        // TODO: Replace with actual API call
        const response = await $fetch('/api/auth/me', {
          headers: {
            Authorization: `Bearer ${token.value}`
          }
        })

        // For now, mock the response
        user.value = {
          id: '1',
          email: 'admin@example.com',
          name: 'Admin User',
          role: 'admin'
        }

        return true
      } catch (error) {
        // Token is invalid, clear auth data
        token.value = null
        user.value = null
        return false
      }
    }

    return isAuthenticated.value
  }

  return {
    user: readonly(user),
    token: readonly(token),
    isAuthenticated: readonly(isAuthenticated),
    isLoading,
    login,
    logout,
    checkAuth
  }
}