<template>
  <template v-if="isLoadingDomains">
    <TableLoading/>
  </template>
  <template v-else-if="!domains.length">
    <Empty class="mt-10" message="You have not registered any domains."/>
  </template>
  <template v-else>
    <DomainsTable :domains="domains"/>
  </template>
</template>

<script setup lang="ts">
import { DomainsTable, Empty, TableLoading } from '@components';
import { useRequest } from '@composables';
import { onMounted, ref } from 'vue';

const domains = ref([]);
const isLoadingDomains = ref(true);

onMounted(async () => {
  const { data, error } = await useRequest({
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
