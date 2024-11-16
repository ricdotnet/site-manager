<template>
  <Stack>
    <div class="title">Login to your account</div>
    <div class="group">
      <form class="flex flex-col space-y-3" @submit="onSubmitLogin">
        <label for="email" class="flex flex-col space-y-2 relative">
          <span class="pl-4">Email Address</span>
          <Input id="email" ref="email" />
        </label>
        <ButtonGroup>
          <Button
            value="login_btn"
            name="login_btn"
            text="Login"
            color="primary"
            type="submit"
            :is-actioning="isLoggingIn"
          />
        </ButtonGroup>
      </form>
    </div>
  </Stack>

  <Dialog
    :is-open="isCodeDialogOpen"
    @on-confirm-dialog="onSubmitCode"
    :is-actioning="isVerifyingCode"
  >
    <label for="code" class="flex flex-col space-y-2 relative">
      <span class="pl-4">Code</span>
      <Input id="code" ref="code" />
    </label>
  </Dialog>
</template>

<script setup lang="ts">
import { Button, ButtonGroup, Dialog, Input, Stack } from '@components';
import { useAuth, useToaster } from '@composables';
import type { InputComponent } from '@types';
import { messages } from '@utils';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();
const { login, verifyCode } = useAuth();
const { addToast } = useToaster();

const email = ref<InputComponent>();
const code = ref<InputComponent>();

const isLoggingIn = ref(false);
const isCodeDialogOpen = ref(false);
const isVerifyingCode = ref(false);

type Error = {
  message_code: 'email_not_found';
};
const onErrorResponse = (error: Error) => {
  if (error.message_code === 'email_not_found') {
    email.value?.setError(true, messages.user[error.message_code]);
  }

  isLoggingIn.value = false;
};

const onSubmitLogin = async (evt: Event) => {
  evt.preventDefault();

  const e = email.value?.getValue();

  if (!e) {
    email.value?.setError(true);
  }

  if (e) {
    isLoggingIn.value = true;
    const { error } = await login(e);
    if (error) return onErrorResponse(error as Error);

    isLoggingIn.value = false;
    isCodeDialogOpen.value = true;
  }
};

const onSubmitCode = async () => {
  const c = code.value?.getValue();

  if (!c) {
    code.value?.setError(true);
  }

  if (c) {
    isVerifyingCode.value = true;
    const { error } = await verifyCode(c);
    if (error) return addToast('error', 'Could not verify login code.');

    isVerifyingCode.value = false;
    isCodeDialogOpen.value = false;
    await router.push('/');
  }
};
</script>

<style scoped lang="scss">
.title {
  @apply text-2xl font-bold tracking-widest mb-6;
}

.group {
  @apply w-full bg-white px-6 py-10 rounded-md shadow-md;
  @apply lg:w-[26rem] dark:bg-dark;
}
</style>
