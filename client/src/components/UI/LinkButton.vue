<template>
  <template v-if="href">
    <router-link :to="href" class="link-button hover">
      {{ text }}
    </router-link>
  </template>
  <template v-else>
    <button @click="onClick" ref="button" type="button" class="link-button hover">{{ text }}</button>
  </template>
</template>

<script setup lang="ts">
  import { ref } from "vue";

  defineProps<{
    text: string,
    href?: string,
  }>();

  const button = ref<HTMLButtonElement>();

  const emits = defineEmits<{
    (event: 'onLinkClick'): void;
  }>();

  const onClick = () => {
    emits('onLinkClick');
  }
</script>

<style scoped lang="scss">
  .link-button {
    @apply
    flex
    items-center
    px-2
    text-dark-darker
    dark:text-light;
  }

  .hover {
    @apply
    dark:hover:text-light-dim
    hover:text-dark-dim
    ease-in-out
    duration-300
    rounded-md;
  }
</style>
