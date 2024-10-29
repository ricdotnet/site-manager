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
      ref="aliasRef"
      id="dns-alias"
      placeholder="Alias"
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
const aliasRef = ref<InputComponent>();
const ttlRef = ref<InputComponent>();

const clearFormError = () => {
  props.onClearFormError();
};

const processInputs = (): TDNSRecordFormProcess => {
  if (
    hostnameRef.value?.hasError ||
    aliasRef.value?.hasError ||
    ttlRef.value?.hasError
  ) {
    return { hasError: true, error: 'Please check your form' };
  }

  const hostname = hostnameRef.value?.getValue();
  const alias = aliasRef.value?.getValue();
  const ttl = ttlRef.value?.getValue();

  if (!hostname || !alias || !ttl) {
    if (!hostname) hostnameRef.value?.setError(true);
    if (!alias) aliasRef.value?.setError(true);
    if (!ttl) ttlRef.value?.setError(true);

    return { hasError: true, error: 'All values are mandatory.' };
  }

  return {
    data: {
      host: hostname,
      value: alias,
      ttl,
    },
  };
};

events.on('process-inputs', processInputs);
</script>
