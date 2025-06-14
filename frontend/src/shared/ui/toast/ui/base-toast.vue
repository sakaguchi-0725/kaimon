<script setup lang="ts">
import { computed } from 'vue'
import { ExclamationCircleIcon, CheckCircleIcon, XMarkIcon } from '@heroicons/vue/24/solid'
import type { ToastType } from '../model'

const props = defineProps<{
  type: ToastType
  message: string
  title: string
}>()

defineEmits<{
  close: []
}>()

const toastClasses = computed(() => {
  const baseClasses = 'p-3 rounded-md w-full border shadow-sm'

  if (props.type === 'error') {
    return `${baseClasses} bg-error-light border-error`
  } else {
    return `${baseClasses} bg-success-light border-success`
  }
})

const textColorClass = computed(() => {
  return props.type === 'error' ? 'text-error' : 'text-success'
})

const icon = computed(() => {
  return props.type === 'error' ? ExclamationCircleIcon : CheckCircleIcon
})
</script>

<template>
  <div :class="toastClasses">
    <div class="flex items-center justify-between mb-1">
      <div class="flex items-center gap-1">
        <component :is="icon" class="w-5 h-5 flex-shrink-0" :class="textColorClass" />
        <p class="font-medium text-sm" :class="textColorClass">{{ title }}</p>
      </div>
      <button @click="$emit('close')" :class="textColorClass">
        <XMarkIcon class="w-5 h-5" />
      </button>
    </div>
    <p class="text-xs text-gray-600 pl-5">{{ message }}</p>
  </div>
</template>
