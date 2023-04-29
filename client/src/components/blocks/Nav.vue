<template>
  <div class="w-full border-b dark:border-dark-border border-light-border">
    <div class="content h-12 flex py-2 px-4 justify-between items-center">
      <span>Site-Manager</span>

      <div class="flex space-x-3">
        <Toggle @ontoggle="toggleTheme"
                :is-checked="isDark"
                :is-theme-toggle="true"
                name="themeToggle"
                title="Theme toggle"/>
        <LinkButton href="login" text="Login"/>
        <LinkButton href="register" text="Register"/>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { computed } from 'vue';
  import { LinkButton, Toggle } from "../";
  import { useTheme } from '../../composables';

  const { toggleTheme, currentTheme } = useTheme();

  // this theme toggle will be checked if the current theme is dark
  const isDark = computed(() => {
    return currentTheme.value === 'dark';
  });
</script>
