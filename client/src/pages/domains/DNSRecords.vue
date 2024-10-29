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
          <PencilIcon class="w-5 cursor-not-allowed" />
          <Button :name="`delete-${record.host}`" color="icon">
            <TrashIcon
              class="w-5"
              @click="() => onClickDeleteDnsRecord(record)"
            />
          </Button>
        </TableCell>
      </TableRow>
    </TableBody>
  </Table>

  <DeleteDnsRecordDialog
    :is-open="isDeletingDnsRecord"
    :on-close-dialog="onCloseDeleteDnsRecordDialog"
    :dns-record="dnsRecord"
  />
</template>

<script setup lang="ts">
import {
  Button,
  DeleteDnsRecordDialog,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@components';
import { PencilIcon, TrashIcon } from '@heroicons/vue/20/solid';
import { onMounted, ref } from 'vue';
import type { TDNSRecord } from '@types';
import { storeToRefs } from 'pinia';
import { useDnsRecordsStore } from '@stores';

const dnsRecordsStore = useDnsRecordsStore();

const isLoadingRecords = ref(true);
const isDeletingDnsRecord = ref(false);
const dnsRecord = ref<undefined | TDNSRecord>();

const { dnsRecords } = storeToRefs(dnsRecordsStore);

onMounted(async () => {
  await dnsRecordsStore.fetchDnsRecords();
  isLoadingRecords.value = false;
});

const onClickDeleteDnsRecord = (record: TDNSRecord) => {
  dnsRecord.value = record;
  isDeletingDnsRecord.value = true;
};

const onCloseDeleteDnsRecordDialog = () => {
  dnsRecord.value = undefined;
  isDeletingDnsRecord.value = false;
};
</script>

<style scoped lang="scss"></style>
