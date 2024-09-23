<template>
  <component :class="`toast-${toast.type || 'info'}`" :is="toastIcon" class="w-6 h-6"/>
  <template v-if="toast.title">
    <div class="flex flex-col">
      <strong>{{ toast.title }}</strong>
      <span>{{ toast.message }}</span>
    </div>
  </template>
  <template v-else>
    <span>{{ toast.message }}</span>
  </template>
</template>

<script setup lang="ts">
import { CheckCircleIcon, ExclamationCircleIcon, InformationCircleIcon, XCircleIcon } from '@heroicons/vue/20/solid';
import type { TToast } from '@types';
import { computed } from 'vue';

const props = defineProps<{
  toast: TToast;
}>();

const toastIcon = computed(() => {
  switch (props.toast.type) {
    case 'info':
      return InformationCircleIcon;
    case 'success':
      return CheckCircleIcon;
    case 'warning':
      return ExclamationCircleIcon;
    case 'error':
      return XCircleIcon;
    default:
      return InformationCircleIcon;
  }
});
</script>

<style scoped lang="scss">
.toast-info {
  @apply text-blue-500;
}

.toast-success {
  @apply text-green-500;
}

.toast-warning {
  @apply text-yellow-500;
}

.toast-error {
  @apply text-red-500;
}
</style>