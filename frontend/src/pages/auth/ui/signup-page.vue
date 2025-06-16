<script setup lang="ts">
import { SignupForm } from '@/features/auth'
import { HorizontalLogo } from '@/shared/ui/logo'
</script>

<template>
  <div class="flex flex-col justify-center px-4 h-screen">
    <div class="mb-6 flex flex-col items-center gap-5">
      <HorizontalLogo />
      <h2 class="text-center text-2xl font-bold">アカウント登録</h2>
    </div>

    <SignupForm />
  </div>
</template>
