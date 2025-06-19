<script setup lang="ts">
import { XMarkIcon } from '@heroicons/vue/24/solid'

defineProps<{
  isOpen: boolean
  title: string
  closeOnOverlayClick?: boolean
}>()

const emit = defineEmits<{
  close: []
}>()

const onClose = () => {
  emit('close')
}

const handleOverlayClick = (event: MouseEvent) => {
  if (event.target === event.currentTarget) {
    onClose()
  }
}
</script>

<template>
  <Teleport to="body">
    <Transition name="modal">
      <div
        v-if="isOpen"
        class="fixed inset-0 bg-black/50 flex justify-center items-center z-50"
        @click="handleOverlayClick"
      >
        <div
          class="bg-white rounded-lg w-11/12 max-w-md max-h-[90vh] overflow-y-auto shadow-lg"
          @click.stop
        >
          <div class="flex justify-between items-center p-4 border-b border-gray-200">
            <h2 class="text-xl font-semibold text-gray-800">{{ title }}</h2>
            <button class="p-1 rounded-full hover:bg-gray-100 transition-colors" @click="onClose">
              <XMarkIcon class="w-6 h-6 text-gray-500" />
            </button>
          </div>
          <div class="p-4">
            <slot></slot>
          </div>
          <div v-if="$slots.footer" class="p-4 border-t border-gray-200 flex justify-end space-x-2">
            <slot name="footer"></slot>
          </div>
        </div>
      </div>
    </Transition>
  </Teleport>
</template>

<style scoped>
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.3s ease;
}

.modal-enter-from,
.modal-leave-to {
  opacity: 0;
}
</style>
