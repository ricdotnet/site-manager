<template>
  <div class="fixed bottom-4 left-4 z-50 w-[300px] flex flex-col gap-y-2">
    <TransitionGroup name="slide-right">
      <div
        v-for="toast in toasts"
        class="py-4 px-8 rounded-md border dark:border-dark-border border-light-border bg-light dark:bg-dark-darker cursor-pointer"
        @click="removeToast(toast.id)"
        :key="toast.id"
      >
        <div class="flex gap-x-2">
          <Toast :toast="toast" />
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup lang="ts">
import { Toast } from '@components';
import { useToaster } from '@composables';

const { toasts, removeToast } = useToaster();
</script>
