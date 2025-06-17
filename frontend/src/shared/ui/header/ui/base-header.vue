<script setup lang="ts">
import { HorizontalLogo } from '@/shared/ui/logo'
</script>

<template>
  <header class="text-gray-600 body-font z-20 relative bg-white">
    <div
      class="container mx-auto flex justify-between items-center py-3 px-4 sm:px-6 md:px-8 lg:px-16 xl:px-24"
    >
      <HorizontalLogo />
      <slot></slot>
    </div>
  </header>
</template>
