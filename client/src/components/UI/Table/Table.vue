<template>
  <template v-if="isLoading">
    <TableLoading />
  </template>
  <template v-else>
    <table class="w-full table-auto">
      <slot />
    </table>
  </template>
</template>

<script setup lang="ts">
import { TableLoading } from '@components';

defineProps<{
  isLoading?: boolean;
}>();
</script>
