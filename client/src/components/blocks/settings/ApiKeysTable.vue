<template>
  <template v-if="!apiKeys.length">
    <Empty message="You have not added any third party API Keys." />
  </template>
  <template v-else>
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead class="w-[25%]">Key</TableHead>
          <TableHead>Value</TableHead>
          <TableHead class="w-[15%]"></TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="apiKey in apiKeys" :key="apiKey.ID">
          <ApiKeysItem :apiKeyItem="apiKey" />
        </TableRow>
      </TableBody>
    </Table>
  </template>
</template>

<script setup lang="ts">
import {
  ApiKeysItem,
  Empty,
  Table,
  TableBody,
  TableHead,
  TableHeader,
  TableRow,
} from '@components';
import { onMounted } from 'vue';
import { storeToRefs } from 'pinia';
import { useApiKeysStore } from '@stores';

const apiKeysStore = useApiKeysStore();
const { apiKeys } = storeToRefs(apiKeysStore);

onMounted(async () => {
  await apiKeysStore.fetchApiKeys();
});
</script>
