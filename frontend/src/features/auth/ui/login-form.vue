<script setup lang="ts">
import { TextInput, PasswordInput } from '@/shared/ui/input'
import { useLogin } from '../lib/use-login'
import { PrimaryButton } from '@/shared/ui/button'

const { defineField, errors, onSubmit } = useLogin()
const [email, emailProps] = defineField('email')
const [password, passwordProps] = defineField('password')
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

    <PasswordInput
      id="password"
      v-model="password"
      v-bind="passwordProps"
      :error-message="errors.password"
      placeholder="パスワード"
      >パスワード</PasswordInput
    >

    <PrimaryButton type="submit">ログイン</PrimaryButton>
  </form>
</template>
