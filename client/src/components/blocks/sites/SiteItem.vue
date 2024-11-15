<template>
  <div class="p-5">
    <div class="mb-5">
      <h1 class="text-2xl mb-5">Domain Settings</h1>
      <div class="grid grid-cols-2 gap-5">
        <div>
          <Input
            ref="domainInputRef"
            id="domain"
            placeholder="Domain"
            :validator="domainValidator"
          />
        </div>
        <div>
          <Input
            ref="configNameInputRef"
            id="domain"
            placeholder="Config name"
          />
        </div>
      </div>
    </div>

    <div class="mb-5">
      <h1 class="text-2xl mb-5">Nginx config</h1>
      <div ref="monacoRef" class="h-[500px] p-1 rounded-md bg-[#1E1E1E]"></div>
    </div>

    <div class="flex gap-x-2">
      <Button
        type="button"
        name="save"
        color="primary"
        text="Save"
        @click="onClickSave"
        :is-actioning="isSaving"
      />
      <Button
        type="button"
        name="reload-nginx"
        text="Reload nginx"
        color="bordered"
        @click="onClickReloadNginx"
        :is-actioning="isReloadingNginx"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import * as monaco from 'monaco-editor';
import { Button, Input } from '@components';
import { onMounted, onUnmounted, ref } from 'vue';
import { useEditor, useRequest, useToaster } from '@composables';
import type { InputComponent } from '@types';
import { domainValidator } from '@validators';
import { messages } from '@utils';
import { useSitesStore } from '@stores';

const { getSite, reloadNginx } = useSitesStore();

const isSaving = ref(false);
const isReloadingNginx = ref(false);

const domainInputRef = ref<InputComponent>();
const configNameInputRef = ref<InputComponent>();

const { buildEditor, monacoRef } = useEditor();
const { addToast } = useToaster();

onMounted(async () => {
  domainInputRef.value?.setValue(getSite().domain);
  configNameInputRef.value?.setValue(getSite().config_name);

  if (monacoRef.value) {
    buildEditor(getSite().config || '');
  }
});

onUnmounted(() => {
  // necessary or else the editor will not be disposed, and we end up with multiple editors
  monaco.editor.getModels()[0].dispose();
});

const onClickSave = async () => {
  isSaving.value = true;

  const { error } = await useRequest<never>({
    endpoint: `/site/`,
    method: 'PATCH',
    payload: {
      site: {
        id: getSite().id,
        domain: domainInputRef.value?.getValue(),
        config_name: configNameInputRef.value?.getValue(),
      },
      config: monaco.editor.getModels()[0].getValue(),
    },
  });

  if (error) {
    addToast('error', 'Failed to update site');
    isSaving.value = false;
    return;
  }

  addToast('success', 'Site updated successfully');
  isSaving.value = false;
};

const onClickReloadNginx = async () => {
  isReloadingNginx.value = true;
  const { error, data } = await reloadNginx();

  if (error) {
    isReloadingNginx.value = false;
    addToast('error', messages.site[error]);
    return;
  }

  isReloadingNginx.value = false;

  if (!data) return;
  addToast('success', messages.site[data]);
};
</script>
