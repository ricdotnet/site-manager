<template>
  <div class="header">
    <div class="header__greeting">
      Hello <span class="text-cobalt-green">{{ userStore.username }}</span>, welcome
    </div>
    <div class="header__nav">
      <LinkButton :class="{ active: isActive('overview') }" href="/dashboard" text="Overview"/>
      <LinkButton :class="{ active: isActive(['sites', 'site-details']) }" href="/dashboard/sites" text="Sites"/>
      <LinkButton :class="{ active: isActive('domains') }" href="/dashboard/domains" text="Domains"/>
      <LinkButton :class="{ active: isActive('settings') }" href="/dashboard/settings" text="Settings"/>
    </div>
  </div>
  <router-view></router-view>
</template>

<script setup lang="ts">
  import { useUserStore } from "../stores/user.store.ts";
  import { LinkButton } from "../components";
  import { useRoute } from "vue-router";

  const userStore = useUserStore();
  const route = useRoute();

  const isActive = (name: string | string[]) => {
    if (name instanceof Array) {
      return name.includes(route.name);
    }
    return route.name === name;
  }
</script>

<style scoped lang="scss">
  .header {
    @apply
    pt-8
    grid
    grid-cols-1
    gap-y-3
    border-b
    border-light-border
    dark:border-dark-border;

    &__greeting {
      @apply text-3xl;
    }

    &__nav {
      @apply flex;

      * {
        @apply pb-5 relative;
      }

      .active {
        @apply
        before:block
        before:absolute
        before:h-0
        before:left-2
        before:right-2
        before:bottom-[-1px]
        before:border-b-2
        before:border-dark
        dark:before:border-white;
      }
    }
  }
</style>
