<template>
  <div class="mb-5 flex gap-2 justify-end">
    <Button
      text="Add API Key"
      color="primary"
      value="add-api-key"
      name="add-api-key"
      @click="onClickAddApiKey"
    />
  </div>
  <div>
    <ApiKeysTable />
  </div>

  <AddApiKeyDialog
    :is-adding-api-key="isAddingApiKey"
    :close-dialog="closeAddApiKeyDialog"
  />
</template>

<script setup lang="ts">
import { AddApiKeyDialog, ApiKeysTable, Button } from '@components';
import { ref } from 'vue';

const isAddingApiKey = ref(false);

const onClickAddApiKey = () => {
  isAddingApiKey.value = true;
};

const closeAddApiKeyDialog = () => {
  isAddingApiKey.value = false;
};
</script>
