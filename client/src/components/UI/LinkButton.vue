<template>
  <template v-if="href">
    <router-link :to="href" class="flex items-center px-2 hover">
      {{ text }}
    </router-link>
  </template>
  <template v-else>
    <button @click="onClick" ref="button" type="button" class="hover">{{ text }}</button>
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

<style scoped>
  .hover {
    /*@apply outline-none hover:outline hover:outline-light-dim dark:hover:outline-dark-dim;*/
    @apply dark:hover:text-light-dim hover:text-dark-dim;
    @apply ease-in-out duration-300 rounded-md;
  }
</style>
