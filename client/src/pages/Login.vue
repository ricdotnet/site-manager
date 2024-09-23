<template>
  <Stack>
    <div class="title">Login to your account</div>
    <div class="group">
      <form class="flex flex-col space-y-3" @submit="onSubmitLogin">
        <label for="username" class="flex flex-col space-y-2 relative">
          <span class="pl-4">Username</span>
          <Input id="username"
                 ref="username"/>
        </label>
        <label for="password" class="flex flex-col space-y-2">
          <span class="pl-4">Password</span>
          <Input id="password"
                 ref="password"
                 type="password"/>
        </label>
        <ButtonGroup>
          <LinkButton text="Forgot password" href="#"/>
          <Button value="login_btn"
                  name="login_btn"
                  text="Login"
                  color="primary"
                  type="submit"
                  :is-actioning="isLoggingIn"/>
        </ButtonGroup>
      </form>
    </div>
  </Stack>
</template>

<script setup lang="ts">
import { Button, ButtonGroup, Input, LinkButton, Stack } from '@components';
import { useAuth } from '@composables';
import { InputComponent } from '@types';
import { messages } from '@utils';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const username = ref<InputComponent>();
const password = ref<InputComponent>();
const isLoggingIn = ref(false);

const router = useRouter();

const { login } = useAuth();

type Error = {
  message_code: 'username_not_found' | 'incorrect_password';
};
const onErrorResponse = (error: Error) => {
  if (error.message_code === 'username_not_found') {
    username.value?.setError(true, messages.user[error.message_code]);
  }
  if (error.message_code === 'incorrect_password') {
    password.value?.setError(true, messages.user[error.message_code]);
  }

  isLoggingIn.value = false;
};

const onSubmitLogin = async (e: Event) => {
  e.preventDefault();

  const u = username.value?.getValue();
  const p = password.value?.getValue();

  if (!u) {
    username.value?.setError(true);
  }

  if (!p) {
    password.value?.setError(true);
  }

  if (u && p) {
    isLoggingIn.value = true;
    const { error } = await login(u, p);
    if (error) return onErrorResponse(error as Error);

    isLoggingIn.value = false;
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
