<template>
  <Nav />

  <div class="content mx-auto px-4">
    <router-view></router-view>
  </div>

  <Toaster />
</template>

<script setup lang="ts">
import { Nav, Toaster } from '@components';
</script>
