<template>
  <template v-if="error">
    <Empty message="This site does not exist"/>
  </template>
  <template v-else>
    {{ site && site.domain }}
  </template>
</template>

<script setup lang="ts">
  import { ref } from "vue";
  import { TSite, TSIteResponse } from "../../types.ts";
  import { useRequest } from "../../composables";
  import { useRoute } from "vue-router";
  import {Empty} from "../";

  const route = useRoute();
  const site = ref<TSite>();

  const { data, error } = await useRequest<TSIteResponse>({
    endpoint: `/site/${route.params['id']}`,
    needsAuth: true,
  });

  site.value = data?.site;
</script>

<style scoped lang="scss">
</style>
