<script setup lang="ts">
import BaseHeader from './base-header.vue'
import { ChevronLeftIcon } from '@heroicons/vue/24/outline'

const props = defineProps<{
  title: string
  isBackButton?: boolean
  onBack?: () => void
}>()

const emit = defineEmits<{
  'on-back': []
}>()

const handleBack = () => {
  if (props.onBack) {
    props.onBack()
  }
  emit('on-back')
}
</script>

<template>
  <BaseHeader classes="shadow-[0_1px_2px_0_rgba(0,0,0,0.05)] shadow-gray-300">
    <div class="flex items-center justify-between w-full">
      <div class="w-10 flex items-center">
        <button
          v-if="isBackButton"
          @click="handleBack"
          class="rounded-full hover:bg-gray-100 transition-colors"
          aria-label="戻る"
        >
          <ChevronLeftIcon class="h-6 w-6 text-gray-600" />
        </button>
      </div>

      <h1 class="text-lg font-bold text-center flex-1">{{ title }}</h1>

      <div class="w-10"></div>
    </div>
  </BaseHeader>
</template>
