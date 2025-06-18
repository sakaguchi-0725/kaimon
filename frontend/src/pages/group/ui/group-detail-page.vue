<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import { GroupInfo, useGroupInfo } from '@/features/group'

const route = useRoute()
const id = computed(() => {
  const paramId = route.params.id
  return typeof paramId === 'string' ? paramId : (paramId?.[0] ?? '')
})

const { group, members } = useGroupInfo(() => id.value)
</script>

<template>
  <GroupInfo :group="group" :members="members" />
</template>
