<script setup lang="ts">
import BaseButton from './base-button.vue'
import type { ButtonSize } from '../model'

defineProps<{
  size?: ButtonSize
  disabled?: boolean
  type?: 'button' | 'submit'
}>()

defineEmits<{
  onClick: []
}>()
</script>

<template>
  <BaseButton
    :size="size"
    variant="primary"
    :disabled="disabled"
    :type="type"
    v-bind="$attrs"
    @on-click="$emit('onClick')"
  >
    <slot></slot>
  </BaseButton>
</template>
