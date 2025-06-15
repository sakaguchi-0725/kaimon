<script setup lang="ts">
import BaseButton from './base-button.vue'
import type { ButtonSize } from '../model'

defineProps<{
  size?: ButtonSize
  disabled?: boolean
  type?: 'button' | 'submit'
}>()
</script>

<template>
  <BaseButton :size="size" variant="outline" :disabled="disabled" v-bind="$attrs" :type="type">
    <slot></slot>
  </BaseButton>
</template>
