<template>
  <div class="py-5 grid place-content-end">
    <Button text="Add Site" color="primary" value="add-site" name="add-site"/>
  </div>
  <Suspense>
    <SitesTable/>
    <template #fallback>
      <div>Loading your sites...</div>
    </template>
  </Suspense>
</template>

<script setup lang="ts">
  import { Button, SitesTable } from "../components";
</script>

<style scoped lang="scss">
</style>
