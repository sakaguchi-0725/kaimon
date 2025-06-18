<script setup lang="ts">
import { computed } from 'vue'
import type { MemberStatus } from '../model'
import { CheckCircleIcon } from '@heroicons/vue/24/solid'
import { ClockIcon } from '@heroicons/vue/24/solid'

const props = defineProps<{
  status: MemberStatus
}>()

const statusText = computed(() => {
  switch (props.status) {
    case 'active':
      return '参加中'
    case 'pending':
      return '招待中'
    default:
      return ''
  }
})

const statusColor = computed(() => {
  switch (props.status) {
    case 'active':
      return 'text-green-500'
    case 'pending':
      return 'text-yellow-500'
    default:
      return 'text-gray-500'
  }
})

const statusIcon = computed(() => {
  switch (props.status) {
    case 'active':
      return CheckCircleIcon
    case 'pending':
    default:
      return ClockIcon
  }
})
</script>

<template>
  <div class="flex items-center gap-2">
    <component :is="statusIcon" :class="statusColor" class="w-5 h-5" />
    <p :class="statusColor" class="text-sm">{{ statusText }}</p>
  </div>
</template>
