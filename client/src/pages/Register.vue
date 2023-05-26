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
          <Input type="password" id="password" ref="password"/>
        </label>
        <label for="password-confirm" class="flex flex-col space-y-2">
          <span class="pl-4">Confirm Password</span>
          <Input type="password" id="password-confirm" ref="passwordConfirm"/>
        </label>

        <ButtonGroup>
          <LinkButton text="Login" href="/login"/>
          <Button value="Register"
                  name="register"
                  text="Register"
                  color="primary"
                  type="submit"
                  :isActioning="isRegistering"/>
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
  import { useRouter } from "vue-router";

  const router = useRouter();

  const username = ref<InputComponent>();
  const email = ref<InputComponent>();
  const password = ref<InputComponent>();
  const passwordConfirm = ref<InputComponent>();

  const isRegistering = ref(false);

  const { register } = useAuth();

  const hasError = ref(false);
  const setHasError = (v: boolean = true) => {
    hasError.value = v;
  }

  const onSubmitRegister = async (e: Event) => {
    e.preventDefault();
    isRegistering.value = true;
    setHasError(false);

    const registerData: RegisterData = {
      username: username.value?.getValue() ?? '',
      email: email.value?.getValue() ?? '',
      password: password.value?.getValue() ?? '',
      password_confirm: passwordConfirm.value?.getValue() ?? '',
    };

    if (!registerData.username) {
      username.value?.setError();
      setHasError();
    }

    if (!registerData.email) {
      email.value?.setError();
      setHasError();
    }

    if (!registerData.password) {
      password.value?.setError();
      setHasError();
    }

    if (!registerData.password_confirm) {
      passwordConfirm.value?.setError();
      setHasError();
    }

    if (registerData.password !== registerData.password_confirm) {
      password.value?.setError();
      passwordConfirm.value?.setError();
      setHasError();
    }

    if ((registerData.password === registerData.password_confirm)
      && registerData.password && registerData.password_confirm) {
      password.value?.setError(false);
      passwordConfirm.value?.setError(false);
    }

    if (!!hasError.value) return isRegistering.value = false;

    const { error } = await register(registerData);
    if (error) return; // deal with some error

    isRegistering.value = false;
    await router.push('/login');
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
