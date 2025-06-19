<script setup lang="ts">
import type { Member } from '../model'
import MemberStatus from './member-status.vue'
import RoleIcon from './role-icon.vue'

defineProps<{
  members: Member[] | undefined
}>()
</script>

<template>
  <ul class="flex flex-col gap-4 w-full">
    <li v-for="member in members ?? []" :key="member.id" class="flex justify-between items-center">
      <div class="flex items-center gap-2">
        <RoleIcon :role="member.role" classes="w-8 h-8 text-primary" />
        <div class="flex flex-col gap-1">
          <p class="text-sm font-bold">{{ member.name }}</p>
          <p class="text-xs text-gray-500">{{ member.joinedAt }}</p>
        </div>
      </div>
      <MemberStatus :status="member.status" />
    </li>
  </ul>
</template>
