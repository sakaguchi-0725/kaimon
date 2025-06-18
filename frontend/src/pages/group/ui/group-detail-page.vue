<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { GroupInfo, useGroupInfo } from '@/features/group'
import { MemberList, useGetMember } from '@/features/member'
import { FullWidthContainer } from '@/shared/ui/container'
import { UserPlusIcon } from '@heroicons/vue/24/solid'

const route = useRoute()
const id = computed(() => {
  const paramId = route.params.id
  return typeof paramId === 'string' ? paramId : (paramId?.[0] ?? '')
})

const { group } = useGroupInfo(() => id.value)
const { members } = useGetMember(() => id.value)
</script>

<template>
  <div class="flex flex-col">
    <FullWidthContainer>
      <GroupInfo :group="group" :members="members" />
    </FullWidthContainer>
    <FullWidthContainer>
      <div class="flex justify-between items-center mb-4">
        <h3 class="text-lg font-bold">メンバー</h3>
        <button>
          <UserPlusIcon class="w-6 h-6 text-primary" />
        </button>
      </div>
      <MemberList :members="members" />
    </FullWidthContainer>
  </div>
</template>
