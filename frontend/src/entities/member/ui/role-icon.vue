<script setup lang="ts">
import { UserGroupIcon, UserIcon } from '@heroicons/vue/24/solid'
import { computed } from 'vue'
import type { Role } from '../model'

const props = defineProps<{
  role: Role
  classes?: string
}>()

const Icon = computed(() => {
  switch (props.role) {
    case 'admin':
      return UserGroupIcon
    case 'member':
      return UserIcon
    default:
      return UserIcon
  }
})
</script>
<template>
  <component :is="Icon" :class="classes" />
</template>
