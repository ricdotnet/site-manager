<template>
  <Nav />

  <div class="content mx-auto px-4">
    <router-view></router-view>
  </div>

  <Toasts />
</template>

<script setup lang="ts">
import { Nav, Toasts } from '@components';
</script>
