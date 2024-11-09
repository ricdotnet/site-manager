<template>
  <div class="grid grid-cols-left-nav gap-2">
    <div class="left-nav">
      <div data-role="header">General</div>
      <LinkButton :class="isActive(['profile', 'settings']) ? 'active' : ''" href="/dashboard/settings/profile" text="Profile"/>
      <LinkButton :class="isActive('api-keys') ? 'active' : ''" href="/dashboard/settings/api-keys" text="API Keys"/>
      <LinkButton :class="isActive('active-sessions') ? 'active' : ''" href="/dashboard/settings/active-sessions" text="Active Sessions"/>
    </div>
    <div>
      <router-view/>
    </div>
  </div>
</template>

<script setup lang="ts">
import { LinkButton } from '@components';
import { useRoute } from 'vue-router';

const route = useRoute();

const isActive = (name: string | string[]) => {
  if (Array.isArray(name)) {
    return name.includes(route.name as string);
  }
  return route.name === name;
};
</script>

<style scoped lang="scss">
.left-nav {
  a {
    @apply py-2 px-4 hover:bg-dark;
  }

  .active::before {
    content: '';
    border-style: solid;
    border-width: 0 0.25em 0.25em 0;
    display: inline-block;
    height: 0.45em;
    transform: rotate(-45deg);
    width: 0.45em;
    margin-right: 0.5em;
  }

  [data-role="header"] {
    @apply text-light font-bold px-4;
  }
}
</style>
