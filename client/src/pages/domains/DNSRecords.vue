<template>
  <Table :is-loading="isLoadingRecords">
    <TableHeader>
      <TableRow>
        <TableHead>Name</TableHead>
        <TableHead>Value</TableHead>
        <TableHead>TTL</TableHead>
        <TableHead>Status</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow v-for="record in records" :key="record.id">
        <TableCell>{{ record.host }}</TableCell>
        <TableCell>{{ record.value }}</TableCell>
        <TableCell>{{ record.ttl }}</TableCell>
        <TableCell>{{ record.status }}</TableCell>
      </TableRow>
    </TableBody>
  </Table>
</template>

<script setup lang="ts">
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@components';
import { onMounted, ref } from 'vue';
import { useRequest } from '@composables';
import { useRoute } from 'vue-router';

const route = useRoute();

// eslint-disable-next-line @typescript-eslint/no-explicit-any
const records = ref<any>([]);
const isLoadingRecords = ref(true);

onMounted(async () => {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const { data, error } = await useRequest<any>({
    endpoint: `/domains/${route.params.domain}/${route.params.type}`,
    needsAuth: true,
  });

  if (error) {
    console.error(error);
  } else {
    records.value = data.message.records.records;
  }

  isLoadingRecords.value = false;
});
</script>

<style scoped lang="scss"></style>
