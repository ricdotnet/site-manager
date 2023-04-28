<template>
  <div class="w-full border-b dark:border-dark-border border-light-border">
    <div class="content h-12 flex py-2 px-4 justify-between items-center">
      <span>Site-Manager</span>

      <div class="flex">
        <Toggle @ontoggle="toggleTheme" :is-checked="isDark" name="themeToggle" title="Theme toggle"/>
        <LinkButton href="login" text="Login"/>
        <LinkButton href="register" text="Register"/>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { computed } from 'vue';
  import Toggle from './UI/Toggle.vue';
  import LinkButton from './LinkButton.vue';
  import { useTheme } from '../composables/useTheme.ts';

  const { toggleTheme, currentTheme } = useTheme();

  // this theme toggle will be checked if the current theme is dark
  const isDark = computed(() => {
    return currentTheme.value === 'dark';
  });
</script>
