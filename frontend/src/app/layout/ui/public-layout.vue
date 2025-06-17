<script setup lang="ts">
import { BaseHeader } from '@/shared/ui/header'
import { ref, computed } from 'vue'
import { RouterView } from 'vue-router'

const isMenuOpen = ref(false)
const isMobile = ref(false)

// 画面サイズの監視
const checkMobile = () => {
  isMobile.value = window.innerWidth < 768
}

// 初期化時と画面サイズ変更時にチェック
if (typeof window !== 'undefined') {
  checkMobile()
  window.addEventListener('resize', checkMobile)
}

const toggleMenu = () => {
  isMenuOpen.value = !isMenuOpen.value

  // メニューが開いている時は背景のスクロールを無効化
  if (isMenuOpen.value) {
    document.body.style.overflow = 'hidden'
  } else {
    document.body.style.overflow = ''
  }
}

// メニューのスタイルをcomputedで動的に計算
const menuStyles = computed(() => {
  if (isMobile.value) {
    return {
      container: 'fixed inset-0 bg-white z-10 pt-20 overflow-y-auto',
      nav: 'container mx-auto px-4 py-6',
      list: 'space-y-6',
      item: 'block text-xl font-medium text-gray-900 hover:text-primary',
    }
  } else {
    return {
      container:
        'absolute right-0 mt-2 w-48 bg-white rounded-md shadow-lg py-1 z-10 mr-4 sm:mr-6 md:mr-8 lg:mr-16 xl:mr-24',
      nav: '',
      list: '',
      item: 'block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100',
    }
  }
})
</script>

<template>
  <div class="flex flex-col min-h-screen">
    <BaseHeader>
      <!-- ハンバーガーボタン -->
      <button
        @click="toggleMenu"
        class="inline-flex items-center justify-center p-2 rounded-md text-gray-600 hover:text-gray-900 hover:bg-gray-100 focus:outline-none z-30"
        aria-label="メニュー"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            v-if="!isMenuOpen"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16M4 18h16"
          />
          <path
            v-else
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M6 18L18 6M6 6l12 12"
          />
        </svg>
      </button>
    </BaseHeader>

    <!-- 統合されたメニュー (SP/PC共通) -->
    <div v-if="isMenuOpen" :class="menuStyles.container" :style="!isMobile ? { top: '4rem' } : {}">
      <nav :class="menuStyles.nav">
        <ul :class="menuStyles.list">
          <li>
            <a href="/login" :class="menuStyles.item">ログイン</a>
          </li>
          <li>
            <a href="/signup" :class="menuStyles.item">会員登録</a>
          </li>
        </ul>
      </nav>
    </div>

    <!-- メインコンテンツ -->
    <main class="flex-grow">
      <div class="container mx-auto px-4 sm:px-6 md:px-8 lg:px-16 xl:px-24">
        <RouterView />
      </div>
    </main>
  </div>
</template>
