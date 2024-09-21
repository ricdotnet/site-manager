<template>
  <div class="toast-container">
    <div v-for="toast in toasts" class="toast" @click="removeToast(toast.id)">
      <div class="flex gap-x-2">
        <component :class="`toast-${toast.type || 'info'}`" :is="getIcon(toast.type)" class="w-6 h-6"/>
        <span>{{ toast.message }}</span>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useToaster } from "@composables";
import { CheckCircleIcon, ExclamationCircleIcon, InformationCircleIcon, XCircleIcon } from '@heroicons/vue/20/solid';

const { toasts, removeToast } = useToaster();

function getIcon(type: string) {
  switch (type) {
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
}
</script>

<style scoped lang="scss">
.toast-container {
  @apply fixed bottom-4 left-4 z-50 flex flex-col gap-y-2 w-[300px];
}

.toast {
  @apply py-4 px-8 rounded-md cursor-pointer border dark:border-dark-border border-light-border;
}

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