<template>
  <Suspense>
    <Site/>
    <template #fallback>
      Loading site...
    </template>
  </Suspense>
</template>

<script setup lang="ts">
  import { Site } from "../components";
</script>

<style scoped lang="scss">
</style>
