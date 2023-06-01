<template>
  <template v-if="_isOpen">
    <div class="wrapper">
      <div class="dialog">
        <div class="w-full flex flex-col space-y-5">
          <Input id="domain" placeholder="your.domain.com"/>
          <Input id="config" placeholder="your.domain.com.conf"/>
        </div>
      </div>
    </div>
  </template>
</template>

<script setup lang="ts">
  import { computed, watch } from "vue";
  import { Input } from "../index.ts";

  const props = defineProps<{
    isOpen: boolean;
  }>();

  const emits = defineEmits<{
    (event: 'onCloseDialog'): void;
  }>();

  const _isOpen = computed(() => props.isOpen);

  const listenerEvent = (e: KeyboardEvent) => {
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
      min-w-[500px]
      max-w-[600px]
      bg-white
      dark:bg-dark
      overflow-auto
      p-10
      rounded-md;
      height: min(400px, 80vh);
    }
  }
</style>
