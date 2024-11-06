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
        <TableRow v-for="site in sitesStore.getSites()" :key="site.ID">
          <TableCell>
            <Checkbox
              :id="'site-' + site.ID"
              :name="'site-' + site.ID"
              :checked="site.checked ?? false"
              @on-change="onCheckSite(site.ID)"
            />
          </TableCell>
          <TableCell>
            <span
              class="w-2.5 h-2.5 rounded-full mr-2 inline-block"
              :class="site.enabled ? 'bg-cobalt-green' : 'bg-red-500'"
              :title="isEnabled(site.enabled)"
            ></span>
            <router-link
              :to="'/dashboard/sites/' + site.ID"
              class="table__body--col--link"
            >
              {{ site.domain }}
            </router-link>
          </TableCell>
          <TableCell>{{ site.config_name }}</TableCell>
          <TableCell>{{ new Date(site.created_at).toDateString() }}</TableCell>
          <TableCell class="text-right">
            {{ site.has_ssl ? 'Yes' : 'No' }}
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
const { sites } = storeToRefs(sitesStore);

const fetchError = ref(false);
const isLoading = ref(false);

const allChecked = computed(
  () => !sites.value.filter((site) => !site.checked).length,
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
</script>

<style scoped lang="scss">
.table {
  @apply w-full;

  &__head {
    @apply uppercase text-sm bg-light-lighter dark:bg-dark-darker;

    &--col {
      @apply py-3 pl-3 text-left;

      &:first-child {
        @apply rounded-l-md;
      }

      &:last-child {
        @apply rounded-r-md;
      }
    }
  }

  &__body {
    @apply divide-y divide-light-border dark:divide-dark-border;

    &--col {
      &:not(:first-child) {
        @apply w-auto
        px-3
        py-5;
      }

      @apply w-12
      pl-3
      transition
      ease-in-out
      duration-200
      whitespace-nowrap
      group-hover:bg-light-lighter
      dark:group-hover:bg-dark-darker;

      &--link {
        @apply underline
        underline-offset-2
        decoration-dashed
        decoration-1
        decoration-dark/60
        dark:decoration-white/40;
      }
    }
  }
}
</style>
