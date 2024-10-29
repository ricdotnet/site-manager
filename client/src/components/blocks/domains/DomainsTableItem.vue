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
    <Dropdown
      @on-click="isDropdownOpen = !isDropdownOpen"
      :is-open="isDropdownOpen"
    >
      <template #items>
        <DropdownItem :href="`domains/${domain.name}/a`" text="A Records" />
        <DropdownItem
          :href="`domains/${domain.name}/aaaa`"
          text="AAAA Records"
        />
        <DropdownItem
          :href="`domains/${domain.name}/cname`"
          text="CNAME Records"
        />
        <DropdownItem :href="`domains/${domain.name}/txt`" text="TXT Records" />
        <!-- <DropdownItem :href="`domains/${domain.name}/srv`" text="SRV Records"/> -->
        <!-- <DropdownItem :href="`domains/${domain.name}/ns`" text="NS Records" /> -->
        <!-- <DropdownItem :href="`domains/${domain.name}/mx`" text="MX Records"/> -->
      </template>
    </Dropdown>
  </TableCell>
</template>

<script setup lang="ts">
import { Dropdown, DropdownItem, TableCell } from '@components';
import { ArrowTopRightOnSquareIcon } from '@heroicons/vue/20/solid';
import type { TDomain } from '@types';
import { ref } from 'vue';

defineProps<{
  domain: TDomain;
}>();

const isDropdownOpen = ref(false);
</script>
