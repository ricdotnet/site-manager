<template>
  <div ref="dropdownRef" class="dropdown">
    <slot name="caller" />
    <template v-if="!slots.caller">
      <Button
        :id="id"
        :name="name ?? `dropdown-button-${randomName}`"
        :color="color ?? 'primary'"
        @click="emits('onClick')"
      >
        <span v-if="text">{{ text }}</span>
        <ChevronDownIcon
          class="w-5 transition-all"
          :class="isOpen ? 'rotate-180' : ''"
        />
      </Button>
    </template>

    <Transition name="slide-down">
      <div
        ref="itemsRef"
        v-if="isOpen && slots.items"
        class="dropdown__items"
        :class="itemsPosition"
      >
        <slot name="items" />
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { computed, ref, useSlots, watch } from 'vue';
import { Button } from '@components';
import type { ButtonColor } from '@types';
import { ChevronDownIcon } from '@heroicons/vue/20/solid';

const props = defineProps<{
  id?: string;
  name?: string;
  text?: string;
  color?: ButtonColor;
  isOpen: boolean;
}>();

const emits = defineEmits<(event: 'onClick') => void>();

const slots = useSlots();

const dropdownRef = ref<HTMLDivElement>();
const itemsRef = ref<HTMLDivElement>();

const randomName = computed(() => Math.random().toString(36).substring(7));

const itemsPosition = computed(() => {
  if (itemsRef.value) {
    const leftPos = itemsRef.value.getBoundingClientRect().left;

    if (leftPos > window.innerWidth / 2) {
      return 'right-0';
    }

    return 'left-0';
  }

  return null;
});

watch(
  () => props.isOpen,
  () => {
    if (!dropdownRef.value) return;
    if (props.isOpen) {
      return dropdownRef.value.addEventListener('focusout', onDropdownFocusOut);
    }
    dropdownRef.value.removeEventListener('focusout', onDropdownFocusOut);
  },
);

const onDropdownFocusOut = (e: FocusEvent) => {
  e.stopPropagation();
  if (props.isOpen) {
    emits('onClick');
  }
};
</script>

<style scoped lang="scss">
.dropdown {
  @apply relative;

  &__items {
    @apply absolute
    w-max
    min-w-[150px]
    max-w-[250px]
    mt-2
    p-3
    bg-light-lighter
    dark:bg-dark-darker
    flex
    flex-col
    rounded-md
    z-10
    border
    border-light-border
    dark:border-dark;
  }
}
</style>
