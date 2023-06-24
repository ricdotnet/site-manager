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
  import { useRoute } from "vue-router";
  import { TSite, TSIteResponse } from "@types";
  import { useRequest } from "@composables";
  import { Empty } from "@components";

  const route = useRoute();
  const site = ref<TSite>();

  const { data, error } = await useRequest<TSIteResponse>({
    endpoint: `/site/${route.params['id']}`,
    needsAuth: true,
  });

  site.value = data?.site;
</script>
