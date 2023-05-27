<template>
  <div class="flex flex-col space-y-2">
    <input class="input"
           :class="errorClasses"
           :id="id"
           :type="type ?? 'text'"
           ref="inputRef"
           @keyup="setError(false)"/>
    <Transition name="slide-down">
      <span v-if="errorMessageRef"
            class="text-red-500 text-sm px-4">{{ errorMessageRef }}</span>
    </Transition>
  </div>
</template>

<script setup lang="ts">
  import { computed, onUnmounted, reactive, ref } from "vue";
  import { validate, Validator } from "../../validators";

  const props = defineProps<{
    id: string;
    type?: string;
    placeholder?: string;
    errorMessage?: string;
    timeout?: number;
    validator?: Validator | Function;
  }>();

  const state = reactive({
    errorMessage: computed({
      get() {
        return props.errorMessage;
      },
      set(v: string) {
        errorMessageRef.value = v;
      },
    }),
  });

  const errorMessageRef = ref(state.errorMessage);

  const inputRef = ref<HTMLInputElement>();
  const hasError = ref(false);
  const timeout = ref<NodeJS.Timeout>();

  console.log('hello world');

  const errorClasses = computed(() => {
    if (hasError.value) {
      if (timeout.value) {
        clearTimeout(timeout.value);
        timeout.value = undefined;
      }
      timeout.value = setTimeout(() => {
        setError(false);
        emits('onResetError', props.id);
      }, props.timeout ?? 3000);
      return 'outline outline-offset-2 outline-2 outline-red-500';
    }
    return 'outline-none focus:outline-cobalt-green'
  });

  const emits = defineEmits<{
    (event: 'onResetError', key: string): void;
  }>();

  const getValue = () => {
    const value = inputRef.value?.value;
    if (props.validator) {
      const error = validate(props.validator, value);
      if (error) {
        state.errorMessage = error;
        return setError(true);
      }
    }
    return value;
  }

  const setError = (bool: boolean, message?: string) => {
    hasError.value = bool;

    if (message) {
      state.errorMessage = message;
    }

    if (!bool) {
      state.errorMessage = '';
      emits('onResetError', props.id);
    }
  }

  onUnmounted(() => {
    clearTimeout(timeout.value);
  });

  defineExpose({ getValue, setError });
</script>

<style scoped lang="scss">
  .input {
    @apply w-full py-3 px-4 bg-white border border-light-border rounded-md transition-[outline] ease-in-out duration-200;
    @apply dark:bg-dark dark:border-dark-border;
  }

  .slide-down-enter-active,
  .slide-down-leave-active {
    transition: all 0.2s ease;
  }

  .slide-down-enter-from,
  .slide-down-leave-to {
    transform: translateY(-10px);
    opacity: 0;
  }
</style>
