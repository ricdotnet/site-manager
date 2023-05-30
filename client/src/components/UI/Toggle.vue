<template>
  <div @click="onToggle"
       class="w-[40px] h-[25px] relative flex items-center rounded-full border border-light-dim dark:border-dark-border cursor-pointer">
    <input ref="toggle"
           type="checkbox"
           :checked="props.isChecked"
           :name="props.name"
           :id="props.id"
           :title="props.title">
    <template v-if="isThemeToggle">
      <span :class="toggledClasses"
            class="w-5 top-[3px] absolute transition-all ease-in-out duration-200">
        <MoonIcon class="w-4 relative top-[2px]" v-if="isChecked"/>
        <SunIcon class="text-dark-dim" v-else/>
      </span>
    </template>
    <template v-else>
      <span :class="toggledClasses"
            class="absolute w-[13px] h-[13px] rounded-full transition-all ease-in-out duration-200"></span>
    </template>
  </div>
</template>

<script setup lang="ts">
  import { computed } from 'vue';
  import { MoonIcon, SunIcon } from "@heroicons/vue/20/solid";

  const props = defineProps<{
    name: string;
    isChecked: boolean;
    id?: string;
    title?: string;
    isThemeToggle?: boolean;
  }>();

  const toggledClasses = computed(() => {
    if (props.isThemeToggle) {
      return props.isChecked ? 'left-[18px] text-cobalt-green' : 'left-[3px]';
    }
    return props.isChecked ? 'left-[20px] bg-cobalt-green' : 'left-[5px] bg-light-dim dark:bg-dark-dim';
  });

  const emits = defineEmits<{
    (event: 'ontoggle'): void;
  }>();

  function onToggle() {
    emits('ontoggle');
  }
</script>

<style scoped>
  input {
    width: 0;
    height: 0;
  }
</style>
