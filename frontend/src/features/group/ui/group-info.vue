<script setup lang="ts">
import type { GetGroupResponse } from '../model'

defineProps<{
  group: GetGroupResponse | undefined
}>()
</script>

<template>
  <div class="flex flex-col gap-1">
    <div class="flex justify-between items-center">
      <h3 class="text-lg font-bold">{{ group?.name }}</h3>
      <p class="text-sm text-gray-500">{{ group?.createdAt }}</p>
    </div>
    <p class="text-sm text-gray-500">{{ group?.description }}</p>
  </div>
</template>
