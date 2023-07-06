<template>
  <div class="p-5">
    <h1 class="text-2xl mb-5">VHosts config</h1>
    <div>Here you can edit your vhosts config file linked to this domain.
      When saving your site will be updated with the new configs.</div>

    <div ref="monacoRef" class="h-[400px]"></div>
    {{ getSite() }}

    <Button type="button" name="save" color="primary" text="Save" @click="onClick"/>
  </div>
</template>

<script setup lang="ts">
  import { useSitesStore } from "@stores";
  import * as monaco from 'monaco-editor';
  import { onMounted, ref } from "vue";
  import Button from "../../UI/Button.vue";

  const sitesStore = useSitesStore();
  const { getSite } = sitesStore;

  const monacoRef = ref<HTMLElement>();

  onMounted(() => {
    if (monacoRef.value) {
      monaco.editor.create(monacoRef.value, {
        value: ['function x() {', '\tconsole.log("Hello world!");', '}'].join('\n'),
        language: 'text',
        theme: 'vs-dark',
      });
    }
  });

  const onClick = () => {
    console.log(monaco.editor.getEditors()[0].getValue().split('\n'));
  }
</script>
