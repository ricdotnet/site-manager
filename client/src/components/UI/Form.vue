<template>
  <form class="flex flex-col gap-5" @submit="handleFormSubmit">
    <slot />

    <Transition name="slide-down">
      <div v-if="errorMessage" class="text-red-500">
        {{ errorMessage }}
      </div>
    </Transition>

    <div class="flex justify-end gap-2">
      <Button
        v-if="cancelable"
        type="button"
        name="cancel-form"
        color="gray"
        text="Cancel"
        :disabled="isSubmitting"
        @click="handleFormCancel"
      />
      <Button
        type="submit"
        name="form-submit"
        color="primary"
        text="Submit"
        :is-actioning="isSubmitting"
      />
    </div>
  </form>
</template>

<script setup lang="ts">
import { onUnmounted, ref } from 'vue';
import { Button } from '@components';

onUnmounted(() => {
  clearErrorTimeout();
});

const props = defineProps<{
  isSubmitting: boolean;
  onSubmitForm: (e: Event) => void;
  cancelable?: boolean;
  onCancelForm?: () => void;
}>();

const errorMessage = ref<string>();
const errorClearTimeout = ref<number>();

const handleFormSubmit = (e: Event) => {
  e.preventDefault();
  props.onSubmitForm(e);
};

const handleFormCancel = () => {
  if (!props.onCancelForm) {
    console.warn(
      'Cancelling a cancelable form without a cancel function handler.',
    );
    return;
  }

  props.onCancelForm();
};

const clearErrorTimeout = () => {
  if (errorClearTimeout.value) {
    clearTimeout(errorClearTimeout.value);
    errorClearTimeout.value = undefined;
  }
};

const setFormError = (error: string | undefined) => {
  if (!error) return;

  errorMessage.value = error;

  clearErrorTimeout();

  // @ts-expect-error nodejs uses NodeJS.Timeout but the browser uses a number
  errorClearTimeout.value = setTimeout(
    () => (errorMessage.value = undefined),
    10000,
  );
};

const clearFormError = () => {
  errorMessage.value = undefined;
  clearErrorTimeout();
};

defineExpose({
  setFormError,
  clearFormError,
});
</script>
