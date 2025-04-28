<template>
  <div v-if="loading">
    <p>Loading...</p>
  </div>
  <div v-else>
    <h1>Docker Containers</h1>
    <p>Docker containers:</p>
    <ul>
      <li v-for="container in containers" :key="container.id">
        {{ container.id }} - {{ container.names }} [{{ container.ports }}] ({{ container.state }})
      </li>
    </ul>
  </div>
</template>

<script setup lang="ts">
  import { onMounted, ref } from 'vue';
  import { useRequest } from '@composables';

  const loading = ref(true);
  const containers = ref<Container[]>([]);

  interface Container {
    id: string;
    names: string;
    state: string;
    ports: string;
  }

  onMounted(async () => {
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const { data, error } = await useRequest<{ data: Container[] }>({
      endpoint: '/docker/containers',
    });

    if (error) {
      console.error('Error fetching containers:', error);
      return;
    }

    if (!data) {
      console.error('No data received');
      return;
    }

    console.log('Containers:', data);
    containers.value = data.data;

    loading.value = false;
  });
</script>