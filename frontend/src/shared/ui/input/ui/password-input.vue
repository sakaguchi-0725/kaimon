<script setup lang="ts">
import { EyeIcon, EyeSlashIcon } from '@heroicons/vue/24/outline'
import { computed, ref } from 'vue'

// 親からの属性継承を無効化
defineOptions({
  inheritAttrs: false,
})

const props = defineProps<{
  id: string
  placeholder?: string
  errorMessage?: string
}>()

const model = defineModel({ required: true })

const isShow = ref(false)

const toggleShow = async (event: Event) => {
  event.preventDefault()
  isShow.value = !isShow.value
}

const icon = computed(() => (isShow.value ? EyeIcon : EyeSlashIcon))

const classes = computed(() => ({
  'border p-2 rounded-sm focus:outline-none focus:ring-2 focus:ring-primary focus:border-transparent w-full':
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
    <div class="relative">
      <input
        :id="id"
        v-model="model"
        v-bind="$attrs"
        :class="classes"
        :placeholder="placeholder"
        :type="isShow ? 'text' : 'password'"
      />
      <button
        type="button"
        @click="toggleShow($event)"
        class="absolute inset-y-0 end-0 flex items-center px-3 cursor-pointer text-gray-400 focus:text-primary"
      >
        <component :is="icon" class="size-6" />
      </button>
    </div>
    <p v-if="errorMessage" class="text-error">{{ errorMessage }}</p>
  </div>
</template>
