<script setup lang="ts">
import { computed } from 'vue'
import {
  DISABLED_CLASS,
  ENABLED_CLASS,
  SIZE_CLASSES,
  VARIANT_CLASSES,
  type ButtonSize,
  type ButtonVariant,
} from '../model'

const props = withDefaults(
  defineProps<{
    type?: 'button' | 'submit'
    size?: ButtonSize
    variant?: ButtonVariant
    disabled?: boolean
  }>(),
  {
    type: 'button',
    size: 'full',
    variant: 'primary',
    disabled: false,
  }
)

defineEmits<{
  onClick: []
}>()

const sizeClass = computed(() => SIZE_CLASSES[props.size])

const variantClass = computed(() => VARIANT_CLASSES[props.variant])

const disabledClass = computed(() => (props.disabled ? DISABLED_CLASS : ENABLED_CLASS))

const classes = computed(() => [
  'rounded-md',
  sizeClass.value,
  variantClass.value,
  disabledClass.value,
])
</script>

<template>
  <button
    :class="classes"
    :type="props.type"
    @click="$emit('onClick')"
    :disabled="props.disabled"
    :aria-disabled="props.disabled"
  >
    <slot></slot>
  </button>
</template>
