<template>
  <div class="loading">
    <div class="loading__header"></div>
    <div class="loading__row"></div>
    <div class="loading__row"></div>
    <div class="loading__row"></div>
    <div class="loading__row"></div>
  </div>
</template>

<script setup lang="ts">
</script>

<style scoped lang="scss">
.loading {
  @apply flex flex-col gap-2;

  &__header {
    @apply w-full py-5 bg-gray-400 dark:bg-zinc-900 animate-pulse rounded-md;
  }

  &__row {
    @apply w-full py-7 bg-gray-300 dark:bg-zinc-800 animate-pulse rounded-md;
  }
}
</style>
