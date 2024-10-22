<template>
  <Dialog title="Add API Key"
          confirm-label="Add"
          :is-open="isAddingApiKey"
          :is-actioning="isPostingApiKey"
          @on-close-dialog="onCloseDialog"
          @on-confirm-dialog="onClickConfirmDialog">
    <div class="flex flex-col gap-5">
      <Input ref="apiKeyInput"
             id="key"
             placeholder="KEY"
             :validator="apiKeyValidator"/>
      <Input ref="apiKeyValueInput"
             id="value"
             placeholder="VALUE"
             :validator="apiKeyValueValidator"/>
      <Transition name="slide-down">
        <div v-if="generalErrorMessage" class="text-center text-red-500" v-html="generalErrorMessage"></div>
      </Transition>
    </div>

    <AddApiKeyExistsDialog
      :is-api-key-exists-dialog-open="isApiKeyExistsDialogOpen"
      :on-confirm-dialog="postApiKey"
      :on-close-dialog="() => isApiKeyExistsDialogOpen = false"
    />
  </Dialog>
</template>

<script setup lang="ts">
import { AddApiKeyExistsDialog, Dialog, Input } from '@components';
import { useApiKeysStore } from '@stores';
import type { InputComponent } from '@types';
import { apiKeyValidator, apiKeyValueValidator } from '@validators';
import { storeToRefs } from 'pinia';
import { ref } from 'vue';

const props = defineProps<{
  isAddingApiKey: boolean;
  closeDialog: () => void;
}>();

const apiKeysStore = useApiKeysStore();
const { apiKeys } = storeToRefs(apiKeysStore);

const isPostingApiKey = ref(false);
const isApiKeyExistsDialogOpen = ref(false);

const apiKeyInput = ref<InputComponent>();
const apiKeyValueInput = ref<InputComponent>();

const formHasError = ref(false);
const generalErrorMessage = ref('');

const onCloseDialog = () => {
  if (isApiKeyExistsDialogOpen.value) {
    isApiKeyExistsDialogOpen.value = false;
    return;
  }

  apiKeyInput.value?.setValue('');
  apiKeyValueInput.value?.setValue('');
  generalErrorMessage.value = '';

  props.closeDialog();
};

const onClickConfirmDialog = async () => {
  formHasError.value = false;

  if (!apiKeyInput.value?.getValue() || !apiKeyValueInput.value?.getValue()) {
    generalErrorMessage.value = 'Please fill in all fields.';
    formHasError.value = true;

    apiKeyInput.value?.setError(true);
    apiKeyValueInput.value?.setError(true);

    return;
  }

  if (apiKeyInput.value?.hasError || apiKeyValueInput.value?.hasError) {
    formHasError.value = true;
    return;
  }

  const apiKeyExists = apiKeys.value.find((apiKey) => apiKey.key === apiKeyInput.value!.getValue());

  if (apiKeyExists?.value) {
    isApiKeyExistsDialogOpen.value = true;
    return;
  }

  postApiKey();
};

const postApiKey = async () => {
  isPostingApiKey.value = true;
  await apiKeysStore.addApiKey({
    key: apiKeyInput.value!.getValue(),
    value: apiKeyValueInput.value!.getValue(),
    is_api_key: true,
  });
  isPostingApiKey.value = false;
  formHasError.value = false;
  isApiKeyExistsDialogOpen.value = false;
  onCloseDialog();
};
</script>
