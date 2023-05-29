<template>
  <div class="empty">
    {{ message }}
  </div>
</template>

<script setup lang="ts">
  defineProps<{
    message: string;
  }>();
</script>

<style scoped lang="scss">
  .empty {
    @apply
    w-full
    h-48
    flex
    justify-center
    items-center
    border
    border-dashed
    border-light-border
    dark:border-dark-border
    rounded-md
    uppercase
    text-sm;
  }
</style>
