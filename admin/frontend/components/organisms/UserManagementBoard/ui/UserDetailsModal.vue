<template>
  <div v-if="isOpen" class="fixed inset-0 z-50 flex items-center justify-center">
    <!-- Backdrop -->
    <div class="fixed inset-0 bg-black/50" @click="$emit('close')"></div>
    
    <!-- Modal Content -->
    <div class="relative max-w-4xl w-full m-4 max-h-[90vh] overflow-y-auto">
      <UCard class="w-full">
      <template #header>
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <UAvatar
              :src="user.avatar"
              :alt="user.name"
              size="md"
            />
            <div>
              <h3 class="text-lg font-semibold">{{ user.name }}</h3>
              <p class="text-sm text-gray-500">{{ user.email }}</p>
            </div>
          </div>
          <UButton
            variant="ghost"
            icon="i-heroicons-x-mark"
            @click="$emit('close')"
          />
        </div>
      </template>

      <div class="space-y-6">
        <!-- User Status and Quick Actions -->
        <div class="flex items-center justify-between p-4 bg-gray-50 dark:bg-gray-800/50 rounded-lg">
          <div class="flex items-center gap-4">
            <div>
              <p class="text-sm text-gray-500">Status</p>
              <UBadge
                :color="getStatusColor(user.status)"
                :variant="user.status === 'active' ? 'solid' : 'soft'"
              >
                {{ user.status }}
              </UBadge>
            </div>
            <div>
              <p class="text-sm text-gray-500">Member Since</p>
              <p class="text-sm font-medium">{{ formatDate(user.registrationDate) }}</p>
            </div>
            <div>
              <p class="text-sm text-gray-500">Last Active</p>
              <p class="text-sm font-medium">{{ formatDate(user.lastActivity) }}</p>
            </div>
          </div>
          <div class="flex items-center gap-2">
            <UButton
              :color="user.status === 'banned' ? 'green' : 'red'"
              :variant="user.status === 'banned' ? 'solid' : 'outline'"
              @click="toggleStatus"
            >
              {{ user.status === 'banned' ? 'Unban User' : 'Ban User' }}
            </UButton>
          </div>
        </div>

        <!-- Tabs -->
        <div class="border-b border-gray-200 dark:border-gray-700 mb-6">
          <nav class="flex space-x-8">
            <button
              v-for="tab in tabs"
              :key="tab.key"
              @click="activeTab = tab.key"
              :class="[
                'py-2 px-1 border-b-2 font-medium text-sm transition-colors',
                activeTab === tab.key
                  ? 'border-blue-500 text-blue-600 dark:text-blue-400'
                  : 'border-transparent text-gray-500 hover:text-gray-700 hover:border-gray-300 dark:text-gray-400 dark:hover:text-gray-300'
              ]"
            >
              {{ tab.label }}
            </button>
          </nav>
        </div>

        <!-- Tab Content -->
        <div class="tab-content">
          <!-- Overview Tab -->
          <div v-if="activeTab === 'overview'" class="space-y-6">
            <div class="space-y-6">
              <!-- Contact Information -->
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <UCard>
                  <template #header>
                    <h4 class="font-medium">Contact Information</h4>
                  </template>
                  <div class="space-y-3">
                    <div>
                      <p class="text-sm text-gray-500">Email</p>
                      <p class="font-medium">{{ user.email }}</p>
                    </div>
                    <div v-if="user.phone">
                      <p class="text-sm text-gray-500">Phone</p>
                      <p class="font-medium">{{ user.phone }}</p>
                    </div>
                    <div>
                      <p class="text-sm text-gray-500">Registration Date</p>
                      <p class="font-medium">{{ formatDate(user.registrationDate) }}</p>
                    </div>
                  </div>
                </UCard>

                <UCard>
                  <template #header>
                    <h4 class="font-medium">Quick Stats</h4>
                  </template>
                  <div class="space-y-3">
                    <div class="flex justify-between">
                      <span class="text-sm text-gray-500">Total Workouts</span>
                      <span class="font-medium">{{ user.stats.totalWorkouts }}</span>
                    </div>
                    <div class="flex justify-between">
                      <span class="text-sm text-gray-500">Current Streak</span>
                      <span class="font-medium">{{ user.stats.streak }} days</span>
                    </div>
                    <div class="flex justify-between">
                      <span class="text-sm text-gray-500">Total Time</span>
                      <span class="font-medium">{{ Math.floor(user.stats.totalTimeSpent / 60) }}h {{ user.stats.totalTimeSpent % 60 }}m</span>
                    </div>
                    <div class="flex justify-between">
                      <span class="text-sm text-gray-500">Avg Session</span>
                      <span class="font-medium">{{ user.stats.averageSessionTime }}m</span>
                    </div>
                  </div>
                </UCard>
              </div>

              <!-- Favorite Exercises -->
              <UCard>
                <template #header>
                  <h4 class="font-medium">Favorite Exercises</h4>
                </template>
                <div class="flex flex-wrap gap-2">
                  <UBadge
                    v-for="exercise in user.stats.favoriteExercises"
                    :key="exercise"
                    variant="soft"
                    color="blue"
                  >
                    {{ exercise }}
                  </UBadge>
                </div>
              </UCard>
            </div>
          </div>

          <!-- Subscription Tab -->
          <div v-if="activeTab === 'subscription'" class="space-y-6">
            <div class="space-y-6">
              <UCard>
                <template #header>
                  <div class="flex items-center justify-between">
                    <h4 class="font-medium">Subscription Details</h4>
                    <UButton
                      variant="outline"
                      size="sm"
                      @click="$emit('update-subscription', user.id, {})"
                    >
                      Edit Subscription
                    </UButton>
                  </div>
                </template>
                
                <div class="space-y-4">
                  <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div>
                      <p class="text-sm text-gray-500">Current Plan</p>
                      <UBadge
                        :color="getSubscriptionPlanColor(user.subscription.plan)"
                        size="lg"
                      >
                        {{ user.subscription.plan.toUpperCase() }}
                      </UBadge>
                    </div>
                    <div>
                      <p class="text-sm text-gray-500">Status</p>
                      <UBadge
                        :color="getSubscriptionStatusColor(user.subscription.status)"
                        variant="soft"
                      >
                        {{ user.subscription.status }}
                      </UBadge>
                    </div>
                  </div>

                  <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
                    <div>
                      <p class="text-sm text-gray-500">Start Date</p>
                      <p class="font-medium">{{ formatDate(user.subscription.startDate) }}</p>
                    </div>
                    <div>
                      <p class="text-sm text-gray-500">End Date</p>
                      <p class="font-medium">{{ formatDate(user.subscription.endDate) }}</p>
                    </div>
                  </div>

                  <div>
                    <p class="text-sm text-gray-500">Auto Renewal</p>
                    <UBadge
                      :color="user.subscription.autoRenew ? 'green' : 'gray'"
                      variant="soft"
                    >
                      {{ user.subscription.autoRenew ? 'Enabled' : 'Disabled' }}
                    </UBadge>
                  </div>
                </div>
              </UCard>
            </div>
          </div>

          <!-- Activity Tab -->
          <div v-if="activeTab === 'activity'" class="space-y-6">
            <div class="space-y-6">
              <!-- Monthly Activity Chart -->
              <UCard>
                <template #header>
                  <h4 class="font-medium">Monthly Workouts (Last 12 Months)</h4>
                </template>
                <div class="h-64">
                  <ActivityChart
                    :data="{
                      labels: getMonthLabels(),
                      values: user.stats.monthlyActivity
                    }"
                    title=""
                    label="Monthly Workouts"
                  />
                </div>
              </UCard>

              <!-- Personal Activity Summary -->
              <UCard>
                <template #header>
                  <h4 class="font-medium">{{ user.name }}'s Activity Summary</h4>
                </template>
                <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
                  <div class="text-center p-4 bg-blue-50 dark:bg-blue-900/20 rounded-lg">
                    <p class="text-2xl font-bold text-blue-600">{{ user.stats.totalWorkouts }}</p>
                    <p class="text-sm text-gray-500">Total Workouts</p>
                  </div>
                  <div class="text-center p-4 bg-green-50 dark:bg-green-900/20 rounded-lg">
                    <p class="text-2xl font-bold text-green-600">{{ user.stats.streak }}</p>
                    <p class="text-sm text-gray-500">Current Streak (Days)</p>
                  </div>
                  <div class="text-center p-4 bg-purple-50 dark:bg-purple-900/20 rounded-lg">
                    <p class="text-2xl font-bold text-purple-600">{{ user.stats.averageSessionTime }}m</p>
                    <p class="text-sm text-gray-500">Avg Session Time</p>
                  </div>
                </div>
              </UCard>
            </div>
          </div>
        </div>
      </div>
      </UCard>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import type { User, UserStatus } from '~/composables/useUserManagement'
