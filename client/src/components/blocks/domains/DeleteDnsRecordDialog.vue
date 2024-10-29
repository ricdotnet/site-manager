<template>
  <Dialog
    :is-open="isOpen"
    :is-actioning="isDeletingDnsRecord"
    @on-close-dialog="onCloseDialog"
    @on-confirm-dialog="onDeleteDnsRecord"
  >
    Are you sure you want to delete this DNS record?
  </Dialog>
</template>

<script setup lang="ts">
import { Dialog } from '@components';
import type { TDNSRecord } from '@types';
import { ref } from 'vue';
import { useDnsRecordsStore } from '@stores';
import { useToaster } from '@composables';

const { addToast } = useToaster();
const { deleteDnsRecord } = useDnsRecordsStore();

const props = defineProps<{
  isOpen: boolean;
  onCloseDialog: () => void;
  dnsRecord?: TDNSRecord;
}>();

const isDeletingDnsRecord = ref(false);

const onDeleteDnsRecord = async () => {
  if (!props.dnsRecord) return;

  isDeletingDnsRecord.value = true;

  try {
    await deleteDnsRecord({
      host: props.dnsRecord.host,
      value: props.dnsRecord.value,
      ttl: props.dnsRecord.ttl,
    });
  } catch (err) {
    console.log(err);
    isDeletingDnsRecord.value = false;
    return;
  }

  addToast('success', 'DNS record deleted.');
  isDeletingDnsRecord.value = false;
  props.onCloseDialog();
};
</script>
