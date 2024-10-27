<template>
  <div class="flex flex-col gap-2">
    <div class="flex gap-2">
      <input
        class="input"
        :class="state.errorClasses"
        :id="id"
        :type="type ?? 'text'"
        :placeholder="placeholder"
        :disabled="disabled"
        :readonly="readonly"
        ref="inputRef"
        @keyup="onKeyUp"
        @dblclick="onDoubleClick"
        @focusout="onFocusOut"
      />
      <Button
        v-if="isEditing"
        name="save"
        id="save"
        color="primary"
        @click="onClickSave"
      >
        <CheckIcon class="w-5 h-5" />
      </Button>
    </div>
    <Transition name="slide-down">
      <span v-if="errorMessageRef" class="text-red-500 text-sm px-4">
        {{ errorMessageRef }}
      </span>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { type Validator, validate } from '@validators';
import { computed, onMounted, onUnmounted, reactive, ref } from 'vue';
import { Button } from '@components';
import { CheckIcon } from '@heroicons/vue/20/solid';
import { useDebounce } from '@composables';

const debounce = useDebounce();

const props = defineProps<{
  id: string;
  type?: string;
  placeholder?: string;
  errorMessage?: string;
  timeout?: number;
  validator?: Validator | (() => void);
  disabled?: boolean;
  readonly?: boolean;
  isEditing?: boolean;
  value?: string;
}>();

onMounted(() => {
  if (inputRef.value && props.value) {
    inputRef.value.value = props.value;
  }
});

onUnmounted(() => {
  clearTimeout(timeoutCounter.value);
});

const state = reactive({
  errorMessage: computed({
    get() {
      // TODO: handle this '' error message...
      // check again the computed definition (getter does not return undefined???? did before????)
      return props.errorMessage ?? '';
    },
    set(v: string) {
      errorMessageRef.value = v;
    },
  }),
  errorClasses: '',
});

const errorMessageRef = ref(state.errorMessage);

const inputRef = ref<HTMLInputElement>();
const hasError = ref(false);
const timeoutCounter = ref<number>();
const oldValue = ref('');

const errorClasses = () => {
  if (hasError.value) {
    if (timeoutCounter.value) {
      clearTimeout(timeoutCounter.value);
      timeoutCounter.value = undefined;
    }

    // @ts-expect-error: the timeout will be a number in the browser
    timeoutCounter.value = setTimeout(() => {
      setError(false);
      emits('onResetError', props.id);
    }, props.timeout ?? 20000);
    state.errorClasses = 'outline outline-offset-2 outline-2 outline-red-500';

    return;
  }
  state.errorClasses = 'outline-none focus:outline-cobalt-green';
};

const emits = defineEmits<{
  (event: 'onResetError' | 'onKeyUp', key: string): void;
  (event: 'onDoubleClick' | 'onSave'): void;
}>();

const getValue = () => {
  const value = inputRef.value?.value;
  if (props.validator && value) {
    const error = validate(props.validator, value);
    if (error) {
      state.errorMessage = error;
      return setError(true);
    }
  }
  return value;
};

const setValue = (value: string) => {
  if (!inputRef.value) return;
  inputRef.value.value = value;
};

const onKeyUp = (e: KeyboardEvent) => {
  if (e.key === 'Escape' && props.isEditing) {
    setValue(oldValue.value);
    inputRef.value?.blur();
    emits('onSave');
    e.stopPropagation();
  }

  setError(false);
  const inputValue = inputRef.value?.value;
  if (!inputValue) return;
  debounce(() => emits('onKeyUp', inputValue));
};

const onDoubleClick = () => {
  emits('onDoubleClick');
  if (!inputRef.value) return;
  oldValue.value = inputRef.value.value;
};

const onFocusOut = () => {
  const value = inputRef.value?.value;

  if (props.validator && value) {
    const error = validate(props.validator, value);
    if (error) {
      state.errorMessage = error;
      return setError(true);
    }
  }

  emits('onSave');
};

const onClickSave = () => {
  emits('onSave');
};

const setError = (bool: boolean, message?: string) => {
  hasError.value = bool ?? true;

  if (message) {
    state.errorMessage = message;
  }

  if (!bool) {
    state.errorMessage = '';
    emits('onResetError', props.id);
  }

  errorClasses();
};

defineExpose({ getValue, setValue, setError, hasError });
</script>

<style scoped lang="scss">
.input {
  @apply w-full
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
