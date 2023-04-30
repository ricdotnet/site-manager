<template>
  <div class="flex justify-center mt-10">
    <div class="w-[26rem] px-6 py-3 rounded-md border border-light-border dark:border-dark-border">
      <div class="text-2xl font-bold tracking-widest mb-6">Login</div>
      <form class="flex flex-col space-y-3" @submit="onSubmitLogin">
        <label for="username" class="flex flex-col space-y-2">
          <span class="pl-4">Username</span>
          <input id="username" ref="username"
                 class="bg-white px-3 py-3 w-full rounded-md outline-none focus:outline-cobalt-green text-dark dark:bg-dark-border dark:text-light ease-in-out duration-200"
                 type="text"/>
        </label>
        <label for="password" class="flex flex-col space-y-2">
          <span class="pl-4">Password</span>
          <input id="password" ref="password"
                 class="bg-white px-3 py-3 w-full rounded-md outline-none focus:outline-cobalt-green text-dark dark:bg-dark-border dark:text-light ease-in-out duration-200"
                 type="password"/>
        </label>
        <div class="flex space-x-3 py-3 justify-end">
          <Button value="login_btn"
                  name="login_btn"
                  text="Login"
                  color="primary"
                  type="submit"/>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { ref } from "vue";
  import { Button } from "../components";

  const username = ref<HTMLInputElement>();
  const password = ref<HTMLInputElement>();

  function onSubmitLogin(e: Event) {
    e.preventDefault();

    console.log(username.value?.value);
    console.log(password.value?.value);
  }
</script>

<style scoped>
</style>
