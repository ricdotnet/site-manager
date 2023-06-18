<template>
  <div class="py-5 flex gap-2 justify-end">
    <template v-if="anySelected">
      <Button text="Delete Selected"
              color="danger"
              value="delete-selected"
              name="delete-selected"
              :disabled="isDeleting"
              :is-actioning="isDeleting"
              @click="onClickDeleteSites"/>
    </template>
    <Button text="Add Site" color="primary" value="add-site" name="add-site" @click="onClickAddSite"/>
  </div>
  <SitesTable/>

  <AddSiteDialog :is-adding-site="isAddingSite"
                 :close-dialog="closeAddSiteDialog"
                 @on-close-dialog="isAddingSite = false"/>
</template>

<script setup lang="ts">
  import { computed, ref } from "vue";
  import { AddSiteDialog, Button, SitesTable } from "../components";
  import { useSitesStore } from "../stores/sites.store.ts";
  import { TSite } from "../types.ts";
  import { useRequest } from "../composables";

  const sitesStore = useSitesStore();
  const isAddingSite = ref(false);

  const anySelected = computed(() => sitesStore.sites.find(site => site.checked));
  const isDeleting = ref(false);

  const onClickAddSite = () => isAddingSite.value = true;

  const closeAddSiteDialog = () => isAddingSite.value = false;

  const onClickDeleteSites = async () => {
    isDeleting.value = true;

    const sitesToDelete = sitesStore.sites.reduce((sites: number[], site: TSite) => {
      if (site.checked) sites.push(site.ID);
      return sites;
    }, []);

    const { error } = await useRequest({
      endpoint: '/site/',
      method: 'DELETE',
      payload: {
        sites: sitesToDelete,
      },
      needsAuth: true,
    });

    if (error) throw new Error(error);

    sitesStore.removeSites(sitesToDelete);
    isDeleting.value = false;
  }
</script>

<style scoped lang="scss">
</style>
