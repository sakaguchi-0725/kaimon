<script setup lang="ts">
import { RoleIcon } from '@/entities/member'
import MemberStatus from './member-status.vue'
import type { GetGroupResponse, Member } from '../model'
import { UserPlusIcon } from '@heroicons/vue/24/solid'

defineProps<{
  group: GetGroupResponse | undefined
  members: Member[] | undefined
}>()
</script>

<template>
  <div class="flex flex-col gap-5">
    <div class="flex flex-col gap-1">
      <div class="flex justify-between items-center">
        <h3 class="text-lg font-bold">{{ group?.name }}</h3>
        <p class="text-sm text-gray-500">{{ group?.createdAt }}</p>
      </div>
      <p class="text-sm text-gray-500">{{ group?.description }}</p>
    </div>
    <div>
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-bold">メンバー</h3>
        <button>
          <UserPlusIcon class="w-6 h-6 text-primary" />
        </button>
      </div>
      <ul class="flex flex-col gap-4 w-full">
        <li
          v-for="member in members ?? []"
          :key="member.id"
          class="flex justify-between items-center"
        >
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
    </div>
  </div>
</template>
