<template>
  <div class="mb-5 flex gap-2 justify-end">
    <template v-if="anySelected">
      <Button
        text="Delete Selected"
        color="danger"
        value="delete-selected"
        name="delete-selected"
        :disabled="isDeleting"
        :is-actioning="isDeleting"
        @click="isOpenDeleteSites = true"
      />
    </template>
    <Button
      text="Add Site"
      color="primary"
      value="add-site"
      name="add-site"
      @click="onClickAddSite"
    />
  </div>
  <SitesTable />

  <template v-if="configOnlySites.length">
    <div
      v-for="configOnlySite of configOnlySites"
      @click="() => onClickAddConfigOnlySite(configOnlySite)"
      :key="configOnlySite.id"
    >
      {{ configOnlySite.config_name }}
    </div>
  </template>

  <AddSiteDialog
    :is-adding-site="isAddingSite"
    :close-dialog="closeAddSiteDialog"
    @on-close-dialog="closeAddSiteDialog"
    :site="configOnlySiteToCreate"
  />
  <DeleteSitesDialog
    :close-dialog="closeDeleteSitesDialog"
    :is-open-delete-sites="isOpenDeleteSites"
    @on-close-dialog="closeDeleteSitesDialog"
  />
</template>

<script setup lang="ts">
import {
  AddSiteDialog,
  Button,
  DeleteSitesDialog,
  SitesTable,
} from '@components';
import { computed, ref } from 'vue';
import type { TSite } from '@types';
import { useSitesStore } from '@stores';

const sitesStore = useSitesStore();
const isAddingSite = ref(false);
const isOpenDeleteSites = ref(false);
const isDeleting = ref(false);
const configOnlySiteToCreate = ref<TSite | undefined>();

const anySelected = computed(() =>
  sitesStore.sites.find((site) => site.checked),
);
const configOnlySites = computed(() =>
  sitesStore.sites.filter((site) => site.config_only),
);

const onClickAddSite = () => {
  isAddingSite.value = true;
};

const onClickAddConfigOnlySite = (site: TSite) => {
  configOnlySiteToCreate.value = site;
  isAddingSite.value = true;
};

const closeAddSiteDialog = () => {
  isAddingSite.value = false;
  configOnlySiteToCreate.value = undefined;
};

const closeDeleteSitesDialog = () => {
  isOpenDeleteSites.value = false;
};
</script>
