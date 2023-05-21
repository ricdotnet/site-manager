<template>
  <Stack>
    <div class="title">Login to your account</div>
    <div class="group">
      <form class="flex flex-col space-y-3" @submit="onSubmitLogin">
        <label for="username" class="flex flex-col space-y-2">
          <span class="pl-4">Username</span>
          <Input id="username" ref="username"/>
        </label>
        <label for="password" class="flex flex-col space-y-2">
          <span class="pl-4">Password</span>
          <Input id="password" ref="password" type="password"/>
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
  import { ref } from "vue";
  import { Button, ButtonGroup, Input, LinkButton, Stack } from "../components";
  import { InputComponent } from "../types";
  import { useAuth } from "../composables";

  const username = ref<InputComponent>();
  const password = ref<InputComponent>();

  const { login } = useAuth();

  const onSubmitLogin = (e: Event) => {
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
      return login(u, p);
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
</style>
