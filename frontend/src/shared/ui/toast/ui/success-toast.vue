<script setup lang="ts">
import BaseToast from './base-toast.vue'

defineProps<{
  title: string
  message: string
}>()
</script>

<template>
  <BaseToast type="success" :title="title" :message="message" @close="$emit('close')" />
</template>
