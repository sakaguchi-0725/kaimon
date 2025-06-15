<script setup lang="ts">
import { LoginForm } from '@/features/auth'
import { HorizontalLogo } from '@/shared/ui/logo'
</script>

<template>
  <div class="flex flex-col justify-center px-4 h-screen">
    <div class="mb-6 flex flex-col items-center gap-5">
      <HorizontalLogo />
      <h2 class="text-center text-2xl font-bold">ログイン</h2>
    </div>

    <LoginForm />
  </div>
</template>
