<script setup lang="ts">
import { useToastManager } from '@/shared/ui/toast'
import { onMounted, ref } from 'vue'

const { getToastContainerComponent, setToastContainerRef, showError, showSuccess } =
  useToastManager()
const ToastContainerComponent = getToastContainerComponent()
const toastContainerRef = ref(null)

onMounted(() => {
  if (toastContainerRef.value) {
    setToastContainerRef(toastContainerRef.value)

    // 3秒後にエラートーストを表示
    setTimeout(() => {
      showError('エラーが発生しました。再試行してください。')

      // エラーの1秒後に成功トーストを表示
      setTimeout(() => {
        showSuccess('データが正常に保存されました。')
      }, 1000)
    }, 3000)
  }
})
</script>

<template>
  <div>
    <h1 class="text-3xl font-bold underline">Hello World</h1>
    <ToastContainerComponent ref="toastContainerRef" />
  </div>
</template>
