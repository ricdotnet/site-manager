<template>
  <div ref="dropdownRef" tabindex="0" class="dropdown">
    <Button id="dropdown_button" name="dropdown" :color="color ?? 'primary'" @click="emits('onClick')">
      <span>Dropdown</span>
      <ChevronDownIcon v-if="!isOpen" class="w-5"/>
      <ChevronUpIcon v-if="isOpen" class="w-5"/>
    </Button>

    <Transition name="slide-down">
      <div ref="itemsRef" v-if="isOpen && slots.items" class="dropdown__items" :class="itemsPosition">
        <slot name="items"/>
      </div>
    </Transition>
  </div>
</template>

<script setup lang="ts">
  import { computed, ref, useSlots, watch } from "vue";
  import { Button } from "@components";
  import { ChevronDownIcon, ChevronUpIcon } from "@heroicons/vue/20/solid";
  import { ButtonColor } from "@types";

  const props = defineProps<{
    isOpen: boolean;
    color?: ButtonColor;
  }>();

  const emits = defineEmits<{
    (event: 'onClick'): void;
  }>();

  const slots = useSlots();

  const dropdownRef = ref<HTMLDivElement>();
  const itemsRef = ref<HTMLDivElement>();

  const itemsPosition = computed(() => {
    if (itemsRef.value) {
      const leftPos = itemsRef.value!.getBoundingClientRect().left;

      if (leftPos > window.innerWidth / 2) {
        return 'right-0';
      }

      return 'left-0';
    }
  });

  watch(() => props.isOpen, () => {
    if (props.isOpen) {
      return dropdownRef.value!.addEventListener('focusout', onDropdownFocusOut);
    }
    dropdownRef.value!.removeEventListener('focusout', onDropdownFocusOut);
  });

  const onDropdownFocusOut = (e: FocusEvent) => {
    e.stopPropagation();
    if (props.isOpen) {
      emits('onClick');
    }
  }
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
