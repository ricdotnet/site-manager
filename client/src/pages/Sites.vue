<template>
  <div class="py-5 flex gap-2 justify-end">
    <template v-if="anySelected">
      <Button text="Delete Selected"
              color="danger"
              value="delete-selected"
              name="delete-selected"
              :disabled="isDeleting"
              :is-actioning="isDeleting"
              @click="isOpenDeleteSites = true"/>
    </template>
    <Button text="Add Site" color="primary" value="add-site" name="add-site" @click="onClickAddSite"/>
  </div>
  <SitesTable/>

  <AddSiteDialog :is-adding-site="isAddingSite"
                 :close-dialog="closeAddSiteDialog"
                 @on-close-dialog="closeAddSiteDialog"/>
  <DeleteSitesDialog :close-dialog="closeDeleteSitesDialog"
                     :is-open-delete-sites="isOpenDeleteSites"
                     @on-close-dialog="closeDeleteSitesDialog"/>
</template>

<script setup lang="ts">
  import { computed, ref } from "vue";
  import { AddSiteDialog, Button, SitesTable } from "@components";
  import { useSitesStore } from "../stores/sites.store.ts";
  import DeleteSitesDialog from "../components/blocks/sites/DeleteSitesDialog.vue";

  const sitesStore = useSitesStore();
  const isAddingSite = ref(false);
  const isOpenDeleteSites = ref(false);

  const anySelected = computed(() => sitesStore.sites.find(site => site.checked));
  const isDeleting = ref(false);

  const onClickAddSite = () => isAddingSite.value = true;

  const closeAddSiteDialog = () => isAddingSite.value = false;
  const closeDeleteSitesDialog = () => isOpenDeleteSites.value = false;
</script>

<style scoped lang="scss">
</style>
