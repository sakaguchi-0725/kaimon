<script setup lang="ts">
import { RouterLink, useRoute } from 'vue-router'
import { UserGroupIcon, ShoppingBagIcon, Cog6ToothIcon } from '@heroicons/vue/24/outline'

const isActive = (path: string) => {
  return useRoute().path.startsWith(path)
}
</script>

<template>
  <div
    class="fixed bottom-0 left-0 right-0 bg-white shadow-[0_-1px_2px_0_rgba(0,0,0,0.05)] shadow-gray-300 z-20"
  >
    <div class="container mx-auto px-4 pb-4">
      <nav class="flex justify-around py-3">
        <RouterLink
          :to="{ name: 'group-list' }"
          class="flex flex-col items-center text-gray-800 hover:text-primary"
          :class="{ 'text-primary font-bold': isActive('/groups') }"
        >
          <UserGroupIcon class="h-6 w-6" />
          <span class="text-xs mt-1">グループ</span>
        </RouterLink>

        <RouterLink
          to="#"
          class="flex flex-col items-center text-gray-800 hover:text-primary"
          :class="{ 'text-primary font-bold': isActive('/shopping') }"
        >
          <ShoppingBagIcon class="h-6 w-6" />
          <span class="text-xs mt-1">買い物</span>
        </RouterLink>

        <RouterLink
          to="#"
          class="flex flex-col items-center text-gray-800 hover:text-primary"
          :class="{ 'text-primary font-bold': isActive('/settings') }"
        >
          <Cog6ToothIcon class="h-6 w-6" />
          <span class="text-xs mt-1">設定</span>
        </RouterLink>
      </nav>
    </div>
  </div>
</template>
