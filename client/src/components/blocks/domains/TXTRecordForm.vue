<template>
  <form class="flex flex-col gap-5">
    <Input
      ref="hostnameRef"
      id="dns-hostname"
      placeholder="Host Name"
      :validator="hostnameValidator"
      @on-key-up="clearFormError"
    />
    <Input
      ref="valueRef"
      id="dns-value"
      placeholder="Value"
      :validator="aliasValidator"
      @on-key-up="clearFormError"
    />
    <Input
      ref="ttlRef"
      id="dns-ttl"
      placeholder="TTL"
      value="38400"
      :validator="ttlValidator"
      @on-key-up="clearFormError"
    />
  </form>
</template>

<script setup lang="ts">
import type {
  InputComponent,
  TDNSRecordFormProcess,
  TDNSRecordFormProps,
} from '@types';
import { aliasValidator, hostnameValidator, ttlValidator } from '@validators';
import { Input } from '@components';
import { ref } from 'vue';
import { useEvents } from '@composables';

const props = defineProps<TDNSRecordFormProps>();

const events = useEvents();

const hostnameRef = ref<InputComponent>();
const valueRef = ref<InputComponent>();
const ttlRef = ref<InputComponent>();

const clearFormError = () => {
  props.onClearFormError();
};

const processInputs = (): TDNSRecordFormProcess => {
  if (
    hostnameRef.value?.hasError ||
    valueRef.value?.hasError ||
    ttlRef.value?.hasError
  ) {
    return { hasError: true, error: 'Please check your form' };
  }

  const hostname = hostnameRef.value?.getValue();
  const value = valueRef.value?.getValue();
  const ttl = ttlRef.value?.getValue();

  if (!hostname || !value || !ttl) {
    if (!hostname) hostnameRef.value?.setError(true);
    if (!value) valueRef.value?.setError(true);
    if (!ttl) ttlRef.value?.setError(true);

    return { hasError: true, error: 'All values are mandatory.' };
  }

  return {
    data: {
      host: hostname,
      value,
      ttl,
    },
  };
};

events.on('process-inputs', processInputs);
</script>
