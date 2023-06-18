<template>
  <template v-if="isLoading">
    <TableLoading/>
  </template>
  <template v-else-if="!sitesStore.sites.length">
    <Empty message="You have no sites to show"/>
  </template>
  <template v-else>
    <div class="dark:bg-dark bg-white p-5 overflow-x-auto rounded-md shadow">
      <table class="table">
        <thead class="table__head">
        <tr>
          <th class="table__head--col">
            <Checkbox id="site-all"
                      name="site-all"
                      :checked="allChecked"
                      @change="onCheckAll"/>
          </th>
          <th class="table__head--col">Domain</th>
          <th class="table__head--col">Config</th>
          <th class="table__head--col">Created on</th>
          <th class="table__head--col">Has SSL</th>
        </tr>
        </thead>
        <tbody class="table__body">
        <tr v-for="site in sitesStore.sites" class="group">
          <td class="table__body--col">
            <Checkbox :id="'site-'+site.ID"
                      :name="'site-'+site.ID"
                      :checked="site.checked ?? false"
                      @on-change="onCheckSite(site.ID)"/>
          </td>
          <td class="table__body--col">
            <span class="w-2.5 h-2.5 rounded-full mr-2 inline-block"
                  :class="site.enabled ? 'bg-cobalt-green' : 'bg-red-500'"></span>
            <router-link :to="'/dashboard/sites/' + site.ID" class="table__body--col--link">
              {{ site.domain }}
            </router-link>
          </td>
          <td class="table__body--col">{{ site.config_name }}</td>
          <td class="table__body--col">{{ new Date(site.created_at).toDateString() }}</td>
          <td class="table__body--col">{{ site.has_ssl ? 'Yes' : 'No' }}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </template>
</template>

<script setup lang="ts">
  import { computed, onMounted, ref } from "vue";
  import { Checkbox, Empty, TableLoading } from "../";
  import { useSitesStore } from "../../stores/sites.store.ts";

  const sitesStore = useSitesStore();

  const fetchError = ref(false);
  const isLoading = ref(false);

  const allChecked = computed(() => !sitesStore.sites.filter((site) => !site.checked).length);

  onMounted(async () => {
    isLoading.value = true;

    const error = await sitesStore.fetchSites();
    if (error) fetchError.value = true;

    isLoading.value = false;
  });

  const onCheckSite = (id: number) => {
    sitesStore.checkSite(id);
  }

  const onCheckAll = (e: Event & { target: HTMLInputElement }) => {
    sitesStore.checkAll(e.target.checked);
  }
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
          @apply
          w-auto
          px-3
          py-5;
        }

        @apply
        w-12
        pl-3
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
