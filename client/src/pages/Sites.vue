<template>
  <div class="py-5 grid place-content-end">
    <Button text="Add Site" color="primary" value="add-site" name="add-site" @click="onClickAddSite"/>
  </div>
  <Suspense>
    <SitesTable/>
    <template #fallback>
      <div>Loading your sites...</div>
    </template>
  </Suspense>

  <AddSiteDialog :is-adding-site="isAddingSite"
                 :close-dialog="closeAddSiteDialog"
                 @on-close-dialog="isAddingSite = false"/>
</template>

<script setup lang="ts">
  import { ref } from "vue";
  import { AddSiteDialog, Button, SitesTable } from "../components";

  const isAddingSite = ref(false);

  const onClickAddSite = () => isAddingSite.value = true;

  const closeAddSiteDialog = () => isAddingSite.value = false;
</script>

<style scoped lang="scss">
</style>
