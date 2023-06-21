<template>
  <Transition name="dialog-wrapper-transition">
    <template v-if="_isOpen">
      <div class="wrapper">
        <div class="dialog dialog-box-transition">
          <h1 v-if="title" class="text-2xl text-center pt-5">{{ title }}</h1>
          <div class="dialog__content">
            <slot/>
          </div>
          <div class="dialog__footer" v-if="props.showActions">
            <Button name="cancel"
                    value="cancel"
                    color="gray"
                    :text="cancelLabel"
                    :disabled="isActioning"
                    @click="emits('onCloseDialog')"/>
            <Button name="continue"
                    value="continue"
                    color="primary"
                    :is-actioning="isActioning"
                    :text="confirmLabel"
                    @click="emits('onConfirmDialog')"/>
          </div>
        </div>
      </div>
    </template>
  </Transition>
</template>

<script setup lang="ts">
  import { computed, watch } from "vue";
  import { Button } from "@components";

  interface DialogProps {
    isOpen: boolean;
    showActions?: boolean;
    title?: string;
    cancelLabel?: string;
    confirmLabel?: string;
    isActioning?: boolean;
  }

  // @ts-ignore
  const props = withDefaults(defineProps<DialogProps>(), {
    cancelLabel: 'Cancel',
    confirmLabel: 'Confirm',
    showActions: true,
    isActioning: false,
  });

  const emits = defineEmits<{
    (event: 'onCloseDialog'): void;
    (event: 'onConfirmDialog'): void;
  }>();

  const _isOpen = computed(() => props.isOpen);

  const listenerEvent = (e: KeyboardEvent) => {
    e.preventDefault();
    if (props.isActioning) return; // do not allow actioning if an action is running
    if (e.key === 'Escape') {
      return emits('onCloseDialog')
    }
  }

  watch(_isOpen, () => {
    if (_isOpen.value) {
      document.body.style.overflow = 'hidden';
      return document.addEventListener('keyup', listenerEvent);
    }

    // @ts-ignore
    document.body.style.overflow = null;
    document.removeEventListener('keyup', listenerEvent)
  });

</script>

<style scoped lang="scss">
  .wrapper {
    @apply
    fixed
    top-0
    left-0
    bottom-0
    w-full
    h-full
    bg-black/50
    flex
    items-center
    justify-center;

    .dialog {
      @apply
      max-w-[600px]
      flex
      flex-col
      bg-white
      dark:bg-dark
      rounded-md
      border
      border-light-border
      dark:border-dark-border;
      min-width: min(500px, 90vw);
      max-height: min(400px, 80vh);

      &__content {
        @apply p-5 overflow-y-auto;
      }

      &__footer {
        @apply
        flex
        justify-between
        p-5
        border-t
        border-light-border
        dark:border-dark-border;
      }
    }
  }

  .dialog-wrapper-transition-enter-active,
  .dialog-wrapper-transition-leave-active {
    transition: opacity 0.25s ease-in-out;
  }

  .dialog-wrapper-transition-enter-from,
  .dialog-wrapper-transition-leave-to {
    opacity: 0;
  }

  .dialog-wrapper-transition-enter-active .dialog-box-transition,
  .dialog-wrapper-transition-leave-active .dialog-box-transition {
    transition: all 0.25s ease-in-out;
  }

  .dialog-wrapper-transition-enter-from .dialog-box-transition {
    transform: translateY(-30px);
  }

  .dialog-wrapper-transition-leave-to .dialog-box-transition {
    transform: translateY(-30px);
  }
</style>
