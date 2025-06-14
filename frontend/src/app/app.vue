<script setup lang="ts">
import {
  getToastContainerComponent,
  setToastContainerRef,
  showError,
  showSuccess,
} from '@/shared/ui/toast'
import { onMounted, ref } from 'vue'
import { RouterView } from 'vue-router'

const ToastContainerComponent = getToastContainerComponent()
const toastContainerRef = ref(null)

onMounted(() => {
  if (toastContainerRef.value) {
    setToastContainerRef(toastContainerRef.value)

    // 3秒後にエラートーストを表示
    setTimeout(() => {
      showError('トーストテスト', 'エラーが発生しました。再試行してください。')

      // エラーの1秒後に成功トーストを表示
      setTimeout(() => {
        showSuccess(undefined, 'データが正常に保存されました。')
      }, 1000)
    }, 4000)
  }
})
</script>

<template>
  <div>
    <RouterView />
    <ToastContainerComponent ref="toastContainerRef" />
  </div>
</template>
