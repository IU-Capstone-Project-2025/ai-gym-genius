<template>
  <div class="space-y-6">
    <!-- Header -->
    <div class="flex items-center justify-between">
      <h1 class="text-2xl font-bold">User Management</h1>
      <div class="flex items-center gap-3">
        <UBadge color="blue" variant="soft">
          {{ totalUsers }} users
        </UBadge>
        <UButton
          icon="i-heroicons-funnel"
          variant="outline"
          @click="showFilters = !showFilters"
        >
          Filters
        </UButton>
      </div>
    </div>

    <!-- Filters -->
    <UCard v-if="showFilters" class="p-4">
      <div class="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 gap-4">
        <!-- Search -->
        <UInput
          v-model="filters.search"
          placeholder="Search users..."
          icon="i-heroicons-magnifying-glass"
        />

        <!-- Status Filter -->
        <select
          v-model="filters.status"
          class="px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          <option value="all">All Statuses</option>
          <option value="active">Active</option>
          <option value="banned">Banned</option>
          <option value="suspended">Suspended</option>
        </select>

        <!-- Subscription Plan Filter -->
        <select
          v-model="filters.subscriptionPlan"
          class="px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          <option value="all">All Plans</option>
          <option value="free">Free</option>
          <option value="basic">Basic</option>
          <option value="premium">Premium</option>
          <option value="pro">Pro</option>
        </select>

        <!-- Subscription Status Filter -->
        <select
          v-model="filters.subscriptionStatus"
          class="px-3 py-2 border border-gray-300 dark:border-gray-600 rounded-md bg-white dark:bg-gray-800 text-sm focus:outline-none focus:ring-2 focus:ring-blue-500"
        >
          <option value="all">All Subscription States</option>
          <option value="active">Active</option>
          <option value="expired">Expired</option>
          <option value="cancelled">Cancelled</option>
          <option value="pending">Pending</option>
        </select>

        <!-- Clear Filters -->
        <UButton
          variant="outline"
          color="gray"
          @click="clearFilters"
          icon="i-heroicons-x-mark"
        >
          Clear
        </UButton>
      </div>
    </UCard>

    <!-- Users Table -->
    <UCard>
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-gray-200 dark:border-gray-700">
              <th class="text-left p-4 font-medium">User</th>
              <th class="text-left p-4 font-medium">Status</th>
              <th class="text-left p-4 font-medium">Subscription</th>
              <th class="text-left p-4 font-medium">Last Activity</th>
              <th class="text-left p-4 font-medium">Workouts</th>
              <th class="text-left p-4 font-medium">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr v-if="isLoading" class="border-b border-gray-200 dark:border-gray-700">
              <td colspan="6" class="p-8 text-center">
                <UIcon name="i-heroicons-arrow-path" class="animate-spin h-6 w-6 mx-auto text-primary-500" />
                <p class="mt-2 text-sm text-gray-500">Loading users...</p>
              </td>
            </tr>
            <tr
              v-else
              v-for="user in users"
              :key="user.id"
              class="border-b border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-800/50 transition-colors"
            >
              <!-- User Info -->
              <td class="p-4">
                <div class="flex items-center gap-3">
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
              </td>

              <!-- Status -->
              <td class="p-4">
                <UBadge
                  :color="getStatusColor(user.status)"
                  :variant="user.status === 'active' ? 'solid' : 'soft'"
                >
                  {{ user.status }}
                </UBadge>
              </td>

              <!-- Subscription -->
              <td class="p-4">
                <div class="space-y-1">
                  <UBadge
                    :color="getSubscriptionPlanColor(user.subscription.plan)"
                    size="sm"
                  >
                    {{ user.subscription.plan.toUpperCase() }}
                  </UBadge>
                  <p class="text-sm text-gray-500">
                    {{ user.subscription.status }}
                  </p>
                </div>
              </td>

              <!-- Last Activity -->
              <td class="p-4">
                <p class="text-sm">{{ formatDate(user.lastActivity) }}</p>
              </td>

              <!-- Workouts -->
              <td class="p-4">
                <div class="space-y-1">
                  <p class="text-sm font-medium">{{ user.stats.totalWorkouts }}</p>
                  <p class="text-xs text-gray-500">{{ user.stats.streak }} day streak</p>
                </div>
              </td>

              <!-- Actions -->
              <td class="p-4">
                <div class="flex items-center gap-2">
                  <UButton
                    size="xs"
                    variant="ghost"
                    icon="i-heroicons-eye"
                    @click="viewUser(user)"
                  />
                  <UButton
                    size="xs"
                    variant="ghost"
                    :icon="user.status === 'banned' ? 'i-heroicons-lock-open' : 'i-heroicons-lock-closed'"
                    :color="user.status === 'banned' ? 'green' : 'red'"
                    @click="toggleUserStatus(user)"
                  />
                  <UDropdown
                    :items="[
                      [
                        {
                          label: 'View Details',
                          icon: 'i-heroicons-eye',
                          click: () => viewUser(user)
                        },
                        {
                          label: 'Edit Subscription',
                          icon: 'i-heroicons-credit-card',
                          click: () => editSubscription(user)
                        }
                      ],
                      [
                        {
                          label: 'Delete User',
                          icon: 'i-heroicons-trash',
                          click: () => confirmDelete(user),
                          color: 'red'
                        }
                      ]
                    ]"
                  >
                    <UButton
                      size="xs"
                      variant="ghost"
                      icon="i-heroicons-ellipsis-vertical"
                    />
                  </UDropdown>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Pagination -->
      <div class="flex items-center justify-between px-4 py-3 border-t border-gray-200 dark:border-gray-700">
        <div class="flex items-center gap-2">
          <span class="text-sm text-gray-500">
            Showing {{ ((currentPage - 1) * pageSize) + 1 }} to {{ Math.min(currentPage * pageSize, totalUsers) }} of {{ totalUsers }} users
          </span>
        </div>
        <div class="flex items-center gap-2">
          <UButton
            size="xs"
            variant="outline"
            :disabled="currentPage === 1"
            @click="currentPage--"
          >
            Previous
          </UButton>
          <span class="text-sm px-2">
            Page {{ currentPage }} of {{ totalPages }}
          </span>
          <UButton
            size="xs"
            variant="outline"
            :disabled="currentPage === totalPages"
            @click="currentPage++"
          >
            Next
          </UButton>
        </div>
      </div>
    </UCard>

    <!-- User Details Modal -->
    <UserDetailsModal
      v-if="selectedUser && showUserDetails"
      :user="selectedUser"
      :is-open="showUserDetails"
      @close="showUserDetails = false"
      @update-status="handleStatusUpdate"
      @update-subscription="handleSubscriptionUpdate"
    />

    <!-- Subscription Edit Modal -->
    <SubscriptionEditModal
      v-if="selectedUser && showSubscriptionEdit"
      :user="selectedUser"
      :is-open="showSubscriptionEdit"
      @close="showSubscriptionEdit = false"
      @update="handleSubscriptionUpdate"
    />

    <!-- Delete Confirmation Modal -->
    <UModal v-if="showDeleteConfirm" :model-value="showDeleteConfirm" @update:model-value="showDeleteConfirm = false">
      <UCard>
        <template #header>
          <h3 class="text-lg font-semibold">Confirm Delete</h3>
        </template>
        
        <p class="text-gray-600 dark:text-gray-400">
          Are you sure you want to delete <strong>{{ userToDelete?.name }}</strong>? This action cannot be undone.
        </p>
        
        <template #footer>
          <div class="flex justify-end gap-2">
            <UButton
              variant="outline"
              @click="showDeleteConfirm = false"
            >
              Cancel
            </UButton>
            <UButton
              color="red"
              @click="handleDelete"
            >
              Delete User
            </UButton>
          </div>
        </template>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useUserManagement } from '~/composables/useUserManagement'
