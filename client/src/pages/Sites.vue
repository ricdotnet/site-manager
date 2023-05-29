<template>
  <Suspense>
    <Sites/>
    <template #fallback>
      <div>Loading your sites...</div>
    </template>
  </Suspense>
</template>

<script setup lang="ts">
import { Sites } from "../components";
</script>

<style scoped lang="scss">
</style>
