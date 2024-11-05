<template>
  <template v-if="isLoading">
    <Loading class="w-full m-auto" />
  </template>
  <template v-else-if="fetchError"> something went wrong...</template>
  <template v-else>
    <SiteInfo />
    <SiteItem />
  </template>
</template>

<script setup lang="ts">
import { Loading, SiteInfo, SiteItem } from '@components';
import { onMounted, ref } from 'vue';
import { useSitesStore } from '@stores';

const { fetchSite } = useSitesStore();

const isLoading = ref(true);
const fetchError = ref(false);

onMounted(async () => {
  const error = await fetchSite();
  if (error) fetchError.value = true;

  isLoading.value = false;
});
</script>
