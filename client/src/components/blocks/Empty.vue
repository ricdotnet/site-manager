<template>
  <div class="empty">
    {{ message }}
    <slot/>
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
    grid
    place-content-center
    border
    border-dashed
    border-light-border
    dark:border-dark-border
    rounded-md
    uppercase
    text-center
    text-dark-darker/40
    dark:text-white/40;
  }
</style>
