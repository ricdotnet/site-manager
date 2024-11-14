<template>
  <div class="flex justify-between">
    <div>
      <Button
        name="reload-nginx"
        color="bordered"
        text="Reload nginx"
        :is-actioning="isReloadingNginx"
        @click="onClickReloadNginx"
      />
    </div>
    <div class="mb-5 flex gap-2">
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
import { messages } from '@utils';
import { useSitesStore } from '@stores';
import { useToaster } from '@composables';

const sitesStore = useSitesStore();
const { addToast } = useToaster();

const isAddingSite = ref(false);
const isOpenDeleteSites = ref(false);
const isDeleting = ref(false);
const isReloadingNginx = ref(false);
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

const onClickReloadNginx = async () => {
  isReloadingNginx.value = true;
  const { error, data } = await sitesStore.reloadNginx();

  if (error) {
    isReloadingNginx.value = false;
    addToast('error', messages.site[error]);
    return;
  }

  isReloadingNginx.value = false;

  if (!data) return;
  addToast('success', messages.site[data]);
};
</script>
