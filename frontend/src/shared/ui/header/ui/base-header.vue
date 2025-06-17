<script setup lang="ts">
defineProps<{
  classes?: string
}>()
</script>

<template>
  <header
    :class="`text-gray-800 body-font z-20 relative bg-white top-0 h-14 sm:h-16 md:h-20 flex items-center ${classes}`"
  >
    <div
      class="container mx-auto flex justify-between items-center px-4 sm:px-6 md:px-8 lg:px-16 xl:px-24"
    >
      <slot></slot>
    </div>
  </header>
</template>
