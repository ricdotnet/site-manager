<template>
  <Dialog title="Add a Site"
          confirm-label="Add"
          :is-open="isAddingSite"
          :is-actioning="isPostingSite"
          @on-close-dialog="onCloseDialog"
          @on-confirm-dialog="onClickConfirmDialog">
    <div class="flex flex-col gap-5">
      <Input ref="domainInput"
             id="domain"
             placeholder="Domain"
             :validator="domainValidator"
             @on-key-up="onDomainKeyUp"/>
      <Input ref="configInput" id="config" placeholder="Config" :validator="configValidator"/>
      <Transition name="slide-down">
        <div v-if="generalErrorMessage" class="text-center text-red-500" v-html="generalErrorMessage"></div>
      </Transition>
    </div>
  </Dialog>
</template>

<script setup lang="ts">
  import { ref } from "vue";
  import { configValidator, domainValidator } from "@validators";
  import { Dialog, Input } from "@components";
  import { InputComponent, TSite } from "@types";
  import { useRequest } from "@composables";
  import { useSitesStore } from "@stores";
  import { messages } from "@utils";

  const sitesStore = useSitesStore();

  const props = defineProps<{
    isAddingSite: boolean;
    closeDialog: () => void;
  }>();
  const isPostingSite = ref(false);

  const domainInput = ref<InputComponent>();
  const configInput = ref<InputComponent>();

  const formHasError = ref(false);
  const generalErrorMessage = ref('');

  interface AddSiteResponse {
    code: number;
    message_code: '';
    site: TSite;
  }

  const onCloseDialog = () => {
    generalErrorMessage.value = '';
  }

  const onClickConfirmDialog = async () => {
    formHasError.value = false;

    if (!domainInput.value?.getValue()) {
      domainInput.value?.setError(true);
      formHasError.value = true;
    }

    if (!configInput.value?.getValue()) {
      configInput.value?.setError(true);
      formHasError.value = true;
    }

    if (formHasError.value) return;

    const { data, error } = await useRequest<AddSiteResponse>({
      endpoint: '/site/',
      method: 'POST',
      needsAuth: true,
      payload: {
        domain: domainInput.value?.getValue(),
        config_name: configInput.value?.getValue(),
      },
    });

    if (error) {
      generalErrorMessage.value = messages.site[error.response.data.message_code];
      return;
    }

    if (data && data.site) {
      sitesStore.addSite(data.site);
    }
    props.closeDialog();
  }

  const onDomainKeyUp = () => {
    const domainValue = domainInput.value?.getValue();

    if (domainValue) {
      configInput.value?.setValue(domainValue + '.conf');
    }
  }
</script>

<style scoped lang="scss">
</style>
