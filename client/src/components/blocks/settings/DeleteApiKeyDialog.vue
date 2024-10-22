<template>
  <Dialog title="Delete API Key"
          confirm-label="Delete"
          :is-open="isOpenDeleteApiKey"
          :is-actioning="isDeletingApiKey"
          @on-close-dialog="onCloseDialog"
          @on-confirm-dialog="onClickConfirmDialog">
    Do you really want to delete this API Key?<br>
    All integrations that use this key will stop working.
  </Dialog>
</template>

<script setup lang="ts">
import { Dialog } from '@components';
import { useApiKeysStore } from '@stores';
import { TApiKey } from '@types';
import { ref } from 'vue';

const apiKeysStore = useApiKeysStore();
const { deleteApiKey } = apiKeysStore;

const props = defineProps<{
  isOpenDeleteApiKey: boolean;
  closeDialog: () => void;
  apiKey: TApiKey;
}>();

const isDeletingApiKey = ref(false);

const onCloseDialog = () => {
  props.closeDialog();
};

const onClickConfirmDialog = async () => {
  isDeletingApiKey.value = true;
  const error = await deleteApiKey(props.apiKey);
  if (error) {
    alert(error);
  }
  isDeletingApiKey.value = false;
  onCloseDialog();
};
</script>
