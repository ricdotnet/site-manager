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
  </Dialog>
</template>

<script setup lang="ts">
import { Dialog, Input } from '@components';
import type { InputComponent } from '@types';
import { apiKeyValidator, apiKeyValueValidator } from '@validators';
import { ref } from 'vue';

const props = defineProps<{
  isAddingApiKey: boolean;
  closeDialog: () => void;
}>();

const isPostingApiKey = ref(false);

const apiKeyInput = ref<InputComponent>();
const apiKeyValueInput = ref<InputComponent>();

const formHasError = ref(false);
const generalErrorMessage = ref('');

const onCloseDialog = () => {
  props.closeDialog();
};

const onClickConfirmDialog = () => {
  formHasError.value = false;

  if (apiKeyInput.value?.hasError || apiKeyValueInput.value?.hasError) {
    formHasError.value = true;
    return;
  }

  isPostingApiKey.value = true;
  setTimeout(() => {
    isPostingApiKey.value = false;
    formHasError.value = false;
    onCloseDialog();
  }, 5000);
};
</script>
