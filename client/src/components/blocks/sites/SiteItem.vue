<template>
  <div class="p-5">
    <div class="mb-5">
      <h1 class="text-2xl mb-5">Domain Settings</h1>
      <div class="grid grid-cols-2 gap-5">
        <div>
          <Input ref="domainInputRef"
                 id="domain"
                 placeholder="Domain"
                 :validator="domainValidator"
          />
        </div>
        <div>
          <Input ref="configNameInputRef"
                 id="domain"
                 placeholder="Config name"
          />
        </div>
      </div>
    </div>

    <div class="mb-5">
      <h1 class="text-2xl mb-5">Nginx config</h1>
      <div ref="monacoRef" class="h-[500px]"></div>
    </div>

    <Button type="button" name="save" color="primary" text="Save" @click="onClick"/>
  </div>
</template>

<script setup lang="ts">
  import * as monaco from 'monaco-editor';
  import { useSitesStore } from "@stores";
  import { onMounted, ref } from "vue";
  import { domainValidator } from "@validators";
  import { InputComponent } from "@types";
  import { useRequest, useEditor } from "@composables";
  import Button from "../../UI/Button.vue";
  import Input from "../../UI/Input.vue";

  const sitesStore = useSitesStore();
  const { getSite } = sitesStore;

  const domainInputRef = ref<InputComponent>();
  const configNameInputRef = ref<InputComponent>();

  const { buildEditor, monacoRef } = useEditor();

  onMounted(async () => {
    domainInputRef?.value.setValue(getSite().domain);
    configNameInputRef?.value.setValue(getSite().config_name);

    if (monacoRef.value) {
      buildEditor(getSite().config);
    }
  });

  const onClick = async () => {
    const { error } = await useRequest<any>({
      endpoint: `/site/${getSite().id}`,
      method: 'PATCH',
      needsAuth: true,
      payload: {
        domain: domainInputRef.value?.getValue(),
        config_name: configNameInputRef.value?.getValue(),
        config: monaco.editor.getModels()[0].getValue(),
      },
    });

    if (error) {
      console.log(error);
      return;
    }
  }
</script>
