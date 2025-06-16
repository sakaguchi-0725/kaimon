<script setup lang="ts">
import { PasswordInput, TextInput } from '@/shared/ui/input'
import { useSignup } from '../lib/use-signup'
import { PrimaryButton } from '@/shared/ui/button'

const { defineField, errors, onSubmit } = useSignup()

const [name, nameProps] = defineField('name')
const [email, emailProps] = defineField('email')
const [password, passwordProps] = defineField('password')
</script>

<template>
  <form @submit="onSubmit" class="flex flex-col gap-5">
    <TextInput
      id="name"
      type="text"
      v-model="name"
      v-bind="nameProps"
      :error-message="errors.name"
      placeholder="山田太郎"
      >ユーザー名</TextInput
    >

    <TextInput
      id="email"
      type="email"
      v-model="email"
      v-bind="emailProps"
      :error-message="errors.email"
      placeholder="sharebasket@example.com"
      >メールアドレス</TextInput
    >

    <PasswordInput
      id="password"
      type="password"
      v-model="password"
      v-bind="passwordProps"
      :error-message="errors.password"
      placeholder="パスワード"
      >パスワード</PasswordInput
    >

    <PrimaryButton type="submit">登録</PrimaryButton>
  </form>
</template>
