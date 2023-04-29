<template>
  <button class="rounded-xl button-hover-effect flex items-center space-x-1"
          :class="computedStyles"
          :name="props.name"
          :value="props.value"
          :type="buttonType"
          :disabled="props.disabled">
    {{ props.text }}
    <slot/>
  </button>
</template>

<script setup lang="ts">
  import { computed } from "vue";
  import { ButtonColor, ButtonSize, ButtonType } from "../../types.ts";

  const props = defineProps<{
    value: string;
    name: string;
    text: string;
    color: ButtonColor;
    disabled?: boolean;
    size?: ButtonSize;
    type?: ButtonType;
  }>();

  const buttonStyle = computed(() => {
    switch (props.color) {
      case 'primary':
        return 'bg-cobalt-green text-dark';
      case 'gray':
        return 'bg-dark-dim text-white';
      case 'danger':
        return 'bg-red-700 text-white';
    }
  });

  const buttonSize = computed(() => {
    const defaultSize = 'px-4 py-2';

    switch (props.size) {
      case 'sm':
        return 'px-2 py-1';
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
