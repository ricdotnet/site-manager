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

  <Dialog title="Add a Site"
          confirm-label="Add"
          :is-open="isAddingSite"
          :is-actioning="isPostingSite"
          @on-close-dialog="isAddingSite = false"
          @on-confirm-dialog="onClickConfirmDialog">
    <div class="flex flex-col gap-5">
      <Input id="domain" placeholder="Domain"/>
      <Input id="config" placeholder="Config"/>
    </div>
  </Dialog>
</template>

<script setup lang="ts">
  import { Button, Dialog, Input, SitesTable } from "../components";
  import { ref } from "vue";

  const isAddingSite = ref(false);
  const isPostingSite = ref(false);

  const onClickAddSite = () => {
    isAddingSite.value = true;
  }

  const onClickConfirmDialog = () => {
    isPostingSite.value = true;
    setTimeout(() => {
      isPostingSite.value = false;
      isAddingSite.value = false;
    }, 5000);
  }
</script>

<style scoped lang="scss">
</style>
