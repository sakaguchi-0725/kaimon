<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  id: string
  type: 'text' | 'email'
  placeholder?: string
  errorMessage?: string
}>()

const model = defineModel({ required: true })

const classes = computed(() => ({
  'border p-2 rounded-sm focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent':
    true,
  'border-error': props.errorMessage !== undefined,
  'border-gray-400': props.errorMessage === undefined,
}))
</script>

<template>
  <div class="flex flex-col">
    <label :for="id" class="text-gray-800 mb-1">
      <slot></slot>
    </label>
    <input
      :id="id"
      :type="type"
      v-model="model"
      v-bind="$attrs"
      :class="classes"
      :placeholder="placeholder"
    />
    <p v-if="errorMessage" class="text-error">{{ errorMessage }}</p>
  </div>
</template>
