<template>
  <Dialog title="Delete selected sites"
          confirm-label="Delete"
          :is-open="isOpenDeleteSites"
          :is-actioning="isDeletingSites"
          @on-confirm-dialog="onConfirmDeleteSites">
    Do you really want to delete the selected sites?
    This will delete <b>all sites</b> from the database, <b>all *.conf</b> files, <b>all ssl certificates</b> and <b>disable
    all</b> of them making them inaccessible.
  </Dialog>
</template>

<script setup lang="ts">
  import { ref } from "vue";
  import { Dialog } from "@components";
  import { TSite } from "@types";
  import { useRequest } from "@composables";
  import { useSitesStore } from "@stores";

  const sitesStore = useSitesStore();

  const props = defineProps<{
    isOpenDeleteSites: boolean;
    closeDialog: () => void;
  }>();

  const isDeletingSites = ref(false);

  const onConfirmDeleteSites = async () => {
    isDeletingSites.value = true;

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
    isDeletingSites.value = false;
    props.closeDialog();
  }
</script>

<style scoped lang="scss">
</style>
