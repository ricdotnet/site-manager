<template>
  <router-link :to="href" class="dropdown-item">
    {{ text }}
  </router-link>
</template>

<script setup lang="ts">
defineProps<{
  href: string;
  text: string;
}>();
</script>

<style scoped lang="scss">
  .dropdown-item {
    @apply
    uppercase
    font-light
    text-dark
    dark:text-light
    text-sm
    p-3
    hover:bg-light-lighter
    dark:hover:bg-dark-border
    rounded-md;
  }
</style>
