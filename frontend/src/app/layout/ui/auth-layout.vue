<script setup lang="ts">
import { HorizontalLogo } from '@/shared/ui/logo'
import { computed } from 'vue'
import { RouterView, useRoute } from 'vue-router'

const route = useRoute()

const title = computed(() => route.meta.title)
</script>

<template>
  <div class="flex flex-col justify-center px-4 h-screen gap-6">
    <HorizontalLogo class="mx-auto" />
    <h2 class="text-center text-xl font-bold">{{ title }}</h2>
    <RouterView />
  </div>
</template>
