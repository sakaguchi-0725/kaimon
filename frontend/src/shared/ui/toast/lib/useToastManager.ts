import { ref, markRaw, type Component } from 'vue'
import ToastContainer from '../ui/toast-container.vue'
import type { ToastContainerRef } from '../model'

const toastContainerRef = ref<ToastContainerRef | null>(null)

export const useToastManager = () => {
  const getToastContainerComponent = (): Component => {
    return markRaw(ToastContainer)
  }

  const setToastContainerRef = (ref: ToastContainerRef) => {
    toastContainerRef.value = ref
  }

  const showError = (message: string, title?: string) => {
    if (toastContainerRef.value) {
      toastContainerRef.value.showError(title, message)
    } else {
      console.error('Toast container is not mounted yet')
    }
  }

  const showSuccess = (message: string, title?: string) => {
    if (toastContainerRef.value) {
      toastContainerRef.value.showSuccess(title, message)
    } else {
      console.error('Toast container is not mounted yet')
    }
  }

  return {
    getToastContainerComponent,
    setToastContainerRef,
    showError,
    showSuccess,
  }
}
