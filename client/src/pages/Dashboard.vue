<template>
  <div class="header">
    <div class="header__greeting">
      Hello <span class="text-cobalt-green">{{ userStore.username }}</span>, welcome
    </div>
    <div class="header__nav">
      <LinkButton href="/dashboard" text="Overview"/>
      <LinkButton href="/dashboard/sites" text="Sites"/>
      <LinkButton href="/dashboard/domains" text="Domains"/>
      <LinkButton href="/dashboard/settings" text="Settings"/>
    </div>
  </div>
  <router-view></router-view>
</template>

<script setup lang="ts">
  import { useUserStore } from "../stores/user.store.ts";
  import { LinkButton } from "../components";

  const userStore = useUserStore();
</script>

<style scoped lang="scss">
  .header {
    @apply py-10 grid grid-cols-1 gap-y-3;

    &__greeting {
      @apply text-2xl
    }

    &__nav {
      @apply flex;
    }
  }
</style>
