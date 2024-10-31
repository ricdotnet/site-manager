<template>
  <TableCell>
    <span class="flex gap-2 items-center">
      <a :href="'https://' + domain.name" target="_blank">
        {{ domain.name }}
        <ArrowTopRightOnSquareIcon class="w-4 h-4 inline-block" />
      </a>
    </span>
  </TableCell>
  <TableCell>
    {{ new Date(parseInt(domain.created_at) * 1000).toDateString() }}
  </TableCell>
  <TableCell>
    {{ new Date(parseInt(domain.renewal_at) * 1000).toDateString() }}
  </TableCell>
  <TableCell class="flex justify-end">
    <DnsRecordTypeDropdown
      :is-open="isDropdownOpen"
      :domain-name="domain.name"
      :on-click-dropdown="onClickDropdown"
    />
  </TableCell>
</template>

<script setup lang="ts">
import { DnsRecordTypeDropdown, TableCell } from '@components';
import { ArrowTopRightOnSquareIcon } from '@heroicons/vue/20/solid';
import type { TDomain } from '@types';
import { ref } from 'vue';

defineProps<{
  domain: TDomain;
}>();

const isDropdownOpen = ref(false);

const onClickDropdown = () => {
  isDropdownOpen.value = !isDropdownOpen.value;
};
</script>
