<template>
  <input class="input"
         :class="errorClasses"
         :id="id"
         :type="type ?? 'text'"
         ref="inputRef"
         @keyup="setError(false)"/>
</template>

<script setup lang="ts">
  import { computed, ref } from "vue";

  defineProps<{
    id: string;
    type?: string;
    placeholder?: string;
  }>();

  const inputRef = ref<HTMLInputElement>();
  const hasError = ref(false);

  const errorClasses = computed(() => {
    if (hasError.value) {
      return 'outline outline-offset-2 outline-2 outline-red-500';
    }
    return 'outline-none focus:outline-cobalt-green'
  });

  const getValue = () => {
    return inputRef.value?.value;
  }

  const setError = (bool = true) => {
    hasError.value = bool;
  }

  defineExpose({ getValue, setError });
</script>

<style scoped lang="scss">
  .input {
    @apply py-3 px-4 bg-white border border-light-border rounded-md transition-[outline] ease-in-out duration-200;
    @apply dark:bg-dark dark:border-dark-border;
  }
</style>
