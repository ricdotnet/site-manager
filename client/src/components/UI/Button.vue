<template>
  <button
    class="rounded-md button-hover-effect flex items-center space-x-1"
    :class="computedStyles"
    :name="name"
    :value="value"
    :type="buttonType"
    :disabled="disabled || isActioning"
  >
    <slot />
    <span v-if="text">{{ text }}</span>
    <ArrowPathIcon v-if="isActioning" class="w-5 animate-spin" />
  </button>
</template>

<script setup lang="ts">
import { ArrowPathIcon } from '@heroicons/vue/20/solid';
import { ButtonColor, ButtonSize, ButtonType } from '@types';
import { computed } from 'vue';

const props = defineProps<{
  name: string;
  color: ButtonColor;
  text?: string;
  value?: string;
  disabled?: boolean;
  size?: ButtonSize;
  type?: ButtonType;
  isActioning?: boolean;
}>();

const buttonStyle = computed(() => {
  switch (props.color) {
    case 'primary':
      return 'bg-cobalt-green text-dark';
    case 'gray':
      return 'bg-dark-dim text-white';
    case 'danger':
      return 'bg-red-700 text-white';
    case 'icon':
      return 'bg-transparent text-white hover:bg-dark';
  }
});

const buttonSize = computed(() => {
  const defaultSize = 'px-4 py-2';

  switch (props.size) {
    case 'sm':
      return 'px-2 py-1 text-sm';
    case 'md':
      return defaultSize;
    case 'lg':
      return 'px-6 py-3';
    default:
      return defaultSize;
  }
});

const buttonType = computed(() => {
  return props.type ?? 'button';
});

const computedStyles = computed(() => {
  return [buttonStyle.value, buttonSize.value].join(' ');
});
</script>
