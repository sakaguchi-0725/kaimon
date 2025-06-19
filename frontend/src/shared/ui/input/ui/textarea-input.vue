<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  id: string
  placeholder?: string
  errorMessage?: string
  disabled?: boolean
  rows?: number
}>()

const model = defineModel<string>()

const classes = computed(() => ({
  'border p-2 rounded-sm focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent':
    true,
  'border-error': props.errorMessage !== undefined,
  'border-gray-400': props.errorMessage === undefined,
  'bg-gray-100 cursor-not-allowed': props.disabled,
}))
</script>

<template>
  <div class="flex flex-col">
    <label :for="id" class="text-gray-800 mb-1">
      <slot></slot>
    </label>
    <textarea
      :id="id"
      v-model="model"
      v-bind="$attrs"
      :class="classes"
      :placeholder="placeholder"
      :disabled="disabled"
      :rows="props.rows"
    ></textarea>
    <p v-if="errorMessage" class="text-error">{{ errorMessage }}</p>
  </div>
</template>
