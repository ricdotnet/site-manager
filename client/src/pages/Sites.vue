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

  <Dialog :is-open="isAddingSite" @on-close-dialog="isAddingSite = false"/>
</template>

<script setup lang="ts">
  import { Button, Dialog, SitesTable } from "../components";
  import { ref } from "vue";

  const isAddingSite = ref(false);

  const onClickAddSite = () => {
    isAddingSite.value = true;
  }
</script>

<style scoped lang="scss">
</style>
