<template>
  <div class="flex justify-between mb-5">
    <Breadcrumbs>
      <span>{{ $route.params.domain }}</span>
      <BreadcrumbSeparator />
      <DnsRecordTypeDropdown
        :is-open="isDropdownOpen"
        :domain-name="<string>$route.params.domain"
        :on-click-dropdown="onClickDropdown"
        :is-breadcrumb-link="true"
      />
    </Breadcrumbs>
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
  <router-view :key="<string>$route.params.type" />

  <AddDnsRecordDialog
    :is-open="isAddingDnsRecord"
    :on-close-dialog="() => (isAddingDnsRecord = false)"
  />
</template>

<script setup lang="ts">
import {
  AddDnsRecordDialog,
  BreadcrumbSeparator,
  Breadcrumbs,
  Button,
  DnsRecordTypeDropdown,
} from '@components';
import { ref } from 'vue';

const isAddingDnsRecord = ref(false);
const isDropdownOpen = ref(false);

const onClickDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value;
};
</script>
