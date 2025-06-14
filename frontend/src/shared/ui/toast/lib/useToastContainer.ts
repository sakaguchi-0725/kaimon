import { ref } from 'vue'
import type { ToastItem, ToastType } from '../model'

export const useToastContainer = () => {
  const toasts = ref<ToastItem[]>([])
  let nextId = 0

  const showToast = (type: ToastType, title: string, message: string) => {
    const id = nextId++
    toasts.value.push({ id, type, title, message, removing: false })

    // 3秒後にトーストを削除
    setTimeout(() => {
      removeToast(id)
    }, 3000)
  }

  const removeToast = (id: number) => {
    const index = toasts.value.findIndex((toast: ToastItem) => toast.id === id)
    if (index !== -1 && toasts.value[index]) {
      // まず削除中フラグを設定
      toasts.value[index].removing = true

      // アニメーション時間（0.3秒）後に実際に削除
      setTimeout(() => {
        const currentIndex = toasts.value.findIndex((toast: ToastItem) => toast.id === id)
        if (currentIndex !== -1) {
          toasts.value.splice(currentIndex, 1)
        }
      }, 300)
    }
  }

  const showError = (title: string | undefined, message: string) => {
    const defaultTitle = title || 'エラーが発生しました'
    showToast('error', defaultTitle, message)
  }

  const showSuccess = (title: string | undefined, message: string) => {
    const defaultTitle = title || '成功しました'
    showToast('success', defaultTitle, message)
  }

  return {
    toasts,
    removeToast,
    showError,
    showSuccess,
  }
}
