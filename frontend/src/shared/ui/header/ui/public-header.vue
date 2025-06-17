<script setup lang="ts">
import { HorizontalLogo } from '@/shared/ui/logo'
import BaseHeader from './base-header.vue'
</script>

<template>
  <BaseHeader>
    <HorizontalLogo />
    <slot></slot>
  </BaseHeader>
</template>
