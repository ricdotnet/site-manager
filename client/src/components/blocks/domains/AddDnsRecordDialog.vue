<template>
  <Dialog
    :is-open="isOpen"
    :title="`Add a new ${route.params.type.toString().toUpperCase()} record`"
    confirm-label="Add"
    @on-close-dialog="closeDialog"
    :show-actions="false"
    :is-actioning="isSubmittingForm"
  >
    <Form
      ref="formRef"
      :cancelable="true"
      :is-submitting="isSubmittingForm"
      :error-message="'Error message'"
      :on-submit-form="onSubmitForm"
      :on-cancel-form="onCloseDialog"
    >
      <component :is="component" :on-clear-form-error="clearFormError" />
    </Form>
  </Dialog>
</template>

<script setup lang="ts">
import {
  AAAARecordForm,
  ARecordForm,
  CNAMERecordForm,
  Dialog,
  Form,
  TXTRecordForm,
} from '@components';
import type { FormComponent, TDNSRecordFormProcess } from '@types';
import { computed, ref } from 'vue';
import { useEvents, useToaster } from '@composables';
import { useDnsRecordsStore } from '@stores';
import { useRoute } from 'vue-router';

const props = defineProps<{
  isOpen: boolean;
  onCloseDialog: () => void;
}>();

const route = useRoute();
const toaster = useToaster();
const events = useEvents();

const { addDnsRecord } = useDnsRecordsStore();

const formRef = ref<FormComponent>();
const isSubmittingForm = ref(false);

const component = computed(() => {
  switch (route.params.type.toString().toLocaleLowerCase()) {
    case 'a':
      return ARecordForm;
    case 'aaaa':
      return AAAARecordForm;
    case 'cname':
      return CNAMERecordForm;
    case 'txt':
      return TXTRecordForm;
    default:
      return ARecordForm;
  }
});

const onSubmitForm = async (): Promise<void> => {
  const { hasError, data, error } = events.emit(
    'process-inputs',
  ) as TDNSRecordFormProcess;
  formRef.value?.setFormError(error);
  if (hasError || !data) {
    isSubmittingForm.value = false;
    return;
  }

  isSubmittingForm.value = true;

  try {
    await addDnsRecord(data);
  } catch (error) {
    let message: string | undefined;

    if (error instanceof Error) {
      message = error.message;
    }

    onError(message);
    return;
  }

  onSuccess();
};

const onSuccess = () => {
  toaster.addToast('success', 'New DNS record added.');
  closeDialog();
};

const onError = (errorMessage?: string) => {
  toaster.addToast(
    'error',
    errorMessage ?? 'Unable to add the new DNS record.',
  );
  closeDialog();
};

const clearFormError = () => {
  formRef.value?.clearFormError();
};

const closeDialog = () => {
  isSubmittingForm.value = false;
  props.onCloseDialog();
};
</script>
