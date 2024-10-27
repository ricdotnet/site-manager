<template>
  <Table :is-loading="isLoadingRecords">
    <TableHeader>
      <TableRow>
        <TableHead>Name</TableHead>
        <TableHead>Value</TableHead>
        <TableHead>TTL</TableHead>
        <TableHead>Status</TableHead>
        <TableHead />
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow v-for="record in dnsRecords" :key="record.id">
        <TableCell>{{ record.host }}</TableCell>
        <TableCell>{{ record.value }}</TableCell>
        <TableCell>{{ record.ttl }}</TableCell>
        <TableCell>{{ record.status }}</TableCell>
        <TableCell class="flex justify-end gap-2">
          <PencilIcon class="w-5" />
          <TrashIcon class="w-5" />
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>
</template>

<script setup lang="ts">
import { PencilIcon, TrashIcon } from '@heroicons/vue/20/solid';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@components';
import { onMounted, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useDnsRecordsStore } from '@stores';

const dnsRecordsStore = useDnsRecordsStore();

const isLoadingRecords = ref(true);

const { dnsRecords } = storeToRefs(dnsRecordsStore);

onMounted(async () => {
  await dnsRecordsStore.fetchDnsRecords();
  isLoadingRecords.value = false;
});
</script>

<style scoped lang="scss"></style>
