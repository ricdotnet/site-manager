<template>
  <Dropdown @on-click="onClickDropdown" :is-open="isOpen">
    <template v-if="isBreadcrumbLink" #caller>
      <span class="cursor-pointer text-cobalt-green" @click="onClickDropdown">
        {{ $route.params.type.toString().toUpperCase() }}
      </span>
    </template>
    <template #items>
      <DropdownItem
        :href="`/dashboard/domains/${domainName}/a`"
        text="A Records"
      />
      <DropdownItem
        :href="`/dashboard/domains/${domainName}/aaaa`"
        text="AAAA Records"
      />
      <DropdownItem
        :href="`/dashboard/domains/${domainName}/cname`"
        text="CNAME Records"
      />
      <DropdownItem
        :href="`/dashboard/domains/${domainName}/txt`"
        text="TXT Records"
      />
      <!-- <DropdownItem :href="`domains/${domain.name}/srv`" text="SRV Records"/> -->
      <!-- <DropdownItem :href="`domains/${domain.name}/ns`" text="NS Records" /> -->
      <!-- <DropdownItem :href="`domains/${domain.name}/mx`" text="MX Records"/> -->
    </template>
  </Dropdown>
</template>

<script setup lang="ts">
import { Dropdown, DropdownItem } from '@components';

defineProps<{
  isOpen: boolean;
  onClickDropdown: () => void;
  domainName: string;
  isBreadcrumbLink?: boolean;
}>();
</script>
