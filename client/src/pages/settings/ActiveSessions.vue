<template>
  <Suspense>
    <ActiveSessionsTable />

    <template #fallback>
      <Loading class="mx-auto" />
    </template>
  </Suspense>
</template>

<script setup lang="ts">
import { ActiveSessionsTable, Loading } from '@components';
</script>
