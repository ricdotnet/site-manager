<template>
  <template v-if="isLoadingRecords">
    <TableLoading />
  </template>
  <template v-else-if="!dnsRecords.length">
    <Empty
      message="No DNS record available."
      v-if="!isLoadingRecords && !dnsRecords.length"
    />
  </template>
  <template v-else>
    <Table :is-loading="isLoadingRecords">
      <TableHeader>
        <TableRow>
          <TableHead class="w-[10%]">Name</TableHead>
          <TableHead>Value</TableHead>
          <TableHead class="hidden lg:table-cell w-[1px] whitespace-nowrap">
            TTL
          </TableHead>
          <TableHead class="hidden lg:table-cell w-[1px] whitespace-nowrap">
            Status
          </TableHead>
          <TableHead class="w-[1px] whitespace-nowrap" />
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="record in dnsRecords" :key="record.id">
          <TableCell class="max-w-[150px] overflow-auto">
            {{ record.host }}
          </TableCell>
          <TableCell class="max-w-[1px] overflow-auto">
            {{ record.value }}
          </TableCell>
          <TableCell class="hidden lg:table-cell">{{ record.ttl }}</TableCell>
          <TableCell class="hidden lg:table-cell">
            {{ record.status }}
          </TableCell>
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
  </template>

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
  Empty,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableLoading,
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
