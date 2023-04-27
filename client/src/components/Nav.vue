<template>
  <div class="w-full border-b dark:border-dark-border border-light-border">
    <div class="content h-12 flex py-2 px-4 justify-between">
      <button @click="toggle">
        Change Theme
      </button>

      <div class="flex">
        <LinkButton href="login" text="Login"/>
        <LinkButton href="register" text="Register"/>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  import LinkButton from './LinkButton.vue';
  import { useTheme } from '../composables/useTheme.ts';

  const { toggle } = useTheme();
</script>
