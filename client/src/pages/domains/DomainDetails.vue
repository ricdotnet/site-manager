<template>
  <div class="flex justify-between mb-5">
    {{ $route.params.domain }}
    <div>
      <Button
        name="add-dns-record"
        color="primary"
        @click="isAddingDnsRecord = true"
      >
        Add DNS Record
      </Button>
    </div>
  </div>
  <router-view />

  <AddDnsRecordDialog
    :is-open="isAddingDnsRecord"
    :on-close-dialog="() => (isAddingDnsRecord = false)"
  />
</template>

<script setup lang="ts">
import { AddDnsRecordDialog, Button } from '@components';
import { ref } from 'vue';

const isAddingDnsRecord = ref(false);
</script>
