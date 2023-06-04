<template>
  <template v-if="error">
    Something went wrong
  </template>
  <template v-else-if="!sitesStore.sites">
    <Empty message="You have no sites to show"/>
  </template>
  <template v-else>
    <div class="dark:bg-dark bg-white p-5 overflow-x-auto rounded-md shadow">
      <table class="table">
        <thead class="table__head">
        <tr>
          <th class="table__head--col">Domain</th>
          <th class="table__head--col">Config</th>
          <th class="table__head--col">Has SSL</th>
          <th class="table__head--col w-0"></th>
        </tr>
        </thead>
        <tbody class="table__body">
        <tr v-for="site in sitesStore.sites" class="group">
          <td class="table__body--col">
            <span class="w-2.5 h-2.5 rounded-full mr-2 inline-block"
                  :class="site.enabled ? 'bg-cobalt-green' : 'bg-red-500'"></span>
            <router-link :to="'/dashboard/sites/' + site.ID" class="table__body--col--link">
              {{ site.domain }}
            </router-link>
          </td>
          <td class="table__body--col">{{ site.config_name }}</td>
          <td class="table__body--col">{{ site.has_ssl ? 'Yes' : 'No' }}</td>
          <td class="table__body--col">
            <Button name="delete" value="delete" id="delete" color="danger">
              <TrashIcon class="w-5"/>
            </Button>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </template>
</template>

<script setup lang="ts">
  import { useRequest } from "../../composables";
  import { TSitesResponse } from "../../types.ts";
  import { Button, Empty } from "../";
  import { TrashIcon } from "@heroicons/vue/20/solid";
  import { useSitesStore } from "../../stores/sites.store.ts";

  const sitesStore = useSitesStore();

  // wait 30 seconds until doing another fetch... a refresh will always fetch
  // TODO: move this to the store instead
  if (sitesStore.lastFetch === 0 || (Date.now() - sitesStore.lastFetch > 30000)) {
    const { data, error } = await useRequest<TSitesResponse>({
      endpoint: '/sites/all',
      needsAuth: true,
    });

    if (error) {
      throw new Error('Something went wrong when fetching your sites...');
    }

    sitesStore.sites = [...data?.sites];
    sitesStore.setLastFetch();
  }
</script>

<style scoped lang="scss">
  .table {
    @apply table-auto w-full;

    &__head {
      @apply uppercase text-sm bg-light-lighter dark:bg-dark-darker;

      &--col {
        @apply p-3 text-left pl-4;
      }
    }

    &__body {
      @apply divide-y divide-light-border dark:divide-dark-border;

      &--col {
        @apply
        px-3
        py-5
        transition
        ease-in-out
        duration-200
        whitespace-nowrap
        group-hover:bg-light-lighter
        dark:group-hover:bg-dark-darker;

        &--link {
          @apply
          underline
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
