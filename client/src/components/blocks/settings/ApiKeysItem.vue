<template>
  <TableCell>{{ apiKeyItem.key }}</TableCell>
  <TableCell class="text-ellipsis overflow-hidden">
    <span v-if="isHidden"> ••••••••••••••• </span>
    <span v-else>
      {{ apiKeyItem.value }}
    </span>
  </TableCell>
  <TableCell class="flex justify-end">
    <Button
      color="icon"
      value="toggle-visibility"
      name="toggle-visibility"
      @click="isHidden = !isHidden"
    >
      <template v-if="isHidden">
        <EyeIcon class="h-5 w-5" />
      </template>
      <template v-else>
        <EyeSlashIcon class="h-5 w-5" />
      </template>
    </Button>
    <!--    <Button color="icon" value="edit-api-key" name="edit-api-key" @click="onClickEditApiKey(key)">-->
    <!--      <PencilIcon class="h-5 w-5"/>-->
    <!--    </Button>-->
    <Button
      color="icon"
      value="delete-api-key"
      name="delete-api-key"
      @click="() => onClickDeleteApiKey()"
    >
      <TrashIcon class="h-5 w-5" />
    </Button>
  </TableCell>
  <DeleteApiKeyDialog
    :is-open-delete-api-key="isOpenDeleteApiKeyDialog"
    :close-dialog="onCloseDeleteApiKeyDialog"
    :api-key="apiKeyItem"
  />
</template>

<script setup lang="ts">
import { Button, DeleteApiKeyDialog, TableCell } from '@components';
import { EyeIcon, EyeSlashIcon, TrashIcon } from '@heroicons/vue/20/solid';
import type { TApiKey } from '@types';
import { ref } from 'vue';

const isHidden = ref(true);
const isOpenDeleteApiKeyDialog = ref(false);

defineProps<{
  apiKeyItem: TApiKey;
}>();

const onClickDeleteApiKey = () => {
  isOpenDeleteApiKeyDialog.value = true;
};

const onCloseDeleteApiKeyDialog = () => {
  isOpenDeleteApiKeyDialog.value = false;
};
</script>
