<template>
  <div class="h-12 w-full flex py-2 px-4">
    <LinkButton href="#" text="Register" />
    <LinkButton href="#" text="Login" />
    <button @click="toggle">
      Change Theme
    </button>
  </div>
</template>

<script setup lang="ts">
  import LinkButton from './LinkButton.vue';
  import { useTheme } from '../composables/useTheme.ts';

  const { toggle } = useTheme();
</script>
