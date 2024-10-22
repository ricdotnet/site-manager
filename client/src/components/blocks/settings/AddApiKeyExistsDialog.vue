<template>
  <Dialog
    title="API Key already exists"
    :is-open="isApiKeyExistsDialogOpen"
    @on-close-dialog="onCloseDialog"
    @on-confirm-dialog="onConfirmDialog"
  >
    This API Key already exists. Confirming will update with the new value.
  </Dialog>
</template>

<script setup lang="ts">
import { Dialog } from '@components';

defineProps<{
  isApiKeyExistsDialogOpen: boolean;
  onConfirmDialog: () => void;
  onCloseDialog: () => void;
}>();
</script>

<style scoped lang="scss">

</style>
