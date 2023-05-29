<template>
  <template v-if="error">
    Something went wrong
  </template>
  <template v-else-if="sites && !sites.length">
    <Empty message="You have no sites to show"/>
  </template>
  <template v-else>
    <div class="dark:bg-dark bg-white p-5 overflow-x-auto">
      <table class="table">
        <thead class="table__head">
        <tr>
          <th class="table__head--row">Domain</th>
          <th class="table__head--row">Config</th>
          <th class="table__head--row">Has SSL</th>
        </tr>
        </thead>
        <tbody class="table__body">
        <tr v-for="site in sites" class="group">
          <td class="table__body--row">
            <span class="w-2.5 h-2.5 rounded-full mr-2 inline-block"
                  :class="site.enabled ? 'bg-cobalt-green' : 'bg-red-500'"></span>
            <router-link :to="'/dashboard/sites/' + site.ID">{{ site.domain }}</router-link>
          </td>
          <td class="table__body--row">{{ site.config_name }}</td>
          <td class="table__body--row">{{ site.has_ssl ? 'Yes' : 'No' }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </template>
</template>

<script setup lang="ts">
  import { useRequest } from "../../composables";
  import { ref } from "vue";
  import { TSite, TSitesResponse } from "../../types.ts";
  import { Empty } from "../";

  const sites = ref<TSite[]>();

  const { data, error } = await useRequest<TSitesResponse>({
    endpoint: '/sites/all',
    needsAuth: true,
  });

  sites.value = data.value?.sites;
</script>

<style scoped lang="scss">
  .table {
    @apply table-auto w-full;

    &__head {
      @apply uppercase text-sm bg-light-lighter dark:bg-dark-darker;

      &--row {
        @apply p-3 text-left pl-4;
      }
    }

    &__body {
      @apply divide-y dark:divide-dark-border divide-light-border;

      &--row {
        @apply px-3 py-5 dark:group-hover:bg-dark-darker group-hover:bg-light-lighter transition ease-in-out duration-200 whitespace-nowrap;
      }
    }
  }
</style>
