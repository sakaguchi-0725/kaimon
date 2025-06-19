<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  color?: 'white' | 'orange' | 'gray'
}>()

const bgColor = computed(() => {
  switch (props.color) {
    case 'orange':
      return 'bg-orange-50'
    case 'gray':
      return 'bg-gray-100'
    case 'white':
    default:
      return 'bg-white'
  }
})
</script>

<template>
  <div class="full-width-wrapper py-5" :class="bgColor">
    <div class="content-container px-4 sm:px-6 md:px-8 lg:px-16 xl:px-24">
      <slot></slot>
    </div>
  </div>
</template>

<style scoped>
.full-width-wrapper {
  width: 100vw;
  position: relative;
  left: 50%;
  right: 50%;
  margin-left: -50vw;
  margin-right: -50vw;
}

.content-container {
  width: 100%;
  max-width: 100%;
  margin: 0 auto;
}
</style>
