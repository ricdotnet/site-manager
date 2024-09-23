<template>
  <div ref="dropdownRef" class="dropdown">
    <slot name="caller"/>
    <template v-if="!slots.caller">
      <!-- TODO find another way to handle an undefined name -->
      <Button :id="id" :name="name ?? ''" :color="color ?? 'primary'" @click="emits('onClick')">
        <span>{{ text }}</span>
        <ChevronDownIcon v-if="!isOpen" class="w-5"/>
        <ChevronUpIcon v-if="isOpen" class="w-5"/>
      </Button>
    </template>

    <Transition name="slide-down">
      <div ref="itemsRef" v-if="isOpen && slots.items" class="dropdown__items" :class="itemsPosition">
        <slot name="items"/>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
import { Button } from '@components';
import { ChevronDownIcon, ChevronUpIcon } from '@heroicons/vue/20/solid';
import { ButtonColor } from '@types';
import { computed, ref, useSlots, watch } from 'vue';

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

const itemsPosition = computed(() => {
  if (itemsRef.value) {
    const leftPos = itemsRef.value.getBoundingClientRect().left;

    if (leftPos > window.innerWidth / 2) {
      return 'right-0';
    }

    return 'left-0';
  }
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
    @apply
    absolute
    w-max
    min-w-[150px]
    max-w-[250px]
    mt-2
    p-3
    bg-light-border
    dark:bg-dark
    flex
    flex-col
    rounded-md;
  }
}
</style>