import { ActivityChart } from '~/components/molecules/ActivityChart'

interface Props {
  user: User
  isOpen: boolean
}

interface Emits {
  (e: 'close'): void
  (e: 'update-status', userId: number, status: UserStatus): void
  (e: 'update-subscription', userId: number, subscription: Partial<User['subscription']>): void
}

const props = defineProps<Props>()
const emit = defineEmits<Emits>()

const activeTab = ref('overview')

const tabs = [
  { key: 'overview', label: 'Overview' },
  { key: 'subscription', label: 'Subscription' },
  { key: 'activity', label: 'Activity' }
]

const getStatusColor = (status: UserStatus) => {
  switch (status) {
    case 'active': return 'green'
    case 'banned': return 'red'
    case 'suspended': return 'yellow'
    default: return 'gray'
  }
}

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

const getMonthLabels = () => {
  const months = []
  const now = new Date()
  for (let i = 11; i >= 0; i--) {
    const date = new Date(now.getFullYear(), now.getMonth() - i, 1)
    months.push(date.toLocaleDateString('en-US', { month: 'short' }))
  }
  return months
}

const toggleStatus = () => {
  const newStatus: UserStatus = props.user.status === 'banned' ? 'active' : 'banned'
  emit('update-status', props.user.id, newStatus)
}
</script>