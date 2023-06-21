<template>
  <div class="w-full border-b dark:border-dark-border border-light-border">
    <div class="content h-12 flex py-2 px-4 justify-between items-center">
      <span class="text-xl">Site-Manager</span>

      <div class="flex space-x-3">
        <Toggle @ontoggle="toggleTheme"
                :is-checked="isDark"
                :is-theme-toggle="true"
                name="themeToggle"
                title="Theme toggle"/>
        <template v-if="!userStore.isAuthed">
          <LinkButton href="login" text="Login"/>
          <LinkButton href="register" text="Register"/>
        </template>
        <template v-else>
          <LinkButton href="/dashboard" text="Dashboard"/>
          <LinkButton @on-link-click="doLogout" text="Logout"/>
        </template>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { computed } from 'vue';
  import { useRouter } from "vue-router";
  import { LinkButton, Toggle } from "@components";
  import { useTheme } from '@composables';
  import { useUserStore } from "@stores";

  const router = useRouter();
  const { toggleTheme, currentTheme } = useTheme();
  const userStore = useUserStore();

  const doLogout = () => {
    localStorage.removeItem('token');
    userStore.setUsername('');
    userStore.setIsAuthed(false);
    userStore.setUserId('');
    router.push('/login');
  }

  // this theme toggle will be checked if the current theme is dark
  const isDark = computed(() => {
    return currentTheme.value === 'dark';
  });
</script>
