<script setup lang="ts">
import { CreateGroupForm, GroupList, useGroupList } from '@/features/group'
import { OutlineButton, PrimaryButton } from '@/shared/ui/button'
import { FullWidthContainer } from '@/shared/ui/container'
import { BaseModal } from '@/shared/ui/modal'
import { useModal } from '@/shared/ui/modal/lib/use-modal'

const { groups, hasGroups, onDetail } = useGroupList()
const { isOpen, open, close } = useModal()
</script>

<template>
  <FullWidthContainer>
    <h2 class="text-lg font-bold mb-4">参加中のグループ</h2>

    <GroupList v-if="hasGroups" :groups="groups" @on-detail="onDetail" />
    <div v-else class="flex flex-col gap-4 justify-center items-center h-full">
      <p class="text-gray-500 text-center text-sm">まだ、グループに参加していません。</p>
      <p class="text-gray-500 text-center text-sm">
        グループを作成するか、他のグループに<br />
        参加してください。
      </p>
    </div>

    <span class="flex justify-center border-b border-gray-200 w-full my-8"></span>

    <div class="flex flex-col gap-2 w-full">
      <PrimaryButton @on-click="open">グループを作成</PrimaryButton>
      <p class="text-sm text-gray-500 text-center">または</p>
      <OutlineButton>グループに参加</OutlineButton>
    </div>
  </FullWidthContainer>

  <BaseModal :is-open="isOpen" title="グループを作成" @close="close">
    <CreateGroupForm @on-success="close" />
  </BaseModal>
</template>
