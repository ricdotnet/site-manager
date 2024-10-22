<template>
  <Table>
    <TableHeader>
      <TableRow>
        <TableHead>Domain</TableHead>
        <TableHead>Registered on</TableHead>
        <TableHead>Renewal date</TableHead>
        <TableHead class="text-right">More</TableHead>
      </TableRow>
    </TableHeader>
    <TableBody>
      <TableRow v-for="domain in domains" :key="domain.order_id">
        <DomainsTableItem :domain="domain"/>
      </TableRow>
    </TableBody>
  </Table>
</template>

<script setup lang="ts">
import { DomainsTableItem, Table, TableBody, TableHead, TableHeader, TableRow } from '@components';
import { TDomain } from '@types';

defineProps<{
  domains: TDomain[];
}>();
</script>

<style scoped lang="scss">

</style>
