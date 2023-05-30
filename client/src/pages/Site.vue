<template>
  <div class="mt-10"></div>
  <Suspense>
    <SiteItem/>
    <template #fallback>
      Loading site...
    </template>
  </Suspense>
</template>

<script setup lang="ts">
  import { SiteItem } from "../components";
</script>

<style scoped lang="scss">
</style>
