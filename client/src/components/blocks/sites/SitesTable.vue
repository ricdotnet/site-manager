<template>
  <template v-if="isLoading">
    <TableLoading />
  </template>
  <template v-else-if="!sitesStore.sites.length">
    <Empty message="You have no sites to show" />
  </template>
  <template v-else>
    <Table>
      <TableHeader>
        <TableRow>
          <TableHead class="w-0 pr-6">
            <Checkbox
              id="site-all"
              name="site-all"
              :checked="allChecked"
              @change="onCheckAll"
              :indeterminate="anyChecked"
            />
          </TableHead>
          <TableHead>Domain</TableHead>
          <TableHead>Config</TableHead>
          <TableHead>Created on</TableHead>
          <TableHead class="text-right">Has SSL</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        <TableRow v-for="site in sitesStore.getSites()" :key="site.id">
          <TableCell>
            <Checkbox
              :id="'site-' + site.id"
              :name="'site-' + site.id"
              :checked="site.checked ?? false"
              @on-change="onCheckSite(site.id)"
            />
          </TableCell>
          <TableCell>
            <span
              class="w-2.5 h-2.5 rounded-full mr-2 inline-block"
              :class="site.enabled ? 'bg-cobalt-green' : 'bg-red-500'"
              :title="isEnabled(site.enabled)"
            ></span>
            <router-link
              :to="'/dashboard/sites/' + site.id"
              class="table__body--col--link"
            >
              {{ site.domain }}
            </router-link>
          </TableCell>
          <TableCell>{{ site.config_name }}</TableCell>
          <TableCell>{{ new Date(site.created_at).toDateString() }}</TableCell>
          <TableCell class="text-right">
            {{ hasSSL(site.domain) }}
          </TableCell>
        </TableRow>
      </TableBody>
    </Table>
  </template>
</template>

<script setup lang="ts">
import {
  Checkbox,
  Empty,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableLoading,
  TableRow,
} from '@components';
import { computed, onMounted, ref } from 'vue';
import { storeToRefs } from 'pinia';
import { useSitesStore } from '@stores';

const sitesStore = useSitesStore();
const { sites, certificates } = storeToRefs(sitesStore);

const fetchError = ref(false);
const isLoading = ref(false);

const allChecked = computed(
  () =>
    !sites.value.filter((site) => !site.checked && !site.config_only).length,
);
const anyChecked = computed(() => sites.value.some((site) => site.checked));

onMounted(async () => {
  isLoading.value = true;

  const error = await sitesStore.fetchSites();
  if (error) fetchError.value = true;

  isLoading.value = false;
});

const onCheckSite = (id: number) => {
  sitesStore.checkSite(id);
};

const onCheckAll = (e: Event & { target: HTMLInputElement }) => {
  sitesStore.checkAll(e.target.checked);
};

const isEnabled = (isEnabled: boolean) => {
  return isEnabled ? 'Site enabled' : 'Site disabled';
};

const hasSSL = (domain: string) => {
  const hasDirectSSL = certificates.value.some(
    (cert) => cert.domains === domain,
  );

  if (hasDirectSSL) {
    return true;
  }

  const domainMatches = domain.match(/([\w-]+.\w{1,4})$/);

  if (!domainMatches || !domainMatches.length) {
    return false;
  }

  const wildcardDomain = `*.${domainMatches[domainMatches.length - 1]}`;
  return certificates.value.some((cert) => cert.domains === wildcardDomain);
};
</script>
