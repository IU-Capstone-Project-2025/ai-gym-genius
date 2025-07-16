<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center">
    <!-- Backdrop -->
    <div class="fixed inset-0 bg-black/50" @click="$emit('close')"></div>
    
    <!-- Modal Content -->
    <div class="relative max-w-2xl w-full m-4 max-h-[90vh] overflow-y-auto">
      <UCard class="w-full">
      <template #header>
        <div class="flex items-center justify-between">
          <h3 class="text-lg font-semibold">Edit Subscription</h3>
          <UButton
            variant="ghost"
            icon="i-heroicons-x-mark"
            @click="$emit('close')"
          />
        </div>
      </template>

      <form @submit.prevent="handleSubmit" class="space-y-6">
        <!-- User Info -->
        <div class="flex items-center gap-3 p-4 bg-gray-50 dark:bg-gray-800/50 rounded-lg">
          <UAvatar
            :src="user.avatar"
            :alt="user.name"
            size="sm"
          />
          <div>
            <p class="font-medium">{{ user.name }}</p>
            <p class="text-sm text-gray-500">{{ user.email }}</p>
          </div>
        </div>

        <!-- Current Subscription -->
        <div class="p-4 border border-gray-200 dark:border-gray-700 rounded-lg">
          <h4 class="font-medium mb-3">Current Subscription</h4>
          <div class="grid grid-cols-2 gap-4 text-sm">
            <div>
              <p class="text-gray-500">Plan</p>
              <UBadge :color="getSubscriptionPlanColor(user.subscription.plan)">
                {{ user.subscription.plan.toUpperCase() }}
              </UBadge>
            </div>
            <div>
              <p class="text-gray-500">Status</p>
              <UBadge
                :color="getSubscriptionStatusColor(user.subscription.status)"
                variant="soft"
              >
                {{ user.subscription.status }}
              </UBadge>
            </div>
            <div>
              <p class="text-gray-500">Start Date</p>
              <p class="font-medium">{{ formatDate(user.subscription.startDate) }}</p>
            </div>
            <div>
              <p class="text-gray-500">End Date</p>
              <p class="font-medium">{{ formatDate(user.subscription.endDate) }}</p>
            </div>
          </div>
        </div>

        <!-- Edit Form -->
        <div class="space-y-4">
          <!-- Subscription Plan -->
          <div>
            <label class="block text-sm font-medium mb-2">Subscription Plan</label>
            <USelect
              v-model="formData.plan"
              :options="[
                { label: 'Free', value: 'free' },
                { label: 'Basic', value: 'basic' },
                { label: 'Premium', value: 'premium' },
                { label: 'Pro', value: 'pro' }
              ]"
              placeholder="Select plan"
            />
          </div>

          <!-- Subscription Status -->
          <div>
            <label class="block text-sm font-medium mb-2">Status</label>
            <USelect
              v-model="formData.status"
              :options="[
                { label: 'Active', value: 'active' },
                { label: 'Expired', value: 'expired' },
                { label: 'Cancelled', value: 'cancelled' },
                { label: 'Pending', value: 'pending' }
              ]"
              placeholder="Select status"
            />
          </div>

          <!-- Start Date -->
          <div>
            <label class="block text-sm font-medium mb-2">Start Date</label>
            <UInput
              v-model="formData.startDate"
              type="date"
              placeholder="Select start date"
            />
          </div>

          <!-- End Date -->
          <div>
            <label class="block text-sm font-medium mb-2">End Date</label>
            <UInput
              v-model="formData.endDate"
              type="date"
              placeholder="Select end date"
            />
          </div>

          <!-- Auto Renewal -->
          <div>
            <label class="block text-sm font-medium mb-2">Auto Renewal</label>
            <UToggle
              v-model="formData.autoRenew"
              :ui="{ ring: 'ring-2 ring-primary-500' }"
            />
            <p class="text-sm text-gray-500 mt-1">
              {{ formData.autoRenew ? 'Subscription will auto-renew' : 'Subscription will not auto-renew' }}
            </p>
          </div>
        </div>

        <!-- Plan Comparison -->
        <div class="p-4 bg-blue-50 dark:bg-blue-900/20 rounded-lg">
          <h4 class="font-medium mb-3">Plan Features</h4>
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 text-sm">
            <div>
              <p class="font-medium">{{ getPlanFeatures(formData.plan).name }}</p>
              <ul class="mt-2 space-y-1 text-gray-600 dark:text-gray-400">
                <li v-for="feature in getPlanFeatures(formData.plan).features" :key="feature">
                  • {{ feature }}
                </li>
              </ul>
            </div>
            <div>
              <p class="font-medium">Price</p>
              <p class="text-lg font-bold text-blue-600">{{ getPlanFeatures(formData.plan).price }}</p>
            </div>
          </div>
        </div>

        <!-- Actions -->
        <div class="flex justify-end gap-3">
          <UButton
            variant="outline"
            @click="$emit('close')"
          >
            Cancel
          </UButton>
          <UButton
            type="submit"
            :loading="isLoading"
          >
            Update Subscription
          </UButton>
        </div>
      </form>
      </UCard>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, watch } from 'vue'
