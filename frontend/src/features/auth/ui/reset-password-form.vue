<script setup lang="ts">
import { PrimaryButton } from '@/shared/ui/button'
import { useResetPassword } from '../lib/use-reset-password'
import { TextInput } from '@/shared/ui/input'

const { defineField, errors, onSubmit } = useResetPassword()

const [email, emailProps] = defineField('email')
</script>

<template>
  <form @submit="onSubmit" class="flex flex-col gap-5">
    <TextInput
      id="email"
      v-model="email"
      v-bind="emailProps"
      :error-message="errors.email"
      type="email"
      placeholder="sharebasket@example.com"
      >メールアドレス</TextInput
    >

    <PrimaryButton type="submit">確認メールを送信</PrimaryButton>
  </form>
</template>
