<script setup lang="ts">
import { useResetPasswordConfirm } from '../lib/use-reset-password-confirm'
import { TextInput } from '@/shared/ui/input'
import { PrimaryButton } from '@/shared/ui/button'

const { email, defineField, errors, onSubmit } = useResetPasswordConfirm()

const [code, codeProps] = defineField('code')
</script>

<template>
  <form @submit="onSubmit" class="flex flex-col gap-5">
    <TextInput
      id="email"
      v-model="email"
      type="email"
      placeholder="sharebasket@example.com"
      :disabled="true"
      >メールアドレス</TextInput
    >

    <TextInput
      id="confirmationCode"
      type="text"
      v-model="code"
      v-bind="codeProps"
      :error-message="errors.code"
      >確認コード</TextInput
    >
    <PrimaryButton type="submit">パスワードリセット</PrimaryButton>
  </form>
</template>
