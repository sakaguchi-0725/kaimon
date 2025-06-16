<script setup lang="ts">
import { HorizontalLogo } from '@/shared/ui/logo'
import { RouterView } from 'vue-router'
</script>

<template>
  <div class="flex flex-col justify-center px-4 h-screen gap-6">
    <HorizontalLogo class="mx-auto" />
    <RouterView />
  </div>
</template>
