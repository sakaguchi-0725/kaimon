<script setup lang="ts">
import { TextInput } from '@/shared/ui/input'
import { useSignupConfirm } from '../lib/use-signup-confirm'
import { PrimaryButton } from '@/shared/ui/button'

const { email, defineField, errors, onSubmit } = useSignupConfirm()

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
    <PrimaryButton type="submit">登録</PrimaryButton>
  </form>
</template>
