<template>
  <Stack>
    <div class="title">Register a new account</div>
    <div class="group">
      <form class="flex flex-col space-y-3" @submit="onSubmitRegister">
        <label for="username" class="flex flex-col space-y-2">
          <span class="pl-4">Username</span>
          <Input id="username" ref="username"/>
        </label>
        <label for="email" class="flex flex-col space-y-2">
          <span class="pl-4">Email</span>
          <Input id="email" ref="email"/>
        </label>
        <label for="password" class="flex flex-col space-y-2">
          <span class="pl-4">Password</span>
          <Input id="password" ref="password"/>
        </label>
        <label for="password-confirm" class="flex flex-col space-y-2">
          <span class="pl-4">Confirm Password</span>
          <Input id="password-confirm" ref="passwordConfirm"/>
        </label>

        <ButtonGroup>
          <LinkButton text="Login" href="/login"/>
          <Button value="Register"
                  name="register"
                  text="Register"
                  color="primary"
                  type="submit"/>
        </ButtonGroup>
      </form>
    </div>
  </Stack>
</template>

<script setup lang="ts">
  import { Button, ButtonGroup, Input, LinkButton, Stack } from "../components";
  import { InputComponent, RegisterData } from "../types.ts";
  import { ref } from "vue";
  import { useAuth } from "../composables";

  const username = ref<InputComponent>();
  const email = ref<InputComponent>();
  const password = ref<InputComponent>();
  const passwordConfirm = ref<InputComponent>();

  const { register } = useAuth();

  const onSubmitRegister = (e: Event) => {
    e.preventDefault();

    const registerData: RegisterData = {
      username: username.value?.getValue() ?? '',
      email: email.value?.getValue() ?? '',
      password: password.value?.getValue() ?? '',
      passwordConfirm: passwordConfirm.value?.getValue() ?? '',
    };

    // TODO: extract these validations to a validator file
    if (!registerData.username) {
      username.value?.setError();
    }

    if (!registerData.email) {
      email.value?.setError();
    }

    if (!registerData.password) {
      password.value?.setError();
    }

    if (!registerData.passwordConfirm) {
      passwordConfirm.value?.setError();
    }

    if (registerData.password !== registerData.passwordConfirm) {
      password.value?.setError();
      passwordConfirm.value?.setError();
    }

    if ((registerData.password === registerData.passwordConfirm)
    && registerData.password && registerData.passwordConfirm) {
      password.value?.setError(false);
      passwordConfirm.value?.setError(false);
    }

    register(registerData);
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
