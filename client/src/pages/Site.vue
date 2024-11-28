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
import { onMounted, onUnmounted, ref } from 'vue';
import { useRoute } from 'vue-router';
import { useSitesStore } from '@stores';
import { useToaster } from '@composables';

const route = useRoute();
const { fetchSite } = useSitesStore();
const { addToast } = useToaster();

const isLoading = ref(true);
const fetchError = ref(false);

const sse = ref<EventSource | null>(null);

onMounted(async () => {
  const error = await fetchSite();
  if (error) fetchError.value = true;

  isLoading.value = false;

  const api = import.meta.env.VITE_API;
  const id = route.params.id;
  sse.value = new EventSource(`${api}/site/${id}/realtime`, {
    withCredentials: true,
  });

  sse.value.onmessage = ({ data }) => {
    addToast('success', data);
  };
});

onUnmounted(() => {
  if (sse.value) {
    sse.value.close();
  }
});
</script>