import type { User, SubscriptionPlan, SubscriptionStatus } from '~/composables/useUserManagement'

interface Props {
  user: User
  isOpen: boolean
}

interface Emits {
  (e: 'close'): void
  (e: 'update', userId: number, subscription: Partial<User['subscription']>): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const isLoading = ref(false)

const formData = reactive({
  plan: props.user.subscription.plan,
  status: props.user.subscription.status,
  startDate: props.user.subscription.startDate.toISOString().split('T')[0],
  endDate: props.user.subscription.endDate.toISOString().split('T')[0],
  autoRenew: props.user.subscription.autoRenew
})

// Reset form when user changes
watch(() => props.user, (newUser) => {
  formData.plan = newUser.subscription.plan
  formData.status = newUser.subscription.status
  formData.startDate = newUser.subscription.startDate.toISOString().split('T')[0]
  formData.endDate = newUser.subscription.endDate.toISOString().split('T')[0]
  formData.autoRenew = newUser.subscription.autoRenew
}, { immediate: true })

const getSubscriptionPlanColor = (plan: string) => {
  switch (plan) {
    case 'free': return 'gray'
    case 'basic': return 'blue'
    case 'premium': return 'purple'
    case 'pro': return 'yellow'
    default: return 'gray'
  }
}

const getSubscriptionStatusColor = (status: string) => {
  switch (status) {
    case 'active': return 'green'
    case 'expired': return 'red'
    case 'cancelled': return 'gray'
    case 'pending': return 'yellow'
    default: return 'gray'
  }
}

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(date)
}

const getPlanFeatures = (plan: SubscriptionPlan) => {
  switch (plan) {
    case 'free':
      return {
        name: 'Free Plan',
        price: '$0/month',
        features: ['Basic workout tracking', 'Limited exercise library', 'Basic stats']
      }
    case 'basic':
      return {
        name: 'Basic Plan',
        price: '$9.99/month',
        features: ['Full workout tracking', 'Complete exercise library', 'Advanced stats', 'Email support']
      }
    case 'premium':
      return {
        name: 'Premium Plan',
        price: '$19.99/month',
        features: ['All Basic features', 'Custom workout plans', 'Nutrition tracking', 'Priority support', 'Progress photos']
      }
    case 'pro':
      return {
        name: 'Pro Plan',
        price: '$29.99/month',
        features: ['All Premium features', 'Personal trainer access', 'Advanced analytics', 'API access', 'White-label options']
      }
    default:
      return {
        name: 'Unknown Plan',
        price: '$0/month',
        features: []
      }
  }
}

const handleSubmit = async () => {
  isLoading.value = true
  
  try {
    const subscriptionUpdate = {
      plan: formData.plan as SubscriptionPlan,
      status: formData.status as SubscriptionStatus,
      startDate: new Date(formData.startDate),
      endDate: new Date(formData.endDate),
      autoRenew: formData.autoRenew
    }
    
    await emit('update', props.user.id, subscriptionUpdate)
    emit('close')
  } catch (error) {
    console.error('Error updating subscription:', error)
  } finally {
    isLoading.value = false
  }
}
</script>