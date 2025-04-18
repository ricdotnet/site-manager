<template>
  <Input
    ref="hostnameRef"
    id="dns-hostname"
    placeholder="Host Name"
    :validator="hostnameValidator"
    @on-key-up="clearFormError"
  />
  <Input
    ref="ipv6addressRef"
    id="dns-ipv6address"
    placeholder="IPv6 Address"
    :validator="ipv6Validator"
    @on-key-up="clearFormError"
  />
  <Input
    ref="ttlRef"
    id="dns-ttl"
    placeholder="TTL"
    value="28800"
    :validator="ttlValidator"
    @on-key-up="clearFormError"
  />
</template>

<script setup lang="ts">
import type {
  InputComponent,
  TDNSRecordFormProcess,
  TDNSRecordFormProps,
} from '@types';
import { hostnameValidator, ipv6Validator, ttlValidator } from '@validators';
import { Input } from '@components';
import { ref } from 'vue';
import { useEvents } from '@composables';

const props = defineProps<TDNSRecordFormProps>();

const events = useEvents();

const hostnameRef = ref<InputComponent>();
const ipv6addressRef = ref<InputComponent>();
const ttlRef = ref<InputComponent>();

const clearFormError = () => {
  props.onClearFormError();
};

const processInputs = (): TDNSRecordFormProcess => {
  if (
    hostnameRef.value?.hasError ||
    ipv6addressRef.value?.hasError ||
    ttlRef.value?.hasError
  ) {
    return { hasError: true, error: 'Please check your form' };
  }

  const hostname = hostnameRef.value?.getValue();
  const ipAddress = ipv6addressRef.value?.getValue();
  const ttl = ttlRef.value?.getValue();

  if (!hostname || !ipAddress || !ttl) {
    if (!hostname) hostnameRef.value?.setError(true);
    if (!ipAddress) ipv6addressRef.value?.setError(true);
    if (!ttl) ttlRef.value?.setError(true);

    return { hasError: true, error: 'All values are mandatory.' };
  }

  return {
    data: {
      host: hostname,
      value: ipAddress,
      ttl,
    },
  };
};

events.on('process-inputs', processInputs);
</script>
