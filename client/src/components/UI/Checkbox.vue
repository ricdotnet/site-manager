<template>
  <label class="checkbox" :for="id">
    <input
      type="checkbox"
      :id="id"
      :name="name"
      class="checkbox__input"
      :checked="checked"
      @change="onChange"
      :indeterminate="indeterminate"
    />
    <span class="checkbox__checkmark">
      <template v-if="indeterminate && !checked">
        <MinusIcon class="checkbox__checkmark--check" />
      </template>
      <template v-else>
        <CheckIcon class="checkbox__checkmark--check" />
      </template>
    </span>
  </label>
</template>

<script setup lang="ts">
import { CheckIcon, MinusIcon } from '@heroicons/vue/20/solid';

defineProps<{
  id: string;
  name: string;
  checked: boolean;
  indeterminate?: boolean;
}>();

const emits = defineEmits<(event: 'onChange') => void>();

const onChange = () => {
  emits('onChange');
};
</script>

<style scoped lang="scss">
.checkbox {
  @apply flex select-none cursor-pointer;

  &__input {
    @apply w-0 h-0 absolute cursor-pointer;

    &:checked + .checkbox__checkmark {
      @apply bg-dark-border dark:bg-light-border;

      .checkbox__checkmark--check {
        @apply block text-white dark:text-black;
      }
    }

    &:indeterminate + .checkbox__checkmark {
      @apply bg-dark-border dark:bg-light-border;

      .checkbox__checkmark--check {
        @apply block text-white dark:text-black;
      }
    }

    &:focus + .checkbox__checkmark {
      @apply ring-2
        ring-dark
        dark:ring-white
        ring-offset-2
        ring-offset-light
        dark:ring-offset-dark;
    }
  }

  &__checkmark {
    @apply transition
      ease-in-out
      duration-150
      w-5
      h-5
      border
      border-dark-dim
      dark:border-light-dim
      rounded-md
      flex-shrink-0
      flex
      items-center
      justify-center;

    &:hover {
      @apply border-dark-darker dark:border-light-border;
    }

    &--check {
      @apply hidden w-3.5 h-3.5;
    }
  }
}
</style>
