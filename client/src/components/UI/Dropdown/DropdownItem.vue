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
  @apply uppercase
    font-light
    text-dark-darker
    dark:text-light
    text-sm
    p-3
    hover:bg-light-border
    dark:hover:bg-dark
    rounded-md;
}
</style>
