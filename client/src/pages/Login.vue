<template>
  <Stack>
    <div class="title">Login to your account</div>
    <div class="group">
      <form class="flex flex-col space-y-3" @submit="onSubmitLogin">
        <label for="username" class="flex flex-col space-y-2 relative">
          <span class="pl-4">Username</span>
          <Input id="username" ref="username" @on-reset-error="onResetError('username')"/>
          <Transition name="slide-down">
            <span v-if="state.hasUsernameError"
                  class="text-red-500 text-sm px-4">User with that username does not exist</span>
          </Transition>
        </label>
        <label for="password" class="flex flex-col space-y-2">
          <span class="pl-4">Password</span>
          <Input id="password" ref="password" type="password" @on-reset-error="onResetError('password')"/>
          <Transition name="slide-down">
            <span v-if="state.hasPasswordError"
                  class="text-red-500 text-sm px-4">You entered the wrong password</span>
          </Transition>
        </label>
        <ButtonGroup>
          <LinkButton text="Forgot password" href="#"/>
          <Button value="login_btn"
                  name="login_btn"
                  text="Login"
                  color="primary"
                  type="submit"/>
        </ButtonGroup>
      </form>
    </div>
  </Stack>
</template>

<script setup lang="ts">
  import { reactive, ref } from "vue";
  import { Button, ButtonGroup, Input, LinkButton, Stack } from "../components";
  import { InputComponent } from "../types";
  import { useAuth } from "../composables";
  import { useRouter } from "vue-router";

  const username = ref<InputComponent>();
  const password = ref<InputComponent>();

  const router = useRouter();

  const state = reactive<{
    hasUsernameError: boolean;
    hasPasswordError: boolean;
    resetTimer: NodeJS.Timeout | undefined;
  }>({
    hasUsernameError: false,
    hasPasswordError: false,
    resetTimer: undefined,
  });

  const { login } = useAuth();

  type Error = {
    message_code: 'username_not_found' | 'incorrect_password';
  }
  const onErrorResponse = (error: Error) => {
    if (error.message_code === 'username_not_found') {
      state.hasUsernameError = true;
    }
    if (error.message_code === 'incorrect_password') {
      state.hasPasswordError = true;
    }
    state.resetTimer = setTimeout(() => {
      state.hasUsernameError = false;
      state.hasPasswordError = false;
      state.resetTimer = undefined;
    }, 15000);
  }

  const onResetError = (key: 'username' | 'password') => {
    if (state.resetTimer) {
      clearTimeout(state.resetTimer);
      state.resetTimer = undefined;
    }
    if (key === 'username') return state.hasUsernameError = false;
    if (key === 'password') return state.hasPasswordError = false;
  }

  const onSubmitLogin = async (e: Event) => {
    e.preventDefault();

    const u = username.value?.getValue();
    const p = password.value?.getValue();

    if (!u) {
      username.value?.setError();
    }

    if (!p) {
      password.value?.setError();
    }

    if (u && p) {
      const { error } = await login(u, p);
      if (error) return onErrorResponse(error);

      await router.push('/');
    }
  }
</script>

<style scoped lang="scss">
  .title {
    @apply text-2xl font-bold tracking-widest mb-6;
  }

  .group {
    @apply w-full bg-white px-6 py-10 rounded-md shadow-md;
    @apply lg:w-[26rem] dark:bg-dark;
  }

  .slide-down-enter-active,
  .slide-down-leave-active {
    transition: all 0.2s ease;
  }

  .slide-down-enter-from,
  .slide-down-leave-to {
    transform: translateY(-10px);
    opacity: 0;
  }
</style>
