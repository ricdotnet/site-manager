<template>
  <template v-if="_isOpen">
    <div class="wrapper">
      hello world
    </div>
  </template>
</template>

<script setup lang="ts">
  import { computed, watch } from "vue";

  const props = defineProps<{
    isOpen: boolean;
  }>();

  const emits = defineEmits<{
    (event: 'onCloseDialog'): void;
  }>();

  const _isOpen = computed(() => props.isOpen);

  const listenerEvent = (e: KeyboardEvent) => {
    console.log('hi');
    if (e.key === 'Escape') {
      return emits('onCloseDialog')
    }
  }

  watch(_isOpen, () => {
    if (_isOpen.value) {
      return document.addEventListener('keyup', listenerEvent);
    }

    document.removeEventListener('keyup', listenerEvent)
  });

</script>

<style scoped lang="scss">
  .wrapper {
    @apply
    absolute
    top-0
    left-0
    bg-black/50
    w-screen
    h-screen;
  }
</style>
