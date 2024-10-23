<template>
  <template v-if="isLoadingDomains">
    <TableLoading />
  </template>
  <template v-else-if="!domains.length">
    <Empty class="mt-10" message="You have not registered any domains." />
  </template>
  <template v-else>
    <DomainsTable :domains="domains" />
  </template>
</template>

<script setup lang="ts">
import { DomainsTable, Empty, TableLoading } from '@components';
import { onMounted, ref } from 'vue';
import { useRequest } from '@composables';

const domains = ref([]);
const isLoadingDomains = ref(true);

onMounted(async () => {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const { data, error } = await useRequest<any>({
    endpoint: '/domains/all',
    needsAuth: true,
  });

  if (error) {
    console.error(error);
  } else {
    domains.value = data.message.domains;
  }

  isLoadingDomains.value = false;
});
</script>
