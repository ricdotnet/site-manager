<template>
  <div @click="onToggle"
       class="w-[40px] h-[25px] relative flex items-center rounded-full border border-light-dim dark:border-dark-border cursor-pointer">
    <input ref="toggle"
           type="checkbox"
           :checked="isChecked"
           :name="props.name"
           :id="props.id"
           :title="props.title">
    <span :class="toggledClasses"
          class="absolute w-[13px] h-[13px] rounded-full transition-all ease-in-out duration-200"></span>
  </div>
</template>

<script setup lang="ts">
  import { computed, ref } from 'vue';

  const props = defineProps<{
    name: string;
    isChecked: boolean;
    id?: string;
    title?: string;
  }>();

  const isChecked = ref(props.isChecked);
  const toggledClasses = computed(() => {
    return isChecked.value ? 'left-[20px] bg-cobalt-green' : 'left-[5px] bg-light-dim dark:bg-dark-dim';
  });

  const emits = defineEmits<{
    (event: 'ontoggle'): void;
  }>();

  function onToggle() {
    isChecked.value = !isChecked.value;
    emits('ontoggle');
  }
</script>

<style scoped>
  input {
    width: 0;
    height: 0;
  }
</style>
