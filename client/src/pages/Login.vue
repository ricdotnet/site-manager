<template>
  <Stack>
    <div class="text-2xl font-bold tracking-widest mb-6">
      Login to your account
    </div>
    <div
      class="w-full bg-white px-6 py-10 rounded-md shadow-md lg:w-[26rem] dark:bg-dark"
    >
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
    :is-cancelable="false"
    :is-actioning="isVerifyingCode"
    title="Verification code"
  >
    <label for="code" class="flex flex-col space-y-2 relative">
      <span class="pl-4">Code</span>
      <Input id="code" ref="code" />
    </label>
    <p class="pt-5 pl-4 text-lg">Check your email for the verification code.</p>
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

const onErrorResponse = (error: string) => {
  if (error === 'email_not_found') {
    email.value?.setError(true, messages.user[error]);
    addToast('error', messages.user[error]);
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
    if (error) return onErrorResponse(error);

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
    if (error) {
      isVerifyingCode.value = false;

      code.value?.setError(true);

      addToast('error', 'Could not verify login code.');
      return;
    }

    isVerifyingCode.value = false;
    isCodeDialogOpen.value = false;
    await router.push('/');
  }
};
</script>
