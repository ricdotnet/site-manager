<template>
  <router-link :to="href" class="flex items-center px-2 hover">
    {{ text }}
  </router-link>
</template>

<script setup lang="ts">
  defineProps<{
    text: string,
    href: string,
  }>();
</script>

<style scoped>
  .hover {
    /*@apply outline-none hover:outline hover:outline-light-dim dark:hover:outline-dark-dim;*/
    @apply dark:hover:text-light-dim hover:text-dark-dim;
    @apply ease-in-out duration-300 rounded-md;
  }
</style>
