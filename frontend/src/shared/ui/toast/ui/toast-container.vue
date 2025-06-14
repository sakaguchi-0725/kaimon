<script setup lang="ts">
import { useToastContainer } from '../lib/use-toast-container'
import ErrorToast from './error-toast.vue'
import SuccessToast from './success-toast.vue'

const { toasts, removeToast, showError, showSuccess } = useToastContainer()

// 外部から使用できるようにメソッドを公開
defineExpose({
  showError,
  showSuccess,
})
</script>

<template>
  <div
    class="toast-container fixed top-4 left-1/2 -translate-x-1/2 z-50 flex flex-col gap-2 w-full max-w-xs mx-auto"
  >
    <template v-for="toast in toasts" :key="toast.id">
      <div class="toast-item w-full" :class="{ removing: toast.removing }">
        <ErrorToast
          v-if="toast.type === 'error'"
          :title="toast.title"
          :message="toast.message"
          @close="removeToast(toast.id)"
        />
        <SuccessToast
          v-else
          :title="toast.title"
          :message="toast.message"
          @close="removeToast(toast.id)"
        />
      </div>
    </template>
  </div>
</template>

<style scoped>
.toast-container {
  pointer-events: none;
}

.toast-container > * {
  pointer-events: auto;
}

.toast-item {
  animation: toast-in 0.3s ease-out forwards;
  transition: all 0.3s ease-out;
}

.toast-item.removing {
  animation: toast-out 0.3s ease-out forwards;
}

@keyframes toast-in {
  from {
    transform: translateY(-20px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}

@keyframes toast-out {
  from {
    transform: translateY(0);
    opacity: 1;
  }
  to {
    transform: translateY(-20px);
    opacity: 0;
  }
}
</style>
