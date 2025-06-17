<script setup lang="ts">
import BaseHeader from './base-header.vue'
import { ChevronLeftIcon } from '@heroicons/vue/24/outline'

defineProps<{
  title: string
  isBackButton?: boolean
  onBack?: () => void
}>()
</script>

<template>
  <BaseHeader>
    <div class="flex items-center justify-between w-full">
      <div class="w-10 flex items-center">
        <button
          v-if="isBackButton"
          @click="onBack"
          class="rounded-full hover:bg-gray-100 transition-colors"
          aria-label="戻る"
        >
          <ChevronLeftIcon class="h-6 w-6 text-gray-600" />
        </button>
      </div>

      <h1 class="text-xl font-medium text-center flex-1">{{ title }}</h1>

      <div class="w-10"></div>
    </div>
  </BaseHeader>
</template>
