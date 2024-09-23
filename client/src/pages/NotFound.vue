<template>
  <div class="h-[calc(100vh-43px)] grid items-center">
    <Empty message="404 Not Found">
      <router-link to="/dashboard">&lt;- dashboard</router-link>
    </Empty>
  </div>
</template>

<script setup lang="ts">
import { Empty } from '@components';
</script>
