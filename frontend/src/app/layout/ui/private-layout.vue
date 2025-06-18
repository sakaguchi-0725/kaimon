<script setup lang="ts">
import { PrivateHeader } from '@/shared/ui/header'
import { RouterView, useRoute, useRouter } from 'vue-router'
import { computed } from 'vue'
import { NavFooter } from '@/shared/ui/navbar'

const route = useRoute()
const router = useRouter()

const title = computed(() => route.meta.title as string)
const isBackButton = computed(() => route.meta.isBackButton as boolean)
const toBack = computed(() => {
  return route.meta.toBack as string
})

const onBack = () => {
  if (toBack.value) {
    router.push(toBack.value)
  }
}
</script>

<template>
  <div class="flex flex-col h-screen">
    <PrivateHeader :title="title" :is-back-button="isBackButton" @on-back="onBack" />

    <main class="flex-1 overflow-y-auto bg-white">
      <div class="container mx-auto px-4 sm:px-6 md:px-8 lg:px-16 xl:px-24 pb-20">
        <RouterView />
      </div>
    </main>

    <NavFooter />
  </div>
</template>
