<template>
  <div class="toast-container">
    <TransitionGroup name="slide-right">
      <div v-for="toast in toasts" class="toast" @click="removeToast(toast.id)" :key="toast.id">
        <div class="flex gap-x-2">
          <Toast :toast="toast"/>
        </div>
      </div>
    </TransitionGroup>
  </div>
</template>

<script setup lang="ts">
import { useToaster } from "@composables";
import { Toast } from "@components";

const { toasts, removeToast } = useToaster();
</script>

<style scoped lang="scss">
.toast-container {
  @apply fixed bottom-4 left-4 z-50 flex flex-col gap-y-2 w-[300px];
}

.toast {
  @apply py-4 px-8 rounded-md cursor-pointer border dark:border-dark-border border-light-border;
}
</style>