import type { User, UserStatus } from '~/composables/useUserManagement'
import UserDetailsModal from './UserDetailsModal.vue'
import SubscriptionEditModal from './SubscriptionEditModal.vue'

const {
  users,
  isLoading,
  selectedUser,
  currentPage,
  pageSize,
  totalUsers,
  totalPages,
  filters,
  updateUserStatus,
  updateUserSubscription,
  deleteUser,
  clearFilters
} = useUserManagement()

const showFilters = ref(false)
const showUserDetails = ref(false)
const showSubscriptionEdit = ref(false)
const showDeleteConfirm = ref(false)
const userToDelete = ref<User | null>(null)


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

const formatDate = (date: Date) => {
  return new Intl.DateTimeFormat('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric'
  }).format(date)
}

const viewUser = (user: User) => {
  selectedUser.value = user
  showUserDetails.value = true
}

const editSubscription = (user: User) => {
  selectedUser.value = user
  showSubscriptionEdit.value = true
}

const confirmDelete = (user: User) => {
  userToDelete.value = user
  showDeleteConfirm.value = true
}

const toggleUserStatus = async (user: User) => {
  const newStatus: UserStatus = user.status === 'banned' ? 'active' : 'banned'
  await updateUserStatus(user.id, newStatus)
}

const handleStatusUpdate = async (userId: number, status: UserStatus) => {
  await updateUserStatus(userId, status)
}

const handleSubscriptionUpdate = async (userId: number, subscription: Partial<User['subscription']>) => {
  await updateUserSubscription(userId, subscription)
}

const handleDelete = async () => {
  if (userToDelete.value) {
    await deleteUser(userToDelete.value.id)
    showDeleteConfirm.value = false
    userToDelete.value = null
  }
}
</script>