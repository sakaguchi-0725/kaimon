<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { GroupInfo, useGroupInfo } from '@/features/group'
import { MemberList, useGetMember } from '@/features/member'

const route = useRoute()
const id = computed(() => {
  const paramId = route.params.id
  return typeof paramId === 'string' ? paramId : (paramId?.[0] ?? '')
})

const { group } = useGroupInfo(() => id.value)
const { members } = useGetMember(() => id.value)
</script>

<template>
  <div class="flex flex-col gap-6">
    <GroupInfo :group="group" :members="members" />
    <MemberList :members="members" />
  </div>
</template>
