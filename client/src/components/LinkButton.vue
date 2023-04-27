<template>
  <a :href="props.href" class="flex items-center px-4 hover">
    {{ props.text }}
  </a>
</template>

<script setup lang="ts">
  const props = defineProps<{
    text: string,
    href: string,
  }>();
</script>

<style scoped>
  .hover {
    @apply hover:outline hover:outline-light-dim dark:hover:outline-dark-dim;
  }
</style>
