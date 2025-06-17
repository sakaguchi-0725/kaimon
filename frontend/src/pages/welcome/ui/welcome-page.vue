<script setup lang="ts">
import { RouterLink } from 'vue-router'
</script>

<template>
  <div class="py-12 text-center">
    <h1 class="text-4xl font-bold text-gray-900 mb-6">Welcome to ShareBasket</h1>
    <p class="text-xl text-gray-600 max-w-2xl mx-auto mb-8">
      買い物リストを家族や友人と簡単に共有できるサービスです。
      リアルタイムで買い物状況を確認し、効率的な買い物をサポートします。
    </p>
    <div class="flex justify-center gap-4">
      <RouterLink
        :to="{ name: 'login' }"
        class="bg-primary text-white px-6 py-2 rounded-md hover:bg-primary-dark transition-colors"
      >
        ログイン
      </RouterLink>
      <RouterLink
        :to="{ name: 'signup' }"
        class="bg-white text-primary border border-primary px-6 py-2 rounded-md hover:bg-gray-100 transition-colors"
      >
        会員登録
      </RouterLink>
    </div>
  </div>
</template>
