<template>
  <div class="flex flex-col gap-2">
    <div class="flex gap-2">
      <input class="input"
             :class="errorClasses"
             :id="id"
             :type="type ?? 'text'"
             :placeholder="placeholder"
             :disabled="disabled"
             :readonly="readonly"
             ref="inputRef"
             @keyup="onKeyUp"
             @dblclick="onDoubleClick"
             @focusout="onFocusOut"/>
      <Button v-if="isEditing" name="save" id="save" color="primary" @click="onClickSave">
        <CheckIcon class="w-5 h-5"/>
      </Button>
    </div>
    <Transition name="slide-down">
      <span v-if="errorMessageRef"
            class="text-red-500 text-sm px-4">{{ errorMessageRef }}</span>
    </Transition>
  </div>
</template>

<script setup lang="ts">
  import { computed, onUnmounted, reactive, ref } from "vue";
  import { validate, Validator } from "@validators";
  import { useDebounce } from "@composables";
  import { Button } from "@components";
  import { CheckIcon } from "@heroicons/vue/20/solid";

  const debounce = useDebounce();

  const props = defineProps<{
    id: string;
    type?: string;
    placeholder?: string;
    errorMessage?: string;
    timeout?: number;
    validator?: Validator | Function;
    disabled?: boolean;
    readonly?: boolean;
    isEditing?: boolean;
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
  const oldValue = ref('');

  const errorClasses = computed(() => {
    if (hasError.value) {
      if (timeout.value) {
        clearTimeout(timeout.value);
        timeout.value = undefined;
      }
      timeout.value = setTimeout(() => {
        setError(false);
        emits('onResetError', props.id);
      }, props.timeout ?? 20000);
      return 'outline outline-offset-2 outline-2 outline-red-500';
    }
    return 'outline-none focus:outline-cobalt-green'
  });

  const emits = defineEmits<{
    (event: 'onResetError', key: string): void;
    (event: 'onKeyUp', value: string): void;
    (event: 'onDoubleClick'): void;
    (event: 'onSave'): void;
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

  const setValue = (value: string) => {
    inputRef.value!.value = value;
  }

  const onKeyUp = (e: KeyboardEvent) => {
    if (e.key === 'Escape' && props.isEditing) {
      setValue(oldValue.value);
      inputRef.value?.blur();
      emits('onSave');
      e.stopPropagation();
    }

    setError(false);
    if (inputRef.value!.value === '') return;
    debounce(() => emits('onKeyUp', inputRef.value!.value));
  }

  const onDoubleClick = () => {
    emits('onDoubleClick');
    oldValue.value = inputRef.value!.value;
  }

  const onFocusOut = () => {
    emits('onSave');
  }

  const onClickSave = () => {
    emits('onSave');
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

  defineExpose({ getValue, setValue, setError });
</script>

<style scoped lang="scss">
  .input {
    @apply
    w-full
    py-3
    px-4
    bg-white
    border
    border-light-border
    rounded-md
    transition-[outline]
    ease-in-out
    duration-200
    dark:bg-dark dark:border-dark-border
    read-only:dark:border-dark-border/70
    read-only:border-light-border/50
    read-only:dark:text-dark-dim
    read-only:text-dark-dim/70
    read-only:outline-none
    read-only:cursor-default;
  }
</style>
