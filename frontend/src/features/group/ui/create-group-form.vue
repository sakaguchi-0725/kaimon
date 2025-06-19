<script setup lang="ts">
import { PrimaryButton } from '@/shared/ui/button'
import { TextareaInput, TextInput } from '@/shared/ui/input'
import { useCreateGroup } from '../lib/use-create-group'

const emit = defineEmits<{
  onSuccess: []
}>()

const { defineField, errors, onSubmit } = useCreateGroup()

const [name, nameProps] = defineField('name')
const [description, descriptionProps] = defineField('description')

const handleSubmit = onSubmit(() => emit('onSuccess'))
</script>

<template>
  <form @submit.prevent="handleSubmit" class="flex flex-col gap-5">
    <TextInput id="name" type="text" v-model="name" v-bind="nameProps" :error-message="errors.name"
      >グループ名</TextInput
    >
    <TextareaInput
      id="description"
      type="text"
      v-model="description"
      v-bind="descriptionProps"
      :error-message="errors.description"
      >グループの説明</TextareaInput
    >
    <PrimaryButton type="submit">作成</PrimaryButton>
  </form>
</template>
