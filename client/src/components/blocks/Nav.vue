<template>
  <div class="w-full border-b dark:border-dark-border border-light-border">
    <div class="content h-12 flex py-2 px-4 justify-between items-center">
      <span class="text-xl">
        <RouterLink to="/">Site-Manager</RouterLink>
      </span>

      <div class="flex space-x-3">
        <Toggle
          @ontoggle="toggleTheme"
          :is-checked="isDark"
          :is-theme-toggle="true"
          name="themeToggle"
          title="Theme toggle"
        />
        <template v-if="!userStore.isAuthed">
          <LinkButton href="login" text="Login" />
        </template>
        <template v-else>
          <LinkButton href="/dashboard" text="Dashboard" />
          <LinkButton @on-link-click="doLogout" text="Logout" />
        </template>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { LinkButton, Toggle } from '@components';
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { useTheme } from '@composables';
import { useUserStore } from '@stores';

const router = useRouter();
const { toggleTheme, currentTheme } = useTheme();
const userStore = useUserStore();

const doLogout = async () => {
  await userStore.logout();
  userStore.setUsername('');
  userStore.setIsAuthed(false);
  userStore.setUserId('');
  router.push('/login');
};

const isDark = computed(() => {
  return currentTheme.value === 'dark';
});
</script>
