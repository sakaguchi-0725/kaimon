<script setup lang="ts">
import type { JoinedGroupsResponse } from '../model'
import { OutlineButton } from '@/shared/ui/button'

defineProps<{
  groups: JoinedGroupsResponse['groups']
}>()

defineEmits<{
  onDetail: [id: string]
}>()
</script>

<template>
  <div
    v-for="group in groups"
    :key="group.id"
    class="flex border border-gray-200 rounded-md p-3 items-center mb-2"
  >
    <div class="flex-1 flex flex-col gap-1">
      <p class="text-sm font-bold">{{ group.name }}</p>
      <p class="text-xs text-gray-500">{{ group.description }}</p>
    </div>

    <div>
      <OutlineButton size="sm" @click="$emit('onDetail', group.id)">詳細</OutlineButton>
    </div>
  </div>
</template>
