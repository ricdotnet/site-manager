<template>
  <Loading />
</template>

<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router';
import { Loading } from '@components';
import { onMounted } from 'vue';
import { useAuth } from '@composables';

const route = useRoute();
const router = useRouter();
const { verifyCode } = useAuth();

const code = route.query.code;

onMounted(async () => {
  if (!code) {
    return router.push('/');
  }

  const { error } = await verifyCode(code as string);
  if (error) {
    return router.push('/');
  }

  router.push('/dashboard');
});
</script>
