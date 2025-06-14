import { ref, markRaw, type Component } from 'vue'
import ToastContainer from '../ui/toast-container.vue'
import type { ToastContainerRef } from '../model'

// グローバルなトーストコンテナの参照
const toastContainerRef = ref<ToastContainerRef | null>(null)

// トーストコンテナコンポーネントを取得
export const getToastContainerComponent = (): Component => {
  return markRaw(ToastContainer)
}

// トーストコンテナの参照を設定する
export const setToastContainerRef = (ref: ToastContainerRef) => {
  toastContainerRef.value = ref
}

// グローバルに使用できるエラートースト表示関数
export const showError = (title: string | undefined, message: string) => {
  if (toastContainerRef.value) {
    toastContainerRef.value.showError(title, message)
  } else {
    console.error('Toast container is not mounted yet')
    console.error(`Error: ${title || 'エラー'} - ${message}`)
  }
}

// グローバルに使用できる成功トースト表示関数
export const showSuccess = (title: string | undefined, message: string) => {
  if (toastContainerRef.value) {
    toastContainerRef.value.showSuccess(title, message)
  } else {
    console.error('Toast container is not mounted yet')
    console.log(`Success: ${title || '成功'} - ${message}`)
  }
}
