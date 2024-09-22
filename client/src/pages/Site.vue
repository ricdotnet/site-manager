<template>
  <template v-if="isLoading">
    Loading...
  </template>
  <template v-else-if="fetchError">
    something went wrong...
  </template>
  <template v-else>
    <SiteInfo/>
    <SiteItem/>
  </template>
</template>

<script setup lang="ts">
import { SiteInfo, SiteItem } from '@components';
import { useSitesStore } from '@stores';
import { onMounted, ref } from 'vue';

const sitesStore = useSitesStore();

const isLoading = ref(true);
const fetchError = ref(false);

onMounted(async () => {
  const error = await sitesStore.fetchSite();
  if (error) fetchError.value = true;

  isLoading.value = false;
});
</script>
